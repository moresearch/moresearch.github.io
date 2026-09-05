---
title: "Every workflow is an FSM. Not every FSM is a workflow."
date: 2026-09-05
slug: every-workflow-is-an-fsm
summary: "The sentence 'every workflow is an FSM, but not vice-versa' is really four claims in a trench coat, and only one of them is the truism it looks like. Every workflow engine is a state machine underneath — that is an implementation fact, not a design directive. The interesting fights are all about who owns the state machine, and in Go the schools have concrete names: replay-based durable execution (the Temporal Go SDK), checkpointed or explicit-state runtimes (DBOS Transact Go, the Restate Go SDK), a Go engine that executes YAML DAGs (Argo Workflows), or a diagram that lies. Reading Maxim Fateev's 'fallacy of the graph', Temporal's own workflow-engine principles, Chris Gillum's durable-execution pitfalls, and Long Quanzheng's reply as a single conversation — then asking what each one looks like when you only use Go."
tags: workflows, state-machines, fsm, durable-execution, temporal, dbos, workflow-as-code, orchestration, agents, essay
---

> "Every workflow is an FSM — but not vice-versa."

There is a sentence that gets rewritten as a slogan at every workflow-engine talk, and it usually comes out exactly this badly. It deserves better, because it is actually several different claims wearing one trench coat, and the entire modern argument about workflow engines is the work of unpacking them. Pin the sentence down and it splits into four arrows:

- **Every workflow really is a state machine.** True, and boring — it is a fact about engines, not about your tools.
- **Not every state machine is a workflow.** True, and it is the arrow pointing somewhere interesting.
- **If it is a state machine, express it as a graph or a diagram.** False — and the best-known refutation of the past year is devoted to exactly this.
- **If it is a workflow, it must be replay-based durable execution.** False — and the people who built one replay engine, then a second engine on top of it, wrote that reply themselves.

Four recent pieces line up almost perfectly along these arrows: [Temporal's workflow-engine design principles](https://temporal.io/blog/workflow-engine-principles#workflow-as-state-machine-515), Maxim Fateev's [_The Fallacy of the Graph_](https://temporal.io/blog/the-fallacy-of-the-graph-why-your-next-workflow-should-be-code-not-a-diagram), Chris Gillum's [_Common Pitfalls with Durable Execution Frameworks_](https://blog.cgillum.tech/common-pitfalls-with-durable-execution-frameworks-like-durable-functions-or-temporal-eaf635d4a8bb), and Long Quanzheng's [_Workflow Should Be Code, but Replay-Based Durable Execution Is NOT the Only Way_](https://medium.com/@qlong/workflow-should-be-code-but-durable-execution-is-not-the-only-way-519f7682360c). Read together they agree on the first arrow, disagree with each other on the third and fourth, and barely notice the second — which is the one this post wants to spend the most time on.

## The state machine lives in the engine, not in the file

The cleanest statement of the first arrow comes from Shawn Wang's annotated summary of Temporal's [_Designing a Workflow Engine_](https://www.youtube.com/watch?v=t524U9CixZ0) talk at Systems @ Scale 2021:

> "At a high level, any workflow engine (including those that use Workflow-as-Code like Temporal, or those that use YAML/JSON/XML) are state machines. You need a state machine to define which order to run tasks and to react to external events."

[![Designing a Workflow Engine — Systems @ Scale 2021](https://img.youtube.com/vi/t524U9CixZ0/hqdefault.jpg)](https://www.youtube.com/watch?v=t524U9CixZ0&t=315)

Watch the [Workflow as State Machine section (5:15)](https://www.youtube.com/watch?v=t524U9CixZ0&t=315) and you get the actual machine, not the metaphor. The engine holds the workflow's current state in durable storage. It feeds that state to the workflow definition. The definition answers with the next commands to issue — run Task 1, arm a timer, wait for a signal. The engine dispatches them, and when a task completes, it notifies the definition and runs the loop again: state in, commands out, event back, repeat. The History component is literally "responsible for state transitions of individual workflows." A stylized version of that loop, in Go, is embarrassingly small:

```go
// drive is the core every workflow engine implements — a state machine
// whose "state" it persists, and whose "transitions" are the answers the
// workflow definition gives to the question: what should happen next?
func drive(engine *Engine, workflowID string) error {
    for {
        state, err := engine.store.Load(workflowID) // durable current state
        if err != nil {
            return err
        }
        commands, err := engine.definition.Step(state) // ask the workflow
        if err != nil || len(commands) == 0 {
            return err // no more commands: the workflow is finished
        }
        // persist state + commands + dispatched tasks in one transaction,
        // then wait for a task completion, timer firing, or external signal
        if err := engine.dispatch(workflowID, state, commands); err != nil {
            return err
        }
        event := <-engine.events(workflowID) // "Task 1 completed", "12h elapsed"...
        engine.store.Append(workflowID, event)
    }
}
```

Everything else in that talk — task queues, timer queues, transactional consistency across state, tasks, and timers, sharding, transfer queues — is the machinery that makes this loop survive process death, wait thirty days, and scale past one machine. The talk's punchline for the state-machine point is that all of that machinery is the hard part, not the transitions: "If you only remember one slide from this presentation, remember this one" — the slide about transactions across state, tasks, and timers, because that is where the race conditions live.

So the first arrow is true at a layer almost nobody is arguing about. Whether you author in YAML, in a graph, or in ordinary Go, the thing that *runs* is an event-driven stored-state machine. The FSM was never in your file format. It is in the engine, and it was there before you chose a format. That is why the format wars below are not about whether the FSM exists — they are about who has to write it, and who has to keep it honest.

## Who owns the FSM: three schools

If every workflow is a state machine, the design question stops being "is it a state machine?" and becomes **"who owns the state machine?"** Three answers are actually in production, and they tax you in different currencies.

**School one — implicit, owned by nobody: replay-based workflow-as-code.** In Go this school is the Temporal Go SDK — `go.temporal.io/sdk`, the descendant of the AWS SWF ideas that became Cadence and then Temporal, each with a Go client at its core. You write a workflow as ordinary Go code. Charge the customer, provision the account, sleep for a day, send the email — plain sequential logic:

```go
// Implicit-state style (Temporal): the FSM is emergent.
// Wherever this code is paused IS the current state; the engine rebuilds
// it by replaying the event history through the same deterministic code.
func OrderWorkflow(ctx workflow.Context, order Order) error {
    // Activity: side-effecting work runs on workers, retried, idempotent
    if err := workflow.ExecuteActivity(ctx, Charge, order).Get(ctx, nil); err != nil {
        return err
    }
    workflow.Sleep(ctx, 24*time.Hour) // durable timer: "we are paused at this line"
    return workflow.ExecuteActivity(ctx, Ship, order).Get(ctx, nil)
}
```

The state machine is not written down anywhere. The workflow's state is the event history plus where the program counter stopped, and the engine's replay machinery re-derives it on every wake-up. Nobody maintains a transition table, because there isn't one to maintain. This is the school Fateev is defending, and for control flow that branches on data — especially data produced at runtime by an LLM — it is genuinely hard to beat: loops, `if/else`, `switch`, and function calls already exist, tested, refactorable, and reviewable in your language.

The tax: replay couples your code to your history. Gillum's [pitfall list](https://blog.cgillum.tech/common-pitfalls-with-durable-execution-frameworks-like-durable-functions-or-temporal-eaf635d4a8bb) was written from the Azure Durable Functions side, but every pitfall maps one-to-one onto the Temporal Go SDK, and it is mostly a catalogue of *implicit-state* problems: workflow code must be deterministic (no native `time.Now()`, no I/O, no randomness at workflow level); activities run at-least-once, so they must be idempotent or you send the email twice; a code change that is incompatible with in-flight event histories fails replays, forcing versioning discipline (side-by-side versions on separate task queues, or version checks in code); every event is serialized into the history, so big payloads and long fan-outs inflate history and memory until you reach for continue-as-new or sub-workflows; and nothing dead-letters a poisoned workflow that can never make progress. Quanzheng's list from running Cadence/Temporal at Indeed is the same story from the operator's chair: versioning incidents, a new programming paradigm with implicit contracts, operations that look nothing like request-response, unit tests that need a `testEnv`, and cooperative multi-threading semantics most engineers cannot debug.

**School two — implicit, checkpointed: workflow-as-code without replay.** DBOS takes the "just write code" promise and a different durability mechanism: every step of a normal function is recorded as a row in Postgres, and on restart the runtime resumes from the last committed step instead of replaying the whole history from event one. No determinism sandbox, no SDK replacements for `time.Sleep` — [`dbos.Sleep`](https://docs.dbos.dev/golang/tutorials/workflows) is durable because the checkpoint is, not because the runtime can rerun your function identically.

```go
// Implicit-state checkpoint style (DBOS Transact Go): no replay, no sandbox.
// Side effects are wrapped in steps whose results are checkpointed to
// Postgres; a crash resumes after the last completed step, never before it.
func OnboardWorkflow(ctx dbos.DBOSContext, input OnboardingInput) error {
    // Ordinary I/O inside a step: the return value is checkpointed.
    chargeID, err := dbos.RunAsStep(ctx, func(ctx context.Context) (string, error) {
        return billing.Charge(ctx, input.CustomerID, input.Amount)
    }, dbos.WithStepName("charge"), dbos.WithStepMaxRetries(3))
    if err != nil {
        return err
    }
    // Durable timer: survives restarts because the checkpoint does.
    if _, err := dbos.Sleep(ctx, 48*time.Hour); err != nil {
        return err
    }
    _, err = dbos.RunAsStep(ctx, func(ctx context.Context) (string, error) {
        return billing.Finalize(ctx, chargeID) // never duplicates the charge
    }, dbos.WithStepName("finalize"), dbos.WithStepMaxRetries(5))
    return err
}
```

I compared the two schools in detail in [Temporal vs DBOS for Go](https://blog.hackspree.com/#temporal-vs-dbos-golang); the one-line version is that Temporal sells you an event-sourced platform and DBOS sells you a checkpointing library, and the difference in operational surface is the whole argument.

**School three — explicit, owned by the developer: state as code.** Quanzheng's [reply to Fateev](https://medium.com/@qlong/workflow-should-be-code-but-durable-execution-is-not-the-only-way-519f7682360c) concedes the "code, not diagrams" half and attacks the "and therefore replay" half. At Indeed he watched a JSON-DSL workflow engine run millions of workflows a day for two years and get shut down for the same reasons Fateev lists — no unit tests, logic split between JSON and Java, copy-paste versioning, no type safety — and then watched replay-based durable execution cost his team a second set of incidents: non-determinism replays that degraded the whole fleet, versioning landmines, a paradigm engineers forgot between uses. His own answer (iWF, now forked as [Dex](https://github.com/superdurable/dex)) is Java-first, but the idea — the state machine written down explicitly, in your language — has two native Go homes. The productized one is the [Restate Go SDK](https://github.com/restatedev/sdk-go): handlers are exported methods on a struct, state is explicit key-value storage scoped to a key-addressable virtual object, and every call a handler makes is journaled by the runtime so it is never repeated:

```go
// Explicit-state style (Restate Go SDK): handlers on a key-addressable
// virtual object. State is explicit K/V; durability comes from the journaled
// execution log, not from replaying your whole call stack.
type Payment struct{}

// Exclusive per object key: this handler owns the object's state while it runs.
func (Payment) AwaitAndSettle(ctx restate.ObjectContext, orderID string) error {
    // Journaled side effect: charged exactly once across retries and crashes.
    if _, err := restate.Run(ctx, func(rc restate.RunContext) (string, error) {
        return billing.Charge(rc, orderID) // HTTP or DB call lives here
    }); err != nil {
        return err
    }
    if err := restate.Sleep(ctx, 12*time.Hour); err != nil { // durable timer
        return err
    }
    restate.Set(ctx, "status", "settled") // the FSM's state, on disk
    return nil
}
```

The self-built home is the pattern this site keeps returning to — [durable daemons](https://blog.hackspree.com/#durable-daemons-execution): plain Go processes, a Postgres row per workflow, an event loop that advances explicit state, and no orchestrator beyond your own binary. Both expressions put the transitions in the developer's hands, which is the point: the FSM is visible, versionable, and unit-testable as ordinary Go instead of being hidden inside a replayable call stack. Quanzheng's framing is exactly the fourth arrow: workflow-as-code is right, and replay-based durable execution is only one engine for it.

And the ghost at this table — **school zero, explicit, owned by the diagram** — is the graph or DSL. In the Go ecosystem it survives as Kubernetes-native YAML-DAG controllers — [Argo Workflows](https://github.com/argoproj/argo-workflows) is itself a Go program, which makes it the purest Go specimen of the school: a Go engine whose users express workflows as YAML, not as Go — plus home-grown node runners and JSON/YAML DSLs. The agent-graph frameworks Fateev is writing against (LangGraph and its kin) are Python/TypeScript products with no mainstream Go equivalent, which is itself evidence: Go teams that want durable logic reach for code, not canvases. Whatever the host language, the school expresses the FSM directly as nodes and edges. This is the school Fateev is dismantling, and its failure is the subject of the next section, because it fails for a reason that none of the code schools share: **the diagram's FSM is never the FSM that runs.**

## The converses that fail, and why each fails differently

The "not vice-versa" tail of the slogan does real work only if you say *which* converse you mean, because three different ones are floating around, and they fail for three different reasons.

**"Not every FSM is a workflow" — true, and it is the point of the word *workflow*.** A checkout-flow state machine, a TCP state machine, a UI wizard, or the claims FSM from this site's [state machines vs. rule engines](https://blog.hackspree.com/#on-rule-engines-state-machines-vs-rule-engines) essay is a perfectly good FSM that lives in memory, restarts at zero, and never needed a queue or a timer. Wang's talk defines a workflow more precisely than "has states": a workflow is a *resilient program* that executes tasks and reacts to external events, including timers and timeouts. Resilience is the load-bearing word. The FSM is the head of the animal; the workflow is the animal — the machinery that keeps the head thinking while the body crashes, restarts, waits thirty days, and scales. Draw the Venn diagram honestly and workflows are a small, demanding subset of FSMs: the subset that must survive mid-flight death and resume from where it was. That is why [durable daemons](https://blog.hackspree.com/#durable-daemons-definition) — agents that persist, remember, and act — specify crash-proof execution as a *separate condition* rather than assuming the state machine gives it to you. An FSM without durability is a diagram; the durability is what makes it a workflow. ([Condition 4, the runtime, is where it actually gets implemented.](https://blog.hackspree.com/#durable-daemons-execution))

**"If it is a state machine, draw it" — false, and Fateev's [whole post](https://temporal.io/blog/the-fallacy-of-the-graph-why-your-next-workflow-should-be-code-not-a-diagram) is the refutation.** A graph cannot hold the real machine. The real machine's transitions branch on data — an `if` statement, a loop counter, an LLM's tool list — so the graph must smuggle code into node bodies and edge conditions, and the true control flow hides in those code blobs. Its data management regresses to a global untyped key-value store addressed by JSONPath strings, because there is no scope to hang types on. Its error handling cannot represent path-dependent compensation — with three conditional steps, the set of "undo what actually ran" edges already explodes past drawable; with a hundred, it is a monstrosity. And the famous benefit — "but I can see it!" — is a picture of what *might* happen, not of what did; an OpenTelemetry trace of the real run is the honest visualization, and it is generated from code, not from the drawing. Even the graph school's defenders concede the deepest point: a drawn FSM never executes itself. It always runs on an engine underneath — which is, as section one established, another state machine. The diagram is an FSM describing an FSM, and the description is worse than the thing it describes. Fateev's conclusion is that code is not merely an acceptable way to express an FSM; it is the way that stops lying to you about which machine is actually running.

**"Workflow-as-code must be replay" — false, and Quanzheng's post plus DBOS are the counterexample.** Two different durability mechanisms already ship under identical "just write code" promises: event-sourced replay (state re-derived from history by deterministic re-execution) and checkpointing (state persisted per step, resumed from the last commit). And a third axis, who writes the transitions, gives you implicit-state code (Temporal Go SDK, DBOS) and explicit-state code (Restate's virtual objects, or a hand-rolled Go daemon over Postgres). The four-way space this opens is the real terrain of the argument, and every cell is shippable from Go today — with the twist that the diagram cell is the only one where you stop writing Go:

| | Diagram / DSL | Implicit-state replay | Implicit-state checkpoint | Explicit-state code |
|---|---|---|---|---|
| **Who writes the transitions** | the author, as nodes/edges in YAML | nobody — they emerge from Go code | nobody — they emerge from Go code | the author, as handlers and keyed state |
| **Where the state lives** | on the drawing (and in the engine under it) | event history + paused code | committed step rows in Postgres | per-key K/V in the runtime, or a Postgres row |
| **How it is made durable** | borrowed from Kubernetes — not owned by the drawing | deterministic replay of event history | per-step Postgres checkpoints | journaled execution log (Restate); your own transactions if self-rolled |
| **Dynamic control flow** | cannot be drawn; leaks into code blobs | native `if`/`for`/`try` | native `if`/`for`/`try` | native code; transitions written by hand |
| **The tax you pay** | the diagram lies; YAML is data, not logic | determinism rules, versioning, history growth, at-least-once (Gillum; Quanzheng) | every side effect wrapped as a step | handlers version-locked while in-flight; side effects journaled — or you build the whole engine |
| **Go implementation** | Argo Workflows (Go engine, YAML DAG) | Temporal Go SDK | DBOS Transact Go | Restate Go SDK; or a hand-rolled Go + Postgres daemon |

The interesting structural fact: **the leftmost column pays the code schools' taxes *and* a diagram tax on top.** It gets the versioning and payload problems of whatever engine executes the graph, plus the lie that the drawing is the logic. Nobody argues the FSM is optional; they argue about who writes it, whether the runtime re-derives it or persists it, and which tax they prefer to pay. Gillum's closing judgment — the pitfalls are real and "far outweighed by the productivity benefits" — and Quanzheng's — the pitfalls were severe enough to build a second programming model — are not contradictions; they are the two edges of the same tradeoff, priced in different currencies.

## What this means when the workflow calls an LLM

The Fallacy of the Graph was written for agentic workflows, and that is where the arrows stop being academic. An LLM tool-calling agent is the extreme case of data-dependent control flow: the set and order of steps is not known at design time, because it is a function of tokens nobody has seen yet. Fateev's example — an LLM returns a tool list at runtime, and the code loops over it — is the difference between a static DAG and a program. Draw that agent as a graph and you are drawing the one shape it is guaranteed not to take. Write it as code and dynamic dispatch is a `for` loop.

The state-machine framing sharpens the agent question the same way it sharpens the workflow question. An agent run that must survive crashes, redeploys, and multi-day horizons is a workflow, which means it needs the machinery — durable state, timers, transactional event storage — not just an FSM-shaped diagram. That is the durable-daemon move, and the site has been making it from several directions: [always-on agents](https://blog.hackspree.com/#always-on-agents) that act without prompts need persistence and memory to be trustworthy; [choreography through shared state](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard) needs the shared state itself to be transactional; and [the event log, not the plan, is the agent's honest memory](https://blog.hackspree.com/#stories-from-events) — which is another way of saying the replay history or checkpoint ledger is the state machine's truth, and any diagram you drew in advance is fiction. Even the process-automation view lands here: [no single role sees the end-to-end flow](https://blog.hackspree.com/#blind-men-elephant-business-process-automation), and the only artifact that records the transitions between roles is the event log. Graphs describe intentions. Logs describe executions. Durable code is the thing that turns log back into execution.

## Key insight

**"Every workflow is an FSM" is true at exactly the layer where it is useless: the engine.** All four sources agree there, whatever else they fight about — Temporal's own design principles state it outright for every engine from workflow-as-code to YAML/XML. The arguments that matter are the converses, and they fail in three different directions that the slogan smushes together. Not every FSM is a workflow: the FSM is the head, and the workflow is the durability machinery — queues, timers, transactional state, versioning, resilience — that the head was never going to provide by itself. Not every workflow should be drawn: an explicit diagrammatic FSM is the one representation guaranteed to diverge from the machine that runs, because real transitions branch on data and the drawing cannot hold data. And not every code workflow needs replay: checkpointing (DBOS Transact Go) and explicit-state code (Restate virtual objects, or a hand-rolled Go daemon) are different mechanisms with the same "write normal Go code" promise and different taxes. So the question to ask about any workflow is never "is it a state machine?" — the answer is always yes, it is, under the hood, and the hood is not where the interesting problems are. The question is **who owns the state machine, what machinery keeps it alive across crashes, and which tax — determinism discipline, checkpoint rows, state-class boilerplate, or a diagram that lies — you are willing to pay.** Choose the owner and the durability mechanism deliberately, and the FSM takes care of itself. Choose a picture instead, and the FSM you drew will not even be the one that runs.

## References

- Wang, S. [_Workflow engine design principles with Temporal_](https://temporal.io/blog/workflow-engine-principles#workflow-as-state-machine-515) — Temporal Blog, 5 April 2021. Annotated summary of the Systems @ Scale 2021 talk _Designing a Workflow Engine_ ([video](https://www.youtube.com/watch?v=t524U9CixZ0)); see especially the [Workflow as State Machine section (5:15)](https://www.youtube.com/watch?v=t524U9CixZ0&t=315), task queues, timer queues, and the transactional-consistency slide.
- Fateev, M. [_The fallacy of the graph: why your next agentic workflow should be code, not a diagram_](https://temporal.io/blog/the-fallacy-of-the-graph-why-your-next-workflow-should-be-code-not-a-diagram) — Temporal Blog, 20 August 2025.
- Gillum, C. [_Common pitfalls with durable execution frameworks, like Durable Functions or Temporal_](https://blog.cgillum.tech/common-pitfalls-with-durable-execution-frameworks-like-durable-functions-or-temporal-eaf635d4a8bb) — blog.cgillum.tech, 27 August 2022. Written after Temporal's Replay conference by the creator of Azure Durable Functions; the pitfalls transfer one-to-one to the Temporal Go SDK.
- Quanzheng, L. [_Workflow should be code, but replay-based durable execution is NOT the only way_](https://medium.com/@qlong/workflow-should-be-code-but-durable-execution-is-not-the-only-way-519f7682360c) — Medium, 20 August 2025 (iWF is now forked as [Dex (Durable-Execution)](https://github.com/superdurable/dex), Java-first). The same explicit-state idea in Go is the Restate virtual-object model or a self-rolled durable daemon.

**The Go implementations, official docs and packages:**

- [Temporal Go SDK](https://pkg.go.dev/go.temporal.io/sdk) — implicit-state replay: deterministic workflow code, activities, durable timers.
- [DBOS Transact Go](https://github.com/dbos-inc/dbos-transact-golang) — implicit-state checkpointing: `RunAsStep` results and sleep recorded to Postgres, no replay.
- [Restate Go SDK](https://github.com/restatedev/sdk-go) — explicit-state code: virtual objects, keyed state, journaled handlers (`restate.Run`, `restate.Sleep`, `restate.Get`/`Set`).
- [Argo Workflows](https://github.com/argoproj/argo-workflows) — the diagram school's Go engine: workflows authored as YAML DAGs on Kubernetes.

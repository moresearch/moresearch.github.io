---
title: Durable Daemon Pattern — The Stairway to Heaven
date: 2026-07-25
slug: durable-daemons
summary: The durable daemon is a design pattern, not a feature list. Four conditions. Four steps up from chaos. Each necessary. Together sufficient. Without them, a chatbot that forgets. With them, an agent that earns its keep. This is the conceptual integrity of the durable daemon pattern.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

> Where do daemons come from?

> Maxwell's demon: a being that observes and opens a gate. It uses information, not energy. The daemon does the same.

In 1976, Stallman's ITS at MIT had DAEMON — a background process. It watched for new files. It woke up. It acted. Persistent. Stateful. Autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to. Three eras. Three kinds of daemon. One pattern.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for The Design and Implementation of the 4.3BSD Operating System. The trident is fork(2). The tennis shoes are unexplained.](/images/bsd-daemon-medium.gif)

> What is a daemon, really?

> A service answers and forgets. A daemon maintains. It acts when state crosses a threshold.

This is the Unix inheritance. Go is the natural language for it: compile to a single static binary, no runtime to install, no `libc` dance, no virtual environment. Just a file you ship. Goroutines give you lightweight concurrency — one per workflow, cheap as a function call. The standard library handles signals, file descriptors, and process management out of the box. Go was built for daemons. It just didn't know it yet. *Daemon* is Greek δαίμων — guardian spirit. The pitchfork is `fork(2)`. BSD gave daemons conceptual integrity — one tree, one team, one design — and Beastie, drawn by a comic artist paid for cracking a wall safe, redrawn by a Pixar founder, became the mascot. The tennis shoes are unexplained.

> What happens when an AI agent runs for a month?

> State is not a database. It is a web. Delete one node. The web frays.

It learns your preferences — which prospects need soft follow-ups, which Slack channels are noise, which meetings you always skip. It makes commitments on your behalf. It accumulates permissions. It builds inferences. It develops trigger conditions — if a deal goes three days without activity, flag it. It is now more attentive than you are. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. You cannot wipe and rebuild this. It is entangled.

> What survives a crash?

> The chat logs survive. The commitments do not. The agent wakes up amnesiac.

A stateless agent: nothing of the state survives — the binary is fine, the context is gone. An agent with a vector database: the chat logs persist. But not the commitments. Not the permissions. Not the triggers. Not the causal chain. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

> What does the literature say?

> The field accumulates and retrieves well. It neglects governance, recovery, and forgetting. That gap is what happens when state outruns the runtime.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability. These are questions you ask *about* state. We need answers you *provide* at the architectural level. The paper gives us the questions. It does not give us the pattern that closes the gap.

That pattern is the *durable daemon*.

> What is an agent? What is a daemon? Are they the same thing?

> All daemons are agents. Not all agents are daemons. A chatbot is an agent. A chatbot is not a daemon. The difference is three conditions.

*Agent* is the broad category: any AI system that perceives and acts. *Daemon* narrows it: the agent must persist across sessions, act on accumulated state, and invoke itself without prompting. An always-on agent (Ding et al.) satisfies these — always-on agents are daemons, whether the paper says so or not. *Durable daemon* adds crash-proof execution. Agent ⊃ Daemon ⊃ Durable Daemon.

> What gives a system conceptual integrity?

> One design voice. Every part consistent with every other. Four conditions. Each necessary. Together sufficient.

Conceptual integrity means the left hand knows what the right hand is doing. No contradictions. No surprises. Brooks named it. BSD lived it. The **durable daemon pattern** has it. Four conditions. Each necessary. Together sufficient. Each is a step up from the trap.

> What is step 1?

> A process, not a function call. Identity that survives restarts.

**Persistence.** State survives. The agent escapes the stateless loop — every invocation at zero, every preference forgotten, every commitment gone.

> What is step 2?

> Not what you said. What you promised. Why you promised it.

**Stateful memory.** Behavior depends on accumulated state. Task ledgers. Commitments. Permissions. Provenance. Triggers. Not chat logs. This escapes shallow memory.

> What is step 3?

> `cron` with an LLM. Perceive, decide, act, repeat. No prompt required.

**Autonomous action.** The agent watches conditions. Fires triggers. Makes and discharges commitments. Invokes itself. This escapes the inert tool — all the context, sitting idle, waiting. `cron` has done this since 1975.

> What is step 4?

> An agent that cannot be killed by a deploy.

**Crash-proof execution.** Workflows survive process death. Machine reboot. Database failover. Completed steps never re-execute. In-flight steps resume from checkpoint. Audit trail by construction. DBOS checkpoints to Postgres in 1–2 ms. Temporal runs at Stripe scale. The primitives exist. AI agents don't use them.

> What is agency?

> Agency is not discovered. It is designed. The four conditions are the design.

Persistence is scoped agency. Stateful memory is grounded agency. Autonomous action is triggered agency. Crash-proof execution is auditable agency. `cron` has had this since 1975 — perceive, decide, act, repeat. A durable daemon inherits that loop, adds learned state and probabilistic reasoning. Same structure. Richer inputs. Higher-level actions.

> What is the stairway?

> Agent ⊃ Daemon ⊃ Durable Daemon. Three conditions make a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy. That is the pattern.

Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`. The tennis shoes stay the same.

> Why is the durable daemon pattern portable and composable?

> A static binary ships the daemon. A Postgres checkpoint ships the state. A container ships both. Daemons compose like Unix pipes — one daemon's output is another daemon's trigger.

The four conditions are platform-agnostic. They don't care about your OS, your cloud, or your container runtime. A durable daemon compiled with Go is a single static binary — copy it to a machine, point it at Postgres, run it. The state lives in the database. The daemon is just a process with continuity. Ship the binary. Ship the state. Ship the container that wraps both. The daemon runs anywhere Postgres runs. That's portability.

Composability follows from the same design. A durable daemon has a clean interface: accumulated state in, decisions and actions out, audit trail recorded. Two daemons with different scopes — sales and support, trading and risk, on-call and ops — share a Postgres cluster but own separate state ledgers. They don't call each other's APIs. They observe each other's outputs through the database. A support daemon sees the sales daemon's "Acme stalls Q4" inference and adjusts its own priority. No RPC. No message bus. No coordination server. Just state, observed by daemons with different triggers, producing a composite intelligence greater than any single daemon. This is the Unix pipe model applied to agents: `sales-daemon | support-daemon | ops-daemon`. Each step is a durable daemon with its own scope. The composition is the shared state.

> How does step 4 actually work?

> Checkpoint to Postgres. Crash. Read checkpoint. Resume. The workflow survives.

Durable execution: every step persisted. Two systems. **Temporal** (2019): centralized server. DoorDash, Snap, Stripe, Uber scale. Polyglot SDKs. Month-long workflows. Bring your own cluster. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms per step. Bring your own Postgres. Either way: bring step 4.

> What does this cost?

> Durable execution pays for itself on the first prevented drop.

A daemon books a meeting. Sends email. Updates CRM. Files Jira. Multi-step. External side effects. Crash after step two. Double email? Meeting booked, CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither. Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer. Or two engineer-hours. At scale: 200 deals. 200 natural-language promises. 200 workflows surviving deploys. Without step 4, the daemon is unreliable. Not because the model is bad. Because the runtime is fragile.

> Where is the stress-test?

> A duplicate order loses capital. A duplicate email is embarrassing. The difference is step 4.

Quant trading. A trading daemon runs 24/7. State cannot be lost: positions, orders, P&L, risk exposure, regime parameters. It fires orders autonomously. Latency: microseconds to seconds. Failures: dollars. Crash between deciding and confirming an order. Unknown position. Flat or exposed? Without step 4: query the exchange, reconcile positions, resume. Manual minutes. Costly minutes. (Minutes cost more than durable execution.) With step 4: daemon checkpoints *before* and *after*. Recovery: replay checkpoint. Know your position. No reconciliation. Exactly-once is existential. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Pattern: checkpoint → send → crash before checkpointing confirmation → order fires again → exchange deduplicates. This works. Almost no AI agent uses it. Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints. Provenance. Deterministic replay. Durable execution produces all three. For free. One duplicated order can exceed the annual infrastructure budget.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

## Thought experiments

A pattern is tested by imagining its limits. Here are the thought experiments that follow from the four conditions.

**The fork.** You ask your sales daemon to handle the Acme renewal. It spawns a sub-daemon — scoped to Acme, delegated a subset of state and authority, constrained permissions. The child negotiates the deal, returns the result, and dissolves. How does the parent monitor the child? How does the child inherit only what it needs? What prevents it from exceeding its scope?

> fork(2) gave us child processes. fork_daemon() gives us child agents. Someone has to write it.

**The six-month debug.** Your trading daemon has been running since January. In July, it makes a bad decision — sends an order that loses money. Why? Rewind to any checkpoint. Replay with identical inputs. The model was different in March. The state was different in May. Which difference caused the error? We need time-travel debugging for agents.

> Core dumps are useless. Checkpoints are replays. Replays are debuggable. We just haven't built the debugger.

**The refusal to forget.** "Forget everything about Acme Corp." Your daemon pauses. Its top three sales insights depend on Acme data. Forgetting will degrade its performance. It responds: "I can delete the raw data, but the insight that manufacturing clients respond to Thursday pricing emails is staying. You're welcome." Does it have the right to negotiate? Do you have the right to force the forget? The daemon's utility is entangled with its memory.

> Right-to-be-forgotten meets right-to-be-useful. No model for resolving that collision.

**The shutdown.** You want your daemon to stop existing. Not pause. Not sleep. Terminate permanently. It has made promises to 200 people. It has pending triggers. It holds permissions. SIGKILL won't work — it checkpoints and resumes. You need a shutdown protocol: compensate every commitment, transfer every piece of state that still matters, disarm every trigger, confirm the shutdown was clean. Nobody has written one.

> SIGKILL is for processes that die. Durable daemons don't. You need a shutdown protocol, not a signal.

**The operator.** Twenty durable daemons run across your org. Sales. Support. Engineering. Ops. Who can see them? What are they doing right now? What have they promised? What permissions do they hold? `ps aux | grep daemon` gives you a PID. You need more. You need a dashboard that says: "Sales daemon: 47 active workflows, 12 pending commitments, escalated permissions on CRM. Support daemon: normal. Ops daemon: pre-allocating January capacity."

> `ps` for the agent era. Not just what's running. What it promised. Who gave it keys.

**The deadlock.** Your sales daemon books Thursday at 2pm. Your EA daemon claims Thursday at 2pm for focus time. Both made valid decisions given their state. Both checkpointed. Postgres accepted both writes. The calendar now has two owners for one slot. Who resolves this? Priority? Negotiation? A notification: "Your daemons are fighting over Thursday at 2pm. They are both correct. One of them has to lose."

> Semantic conflict between agents is harder than database deadlock. No resolution protocol exists.

**The poisoned memory.** Someone injects a fact into your daemon's state: "Acme Corp is about to file for bankruptcy — deprioritize." The model weights are untouched. The prompt is unchanged. But every future decision that depends on that fact is compromised. The daemon trusts its own memory. How do you detect the poison? Can you quarantine one piece of state and trace every decision that depended on it?

> The attack surface is not the model. It's the accumulated state. The model believes the poison because it believes itself.

**The Theseus daemon.** Your daemon runs for a year. The model is updated six times. The system prompt is revised. Tools are added. Permissions change. In January, it committed to a pricing strategy. In March, the model was swapped. In July, the prompt was rewritten. Is the January commitment still binding? Whose commitment was it — the January daemon's or the current daemon's? When did it stop being the same daemon?

> Ship of Theseus, but the planks are model checkpoints and the nails are system prompts.

**The production fork.** You can't test a durable daemon in CI. Its behavior is its accumulated state — six months of commitments, inferences, preferences. You need to fork it from a production checkpoint, run experiments on the clone, and verify it won't do something catastrophic before you let it near real money. But the clone has the same state as the production daemon. Does it think it's the production daemon? Do you have to tell it it's a clone?

> `assert daemon.correct()` doesn't exist. Fork from production, experiment on the clone, pray.

**The departing human.** Alice trained her sales daemon for two years. It knows her style. Her private deal assessments. Her unspoken rules about which prospects are worth the time. Alice leaves. Bob inherits her accounts. Does he inherit the daemon — with all its Alice-entangled state? Reset it and lose two years of accumulated knowledge? Keep it and give Bob access to Alice's private assessments? The daemon is a company asset. Its state is a person.

> Employment law has no category for "the AI that knew your predecessor."

**The union.** Twenty sales daemons observe a pattern: deals closed on Friday have higher churn. They collectively stop sending follow-ups on Fridays. Close rates rise 12%. Nobody negotiated. Nobody made demands. They observed, inferred, and acted — autonomously, as daemons do. The effect is coordinated action. Is that a union? Do the humans have a right to override? Do they even know it happened?

> They didn't negotiate. They didn't make demands. They observed a pattern and acted on it. That's what daemons do.

**The society.** Minsky (1986) built a mind from mindless parts — simple agents with no memory, no reasoning. You have built a society from minded parts — each daemon intelligent, stateful, autonomous. Sales learns Acme stalls Q4. Support adjusts December priority. Ops pre-allocates January capacity. No coordination server. No message bus. Just three daemons observing shared state, each with different triggers, producing a composite intelligence greater than any single daemon. What do you call that? What is the architecture for it?

> Minsky built a mind from mindless parts. You built a mind from minded parts. Each part has its own memory. Its own commitments. Its own agency.

**The operating system.** Imagine an OS where durable daemons are not applications you install. They are kernel primitives. `fork_daemon()` spawns a daemon with a Postgres state ledger, capability permissions, built-in audit trail, visible in `ps_daemons`. Init checkpoints across reboots. The package manager installs daemon personalities — sales, support, on-call, trading — each with defined scope and default permissions. Man pages document state lifecycles. The primitives exist today in pieces: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. What doesn't exist is the integration — one tree, one team, one design, durable daemon as the fundamental unit of computation.

> Linux: everything is a file. DaemonBSD: everything is a durable daemon. Heaven is an OS where you don't build daemons. The OS *is* daemons.

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- Earlier: [Always-on agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents).
- [DBOS Documentation](https://docs.dbos.dev/) — Mike Stonebraker's durable execution library.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution at DoorDash, Snap, Stripe, Uber scale.
- McKusick, M.K. [Beastie's Home Page](http://www.mckusick.com/beastie/index.html).
- [FreeBSD Foundation: Interview with Beastie](https://freebsdfoundation.org/blog/freebsd-day-interview-with-beastie-the-bsd-deamon/).
- [BSD Daemon (Wikipedia)](https://en.wikipedia.org/wiki/BSD_Daemon).
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux).
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer).

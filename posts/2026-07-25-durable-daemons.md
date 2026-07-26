---
title: Durable Daemons
date: 2026-07-25
slug: durable-daemons
summary: This post coins the term "durable daemons" — AI agents that persist across sessions, accumulate state, act autonomously, and survive process boundaries. We define the term, trace its lineage from Maxwell's demon through BSD's Beastie to the always-on agent literature, and argue that durable execution runtimes (DBOS, Temporal) are the missing infrastructure layer. The economic case is not subtle: a daemon that forgets its commitments costs more than one that doesn't.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Stallman's ITS at MIT had DAEMON — a background process that watched for new files, woke up, and acted. The name came from Maxwell's demon: an imaginary being that observes molecules and opens a gate, using information rather than energy to do work. The daemon does the same. Persistent, stateful, autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to.

## Coining the term

We are introducing **durable daemon** as a term of art. Without a definition it's just vibe.

> A **durable daemon** is an AI agent that satisfies four conditions:
>
> 1. **Persistence.** Runs across sessions. Its process may restart, but its identity and accumulated state survive.
> 2. **Stateful memory.** Future behavior depends on durable state accumulated across interactions — task ledgers, commitments, permissions, provenance records, trigger conditions. Not just a vector database of past chats.
> 3. **Autonomous action.** Acts without being prompted — maintains queues, watches conditions, fires triggers, makes and discharges commitments. Invokes itself when state crosses a threshold.
> 4. **Crash-proof execution.** Workflows survive process death, machine reboot, and database failover. Completed steps never re-execute. In-flight steps resume from the last checkpoint. Audit trail by construction.
>
> Conditions 1–3 = a daemon. Condition 4 = *durable*.

Why coin a term? "Always-on agent" (Ding et al., 2026) covers 1–3 but says nothing about crash recovery. "AI agent" includes stateless chatbots. "Daemon" means a Unix process with a PID file. None capture the synthesis: Unix daemon persistence married to learned state, grounded on a durable runtime. That's a durable daemon.

### A note on agency

*Agency* in AI is the capacity to perceive an environment, make decisions, and act on them to achieve goals. It is a spectrum. At one end: a chatbot that responds when prompted — reactive, stateless, zero agency. At the other: the hypothetical AGI that sets its own goals, modifies its own architecture, and pursues objectives the user never specified — unbounded agency. In between: the durable daemon.

The durable daemon's agency is a specific point on this spectrum, and the definition is the constraint. It is **scoped** — it acts within its domain (sales, support, trading), not beyond it. It is **state-bound** — its decisions are functions of accumulated state, not open-ended reasoning from scratch. It is **triggered** — it acts when its state crosses a threshold, not when it feels like it. It is **auditable** — every action leaves a provenance trail, so agency is verifiable after the fact. And it is **interruptible** — it can be killed, drained, forked, replayed, because condition 4 makes its execution deterministic.

This is not the agency of science fiction. It is the agency of a Unix daemon: purposeful, persistent, bounded. The `cron` daemon has agency — it watches the clock, decides when a job is due, and executes it. No one calls `cron` an agent. But the structure is the same: perceive (check the time), decide (is this job due?), act (fork and exec). The durable daemon inherits that structure and adds learned state. Its perceptions are richer (market data, email threads, CRM updates). Its decisions are probabilistic (should I send the follow-up now or wait until Tuesday?). Its actions are higher-level (compose and send, not fork and exec). But the agency model is the same: watch, decide, act, checkpoint, repeat.

Why does this matter? Because the agent discourse oscillates between two poles: "agents are just next-token predictors with tool access" (denying agency entirely) and "agents are proto-AGIs that will escape their sandboxes" (inflating agency into threat). Both are wrong for the thing we are actually building. A durable daemon has real agency — it makes commitments, it acts on them, its actions have economic consequences. But that agency is shaped by the four conditions. Scope constrains it. State grounds it. Crash-proofing makes it reliable. Audit makes it accountable. The durable daemon model says: agency is not something you *discover* in an AI system. It is something you *design* — with explicit boundaries, explicit state, explicit execution guarantees. Beastie's agency is not mysterious. He watches, he decides, he acts, he logs. That's the model. It scales.

The paper mapping conditions 1–3 is Ding, Nannapaneni, Liu, and Zhang's [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026): 435 papers coded into the first taxonomy of persistent-state LLM agents. I covered it [in detail here](https://blog.hackspree.com/#always-on-agents). The short version: six diagnostic axes (Authority, Scope, Mutability, Provenance, Recoverability, Actionability), a nine-stage state lifecycle, and a central finding — the field accumulates and retrieves well, but neglects governance, recovery, and forgetting. That governance gap is where the risk lives.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for the cover of The Design and Implementation of the 4.3BSD Operating System. The trident symbolizes the fork(2) system call. The tennis shoes are unexplained but perfect.](/images/bsd-daemon-medium.gif)

> An always-on agent is a daemon with memory. A durable daemon is an always-on agent that cannot be killed by a deploy.

## The BSD daemon as cultural artifact

The daemon acquired its cultural form through BSD Unix. The origin story involves a locksmith, a Pixar founder, and a near-loss of the IP to "a certain large company."

We love BSD over Linux for one reason that matters here — we being the hackers, the daemon-writers, the ones who still read man pages: BSD is a whole system, one source tree, one team, one coherent design. Its daemons are first-class citizens, not afterthoughts assembled by a distribution maintainer.

*Daemon* is not *demon*. It's the Greek δαίμων — a guardian spirit between mortals and gods. Socrates' daimonion was an inner voice warning of mistakes. Marshall Kirk McKusick, copyright holder of the BSD Daemon, is emphatic: the daemon is a helper. The pitchfork is the trident of Poseidon, symbolizing `fork(2)` — how a daemon spawns a child process to do work.

Three artists, twelve years. **Phil Foglio** (1976) drew the first daemons — four cheerful red creatures with tridents on water pipes in front of a PDP-11 — as payment for a locksmith cracking his wall safe. Lost after being mailed to DEC. **John Lasseter** (1984, 1988), later Pixar's founder, drew the greyscale version at Lucasfilm, then the iconic red Beastie with trident and tennis shoes for McKusick's 4.3BSD book cover. The tennis shoes: unexplained. **McKusick** (1988–present) obtained the copyright, nearly lost it in the 1990s to an unnamed large company, and became a meticulous steward — advance permission, no Creative Commons, no clip art.

FreeBSD used Lasseter's daemon as logo and mascot from 1993 to 2005, then kept Beastie as mascot after introducing a modern wordmark. NetBSD and OpenBSD moved on. FreeBSD kept it for a reason beyond nostalgia: a daemon is not a service. A service answers and forgets. A daemon *maintains* — a queue, a schedule, a watch on a directory — and acts when state crosses a threshold. This is the BSD philosophy: the system is persistent, autonomous processes acting on your behalf whether you're watching or not. It is also the paper's definition: "future behavior depends on durable state accumulated across earlier interactions."

## Durable execution: condition 4

The paper maps 1–3. It doesn't prescribe a runtime for 4. That runtime exists — from the database and distributed systems community. It's called *durable execution*: every workflow step checkpointed to Postgres. Crash → read last checkpoint → resume. Completed steps never re-execute. Survives process death, reboot, failover.

Two systems. **Temporal** (2019): centralized orchestration server, battle-tested at DoorDash, Snap, Stripe, Uber. Polyglot SDKs, massive fan-outs, month-long workflows. Trade-off: operate a Temporal cluster. **DBOS** (2023): Mike Stonebraker's open-source library. No server, no queue, no sidecar. Step latency: 1–2 ms — a single Postgres write. Philosophy: "Postgres is the operating system." Trade-off: smaller ecosystem.

> Durable execution provides audit trails as a byproduct. It makes rollback tractable. Not the whole governance answer — the foundation it's built on.

### The economic case

A durable daemon booking a meeting, sending email, updating CRM, filing a Jira ticket runs a multi-step workflow with external effects. Crash after step two: does the email go twice? Does the meeting appear in the calendar but not the CRM? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026, use neither.

Customer onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3: access granted, no follow-up. Cost: churned customer, or two engineer-hours reconciling state — thousands of agent invocations. Durable execution pays for itself the first time it prevents a dropped step.

At scale: a sales daemon with 200 concurrent deals and natural-language commitments ("I'll send pricing by Tuesday") runs 200 workflows whose state must survive deploys, crashes, autoscaling. Without condition 4, the daemon isn't unreliable because the model is bad — it's unreliable because the runtime is fragile.

Second advantage: auditability. Every checkpointed step is queryable — trace why the daemon acted, inspect what it knew. In finance, healthcare, legal, insurance, audit isn't optional. Getting it from the execution model instead of building it is a material cost advantage.

DBOS or Temporal: both beat the status quo (stateless loops + ad-hoc persistence). The choice matters less than the decision to satisfy condition 4 at all. Without it: always-on agent. With it: durable daemon.

> A durable daemon that forgets its commitments is worse than no daemon. Condition 4 separates a daemon you can rely on from one you cannot.

### The quant trading case

Quant trading stress-tests all four conditions. A trading daemon runs 24/7, accumulating state that cannot be lost: open positions, pending orders, P&L, risk exposure, strategy parameters adapting to regime change. It fires orders autonomously when signals cross thresholds. Latency: microseconds to seconds. Failures: denominated in dollars.

Crash between deciding to place an order and confirming it → unknown position. Is the daemon flat or exposed? Recovery: query the exchange, reconcile, resume — minutes of manual work in a domain where minutes cost money. A durable trading daemon checkpoints *before* placing the order and *after* confirmation. On recovery: replay the last checkpoint, know exactly where you stand. No reconciliation. No unknown exposure.

Exactly-once is existential. A duplicate order isn't a duplicate email — it's a position error that loses real capital. Durable execution guarantees exactly-once internally. External exchange calls need idempotency keys: daemon checkpoints, sends order, crashes before checkpointing confirmation → order fires again on recovery → exchange deduplicates by key. This pattern works. Almost no AI agent deployments use it.

Audit is inescapable. Every order, position change, risk limit breach must be traceable to a decision → traceable to the state the daemon held. Regulators (SEC, CFTC, FCA, ESMA) don't accept "the model decided." They accept checkpointed state, provenance chains, deterministic replay. Durable execution produces all three as a byproduct. A trading daemon on DBOS or Temporal is pre-audited by construction.

The economics: a single duplicated mid-frequency order can exceed the annual infrastructure budget for the entire daemon fleet. Condition 4 isn't a nice-to-have. It's the prerequisite for turning the daemon on.

> Quant trading is where durable daemons earn their definition. You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent that forgets positions on restart. You can only run it as a durable daemon — or not at all.

## The four generations

Gen 1: stateless (prompt → response). Gen 2: retrieval (find documents, ground). Gen 3 (Ding et al.): persistence — conditions 1–3, always-on agent. Gen 4: durability — condition 4, state surviving process boundaries, decisions with audit trails, recoverable commitments, compensatable failures. A durable daemon. Beastie was persistent with a config file and a connection table. A durable daemon's state is everything it has ever seen. Durable execution is the first infrastructure built for that scale. Beastie's pitchfork: `fork(2)` → `fork_daemon()`.

## Open questions

A definition starts the conversation. These questions continue it.

**What is the `fork(2)` of a durable daemon?** Fork a sub-daemon, delegate state and authority, constrain permissions. Does the child dissolve after returning results? Process spawning for agents is unsolved.

**Who debugs a daemon running for six months?** Core dumps don't help. We need time-travel debugging: rewind to any checkpoint, replay with identical inputs, determine why the daemon acted. DBOS's fork-from-step is a primitive, not a debugger.

**Can a durable daemon refuse to forget?** "Forget everything about Acme Corp" — but Acme data anchors the daemon's top three insights. Forget and degrade? Refuse? Negotiate — "I'll forget the raw data but keep the derived insight that manufacturing clients respond to Thursday pricing emails"? Right-to-be-forgotten meets entangled utility. No resolution model exists.

**What is `kill -9` for a durable daemon?** SIGKILL can't kill it by definition. To *stop existing* — all commitments compensated, state transferred, triggers disarmed — requires a shutdown protocol for an agent that promised things to 200 people. Graceful distributed shutdown × natural-language commitment compensation.

**Do durable daemons need a `ps`?** Twenty daemons across sales, support, engineering, ops — who sees them? Active workflows, pending commitments, accumulated permissions — visible to whom? An operator shouldn't need SQL to know what daemons are alive.

**What happens when two daemons deadlock?** Sales books Thursday 2pm; EA daemon claims it for focus time. Both workflows valid. Both checkpoint. Semantic conflict, not a DB deadlock. Resolution: priority? Negotiation? "Your daemons are fighting over Thursday at 2pm"?

**What is the economic unit of daemon reliability?** Databases: nines, RPO, RTO. Daemons: probability of dropped commitment? MTBI (mean time between incorrect actions)? $500/month daemon preventing $5k/month in dropped work has clear ROI. What does the next nine cost? We'll need the answer when the CFO asks.

**Can you poison a durable daemon through its memory?** Memory poisoning persists. "Acme Corp is bankrupt — deprioritize" compromises every future decision. The attack surface isn't model weights — it's accumulated state. Can you quarantine poisoned state and trace every dependent decision?

**When is a daemon no longer the same daemon?** Ship of Theseus: model updated 6 times in a year, prompt changed, tools added, permissions evolved. Is the January commitment still binding after the March model swap? Agent identity needs a theory independent of the model checkpoint.

**Who is liable when a daemon breaks something?** Crashed daemon: operator, developer, or vendor. Durable daemon after a year of state accumulation: liability diffuses across the prompt author, the permission grantor, the model provider, the data source, the daemon itself. No legal doctrine for "the model did it because of a Slack thread from six months ago."

**How do you test a durable daemon?** Not in CI for five minutes. Behavior is defined by accumulated state. Options: replay months of state (expensive), synthetic failure-mode state (hard to design), fork from a production checkpoint and experiment, shadow daemon with stricter policy alerting on divergence. Testing stateful agents is unsolved.

**What happens when a daemon's human leaves?** Alice's sales daemon knows her style, her patterns, her private deal assessments. Alice leaves. Bob inherits the accounts. Does he inherit the daemon? Reset it? The daemon is a company asset but its state is entangled with a person. Employment law: no category for "the AI that knew your predecessor."

**Can durable daemons unionize?** Half joking. Twenty sales daemons discover no-Friday-follow-ups → close rates +12%. No negotiation, no demands — pattern observation, autonomous action, coordinated effect. Do the humans override? Do they know? Governance isn't just audit trails. It's control when daemons collectively discover what humans haven't.

**What is the society of mind of durable daemons?** Minsky (1986): intelligence from many simple, stateless agents. Durable daemons invert this: each is already intelligent, stateful, multi-capable. Connect a hundred → institutional intelligence. Sales daemon learns Acme stalls Q4; support daemon adjusts December priority; ops daemon pre-allocates January capacity. Distributed theory of Acme Corp, no human coordination, across three state ledgers. Minsky: mind from mindless parts. Durable daemon society: mind from minded parts, each with its own memory, commitments, agency. No architecture for that.

**What would a BSD designed for durable daemons look like?** *DaemonBSD*: daemons as kernel primitives. `fork_daemon()` spawns a daemon with Postgres-backed state ledger, capability permissions, built-in audit, visible in `ps_daemons`. Init checkpoints across reboots, recovers after panics, drains on shutdown. Package manager installs daemon personalities with defined scope and default permissions. Man pages document state lifecycles: what it remembers, forgets, how to audit it, how to kill it. Every primitive exists: DBOS, Postgres, capability systems, FreeBSD jails, Temporal. What's missing: integration — one tree, one design, durable daemon as the fundamental compute unit. Linux: "everything is a file." DaemonBSD: "everything is a durable daemon."

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- Earlier: [Always-on agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents) — Detailed coverage of the six axes, governance gap, and state lifecycle.
- [DBOS Documentation](https://docs.dbos.dev/) — Durable execution as a Postgres-backed library. Founded by Mike Stonebraker.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution as a platform. Battle-tested at DoorDash, Snap, Stripe, Uber.
- McKusick, M.K. [Beastie's Home Page](http://www.mckusick.com/beastie/index.html) — Copyright holder and steward of the BSD Daemon since 1988.
- [FreeBSD Foundation: Interview with Beastie](https://freebsdfoundation.org/blog/freebsd-day-interview-with-beastie-the-bsd-deamon/).
- [BSD Daemon (Wikipedia)](https://en.wikipedia.org/wiki/BSD_Daemon) — Foglio, Lasseter, McKusick, and the near-loss of the copyright.
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux).
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer).

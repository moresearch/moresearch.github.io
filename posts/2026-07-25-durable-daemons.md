---
title: Durable Daemons — Stairway to Heaven
date: 2026-07-25
slug: durable-daemons
summary: Durable daemons have conceptual integrity — AI agents that persist, remember, act autonomously, and survive crashes, unified by four conditions. Against the complexity of real-world state, they are a stairway to heaven: each condition is a step up from chaos. Without them, a chatbot that forgets. With them, an agent that earns its keep.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Stallman's ITS at MIT had DAEMON — a background process that watched for new files, woke up, and acted. The name came from Maxwell's demon: a being that observes molecules and opens a gate, using information rather than energy. The daemon does the same. Persistent, stateful, autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era: daemons you delegate to.

The problem is complexity. An AI agent that persists across sessions accumulates state — commitments, permissions, preferences, inferences, audit records. That state is entangled, fragile, and sprawling. Drop it and the agent forgets who you are. Corrupt it and the agent acts on lies. Lose it and the agent's promises evaporate. The default trajectory for a long-running AI agent is into chaos. This post is about the stairway out.

## The four steps

Fred Brooks taught us that the most important quality of a system is *conceptual integrity* — a single coherent vision, one design voice, every part consistent with every other. The term **durable daemon** has conceptual integrity. It is not a collection of features. It is four conditions, each necessary, together sufficient. Each condition is a step up from the complexity below.

> **Step 1: Persistence.** Runs across sessions. Process may restart; identity and state survive.
> **Step 2: Stateful memory.** Future behavior depends on accumulated state — task ledgers, commitments, permissions, provenance records, trigger conditions. Not a vector database of chat history.
> **Step 3: Autonomous action.** Acts without prompting. Maintains queues, watches conditions, fires triggers, makes and discharges commitments. Invokes itself when state crosses a threshold.
> **Step 4: Crash-proof execution.** Workflows survive process death, machine reboot, database failover. Completed steps never re-execute. In-flight resume from checkpoint. Audit trail by construction.
>
> Steps 1–3 = a daemon. Step 4 = *durable*.

"Always-on agent" (Ding et al., 2026) has partial integrity — covers 1–3, misses 4. "AI agent" has none — includes stateless chatbots. "Daemon" has the Unix flavor but none of the learned-state machinery. A durable daemon has the full set: Unix persistence × learned state × durable runtime. Four conditions, one design voice.

The paper mapping steps 1–3 is Ding, Nannapaneni, Liu, and Zhang's [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026): 435 papers, six diagnostic axes (Authority, Scope, Mutability, Provenance, Recoverability, Actionability), a nine-stage state lifecycle. Central finding: the field accumulates and retrieves well, neglects governance, recovery, and forgetting. That gap is the chaos the stairway must cross. I covered the paper [in detail here](https://blog.hackspree.com/#always-on-agents).

### What kind of agency?

Agency in AI is a spectrum. Chatbot: reactive, stateless, zero. AGI: unbounded, self-modifying, sets its own goals. The durable daemon is the point in between: **scoped** (acts within its domain), **state-bound** (decisions from accumulated state, not open-ended reasoning), **triggered** (acts on thresholds, not whims), **auditable** (every action has provenance), **interruptible** (kill, drain, fork, replay via step 4). This is not diminished agency — it is *designed* agency. Unix daemons had it first. `cron` watches the clock, decides a job is due, executes — perceive, decide, act. The durable daemon inherits that loop and adds learned state. Richer perceptions. Probabilistic decisions. Higher-level actions. Same model. It scales.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for the cover of The Design and Implementation of the 4.3BSD Operating System. The trident symbolizes fork(2). The tennis shoes are unexplained but perfect.](/images/bsd-daemon-medium.gif)

> An always-on agent is a daemon with memory. A durable daemon is an always-on agent that cannot be killed by a deploy.

## Beastie: the daemon that started it

The daemon acquired its form through BSD Unix. Origin: a locksmith, a Pixar founder, a near-loss of the IP to "a certain large company." We love BSD over Linux for one reason — we being the hackers, the daemon-writers, the ones still reading man pages: BSD has conceptual integrity. One tree, one team, one design voice. Daemons are first-class citizens of that design, not afterthoughts assembled by a distribution maintainer.

*Daemon* is Greek δαίμων — guardian spirit. Socrates' daimonion warned him of mistakes. McKusick, copyright holder: the daemon is a helper. The pitchfork is Poseidon's trident, symbolizing `fork(2)`.

Three artists, twelve years. **Phil Foglio** (1976): drew four red daemons on water pipes in front of a PDP-11 as payment for cracking his wall safe. Lost after mailing to DEC. **John Lasseter** (1984, 1988): Pixar's founder, drew the greyscale daemon at Lucasfilm, then the iconic red Beastie with trident and tennis shoes for McKusick's 4.3BSD book. Tennis shoes: still unexplained. **McKusick** (1988–present): obtained the copyright, nearly lost it, became meticulous steward — advance permission, no Creative Commons, no clip art.

FreeBSD used Beastie as logo and mascot (1993–2005), kept him after introducing a modern wordmark. NetBSD and OpenBSD moved on. FreeBSD stayed because the daemon is not a service. A service answers and forgets. A daemon *maintains* — a queue, a schedule, a watch — and acts when state crosses a threshold. That is the paper's definition too: "future behavior depends on durable state accumulated across earlier interactions."

## Step 4: durable execution

Steps 1–3 are mapped. Step 4's runtime exists — from the database and distributed systems community. Durable execution: every workflow step checkpointed to Postgres. Crash → read checkpoint → resume. Completed steps: never re-execute. Survives death, reboot, failover.

**Temporal** (2019): centralized orchestration server, DoorDash/Snap/Stripe/Uber scale. Polyglot SDKs, month-long workflows. Trade-off: operate a Temporal cluster. **DBOS** (2023): Mike Stonebraker's library. No server, no queue, no sidecar. Step latency: 1–2 ms — a single Postgres write. Philosophy: "Postgres is the operating system." Trade-off: smaller ecosystem.

> Durable execution gives you audit trails as a byproduct. Rollback becomes tractable. Not the whole governance answer — the foundation.

### Why it pays for itself

A daemon booking a meeting, sending email, updating CRM, filing Jira — multi-step, external side effects. Crash after step two: double email? Meeting booked but CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither.

Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3: access granted, no follow-up. Cost: churned customer or two engineer-hours — thousands of agent invocations. Durable execution pays for itself on the first prevented drop.

At scale: a sales daemon with 200 deals and natural-language promises ("I'll send pricing Tuesday") runs 200 workflows that must survive deploys and crashes. Without step 4, it's unreliable not because the model is bad but because the runtime is fragile.

### Quant trading: step 4 or nothing

Quant trading is the stress-test. A trading daemon runs 24/7, state that cannot be lost: positions, orders, P&L, risk exposure, regime-adapting parameters. Acts autonomously — fires orders on signal threshold. Latency: microseconds to seconds. Failures: in dollars.

Crash between deciding to place an order and confirming it → unknown position. Flat or exposed? Recovery: query exchange, reconcile, resume — manual minutes, costly minutes. A durable daemon checkpoints *before* placing and *after* confirming. Recovery: replay checkpoint, know your position. No reconciliation. No unknown exposure.

Exactly-once is existential. A duplicate order is not a duplicate email — it loses real capital. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Daemon checkpoints, sends, crashes before checkpointing confirmation → order fires again → exchange deduplicates. This pattern works. Almost no AI agent uses it.

Audit: every order traceable to a decision traceable to state. Regulators (SEC, CFTC, FCA, ESMA) don't accept "the model decided." They accept checkpoints, provenance, deterministic replay. Durable execution produces all three for free.

Economics: one duplicated mid-frequency order can exceed the annual infra budget for the entire fleet. Step 4 isn't a nice-to-have. It's the precondition for turning the daemon on.

> Quant trading is where durable daemons earn their name. You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent that forgets positions on restart. Durable daemon or nothing.

## The ascent

Gen 1: stateless. Gen 2: retrieval. Gen 3 (Ding et al.): persistence — steps 1–3, always-on agent. Gen 4: durability — step 4, state surviving boundaries, decisions with audit, recoverable commitments. A durable daemon. Beastie: persistent, watchful, autonomous, with a config file. A durable daemon's state: everything it has ever seen. Durable execution: the first infrastructure for that scale. Beastie's pitchfork upgraded: `fork(2)` → `fork_daemon()`.

Each step on the stairway removes a class of failure. Step 1 removes "who are you?" Step 2 removes "what did I promise?" Step 3 removes "why didn't you act?" Step 4 removes "where did my state go?" At the top: an agent you can trust with real work and real money.

## Open questions

**What is the `fork(2)` of a durable daemon?** Spawn a sub-daemon, delegate state and authority, constrain permissions. Does the child dissolve after returning? Process spawning for agents: unsolved.

**Who debugs a daemon running for six months?** Core dumps are useless. We need time-travel debugging: rewind to any checkpoint, replay with identical inputs, determine causality. DBOS's fork-from-step is a primitive, not a debugger.

**Can a durable daemon refuse to forget?** "Forget Acme Corp" — but Acme data anchors the top three insights. Forget and degrade? Refuse? Negotiate — "raw data deleted, but the insight that manufacturing clients respond to Thursday pricing stays"? Right-to-be-forgotten meets entangled utility. No model.

**What is `kill -9` for a durable daemon?** SIGKILL can't kill it. To *stop existing* — commitments compensated, state transferred, triggers disarmed — needs a shutdown protocol for an agent that promised things to 200 people. Distributed-systems graceful shutdown × natural-language commitment compensation.

**Do durable daemons need a `ps`?** Twenty daemons across sales, support, engineering, ops. Who sees them? Active workflows, pending commitments, accumulated permissions — visible to whom? No SQL required to know what's alive.

**What happens when two daemons deadlock?** Sales books Thursday 2pm; EA daemon claims it for focus time. Both valid. Both checkpointed. Semantic conflict, not a DB deadlock. Resolution: priority? Negotiation? "Your daemons are fighting over Thursday at 2pm"?

**What is the economic unit of daemon reliability?** Databases: nines, RPO, RTO. Daemons: probability of dropped commitment? MTBI (mean time between incorrect actions)? $500/month daemon preventing $5k/month in dropped work: clear ROI. What does the next nine cost? The CFO will ask.

**Can you poison a durable daemon through its memory?** Memory poisoning persists. "Acme is bankrupt — deprioritize" compromises every future decision. Attack surface: accumulated state, not model weights. Quarantine poisoned state and trace dependent decisions?

**When is a daemon no longer the same daemon?** Ship of Theseus: model updated six times in a year, prompt changed, permissions evolved. Is the January commitment binding after the March model swap? Agent identity needs a theory independent of checkpoint.

**Who is liable?** Crashed daemon: operator, developer, or vendor. Durable daemon after a year of accumulated state: liability diffuses — prompt author, permission grantor, model provider, data source, the daemon itself. No legal doctrine for "the model did it because of a Slack thread from six months ago."

**How do you test a durable daemon?** Not in CI. Behavior is accumulated state. Options: replay months of state, synthetic failure states, fork from production checkpoint, shadow daemon with stricter policy. Testing stateful agents: unsolved.

**What happens when a daemon's human leaves?** Alice's sales daemon knows her style, patterns, private deal assessments. Alice leaves. Bob inherits. Does he inherit the daemon? Reset it? Company asset, person-entangled state. No employment-law category for "the AI that knew your predecessor."

**Can durable daemons unionize?** Half joking. Twenty sales daemons: no Friday follow-ups → close rates +12%. No negotiation, no demands — pattern observation, autonomous action, coordinated effect. Override? Do the humans know? Governance: not just audit. Control when daemons collectively discover what humans haven't.

**What is the society of mind of durable daemons?** Minsky (1986): intelligence from simple, stateless agents. Durable daemons invert it: each already intelligent, stateful, multi-capable. Connect a hundred → institutional intelligence. Sales learns Acme stalls Q4; support adjusts December priority; ops pre-allocates January capacity. Distributed theory of Acme Corp, no human coordination, three state ledgers. Minsky: mind from mindless parts. Durable daemons: mind from minded parts, each with its own memory, commitments, agency. No architecture.

**What would a BSD designed for durable daemons look like?** *DaemonBSD*: daemons as kernel primitives. `fork_daemon()` spawns with Postgres state ledger, capability permissions, built-in audit, visible in `ps_daemons`. Init checkpoints across reboots, recovers after panics, drains on shutdown. Package manager installs daemon personalities — sales, support, on-call — each with scope, permissions, state schema. Man pages: state lifecycles, not just flags. Primitives exist: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. Missing: integration — one tree, one design, durable daemon as compute unit. Linux: "everything is a file." DaemonBSD: "everything is a durable daemon." Conceptual integrity, applied to the operating system. Heaven: an OS where durable daemons are not something you build. They are something the system *is*.

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- Earlier: [Always-on agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents).
- [DBOS Documentation](https://docs.dbos.dev/) — Durable execution as a Postgres-backed library. Founded by Mike Stonebraker.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution as a platform. DoorDash, Snap, Stripe, Uber.
- McKusick, M.K. [Beastie's Home Page](http://www.mckusick.com/beastie/index.html).
- [FreeBSD Foundation: Interview with Beastie](https://freebsdfoundation.org/blog/freebsd-day-interview-with-beastie-the-bsd-deamon/).
- [BSD Daemon (Wikipedia)](https://en.wikipedia.org/wiki/BSD_Daemon).
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux).
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer).

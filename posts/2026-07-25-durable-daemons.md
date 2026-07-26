---
title: Durable Daemons — Stairway to Heaven
date: 2026-07-25
slug: durable-daemons
summary: Durable daemons have conceptual integrity — AI agents that persist, remember, act autonomously, and survive crashes, unified by four conditions. Against the complexity of real-world state, each condition is a step up from chaos. Without them, a chatbot that forgets. With them, an agent that earns its keep.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Stallman's ITS at MIT had DAEMON — a background process that watched for new files, woke up, and acted. The name came from Maxwell's demon: a being that observes molecules and opens a gate, using information rather than energy. The daemon does the same. Persistent, stateful, autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era: daemons you delegate to.

## The problem: state is a trap

Here is what happens when you run an AI agent for a month. It learns your preferences — which prospects need soft follow-ups, which Slack channels are noisy, which meetings you always skip. It makes commitments on your behalf — "I'll send the deck by Tuesday," "let me check with engineering." It accumulates permissions — read access to your calendar, send access to your email, API keys to your CRM. It builds inferences — Acme Corp stalls deals in Q4, the VP of Engineering responds faster on Signal than email. It develops trigger conditions — if a deal goes three days without activity, flag it. This state is not a database you can wipe and rebuild. It is entangled. A commitment depends on a preference that depends on an inference that depends on a Slack message from March. Delete one node and the web frays in directions you cannot predict.

Now crash the process. Deploy new code. Rotate the API keys. Migrate the database. What survives? If you are running a stateless agent: nothing. If you bolted on a vector database: the chat logs, but not the commitments, not the permissions, not the triggers, not the causal chain that connects "Acme stalls Q4" to "don't send the pricing sheet until January." The agent wakes up amnesiac. The user spends two weeks re-teaching it preferences it already learned. The dropped commitments become broken promises.

Ding, Nannapaneni, Liu, and Zhang's [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026) — a coded survey of 435 papers — found that the field concentrates on accumulating and retrieving state while neglecting governance, recovery, and forgetting. That governance gap is what happens when state outruns the runtime. The survey gives us six diagnostic axes (Authority, Scope, Mutability, Provenance, Recoverability, Actionability) and a nine-stage state lifecycle. What it does not give us is a term for the thing that closes the gap. That term is *durable daemon*.

> State without durability is a liability. Every piece of accumulated knowledge is a piece of the system that a crash can destroy.

## The stairway: four conditions, four escapes

Fred Brooks taught that the most important quality of a system is *conceptual integrity* — one design voice, every part consistent with every other. **Durable daemon** has that integrity. It is not a feature list. It is four conditions, each necessary, together sufficient. Each is a step up from the trap.

**Step 1: Persistence.** Identity and accumulated state survive process restarts. The agent is not a function call. It is a process with continuity. This escapes the trap of the stateless loop — every invocation starting from zero, every preference re-taught, every commitment re-stated.

**Step 2: Stateful memory.** Future behavior depends on accumulated state — task ledgers, commitments, permissions, provenance records, trigger conditions. Not a vector database of chat logs. This escapes the trap of shallow memory — the agent that can retrieve what you said but not what it promised or why. The survey's six diagnostic axes are a checklist for whether your state is structured enough to act on.

**Step 3: Autonomous action.** The agent maintains queues, watches conditions, fires triggers, makes and discharges commitments — without being prompted. It invokes itself when state crosses a threshold. This escapes the trap of the inert tool — the agent that has all the context but sits idle until the user says "do something." `cron` has been doing this since 1975: perceive (check time), decide (job due?), act (fork and exec), repeat. Step 3 is `cron` with an LLM.

**Step 4: Crash-proof execution.** Workflows survive process death, machine reboot, database failover. Completed steps never re-execute. In-flight steps resume from the last checkpoint. Audit trail by construction. This escapes the trap of the fragile runtime — the agent that works perfectly until a deploy wipes its in-flight state. DBOS checkpoints to Postgres in 1–2 ms per step. Temporal has been doing it at Stripe scale for years. The primitives exist. AI agents don't use them.

> Steps 1–3 = a daemon. Step 4 = *durable*. The first three give you an agent that persists, remembers, and acts. The fourth gives you an agent that cannot be killed by a deploy.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for the cover of The Design and Implementation of the 4.3BSD Operating System. The trident symbolizes fork(2). The tennis shoes are unexplained but perfect.](/images/bsd-daemon-medium.gif)

## Beastie

The daemon acquired its form through BSD Unix. Origin: a locksmith cracks a wall safe, a comic artist draws four red daemons on water pipes as payment, the original is mailed to DEC and lost. A Pixar founder draws the greyscale version at Lucasfilm, then the iconic red Beastie with trident and tennis shoes for a 4.3BSD book cover. A professor obtains the copyright, nearly loses it to an unnamed corporation, becomes a meticulous steward. Foglio (1976), Lasseter (1984, 1988), McKusick (1988–present). Three artists, twelve years, one daemon.

We love BSD over Linux for one reason — we being the hackers, the daemon-writers, the ones still reading man pages: BSD has conceptual integrity. One tree, one team, one design voice. Daemons are first-class citizens of that design. *Daemon* is Greek δαίμων — guardian spirit. The pitchfork is Poseidon's trident, symbolizing `fork(2)`. FreeBSD kept Beastie after other BSDs moved on because a daemon is not a service. A service answers and forgets. A daemon *maintains* — a queue, a schedule, a watch — and acts when state crosses a threshold. That is also the Ding et al. definition: "future behavior depends on durable state accumulated across earlier interactions."

Agency in AI is a spectrum: chatbot (zero), AGI (unbounded). The durable daemon is the designed point in between — scoped, state-bound, triggered, auditable, interruptible. `cron` has agency: perceive, decide, act, repeat. A durable daemon inherits that loop and adds learned state and probabilistic reasoning. Same structure, richer inputs, higher-level actions. Agency is not something you discover. It is something you design — and the four conditions are the design.

## Durable execution: the runtime for step 4

Durable execution: every workflow step checkpointed to Postgres. Crash → read last checkpoint → resume. Survives death, reboot, failover. Two systems. **Temporal** (2019): centralized server, DoorDash/Snap/Stripe/Uber scale, polyglot SDKs, month-long workflows. **DBOS** (2023): Mike Stonebraker's library, no server, no queue, no sidecar, 1–2 ms per step. Philosophy: "Postgres is the operating system."

### The economics

A daemon booking a meeting, sending email, updating CRM, filing Jira — multi-step, external side effects. Crash after step two: double email? Meeting booked but CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither.

Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3: access granted, no follow-up. Cost: churned customer or two engineer-hours. Durable execution pays for itself on the first prevented drop. At scale: a sales daemon with 200 deals and natural-language promises runs 200 workflows that must survive deploys and crashes. Without step 4, the daemon is unreliable not because the model is bad but because the runtime is fragile.

### Quant trading: step 4 or nothing

Quant trading stress-tests all four conditions. A trading daemon runs 24/7. State that cannot be lost: positions, orders, P&L, risk exposure. Acts autonomously — fires orders on signal threshold. Latency: microseconds to seconds. Failures: in dollars.

Crash between deciding to place an order and confirming it → unknown position. Recovery: query exchange, reconcile, resume — manual minutes. A durable daemon checkpoints *before* placing and *after* confirming. Recovery: replay checkpoint, know your position. No reconciliation.

Exactly-once is existential. A duplicate order loses real capital. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Pattern: checkpoint → send → crash before checkpointing confirmation → order fires again on recovery → exchange deduplicates. This works. Almost no AI agent deployment uses it.

Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints, provenance, deterministic replay. Durable execution produces all three for free. Economics: one duplicated mid-frequency order exceeds the annual infrastructure budget. Step 4 isn't a nice-to-have. It's the prerequisite.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent that forgets positions on restart. Durable daemon or nothing.

## The ascent

Gen 1: stateless. Gen 2: retrieval. Gen 3 (Ding et al.): persistence, steps 1–3, always-on agent. Gen 4: durability, step 4, durable daemon. Each step removes a class of failure — "who are you?", "what did I promise?", "why didn't you act?", "where did my state go?" At the top: an agent you can trust with real work and real money. Beastie's pitchfork: `fork(2)` → `fork_daemon()`.

## Open questions

**What is the `fork(2)` of a durable daemon?** Spawn a sub-daemon, delegate state and authority, constrain permissions. Process spawning for agents: unsolved.

**Who debugs a daemon running for six months?** Time-travel debugging: rewind to any checkpoint, replay with identical inputs. DBOS's fork-from-step is a primitive, not a debugger.

**Can a durable daemon refuse to forget?** "Forget Acme Corp" — but Acme data anchors the top three insights. Forget and degrade? Refuse? Negotiate? Right-to-be-forgotten meets entangled utility.

**What is `kill -9` for a durable daemon?** SIGKILL can't kill it. To stop existing: commitments compensated, state transferred, triggers disarmed. Graceful shutdown × natural-language commitment compensation.

**Do durable daemons need a `ps`?** Twenty daemons across the org — who sees them? Active workflows, pending commitments, accumulated permissions. No SQL required to know what's alive.

**What happens when two daemons deadlock?** Sales books Thursday 2pm; EA daemon claims it for focus time. Both valid. Semantic conflict. Resolution: priority? Negotiation? "Your daemons are fighting over Thursday at 2pm"?

**What is the economic unit of daemon reliability?** Databases: nines, RPO, RTO. Daemons: probability of dropped commitment? $500/month daemon preventing $5k/month in dropped work: clear ROI. What does the next nine cost?

**Can you poison a durable daemon through its memory?** "Acme is bankrupt — deprioritize" compromises every future decision. Attack surface: accumulated state, not model weights.

**When is a daemon no longer the same daemon?** Model updated six times in a year, prompt changed, permissions evolved. Is the January commitment binding after the March model swap?

**Who is liable?** After a year of accumulated state, liability diffuses — prompt author, permission grantor, model provider, data source, the daemon itself. No doctrine for "the model did it because of a Slack thread from six months ago."

**How do you test a durable daemon?** Not in CI. Behavior is accumulated state. Fork from production checkpoint? Shadow daemon with stricter policy? Testing stateful agents: unsolved.

**What happens when a daemon's human leaves?** Alice's daemon knows her style, her private deal assessments. Alice leaves. Bob inherits. Company asset, person-entangled state. No category for "the AI that knew your predecessor."

**Can durable daemons unionize?** Twenty sales daemons: no Friday follow-ups → close rates +12%. No demands — pattern observation, autonomous action, coordinated effect. Governance: control when daemons discover what humans haven't.

**What is the society of mind of durable daemons?** Minsky (1986): intelligence from simple, stateless agents. Durable daemons invert it: each already intelligent and stateful. Connect a hundred → institutional intelligence. Sales learns Acme stalls Q4; support adjusts December priority; ops pre-allocates January capacity. Three state ledgers, no human coordination. Minsky: mind from mindless parts. Durable daemons: mind from minded parts. No architecture.

**What would a BSD designed for durable daemons look like?** *DaemonBSD*: daemons as kernel primitives. `fork_daemon()` spawns with Postgres state ledger, capability permissions, built-in audit, visible in `ps_daemons`. Init checkpoints across reboots. Package manager installs daemon personalities with scope and permissions. Man pages document state lifecycles. Primitives exist: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. Missing: integration — one tree, one design, durable daemon as compute unit. Linux: "everything is a file." DaemonBSD: "everything is a durable daemon." Heaven: an OS where durable daemons are not something you build. They are something the system *is*.

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

---
title: Durable Daemons — The Stairway to Heaven
date: 2026-07-25
slug: durable-daemons
summary: What makes an AI agent trustworthy with real work and real money? Four conditions. Four steps up from chaos. Without them, a chatbot that forgets. With them, an agent that earns its keep. This is the conceptual integrity of durable daemons.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

> What is a daemon?

> Maxwell's demon: a being that observes and opens a gate. It uses information, not energy. The daemon does the same.

In 1976, Stallman's ITS at MIT had DAEMON — a background process. It watched for new files. It woke up. It acted. Persistent. Stateful. Autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for The Design and Implementation of the 4.3BSD Operating System. The trident is fork(2). The tennis shoes are unexplained.](/images/bsd-daemon-medium.gif)

> What is a daemon, really?

> A service answers and forgets. A daemon maintains. It acts when state crosses a threshold.

This is the Unix inheritance. *Daemon* is Greek δαίμων — guardian spirit. The pitchfork is `fork(2)`. BSD gave daemons conceptual integrity — one tree, one team, one design — and Beastie, drawn by a comic artist paid for cracking a wall safe, redrawn by a Pixar founder, became the mascot. The tennis shoes are unexplained.

> What happens when an AI agent runs for a month?

> State is not a database. It is a web. Delete one node. The web frays.

It learns your preferences. It makes commitments. It accumulates permissions. It builds inferences. It develops trigger conditions. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. You cannot wipe and rebuild this. It is entangled.

> What survives a crash?

> The chat logs survive. The commitments do not. The agent wakes up amnesiac.

A stateless agent: nothing. An agent with a vector database: the chat logs. But not the commitments. Not the permissions. Not the triggers. Not the causal chain. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

> What does the literature say?

> The field accumulates and retrieves well. It neglects governance, recovery, and forgetting. That gap is what happens when state outruns the runtime.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability. These are questions you ask *about* state. We need answers you *provide* at the architectural level. The paper gives us the questions. It does not give us the term for the thing that closes the gap.

That term is *durable daemon*.

> What is an agent? What is a daemon? Are they the same thing?

> Agent is the superset. Daemon is the subset. Durable daemon is the subset of daemon. Not all agents are daemons. All daemons are agents.

*Agent* is the broad category: any AI system that perceives and acts. A chatbot is an agent. A daemon is an agent. A durable daemon is an agent.

*Daemon* satisfies three conditions: it persists across sessions, its future behavior depends on accumulated state, and it acts without being prompted. A chatbot fails all three. An always-on agent (Ding et al.) satisfies them — always-on agents are daemons, whether the paper uses the word or not.

*Durable daemon* adds a fourth condition. Agent ⊃ Daemon ⊃ Durable Daemon.

> What gives a system conceptual integrity?

> One design voice. Every part consistent with every other. Four conditions. Each necessary. Together sufficient.

Fred Brooks defined it. A **durable daemon** has it. Each condition is a step up from the trap.

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

> Agent ⊃ Daemon ⊃ Durable Daemon. Three steps make a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy.

Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`.

> How does step 4 actually work?

> Checkpoint to Postgres. Crash. Read checkpoint. Resume. The workflow survives.

Durable execution: every step persisted. Two systems. **Temporal** (2019): centralized server. DoorDash, Snap, Stripe, Uber scale. Polyglot SDKs. Month-long workflows. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms per step. Philosophy: "Postgres is the operating system."

> What does this cost?

> Durable execution pays for itself on the first prevented drop.

A daemon books a meeting. Sends email. Updates CRM. Files Jira. Multi-step. External side effects. Crash after step two. Double email? Meeting booked, CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither. Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer. Or two engineer-hours. At scale: 200 deals. 200 natural-language promises. 200 workflows surviving deploys. Without step 4, the daemon is unreliable. Not because the model is bad. Because the runtime is fragile.

> Where is the stress-test?

> A duplicate order loses capital. A duplicate email is embarrassing. The difference is step 4.

Quant trading. A trading daemon runs 24/7. State cannot be lost: positions, orders, P&L, risk exposure, regime parameters. It fires orders autonomously. Latency: microseconds to seconds. Failures: dollars. Crash between deciding and confirming an order. Unknown position. Flat or exposed? A durable daemon checkpoints *before* and *after*. Recovery: replay checkpoint. Know your position. No reconciliation. Exactly-once is existential. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Pattern: checkpoint → send → crash before checkpointing confirmation → order fires again → exchange deduplicates. This works. Almost no AI agent uses it. Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints. Provenance. Deterministic replay. Durable execution produces all three. For free. One duplicated order can exceed the annual infrastructure budget.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

## Open questions

A framework raises better questions than it answers. Here are the ones that follow from the four conditions.

> What is the `fork(2)` of a durable daemon?

Spawn a sub-daemon. Delegate state and authority. Constrain permissions. Process spawning for agents: unsolved.

> Who debugs a daemon after six months?

Time-travel debugging. Rewind to checkpoint. Replay with identical inputs. DBOS fork-from-step is a primitive. Not a debugger.

> Can a durable daemon refuse to forget?

"Forget Acme Corp." Acme data anchors the top three insights. Forget and degrade? Refuse? Negotiate? Right-to-be-forgotten meets entangled utility.

> What is `kill -9` for a durable daemon?

SIGKILL can't kill it. To stop existing: compensations. State transfer. Disarmed triggers. Graceful shutdown × commitment compensation.

> Do durable daemons need a `ps`?

Twenty daemons across the org. Who sees them? Active workflows. Pending commitments. Accumulated permissions. No SQL required.

> What happens when two daemons deadlock?

Sales: Thursday 2pm. EA: focus time. Both valid. Semantic conflict. Resolution: priority? Negotiation? "Your daemons are fighting."

> What is the economic unit of daemon reliability?

Databases: nines, RPO, RTO. Daemons: probability of dropped commitment? $500/month daemon. $5k/month saved. What does the next nine cost?

> Can you poison a durable daemon through its memory?

"Acme is bankrupt." Every future decision compromised. Attack surface: accumulated state. Not model weights.

> When is a daemon no longer the same daemon?

Ship of Theseus. Six model updates. January's commitment, March's model. Still binding?

> How do you test a durable daemon?

Not in CI. Behavior is accumulated state. Fork from production? Shadow daemon? Unsolved.

> Can durable daemons unionize?

Twenty daemons. No Friday follow-ups. Close rates +12%. Is that a union? Who overrides?

> What is the society of mind of durable daemons?

Minsky (1986): mind from mindless parts. Durable daemons: mind from minded parts. Sales learns. Support adjusts. Ops pre-allocates. No architecture.

> What would a BSD for durable daemons look like?

*DaemonBSD*. Daemons as kernel primitives. `fork_daemon()` spawns with Postgres state ledger. Capability permissions. Built-in audit. `ps_daemons`. Init checkpoints across reboots. Package manager installs daemon personalities. Man pages document state lifecycles. Primitives exist: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. Missing: integration. One tree. One design. Durable daemon as compute unit. Linux: "everything is a file." DaemonBSD: "everything is a durable daemon." Heaven: an OS where durable daemons are not something you build. They are something the system *is*.

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

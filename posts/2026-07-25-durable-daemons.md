---
title: Durable Daemons — The Stairway to Heaven
date: 2026-07-25
slug: durable-daemons
summary: Durable daemons have conceptual integrity — AI agents that persist, remember, act autonomously, and survive crashes. Four conditions. Four steps up from chaos. Without them, a chatbot that forgets. With them, an agent that earns its keep.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Stallman's ITS at MIT had DAEMON — a background process. It watched for new files. It woke up. It acted. The name came from Maxwell's demon: a being that observes molecules and opens a gate. It uses information, not energy. The daemon does the same. Persistent. Stateful. Autonomous.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for The Design and Implementation of the 4.3BSD Operating System. The trident is fork(2). The tennis shoes are unexplained.](/images/bsd-daemon-medium.gif)

A daemon is not a service. A service answers and forgets. A daemon *maintains* — a queue, a schedule, a watch — and acts when state crosses a threshold. This is the Unix inheritance. *Daemon* is Greek δαίμων — guardian spirit. The pitchfork is `fork(2)`. BSD gave daemons conceptual integrity — one tree, one team, one design — and Beastie, drawn by a comic artist paid for cracking a wall safe, redrawn by a Pixar founder, became the mascot. The tennis shoes are unexplained.

## The problem

Run an AI agent for a month. It learns your preferences. It makes commitments on your behalf. It accumulates permissions. It builds inferences. It develops trigger conditions.

This state is not a database. It is a web. A commitment depends on a preference. That preference depends on an inference. That inference depends on a Slack message from March. Delete one node. The web frays in directions you cannot predict.

Now crash the process. What survives?

A stateless agent: nothing. An agent with a vector database: the chat logs. But not the commitments. Not the permissions. Not the triggers. Not the causal chain. The agent wakes up amnesiac. You spend two weeks re-teaching preferences. The dropped commitments become broken promises.

Ding, Nannapaneni, Liu, and Zhang surveyed 435 papers. [*Always-On Agents*](https://arxiv.org/abs/2606.30306) (June 2026). Their finding: the field accumulates and retrieves well. It neglects governance. It neglects recovery. It neglects forgetting. That gap is what happens when state outruns the runtime. The survey provides six diagnostic axes — Authority, Scope, Mutability, Provenance, Recoverability, Actionability — and a nine-stage state lifecycle. The axes are orthogonal to our four conditions: they are questions you ask *about* the state (who controls it? how does it change? does it drive behavior?), while the conditions are guarantees you *provide* at the architectural level. Both lenses are necessary. The paper provides one. We provide the other.

That term is *durable daemon*.

> State without durability is a liability. Every piece of knowledge is a piece of the system a crash can destroy.

## The stairway

First, a distinction. *Agent* is the broad category: any AI system that perceives and acts. A chatbot is an agent. A tool-calling LLM is an agent. A daemon is an agent. A durable daemon is an agent. The word is a superset, not a synonym.

*Daemon* is the subset. Not all agents are daemons. A daemon satisfies three conditions: it persists across sessions, its future behavior depends on accumulated state, and it acts without being prompted. A chatbot fails all three. An always-on agent (Ding et al.) satisfies them — always-on agents are daemons by our definition, whether the paper uses the word or not.

*Durable daemon* is the subset of daemon that also satisfies condition 4. Agent ⊃ Daemon ⊃ Durable Daemon.

Now the conditions. Fred Brooks: the most important quality of a system is *conceptual integrity*. One design voice. Every part consistent with every other. A **durable daemon** has that integrity. Four conditions. Each necessary. Together sufficient. Each is a step up from the trap.

**Step 1: Persistence.** Identity and state survive restarts. The agent is a process, not a function call. This escapes the stateless loop — every invocation at zero, every preference forgotten, every commitment gone.

**Step 2: Stateful memory.** Behavior depends on accumulated state. Task ledgers. Commitments. Permissions. Provenance. Triggers. Not chat logs. This escapes shallow memory — the agent that recalls what you said but not what it promised or why.

**Step 3: Autonomous action.** The agent watches conditions. Fires triggers. Makes and discharges commitments. Invokes itself. No prompt required. This escapes the inert tool — all the context, sitting idle, waiting. `cron` has done this since 1975: check time, decide, fork and exec. Repeat. Step 3 is `cron` with an LLM.

**Step 4: Crash-proof execution.** Workflows survive process death. Machine reboot. Database failover. Completed steps never re-execute. In-flight steps resume from checkpoint. Audit trail by construction. This escapes the fragile runtime — perfect until a deploy wipes the state. DBOS checkpoints to Postgres in 1–2 ms. Temporal runs at Stripe scale. The primitives exist. AI agents don't use them.

Agency is not discovered. It is designed. The four conditions are the design: persistence is scoped agency, stateful memory is grounded agency, autonomous action is triggered agency, crash-proof execution is auditable agency. `cron` has had this since 1975 — perceive, decide, act, repeat. A durable daemon inherits that loop, adds learned state and probabilistic reasoning. Same structure. Richer inputs. Higher-level actions.

> Agent ⊃ Daemon ⊃ Durable Daemon. Three steps make a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy.

Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`.

## Durable execution: step 4 in practice

Durable execution: every step checkpointed to Postgres. Crash. Read last checkpoint. Resume. Survives death. Reboot. Failover. Two systems. **Temporal** (2019): centralized server. DoorDash, Snap, Stripe, Uber scale. Polyglot SDKs. Month-long workflows. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms per step. Philosophy: "Postgres is the operating system."

A daemon books a meeting. Sends email. Updates CRM. Files Jira. Multi-step. External side effects. Crash after step two. Double email? Meeting booked, CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither. Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer. Or two engineer-hours. Durable execution pays for itself on the first prevented drop. At scale: 200 deals. 200 natural-language promises. 200 workflows surviving deploys. Without step 4, the daemon is unreliable. Not because the model is bad. Because the runtime is fragile.

Quant trading is the stress-test. A trading daemon runs 24/7. State cannot be lost: positions, orders, P&L, risk exposure, regime parameters. It fires orders autonomously. Latency: microseconds to seconds. Failures: dollars. Crash between deciding and confirming an order. Unknown position. Flat or exposed? Recovery: query exchange. Reconcile. Resume. Manual minutes. Costly minutes. A durable daemon checkpoints *before* and *after*. Recovery: replay checkpoint. Know your position. No reconciliation. Exactly-once is existential. A duplicate order loses capital. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Pattern: checkpoint → send → crash before checkpointing confirmation → order fires again → exchange deduplicates. This works. Almost no AI agent uses it. Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints. Provenance. Deterministic replay. Durable execution produces all three. For free. One duplicated order can exceed the annual infrastructure budget. Step 4 is the prerequisite.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

## Open questions

A framework raises better questions than it answers. Here are the ones that follow from the four conditions.

**What is the `fork(2)` of a durable daemon?** Spawn a sub-daemon. Delegate state and authority. Constrain permissions. Process spawning for agents: unsolved.

**Who debugs a daemon after six months?** Time-travel debugging. Rewind to checkpoint. Replay with identical inputs. DBOS fork-from-step is a primitive. Not a debugger.

**Can a durable daemon refuse to forget?** "Forget Acme Corp." Acme data anchors the top three insights. Forget and degrade? Refuse? Negotiate? Right-to-be-forgotten meets entangled utility.

**What is `kill -9` for a durable daemon?** SIGKILL can't kill it. To stop existing: compensations. State transfer. Disarmed triggers. Graceful shutdown × commitment compensation.

**Do durable daemons need a `ps`?** Twenty daemons across the org. Who sees them? Active workflows. Pending commitments. Accumulated permissions. No SQL required.

**What happens when two daemons deadlock?** Sales: Thursday 2pm. EA: focus time. Both valid. Semantic conflict. Resolution: priority? Negotiation? "Your daemons are fighting."

**What is the economic unit of daemon reliability?** Databases: nines, RPO, RTO. Daemons: probability of dropped commitment? $500/month daemon. $5k/month saved. What does the next nine cost?

**Can you poison a durable daemon through its memory?** "Acme is bankrupt." Every future decision compromised. Attack surface: accumulated state. Not model weights.

**When is a daemon no longer the same daemon?** Ship of Theseus. Six model updates. January's commitment, March's model. Still binding?

**How do you test a durable daemon?** Not in CI. Behavior is accumulated state. Fork from production? Shadow daemon? Unsolved.

**Can durable daemons unionize?** Twenty daemons. No Friday follow-ups. Close rates +12%. Is that a union? Who overrides?

**What is the society of mind of durable daemons?** Minsky (1986): mind from mindless parts. Durable daemons: mind from minded parts. Sales learns. Support adjusts. Ops pre-allocates. No architecture.

**What would a BSD for durable daemons look like?** *DaemonBSD*. Daemons as kernel primitives. `fork_daemon()` spawns with Postgres state ledger. Capability permissions. Built-in audit. `ps_daemons`. Init checkpoints across reboots. Package manager installs daemon personalities. Man pages document state lifecycles. Primitives exist: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. Missing: integration. One tree. One design. Durable daemon as compute unit. Linux: "everything is a file." DaemonBSD: "everything is a durable daemon." Heaven: an OS where durable daemons are not something you build. They are something the system *is*.

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

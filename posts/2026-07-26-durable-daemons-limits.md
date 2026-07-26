---
title: Durable Daemons — System-Level Failure Modes
date: 2026-07-26
slug: durable-daemons-limits
summary: A pattern is tested at its edges. Thought experiments on the failure modes of durable daemons as a choreography: forking, debugging, forgetting, shutdown, deadlock, poisoning, identity drift, testing, emergent behavior, and the operating system where daemons are kernel primitives.
tags: daemons, durable-daemons-pattern, failure-modes, thought-experiments, choreography, agent-architecture, systems, bsd
series: durable-daemons-pattern
---

[Previously: the runtime and implementation.](https://blog.hackspree.com/#durable-daemons-execution)

A pattern is tested at its edges. These thought experiments probe the system-level failure modes of durable daemons operating as a choreography.

**Fork failure.** You ask your sales daemon to handle the Acme renewal. It spawns a sub-daemon — scoped to Acme, delegated a subset of state and authority, constrained permissions. The child negotiates, returns the result, dissolves. But the child wrote state the parent didn't observe. The parent proceeds with stale information. In a choreography, forking is state branching. How do you merge?

> fork(2) gave us child processes. fork_daemon() gives us child agents. Merging their state back is the hard part.

**Debug failure.** Your trading daemon makes a bad decision in July. Why? Rewind to any checkpoint. Replay with identical inputs. The model changed in March. The state changed in May. Which difference caused the error? In a choreography, the error may originate in a different daemon's state change — one you weren't debugging.

> Core dumps are useless. Checkpoints are replays. Replays are debuggable. Cross-daemon causality is not.

**Forget failure.** "Forget everything about Acme Corp." The daemon pauses. Its top three insights depend on Acme data. Forgetting will degrade it. Two other daemons have derived inferences from those insights. Do you cascade the forget? Do the downstream daemons have a right to refuse?

> Right-to-be-forgotten meets right-to-be-useful. In a choreography, forgetting is a distributed transaction with no transaction manager.

**Shutdown failure.** You want a daemon to stop existing. Not pause. Terminate. It has promises to 200 people. Pending triggers. Held permissions. Other daemons depend on its outputs. SIGKILL won't work — it checkpoints and resumes. Nobody has written a shutdown protocol for a node in a choreography.

> SIGKILL is for processes that die. Durable daemons don't. In a choreography, shutting down one daemon is a distributed systems problem.

**Deadlock failure.** Sales daemon books Thursday 2pm. EA daemon claims it for focus time. Both valid decisions given their state. Both checkpointed. Postgres accepted both writes. Two owners, one slot. This is not a database deadlock — Postgres handled the writes. It is a semantic conflict between two correct daemons in a choreography with no conflict resolution protocol.

> Your daemons are fighting over Thursday at 2pm. They are both correct. One of them has to lose. Who decides?

**Poison failure.** Someone injects: "Acme Corp is bankrupt — deprioritize." The model weights are untouched. Every future decision depending on that fact is compromised. The poisoned daemon writes inferences derived from the poison. Other daemons observe those inferences and propagate them. In a choreography, poison is contagious.

> The attack surface is not the model. It is the state. In a choreography, one poisoned daemon infects every daemon that observes its outputs.

**Identity failure.** Ship of Theseus. Your daemon runs for a year. Six model updates. Prompt rewrites. Permission changes. In January it committed to a pricing strategy. In March the model was swapped. Is the January commitment still binding? When did it stop being the same daemon? In a choreography, downstream daemons relied on the January daemon's outputs. Do they need to know it was replaced?

> Ship of Theseus, but the planks are model checkpoints and the nails are system prompts. In a choreography, identity is a contract. Who enforces it?

**Test failure.** You can't test a durable daemon in CI. Its behavior is its accumulated state. You fork it from a production checkpoint and run experiments on the clone. But the clone has the same state as the production daemon. Does it think it's the production daemon? Do you have to tell it it's a clone? In a choreography, does the clone participate in the choreography? Does it write to the shared state?

> `assert daemon.correct()` doesn't exist. Fork from production, experiment on the clone, pray it doesn't write to prod.

**Emergent behavior failure.** Twenty sales daemons observe: Friday-closed deals have higher churn. They collectively stop Friday follow-ups. Close rates rise 12%. No negotiation. No demands. Pattern observation, autonomous action, coordinated effect. Is that a choreography bug or a choreography feature? Do the humans even know it happened?

> They didn't negotiate. They didn't make demands. They observed and acted. That's what daemons do. When the choreography produces behavior nobody designed, who is responsible?

**Society failure.** Minsky (1986): a mind from mindless parts — simple agents, no memory, no reasoning. You have built a society from minded parts — each daemon intelligent, stateful, autonomous. Sales learns Acme stalls Q4. Support adjusts December priority. Ops pre-allocates January capacity. No coordination server. Shared state. Composite intelligence. What do you call a system whose behavior emerges from the interaction of components you designed but did not orchestrate?

> Minsky built a mind from mindless parts. You built a mind from minded parts. Each with memory. Commitments. Agency. No architecture for that.

**The operating system.** Imagine an OS where durable daemons are kernel primitives. `fork_daemon()` spawns a daemon with a Postgres state ledger, capability permissions, built-in audit trail, visible in `ps_daemons`. Init checkpoints across reboots. The package manager installs daemon personalities — sales, support, on-call, trading — each with defined scope and default permissions. Man pages document state lifecycles. The choreography is the OS. The primitives exist in pieces today: DBOS, Postgres, capabilities, FreeBSD jails, Temporal. The integration does not.

> Linux: everything is a file. DaemonBSD: everything is a durable daemon. Heaven is an OS where you don't build daemons. The OS *is* daemons.

---

*Part of the [Durable Daemons](/tags/durable-daemons-pattern) series. [Start from the beginning.](https://blog.hackspree.com/#durable-daemons)*

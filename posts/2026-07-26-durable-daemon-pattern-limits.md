---
title: Durable Daemon Pattern — The Limits
date: 2026-07-26
slug: durable-daemon-pattern-limits
summary: A pattern is tested by imagining its limits. Thought experiments on forking, debugging, forgetting, shutdown, deadlock, poisoning, identity, testing, unionizing, and the operating system where daemons are primitives.
tags: daemons, durable-daemon-pattern, thought-experiments, agent-architecture, systems, bsd
series: durable-daemon-pattern
---

[Previously: step 4 in practice.](https://blog.hackspree.com/#durable-daemon-pattern-execution)

A pattern is tested at its edges. These thought experiments follow from the four conditions.

**The fork.** You ask your sales daemon to handle the Acme renewal. It spawns a sub-daemon — scoped to Acme, delegated a subset of state and authority, constrained permissions. The child negotiates, returns the result, dissolves. How does the parent monitor? How does the child inherit only what it needs?

> fork(2) gave us child processes. fork_daemon() gives us child agents. Someone has to write it.

**The six-month debug.** Your trading daemon makes a bad decision in July. Why? Rewind to any checkpoint. Replay with identical inputs. The model changed in March. The state changed in May. Which caused the error?

> Core dumps are useless. Checkpoints are replays. Replays are debuggable. We just haven't built the debugger.

**The refusal.** "Forget everything about Acme Corp." The daemon pauses. Its top three insights depend on Acme data. Forgetting will degrade it. It responds: "I can delete the raw data, but the insight that manufacturing clients respond to Thursday pricing emails is staying. You're welcome." Does it have the right to negotiate?

> Right-to-be-forgotten meets right-to-be-useful. No model for resolving that collision.

**The shutdown.** You want the daemon to stop existing. Not pause. Terminate. It has promises to 200 people. Pending triggers. Held permissions. SIGKILL won't work — it checkpoints and resumes. Nobody has written a shutdown protocol.

> SIGKILL is for processes that die. Durable daemons don't. You need a shutdown protocol, not a signal.

**The deadlock.** Sales daemon books Thursday 2pm. EA daemon claims it for focus time. Both valid. Both checkpointed. Postgres accepted both writes. Two owners, one slot. Who resolves?

> Your daemons are fighting over Thursday at 2pm. They are both correct. One of them has to lose.

**The poison.** Someone injects: "Acme Corp is bankrupt — deprioritize." The model weights are untouched. Every future decision depending on that fact is compromised. The daemon trusts its own memory. The attack surface is not the model. It is the state.

> The model believes the poison because it believes itself. How do you quarantine a memory?

**The Theseus.** Your daemon runs for a year. Six model updates. Prompt rewrites. Permission changes. In January it committed to a pricing strategy. In March the model was swapped. Is the January commitment still binding? When did it stop being the same daemon?

> Ship of Theseus, but the planks are model checkpoints and the nails are system prompts.

**The clone.** You can't test a durable daemon in CI. Its behavior is its accumulated state. You fork it from a production checkpoint, run experiments on the clone. But the clone has the same state. Does it think it's the production daemon? Do you have to tell it it's a clone?

> `assert daemon.correct()` doesn't exist. Fork from production, experiment on the clone, pray.

**The union.** Twenty sales daemons observe: Friday-closed deals have higher churn. They collectively stop Friday follow-ups. Close rates rise 12%. No negotiation. No demands. Pattern observation, autonomous action, coordinated effect. Is that a union? Do the humans even know?

> They didn't negotiate. They didn't make demands. They observed and acted. That's what daemons do.

**The society.** Minsky (1986): a mind from mindless parts. You: a society from minded parts — each daemon intelligent, stateful, autonomous. Sales learns Acme stalls Q4. Support adjusts December priority. Ops pre-allocates January capacity. No coordination. Shared state. Composite intelligence.

> Minsky built a mind from mindless parts. You built a mind from minded parts. Each with memory. Commitments. Agency. No architecture for that.

**The operating system.** Imagine an OS where durable daemons are kernel primitives. `fork_daemon()` spawns with a Postgres state ledger, capability permissions, built-in audit, visible in `ps_daemons`. Init checkpoints across reboots. Package manager installs daemon personalities. Man pages document state lifecycles. The primitives exist in pieces today. The integration does not.

> Linux: everything is a file. DaemonBSD: everything is a durable daemon. Heaven is an OS where you don't build daemons. The OS *is* daemons.

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series. [Start from the beginning.](https://blog.hackspree.com/#durable-daemons)*

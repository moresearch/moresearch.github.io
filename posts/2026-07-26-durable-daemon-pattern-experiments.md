---
title: Durable Daemon Pattern — The Thought Experiments
date: 2026-07-26
slug: durable-daemon-pattern-experiments
summary: A pattern is tested by imagining its limits. Thirteen thought experiments that follow from the four conditions — the fork, the debug, the forget, the shutdown, the deadlock, the poison, the Theseus, the union, the society, and the operating system where daemons are primitives.
tags: daemons, durable-daemon-pattern, thought-experiments, agent-architecture, systems, bsd
series: durable-daemon-pattern
---

[Previously: step 4 in practice — execution, economics, and quant trading.](https://blog.hackspree.com/#durable-daemon-pattern-execution)

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

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series. [Start from the beginning.](https://blog.hackspree.com/#durable-daemon-pattern-genealogy)*

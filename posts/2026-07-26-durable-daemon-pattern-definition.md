---
title: Durable Daemon Pattern — The Definition
date: 2026-07-26
slug: durable-daemon-pattern-definition
summary: The durable daemon pattern is four conditions — persistence, stateful memory, autonomous action, crash-proof execution. Each necessary. Together sufficient. Agent ⊃ Daemon ⊃ Durable Daemon. This is the conceptual integrity of the pattern.
tags: daemons, durable-daemon-pattern, ai-agents, always-on-agents, agent-architecture, systems
series: durable-daemon-pattern
---

[Previously: the problem the pattern solves.](https://blog.hackspree.com/#durable-daemon-pattern-problem)

> What is an agent? What is a daemon? Are they the same thing?

> All daemons are agents. Not all agents are daemons. A chatbot is an agent. A chatbot is not a daemon. The difference is three conditions.

*Agent* is the broad category: any AI system that perceives and acts. *Daemon* narrows it: the agent must persist across sessions, act on accumulated state, and invoke itself without prompting. An always-on agent (Ding et al.) satisfies these — always-on agents are daemons, whether the paper says so or not. *Durable daemon* adds crash-proof execution. Agent ⊃ Daemon ⊃ Durable Daemon.

> What gives a system conceptual integrity?

> One design voice. Every part consistent with every other. Four conditions. Each necessary. Together sufficient.

Conceptual integrity means the left hand knows what the right hand is doing. No contradictions. No surprises. Brooks named it. BSD lived it. The durable daemon pattern has it. Each condition is a step up from the trap.

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

> Agent ⊃ Daemon ⊃ Durable Daemon. Three conditions make a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy. That is the pattern.

Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`. The tennis shoes stay the same.

[Next: step 4 in practice — durable execution, economics, and the quant trading stress-test.](https://blog.hackspree.com/#durable-daemon-pattern-execution)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*

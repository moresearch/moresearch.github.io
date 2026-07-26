---
title: Durable Daemons — Pattern Specification
date: 2026-07-26
slug: durable-daemons-definition
summary: The durable daemons pattern specifies four conditions. Each necessary. Together sufficient. Agent ⊃ Daemon ⊃ Durable Daemon. Persistence, stateful memory, autonomous action, crash-proof execution. Daemons coordinate through shared state — an event-driven choreography with no central orchestrator.
tags: daemons, durable-daemons-pattern, pattern-specification, agent-architecture, choreography, systems
series: durable-daemons-pattern
---

[Previously: the state persistence problem.](https://blog.hackspree.com/#durable-daemons-trap)

Imagine you are asked to trust an AI agent with your calendar. Your email. Your money. What would you need to know? You would need to know that it persists — it does not start over every time you open the app. You would need to know that it remembers — not just what you said, but what it promised and why. You would need to know that it acts — not when you tell it to, but when conditions require it. You would need to know that it survives — a deploy, a crash, a reboot does not erase its commitments. These are not preferences. They are preconditions.

> Trust requires guarantees, not vibes. Four guarantees. Four conditions. That is the specification.

First, the type hierarchy. *Agent* is the broad category: any AI system that perceives and acts. A chatbot is an agent. *Daemon* is the subtype: an agent that persists across sessions, whose future behavior depends on accumulated state, and which acts without being prompted. A chatbot fails all three. An always-on agent (Ding et al.) satisfies them. *Durable daemon* adds a fourth condition. Agent ⊃ Daemon ⊃ Durable Daemon.

Fred Brooks taught that the most important quality of a system is *conceptual integrity* — the left hand knows what the right hand is doing. No contradictions. No surprises. Brooks named it. BSD lived it. The durable daemons pattern has it.

> One design voice. Every part consistent with every other. Four conditions. Each necessary. Together sufficient.

**Condition 1: Persistence.** Identity and state survive restarts. The daemon is a process, not a function call. This eliminates the stateless loop failure mode — every invocation at zero, every preference forgotten, every commitment gone.

**Condition 2: Stateful memory.** Behavior depends on accumulated state. Task ledgers. Commitments. Permissions. Provenance records. Trigger conditions. Not a vector database of chat logs. This eliminates the shallow memory failure mode — the agent that recalls what you said but not what it promised or why.

**Condition 3: Autonomous action.** The daemon watches conditions. Fires triggers. Makes and discharges commitments. Invokes itself. No prompt required. This eliminates the inert tool failure mode — all the context, sitting idle, waiting. `cron` has done this since 1975: perceive, decide, act, repeat. Condition 3 is `cron` with an LLM.

**Condition 4: Crash-proof execution.** Workflows survive process death. Machine reboot. Database failover. Completed steps never re-execute. In-flight steps resume from the last checkpoint. Audit trail by construction. This eliminates the fragile runtime failure mode — perfect until a deploy wipes the in-flight state.

These conditions compose. A durable daemon writes its state to Postgres. Another durable daemon observes that state and reacts. A third observes the second and reacts. Each daemon satisfies all four conditions independently. The choreography emerges from shared state — no orchestrator, no message bus, no coordination server. This is an event-driven architecture where the event store *is* the state store, and every daemon is both a producer and a consumer.

> Agency is not discovered. It is designed. Persistence is scoped agency. Stateful memory is grounded agency. Autonomous action is triggered agency. Crash-proof execution is auditable agency. The choreography is the composition.

Three conditions define a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy. The pattern composes them into a system. Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`. The tennis shoes stay the same.

[Next: the runtime and implementation — step 4 in practice.](https://blog.hackspree.com/#durable-daemons-execution)

---

*Part of the [Durable Daemons](/tags/durable-daemons-pattern) series.*

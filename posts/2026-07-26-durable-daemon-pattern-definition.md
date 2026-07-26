---
title: Durable Daemon Pattern — The Definition
date: 2026-07-26
slug: durable-daemon-pattern-definition
summary: Four conditions. Each necessary. Together sufficient. Agent ⊃ Daemon ⊃ Durable Daemon. Persistence, stateful memory, autonomous action, crash-proof execution. This is the conceptual integrity of the durable daemon pattern.
tags: daemons, durable-daemon-pattern, ai-agents, agent-architecture, systems
series: durable-daemon-pattern
---

[Previously: the trap of long-running AI state.](https://blog.hackspree.com/#durable-daemon-pattern-trap)

Imagine you are asked to trust an AI agent with your calendar. Your email. Your money. What would you need to know? You would need to know that it persists — it does not start over every time you open the app. You would need to know that it remembers — not just what you said, but what it promised and why. You would need to know that it acts — not when you tell it to, but when conditions require it. You would need to know that it survives — a deploy, a crash, a reboot does not erase its commitments. These are not preferences. They are conditions.

> Trust requires guarantees, not vibes. Four guarantees. Four conditions. That is the pattern.

First, a distinction. *Agent* is the broad category: any AI system that perceives and acts. A chatbot is an agent. *Daemon* is the subset: an agent that persists across sessions, whose future behavior depends on accumulated state, and which acts without being prompted. A chatbot fails all three. An always-on agent (Ding et al.) satisfies them — always-on agents are daemons, whether the paper says so or not. *Durable daemon* adds a fourth condition. Agent ⊃ Daemon ⊃ Durable Daemon.

Fred Brooks taught that the most important quality of a system is *conceptual integrity* — the left hand knows what the right hand is doing. No contradictions. No surprises. Brooks named it. BSD lived it. The durable daemon pattern has it.

> One design voice. Every part consistent with every other. Four conditions. Each necessary. Together sufficient.

**Step 1: Persistence.** Identity and state survive restarts. The agent is a process, not a function call. This escapes the stateless loop — every invocation at zero, every preference forgotten, every commitment gone.

**Step 2: Stateful memory.** Behavior depends on accumulated state. Task ledgers. Commitments. Permissions. Provenance. Triggers. Not chat logs. This escapes shallow memory — the agent that recalls what you said but not what it promised or why.

**Step 3: Autonomous action.** The agent watches conditions. Fires triggers. Makes and discharges commitments. Invokes itself. No prompt required. This escapes the inert tool — all the context, sitting idle, waiting. `cron` has done this since 1975: perceive, decide, act, repeat. Step 3 is `cron` with an LLM.

**Step 4: Crash-proof execution.** Workflows survive process death. Machine reboot. Database failover. Completed steps never re-execute. In-flight steps resume from checkpoint. Audit trail by construction. This escapes the fragile runtime — perfect until a deploy wipes the state.

> Agency is not discovered. It is designed. Persistence is scoped agency. Stateful memory is grounded agency. Autonomous action is triggered agency. Crash-proof execution is auditable agency.

Three conditions make a daemon. The fourth makes it durable — an agent that cannot be killed by a deploy. That is the pattern. Beastie's pitchfork gets an upgrade: `fork(2)` → `fork_daemon()`. The tennis shoes stay the same.

[Next: step 4 in practice — the runtime, the economics, and the quant trading stress-test.](https://blog.hackspree.com/#durable-daemon-pattern-execution)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*

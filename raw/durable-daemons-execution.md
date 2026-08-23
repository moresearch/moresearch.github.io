---
title: Durable Daemons — Runtime and Implementation
date: 2026-07-26
slug: durable-daemons-execution
summary: Condition 4 requires crash-proof execution. DBOS and Temporal provide it — checkpointing workflow state to Postgres, surviving process death, reboot, and failover. Quant trading is the stress-test. The implementation is portable (Go binary + Postgres) and composable (event-driven choreography through shared state).
tags: daemons, durable-daemons-pattern, dbos, temporal, durable-execution, quant-trading, go, choreography, event-driven
series: durable-daemons-pattern
---

[Previously: the pattern specification.](https://blog.hackspree.com/#durable-daemons-definition)

Imagine a trading daemon. It runs 24/7 across market sessions. Its state cannot be lost: open positions, pending orders, realized P&L, risk exposure, regime-adapting strategy parameters. A signal crosses a threshold. The daemon decides to place an order. It sends the order to the exchange. Before the confirmation arrives, the process crashes.

The daemon restarts. What is its position? Did the order reach the exchange? Is it flat or exposed? To know, you must query the exchange, reconcile positions, and resume — manual minutes in a domain where minutes cost money. Now imagine the daemon checkpointed its state *before* placing the order and checkpointed the confirmation *after*. On recovery, it replays the last checkpoint and knows exactly where it stands. No reconciliation. No unknown exposure.

> A duplicate order loses capital. A duplicate email is embarrassing. The difference is condition 4.

This is durable execution: every workflow step persisted to Postgres. Crash. Read last checkpoint. Resume. Completed steps never re-execute. Survives process death, machine reboot, database failover. Two production-grade implementations. **Temporal** (2019): centralized orchestration server, DoorDash/Snap/Stripe/Uber scale, polyglot SDKs, month-long workflows. Bring your own cluster. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms step latency — a single Postgres write. Bring your own Postgres. Either way: bring condition 4.

> Durable execution pays for itself on the first prevented state loss.

Onboarding workflow: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer, or two engineer-hours manually reconciling state — thousands of agent invocations. At scale: 200 deals, 200 natural-language commitments, 200 concurrent workflows that must survive deploys. Without condition 4, the daemon is unreliable. Not because the model is bad. Because the runtime cannot guarantee state survival.

Exactly-once semantics are existential in trading. Durable execution guarantees exactly-once within its control boundary. External calls to exchange APIs require idempotency keys. Pattern: checkpoint → send order → crash before checkpointing confirmation → order fires again on recovery → exchange deduplicates by key. This works. Almost no AI agent deployment implements it. Audit: every order traceable to a decision, every decision traceable to the state the daemon held at the time. SEC, CFTC, FCA, ESMA do not accept "the model decided." They accept checkpointed state, provenance chains, and deterministic replay. Durable execution produces all three as a byproduct of normal operation. One duplicated mid-frequency order can exceed the annual infrastructure budget.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

The implementation is portable. A Go static binary ships the daemon. A Postgres checkpoint ships the state. A container ships both. The daemon runs anywhere Postgres runs — no additional infrastructure dependencies beyond the database.

The implementation is composable. This is the choreography. Two daemons with different scopes — sales and support — share a Postgres cluster but own separate state ledgers. They observe each other's outputs through the database. No RPC. No message bus. No central orchestrator. The sales daemon writes an inference: "Acme stalls Q4." The support daemon observes it and adjusts December ticket priority. The ops daemon observes both and pre-allocates January capacity. Each daemon satisfies all four conditions independently. The choreography is the shared state. `sales-daemon | support-daemon | ops-daemon`. Unix pipes, applied to AI agents.

> Each daemon is a process. Each daemon is durable. The choreography is the composition. No orchestrator. Just state.

[Next: the system-level failure modes — thought experiments at the edges of the pattern.](https://blog.hackspree.com/#durable-daemons-limits)

---

*Part of the [Durable Daemons](/tags/durable-daemons-pattern) series.*

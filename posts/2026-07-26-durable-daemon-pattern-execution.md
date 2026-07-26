---
title: Durable Daemon Pattern — The Execution
date: 2026-07-26
slug: durable-daemon-pattern-execution
summary: Step 4 is crash-proof execution — the runtime that makes a daemon durable. DBOS and Temporal provide it. The pattern is portable (static binary + Postgres) and composable (Unix pipes for agents). Quant trading proves it: durable daemon or nothing.
tags: daemons, durable-daemon-pattern, dbos, temporal, durable-execution, quant-trading, go
series: durable-daemon-pattern
---

[Previously: the four conditions of the pattern.](https://blog.hackspree.com/#durable-daemon-pattern-definition)

Imagine a trading daemon. It runs 24/7 across market sessions. Its state cannot be lost: open positions, pending orders, realized P&L, risk exposure, regime-adapting strategy parameters. A signal crosses a threshold. The daemon decides to place an order. It sends the order to the exchange. Before the confirmation arrives, the process crashes.

The daemon restarts. What is its position? Did the order reach the exchange? Is it flat or exposed? To know, you must query the exchange, reconcile positions, and resume — manual minutes in a domain where minutes cost money. Now imagine the daemon checkpointed its state *before* placing the order and checkpointed the confirmation *after*. On recovery, it replays the last checkpoint and knows exactly where it stands. No reconciliation. No unknown exposure.

> A duplicate order loses capital. A duplicate email is embarrassing. The difference is step 4.

This is durable execution: every workflow step persisted to Postgres. Crash. Read last checkpoint. Resume. Completed steps never re-execute. Survives process death, machine reboot, database failover. Two systems define the category. **Temporal** (2019): centralized server, DoorDash/Snap/Stripe/Uber scale, polyglot SDKs, month-long workflows. Bring your own cluster. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms per step. Bring your own Postgres. Either way: bring step 4.

> Durable execution pays for itself on the first prevented drop.

Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer. Or two engineer-hours — thousands of agent invocations. At scale: 200 deals. 200 natural-language promises. 200 workflows that must survive deploys. Without step 4, the daemon is unreliable. Not because the model is bad. Because the runtime is fragile.

Exactly-once is existential in trading. Durable execution guarantees exactly-once internally. Exchange APIs need idempotency keys. Pattern: checkpoint → send order → crash before checkpointing confirmation → order fires again on recovery → exchange deduplicates by key. This works. Almost no AI agent uses it. Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints. Provenance. Deterministic replay. Durable execution produces all three. For free. One duplicated mid-frequency order can exceed the annual infrastructure budget.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

The pattern is portable. A Go static binary ships the daemon. A Postgres checkpoint ships the state. A container ships both. The daemon runs anywhere Postgres runs. The pattern is composable. Two daemons with different scopes — sales and support — share a Postgres cluster but own separate state ledgers. They observe each other's outputs through the database. No RPC. No message bus. Just state. `sales-daemon | support-daemon | ops-daemon`. Unix pipes for agents.

[Next: the thought experiments that test the pattern's limits.](https://blog.hackspree.com/#durable-daemon-pattern-limits)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*

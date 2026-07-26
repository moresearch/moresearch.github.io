---
title: Durable Daemon Pattern — The Execution
date: 2026-07-26
slug: durable-daemon-pattern-execution
summary: Step 4 is crash-proof execution — the runtime that makes a daemon durable. DBOS and Temporal provide it. Quant trading proves it. The pattern is portable (static binary + Postgres state) and composable (Unix pipes for agents).
tags: daemons, durable-daemon-pattern, dbos, temporal, durable-execution, quant-trading, go
series: durable-daemon-pattern
---

[Previously: the four conditions of the pattern.](https://blog.hackspree.com/#durable-daemon-pattern-definition)

> How does step 4 actually work?

> Checkpoint to Postgres. Crash. Read checkpoint. Resume. The workflow survives.

Durable execution: every step persisted. Two systems. **Temporal** (2019): centralized server. DoorDash, Snap, Stripe, Uber scale. Polyglot SDKs. Month-long workflows. Bring your own cluster. **DBOS** (2023): Mike Stonebraker's library. No server. No queue. No sidecar. 1–2 ms per step. Bring your own Postgres. Either way: bring step 4.

> What does this cost?

> Durable execution pays for itself on the first prevented drop.

A daemon books a meeting. Sends email. Updates CRM. Files Jira. Multi-step. External side effects. Crash after step two. Double email? Meeting booked, CRM not updated? Databases solved this with transactions. Workflow engines with checkpoint-and-compensate. AI agents, mid-2026: neither. Onboarding: verify identity → create account → provision access → send welcome email → schedule call. Crash after step 3. Access granted. No follow-up. Cost: churned customer. Or two engineer-hours. At scale: 200 deals. 200 natural-language promises. 200 workflows surviving deploys. Without step 4, the daemon is unreliable. Not because the model is bad. Because the runtime is fragile.

> Where is the stress-test?

> A duplicate order loses capital. A duplicate email is embarrassing. The difference is step 4.

Quant trading. A trading daemon runs 24/7. State cannot be lost: positions, orders, P&L, risk exposure, regime parameters. It fires orders autonomously. Latency: microseconds to seconds. Failures: dollars. Crash between deciding and confirming an order. Unknown position. Flat or exposed? Without step 4: query the exchange, reconcile positions, resume. Manual minutes. Costly minutes. (Minutes cost more than durable execution.) With step 4: daemon checkpoints *before* and *after*. Recovery: replay checkpoint. Know your position. No reconciliation. Exactly-once is existential. Durable execution: exactly-once internally. Exchange APIs: idempotency keys. Pattern: checkpoint → send → crash before checkpointing confirmation → order fires again → exchange deduplicates. This works. Almost no AI agent uses it. Audit: every order traceable to a decision traceable to state. SEC, CFTC, FCA, ESMA don't accept "the model decided." They accept checkpoints. Provenance. Deterministic replay. Durable execution produces all three. For free. One duplicated order can exceed the annual infrastructure budget.

> You can't run a trading strategy as a stateless loop. You can't run it as an always-on agent. Durable daemon or nothing.

> Why is the pattern portable and composable?

> A static binary ships the daemon. A Postgres checkpoint ships the state. A container ships both. Daemons compose like Unix pipes — one daemon's output is another daemon's trigger.

The four conditions are platform-agnostic. They don't care about your OS, your cloud, or your container runtime. A durable daemon compiled with Go is a single static binary — copy it to a machine, point it at Postgres, run it. The state lives in the database. The daemon is just a process with continuity. Ship the binary. Ship the state. Ship the container that wraps both. The daemon runs anywhere Postgres runs. That's portability.

Composability follows from the same design. A durable daemon has a clean interface: accumulated state in, decisions and actions out, audit trail recorded. Two daemons with different scopes — sales and support, trading and risk, on-call and ops — share a Postgres cluster but own separate state ledgers. They don't call each other's APIs. They observe each other's outputs through the database. A support daemon sees the sales daemon's "Acme stalls Q4" inference and adjusts its own priority. No RPC. No message bus. No coordination server. Just state, observed by daemons with different triggers, producing a composite intelligence greater than any single daemon. This is the Unix pipe model applied to agents: `sales-daemon | support-daemon | ops-daemon`. Each step is a durable daemon with its own scope. The composition is the shared state.

[Next: the thought experiments that test the pattern's limits.](https://blog.hackspree.com/#durable-daemon-pattern-experiments)

---

*Part of the [Durable Daemon Pattern](/tags/durable-daemon-pattern) series.*

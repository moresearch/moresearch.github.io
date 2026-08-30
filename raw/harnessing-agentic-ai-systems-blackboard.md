---
title: "Harnessing Agentic AI Systems: Blackboard Pattern"
date: 2026-08-28
slug: harnessing-agentic-ai-systems-blackboard
summary: "Problem 13 of 15: coordinating through shared state without corruption. The Blackboard (Shared Workspace) pattern vs the State Race Conditions anti-pattern — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, orchestration, shared-state, choreography
series: harnessing-agentic-ai-systems
---

**Problem 13 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Orchestrator-Worker Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker) · Next: [Sequential Pipeline Routing Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-sequential-pipeline-routing).

## The Problem — Coordinating through shared state without corruption

Multiple agents need a shared working memory, and unsynchronized writes corrupt it. The board must be shared and safe at once.

| Field | P14 — Blackboard (Shared Workspace) (pattern) | A12 — State Race Conditions (anti-pattern) |
|---|---|---|
| **Forces / Smell** | Unified schema vs per-agent scopes; parallel writes vs ordered consistency; coordinator vs choreography. | Async multi-agent writes to shared memory with no transactional locks; corruption that appears "random." |
| **Solution / Anti-solution** | Connect disjointed agents to a unified state schema where the harness coordinates simultaneous data writes and events. | "It usually works" — hope as a concurrency strategy. |
| **Consequences / Failure** | Choreography without an orchestrator — "No orchestrator. Just state." Conflicts detected at write time, not merge time (GitButler's virtual branches). | Two agents writing the same ledger, one overwriting the other's checkpoint — the blackboard without its harness. |
| **Tradeoffs / Refactoring** | Shared state is shared risk; the board fights per-agent scope; the resolution is explicit scope discipline — an ownerless board is a tragedy of the commons for state. | P7 with the durability discipline: writes through a single ordered log, exactly-once, idempotency keys, transactional ownership. |
| **Evidence** | CrewAI's Flows ([docs](https://docs.crewai.com/en/concepts/flows)); the durable-daemons choreography ([definition](https://blog.hackspree.com/#durable-daemons-definition)); GitButler's collaboration patterns ([Buzz](https://blog.hackspree.com/#buzz-block-agents)). | Temporal's event-sourced, deterministic execution ([docs](https://docs.temporal.io/)). |
| **Related** | Composes with P7 (the board must be snapshotable) and P10 (the board's slow writers). | Is the absence of P7 and P14's ownership discipline; the async risk of P10. |

## Discussion

The blackboard is choreography made concrete — shared state with no orchestrator. Shared state is shared risk: the board concentrates races, and its durability decides whether corruption is recoverable, which is why snapshots (P7) compose with it. The fix is the durable-daemons discipline: "no RPC. No message bus. No central orchestrator" is only safe when the shared state itself is the coordination mechanism, and the coordination mechanism must be transactional.

## Key Insight

**Shared state is shared risk.** Choreography without an orchestrator requires transactional ownership — writes through a single ordered log, exactly-once, idempotency keys — and the ownership question, which state is shared and who owns the boundary, is the least-solved governance problem in the catalog.

## References

CrewAI Flows ([docs](https://docs.crewai.com/en/concepts/flows)); Temporal ([docs](https://docs.temporal.io/)); archive: [durable daemons series](https://blog.hackspree.com/#durable-daemons), [Buzz and the Identity Problem](https://blog.hackspree.com/#buzz-block-agents), [always-on agents](https://blog.hackspree.com/#always-on-agents).

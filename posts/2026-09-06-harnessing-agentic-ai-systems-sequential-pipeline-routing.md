---
title: "Harnessing Agentic AI Systems: Sequential Pipeline Routing Pattern"
date: 2026-08-28
slug: harnessing-agentic-ai-systems-sequential-pipeline-routing
summary: "Problem 14 of 15: keeping linear flows linear. The Sequential Pipeline Routing pattern — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, orchestration, pipelines, chains
series: harnessing-agentic-ai-systems
---

**Problem 14 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Blackboard Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard) · Next: [Voting / Consensual Ensemble Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).

## The Problem — Keeping linear flows linear

Some flows are genuinely linear — classify, transform, emit — and the simplest verifiable shape is the right one. There is no named anti-pattern; its failure mode — forcing non-linear flows into pipelines — is in the tradeoffs.

| Field | P15 — Sequential Pipeline Routing (pattern) |
|---|---|
| **Forces** | Simplicity vs branching; fixed routing vs recovery; few calls vs specialized hops. |
| **Solution** | Pass analytical payloads through rigid linear stages, using LLMs solely for classification or transformation tasks at each hop. |
| **Consequences** | The easiest harness to verify, replay, and bill: each hop has one job and a deterministic contract — "simple, composable patterns rather than complex frameworks." Each hop is a tool with a `--json` contract, exit codes, and honest `--help`. |
| **Tradeoffs** | Rigid pipelines cannot route around a failed stage; pipelines serialize at the slowest hop, and each hop is a full request. Chains are right when the flow is linear, wrong when it needs to branch, loop, or recover. |
| **Evidence** | LangChain's chains, with the migration-era docs explicit about when the rigid linear form is right ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)). |
| **Related** | Is the disciplined baseline; the opposite of A11 at orchestration scale; composes with P9. |

## Discussion

The pipeline is the disciplined baseline for genuinely linear flows: each hop has one job and a deterministic contract, which makes the system verifiable, replayable, and billable. The boundary is the pattern's own warning: chains are right when the flow is linear, wrong when it needs to branch or recover — forcing non-linear flows into a rigid shape pays the costs (no rerouting, serialization, poisoned stages) without the benefits.

## Key Insight

**Linear flows deserve linear shapes.** The pipeline is the disciplined baseline — and its boundary is explicit: it is the default for the flows that are already linear, not for everything. When the flow needs to branch, loop, or recover, an orchestrator or a graph wins.

## References

LangChain chains ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)); Anthropic, Building Effective Agents ([post](https://www.anthropic.com/engineering/building-effective-agents)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [Loop Engineering](https://blog.hackspree.com/#loop-engineering).

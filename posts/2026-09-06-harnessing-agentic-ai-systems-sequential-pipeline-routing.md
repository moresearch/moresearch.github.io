---
title: "Harnessing Agentic AI Systems: Sequential Pipeline Routing Pattern"
date: 2026-09-06
slug: harnessing-agentic-ai-systems-sequential-pipeline-routing
summary: "Problem 14 of 15 in the Harnessing Agentic AI Systems pattern-language series: keeping linear flows linear. One table — the Sequential Pipeline Routing pattern — followed by the discussion, the key insight, and the problem's references."
tags: harness, pattern-language, agentic-ai, series, orchestration, pipelines, chains
series: harnessing-agentic-ai-systems
---

This is **Problem 14 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern-language series — read the index for the framing and the map. Previous: [Blackboard Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard) · Next: [Voting / Consensual Ensemble Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).

## The Problem — Keeping linear flows linear

Some flows are genuinely linear — classify, transform, emit — and the simplest verifiable shape is the right one. There is no named anti-pattern for this problem; its failure mode — forcing non-linear flows into pipelines — is covered in the tradeoffs.

| Field | P15 — Sequential Pipeline Routing (pattern) |
|---|---|
| **Forces** | Simplicity wants rigid stages; adaptivity wants branching. Determinism wants fixed routing; robustness wants recovery. Billing wants few calls; quality wants specialized hops. |
| **Solution** | Pass analytical payloads through rigid linear stages, using LLMs solely for classification or transformation tasks at each hop. |
| **Consequences** | The easiest harness to verify, replay, and bill: a linear stage where each hop has one job and a deterministic contract, in the spirit of "the most successful implementations use simple, composable patterns rather than complex frameworks." Each hop is a tool with a `--json` contract, exit codes, and honest `--help`. |
| **Tradeoffs** | Rigid pipelines cannot route around failure — a misclassifying stage poisons every downstream stage; pipelines serialize at the slowest hop, and each hop is a full request. Chains are right when the flow is linear, wrong when it needs to branch, loop, or recover. |
| **Evidence** | LangChain's chains, with the migration-era docs explicit about when the rigid linear form is right ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)). |
| **Related** | Is the disciplined baseline; the opposite of A11 at orchestration scale; composes with P9 (each hop's output is schema-enforced). |

## Discussion

The pipeline is the disciplined baseline for genuinely linear flows: each hop has one job and a deterministic contract, which makes the system verifiable, replayable, and billable — the [loop engineering](https://blog.hackspree.com/#loop-engineering) answer to "simple, composable patterns rather than complex frameworks." The boundary is explicit in the pattern's own lineage: chains are right when the flow is linear, wrong when it needs to branch or recover, which is when P13 or a graph shape wins. The pipeline is not the default for everything; it is the default for the flows that are already linear — the failure mode is forcing non-linear flows into a rigid shape, which pays the pipeline's costs (no rerouting, serialization, poisoned stages) without its benefits.

## Key Insight

**Linear flows deserve linear shapes.** The pipeline is the disciplined baseline: each hop has one job and a deterministic contract, which makes the system verifiable, replayable, and billable — the "simple, composable patterns rather than complex frameworks" answer. The boundary is the pattern's own warning: forcing non-linear flows into a rigid shape pays the pipeline's costs — no rerouting, serialization, poisoned stages — without its benefits. Chains are right when the flow is linear, wrong when it needs to branch, loop, or recover.

## References

LangChain chains ([docs](https://python.langchain.com/docs/versions/migrating_chains/overview/)); Anthropic, Building Effective Agents ([post](https://www.anthropic.com/engineering/building-effective-agents)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [Loop Engineering](https://blog.hackspree.com/#loop-engineering), [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag).

Next in the series: [Problem 15 — Voting / Consensual Ensemble Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-voting-ensemble).

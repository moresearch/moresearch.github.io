---
title: "Harnessing Agentic AI Systems: Orchestrator-Worker Pattern"
date: 2026-08-28
slug: harnessing-agentic-ai-systems-orchestrator-worker
summary: "Problem 12 of 15: dividing a workflow across agents. The Orchestrator-Worker pattern, the God Agent anti-pattern, the Sprint Contracts frontier — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, orchestration, delegation, multi-agent
series: harnessing-agentic-ai-systems
---

**Problem 12 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Dynamic Tool Discovery Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-dynamic-tool-discovery) · Next: [Blackboard Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard).

## The Problem — Dividing a workflow across agents

No single context can hold a whole workflow. The work must be divided, delegated, and handed off — and the frontier has added the contract that makes delegation safe (F2).

| Field | P13 — Orchestrator-Worker (pattern) | A11 — The God Agent (anti-pattern) | F2 — Sprint Contracts (frontier) |
|---|---|---|---|
| **Forces / Smell** | Context limits vs a single plan; specialists vs oversight; termination vs conversation. | One massive agent, immense prompt, every phase in one context window. | Specification vs latitude; criteria vs freedom. |
| **Solution / Anti-solution** | Direct traffic using a highly capable central agent that delegates atomic sub-tasks to smaller, faster, specialized agents. | "One agent, all context, total control." | Before each sprint, generator and evaluator negotiate what "done" looks like — and how success will be verified — before any code is written. |
| **Consequences / Failure** | Bounded contexts, parallelizable work, a plan-then-execute shape. | Context avalanche and lost-in-the-middle degradation by construction — the window degrades as it fills. | Ambiguity resolved at the moment of maximum leverage; the agent defines the verifier before the artifact. |
| **Tradeoffs / Refactoring** | The orchestrator is a single point of failure; spec errors cascade; handoffs can lose state; termination must be a system property. | P13 with decomposition and handoff: one feature at a time, structured artifacts between sessions. | Negotiation overhead — bounded by keeping the contract to a sprint-sized chunk. |
| **Evidence** | AutoGen ([paper](https://arxiv.org/abs/2308.08155)); Anthropic's planner-generator-evaluator ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); DeepSeek's subagent registry ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | CrewAI's crews argument ([docs](https://docs.crewai.com/en/concepts/crews)). | Anthropic's harness ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). |
| **Related** | Refactoring for A11; composes with P16 and F2. | Is the absence of P13; composes A4 and A8. | Composes with P13 and F1. |

## Discussion

Delegation is how a system outgrows one context window: the orchestrator curates a delegation tree instead of drowning in context. The single point of failure is the orchestrator — the spec must stay high-level because errors cascade downstream, and termination must be a system property, not the conversation's. The god agent fails by construction; F2 makes each delegation safe by negotiating "done" before the work exists.

## Key Insight

**No single context can hold a workflow.** Delegation is how systems scale, the spec must stay high-level because errors cascade, and termination is a system property, not a conversation's. One window holding every phase degrades exactly as the middle fills.

## References

AutoGen ([arXiv:2308.08155](https://arxiv.org/abs/2308.08155)); CrewAI crews ([docs](https://docs.crewai.com/en/concepts/crews)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents).

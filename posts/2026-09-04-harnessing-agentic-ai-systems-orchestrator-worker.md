---
title: "Harnessing Agentic AI Systems: Orchestrator-Worker Pattern"
date: 2026-09-04
slug: harnessing-agentic-ai-systems-orchestrator-worker
summary: "Problem 12 of 15 in the Harnessing Agentic AI Systems pattern-language series: dividing a workflow across agents. One table — the Orchestrator-Worker pattern, the God Agent anti-pattern, the Sprint Contracts frontier — followed by the discussion, the key insight, and the problem's references."
tags: harness, pattern-language, agentic-ai, series, orchestration, delegation, multi-agent
series: harnessing-agentic-ai-systems
---

This is **Problem 12 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern-language series — read the index for the framing and the map. Previous: [Dynamic Tool Discovery Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-dynamic-tool-discovery) · Next: [Blackboard Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard).

## The Problem — Dividing a workflow across agents

No single context can hold a whole workflow. The work must be divided, delegated, and handed off — and the frontier has added the contract that makes delegation safe (F2).

| Field | P13 — Orchestrator-Worker (pattern) | A11 — The God Agent (anti-pattern) | F2 — Sprint Contracts (frontier) |
|---|---|---|---|
| **Forces / Smell** | Context limits vs a single plan; specialists vs oversight; termination vs conversation. | One massive agent with an immense prompt managing every phase; every tool, rule, and stage in one context window. | Specification wants precision; agility wants latitude. Verification wants criteria; creativity wants freedom. |
| **Solution / Anti-solution** | Direct traffic using a highly capable central agent that delegates atomic sub-tasks to smaller, faster, specialized agents. | "One agent, all context, total control." | Before each sprint, generator and evaluator negotiate a contract: what "done" looks like, and how success will be verified, before any code is written. |
| **Consequences / Failure** | Bounded contexts, parallelizable work, a plan-then-execute shape; the orchestrator curates a delegation tree instead of drowning in context. | The context avalanche and lost-in-the-middle degradation are guaranteed by construction — the window degrades exactly as the middle fills. | The spec's ambiguity is resolved at the moment of maximum leverage — before the work exists; the agent defines the verifier before it defines the artifact. |
| **Tradeoffs / Refactoring** | The orchestrator is a single point of failure; spec errors cascade; handoffs can lose state; termination must be a system property, not the conversation's. | P13 with decomposition and handoff: one feature at a time, structured artifacts carrying state between sessions. | Negotiation overhead — two agents spending tokens agreeing on "done" — bounded by keeping the contract to a sprint-sized chunk. |
| **Evidence** | AutoGen ([paper](https://arxiv.org/abs/2308.08155)); Anthropic's planner-generator-evaluator ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); DeepSeek's subagent registry ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | CrewAI's crews argument: role-specialized agents with focused prompts beat one agent doing everything ([docs](https://docs.crewai.com/en/concepts/crews)). | Anthropic's harness ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). |
| **Related** | Refactoring for A11; composes with P16 (the ensemble as the delegator's judge) and F2. | Is the absence of P13; composes A4 and A8. | Composes with P13 (the delegation's contract) and F1 (the evaluator grades against the contract). |

## Discussion

Delegation is how a system outgrows one context window: the orchestrator curates a delegation tree instead of drowning in context — a swarm of [tiny specialists coordinated by a router](https://blog.hackspree.com/#agents-are-distillation-at-scale). The single point of failure is the orchestrator, which is why the spec must stay high-level (Anthropic's planner deliberately avoids granular detail because "errors in the spec would cascade") and why termination must be a system property. The anti-pattern is the belief that one context can hold everything, falsified by construction — the god agent degrades exactly as its window fills. F2 is the frontier addition that makes each delegation safe: the contract negotiates "done" before the work exists, which is the [tasks-that-fight-back](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) principle applied to the contract itself.

## Key Insight

**No single context can hold a workflow.** Delegation is how systems scale: the orchestrator curates a delegation tree instead of drowning in context, and every handoff carries state through a structured artifact. The spec must stay high-level because errors cascade downstream, and termination must be a system property, not a conversation's. The god agent fails by construction: one window holding every phase degrades exactly as the middle fills.

## References

AutoGen ([arXiv:2308.08155](https://arxiv.org/abs/2308.08155)); CrewAI crews ([docs](https://docs.crewai.com/en/concepts/crews)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (workflow seam, subagent registry), [Agents Aren't Magic. They're Distillation at Scale.](https://blog.hackspree.com/#agents-are-distillation-at-scale), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents).

Next in the series: [Problem 13 — Blackboard Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-blackboard).

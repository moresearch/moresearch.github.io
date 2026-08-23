---
title: "Harnessing Agentic AI Systems: Static Intercepting Gatekeeper Pattern"
date: 2026-08-25
slug: harnessing-agentic-ai-systems-static-intercepting-gatekeeper
summary: "Problem 2 of 15: refusing a tool call before it happens. The Static Intercepting Gatekeeper pattern vs the Prompt-Driven Authorization anti-pattern — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, safety, authorization, interception
series: harnessing-agentic-ai-systems
---

**Problem 2 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Ephemeral Sandbox Wrapper Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-ephemeral-sandbox-wrapper) · Next: [Human-in-the-Loop Breakpoint Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-hitl-breakpoint).

## The Problem — Refusing a tool call before it happens

The system's hands must be able to say no before anything reaches an external API — and the denial must be final: structural, not prose. Interception and authorization are the same seam.

| Field | P2 — Static Intercepting Gatekeeper (pattern) | A3 — Prompt-Driven Authorization (anti-pattern) |
|---|---|---|
| **Forces / Smell** | Security wants denial final; usability wants appeals. Determinism vs adaptivity; audit vs latency. | "Do not delete user data" in the system prompt; permission checks that are sentences, not code. |
| **Solution / Anti-solution** | Intercept model-generated tool calls against a strict blocklist before passing them to external APIs. | Policy as prose — the belief that the model will read and obey the instructions. |
| **Consequences / Failure** | A deterministic, auditable floor that cannot be argued around — the system's `pledge(2)`: a restricted interface where the wrong thing is unexpressible. | Instructions are data; a system prompt is a document the model may be instructed to ignore. |
| **Tradeoffs / Refactoring** | Llama Guard is *not* static — it is a model that can be fooled; a true blocklist catches only what it enumerates; static floor for denial, model judgment above, never below. | Authorization must be monotonic, structural, and fail-closed: monotonic guards "deny or abstain and can never force-allow." |
| **Evidence** | Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); DeepSeek tool pipeline — waterfalls, monotonic guards, `allowed-once` ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); the DeepSeek monotonic-guard doctrine ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). |
| **Related** | Composes with P1 and P3; refactoring for A3. | Is the deeper form of A1; fixed by P2. |

## Discussion

The gatekeeper makes denial a system property rather than a model preference: the static floor does the deterministic denial, and the model-based layer adds judgment *above* it, never below — the DeepSeek ordering doctrine. The anti-pattern inverts the seam: policy as prose asks the model to obey a rule it can be told to ignore, so the refactoring is never a better instruction — it is moving the check into the tool, where injection cannot reach it.

## Key Insight

**Denial must be structural to be final.** The blocklist decides before the model can be persuaded — deny by default, allow by exception — and authorization belongs in the tool, not the prompt: who may modify the system's state is a property of the harness.

## References

Meta's Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); OWASP Top 10 ([2025](https://genai.owasp.org/llm-top-10/)); Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering), [always-on agents](https://blog.hackspree.com/#always-on-agents).

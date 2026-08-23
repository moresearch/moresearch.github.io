---
title: "Harnessing Agentic AI Systems: Static Intercepting Gatekeeper Pattern"
date: 2026-08-25
slug: harnessing-agentic-ai-systems-static-intercepting-gatekeeper
summary: "Problem 2 of 15 in the Harnessing Agentic AI Systems pattern-language series: refusing a tool call before it happens. One table — the Static Intercepting Gatekeeper pattern, the Prompt-Driven Authorization anti-pattern — followed by the discussion, the key insight, and the problem's references."
tags: harness, pattern-language, agentic-ai, series, safety, authorization, interception
series: harnessing-agentic-ai-systems
---

This is **Problem 2 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern-language series — read the index for the framing and the map. Previous: [Ephemeral Sandbox Wrapper Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-ephemeral-sandbox-wrapper) · Next: [Human-in-the-Loop Breakpoint Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-hitl-breakpoint).

## The Problem — Refusing a tool call before it happens

The system's hands must be able to say no before anything reaches an external API — and the denial must be final: structural, not prose. This is the "who may act" problem: interception and authorization are the same seam.

| Field | P2 — Static Intercepting Gatekeeper (pattern) | A3 — Prompt-Driven Authorization (anti-pattern) |
|---|---|---|
| **Forces / Smell** | Security wants denial to be final; usability wants appeals. Determinism vs adaptivity; audit vs latency. | "Do not delete user data" in the system prompt; permission checks that are sentences, not code. |
| **Solution / Anti-solution** | Intercept model-generated tool calls against a strict blocklist before passing them to external APIs. | Policy as prose — the belief that the model will read and obey the instructions. |
| **Consequences / Failure** | A deterministic, auditable floor that cannot be argued around — the system's `pledge(2)`: a restricted interface where the wrong thing is unexpressible. | Instructions are data; prompt injection is the mechanism; a system prompt is a document the model may be instructed to ignore. |
| **Tradeoffs / Refactoring** | Llama Guard is *not* static — it is itself a model that can be fooled or prompted around; a true blocklist catches only what it enumerates; the pattern works when the static floor does the denial and the model adds judgment above, never below. | Authorization must be monotonic, structural, and fail-closed: monotonic guards "deny or abstain and can never force-allow, so owner policy that must not be reordered cannot be argued around." |
| **Evidence** | Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); DeepSeek tool pipeline — waterfalls, monotonic guards, `allowed-once` ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); the DeepSeek monotonic-guard doctrine ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)); [always-on agents](https://blog.hackspree.com/#always-on-agents) (authority axis). |
| **Related** | Composes with P1 (the wrapper contains what the gatekeeper misses) and P3 (the breakpoint is the gatekeeper with a human in it); refactoring for A3. | Is the deeper form of A1; is fixed by P2. |

## Discussion

Interception and authorization are the same seam, and the gatekeeper is where denial becomes a system property rather than a model preference. The honest caveat is that the reference classifier is itself a model — the pattern works when the static floor does the deterministic denial and the model-based layer adds judgment *above* it, never below, which is exactly the DeepSeek ordering doctrine: "pre-execute listeners, approval, and guards must never observe — or worse, approve — a call that can only fail" ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The anti-pattern inverts the seam: policy as prose asks the model to remember and obey a rule it can be told to ignore, and the refactoring is never to write a better instruction — it is to move the check into the tool, where injection cannot reach it. The DSL idea from Fowler's retreat is this problem's constructive form: restrict the vocabulary until the wrong thing is unexpressible ([Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering)).

## Key Insight

**Denial must be structural to be final.** The blocklist decides before the model can be persuaded — deny by default, allow by exception, never let policy be argued around — and authorization belongs in the tool, not the prompt. The systems view names the mechanism: who may modify the system's state is a property of the harness, which is the authority axis of the [always-on survey](https://blog.hackspree.com/#always-on-agents) made concrete. The anti-pattern fails because a system prompt is a document the model may be instructed to ignore; the gatekeeper succeeds because it is a check the model cannot reach.

## References

Meta's Llama Guard ([publication](https://ai.meta.com/research/publications/llama-guard-llm-based-input-output-safeguard-for-human-ai-conversations/)); OWASP Top 10 ([2025](https://genai.owasp.org/llm-top-10/)); Willison's prompt injection series ([series](https://simonwillison.net/series/prompt-injection/)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (monotonic guards, approval seam), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) (DSLs as the pledge(2) bridge), [always-on agents](https://blog.hackspree.com/#always-on-agents) (authority axis).

Next in the series: [Problem 3 — Human-in-the-Loop Breakpoint Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-hitl-breakpoint).

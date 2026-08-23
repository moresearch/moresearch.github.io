---
title: "Harnessing Agentic AI Systems: Rolling Window Compression Pattern"
date: 2026-08-28
slug: harnessing-agentic-ai-systems-rolling-window-compression
summary: "Problem 5 of 15: keeping a long session in a lean window. The Rolling Window Compression pattern, the Context Avalanche anti-pattern, the Context Resets and Context Engineering frontiers — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, memory, context, compression
series: harnessing-agentic-ai-systems
---

**Problem 5 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Token & Time Budget Throttler Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-budget-throttler) · Next: [Semantic Memory Router Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-semantic-memory-router).

## The Problem — Keeping a long session in a lean window

Context is finite and conversation is not; raw dumps degrade reasoning. The frontier has split the answer into three: compression (P5), resets (F3), and the umbrella discipline of context engineering (F6).

| Field | P5 — Rolling Window Compression (pattern) | A4 — The Context Avalanche (anti-pattern) | F3 — Context Resets (frontier) | F6 — Context Engineering (frontier) |
|---|---|---|---|---|
| **Forces / Smell** | Fidelity vs summary; continuity vs clean slate; stable cache prefix vs rewrites. | Raw logs and full transcripts in the history; context near capacity; performance degrading as the session grows. | Continuity vs reset; state survival vs emptied window; cost vs quality. | Right words vs right state; curation vs accumulation. |
| **Solution / Anti-solution** | Automatically summarize older conversation histories in background threads; keep the active window lean. | "Long context solves it" — dump everything. | Distinguish compaction from resets: clear the context and hand state to a fresh agent through a structured artifact; use the reset when the model exhibits context anxiety. | Engineer the whole context state — instructions, tools, MCP servers, data, history — toward the desired behavior. |
| **Consequences / Failure** | Long sessions at bounded cost; a paging discipline — window is RAM, log is disk, summary is the page table. | Lost in the Middle: performance highest at the ends, degrades in the middle; a filled context reads the middle worst exactly when the middle holds the answer. | "A reset provides a clean slate, at the cost of the handoff artifact having enough state for the next agent to pick up the work cleanly." The log never rewrites — the artifact is a new prefix, not an edit. | The unit of design becomes the system's context state — the systems argument in one sentence. |
| **Tradeoffs / Refactoring** | Summary loss is permanent unless the log is preserved; compaction alone does not fix context anxiety; every compaction must be a genuine prefix-extension of the warm request. | P5 over a derived view: 44 event types in the DeepSeek log, exactly three visible to the model. | Resets add orchestration complexity, token overhead, latency; the need is a function of the model generation (Opus 4.5 removed the behavior). | Every refinement risks a cache-prefix violation and a governance gap. |
| **Evidence** | MemGPT ([paper](https://arxiv.org/abs/2310.08560)); DeepSeek's compaction fix ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)); Anthropic's context anxiety ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). | Liu et al., Lost in the Middle ([paper](https://arxiv.org/abs/2307.03172)); DeepSeek's logged-surface invariant ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | Anthropic's harness work ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). | Anthropic, Effective context engineering ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)). |
| **Related** | Refactoring for A4; conflicts with and composes with the append-only log; pairs with F3. | Is the absence of P5 and P8; substrate of A11; retrieval form of A6. | Companion of P5; composes with P13; umbrella is F6. | Umbrella over P5, P6, F3; "the system, not the agent." |

## Discussion

The memory problem contains the catalog's most honest war: the [append-only log](https://blog.hackspree.com/#deepseek-harness) says the past is immutable, compression rewrites it as a projection, tiers store it in parallel — the resolution is layering, and the [compaction prefix bug](https://blog.hackspree.com/#deepseek-harness) is what happens when the layers touch. The boundary is context anxiety: compaction preserves continuity but not a clean slate, and whether you need resets is a function of the model generation. F6 is the umbrella coordinating all three answers against the cache prefix and the governance gap.

## Key Insight

**The past is immutable and the view is derived.** Compression is a projection over the log — if the summarizer rewrites history, the past becomes negotiable and the cache prefix dies. Context anxiety is a model property the system must adapt to; the harness is coupled to the model's psychology, and must be re-examined every time the model changes.

## References

MemGPT ([arXiv:2310.08560](https://arxiv.org/abs/2310.08560)); Liu et al., Lost in the Middle ([arXiv:2307.03172](https://arxiv.org/abs/2307.03172)); Anthropic, Effective context engineering ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)); Anthropic, Harness design for long-running applications ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness).

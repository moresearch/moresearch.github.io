---
title: "Harnessing Agentic AI Systems: Rolling Window Compression Pattern"
date: 2026-08-28
slug: harnessing-agentic-ai-systems-rolling-window-compression
summary: "Problem 5 of 15 in the Harnessing Agentic AI Systems pattern-language series: keeping a long session in a lean window. One table — the Rolling Window Compression pattern, the Context Avalanche anti-pattern, the Context Resets and Context Engineering frontiers — followed by the discussion, the key insight, and the problem's references."
tags: harness, pattern-language, agentic-ai, series, memory, context, compression
series: harnessing-agentic-ai-systems
---

This is **Problem 5 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern-language series — read the index for the framing and the map. Previous: [Token & Time Budget Throttler Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-budget-throttler) · Next: [Semantic Memory Router Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-semantic-memory-router).

## The Problem — Keeping a long session in a lean window

Context is finite and conversation is not; raw dumps degrade reasoning. The system must keep the session inside the window — and the frontier has split the answer into three: compression (P5), resets (F3), and the umbrella discipline of context engineering (F6).

| Field | P5 — Rolling Window Compression (pattern) | A4 — The Context Avalanche (anti-pattern) | F3 — Context Resets (frontier) | F6 — Context Engineering (frontier) |
|---|---|---|---|---|
| **Forces / Smell** | Fidelity vs summary; continuity vs clean slate; stable cache prefix vs rewrites. | Raw logs and full transcripts in the history; context perpetually near capacity; performance degrading as the session grows. | Continuity wants compaction; clarity wants a reset. State wants to survive; the window wants to be emptied. | Prompting wants the right words; context wants the right state. |
| **Solution / Anti-solution** | Automatically summarize older conversation histories in background threads; keep the active window lean. | "Long context solves it" — dump everything and let the model sort it out. | Distinguish compaction from resets: clear the context entirely and hand state to a fresh agent through a structured artifact. Use the reset when the model exhibits context anxiety. | Engineer the whole context state — instructions, tools, MCP servers, data, history — toward "what configuration of context is most likely to generate our model's desired behavior?" |
| **Consequences / Failure** | Long sessions at bounded cost; a paging discipline — window is RAM, log is disk, summary is the page table. | Lost in the Middle: performance highest at the ends, degrades in the middle; a filled context is a degraded one that reads the middle worst exactly when the middle holds the answer. | "A reset provides a clean slate, at the cost of the handoff artifact having enough state for the next agent to pick up the work cleanly." The log never rewrites — the artifact is a new prefix, not an edit. | The unit of design becomes the system's context state — the systems argument in one sentence. |
| **Tradeoffs / Refactoring** | Summary loss is permanent unless the log is preserved; compaction alone does not fix context anxiety; every compaction must be a genuine prefix-extension of the warm request. | P5 over a derived view: 44 event types in the DeepSeek log, exactly three visible to the model. | Resets add orchestration complexity, token overhead, latency; the need is a function of the model generation (Opus 4.5 removed the behavior). | Every refinement risks a cache-prefix violation and a governance gap. |
| **Evidence** | MemGPT ([paper](https://arxiv.org/abs/2310.08560)); DeepSeek's compaction fix ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)); Anthropic's context anxiety ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). | Liu et al., Lost in the Middle ([paper](https://arxiv.org/abs/2307.03172)); DeepSeek's logged-surface invariant ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). | Anthropic's harness work ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)). | Anthropic, Effective context engineering for AI agents ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)). |
| **Related** | Refactoring for A4; conflicts with and composes with the append-only log (first edition); pairs with F3. | Is the absence of P5 and P8; is the substrate of A11 and the retrieval form of A6. | Companion of P5; composes with P13 (the reset is the delegation handoff); umbrella is F6. | Umbrella over P5, P6, F3; the memory problem's frontier expression of "the system, not the agent." |

## Discussion

The memory problem contains the catalog's most honest war: the [append-only log](https://blog.hackspree.com/#deepseek-harness) (first edition) says the past is immutable, compression rewrites it as a projection, and tiers store it in parallel — the resolution is layering, and the [compaction prefix bug](https://blog.hackspree.com/#deepseek-harness) is what happens when the layers touch. The pattern's boundary is context anxiety: compaction preserves continuity but not a clean slate, models "begin wrapping up work prematurely as they approach what they believe is their context limit," and for Sonnet 4.5 compaction alone was insufficient — which is why the frontier added resets, and why the frontier's frontier twist is the most honest statement in this catalog about the model/system boundary: whether resets are needed is a function of the model generation. F6 is the umbrella that coordinates all three answers against the two constraints every one of them must respect — the cache prefix and the governance gap.

## Key Insight

**The past is immutable and the view is derived.** Compression is a projection over the log — the window is RAM, the log is disk, the summary is the page table — and if the summarizer rewrites history, the past becomes negotiable and the cache prefix dies. Context anxiety is a model property the system must adapt to: compaction preserves continuity but not a clean slate, which is why resets exist, and why whether you need them is a function of the model generation. The harness is coupled to the model's psychology, and the harness must be re-examined every time the model changes.

## References

MemGPT ([arXiv:2310.08560](https://arxiv.org/abs/2310.08560)); Liu et al., Lost in the Middle ([arXiv:2307.03172](https://arxiv.org/abs/2307.03172)); Anthropic, Effective context engineering for AI agents ([post](https://www.anthropic.com/engineering/effective-context-engineering-for-ai-agents)); Anthropic, Harness design for long-running application development ([post](https://www.anthropic.com/engineering/harness-design-long-running-apps)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (compaction, logged-surface invariant), [always-on agents](https://blog.hackspree.com/#always-on-agents).

Next in the series: [Problem 6 — Semantic Memory Router Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-semantic-memory-router).

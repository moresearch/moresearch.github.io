---
title: "Harnessing Agentic AI Systems: Dynamic Tool Discovery Pattern"
date: 2026-09-03
slug: harnessing-agentic-ai-systems-dynamic-tool-discovery
summary: "Problem 11 of 15: discovering capabilities without bloating the prompt. The Dynamic Tool Discovery / Registry pattern, the Bloated Utility Belt anti-pattern, the Interop Layer frontier — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, tool-binding, discovery, registry, interop
series: harnessing-agentic-ai-systems
---

**Problem 11 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [Mock Tool Virtualization Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-mock-tool-virtualization) · Next: [Orchestrator-Worker Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-orchestrator-worker).

## The Problem — Discovering capabilities without bloating the prompt

Tool spaces grow, and static lists bloat the prompt and confuse selection. Capabilities must be discoverable — and the frontier has standardized the discovery seam itself (F4).

| Field | P12 — Dynamic Tool Discovery / Registry (pattern) | A8 — Bloated Utility Belt (anti-pattern) | F4 — The Interop Layer (frontier) |
|---|---|---|---|
| **Forces / Smell** | Catalog vs small prompt; reordering vs fixed cache prefix; discovery vs authorization. | Dozens of complex tools per agent; incorrect selection; the model "forgetting" tools exist. | Open protocols vs vendor moats; portability vs control. |
| **Solution / Anti-solution** | Store tool specs in databases, matching and surfacing capabilities dynamically to the agent based on semantic text queries. | "More tools equals more capable." | Adopt the emerging protocol layer: MCP for tools, ACP for editor-agent connection, AGENTS.md for where agent instructions live. |
| **Consequences / Failure** | Many tools behind small prompts; the registry as single source of truth; the type-graph mirror keeps specs from drifting. | Every tool in the prompt is a decision the model must make; a bloated registry turns selection into a needle-in-a-haystack retrieval problem. | The seam becomes the ecosystem; your files become portable, whichever harness you run. |
| **Tradeoffs / Refactoring** | The registry is an injection surface; reordering kills the cache prefix; discovery and authorization are separate seams. | P12 with per-session tool compositions — DeepSeek's presets, and code mode replacing the tool list with a generated SDK. | Every protocol you adopt is a contract you do not control; every standard is a boundary where a substitution can hide. |
| **Evidence** | Toolformer ([paper](https://arxiv.org/abs/2302.04761)); DeepSeek's capability seam ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)); MCP ([architecture](https://modelcontextprotocol.io/docs/learn/architecture)). | Toolformer ([paper](https://arxiv.org/abs/2302.04761)); SWE-agent's ACI ([Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design)). | MCP ([docs](https://modelcontextprotocol.io/docs/learn/architecture)); ACP ([agentclientprotocol.com](https://agentclientprotocol.com/)); AGENTS.md ([agents.md](https://agents.md/)). |
| **Related** | Refactoring for A8; composes with P2 (the gatekeeper authorizes what the registry discovered). | Is the absence of P12; tool-scale form of A11. | Protocol face of P12; interop layer for P13. |

## Discussion

The registry is the capability seam made discoverable: many tools behind small prompts, with the type-graph mirror keeping specs from drifting. The two boundaries are the interesting parts: discovery is not authorization (the gatekeeper still decides), and the tool list must not reorder, or the cache prefix dies — the registry and the bill are the same seam. The anti-pattern is the registry without curation: with the same model, a minimal, well-designed interface more than doubled state-of-the-art ([SWE-agent](https://blog.hackspree.com/#agentic-first-cli-design)).

## Key Insight

**Discovery is not authorization.** Many tools behind small prompts, specs that cannot drift, a tool list that never reorders — and the gatekeeper still decides what the registry surfaced may do. The model should learn *which* tool; the system should not trust it with *all* tools.

## References

Toolformer ([arXiv:2302.04761](https://arxiv.org/abs/2302.04761)); Model Context Protocol ([architecture](https://modelcontextprotocol.io/docs/learn/architecture)); Agent Client Protocol ([agentclientprotocol.com](https://agentclientprotocol.com/)); AGENTS.md ([agents.md](https://agents.md/)); archive: [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness), [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design).

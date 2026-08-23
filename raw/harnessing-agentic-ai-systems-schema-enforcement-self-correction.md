---
title: "Harnessing Agentic AI Systems: Schema Enforcement & Self-Correction Pattern"
date: 2026-08-31
slug: harnessing-agentic-ai-systems-schema-enforcement-self-correction
summary: "Problem 8 of 15: typing tool output and keeping errors legible. The Schema Enforcement & Self-Correction pattern vs the Silent Crash and Schema Free-for-All anti-patterns — one table, a short discussion, the key insight, and the important references."
tags: harness, pattern-language, agentic-ai, series, tool-binding, schemas, errors
series: harnessing-agentic-ai-systems
---

**Problem 8 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) series — read the index for the framing. Previous: [State Snapshot & Rollback Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-state-snapshot-rollback) · Next: [Asynchronous Tool Worker Queue Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-async-tool-worker-queue).

## The Problem — Typing tool output and keeping errors legible

Model output is text; tools and downstream systems need types. Malformed values fail far from their cause, and swallowed errors become hallucinated successes. The boundary must be typed, and failures must stay legible.

| Field | P9 — Schema Enforcement & Self-Correction (pattern) | A7 — The Silent Crash (anti-pattern) | A9 — The Schema Free-for-All (anti-pattern) |
|---|---|---|---|
| **Forces / Smell** | Free text vs types; bounded retries vs bad data; append-only corrections. | Errors caught in the background; blank strings returned; no stderr, no verdict. | Complex arguments as raw strings; parsing deferred; "the model formats it." |
| **Solution / Anti-solution** | Force raw LLM text into JSON Schema, catching parsing failures and feeding structural fixes back internally. | Catch-and-continue: "the agent doesn't need to know." | Trust the model's output format. |
| **Consequences / Failure** | A typed contract at the harness boundary — the same contract as `--json` and exit codes; corrections are appends, not rewrites. | A real error becomes a hallucinated success; destroys the stdout/stderr/exit-code contract. | Parsing moves downstream where no model can correct it; errors surface far from their cause. |
| **Tradeoffs / Refactoring** | Retries cost tokens — the retry budget is part of the pattern; the error message must name the field. | P9 with errors surfaced and bounded; layered checks: "the planner chose the wrong tool," not "the agent failed." | P9 at the boundary with feedback while the model is still in the loop. |
| **Evidence** | Instructor — Pydantic validation with `max_retries` and `token_budget` ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)). | Instructor's retry mechanics ([retrying](https://python.useinstructor.com/concepts/retrying/)). | Pydantic — core validation ([docs](https://docs.pydantic.dev/)). |
| **Related** | Refactoring for A9 and A7; composes with the frozen-request pattern. | Is the absence of P9; feeds A9. | Is the absence of P9; feeds A7's downstream. |

## Discussion

The schema is the grammar of the contract: output is converted from text to structure at the boundary where the model is still in the loop to fix it, and the correction is an append, not a rewrite ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The two anti-patterns are the boundary's two failure directions — the silent crash hides the error (a hallucinated success), and the free-for-all defers parsing until no model is present to correct it. Both are fixed by enforcement at the boundary with legible feedback.

## Key Insight

**The schema is the grammar of the contract.** Malformed output is fixed while the model is still in the loop, the retry budget is part of the pattern (unbounded self-correction is the schema version of the vortex), and swallowed errors become hallucinated successes.

## References

Instructor ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)); Pydantic ([docs](https://docs.pydantic.dev/)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness).

---
title: "Harnessing Agentic AI Systems: Schema Enforcement & Self-Correction Pattern"
date: 2026-08-31
slug: harnessing-agentic-ai-systems-schema-enforcement-self-correction
summary: "Problem 8 of 15 in the Harnessing Agentic AI Systems pattern-language series: typing tool output and keeping errors legible. One table — the Schema Enforcement & Self-Correction pattern, the Silent Crash and Schema Free-for-All anti-patterns — followed by the discussion, the key insight, and the problem's references."
tags: harness, pattern-language, agentic-ai, series, tool-binding, schemas, errors
series: harnessing-agentic-ai-systems
---

This is **Problem 8 of 15** in the [Harnessing Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) pattern-language series — read the index for the framing and the map. Previous: [State Snapshot & Rollback Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-state-snapshot-rollback) · Next: [Asynchronous Tool Worker Queue Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-async-tool-worker-queue).

## The Problem — Typing tool output and keeping errors legible

Model output is text; tools and downstream systems need types. Malformed values fail far from their cause, and swallowed errors become hallucinated successes. The boundary must be typed, and failures must stay legible.

| Field | P9 — Schema Enforcement & Self-Correction (pattern) | A7 — The Silent Crash (anti-pattern) | A9 — The Schema Free-for-All (anti-pattern) |
|---|---|---|---|
| **Forces / Smell** | Free text vs types; bounded retries vs bad data; append-only corrections. | API errors caught in the background; blank or generic strings returned to the agent; no stderr, no exit-code verdict. | Complex arguments passed as raw strings; parsing deferred to the consumer; "the model formats it." |
| **Solution / Anti-solution** | Force raw LLM text into JSON Schema, catching parsing failures and feeding structural fixes back internally. | Catch-and-continue: "the agent doesn't need to know about the error." | Trust the model's output format. |
| **Consequences / Failure** | A typed contract at the harness boundary — the same contract as `--json` and exit codes; validation failures become part of the log: the correction is an append, not a rewrite. | A real error becomes a hallucinated success; swallowing errors destroys the stdout/stderr/exit-code contract. | The parsing problem moves downstream where there is no model to correct it; errors surface far from their cause — the silent crash with more steps. |
| **Tradeoffs / Refactoring** | Retries cost tokens — the retry budget is part of the pattern; strict schemas over-constrain, loose ones under-catch; the error message must name the field. | P9 with errors surfaced and bounded by retry and token budgets; layered checks so failures stay legible — "the planner chose the wrong tool," not "the agent failed." | P9 at the boundary with validation errors fed back while the model is still in the loop, and the corrections recorded in the append-only log. |
| **Evidence** | Instructor — Pydantic validation with `max_retries` and `token_budget` ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)). | Instructor's retry mechanics as the counter-pattern ([retrying](https://python.useinstructor.com/concepts/retrying/)). | Pydantic — core validation ([docs](https://docs.pydantic.dev/)). |
| **Related** | Refactoring for A9 and A7; composes with the frozen-request pattern (first edition). | Is the absence of P9; feeds A9. | Is the absence of P9; feeds A7's downstream. |

## Discussion

The schema is the grammar of the contract: it converts model output from text to structure at the boundary where the model is still in the loop to fix it — and it does so in the append-only spirit, because the correction is a new log entry, not a rewrite ([DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness)). The two anti-patterns are the boundary's two failure directions: the silent crash hides the error (the agent receives `""` or `"ok"` and confidently proceeds on a false premise), and the free-for-all defers the parsing until no model is present to correct it. Both are fixed by the same move — enforcement at the boundary with legible feedback while the model is still in the loop — and both connect to the [agentic-first CLI](https://blog.hackspree.com/#agentic-first-cli-design) contract: the three channels (stdout data, stderr diagnostics, exit code verdict) exist precisely so failures are legible.

## Key Insight

**The schema is the grammar of the contract.** Malformed output is fixed while the model is still in the loop — the correction is an append, not a rewrite — and the retry budget is part of the pattern, not an afterthought: unbounded self-correction is the schema version of the vortex. Swallowed errors become hallucinated successes: the agent receives `""` or `"ok"` and confidently proceeds on a false premise. The three channels — stdout data, stderr diagnostics, exit code verdict — exist precisely so failures are legible, and the harness must keep them legible at the layer where they occurred.

## References

Instructor ([docs](https://python.useinstructor.com/), [retry logic](https://python.useinstructor.com/concepts/retrying/)); Pydantic core validation ([docs](https://docs.pydantic.dev/)); archive: [Agentic-First CLI](https://blog.hackspree.com/#agentic-first-cli-design) (stdout/stderr/exit-code contract), [harness canon](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) (layered checks), [DeepSeek teardown](https://blog.hackspree.com/#deepseek-harness) (frozen request).

Next in the series: [Problem 9 — Asynchronous Tool Worker Queue Pattern](https://blog.hackspree.com/#harnessing-agentic-ai-systems-async-tool-worker-queue).

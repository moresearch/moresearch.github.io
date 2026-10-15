---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal on whether learned selection over structured event histories can improve temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and evaluates it against retrieval and rule-based baselines."
tags: event-sourcing, complex-event-processing, context-engineering, temporal-reasoning, long-context, agents, research-direction, llm-architecture
---

One question: does a learned selector over typed event histories improve temporal question answering under a fixed context-token budget, versus strong retrieval and rule-based baselines?

Long-running agents keep histories of events — decisions, observations, revisions — but hand a language model a flat token string and make it infer what changed and when. This proposal instead assembles context from an explicit event log: deterministic temporal rules constrain the candidates, a learned selector picks the relevant ones, and a frozen Qwen3-0.6B answers from the result. Token attention is unchanged; the change is in how context is selected.

The direct inspiration is IBM Technology's [Context Engineering for AI Agents: 4 Failure Modes to Avoid](https://www.youtube.com/watch?v=BwXremQ-i6I), released today. It names four ways context fails long-running agents — poisoning, distraction, confusion, clash. This idea explores only the temporal slice: retrieval, revision, and answer accuracy.

[![IBM Technology — Context Engineering for AI Agents: 4 Failure Modes to Avoid](https://img.youtube.com/vi/BwXremQ-i6I/hqdefault.jpg)](https://www.youtube.com/watch?v=BwXremQ-i6I)

## The idea

The useful distinction is **representing history vs reasoning over it**. A transcript records what was said; it does not make temporal order, superseded facts, or causal links explicit. An event log does — an ordered record of state changes, where a retraction is a pointer, not a deletion:

```json
{ "id": "evt-0042", "type": "constraint_updated", "t": 1734200000,
  "entity": "network_policy", "payload": "...", "supersedes": "evt-0019" }
```

Three pieces, one contribution:

```text
candidates = filter(log, entity, type, window, reachability)     -- deterministic CEP
score(e,q) = σ( MLP([type_e, type_q, Δt, causal_edge, emb_e, emb_q]) )  -- learned
context    = top_k(candidates by score)
```

- **Event store** — append-only; revision is a pointer.
- **Deterministic filter** — `latest_version(entity)`, time windows, reachability. Fixed, not learned.
- **Learned selector** — relevance ranking only; it does not learn `SEQ`/`NOT`/`COUNT`.
- **Frozen generator** — Qwen3-0.6B reads question + selected events.

So the contribution under test is narrow: learned selection over typed, temporally constrained candidates. (`t` isn't automatically monotonic — pick event time, ingestion time, or a sequence; and end-to-end cost must count candidate retrieval, embeddings, traversal, and serialization, cold-start and repeated-query.)

## Related work and the gap

Structured temporal memory and event retrieval already exist: [APEX-MEM](https://arxiv.org/abs/2604.14362) (event graphs + append-only storage), [Chronos](https://arxiv.org/abs/2603.16862) (timestamped event tuples for time-sensitive QA), ARTEM (event extraction + temporal memory + selective retrieval). Learned selection also exists — sparse attention and differentiable retrieval — and deterministic CEP gives `SEQ`/`NOT`/`WITHIN`/`latest_version` as hand-written rules.

So the gap isn't "events" or "learned selection." It's whether a learned selector, constrained by explicit temporal rules, beats both strong retrieval and a well-designed deterministic query under equal budgets. That's an empirical question, not an assumed gap.

## The experiment

Synthetic temporal QA with facts, updates, and retractions (a preference asserted at `t1`, retracted at `t2`, re-asserted at `t3`; "what is the current policy?").

- Representation: append-only events (entity, type, timestamp/sequence, payload, optional supersession).
- Selector: deterministic filter → small learned scorer.
- Generator: frozen Qwen3-0.6B, question + selected events only.
- Baselines: full context, embedding retrieval, recency, deterministic selection.
- Primary outcome: answer accuracy at a fixed token budget as history grows. Secondary: selection precision/recall, tokens, latency.

Three rules: gold evidence must be independent of the learned model; baselines must see the same information (including a retraction-aware retrieval); and isolate the variable — same schema, rules, generator, prompt, and budget, with indexing overhead counted.

The idea is falsified if a hand-written rule or retraction-aware retrieval matches the selector everywhere, or if full context wins at every affordable length.

## Limitations

Rules may suffice; selection can drop the statement a correction fixes; overhead can erase the token savings. Deferred until the frozen-generator run shows a gap: generator fine-tuning, end-to-end gradients, learned operators, Decision Transformer-style conditioning. Generalization beyond synthetic data is open.

## References

**Event-based memory and temporal retrieval**

- [APEX-MEM: Agentic Semi-Structured Memory with Temporal Reasoning for Long-Term Conversational AI](https://arxiv.org/abs/2604.14362) — temporally grounded event graphs, append-only storage, retrieval-time resolution of evolving information.
- [Chronos: Temporal-Aware Conversational Agents with Structured Event Retrieval for Long-Term Memory](https://arxiv.org/abs/2603.16862) — timestamped event tuples for time-sensitive QA.
- ARTEM (AAAI) — event extraction, temporal memory, and selective retrieval.

**Retrieval and event processing**

- Packer, C. et al. [*MemGPT: Towards LLMs as Operating Systems*](https://arxiv.org/abs/2310.08560) — 2023.
- Fowler, M. [*Event Sourcing*](https://martinfowler.com/eaaDev/EventSourcing.html) — the append-only-log discipline.
- [Apache Flink: Complex Event Processing](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/) — `SEQ`, `NOT`, `WITHIN` patterns over streams.

**Experimental backbone**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — Hugging Face.

**Motivation**

- Keen, M. (IBM Technology). [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I) — released 2026-10-15; the inspiration for this idea.

**Related ideas (future work)**

- Chen, L. et al. [*Decision Transformer: Reinforcement Learning via Sequence Modeling*](https://arxiv.org/abs/2106.01345) — NeurIPS 2021.

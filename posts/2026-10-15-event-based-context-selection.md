---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal on whether learned selection over structured event histories can improve temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and evaluates it against retrieval and rule-based baselines."
tags: event-sourcing, complex-event-processing, context-engineering, temporal-reasoning, long-context, agents, research-direction, llm-architecture
---

The question, narrowed to one sentence: **does a learned selector over typed event histories improve temporal question answering under a fixed context-token budget, compared with strong retrieval and rule-based baselines?**

Long-running agents keep histories of events, decisions, observations, and revisions, but usually hand a language model a flat token string and let it infer what changed and when. The proposal is to assemble that context from an explicit event history instead: deterministic temporal rules constrain the candidates, a learned selector picks the relevant ones, and a frozen generator answers from the result. Token attention stays unchanged; the change is in how context is selected.

This is not a claim that event-based context is inherently better. It is a falsifiable question about the accuracy–cost trade-off.

## Intuition

A long-running agent's history isn't a flat transcript — it's a record of events: decisions, actions, observations, revised facts, commitments. An event's significance depends on its content, when it happened, which entity it touches, what it changes, and how it relates to later events.

That is the difference between **representing history and reasoning over it**. A transcript records what was said; it does not make temporal order, superseded facts, or causal links explicit. Models can infer these from text, and agents can patch around it with summaries or retrieval — but neither gives a uniform way to query an evolving history as typed, temporally related events.

Event-based systems do. They store history as an ordered log of state changes; CEP matches temporal patterns with explicit rules; a learned selector picks the events relevant to a question when relevance is semantic rather than rule-shaped.

So the idea isn't "events instead of attention." It's separating two jobs agents currently entangle: **querying an evolving history**, and **generating from the relevant slice**. Event processing selects; the language model interprets.

[![IBM Technology — Context Engineering for AI Agents: 4 Failure Modes to Avoid](https://img.youtube.com/vi/BwXremQ-i6I/hqdefault.jpg)](https://www.youtube.com/watch?v=BwXremQ-i6I)

IBM's [context-engineering talk](https://www.youtube.com/watch?v=BwXremQ-i6I) names four ways context fails long-running agents — poisoning, distraction, confusion, clash. This proposal evaluates only the temporal slice: retrieval, revision, and answer accuracy.

## Proposed mechanism

Three components, and only one is the contribution.

**Event store** — append-only records. Revision is a pointer, not a deletion:

```json
{ "id": "evt-0042", "type": "constraint_updated", "t": 1734200000,
  "entity": "network_policy", "payload": "...", "supersedes": "evt-0019" }
```

**Deterministic candidate filter** — CEP-style rules (`latest_version(entity)`, time windows, causal reachability) narrow the log to a candidate set. These operators stay fixed.

**Learned selector** — a small model scores each candidate for relevance to the query; the top-k become the context:

```text
candidates = filter(log, entity, type, window, reachability)     -- deterministic
score(e,q) = σ( MLP([type_e, type_q, Δt, causal_edge, emb_e, emb_q]) )  -- learned
context    = top_k(candidates by score)
```

**Frozen generator** — [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) reads the question plus the selected events and answers. Token attention is untouched.

The selector is learned, but it is not a learned CEP engine: it learns relevance ranking only, not `SEQ`/`NOT`/`COUNT`. The contribution under test is **learned selection over typed, temporally constrained candidates** — nothing more.

Two caveats. `t` is not automatically monotonic (out-of-order producers, clock drift): pick event time, ingestion time, or a logical sequence. And scoring is `O(C)` only if each candidate scores in constant time — candidate retrieval, embeddings, causal traversal, and serialization are separate costs, and the end-to-end comparison must include them (cold-start and repeated-query).

## Related work and the gap

Three families already cover much of this, so the novelty claim must be narrow.

**Event-based memory.** [APEX-MEM](https://arxiv.org/abs/2604.14362) builds temporally grounded event graphs with append-only storage and resolves evolving information at retrieval; [Chronos](https://arxiv.org/abs/2603.16862) stores dialogue as timestamped event tuples and retrieves evidence for time-sensitive questions; ARTEM (AAAI) combines event extraction, temporal memory, and selective retrieval. Structured temporal memory and event retrieval are established — they are not the gap.

**Retrieval and learned selection.** [MemGPT](https://arxiv.org/abs/2310.08560) and agent frameworks already retrieve a subset of history by similarity, paging, or recency. Learning to select is also not new: sparse attention and differentiable retrieval do it inside the model.

**Deterministic CEP.** Rule engines and event stores provide `SEQ`, `NOT`, `WITHIN`, `latest_version`, but as hand-written, non-differentiable rules ([Flink CEP](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/), [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html)).

The narrower question: **does a learned event selector, operating over candidates constrained by explicit temporal rules, improve the accuracy–cost trade-off over both strong retrieval and a well-designed deterministic event query, under equal context budgets?** Whether this combination adds value is an empirical question, not an assumed gap.

## Experiment (version 1 — the minimum credible study)

**Task.** Synthetic temporal QA with facts, updates, and retractions — a preference asserted at `t1`, retracted at `t2`, re-asserted at `t3`; "what is the current policy?".

**Representation.** Append-only events: entity, type, timestamp/sequence, payload, optional supersession link.

**Selector.** Deterministic candidate filtering → small learned relevance scorer.

**Generator.** Frozen Qwen3-0.6B, question + selected events only.

**Baselines.** Full context; embedding retrieval; recency selection; deterministic event selection.

**Outcomes.** Primary: answer accuracy at a fixed token budget as history length grows. Secondary: evidence-selection precision/recall, token count, end-to-end latency.

Three design rules make it credible:

1. **Gold evidence independent of the learned model.** A deterministic oracle labels rule-defined cases; semantic or ambiguous cases use controlled generation with known provenance or manually validated labels. Otherwise the router is trained and evaluated against the same rule engine's assumptions.
2. **Baselines see the same information.** Give retrieval the same events and metadata, including a stronger variant that resolves correction chains. Don't make a vector baseline deliberately retraction-blind and then claim revision handling wins.
3. **Isolate the variable.** Same schema, candidate rules, generator, prompt, and token budget across selectors; include indexing and routing overhead in cost.

**Falsification.** If a hand-written rule or a retraction-aware retrieval matches the learned selector everywhere, learning is not justified. If full context wins at every affordable length, selection adds nothing at this scale.

## Limitations and next steps

Three ways this can fail, and what comes after.

- **Rules may suffice.** If deterministic CEP already solves the synthetic task, the learned selector is redundant. The discriminating cases — ambiguous payloads, noisy typing, interleaved entities — are where learning must earn its keep, and they must not leak from training.
- **Selection can drop evidence.** A scorer that picks a correction but omits the statement it corrects fails downstream; temporal dependencies and grouping need explicit handling.
- **Overhead can erase the savings.** Candidate retrieval, embeddings, and serialization may cost more than the tokens saved. Measure end-to-end, cold-start and repeated-query.

Deferred until the frozen-generator experiment shows a gap: fine-tuning the generator, end-to-end gradients through a soft mask (Gumbel-Softmax), learning the temporal operators themselves, and [Decision Transformer](https://arxiv.org/abs/2106.01345)-style outcome conditioning. Generalization beyond synthetic histories is also open.

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

- Keen, M. (IBM Technology). [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I).

**Related ideas (future work)**

- Chen, L. et al. [*Decision Transformer: Reinforcement Learning via Sequence Modeling*](https://arxiv.org/abs/2106.01345) — NeurIPS 2021.

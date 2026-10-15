---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal on whether learned selection over structured event histories can improve temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and evaluates it against retrieval and rule-based baselines."
tags: event-sourcing, temporal-reasoning, retrieval, language-models
---

The question: does a learned selector over typed event histories improve temporal question answering under a fixed context-token budget, versus strong retrieval and rule-based baselines?

Long-running agents keep histories of events but hand a model a flat token string and let it infer what changed and when. This proposal assembles context from an explicit event log instead: deterministic rules constrain the candidates, a learned selector picks the relevant ones, and a frozen generator answers from the result. Token attention is unchanged; the change is in how context is selected.

The motivation is IBM Technology's [Context Engineering for AI Agents](https://www.youtube.com/watch?v=BwXremQ-i6I), released today. It names four ways context fails long-running agents — poisoning, distraction, confusion, clash. This idea tests only the temporal slice: retrieval, revision, and answer accuracy.

## The contribution

One question: does a learned selector improve temporal question answering when deterministic rules have already filtered the event history?

The event store and temporal filter are supporting infrastructure; the learned selector is the component under evaluation. Events are append-only records (type, time, entity, payload, and revision links keyed by fact or attribute — not by entity alone, so updating one fact doesn't discard unrelated facts). Deterministic rules narrow the log to candidates without deciding the answer; a fine-tuned decision model ranks the events needed to answer; a frozen Qwen3-0.6B answers from the selection. The decision model is fine-tuned — Qwen is not.

## The research question

Temporal memory and event retrieval are established: APEX-MEM, Chronos, and ARTEM already combine event structure, temporal reasoning, and selective retrieval, and learned selection exists in sparse attention and differentiable retrieval. Prior work motivates the approach but does not, by itself, establish whether this combination improves performance.

The narrower question: **does a learned selector, constrained by explicit temporal rules, improve the accuracy–cost trade-off over strong retrieval and a well-designed deterministic query under equal context budgets?**

## The experiment

I will evaluate synthetic temporal question answering over event histories containing facts, revisions, and retractions.

**Method.** Question plus typed historical events (timestamps, revision links) go in. Deterministic temporal filters remove ineligible candidates. A fine-tuned decision model ranks the remaining events within a fixed token budget, and a frozen Qwen3-0.6B generates the answer. The decision model is fine-tuned; Qwen stays frozen.

**Backbone.** Start from a pretrained decision model, not an MLP from scratch. [Tev1](https://ollama.com/library/tev1) is a decision model (fine-tuned from Qwen3.5) that selects among predefined options, and its 0.8B variant (~0.8–1.5 GB) fits a CPU-oriented setup. Its published LoRA recipe targets Qwen3.5-4B, so verify the training setup supports the 0.8B base before committing.

One caveat: Tev1 picks a single option, not an evidence *subset* under a token budget. Feasibility test first — on a tiny synthetic dataset, confirm Tev1 selects the relevant event from options, then design the ranking or sequential-selection formulation for subsets.

**Baselines.** Compare full context, embedding retrieval, recency selection, a deterministic revision-aware selector, and the original (un-fine-tuned) Tev1 — under comparable token budgets.

**Evaluation.** Measure answer accuracy as history length increases, evidence-selection precision and recall, actual input tokens, and end-to-end latency. First verify that the generator can answer correctly when given gold evidence.

**Falsification.** The hypothesis is unsupported if the learned selector does not improve the accuracy–cost trade-off over strong baselines. If a deterministic method performs equally well at lower cost, prefer the simpler method.

## Limitations

Rules may suffice; selection can drop the statement a correction fixes; overhead can erase the token savings; and generalization beyond synthetic data is open. Fine-tuning the generator Qwen, end-to-end gradients, and learned operators are deferred until the frozen-generator run shows a gap.

## References

**Related work**

- [APEX-MEM](https://arxiv.org/abs/2604.14362) — temporally grounded event graphs, append-only storage, retrieval-time resolution.
- [Chronos](https://arxiv.org/abs/2603.16862) — timestamped event tuples for time-sensitive QA.
- ARTEM (AAAI) — event extraction, temporal memory, and selective retrieval.

**Experimental backbone**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — Hugging Face.
- [Tev1](https://ollama.com/library/tev1) (0.8B) — fine-tuned decision model; LoRA recipe (Qwen3.5-4B) in the [Together AI repo](https://github.com/togethercomputer/tev1).

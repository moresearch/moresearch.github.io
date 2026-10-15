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

## The research question

**Does a learned selector, constrained by explicit temporal rules, improve the accuracy–cost trade-off over strong retrieval and a well-designed deterministic query, under equal context budgets?**

The pipeline: an append-only event log (type, time, entity, payload, and revision links keyed by fact or attribute — not by entity alone) feeds deterministic filters that narrow the log to candidates; a fine-tuned decision model ranks the events needed to answer; a frozen Qwen3-0.6B answers from the selection. The learned selector is the component under evaluation — the event store and filter are infrastructure, and the decision model is fine-tuned, not Qwen.

Temporal memory and event retrieval already exist (APEX-MEM, Chronos, ARTEM), and learned selection exists in sparse attention and differentiable retrieval. Prior work motivates the approach but does not by itself establish whether this combination improves performance.

## The experiment

I will evaluate synthetic temporal question answering over event histories containing facts, revisions, and retractions.

**Method.** Question plus typed historical events (timestamps, revision links) go in. Deterministic temporal filters remove ineligible candidates. A fine-tuned decision model — trained to identify the evidence useful for each question — ranks the remaining events within a fixed token budget, and a frozen Qwen3-0.6B generates the answer. The decision model is fine-tuned; Qwen stays frozen.

**Selector model.** A pretrained decision model is the starting point; [Tev1](https://ollama.com/library/tev1) is a candidate, not a confirmed backbone. Its standard interface selects a single option, so whether it can be adapted to rank events or select several within a token budget is the key feasibility question — establish that before fine-tuning. (LoRA recipe targets Qwen3.5-4B; verify 0.8B support.)

**Baselines.** Full context, embedding retrieval, recency selection, a deterministic revision-aware selector, and the original (un-fine-tuned) Tev1 through the same selection interface — under comparable token budgets. The Tev1 control separates what fine-tuning adds from what a pretrained decision model already provides.

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

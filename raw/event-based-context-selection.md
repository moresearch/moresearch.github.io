---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal on whether learned selection over structured event histories can improve temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and evaluates it against retrieval and rule-based baselines."
tags: event-sourcing, temporal-reasoning, retrieval, language-models
---

Long-running agents must retrieve relevant information from histories in which facts change over time. Flat-context prompting is one approach, while structured temporal memory and selective retrieval offer alternatives. This proposal tests whether a learned event selector can improve temporal question answering under a fixed context budget, beyond strong retrieval and deterministic baselines.

The motivation is IBM Technology's [Context Engineering for AI Agents](https://www.youtube.com/watch?v=BwXremQ-i6I). It names four ways context fails long-running agents — poisoning, distraction, confusion, clash. This idea tests only the temporal slice: retrieval, revision, and answer accuracy.

## The research question

**Does a learned selector, constrained by explicit temporal rules, improve the accuracy–cost trade-off over strong retrieval and a well-designed deterministic query, under equal context budgets?**

The pipeline: an append-only event log (type, time, entity, payload, and revision links keyed by fact or attribute — not by entity alone) feeds deterministic filters that narrow the log to candidates; a fine-tuned decision model ranks the events needed to answer; a frozen Qwen3-0.6B answers from the selection. The learned selector is the component under evaluation — the event store and filter are infrastructure, and the decision model is fine-tuned, not Qwen.

APEX-MEM (temporally grounded event graphs, retrieval-time resolution), Chronos (structured temporal events, query-guided retrieval), and ARTEM (spatio-temporal episodic memory and retrieval) already cover much of event-based memory and temporal retrieval, and learned selection exists in sparse attention and differentiable retrieval. The contribution here is a controlled comparison of learned evidence selection against strong deterministic and retrieval baselines under a fixed context budget — not the introduction of event-based memory.

## The experiment

I will evaluate synthetic temporal question answering over event histories containing facts, revisions, and retractions.

**Data.** Split by event-history template or scenario, not random questions; hold out combinations of revision, retraction, temporal query, and unrelated updates, so the selector can't memorize patterns.

**Method.** Question plus typed historical events (timestamps, revision links) go in. Deterministic temporal filters remove ineligible candidates. A fine-tuned decision model — trained to identify the evidence useful for each question — ranks the remaining events within a fixed token budget, and a frozen Qwen3-0.6B generates the answer. The decision model is fine-tuned; Qwen stays frozen. Queries distinguish "what is true now" from "what was true at t" so recency alone can't win; for the first experiment, order events by a monotonic sequence number with explicit query-time semantics, deferring occurrence-vs-recording.

**Selector model.** A pretrained decision model is the starting point; [Tev1](https://ollama.com/library/tev1) is a candidate, not a confirmed backbone. Its standard interface selects a single option, so the initial experiment tests single-event selection; multi-event selection under a token budget follows only if that succeeds. (Its LoRA recipe targets Qwen3.5-4B; verify the separate 0.8B variant is supported.)

**Baselines.** Full context, embedding retrieval, recency selection, a deterministic revision-aware selector, and the original (un-fine-tuned) Tev1 through the same selection interface — under comparable token budgets. The deterministic selector resolves revisions per fact or attribute, respects the question's temporal reference, and keeps the evidence needed to interpret a correction or retraction. The Tev1 control separates what fine-tuning adds from what a pretrained decision model already provides.

**Evaluation.** Fix an equal context budget — the same maximum generator-input tokens, same tokenizer, counting question, instructions, and selected events. Measure answer accuracy as history length increases, evidence-selection precision and recall, actual tokens, and end-to-end latency including selector inference, retrieval, and serialization. First verify that the generator can answer correctly when given gold evidence.

**Falsification.** The hypothesis is unsupported if the learned selector does not improve the accuracy–cost trade-off over strong baselines. If a deterministic method performs equally well at lower cost, prefer the simpler method.

## Limitations

Rules may suffice; selection can drop the statement a correction fixes; overhead can erase the token savings; and generalization beyond synthetic data is open. Fine-tuning the generator Qwen, end-to-end gradients, and learned operators are deferred until the frozen-generator run shows a gap.

## References

**Related work**

- [APEX-MEM](https://arxiv.org/abs/2604.14362) — temporally grounded event graphs, append-only storage, retrieval-time resolution.
- [Chronos](https://arxiv.org/abs/2603.16862) — timestamped event tuples for time-sensitive QA.
- ARTEM (AAAI) — spatio-temporal episodic memory and retrieval.

**Experimental backbone**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — Hugging Face.
- [Tev1](https://ollama.com/library/tev1) (0.8B) — fine-tuned decision model; LoRA recipe (Qwen3.5-4B) in the [Together AI repo](https://github.com/togethercomputer/tev1).

---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal on whether learned selection over structured event histories can improve temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and evaluates it against retrieval and rule-based baselines."
tags: event-sourcing, temporal-reasoning, retrieval, language-models
---

An AI agent may need to answer a question using a history of events in which facts change over time. The latest event is not always the relevant one: a change to an account's address does not tell us whether its permissions have changed, and a correction may invalidate an earlier fact.

One solution is to give the model the entire history. Another is to retrieve only the information needed for the current question. This proposal investigates a narrower possibility: can a learned model select better evidence from a structured event history than existing retrieval and rule-based methods, given the same context budget?

The proposal is motivated in part by IBM Technology's [discussion of context engineering for AI agents](https://www.youtube.com/watch?v=BwXremQ-i6I), which identifies poisoning, distraction, confusion, and clash as context failure modes. This study focuses on one specific problem: selecting evidence from histories in which facts are revised, retracted, or superseded.

## The research question

**Can a learned model select evidence that improves temporal question answering, compared with strong retrieval and rule-based methods, under the same context budget?**

The proposed system has three stages.

First, it stores events in an append-only history. Each event records what changed, when it was recorded in the sequence, which entity and attribute it concerns, and whether it revises an earlier event.

Second, deterministic rules remove events that violate the question's temporal constraints. The learned selector chooses among the remaining candidates, which may still contain irrelevant events or competing evidence.

Third, a frozen Qwen3-0.6B model answers using only the selected evidence.

The experiment evaluates the selector, not the storage system or the answer generator. The selector is fine-tuned; Qwen remains unchanged. The central question is whether learned selection adds value beyond deterministic filtering and conventional retrieval.

One technical point is worth preserving: revisions should be linked to the specific fact or attribute they change, not merely to an entity. Otherwise, a change to one attribute could incorrectly displace unrelated information about the same entity.

## Related work

Event-based memory and temporal retrieval are established research directions. [APEX-MEM](https://arxiv.org/abs/2604.14362) explores temporally grounded event graphs and retrieval-time resolution; [Chronos](https://arxiv.org/abs/2603.16862) uses structured temporal events for time-sensitive question answering; and [ARTEM](https://ojs.aaai.org/index.php/AAAI/article/view/39773) investigates spatio-temporal episodic memory and retrieval.

This proposal does not introduce event-based memory or temporal retrieval. It asks a narrower empirical question: after deterministic temporal filtering, does a learned selector improve the quality–cost trade-off compared with strong retrieval and rule-based alternatives?

## The experiment

I will generate synthetic event histories containing facts, revisions, retractions, and unrelated updates. Each history will support questions about both the current state and past states.

**First, test single-event selection.** Begin with questions for which one event uniquely identifies the requested fact. Include distractors that occur later but concern unrelated attributes. For example, if an event changes an allowed port from 80 to 443 and a later event changes the authentication method, the port-change event remains the correct evidence. This initial experiment tests whether the selector can identify relevant evidence without confusing recency with relevance.

**Next, train the selector.** Start from a pretrained decision model and fine-tune it on examples that pair a question and candidate events with the correct evidence. Split the data by underlying history template or scenario, rather than randomly splitting individual questions. Hold out combinations of revisions, retractions, temporal queries, and unrelated updates to test generalization.

**Then, compare against baselines.** Evaluate full-context input when it fits, embedding retrieval, recency selection, a deterministic revision-aware selector, and the original unfine-tuned decision model. When a history exceeds the context limit, specify a truncation strategy for the full-context baseline. Give every method the same maximum generator-input token budget.

The first experiment should establish whether fine-tuning improves single-event selection over both the original model and the deterministic baseline. Only if it does should the study extend to selecting multiple events under a token budget.

**Candidate model.** [Tev1](https://ollama.com/library/tev1) is a candidate because it is a pretrained decision model that selects among predefined options. Its standard interface selects one option, so the initial experiment will use single-event selection. Multi-event selection is a later extension, not an assumption of the initial design. The published LoRA recipe targets Qwen3.5-4B; compatibility with the separate 0.8B variant must be verified before implementation. See the [training repository](https://github.com/togethercomputer/tev1).

**Evaluation.** First verify that Qwen3-0.6B answers correctly when given the gold evidence. Then compare all selectors using the same tokenizer and maximum generator-input token budget, counting the question, instructions, and selected events. Measure answer accuracy as history length increases, evidence-selection precision and recall, actual token use, and end-to-end latency, including retrieval, selection, and serialization. When full context exceeds the budget, use a predefined truncation strategy. Report accuracy and end-to-end cost separately rather than as a single accuracy–cost score.

**Falsification.** The hypothesis is unsupported if fine-tuning does not improve the accuracy–cost trade-off over the original decision model and strong deterministic and retrieval baselines. If a deterministic selector achieves comparable accuracy at lower cost, prefer the deterministic method. If Qwen cannot answer reliably with gold evidence, the generator must be addressed before conclusions about selector quality can be drawn.

## Limitations

Deterministic rules may solve the task without a learned model. A selector may omit an event needed to interpret a correction or retraction. Its inference cost may outweigh the savings from a shorter context. Finally, success on synthetic histories would not establish generalization to real-world agent histories.

The initial study therefore keeps the answer generator frozen and defers generator fine-tuning, end-to-end training, and learned temporal operators until the evidence-selection experiment demonstrates a measurable limitation that justifies them.

## References

**Related research**

- [APEX-MEM](https://arxiv.org/abs/2604.14362) — temporally grounded event graphs and retrieval-time resolution.
- [Chronos](https://arxiv.org/abs/2603.16862) — structured temporal events for time-sensitive question answering.
- [ARTEM](https://ojs.aaai.org/index.php/AAAI/article/view/39773) — spatio-temporal episodic memory and retrieval.

**Implementation resources**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — frozen answer generator.
- [Tev1](https://ollama.com/library/tev1) — candidate pretrained decision model.
- [Tev1 training repository](https://github.com/togethercomputer/tev1) — published fine-tuning recipe.

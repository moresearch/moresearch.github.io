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

Second, deterministic rules remove events that are ineligible under the question's temporal constraints, without resolving all remaining evidence-selection decisions. A learned selector then chooses which remaining events to pass to the language model.

Third, a frozen Qwen3-0.6B model answers using only the selected evidence.

The experiment evaluates the selector, not the storage system or the answer generator. The selector is fine-tuned; Qwen remains unchanged. The central question is whether learned selection adds value beyond deterministic filtering and conventional retrieval.

One technical point is worth preserving: revisions should be linked to the specific fact or attribute they change, not merely to an entity. Otherwise, a change to one attribute could incorrectly displace unrelated information about the same entity.

## Related work

Event-based memory and temporal retrieval are established research directions. [APEX-MEM](https://arxiv.org/abs/2604.14362) explores temporally grounded event graphs and retrieval-time resolution; [Chronos](https://arxiv.org/abs/2603.16862) uses structured temporal events for time-sensitive question answering; and [ARTEM](https://ojs.aaai.org/index.php/AAAI/article/view/39773) investigates spatio-temporal episodic memory and retrieval.

This proposal does not introduce event-based memory or temporal retrieval. It asks a narrower empirical question: after deterministic temporal filtering, does a learned selector improve the quality–cost trade-off compared with strong retrieval and rule-based alternatives?

## The experiment

I will generate synthetic event histories containing facts, revisions, retractions, and unrelated updates. Each history will support questions about both the current state and past states.

**First, test single-event selection.** Begin with questions for which one event is sufficient to identify the answer. Include distractors that are newer than the correct event but concern unrelated attributes. Compare the fine-tuned selector with the original model and a deterministic selector. This isolates whether fine-tuning improves evidence selection before introducing questions that require multiple events.

**Next, train the selector.** Start from a pretrained decision model and fine-tune it on examples that pair a question and candidate events with the correct evidence. Split the data by underlying history template or scenario, rather than randomly splitting individual questions. Hold out combinations of revisions, retractions, temporal queries, and unrelated updates to test generalization.

**Then, compare against baselines.** Evaluate full-context input when it fits, embedding retrieval, recency selection, a deterministic revision-aware selector, and the original unfine-tuned decision model. When a history exceeds the context limit, specify a truncation strategy for the full-context baseline. Give every method the same maximum generator-input token budget.

**Finally, measure the result.** Keep Qwen3-0.6B frozen and first verify that it can answer correctly when given the gold evidence.

The first experiment should establish whether fine-tuning improves single-event selection over both the original model and the deterministic baseline. Only if it does should the study extend to selecting multiple events under a token budget.

**Candidate model.** [Tev1](https://ollama.com/library/tev1) is a candidate because it is a pretrained decision model that selects among predefined options. Its standard interface selects one option, so the initial experiment will use single-event selection. Multi-event selection is a later extension, not an assumption of the initial design. The published LoRA recipe targets Qwen3.5-4B; compatibility with the separate 0.8B variant must be verified before implementation. See the [training repository](https://github.com/togethercomputer/tev1).

**Evaluation.** Use the same tokenizer and maximum generator-input token budget for every method, counting the question, instructions, and selected events. Measure answer accuracy as history length increases, evidence-selection precision and recall, actual token use, and end-to-end latency including retrieval and selector overhead. Report accuracy and end-to-end cost separately, and compare methods across context budgets — a method that is more accurate but slower may still be useful, and one scalar should not hide that.

**Falsification.** The hypothesis is not supported if learned selection fails to improve the accuracy–cost trade-off over strong baselines. If a deterministic selector achieves comparable accuracy at lower cost, the deterministic method is preferable. If the frozen generator cannot answer reliably even with gold evidence, the selection experiment must be revised before drawing conclusions about the selector.

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

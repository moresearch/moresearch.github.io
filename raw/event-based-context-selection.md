---
title: "Event-Based Context Selection for Temporal Reasoning in Language Models"
date: 2026-10-15
slug: event-based-context-selection
summary: "An exploratory research proposal testing whether learned selection over structured event histories improves temporal question answering under fixed context budgets. The approach combines deterministic temporal filtering with a learned event selector and compares it against retrieval and rule-based baselines."
tags: event-sourcing, temporal-reasoning, retrieval, language-models
---

An AI agent may need to answer a question using a history of events in which facts change over time. The latest event is not always the relevant one: a change to an account's address does not tell us whether its permissions have changed, and a correction may invalidate an earlier fact.

One solution is to give the model the entire history. Another is to retrieve only the information needed for the current question. This proposal investigates a narrower possibility: can a learned model select better evidence from a structured event history than existing retrieval and rule-based methods, given the same context budget?

The proposal is motivated in part by IBM Technology's [discussion of context engineering for AI agents](https://www.youtube.com/watch?v=BwXremQ-i6I), which identifies poisoning, distraction, confusion, and clash as context failure modes. This study focuses on one specific problem: selecting evidence from histories in which facts are revised, retracted, or superseded.

## The research question

**Can a learned model select evidence that improves temporal question answering, compared with strong retrieval and rule-based methods, under the same context budget?**

The proposed system has three stages.

First, it stores events in an append-only history. Each event records what changed, its sequence position, the entity and attribute it concerns, and whether it revises or retracts an earlier event. Revisions are linked to the specific fact or attribute they change, rather than merely to an entity. This prevents a change to one attribute from incorrectly displacing unrelated information about the same entity.

Second, deterministic rules remove events that are temporally ineligible for the question. Each query specifies a reference point in the event sequence. For a historical query, events recorded after that point are excluded; for a current-state query, the full history is considered. Revisions and retractions affect the interpretation of a fact according to when they occur, so a later correction does not retroactively change what was recorded at an earlier reference point. The learned selector then chooses among the remaining candidates, which may still contain irrelevant events or competing evidence.

Third, a frozen Qwen3-0.6B model answers using only the selected evidence.

The experiment evaluates the selector, not the storage system or the answer generator. The selector is fine-tuned; Qwen remains unchanged. The central question is whether learned selection adds value beyond deterministic filtering and conventional retrieval.

## Related work

Event-based memory and temporal retrieval are established research directions. [APEX-MEM](https://arxiv.org/abs/2604.14362) explores temporally grounded event graphs and retrieval-time resolution; [Chronos](https://arxiv.org/abs/2603.16862) uses structured temporal events for time-sensitive question answering; and [ARTEM](https://ojs.aaai.org/index.php/AAAI/article/view/39773) investigates spatio-temporal episodic memory and retrieval.

This proposal does not introduce event-based memory or temporal retrieval. It asks a narrower empirical question: after deterministic temporal filtering, does a learned selector improve evidence selection and the accuracy–cost trade-off compared with strong retrieval and rule-based alternatives?

## The experiment

I will generate synthetic event histories containing facts, revisions, retractions, and unrelated updates. Each history will support questions about both the current state and past states. Each question will specify its reference point in the event sequence. The benchmark will define the temporal eligibility rules and retain gold evidence labels, allowing selection quality to be measured independently of answer quality.

**First, test single-event selection.** Begin with questions for which one event uniquely identifies the requested fact. Include newer distractors that concern unrelated attributes. For example, if an event changes an allowed port from 80 to 443 and a later event changes the authentication method, the port-change event remains the correct evidence for a question about the port. This initial benchmark tests whether the selector can identify relevant evidence without confusing recency with relevance. The correct event must always be present among the candidates.

**Next, train the selector.** Start from a pretrained decision model and fine-tune it on examples pairing a question and candidate events with the correct evidence. Each candidate option represents an event identified by its event ID. The model selects the ID of the event that provides sufficient evidence to answer the question. Split the data by underlying history template or scenario, rather than randomly splitting individual questions. Hold out combinations of revisions, retractions, temporal queries, and unrelated updates to test generalization.

**Then, compare against baselines.** Evaluate embedding retrieval, recency selection, a deterministic revision-aware selector, and the original unfine-tuned decision model. Also evaluate full-context input when it fits within the context limit. All selection methods receive the same question and temporally eligible candidate events. Apply the same maximum evidence-token budget to the selection methods, using the same tokenizer. For full-context input, use a predefined truncation strategy when the history exceeds the context limit, and report this baseline separately because it does not perform evidence selection in the same way.

The first experiment should establish whether fine-tuning improves single-event selection over the original decision model and strong deterministic and retrieval baselines. Only if it does should the study extend to selecting multiple events under a token budget.

**Candidate model.** [Tev1](https://ollama.com/library/tev1) is a candidate because it is a pretrained decision model that selects among predefined options. Its standard interface selects one option, making it suitable in principle for the initial single-event experiment. Multi-event selection is a later extension, not an assumption of the initial design. The published LoRA recipe targets Qwen3.5-4B; compatibility with the separate 0.8B variant must be verified before implementation. See the [training repository](https://github.com/togethercomputer/tev1).

**Evaluation.** First verify that frozen Qwen3-0.6B can answer reliably when given the gold evidence. This establishes whether errors in downstream answers can reasonably be attributed to evidence selection. If Qwen cannot answer reliably with gold evidence, the benchmark or generator must be addressed before drawing conclusions about the selector.

Measure evidence-selection accuracy as the primary metric for the single-event task. Also report downstream answer accuracy, evidence-selection precision and recall, actual generator-input tokens, selector inference time, and end-to-end latency, including retrieval, selection, and serialization. Evaluate performance as history length increases. Count the question, instructions, and selected evidence in the generator-input token budget. Report answer accuracy and computational cost separately rather than combining them into a single score. The main comparison is whether fine-tuning improves evidence selection and downstream answering enough to justify its additional cost.

**Falsification.** The hypothesis is unsupported if fine-tuning does not improve evidence-selection performance or the accuracy–cost trade-off over the original decision model and strong deterministic and retrieval baselines. If a deterministic selector achieves comparable accuracy at lower cost, the deterministic method is preferable. If the frozen generator cannot answer reliably even with gold evidence, the experiment cannot isolate the selector's contribution and must be revised before drawing conclusions.

## Limitations

Deterministic rules may solve the task without a learned model. A selector may omit evidence needed to interpret a correction or retraction. Its inference cost may outweigh the savings from a shorter context. Synthetic histories may also fail to capture the ambiguity, noise, and complexity of real-world agent histories.

The initial study therefore keeps the answer generator frozen and defers generator fine-tuning, end-to-end training, multi-event selection, and learned temporal operators until the single-event experiment demonstrates a measurable limitation that justifies further work.

## References

**Related research**

- [APEX-MEM](https://arxiv.org/abs/2604.14362) — temporally grounded event graphs and retrieval-time resolution.
- [Chronos](https://arxiv.org/abs/2603.16862) — structured temporal events for time-sensitive question answering.
- [ARTEM](https://ojs.aaai.org/index.php/AAAI/article/view/39773) — spatio-temporal episodic memory and retrieval.

**Implementation resources**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — frozen answer generator.
- [Tev1](https://ollama.com/library/tev1) — candidate pretrained decision model.
- [Tev1 training repository](https://github.com/togethercomputer/tev1) — published fine-tuning recipe.

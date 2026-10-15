---
title: "Temporal Reasoning in Language Models: An Exploratory Research Idea on Event-Based Context and Complex Event Processing"
date: 2026-10-15
slug: what-if-attention-ran-on-events-not-tokens
summary: "A research idea: context as an append-only event log, attention as complex event processing. It targets the four context failure modes in agents and is falsifiable with a typed event log and a temporal-retrieval benchmark."
tags: event-sourcing, complex-event-processing, context-engineering, decision-transformer, attention, transformers, state-space-models, mamba, rwkv, temporal-reasoning, long-context, agents, research-direction, llm-architecture
---

LLMs treat context as a **static token string** and attention as "look at everything, score by similarity." Flip both: context is an **append-only event log**; attention is complex event processing — "which past events matter, by a temporal rule."

This is a reframing, not a new architecture. It borrows Decision Transformer's outcome-conditioned idea but applies a different mechanism: learned selection over an explicit event history.

## Intuition

A long-running agent's history isn't a flat transcript — it's a record of events: decisions, actions, observations, revised facts, commitments. An event's significance depends on its content, when it happened, which entity it touches, what it changes, and how it relates to later events.

That is the difference between **representing history and reasoning over it**. A transcript records what was said; it does not make temporal order, superseded facts, or causal links explicit. Models can infer these from text, and agents can patch around it with summaries or retrieval — but neither gives a uniform way to query an evolving history as typed, temporally related events.

Event-based systems do. They store history as an ordered log of state changes; CEP matches temporal patterns with explicit rules; a learned selector picks the events relevant to a question when relevance is semantic rather than rule-shaped.

So the idea isn't "events instead of attention." It's separating two jobs agents currently entangle: **querying an evolving history**, and **generating from the relevant slice**. Event processing selects; the language model interprets.

The open question: **can a model reason more reliably and economically when its context is assembled from a structured stream of relevant events, rather than a growing transcript or generic retrieval?** The answer isn't obvious — event selection can drop evidence, get temporal relations wrong, or lean on brittle rules. It must beat retrieval and context-selection baselines on tasks where order, revision, and relationships genuinely matter.

## Context is an event log, not a buffer

A normal model stores:

```text
[tok_1, tok_2, tok_3, ..., tok_n]
```

An event-based model stores:

```json
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000, "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": null }
```

The log is append-only: you add events, never rewrite the prompt. Appending isn't cheaper by itself — it buys incremental indexing and cached per-event representations, so re-encoding the whole history stops being the default. This is [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html): [the event log is the honest memory](https://blog.hackspree.com/#events-as-the-source-of-truth), [narratives are reconstructions built on top of it](https://blog.hackspree.com/#stories-from-events).

## Attention is complex event processing, not a dot product

Standard attention:

```text
Attention(Q, K, V) = softmax(Q·Kᵀ / √d) · V
```

Event attention:

```text
Attention(e_now) = f({ e_i | Pattern(e_i, ..., e_now) matches })
```

CEP patterns — the vocabulary of [rule engines](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp):

```text
SEQ(A, B) WITHIN 5 minutes      -- B right after A
A CAUSES B                      -- causal edge, not positional proximity
NOT(C) AFTER B                  -- negation, made explicit
COUNT(X) > 3 IN window W        -- aggregation over time
```

The CEP layer emits a *relevant substream* — `SELECT events WHERE type=entitlement AND user=me AND t > last_payment`. That substream feeds the generator. Attention becomes a **router**: it still scores candidates, but it changes what reaches the generator.

## Four failure modes this might mitigate

Keen (IBM) names four ways context breaks long-running agents in [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I). Each maps to a CEP operator.

[![IBM Technology — Context Engineering for AI Agents: 4 Failure Modes to Avoid](https://img.youtube.com/vi/BwXremQ-i6I/hqdefault.jpg)](https://www.youtube.com/watch?v=BwXremQ-i6I)

**1. Poisoning** — bad info becomes "truth." A token buffer can't un-say it; an event log appends a `retraction` and `latest_version(entity)` hides it — if the retraction is recorded correctly.

**2. Distraction** — long histories derail the model. Dense attention reads everything; a router reads only the matched substream.

**3. Confusion** — too many tools. Typed events let the router filter by `type=tool_returned` and causal reachability.

**4. Clash** — conflicting facts. `supersedes` points at the old fact, `causes` orders the chain. Resolving it becomes a rule, not an inference.

The video's fixes — validate, compact, curate, prune — are hand-written CEP rules. The claim is they can be *learned*, not that they're solved. Punchline: agents do better with **less** context, if it's the **right** context.

## Why it might be useful

**1. Explicit time and causality.** Positional encodings only imply order; events can carry timestamps and causal links — if those fields are accurate.

**2. Bounded compute as history grows.** Savings come from incremental indexing and selective retrieval, not the log itself. "Unbounded" is never literal: storage, latency, and budgets stay finite. [MemGPT](https://arxiv.org/abs/2310.08560) approximates this with vector DBs and paging.

**3. Revision, not rewrite.** Append a `retraction`; `latest_version(entity)` hides the stale fact. Forgetting is a query.

**4. Symbolic–neural seam.** CEP rules can be hand-written *or* learned; a small network learns which patterns to listen to.

Framing: agents live in irregular time, not regular timesteps — [the FSM's memory](https://blog.hackspree.com/#every-workflow-is-an-fsm).

## What exists, and the gap

Three strands, each half-built.

**Event processing: operators, not learning.** CEP and rule engines handle time, causality, retraction, aggregation — but rules are hand-written and non-differentiable ([Flink CEP](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/), [RETE](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp)).

**Temporal retrieval: selection, not operators.** [MemGPT](https://arxiv.org/abs/2310.08560), temporal knowledge graphs, and agent frameworks ([LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), [CrewAI](https://docs.crewai.com/)) retrieve a subset of history, but the selector is frozen — cosine similarity, paging, recency. In those implementations a `retraction` may be just another embedding — a question for comparison, not a property of the area.

**Sequence modeling: learning, not an explicit substream.** Temporal point processes ([Transformer Hawkes Process](https://arxiv.org/abs/2002.09291)), [Time2Vec](https://arxiv.org/abs/1907.05321), [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363), [Decision Transformer](https://arxiv.org/abs/2106.01345)/[Decision Mamba](https://arxiv.org/abs/2406.00079), and state-space models ([Mamba](https://arxiv.org/abs/2312.00752), [RWKV](https://arxiv.org/abs/2305.13048)) model irregular sequences. Transformers attend densely; SSMs fold history into a hidden state. They don't expose a retrievable substream.

Gap (a question, not a proven absence): **does a learned, differentiable router that emits an explicit event substream add anything?** This argues yes — to test, not to assume.

This is an *integration*, not a new primitive. Learned sparse attention, differentiable retrieval, and neural-symbolic temporal reasoning are close neighbors; the absence of this exact combination can't be assumed. If a hand-written CEP rule matches the learned router, the idea dies.

## The Decision Transformer reframing

Decision Transformer learned a sequence model over trajectories instead of a policy:

```text
[G_t, S_t, A_t, G_{t+1}, S_{t+1}, A_{t+1}, ...]
```

Predict `A` given a desired `G` ([Chen et al., 2021](https://arxiv.org/abs/2106.01345)). Swap the labels:

| Decision LM | Event LM |
| :--- | :--- |
| Return-to-go `G` | Desired CEP pattern, e.g. `COMPLETED(booking) WITHIN 5m` |
| State `S` | World-snapshot event |
| Action `A` | Emitted event — `say()`, `call_tool()`, `infer()` |
| Trajectory | Event trace |
| Reward | Did the pattern fire? |

One could train an event model the same way. That's inspiration, not validation: outcome-conditioning doesn't teach event selection. The router needs its own objective and labels (below). A differentiable attention version, for later work:

```text
score_ij = MLP([type_i, type_j, Δt_ij, causal_edge_ij]) · (Q_i · K_j)
```

That MLP is the proposal's own scoring mechanism. [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363) is related work on temporal modeling and selection (variable selection, gating, interpretable attention) — not the source of the formula. This attention-level version remains a possible extension, not the first experiment.

## The minimum novel mechanism

Boring part: append-only log, index, scan. Three mechanisms, named: **deterministic CEP** (hand-written rules), **learned router** (ranks/selects events), **differentiable temporal attention** (learned relationship scoring). The initial contribution is the second only.

**Schema.** Six fields:

```json
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000,
  "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": "evt-0019" }
```

`supersedes` is retraction as a pointer, not a deletion. `t` isn't monotonic — out-of-order producers and clock drift — so pick event time, ingestion time, or a logical sequence number.

**Inputs/outputs.** In: log prefix + query. Out: substream `S = {e_1 ... e_k}`, each tagged with a matched operator and weight. The generator reads `S`, not the log.

**Two-stage selection.** Stage one (non-learned): filter by type, entity, time window, causal reachability → candidate set `C`. Stage two (learned):

```text
score(e_i, q) = σ( MLP([type_i, type_q, Δt, causal_edge, payload_emb_i, payload_emb_q]) )
```

Top-k, tag, emit. Hard top-k is discrete — no gradient through it. So the first experiment trains supervised (predict gold relevance; hard top-k only at inference). End-to-end is separate.

Cost: with incremental indexing and cached embeddings, the router avoids re-encoding full history; cost depends on candidate discovery, scoring, and selected-context size. (`O(C)` assumes constant-time scoring; indexing, embedding, and causal traversal are separate.)

**Training.** *Hindsight supervision*: a CEP oracle marks the correct substream; the router imitates. *End-to-end*: generator loss flows back through a soft mask / Gumbel-Softmax. The first runs first.

Pipeline:

```text
[Event Stream] -> [Event Embedder] -> [Temporal Index]      (off-the-shelf)
        |
        v
[Learnable CEP Router] -> selects k events                  (the novel part)
        |
        v
[Small LLM / SSM] -> generates the next event               (already exists)
```

Embedder = content + time + type + causal links ([Time2Vec](https://arxiv.org/abs/1907.05321) + entity embeddings). [Mamba](https://arxiv.org/abs/2312.00752), [Mamba-2](https://arxiv.org/abs/2405.21060), [RWKV](https://arxiv.org/abs/2305.13048) fold the past into hidden state; the router keeps it retrievable. Whether explicitness pays is what's left to show.

## A falsifiable experiment

**Task: Temporal Retrieval QA with Retractions**, two tiers.
- *Control* (deterministic): assert `t1`, retract `t2`, re-assert `t3`. `latest_version(entity)` solves it — checks correctness, not learning.
- *Discriminating*: ambiguous payloads, noisy typing, interleaved entities, structured+semantic constraints.

The test set must contain patterns and combinations not in training — else it measures memorization, not routing.

**Hypothesis.** **H1**: selective event context maintains answer accuracy as history length grows, within a fixed context budget. **H2**: the learned router improves the accuracy–cost trade-off over embedding retrieval, recency selection, and hand-written CEP on the discriminating task.

**Baselines.** Full history (reference, not a guaranteed loser — it may win at short lengths; its cost grows with the log). Embedding retrieval (misses retractions). Recency window (catches recency bias). Hand-written CEP rule (does learning beat writing the rule down?).

**Metrics.** Accuracy; retrieval precision/recall vs the gold substream; context size at matched accuracy. Key curve: accuracy vs log length, equal budgets. Dense attention may win short; the claim is only that the router *holds*. If the hand-written rule matches everywhere, H2 dies: "use CEP, skip the ML."

## A concrete plan: Qwen, frozen

Run it with [`Qwen3-0.6B`](https://huggingface.co/Qwen/Qwen3-0.6B), frozen. Three parts: **event store** (append-only), **temporal event router** (learned scorer + deterministic CEP), **frozen Qwen generator** (question + selected events → answer). Freeze Qwen first to isolate the context's effect.

**Data + baselines.** Synthetic histories, ground-truth labels, difficult negatives (outdated, superseded, or wrong-time events for the same entity). Split by history, not pairs. Run four baselines first: full history, retrieval, rule-based CEP, learned router.

**Train the router** two ways. *A — Qwen as classifier:* sequence-classification head, LoRA if tight; trains relevance decisions, not generation. *B — smaller router:* a compact model over `[query_emb, event_emb, type, entity, Δt, supersedes]`.

Hard top-k at inference. Gap to check: independent scoring can pick a correction and drop the statement it corrects.

**Evaluate** (generator fixed): recall@k, accuracy, temporal consistency, tokens, latency, cost. Sweep length, entities, density, conflicts across budgets. Central question: accuracy up at comparable tokens, or tokens down at comparable accuracy?

**Only then fine-tune the generator** (SFT/LoRA on `(question, events) → answer`). Keep a frozen-Qwen control. Key decision: **don't fine-tune the generator first.**

## Three objections

**A. Tokens are still tokens.** CEP doesn't replace soft attention; it selects over event boundaries. Naive all-pairs is `O(E²)`; the router scores `C` candidates and pays token attention only over `k` selected events.

**B. `SEQ`, `NOT`, `COUNT` are discrete.** Hard rules break backprop. Fix: a neural-symbolic soft mask — sigmoid over the MLP, or Gumbel-Softmax routing.

**C. No benchmark.** Frozen-LLM + event log is standard. Need a benchmark where routing retains accuracy and cuts cost as history grows — not one rigged so dense attention fails.

Caveat: **messy payloads** — embedding `{"JSON": ...}` and natural language in one space is unsolved.

Answer: **two layers.** Token attention *within* an event; event attention *across* events via a soft CEP mask.

## Why this is worth doing

The storage metaphor in [*The Brain Is a Key, Not a Box*](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box) fails three ways — recollection rewrites, nothing located, never powered off. An append-only event log has that shape. [*LLM at Runtime Is an Antipattern*](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) wants a hard wall between compute and storage — that's router vs generator.

If you want agents that remember *when, why, for how long* — not chatbots that re-read a long string — this is the direction to bet on.

## References

**The reframing**

- Chen, L. et al. [*Decision Transformer: Reinforcement Learning via Sequence Modeling*](https://arxiv.org/abs/2106.01345) — NeurIPS 2021. The sequence-model-over-trajectories trick transplanted onto events.
- Huang, S., Hu, J., Yang, Z., Yang, L., and Luo, T. [*Decision Mamba: Reinforcement Learning via Hybrid Selective Sequence Modeling*](https://arxiv.org/abs/2406.00079) — NeurIPS 2024. Decision-model conditioning on a selective state-space backbone. (Distinct from Lv, Q. et al., [*Decision Mamba: A Multi-Grained State Space Model with Self-Evolution Regularization for Offline RL*](https://arxiv.org/abs/2406.05427), also NeurIPS 2024.)

**Event-based in spirit already**

- Gu, A. and Dao, T. [*Mamba: Linear-Time Sequence Modeling with Selective State Spaces*](https://arxiv.org/abs/2312.00752) — 2023.
- Dao, T. and Gu, A. [*Transformers are SSMs*](https://arxiv.org/abs/2405.21060) — Mamba-2, 2024.
- Peng, B. et al. [*RWKV: Reinventing RNNs for the Transformer Era*](https://arxiv.org/abs/2305.13048) — 2023.

**Time, events, and transformers**

- Zuo, S. et al. [*Transformer Hawkes Process*](https://arxiv.org/abs/2002.09291) — ICML 2020. Irregular event sequences with a transformer.
- Kazemi, S. M. et al. [*Time2Vec: Learning a Vector Representation of Time*](https://arxiv.org/abs/1907.05321) — 2019.
- Lim, B. et al. [*Temporal Fusion Transformers for Interpretable Multi-horizon Time Series Forecasting*](https://arxiv.org/abs/1912.09363) — 2019. Related work on temporal modeling and selection: variable selection, gating, and interpretable temporal attention.

**Learned sparse attention, differentiable retrieval, and neural-symbolic temporal reasoning**

- Child, R. et al. [*Generating Long Sequences with Sparse Transformers*](https://arxiv.org/abs/1904.10509) — 2019; Kitaev, N. et al. [*Reformer: The Efficient Transformer*](https://arxiv.org/abs/2001.04451) — 2020; Roy, A. et al. [*Efficient Content-Based Sparse Attention with Routing Transformers*](https://arxiv.org/abs/2003.05997) — 2020. Learned sparse attention, the family closest to the router's selection step.
- Guu, K. et al. [*REALM: Retrieval-Augmented Language Model Pre-Training*](https://arxiv.org/abs/2002.08909) — 2020; Tay, Y. et al. [*Transformer Memory as a Differentiable Search Index*](https://arxiv.org/abs/2202.06991) — 2022. Differentiable retrieval.
- Manhaeve, R. et al. [*DeepProbLog: Neural Probabilistic Logic Programming*](https://arxiv.org/abs/1805.10872) — 2018; Mei, H. et al. [*Neural Datalog Through Time*](https://arxiv.org/abs/2006.06723) — 2020. Neural-symbolic temporal reasoning.

**Memory and selective retrieval**

- Packer, C. et al. [*MemGPT: Towards LLMs as Operating Systems*](https://arxiv.org/abs/2310.08560) — 2023. Selective memory via paging over a vector store.

**Agent frameworks that already treat history as events**

- [LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), and [CrewAI](https://docs.crewai.com/) — serialize history as structured messages/events; the extent to which they treat retraction as negation is an open comparison point, and a learned router must beat this baseline.

**Context engineering and the four failure modes**

- Keen, M. (IBM Technology). [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I) — context poisoning, distraction, confusion, and clash; fixes: validate, compact, curate, prune.

**Implementation**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — Hugging Face. The small, cheap backbone for the event-selection experiment.

**Context and systems**

- Fowler, M. [*Event Sourcing*](https://martinfowler.com/eaaDev/EventSourcing.html) — the append-only-log discipline behind the event log.
- [Apache Flink: Complex Event Processing](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/) — reference CEP vocabulary (`SEQ`, `NOT`, `WITHIN`, patterns over streams).

**This blog, on events and memory**

- [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [On Rule Engines 1: From RETE to MCP](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp), [Every Workflow Is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm), [The Brain Is a Key, Not a Box](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern).

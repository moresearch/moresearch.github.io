---
title: "What If Attention Ran on Events, Not Tokens?"
date: 2026-10-15
slug: what-if-attention-ran-on-events-not-tokens
summary: "A transformer treats context as a static string of tokens and attention as all-to-all similarity. Flip both: context becomes an append-only event log and attention becomes complex event processing — a set of temporal patterns that select which past events matter. It is the same trick Decision Transformer pulled on reinforcement learning, applied to memory, and it is a research direction you can test this weekend with a frozen model and LoRA."
tags: event-sourcing, complex-event-processing, decision-transformer, attention, transformers, state-space-models, mamba, rwkv, temporal-reasoning, long-context, agents, research-direction, llm-architecture
---

Here is the assumption almost nobody questions: an LLM's context is a **static string of tokens**, and attention is the operation that says **"look at everything at once and score it by similarity."** Those two choices are so baked in that we treat them as what a language model *is*, rather than as two design decisions we happened to inherit.

An event-based language model flips both. Context becomes a **stream of state changes over time**, and attention becomes **"which past events matter for this event, according to a temporal rule."** I want to argue that this is not a stretch goal or a different model — it is a *reframing*, the same kind that made Decision Transformer work, and it is close enough to things that already exist that you can prototype the core idea over a weekend.

## Context is an event log, not a buffer

A normal model stores this:

```text
[tok_1, tok_2, tok_3, ..., tok_n]
```

An event-based model stores this:

```json
{ "id": "evt-0042", "type": "tool_returned", "timestamp": 1734200000, "payload": { "price": 320 }, "causes": ["evt-0041"] }
```

The payload can be anything: `user_said: "book a flight"`, `tool_returned: {"price": 320}`, `inference: user_intent = travel`, `retraction: fact X no longer valid`. The log is **append-only**. You never rewrite the prompt — you only add events. That is the property that matters for long context, because you stop paying the cost of re-encoding a hundred thousand tokens on every turn. This is the same instinct as [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html), which this blog has been circling for a while: [the event log is the honest memory](https://blog.hackspree.com/#events-as-the-source-of-truth), and [narratives are reconstructions built on top of it](https://blog.hackspree.com/#stories-from-events).

## Attention is complex event processing, not a dot product

Standard attention:

```text
Attention(Q, K, V) = softmax(Q·Kᵀ / √d) · V
```

Event attention:

```text
Attention(e_now) = f({ e_i | Pattern(e_i, ..., e_now) matches })
```

Instead of "all tokens are equally eligible, weight by similarity," you define **complex event processing (CEP)** patterns — the vocabulary of [stream processing and rule engines](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp):

```text
SEQ(A, B) WITHIN 5 minutes      -- B happened right after A
A CAUSES B                      -- a causal edge, not positional proximity
NOT(C) AFTER B                  -- negation, which transformers are bad at
COUNT(X) > 3 IN window W        -- aggregation over time
```

So when the model needs to answer, it does not scan the whole history. The CEP layer emits a *relevant event substream* — a query like `SELECT events WHERE type=entitlement AND user=current_user AND timestamp > last_payment`. That substream, tiny and temporally coherent, is what actually feeds the generator. Attention becomes a *router* instead of a *scorer*.

## Why it might be useful

Four concrete reasons, each one a thing the current architecture buys at a premium or fakes.

**1. Native time and causality.** Transformers fake time with positional encodings. An event log has real `t0 < t1 < t2`, with intervals and durations. You get "before," "after," and "during" for free instead of teaching them as an emergent trick.

**2. Bounded compute with unbounded history.** CEP engines are built to run over millions of events. You keep the log on disk and materialize only the matched pattern. This is what memory managers like [MemGPT](https://arxiv.org/abs/2310.08560) and temporal knowledge graphs are approximating with vector databases and paging — but with events, selective retrieval is first-class rather than retrofitted.

**3. Revision instead of retraction-by-rewrite.** In a token buffer you cannot un-say something. In an event log you append a `retraction` or `superseded` event, and the rule `latest_version(entity)` hides the stale fact automatically. Selective forgetting becomes a query, not a surgery.

**4. A symbolic–neural hybrid with an obvious seam.** The big one. CEP rules can be hand-written *or* learned. A small network can learn which event patterns to listen to — your "event attention heads" — with different heads specializing in causality, recency, and frequency. This is where the two halves of the field stop being enemies and become layers.

There is a fifth, framing reason. Decision models assume regular timesteps; agent life does not happen at regular timesteps. If you want LLMs that act as **agents over months** rather than **chatbots over minutes**, the event abstraction is closer to what the system actually lives in. I wrote about this divergence in [*Every Workflow Is an FSM*](https://blog.hackspree.com/#every-workflow-is-an-fsm); the event model is the FSM's memory.

## The Decision Transformer trick

Decision Transformer worked because it refused to learn a policy and instead learned a **sequence model over trajectories**. A trajectory is just text:

```text
[G_t, S_t, A_t, G_{t+1}, S_{t+1}, A_{t+1}, ...]
```

where `G` is return-to-go, `S` is state, and `A` is action. You train a GPT to predict the next `A` conditioned on a *desired* `G` — no value functions, no actor-critic. At inference you condition on a high return and the model produces the behavior that gets you there ([Chen et al., 2021](https://arxiv.org/abs/2106.01345)).

The event model does the same trick with the labels swapped:

| Decision LM concept | Event LM equivalent |
| :--- | :--- |
| Return-to-go `G` | Desired future event / CEP pattern to achieve, e.g. `COMPLETED(booking) WITHIN 5m` |
| State `S` | World-snapshot event — current beliefs, tool states |
| Action `A` | Emitted event — `say()`, `call_tool()`, `infer()` |
| Trajectory | Event trace — the full log of one task or session |
| Reward | Pattern-match score — did the desired CEP pattern fire? |

Training is the same as Decision Transformer: collect traces, compute the outcome *in hindsight*, prepend it as conditioning, and predict the next event with causal masking. The conditioning signal is the insight worth stealing — conditioning on the **future you want** turns a generation problem into a selection problem.

And the attention injection is where CEP becomes differentiable. A vanilla event model can leave attention alone and let it *implicitly* learn temporal patterns, because the events are already structured. The interesting version replaces the score:

```text
score_ij = MLP([type_i, type_j, Δt_ij, causal_edge_ij]) · (Q_i · K_j)
```

That MLP learns rules like "if `j` is an observation within 10 seconds, boost the score" — a learnable `SEQ` operator. This is precisely the "temporal bias" move from [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363), and it is the line where event attention stops being a metaphor and becomes a mechanism.

## A minimal experiment you can actually run

You do not need to pre-train from scratch. The cheap version:

1. Take a small base model — Qwen2-1.5B or TinyLlama — freeze it, add LoRA.
2. Build ~10k event traces from your own agent logs. Instrument an existing LangChain/CrewAI-style agent so every LLM call, tool return, and observation is logged as a typed event. Label each trace with its final outcome event.
3. Format each trace as `[GOAL, Event_0, Event_1, ..., Event_N]` and train to predict `Event_{t+1}` given the goal and the prefix.
4. At inference, condition on a *pattern* instead of a prompt: `NOT(payment_failed) AFTER search`, plus the current log, and let it generate the next event that avoids the failure.

Step one in the smallest form: **don't touch attention at all.** Train a Decision Transformer over typed events and measure whether the model already learns temporal rules implicitly. That is the control experiment, and it is a weekend of work. The follow-ups — the temporal-bias MLP, learned CEP heads, differentiable pattern matching — are where the real research is.

The architecture most people converge on looks like this:

```text
[Event Stream] -> [Event Embedder] -> [Temporal Index]
        |
        v
[Learnable CEP Router] -> selects k events via patterns
        |
        v
[Small LLM / SSM] -> generates the next event, not the next token
```

The embedder is content + time + type + causal links ([Time2Vec](https://arxiv.org/abs/1907.05321) plus entity embeddings). The router replaces QK attention over 32k tokens with maybe eight learnable queries like "find the last user goal not yet marked complete." And here is the part I find satisfying: the generator's neighbors already exist. [Mamba](https://arxiv.org/abs/2312.00752), [Mamba-2](https://arxiv.org/abs/2405.21060), and [RWKV](https://arxiv.org/abs/2305.13048) are event-based *in spirit* — they keep a recurrent state updated per event and propagate only when the event is relevant. A selective scan is a continuous, differentiable CEP. [Decision Mamba](https://arxiv.org/abs/2403.19979) has already spliced this lineage with the Decision Transformer lineage; adding typed, irregular-time events is the natural next step.

## What would have to be true

Three hard parts, stated as problems rather than aspirations.

**Learning the patterns.** Hand-written CEP rules do not scale. The model needs to learn its own event patterns from data, which means *differentiable* CEP — a genuinely open problem. Neural Datalog and temporal-logic networks are attempts, not answers. Until this works, the "symbolic–neural hybrid" claim is a proposal, not a result.

**Messy payloads.** Real events mix structured and unstructured parts. How do you embed `{"JSON from an API": ...}` and natural language in the same space so CEP matching still works over both? This is the embedder problem, and it is unsolved in the general case.

**Evaluation.** All our benchmarks are token-based. An event model needs benchmarks where *temporal logic matters* — long-horizon agent tasks with deadlines, retractions, and causally ordered constraints. Without those, you cannot tell whether the model is reasoning about time or just correlating.

The honest version of all of this: you probably would not replace attention entirely. You would have two layers — fast, low-level token attention *within* an event to read its payload, and slow, high-level event attention *across* events via CEP to decide which memories matter. That division of labor is the whole thesis in one sentence.

## Why this is worth doing

This is the part I keep coming back to. The storage metaphor this blog examined in [*The Brain Is a Key, Not a Box*](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box) fails in three specific ways: recollection rewrites, nothing is located, and the store is never powered off. An event log is not a filing cabinet — it is an append-only stream you re-derive state from, which is exactly the shape those three failures *demand*. And the economics point from [*LLM at Runtime Is an Antipattern*](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) — that the fix is a hard wall between what is computed and what is stored — is the same move as separating the CEP router from the generator.

I think the field is converging on this from three directions at once: state-space models that are already event-based in spirit, memory systems that are already doing selective retrieval under the table, and decision models that are already conditioning on the future. The event abstraction just names what they have in common. Whether it is the right abstraction is an empirical question, and it is testable — which is rare for a claim this sweeping. The experiment is small, the base model already exists, and the data is sitting in your agent traces.

If you want agents that remember *when* and *why* and *for how long*, rather than chatbots that re-read a long string, this is the direction I would bet a weekend on.

## References

**The reframing**

- Chen, L. et al. [*Decision Transformer: Reinforcement Learning via Sequence Modeling*](https://arxiv.org/abs/2106.01345) — NeurIPS 2021. The sequence-model-over-trajectories trick this post transplants onto events.
- Ota, T. [*Decision Mamba: Reinforcement Learning via Hybrid Selective Sequence Modeling*](https://arxiv.org/abs/2403.19979) — 2024. Decision-model conditioning on a selective state-space backbone.

**Event-based in spirit already**

- Gu, A. and Dao, T. [*Mamba: Linear-Time Sequence Modeling with Selective State Spaces*](https://arxiv.org/abs/2312.00752) — 2023.
- Dao, T. and Gu, A. [*Transformers are SSMs*](https://arxiv.org/abs/2405.21060) — Mamba-2, 2024.
- Peng, B. et al. [*RWKV: Reinventing RNNs for the Transformer Era*](https://arxiv.org/abs/2305.13048) — 2023.

**Time, events, and transformers**

- Zuo, S. et al. [*Transformer Hawkes Process*](https://arxiv.org/abs/2002.09291) — ICML 2020. Modeling irregular event sequences with a transformer, the temporal-point-process side of the idea.
- Kazemi, S. M. et al. [*Time2Vec: Learning a Vector Representation of Time*](https://arxiv.org/abs/1907.05321) — 2019. Time as a learned embedding rather than a scalar.
- Lim, B. et al. [*Temporal Fusion Transformers for Interpretable Multi-horizon Time Series Forecasting*](https://arxiv.org/abs/1912.09363) — 2019. The temporal-bias move the learnable CEP mask generalizes.

**Memory and selective retrieval**

- Packer, C. et al. [*MemGPT: Towards LLMs as Operating Systems*](https://arxiv.org/abs/2310.08560) — 2023. Approximating selective memory with paging over a vector store.

**Context and systems**

- Fowler, M. [*Event Sourcing*](https://martinfowler.com/eaaDev/EventSourcing.html) — the append-only-log discipline behind the event log.
- [Apache Flink: Complex Event Processing](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/) — a reference CEP vocabulary (`SEQ`, `NOT`, `WITHIN`, patterns over streams).

**This blog, on events and memory**

- [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [On Rule Engines 1: From RETE to MCP](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp), [Every Workflow Is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm), [The Brain Is a Key, Not a Box](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern), [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models).

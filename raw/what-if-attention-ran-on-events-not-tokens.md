---
title: "What If Attention Ran on Events, Not Tokens?"
date: 2026-10-15
slug: what-if-attention-ran-on-events-not-tokens
summary: "A transformer treats context as a static string of tokens and attention as all-to-all similarity. Flip both: context becomes an append-only event log and attention becomes complex event processing — a set of temporal patterns that select which past events matter. It is the same trick Decision Transformer pulled on reinforcement learning, applied to memory, and it is a research direction you can falsify with a typed event log, a frozen base model, and a temporal-retrieval benchmark."
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
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000, "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": null }
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

## What already exists, and the gap

Three strands have each built one half of this, and none of them meet in the middle.

**Event processing has the operators, not the learning.** CEP engines, rule engines, and event-sourced systems already handle irregular time, causality, retraction, and aggregation over millions of events ([Flink CEP](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/), [RETE](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp), [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html)). They know *how* to select events — `SEQ`, `NOT`, `WITHIN`, `latest_version` — but the rules are hand-written and non-differentiable, so nothing in this stack learns.

**Temporal retrieval has the selection, not the operators.** Memory systems like [MemGPT](https://arxiv.org/abs/2310.08560), temporal knowledge graphs, and the agent frameworks that already serialize history as events ([LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), [CrewAI](https://docs.crewai.com/)) all retrieve a subset of history instead of feeding everything. But the selector is frozen — cosine similarity over embeddings, paging, or recency — not a learned temporal pattern. A `retraction` event is just another embedding; nothing in the retrieval path treats it as the negation of an earlier fact.

**Sequence modeling has the learning, not the sparsity.** Temporal point processes ([Transformer Hawkes Process](https://arxiv.org/abs/2002.09291)), time embeddings ([Time2Vec](https://arxiv.org/abs/1907.05321)), [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363), [Decision Transformer](https://arxiv.org/abs/2106.01345) and [Decision Mamba](https://arxiv.org/abs/2403.19979), and state-space models ([Mamba](https://arxiv.org/abs/2312.00752), [RWKV](https://arxiv.org/abs/2305.13048)) already model irregular event sequences with learned, differentiable parameters. But they score *every* event densely or compress history into a fixed-width state; they do not emit an explicit, typed, explainable substream.

The gap is therefore precise: what does not exist is a **learned, differentiable router whose output is an explicit event substream** — a CEP-style pattern match produced by a network rather than a hand-written rule, with `SEQ`, `NOT`, `COUNT`, and retraction as first-class, explainable operators instead of emergent side effects.

Novelty, stated honestly: this is an *integration* problem, not a new primitive. Every component already exists in one of the three strands; the claim is only that joining them is worth doing. And the join is falsifiable — if a hand-written CEP rule already gets everything the learned router gets, the whole proposal dies. That is exactly what the experiment at the end is designed to find out.

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

## The minimum novel mechanism

The proposal splits into a boring part and a novel part, and it is worth being explicit about which is which. The boring part — the append-only log, the index, the time-ordered scan — is ordinary event-sourcing infrastructure and needs no new research. The novel part is a single component: **a differentiable router that turns a query plus the log into a small, typed, explainable substream.**

**Event schema.** Six fields, nothing more:

```json
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000,
  "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": "evt-0019" }
```

`id` is unique; `type` comes from a fixed vocabulary (`user_said`, `tool_returned`, `inference`, `retraction`, `goal`, ...); `t` is a monotonic timestamp; `payload` is structured, unstructured, or both; `causes` holds explicit causal edges; `supersedes` points at the event this one retracts, and is null otherwise. That `supersedes` field is the whole story of selective forgetting — a retraction is a pointer, not a deletion.

**Router inputs and outputs.** In: the log prefix up to now, plus a query — either the event being answered or a goal pattern such as `latest(preference)`. Out: a substream `S = {e_1 ... e_k}` of at most `k` events, each carrying a matched-operator label (`SEQ`, `NOT`, `COUNT`, `retraction`, ...) and a scalar weight. The generator reads `S`, not the log. That is the entire interface.

**Candidate selection is two-stage.** Stage one is non-learned: filter the log by type, entity, time window, and causal reachability down to a candidate set `C` — exactly what an event-store index already returns for free. Stage two is the learned part, the score from the previous section given a sigmoid:

```text
score(e_i, q) = σ( MLP([type_i, type_q, Δt, causal_edge, payload_emb_i, payload_emb_q]) )
```

Take the top-k, attach the operator labels, emit. The MLP is the same "temporal bias" idea, but now used to *select* events rather than gate attention, and the sigmoid turns a match into a probability so the route stays differentiable. This one MLP is the entire machine-learning contribution.

**Training the router.** Two options, in increasing order of ambition. *Hindsight supervision*: on traces where a hand-written CEP pattern — the oracle — marks the correct substream for each query, train the router to imitate that selection. It is supervised and safe, and it teaches the router to reproduce rules it can later relax. *End-to-end*: let the generator's loss flow back through the sigmoid (or Gumbel-Softmax) mask, so the router discovers patterns the oracle never wrote. The first is the weekend version; the second is the actual research.

What stays separate, deliberately: the append-only log, the index, and even the CEP engine used as the training oracle are ordinary infrastructure. The only machine-learning contribution is that scoring MLP and its training signal. Someone who already has an event-sourced agent bolts on one small model — not a new database.

The full pipeline, with the boring half drawn so it can be ignored:

```text
[Event Stream] -> [Event Embedder] -> [Temporal Index]      (off-the-shelf)
        |
        v
[Learnable CEP Router] -> selects k events via patterns     (the novel part)
        |
        v
[Small LLM / SSM] -> generates the next event               (already exists)
```

The embedder is content + time + type + causal links ([Time2Vec](https://arxiv.org/abs/1907.05321) plus entity embeddings). The generator's neighbors already exist — [Mamba](https://arxiv.org/abs/2312.00752), [Mamba-2](https://arxiv.org/abs/2405.21060), and [RWKV](https://arxiv.org/abs/2305.13048) are event-based in spirit, and [Decision Mamba](https://arxiv.org/abs/2403.19979) has already spliced the selective-scan lineage with the Decision Transformer lineage. What none of them do is emit an explicit, explainable substream — which is the one thing the router adds.

## A falsifiable experiment

The weekend demo is not the argument. The argument is a hypothesis that can be wrong, four baselines that can beat it, and a metric that decides. The toy task is the one from the objections: **Temporal Retrieval QA with Retractions.** A user preference is asserted at `t1`, retracted at `t2`, re-asserted at `t3`; the generator must answer "what is the current preference?" — and the correct answer depends on reading the `retraction`, not on which embedding happens to be closest.

**Hypothesis.** A learned event router that emits a sparse, retraction-aware substream will reach *higher answer accuracy at fixed context* than any baseline that does not model retraction explicitly, and it will hold that accuracy as the log grows. The claim is two parts, and each can fail on its own: **H1** — bounded context is sufficient, top-k selected events contain enough to answer; **H2** — learning beats frozen selection, the trained router matches or exceeds a hand-written CEP rule on the same task.

**Baselines.** Four, each aimed at killing one part of the claim. *Full context*: feed the whole log with dense attention — the information ceiling, and the thing that dies at long horizons. *Vector retrieval*: embed events, retrieve top-k by cosine similarity — the agent-framework default, and the thing that misses retractions because a `retraction` is just another vector. *Recency window*: keep the last `k` events — the trivial temporal baseline that catches the recency-bias failure mode. *Hand-written CEP rule*: the symbolic oracle (`latest_version(entity)`, retraction filtering) — the control that asks whether learning adds anything over writing the rule down.

**Metrics.** Answer accuracy on the QA task is the headline. Behind it: retrieval precision/recall against the gold substream the oracle would select — does the router actually pick the right events? — and context size at matched accuracy, which is the efficiency claim. The interesting curve is accuracy *against log length*. If the learned router stays flat while full-context and vector retrieval decay, H1 holds. If the hand-written rule matches the learned router everywhere, H2 is dead and the honest conclusion is "use CEP, skip the ML" — still a result, just not the one this post is rooting for.

## Three objections a reviewer would raise

Written as a formal paper, this has three gaps, and each one has a concrete fix.

**A. The equivalence fallacy — tokens inside an event are still tokens.** CEP does not replace soft attention; it sits above it. Standard transformers cross-attend over tokens because natural-language ambiguity requires sub-symbolic, dense similarity scoring, and that does not stop being true because the tokens are wrapped in an event envelope. The fix is to state the hierarchy explicitly: CEP is a *sparsity mask over event boundaries*, costing `O(E²)` over `E` events, while soft attention over the selected events' payloads still costs `O(N²)` over `N` tokens. The two levels coexist — the win is that the all-to-all pass happens over a handful of selected events, not the whole log.

**B. Differentiability — `SEQ`, `NOT`, and `COUNT` are discrete logic.** Hard rules are non-differentiable, and replacing `Q·Kᵀ` with them breaks backpropagation during end-to-end training. The fix is to frame the router as a *neural-symbolic soft mask*: a sigmoid over the temporal-bias MLP (so a match is a probability, not a boolean), or Gumbel-Softmax routing over log predicates (so the discrete selection keeps a gradient). This is the continuous relaxation of event matching, and the machinery already exists — temporal point processes are exactly "an event of type `X` is more likely at time `t` given this history," and they are differentiable by construction.

**C. Evaluation and baselines — the hypothesis has no benchmark.** Prompting a frozen LLM with an event log is already standard practice in agent frameworks — AutoGen, CrewAI, and LangGraph all serialize history as structured events or messages. Without an empirical result showing learned CEP attention heads beat plain RAG / vector search over the same log, the proposal stays a hypothesis. The fix is a concrete toy task where dense attention demonstrably fails: *Temporal Retrieval QA with Retractions* — "find the user preference updated at `t1`, retracted at `t2`, re-asserted at `t3`." Dense attention loses to recency bias or context limits; the `retraction` operator is the one thing that gets the answer right. It is small enough to run in a weekend, and it is the control experiment the falsifiable setup above is built around.

One caveat survives all three objections: the **messy payload** problem. Real events mix structured and unstructured parts, and how you embed `{"JSON from an API": ...}` and natural language in the same space so CEP matching still works over both is unsolved in general. It is not an objection to the direction — it is the part of the embedder nobody has solved yet.

The three objections land on the same architecture from different directions, and it is the one this post already argued for: **two layers.** Fast, low-level token attention *within* an event to read its payload (A's point), slow, high-level event attention *across* events through a soft, differentiable CEP mask (B's point), measured on a task with retractions (C's point). That division of labor is the whole thesis in one sentence.

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

**Agent frameworks that already treat history as events**

- [LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), and [CrewAI](https://docs.crewai.com/) — serializing agent history as structured events or messages is already default practice; this is the baseline a learned CEP router must beat.

**Context and systems**

- Fowler, M. [*Event Sourcing*](https://martinfowler.com/eaaDev/EventSourcing.html) — the append-only-log discipline behind the event log.
- [Apache Flink: Complex Event Processing](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/) — a reference CEP vocabulary (`SEQ`, `NOT`, `WITHIN`, patterns over streams).

**This blog, on events and memory**

- [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [On Rule Engines 1: From RETE to MCP](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp), [Every Workflow Is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm), [The Brain Is a Key, Not a Box](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern), [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models).

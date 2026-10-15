---
title: "Temporal Reasoning in Language Models: An Exploratory Research Idea on Event-Based Context and Complex Event Processing"
date: 2026-10-15
slug: what-if-attention-ran-on-events-not-tokens
summary: "A research idea: treat context as an append-only event log and attention as complex event processing, not a dot product. It fixes the four context failure modes in agents, reuses the Decision Transformer trick, and can be falsified with a typed event log and a temporal-retrieval benchmark."
tags: event-sourcing, complex-event-processing, context-engineering, decision-transformer, attention, transformers, state-space-models, mamba, rwkv, temporal-reasoning, long-context, agents, research-direction, llm-architecture
---

Most LLMs treat context as a **static string of tokens**, and attention as "look at everything, score by similarity." An event-based model flips both: context is a **stream of state changes over time**, and attention is "which past events matter, by a temporal rule."

This is not a new architecture. It is a reframing — the same one Decision Transformer made for RL — and it is testable in a weekend.

## Context is an event log, not a buffer

A normal model stores:

```text
[tok_1, tok_2, tok_3, ..., tok_n]
```

An event-based model stores:

```json
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000, "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": null }
```

The payload can be anything: `user_said: "book a flight"`, `tool_returned: {"price": 320}`, `inference: user_intent = travel`, `retraction: fact X no longer valid`. The log is **append-only** — you add events, you never rewrite the prompt.

Appending alone is not cheaper. What it buys is the *option* to index incrementally and cache per-event representations, so re-encoding the whole history stops being the default. A storage property enabling a compute property, not the same thing. This is the same instinct as [event sourcing](https://martinfowler.com/eaaDev/EventSourcing.html): [the event log is the honest memory](https://blog.hackspree.com/#events-as-the-source-of-truth), and [narratives are reconstructions built on top of it](https://blog.hackspree.com/#stories-from-events).

## Attention is complex event processing, not a dot product

Standard attention:

```text
Attention(Q, K, V) = softmax(Q·Kᵀ / √d) · V
```

Event attention:

```text
Attention(e_now) = f({ e_i | Pattern(e_i, ..., e_now) matches })
```

Instead of "all tokens eligible, weighted by similarity," you define **complex event processing (CEP)** patterns — the vocabulary of [stream processing and rule engines](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp):

```text
SEQ(A, B) WITHIN 5 minutes      -- B happened right after A
A CAUSES B                      -- a causal edge, not positional proximity
NOT(C) AFTER B                  -- negation, which transformers are bad at
COUNT(X) > 3 IN window W        -- aggregation over time
```

The CEP layer emits a *relevant substream* — a query like `SELECT events WHERE type=entitlement AND user=me AND t > last_payment`. That tiny, temporally coherent substream feeds the generator. Attention becomes a **router**, not a scorer.

## The four failure modes this fixes

IBM's Martin Keen names four ways context breaks long-running agents in [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I). Each maps to a CEP operator.

[![IBM Technology — Context Engineering for AI Agents: 4 Failure Modes to Avoid](https://img.youtube.com/vi/BwXremQ-i6I/hqdefault.jpg)](https://www.youtube.com/watch?v=BwXremQ-i6I)

**1. Context poisoning** — bad information becomes "truth." A token buffer cannot un-say it. An event log appends a `retraction`, and `latest_version(entity)` hides the poison. Negation is first-class.

**2. Context distraction** — long histories derail the model. Dense attention reads everything; the router reads only the matched substream. The generator never sees the noise.

**3. Context confusion** — too many tools hurt accuracy. Events are typed, so the router filters by `type=tool_returned` and causal reachability, keeping only the tools relevant right now.

**4. Context clash** — conflicting facts break reasoning. A flat context holds two versions and no way to say which wins. The event log says it explicitly: `supersedes` points at the old fact, `causes` orders the chain. The clash is resolved before generation.

The video's own fixes — validate, compact, curate, prune — are hand-written CEP rules. This post's claim is that they can be *learned*. And the video's punchline is the same thesis: agents often do better with **less** context, if it is the **right** context.

## Why it might be useful

Four gains, plus one framing reason.

**1. Native time and causality.** Transformers fake time with positional encodings. Events carry real order, intervals, and durations — "before," "after," "during" for free.

**2. Bounded compute, unbounded history.** The savings come from incremental indexing and selective retrieval, not from the log itself. Keep the log on disk, materialize only the match. Memory managers like [MemGPT](https://arxiv.org/abs/2310.08560) approximate this with vector DBs and paging; events make it first-class.

**3. Revision, not rewrite.** In a token buffer you cannot un-say something. In an event log you append a `retraction` and `latest_version(entity)` hides the stale fact. Forgetting becomes a query.

**4. A symbolic–neural seam.** CEP rules can be hand-written *or* learned. A small network learns which patterns to listen to — your "event attention heads," one for causality, one for recency, one for frequency.

Framing reason: agents live in irregular time, not regular timesteps. If you want agents over months rather than chatbots over minutes, the event abstraction matches reality — [the FSM's memory](https://blog.hackspree.com/#every-workflow-is-an-fsm).

## What already exists, and the gap

Three strands each built half of this; they do not obviously meet.

**Event processing has operators, not learning.** CEP engines, rule engines, and event-sourced systems already handle irregular time, causality, retraction, and aggregation over millions of events ([Flink CEP](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/), [RETE](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp)). They know *how* to select — `SEQ`, `NOT`, `WITHIN`, `latest_version` — but the rules are hand-written and non-differentiable.

**Temporal retrieval has selection, not operators.** Memory systems like [MemGPT](https://arxiv.org/abs/2310.08560), temporal knowledge graphs, and agent frameworks ([LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), [CrewAI](https://docs.crewai.com/)) already retrieve a subset of history. But the selector is frozen — cosine similarity, paging, or recency. A `retraction` is just another embedding; nothing treats it as negation.

**Sequence modeling has learning, not an explicit substream.** Temporal point processes ([Transformer Hawkes Process](https://arxiv.org/abs/2002.09291)), time embeddings ([Time2Vec](https://arxiv.org/abs/1907.05321)), [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363), [Decision Transformer](https://arxiv.org/abs/2106.01345) and [Decision Mamba](https://arxiv.org/abs/2403.19979), and state-space models ([Mamba](https://arxiv.org/abs/2312.00752), [RWKV](https://arxiv.org/abs/2305.13048)) model irregular sequences with learned parameters. Transformers attend densely; SSMs fold history into an input-dependent hidden state. None exposes a *retrievable, explicit* substream.

So the gap is a question, not a proven absence: **does a learned, differentiable router that emits an explicit event substream add anything?** This post argues yes — but that is a claim to test, not a gap the citations prove.

Novelty, honestly: this is an *integration*, not a new primitive. Learned sparse attention, differentiable retrieval, and neural-symbolic temporal reasoning are close enough neighbors that the absence of exactly this combination cannot be assumed from the references alone. The join is falsifiable — if a hand-written CEP rule matches the learned router, the idea dies.

## The Decision Transformer trick

Decision Transformer refused to learn a policy; it learned a **sequence model over trajectories**:

```text
[G_t, S_t, A_t, G_{t+1}, S_{t+1}, A_{t+1}, ...]
```

Predict the next `A` given a desired `G`. Condition on high return, get good behavior ([Chen et al., 2021](https://arxiv.org/abs/2106.01345)).

Swap the labels:

| Decision LM concept | Event LM equivalent |
| :--- | :--- |
| Return-to-go `G` | Desired CEP pattern, e.g. `COMPLETED(booking) WITHIN 5m` |
| State `S` | World-snapshot event — beliefs, tool states |
| Action `A` | Emitted event — `say()`, `call_tool()`, `infer()` |
| Trajectory | Event trace |
| Reward | Did the pattern fire? |

Train the same way: collect traces, label the outcome in hindsight, prepend it as conditioning, predict the next event. Conditioning on the future you want turns generation into selection.

The attention injection is where CEP becomes differentiable:

```text
score_ij = MLP([type_i, type_j, Δt_ij, causal_edge_ij]) · (Q_i · K_j)
```

That MLP learns "if `j` is an observation within 10s, boost" — a learnable `SEQ`. This is the temporal-bias move from [Temporal Fusion Transformers](https://arxiv.org/abs/1912.09363).

## The minimum novel mechanism

Split it. The **boring** part — append-only log, index, scan — is ordinary event sourcing. The **novel** part is one component: a differentiable router that turns query plus log into a small, typed, explainable substream.

**Schema.** Six fields:

```json
{ "id": "evt-0042", "type": "tool_returned", "t": 1734200000,
  "payload": { "price": 320 }, "causes": ["evt-0041"], "supersedes": "evt-0019" }
```

`supersedes` is the whole story of forgetting — a retraction is a pointer, not a deletion. Note `t` is not monotonic: producers emit out of order and clocks drift, so the system must pick event time, ingestion time, or a logical sequence number — the question event streaming answered decades ago. The router trusts the order only as far as the log guarantees it.

**Inputs and outputs.** In: log prefix + query (the event being answered, or a goal pattern like `latest(preference)`). Out: substream `S = {e_1 ... e_k}`, each event tagged with a matched operator and a weight. The generator reads `S`, not the log.

**Two-stage selection.** Stage one (non-learned): filter by type, entity, time window, causal reachability → candidate set `C`. The index does this for free, keeping scoring at `O(C)`, not `O(E²)`. Stage two (learned):

```text
score(e_i, q) = σ( MLP([type_i, type_q, Δt, causal_edge, payload_emb_i, payload_emb_q]) )
```

Top-k, tag, emit. The sigmoid gives continuous scores, but hard top-k is discrete — the gradient does not flow through selection. So the initial experiment trains *supervised* (predict gold relevance; hard top-k only at inference). End-to-end is a separate experiment.

**Training.** Two options. *Hindsight supervision*: a hand-written CEP oracle marks the correct substream; the router imitates it. *End-to-end*: generator loss flows back through a soft mask or Gumbel-Softmax. The first is what the initial experiment runs.

The pipeline:

```text
[Event Stream] -> [Event Embedder] -> [Temporal Index]      (off-the-shelf)
        |
        v
[Learnable CEP Router] -> selects k events                  (the novel part)
        |
        v
[Small LLM / SSM] -> generates the next event               (already exists)
```

The embedder is content + time + type + causal links ([Time2Vec](https://arxiv.org/abs/1907.05321) plus entity embeddings). [Mamba](https://arxiv.org/abs/2312.00752), [Mamba-2](https://arxiv.org/abs/2405.21060), and [RWKV](https://arxiv.org/abs/2305.13048) already process history through input-dependent state updates; [Decision Mamba](https://arxiv.org/abs/2403.19979) spliced the selective-scan lineage with Decision Transformer. The distinction: an SSM *folds* the past into hidden state; the router *keeps* it retrievable. Whether that explicitness pays is what is left to show.

## A falsifiable experiment

The weekend demo is not the argument. A hypothesis, four baselines, and a metric are.

**Task: Temporal Retrieval QA with Retractions**, in two tiers.
- *Control tier* (deterministic): preference asserted at `t1`, retracted at `t2`, re-asserted at `t3`. `latest_version(entity)` solves it — checks correctness, not whether learning adds anything.
- *Discriminating tier*: ambiguous payloads whose retraction target must be inferred, noisy typing, interleaved entities, structured+semantic constraints. This is where learning must earn its keep.

**Hypothesis.** A retraction-aware router reaches higher accuracy at fixed context, and holds it as the log grows. Two parts, each falsifiable: **H1** bounded context is sufficient; **H2** learning beats frozen selection — the trained router matches or beats a hand-written CEP rule.

**Baselines.** *Full context* — dense attention, the information ceiling that dies at long horizons. *Vector retrieval* — agent default, misses retractions. *Recency window* — catches recency bias. *Hand-written CEP rule* — does learning beat writing the rule down?

**Metrics.** Accuracy (headline); retrieval precision/recall against the gold substream; context size at matched accuracy. The key curve is accuracy vs log length, all methods under equal context and compute budgets. Dense attention may win on short histories — the claim is only that the router *holds* as the log grows. If the hand-written rule matches the learned router everywhere, H2 is dead: "use CEP, skip the ML." Still a result.

## A concrete plan: Qwen, frozen

The experiment above can be run end-to-end without training a foundation model from scratch. Use [`Qwen3-0.6B`](https://huggingface.co/Qwen/Qwen3-0.6B) — small enough for cheap iteration.

**Three components.** An **event store** (append-only; id, event time, type, entity, payload, causal/supersession links); a **temporal event router** (learned scorer + deterministic CEP rules); and a **frozen Qwen generator** that reads question + selected events and writes the answer. Freeze Qwen first so any change is attributable to the selected context, not the model.

**Dataset and baselines.** Generate synthetic histories with known ground truth. Each example is a query + candidate event + binary relevance label, with *difficult negatives* — outdated, superseded, or wrong-time events for the same entity. Split by history or scenario, not by random pairs, so near-identical histories don't leak across train/test. Run four baselines before training anything: full history, text/embedding retrieval, rule-based CEP, and the learned router. This stage decides whether there is a real problem for learning to solve.

**Train the router** two ways. *A — Qwen as a classifier:* add a sequence-classification head, feed `[question, candidate event]`, output a relevance score; LoRA if compute is tight. This trains relevance decisions, not generation. *B — a smaller router:* embed query and event, then a compact model over `[query_emb, event_emb, type, entity, Δt, supersedes]`. Cheaper than running Qwen per candidate, and a baseline against the Qwen classifier.

Selection is hard top-k at inference — no gradients through it. One known gap to check: independent scoring can select a correction and drop the statement it corrects, so temporal dependencies and event grouping need explicit evaluation.

**Evaluate** with the generator fixed: event recall@k, answer accuracy, temporal consistency, input tokens, latency, total cost. Sweep history length, entity count, event density, and conflicts across several context budgets. The central question: does the router raise accuracy at a comparable token budget, or cut tokens at comparable accuracy?

**Only then fine-tune the generator.** If selection works but Qwen still misreads the selected history, add supervised fine-tuning (or LoRA) on `(question, selected events) → answer` pairs — a different objective from router training. Keep a frozen-Qwen control so you know which component moved the needle.

The one decision that keeps the whole thing defensible: **don't fine-tune the generator first.** Test whether explicit event selection helps with Qwen unchanged. Cleaner experiment, lower cost, stronger first result.

## Three objections a reviewer would raise

**A. Tokens inside an event are still tokens.** CEP does not replace soft attention; it sits above it, selecting over event *boundaries*. Complexity: naive all-pairs event attention is `O(E²)`, but the router scores `C` candidates (`O(C)` plus indexing) and pays token attention only over the `k` selected events.

**B. `SEQ`, `NOT`, `COUNT` are discrete logic.** Hard rules break backprop. Fix: a neural-symbolic soft mask — sigmoid over the temporal-bias MLP (a match is a probability), or Gumbel-Softmax routing. This is the continuous relaxation temporal point processes already know.

**C. No benchmark.** Frozen-LLM + event log is already standard (AutoGen, CrewAI, LangGraph). The claim needs a benchmark where dense attention fails — that is the Temporal Retrieval QA task above.

One caveat survives: **messy payloads.** Embedding `{"JSON": ...}` and natural language in one space so CEP matches both is unsolved.

All three land on one answer: **two layers.** Token attention *within* an event (A), event attention *across* events through a soft CEP mask (B), measured on retraction tasks (C).

## Why this is worth doing

The storage metaphor in [*The Brain Is a Key, Not a Box*](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box) fails three ways: recollection rewrites, nothing is located, the store never powers off. An append-only event log has exactly that shape. And [*LLM at Runtime Is an Antipattern*](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) argues for a hard wall between compute and storage — the same wall between router and generator.

The field is converging from three directions: state-space models that are event-based in spirit, memory systems doing selective retrieval under the table, and decision models conditioning on the future. The event abstraction names the common thread. Whether it is right is empirical and testable.

If you want agents that remember *when, why, and for how long* — not chatbots that re-read a long string — this is the direction to bet on.

## References

**The reframing**

- Chen, L. et al. [*Decision Transformer: Reinforcement Learning via Sequence Modeling*](https://arxiv.org/abs/2106.01345) — NeurIPS 2021. The sequence-model-over-trajectories trick transplanted onto events.
- Ota, T. [*Decision Mamba: Reinforcement Learning via Hybrid Selective Sequence Modeling*](https://arxiv.org/abs/2403.19979) — 2024. Decision-model conditioning on a selective state-space backbone.

**Event-based in spirit already**

- Gu, A. and Dao, T. [*Mamba: Linear-Time Sequence Modeling with Selective State Spaces*](https://arxiv.org/abs/2312.00752) — 2023.
- Dao, T. and Gu, A. [*Transformers are SSMs*](https://arxiv.org/abs/2405.21060) — Mamba-2, 2024.
- Peng, B. et al. [*RWKV: Reinventing RNNs for the Transformer Era*](https://arxiv.org/abs/2305.13048) — 2023.

**Time, events, and transformers**

- Zuo, S. et al. [*Transformer Hawkes Process*](https://arxiv.org/abs/2002.09291) — ICML 2020. Irregular event sequences with a transformer.
- Kazemi, S. M. et al. [*Time2Vec: Learning a Vector Representation of Time*](https://arxiv.org/abs/1907.05321) — 2019.
- Lim, B. et al. [*Temporal Fusion Transformers for Interpretable Multi-horizon Time Series Forecasting*](https://arxiv.org/abs/1912.09363) — 2019. The temporal-bias move the CEP mask generalizes.

**Learned sparse attention, differentiable retrieval, and neural-symbolic temporal reasoning**

- Child, R. et al. [*Generating Long Sequences with Sparse Transformers*](https://arxiv.org/abs/1904.10509) — 2019; Kitaev, N. et al. [*Reformer: The Efficient Transformer*](https://arxiv.org/abs/2001.04451) — 2020; Roy, A. et al. [*Efficient Content-Based Sparse Attention with Routing Transformers*](https://arxiv.org/abs/2003.05997) — 2020. Learned sparse attention, the family closest to the router's selection step.
- Guu, K. et al. [*REALM: Retrieval-Augmented Language Model Pre-Training*](https://arxiv.org/abs/2002.08909) — 2020; Tay, Y. et al. [*Transformer Memory as a Differentiable Search Index*](https://arxiv.org/abs/2202.06991) — 2022. Differentiable retrieval.
- Manhaeve, R. et al. [*DeepProbLog: Neural Probabilistic Logic Programming*](https://arxiv.org/abs/1805.10872) — 2018; Mei, H. et al. [*Neural Datalog Through Time*](https://arxiv.org/abs/2006.06723) — 2020. Neural-symbolic temporal reasoning.

**Memory and selective retrieval**

- Packer, C. et al. [*MemGPT: Towards LLMs as Operating Systems*](https://arxiv.org/abs/2310.08560) — 2023. Selective memory via paging over a vector store.

**Agent frameworks that already treat history as events**

- [LangGraph](https://langchain-ai.github.io/langgraph/), [AutoGen](https://microsoft.github.io/autogen/), and [CrewAI](https://docs.crewai.com/) — serializing history as structured events is default practice; this is the baseline a learned router must beat.

**Context engineering and the four failure modes**

- Keen, M. (IBM Technology). [*Context Engineering for AI Agents: 4 Failure Modes to Avoid*](https://www.youtube.com/watch?v=BwXremQ-i6I) — context poisoning, distraction, confusion, and clash; fixes: validate, compact, curate, prune.

**Implementation**

- [Qwen3-0.6B](https://huggingface.co/Qwen/Qwen3-0.6B) — Hugging Face. The small, cheap backbone for the event-selection experiment.

**Context and systems**

- Fowler, M. [*Event Sourcing*](https://martinfowler.com/eaaDev/EventSourcing.html) — the append-only-log discipline behind the event log.
- [Apache Flink: Complex Event Processing](https://nightlies.apache.org/flink/flink-docs-stable/docs/libs/cep/) — reference CEP vocabulary (`SEQ`, `NOT`, `WITHIN`, patterns over streams).

**This blog, on events and memory**

- [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [On Rule Engines 1: From RETE to MCP](https://blog.hackspree.com/#on-rule-engines-1-rete-to-mcp), [Every Workflow Is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm), [The Brain Is a Key, Not a Box](https://blog.hackspree.com/#the-brain-is-a-key-not-a-box), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern).

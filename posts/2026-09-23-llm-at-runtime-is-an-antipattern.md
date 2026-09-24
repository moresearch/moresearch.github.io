---
title: "LLM at Runtime Is an Antipattern"
date: 2026-09-23
slug: llm-at-runtime-is-an-antipattern
summary: "An LLM in the event loop converts a fixed cost into a variable cost, and gives up replayability, auditability, versioning, and determinism in the process. The fix is a hard wall: the LLM belongs at compile time, as the semantic layer of the AutoML stack, proposing candidate features that an engine compiles and a small tree model freezes. Jev and Laya are a real advance over generative LLMs at runtime — and still the wrong shape for the event loop, because a decision model is a smaller variable cost, not a fixed one."
tags: llm, runtime, compile-time, automl, semantic-feature-engineering, decision-models, system-one, jev, laya, jevons, catboost, tabular-ml, software-economics, brooks, boehm, determinism, auditability, antipattern, essay
---

Fred Brooks observed that the hardest part of software engineering is not building the thing, but building it in a way that survives change:

> "The hardest single part of building a software system is deciding precisely what to build. No other part of the conceptual work is so difficult as establishing the detailed technical requirements... No other part of the work so cripples the resulting system if done wrong." — [No Silver Bullet — Essence and Accident in Software Engineering](https://www.cs.unc.edu/techreports/86-020.pdf) (1986)

Barry Boehm quantified what that survival costs. A defect caught at design time costs one unit; caught at test time, ten; caught in production, a hundred. Boehm and Basili's [Software Defect Reduction Top 10 List](https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf) opens with exactly that number — "Finding and fixing a software problem after delivery is often 100 times more expensive than finding and fixing it during the requirements and design phase" — and adds that 40 to 50 percent of project effort goes to avoidable rework.

Both men were writing about the same underlying phenomenon, and neither was writing about AI. They were writing about a discipline that had learned, painfully, that **where you pay a cost determines everything about the system that results**.

We are now building AI systems, and we are forgetting this lesson at scale. The current fashion — announced with a manifesto, priced per million tokens, and running inside the event loop of production systems — is precisely the move that both of them warned against.

---

## The Pattern We Keep Repeating

Every generation of software engineering has a moment where the industry realizes it has been doing something convenient instead of something correct.

We built stateful servers because state was easier to keep in memory. We built monoliths because one deployable was easier to reason about. We built services that called each other synchronously because the call stack was easier to follow.

None of these were mistakes. They were the right shape for the problem at the time. They became mistakes when the problem grew and the shape didn't.

Today, the shape is this:

```text
raw_event ──▶ LLM ──▶ action
             (per event, forever)
```

Every decision. Every event. Every tick. A network call to a nondeterministic, slow, expensive, unversionable model.

It works. It demos. It ships. And then it bleeds. The industry already documented this class of failure once — Google's [Machine Learning: The High-Interest Credit Card of Technical Debt](https://research.google/pubs/pub43146/) and the NeurIPS paper that followed, [Hidden Technical Debt in Machine Learning Systems](https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf), are catalogues of it, and their "CACE" principle — *Changing Anything Changes Everything* — is what an unversioned model in the loop does to every downstream decision.

## The Marginal Cost Problem

Here is the economic fact that decides everything else:

**An LLM in the event loop has a nonzero marginal cost per decision. Forever.**

Every call costs tokens. Every token costs money. Every dollar scales linearly with traffic. There is no volume discount that brings the marginal cost to zero. There is no learning curve that bends the cost downward. There is no amortization that spreads the cost across future decisions.

You have taken what should be a **fixed cost** — the design of a decision procedure — and turned it into a **variable cost** that compounds with every event your system processes.

This is the same mistake as building a service whose cost scales with every request instead of every deployment. It is the same mistake as paying per query instead of per schema. It is the same mistake as renting intelligence by the minute when you could own it by the artifact.

Software economics has a name for what you want here: **convert variable costs into fixed costs as early as possible.** You want the cost of a decision to be paid once, at design time, and then to be free at runtime. That is what compilation is — the [Futamura projections](https://link.springer.com/article/10.1023/A:1010095604496) are the purest statement of it. That is what caching is. That is what an [index](https://use-the-index-luke.com/sql/anatomy) is: a precomputed answer to a question you would otherwise re-derive on every scan. That is what every mature engineering discipline does.

Brooks called the residual difficulty that survives good design "essential complexity." Boehm & Basili priced the delay in discovering it. An LLM in the event loop does not reduce essential complexity. It simply pays for it again, on every event, in the most expensive currency available.

It is not flexibility. It is a subscription to your own architecture.

## What You Actually Pay

The cost of an LLM in the event loop is not just latency and dollars. It is the set of engineering properties you give up — and each of those properties has a price.

**You give up replayability.** You cannot reproduce yesterday's decision, because the model that produced it no longer exists in the form it existed yesterday — and even if it did, [inference is not bit-deterministic](https://thinkingmachines.ai/blog/defeating-nondeterminism-in-llm-inference/). Every incident review becomes guesswork. Every regression becomes a mystery. You pay for this in engineer-hours, forever. It is the same problem durable-execution systems solved for code, and the same lesson this blog keeps arriving at from the workflow side: [graphs describe intentions, logs describe executions](https://blog.hackspree.com/#every-workflow-is-an-fsm), and [the event log, not the plan, is the agent's honest memory](https://blog.hackspree.com/#stories-from-events).

**You give up auditability.** You cannot explain why the system did what it did, because the reasoning lives inside weights you do not own. Compliance costs go up — the EU AI Act's [record-keeping and transparency obligations](https://eur-lex.europa.eu/eli/reg/2024/1689/oj) assume you *can* record and explain the system's behaviour. Trust costs go down. Your ability to defend a decision to a customer, a regulator, or a court disappears.

**You give up versionability.** You cannot hash "GPT-4-turbo-as-of-Tuesday." You cannot pin it. You cannot roll it back. You cannot A/B test two versions of it without maintaining two vendor contracts. Model cards and dataset documentation ([Mitchell et al.](https://arxiv.org/abs/1810.03993), Pineau's [reproducibility checklist](https://www.cs.mcgill.ca/~jpineau/ReproducibilityChecklist.pdf)) were invented for exactly this gap, and they only work if the artifact is yours. You pay for the loss in the cost of change — and Boehm's curve is unforgiving.

**You give up determinism.** The same input produces different outputs. Your tests are probabilistic. Your SLAs are guesses. Your incident response is firefighting instead of forensics.

**You give up cost predictability.** Your bill scales linearly with traffic, forever, with no ceiling — and frontier input pricing is a band, not a number: $0.20 to $10 per million tokens, with output tokens about five times that, [per TypeSafe's own comparison table](https://typesafe.ai/blog/introducing-system-one-models-and-jev). Your CFO cannot model it. Your engineering team cannot defend it. Your pricing team cannot pass it through.

**You give up locality.** Your agent cannot run on a laptop, on an [edge device](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem), or in an air-gapped environment without a network dependency. Every decision is a round trip. Every round trip is a point of failure. [llama.cpp](https://github.com/ggml-org/llama.cpp) exists precisely because that dependency is not always acceptable.

These are not minor inconveniences. They are the properties that make software engineering a discipline rather than a craft. When you give them up, you are not building a system. You are building a demo that happens to be in production.

And you are paying for it. Every single event. Forever.

## Why It Feels Right Anyway

The antipattern survives because it maps to how we think about intelligence.

We imagine the agent as a brain. The brain sees the event, thinks, and acts. So we reach for the most brain-like thing we have and wire it into the event loop.

But production agents are not brains. They are *systems*. And systems have a distinction the brain metaphor hides:

**Some work is done once. Some work is done forever.**

Figuring out *what matters* about an event is design work. It happens once. It can be slow, expensive, and messy. It can involve a human staring at a whiteboard. It can involve an LLM proposing a hundred candidate features. It can involve weeks of measurement.

Deciding *what to do* about a specific event is operational work. It happens millions of times. It must be fast, cheap, and deterministic.

The antipattern collapses these two. It re-derives the design from scratch on every event.

This is the same confusion that shows up in every generation. We confuse **capital expenditure** with **operational expenditure**. We confuse **the cost of building a thing** with **the cost of running a thing**. We optimize for the demo and pay for it in production. The dark-factory arithmetic makes the same point from the other direction: an agent run is [not the same thing as automation](https://blog.hackspree.com/#task-automation-economics) — automation is the point at which the marginal cost of the next unit falls to near zero.

A compiler would never do this. You do not re-parse your source code every time you run your program. The entire discipline of software engineering is built on the premise that design-time and runtime are different places, and that artifacts move between them.

AI systems are now converging on the same premise. The question is whether we converge deliberately or after the bill arrives.

## The Sophisticated Version of the Same Mistake

The first correction people reach for is the **decision model**.

Models like **Jev** and **Laya** are a real improvement over a generative LLM in the loop. They strip out token generation entirely. They accept unstructured state and return typed answers — yes/no probabilities, a choice from a fixed set, a score along a rubric — with calibrated confidence. TypeSafe's pitch for Jev is a "frontier-intelligence function call: unstructured state in, typed probabilistic decisions out" ([Introducing System One Models & Jev](https://typesafe.ai/blog/introducing-system-one-models-and-jev), 15 September 2026). Convai's Laya is the open-weights answer in the same class: an Apache-2.0, non-autoregressive [System 1 decision engine](https://huggingface.co/convaiinnovations/laya) that answers typed questions in a single forward pass (~33 ms on GPU, 7.2 ms per question batched) and never emits text, "[so] there is nothing to parse and nothing to hallucinate."

These are real advances. They eliminate the token-generation tax. They are a step in the right direction.

**And they are still the wrong shape for runtime.**

From an economics perspective, a decision model is a smaller variable cost, not a fixed cost. It still scales with traffic. It still requires a GPU (or optimized runtime) to hit its stated latency. It still has a context window. It still has weights you do not own unless you self-host — and if you self-host, you have turned a per-query cost into a per-GPU-hour cost that is still variable.

Even a 322M–421M parameter decision model carries the weight of a neural network: hundreds of megabytes of weights, a forward pass through bidirectional encoder layers plus a decision head, a GPU to hit its stated latency. Laya's own model card is refreshingly blunt about the boundary: out of the box, the base checkpoints score around **0.35** on the typed-decisions benchmark against a **0.318** random baseline — near chance — and the 0.766 headline number comes from fine-tuning on that benchmark's training split. The card's advice: "Treat Laya as a fast foundation model to specialize, not as an omniscient zero-shot oracle." That sentence is the whole post in miniature. The value is created by *specialization*, which is a design-time act.

A decision model is a general-purpose classifier doing work that a feature vector and a gradient-boosted tree can do better, faster, and smaller on a narrow task. It is not the end of the line. It is a waypoint. It buys you time. It does not buy you economics.

## Reading the System One Post Closely

The best-documented decision model we have is Jev, and its launch post is worth reading as an economic document, not just a product announcement. TypeSafe is candid about the numbers, which makes the numbers usable.

The workload it publishes is not a chat benchmark. It is a set of four **decision DAGs** — the shape an actual production workflow takes once you stop prompting for prose and start writing code around a model:

[![TypeSafe AI's simplest published workflow — a decision compute graph: class/urgency/theme fanned out in parallel, then deterministic code for renewal date, overage, churn likelihood level, then four branches into save/loop, billing/inquiry, inquiry, and user conflict; Jev sits in the decisions and ordinary code sits everywhere else](https://framerusercontent.com/images/ih1bFwZGYJxlnijbTuXx3f9NeM.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

Run the same workflow against frontier LLMs and against Jev, and the Pareto picture is stark — Jev sits alone at roughly 300 ms and ~78% agreement with the smartest models, against 1.4 s to 18 s for everything else:

[![TypeSafe workflow evals — System One task accuracy vs. median latency per query: Jev at ~296 ms and 78%, Fable 5.1 at ~1.43 s and 73%, GPT-5.6 Terra at ~2.19 s and 74%, GPT-6 Astra at ~15.6 s and 72%, Gemini 3.5 Pro at ~18.1 s and 64%](https://framerusercontent.com/images/z4Uu1YpJeEZPBSMTCMI0CN2PX0.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

And the type-safety claims — the two columns that matter if a decision is buried in a dependency chain — are a different order of magnitude again:

[![TypeSafe type-safety eval — wrong tool calls and type errors over 1,000 samples: Jev 0% / 0% (schema matching guaranteed), GPT-5.6 Terra 1.0% wrong tool calls and 10% type errors, GPT-6 Astra 0.4% and 0%, with latency bands of ~289 ms vs ~1.6 s to ~16.1 s](https://framerusercontent.com/images/KEoJ6ZaJkOZG6mcjBsOlB3NCqek.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

What this post establishes is real, and I want to be precise about it:

- **Token generation was never free, and the tax was enormous.** Input tokens at $0.042 per million against a frontier band of $0.20–$10, with output tokens free instead of ~5x input, is a structural price difference, not a discount.
- **Type errors are an architectural property, not a model-quality property.** If the output space is defined at request time and the model never writes strings, schema violations stop being a probability and become an impossibility. That is a *compile-time* guarantee smuggled into a runtime call, and TypeSafe is right to call it table stakes for automation.
- **Calibration is what makes a decision usable.** An LLM that can do a task 95% of the time but cannot say which 5% is not automatable; the graph has no branch for it. TypeSafe's RLCD trains honesty directly — a strictly proper scoring rule means reporting truthful probabilities is the only way to maximise reward.
- **The published numbers are footnoted honestly.** The workflows were authored by TypeSafe's own capabilities team, the reference probabilities are the average of GPT-6 Astra and Fable 5.1 (which biases toward OpenAI and Anthropic), the baseline LLMs run inside TypeSafe's own structured-output wrapper, and the type-error figure is "not empirical — schema matching is guaranteed." The [third-party coverage](https://aimodelreport.com/articles/2026-09-21-jev-and-laya-introduce-a-new-model-class-non-autoregressive-decision-models-that/) flags the same caveats.

Now the part that interests me more than the speedups.

**The Doom demo.** TypeSafe's team wired Jev into Doom at ten queries per second and were pleased that it cost "~$7/hour" — which is roughly $61,000 a year to play one game, forever, at a rate that scales with every additional player, every additional query, and every additional frame. That number is the marginal-cost problem stated by the model's own authors, in a demo they posted as a celebration. It is a faster variable cost, not a fixed one.

**The Jevons naming.** TypeSafe named the model after William Stanley Jevons on purpose: "Every order of magnitude drop in the cost of intelligence unlocks orders of magnitude more use cases." I think that is exactly right, and I think it is the strongest argument *against* putting intelligence in the event loop. Jevons' paradox says efficiency gains increase total consumption. Cheaper per-call decisions do not shrink the variable-cost line on your P&L; they grow it, because now you call the model where you previously wrote an `if`. The bill you were hoping to avoid arrives with more line items.

**What the waypoint costs.** Jev is a hosted, closed API: weights, parameter counts, and self-hosting are not on offer, and its behaviour is versioned on someone else's release schedule. Laya is the open-weight counterpart, and its numbers show what "open" costs in practice: 421M (English) and 322M (multilingual) parameters, an ~808 MB English checkpoint in a 2.5 GB bundle, a GPU to hit 33 ms, a token budget per option that degrades past ~20 options — on Banking77's 77 labels Laya scores 0.425 where Jev scores 0.870 — calibration error that starts at 0.466 and only reaches 0.081 after fitting a temperature per question type on your domain, and an English checkpoint that shreds non-Latin scripts *while staying confident*: 0.000 accuracy on Khmer at 0.952 mean confidence. Confidence gating cannot rescue that, because the model does not know it cannot read the input. (All of it is in Laya's own [benchmark report](https://laya.convaiinnovations.com/), which is unusually honest about ceilings.)

None of this is an argument that Jev or Laya are bad work. They are the best available demonstration that the decision layer is a real, separate layer from the language layer — and that once you have separated it, what you have is a **classifier with a GPU, a context window, and a per-call price**. Which is the same thing a feature vector plus a gradient-boosted tree is, at a fraction of the weight, with a hash you can pin.

So take the waypoint seriously, use it where it fits, and treat it as the on-ramp rather than the destination. The decision model is how you *discover* the decision procedure. It is not how you *own* it.

## The Missing Layer in AutoML

Here is where the conversation usually turns toward AutoML, and where it usually stops too early.

Traditional AutoML — [auto-sklearn](https://arxiv.org/abs/2007.04074), [TPOT](https://epistasislab.github.io/tpot/), [H2O AutoML](https://docs.h2o.ai/h2o/latest-stable/h2o-docs/automl.html), and their descendants — automates the *syntactic* parts of the ML pipeline. It searches a fixed space of preprocessing steps, model families, and hyperparameters. It can impute, encode, scale, select, and ensemble. It is genuinely powerful.

What it cannot do is *understand* your event schemas.

It cannot look at a stream of `OrderPlaced`, `RefundRequested`, and `SupportTicketOpened` events and say: *"the ratio of refund requests to purchases over a rolling 90-day window is probably predictive of churn."* That is a semantic step. It requires knowing what the events mean, what the objective is, and how the domain works.

AutoML searches. It does not propose. The closest classical answer — [Deep Feature Synthesis](https://featuretools.alteryx.com/en/stable/) in Featuretools — mechanically composes primitives over relational tables, which is more ambition than most pipelines have and still a syntactic move.

That is precisely the gap where an LLM belongs. Not at runtime. At compile time. As the **semantic layer of the AutoML stack**. And it is no longer speculative: CAAFE ([Hollmann, Müller, Hutter](https://arxiv.org/abs/2305.03403), NeurIPS 2024) puts an LLM in the loop to read dataset context and *propose new features* in Python, then evaluates them with cross-validation, reporting gains over plain AutoML baselines. [AutoML-GPT](https://arxiv.org/abs/2305.02499) takes the same idea to pipeline construction. The pattern is being published; it is mostly being published on the wrong side of the wall.

The LLM reads the declared event schemas, the tool registry, and the objective, and proposes candidate features that a syntactic search would never generate. The engine compiles those features into point-in-time-correct folds. Traditional AutoML — or just CatBoost plus measurement — then trains, prunes, and freezes.

The LLM does not replace AutoML. It feeds it. It does the part AutoML was never able to do — and this blog has already argued the same shape from the enterprise rule side, where the LLM's job is [extracting decision models at design time](https://blog.hackspree.com/#on-rule-engines-automating-decision-models) rather than adjudicating at runtime.

## The Wall

The fix is a hard wall between compile time and runtime.

```text
COMPILE TIME — slow, messy, expensive, human-reviewed
  events   ─┐
  actions  ─┼─▶ LLM ─▶ candidate   ─▶ engine ─▶ AutoML ─▶ frozen
  objective ┘  semantic  feature spec   compiles   trains    artifacts
               layer     (40-100)      PIT folds  prunes    + hashes
═══════════════════════ HARD WALL ═══════════════════════
RUNTIME — fast, deterministic, replayable, versioned
  artifacts ─▶ verify hashes ─▶ Vector() ─▶ tree model ─▶ action ─▶ tool
```

Everything to the left is slow, expensive, and allowed to be messy. Everything to the right is pure code and frozen data.

The only things that cross are declarative artifacts and their hashes. No prompts. No API calls. No model weights except the `.cbm`.

**This is a capital expenditure.** You pay once, at design time, for an artifact. You amortize it across every decision the system ever makes. The marginal cost of the millionth decision is zero.

This is the same move that turned infrastructure into something deployable — Bazel and Nix building content-addressed, reproducible artifacts, [SLSA](https://slsa.dev/) and [in-toto](https://in-toto.io/) binding provenance to digests, ML compilers from [MLIR](https://mlir.llvm.org/) to [TVM](https://tvm.apache.org/) and [ONNX Runtime](https://onnxruntime.ai/) lowering a graph into something a runtime can execute without a scientist in the loop. You define the artifact, version it, hash it, verify it at runtime. The wall is what makes the system engineering-grade. It is also what makes it economically sustainable.

## What Replaces the Decision Model

A decision model — Jev, Laya, or any 300M–500M parameter classifier — is replaced by a **feature vector and a small traditional model**: [CatBoost](https://catboost.ai/), [XGBoost](https://xgboost.readthedocs.io/en/stable/tutorials/saving_model.html), [LightGBM](https://lightgbm.readthedocs.io/en/latest/), a [scikit-learn histogram gradient booster](https://scikit-learn.org/stable/modules/ensemble.html), a logistic regression, or whatever the AutoML search selects as best for the task. Saving and freezing the model is a one-liner, and the frozen file is a few megabytes, not a checkpoint directory.

The traditional model is not a black box. You can inspect feature importance, compute [SHAP values](https://arxiv.org/abs/1705.07874), and explain every decision to a regulator, a customer, or yourself at 3 AM — with [Molnar's Interpretable Machine Learning](https://christophm.github.io/interpretable-ml-book/) as the reference for what those explanations do and do not mean.

But how do you get a small tree ensemble to match a 400M-parameter neural network on your specific task?

**You let the LLM do the semantic feature engineering, and you let AutoML do the rest.**

A decision model is a general-purpose classifier. A CatBoost model trained on features engineered for your specific event stream is a special-purpose classifier. Special-purpose classifiers routinely beat general-purpose ones on narrow domains — that is not a novel claim, it is the entire history of applied machine learning. On tabular data specifically, the evidence is not close: Shwartz-Ziv and Armon's [Tabular Data: Deep Learning Is Not All You Need](https://arxiv.org/abs/2106.03253) shows XGBoost beating the deep tabular models on 11 datasets, and Grinsztajn, Oyallon, and Varoquaux's [Why do tree-based models still outperform deep learning on typical tabular data?](https://arxiv.org/abs/2207.08815) attributes the gap to trees' robustness to uninformative features and their bias toward axis-aligned splits — which is exactly what a hand-engineered feature vector looks like. CatBoost's own contribution ([Prokhorenkova et al.](https://arxiv.org/abs/1706.09516)) was ordered boosting and native categorical handling, so the boring default also handles your `category` column without a bespoke encoder.

This blog has been circling the same conclusion from the small-model side: [specialization, not size, is the lever](https://blog.hackspree.com/#specialized-small-language-models), and the deployable unit is [the model plus the harness](https://blog.hackspree.com/#better-harnesses-smaller-models) — the harness being where the shared difficulty gets lifted out of the model and amortized.

From an economics perspective it is **economies of scope, not economies of scale**. You are not winning by making the model bigger. You are winning by making the problem narrower. The narrower the problem, the smaller the model, the cheaper the inference, the more explainable the decision.

This is the same reason a database index beats a full table scan. This is the same reason a compiled function beats an interpreter. This is the same reason a cache beats a recomputation. Specialization is the oldest economic move in computing.

## The Pipeline, Generically

Forget the specific domain. Here is the pattern, stripped of any particular application.

**Inputs declared by a human:**

```yaml
events:
  - name: EntityCreated
    fields: [id, category, timestamp, attributes]
  - name: EntityUpdated
    fields: [id, changed_fields, timestamp]
  - name: EntityObserved
    fields: [id, signal, timestamp]

actions:
  - route
  - escalate
  - defer
  - ignore

objective: maximize_decision_accuracy
```

**Compile-time LLM proposes candidate features:**

```yaml
candidate_features:
  - name: entity_age_days
    kind: numeric
    from: EntityCreated.timestamp
  - name: update_velocity_30d
    kind: numeric
    from: EntityUpdated count over 30d window
  - name: category_onehot
    kind: categorical
    from: EntityCreated.category
  - name: signal_ratio_90d
    kind: numeric
    from: EntityObserved.signal / EntityCreated count over 90d
  - name: time_since_last_escalation
    kind: numeric
    from: EntityObserved.timestamp
  # ... 40-100 candidates total
```

**The engine compiles the spec:**

- Folds the event stream into a projection
- Generates a `Vector()` function that extracts the features
- Enforces point-in-time correctness structurally (no leakage from future events) — the same guarantee feature platforms sell as [point-in-time joins](https://docs.feast.dev/getting-started/concepts/point-in-time-joins), and the reason [Chronon](https://github.com/airbnb/chronon) exists
- Emits a `.cd` file and a content hash

**AutoML trains and prunes:**

- CatBoost (or whatever the search selects) trains on the full candidate set
- Permutation importance ranks features
- Temporal cross-validation measures generalization
- The bottom 60–80% of features are dropped
- The survivors are frozen

**The output:**

- A frozen spec with N features
- A trained model
- Hashes binding everything together
- A daemon that loads the artifacts and never looks back

This is not a support router. It is not a fraud detector. It is not a recommendation engine. It is the *general shape* of any event-driven decision system where the events are structured, the objective is measurable, and the labels are obtainable.

The data-engineering version of this is already mainstream: validate the inputs ([JSON Schema](https://json-schema.org/), [Great Expectations](https://greatexpectations.io/), [Pandera](https://pandera.readthedocs.io/)), test them in CI ([dbt tests](https://docs.getdbt.com/docs/build/data-tests)), version the artifacts ([DVC](https://dvc.org/)). The feature contract is a data contract with a model attached.

The LLM does the part AutoML cannot: it looks at your event schemas and proposes features that are *semantically meaningful for your domain*. AutoML does the part the LLM cannot: it searches the space rigorously, measures generalization, and prunes without sentiment.

The two are complementary. Neither is a runtime dependency.

## What Crosses the Wall

Only these, and only with hashes:

```text
spec.yaml      → H1   → binds features to model
model.cbm      → H2   → binds weights to features
schemas.yaml   → H3   → binds events to features
actions.yaml   → H4   → binds decisions to tools
triggers.yaml  → H5   → binds conditions to actions
```

At runtime, before acting, the daemon verifies:

```text
H1(spec)     == model.spec_hash
H3(schemas)  == spec.schema_hash
H4(actions)  == spec.action_hash
```

Mismatch is a hard refusal. This is the mechanism that guarantees no LLM-shaped or neural-network-shaped artifact leaked into runtime. If it is not in the frozen set, it cannot affect the decision. It is the same discipline reproducible builds apply to binaries, applied to a decision procedure.

Every way the antipattern sneaks back in — an LLM call inside the fold, a decision model inside `Vector()`, a neural fallback for low-confidence cases — breaks the same properties at once: replayability, auditability, versioning, determinism. There is no partial wall.

It also breaks the economics. Every sneaked-in call is a variable cost that survives the compile-time conversion. Every exception is a subscription you forgot you signed.

## The Broader Trend

Zoom out. The pattern is recognizable.

**AI systems are getting the treatment distributed systems got a decade ago.** Versioning, hashing, replayability, content-addressed artifacts, compile-time validation, hard boundaries between design and execution. These are not AI-specific ideas. They are software engineering ideas that AI systems are adopting because they have to — the MLOps literature has been [arguing for the discipline](https://arxiv.org/abs/2205.02302) since before the current wave, and Google's [Rules of ML](https://developers.google.com/machine-learning/guides/rules-of-ml) and [ML Test Score](https://research.google/pubs/the-ml-test-score-a-rubric-for-ml-production-readiness-and-technical-debt-reduction/) were early attempts to write it down.

**AutoML is getting the semantic layer it always lacked.** Traditional AutoML is syntactic. It searches transformations of features you give it. It cannot invent features from domain knowledge. LLMs can — CAAFE is the proof of concept — and that is the missing piece, and it belongs at compile time.

**The pendulum is swinging from monolithic to modular.** One giant model doing everything at runtime is being replaced by specialized components doing narrow things well: a semantic feature proposer here, a syntactic feature pruner there, a tree ensemble for the decision.

**The shift-left is happening.** Design-time tools — LLMs, compilers, measurement loops — are absorbing work that used to happen at runtime. Verification is the bottleneck, and [the retreat to harness engineering](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering) is the industry noticing that the expensive part was always the check, not the generation.

**Determinism is becoming a first-class requirement.** Compliance, debugging, trust, and cost are forcing AI systems to be reproducible. Replayable pipelines with hashed artifacts are becoming the norm.

**Small models are winning at the edge of the stack.** Not because they are smarter, but because the task does not require smart. It requires specialized. And specialized can be small — [a 3B agentic model](https://blog.hackspree.com/#nanbeige4-2-3b-agentic-model) trained on the right distribution beats a generalist many times its size in-domain.

The wall is not a niche architectural preference. It is where the industry is heading. The only question is whether we get there deliberately or after two years of untangling.

## The Economics, Stated Plainly

If you take nothing else from this post, take this:

**Software economics has one rule that matters more than any other: push costs from runtime to design time whenever you can.**

Compilers do this. Caches do this. Indexes do this. Schemas do this. Type systems do this. Every mature engineering discipline has figured out that the cheapest decision is the one you don't have to make again. Boehm's [Software Engineering Economics](https://www.informit.com/store/software-engineering-economics-9780138221225) is the book that made this arithmetic respectable, and the [NIST study on inadequate software testing infrastructure](https://www.nist.gov/system/files/documents/director/planning/report02-3.pdf) is the bill that arrives when a discipline ignores it.

An LLM in the event loop does the opposite. It takes a decision that could be made once and forces you to make it forever. It converts a fixed cost into a variable cost. It converts an asset into a subscription. It converts a design into a dependency.

This is the antipattern, stated economically: **you are paying rent on intelligence you could have bought.**

The wall is how you buy it. The compile-time LLM — as the semantic layer of AutoML — is how you design it. Traditional ML is how you run it. The marginal cost of the millionth decision is zero.

That is not a technical claim. It is an economic one. And it is the reason this pattern wins.

## When the Wall Is Wrong

I want to be honest about the boundaries of this argument.

The wall is not universal. There are cases where runtime LLM calls or runtime decision models are the correct choice:

- **Decisions depend on rare, nuanced, or unstructured signals** that cannot be summarized into features.
- **Event schemas change frequently**, making retraining expensive.
- **You need zero-shot generalization** to situations the model has never seen — which is exactly what Laya's base checkpoints are honest about not providing.
- **You cannot get reliable labels**, so you cannot train a supervised model.
- **You need a conversational interface** — the user expects prose, not a probability. This is the one place where a language model is not an antipattern but the product.

There is also a cold-start problem. Without labeled examples, traditional ML has nothing to learn from. You can bootstrap with LLM-generated labels — and this is a legitimate use of a runtime model as a *labeler*, writing into the training set rather than into the decision path — but those labels are only as good as the LLM's judgment, and the gap between expressed and actual confidence in language models is [documented](https://arxiv.org/abs/2205.14334) and [still an open problem](https://arxiv.org/abs/2409.00352). Validate against real, human-labeled data before trusting the model in production.

In those cases, the flexibility of an LLM or a decision model is worth the cost. But be honest about the trade: you are choosing flexibility over speed, cost, determinism, and auditability. Those are real trade-offs, not free lunches — the same point this blog makes about every other "obvious" engineering answer: [there are no solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs).

Most production agents do not fall into these categories. Most decisions are structured, repetitive, and labelable. Most of the time, the wall wins.

## What This Changed in My View

Two things moved while I was writing this, and both moved toward the wall rather than away from it.

The first is that the decision-model waypoint arrived faster than I expected, and it is better than I expected. I had been treating "put a classifier in the loop instead of an LLM" as an obvious-but-unimplemented idea. TypeSafe shipped it, published the workflow evals, footnoted their own biases, and priced the calls. Laya shipped open weights, a typed-decision benchmark, and a limitations section that says out loud that the base model is near random on the task it is being advertised for. That honesty is what makes the case *for the wall* — because the same post that shows Jev on a Pareto frontier also shows the "$7/hour to play Doom" line, and that line is a variable cost doing what variable costs do.

The second is that I had underweighted how much of the argument is about **artifacts rather than models**. I came in thinking the choice was "big model at runtime" versus "small model at compile time." The stronger framing is that the choice is "a call you make forever" versus "an artifact you own." A compiled spec, a frozen tree ensemble, and five hashes are not a smaller AI system. They are a different kind of thing: something you can diff in a pull request, roll back, replay, audit, and hand to someone else. That is what [conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) looks like for a decision procedure, and it is why the cost argument and the engineering argument keep landing in the same place.

What did not move: the honest cases at the boundary. Conversational interfaces are not going away, and unlabelable decisions still need something flexible. The wall is a default, not a religion.

## The Point

Compile time is where the LLM lives. It acts as the semantic layer of the AutoML stack: it reads your event schemas, your tool registry, and your objective, and proposes candidate features that a syntactic search would never generate. An engine compiles those features into point-in-time-correct folds. Traditional ML — CatBoost, XGBoost, LightGBM, or whatever the search selects — trains and prunes. The survivors are frozen.

Runtime is where neither the LLM nor the decision model appears. The daemon loads the frozen artifacts, evaluates the compiled spec, builds a feature vector, asks the traditional model for an action, maps that action to a tool, executes, and publishes. You can hand that artifact to an [ONNX Runtime](https://onnxruntime.ai/) session, an [ONNX Runtime Go binding](https://github.com/yalue/onnxruntime_go), or a CatBoost C library inside a static binary and run it on a laptop with the network unplugged.

Hashes bind model to spec to schemas to actions. Mismatch is a hard refusal.

Decision models like Jev and Laya are a meaningful improvement over runtime LLMs. They eliminate token generation, cut latency to hundreds of milliseconds, guarantee type safety, and bring calibration to the decision. But they are still neural networks with per-call pricing. They still carry hundreds of megabytes of weights. They still need a GPU to hit their stated latency. They still have a context window and a token budget per option that degrades as your label space grows. They are a waypoint, not a destination.

The destination is a feature contract and a small traditional model. The compile-time LLM is what gets you there. The wall is what keeps you there.

---

## A Call to the Community

We are at the beginning of a discipline. AI systems engineering does not have its *Design Patterns* yet. It does not have its *Site Reliability Engineering* yet. It does not have its *Choose Boring Technology* yet.

What it has is a lot of demos in production and a growing sense that something is wrong. The sense is right. The wrongness is economic as much as it is architectural.

Here is what I am asking:

**Stop putting LLMs in the event loop.** Not because they are bad. Because the event loop is the wrong place for them. The right place is design time, as the semantic layer of your AutoML pipeline.

**Start treating prompts as source code.** Version them. Review them. Hash them. Compile them into artifacts. Do not let them run free at runtime. Do not let them be a subscription you cannot cancel.

**Measure your traditional models.** Do not assume a 400M-parameter model beats a 2 MB tree on your task. Test it. You will be surprised more often than you expect.

**Build the wall.** Draw the line between compile time and runtime in your architecture. Enforce it with hashes. Refuse to run when the wall is broken. Refuse to run when a variable cost sneaks past.

**Write down your unit economics.** Know your cost per decision. Know how it scales. Know what it would look like if it were zero. Then ask yourself why it isn't. An agent run is not automation until the marginal cost of the next one is yours to decide.

**Write down what you learn.** This discipline is being invented right now, in production systems, by people who are too busy shipping to write blog posts. If you have built one of these systems, write it up. The community needs your scars.

The next decade of AI systems will not be won by the teams with the biggest models. It will be won by the teams who understand that intelligence belongs at design time and execution belongs at runtime — and who build the wall between them.

Brooks told us there is no silver bullet. Boehm told us what the delay costs. We have known this for forty years.

Let's apply it.

---

## References

**The economics of where you pay**

- Brooks, F. P. [*No Silver Bullet — Essence and Accident in Software Engineering*](https://www.cs.unc.edu/techreports/86-020.pdf) — UNC tech report TR86-020, 1986 (later IEEE Computer 20(4), 1987). The source of "essence vs. accident," "there is no single development... which by itself promises even one order-of-magnitude improvement," and the "hardest single part... is deciding precisely what to build" quote used above. See also [*The Mythical Man-Month*](https://archive.org/details/mythicalmanmonth00broo) (1975) for conceptual integrity, and this blog's [Brooks series](https://blog.hackspree.com/#brooks-design-conceptual-integrity) and [Accidental Complexity Is the Only Complexity You Can Remove](https://blog.hackspree.com/#brooks-accidental-complexity) for the applied version.
- Boehm, B. and Basili, V. [*Software Defect Reduction Top 10 List*](https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf) — IEEE Computer 34(1), January 2001, pp. 135–137. Item 1 is the 100x-after-delivery figure; the same page gives the 40–50% avoidable-rework estimate. The cost gradient the essay opens with (1 : 10 : 100 across design, test, production) is Boehm's cost-of-change curve, from [*Software Engineering Economics*](https://www.informit.com/store/software-engineering-economics-9780138221225) (Prentice Hall, 1981).
- NIST. [*The Economic Impacts of Inadequate Infrastructure for Software Testing*](https://www.nist.gov/system/files/documents/director/planning/report02-3.pdf) — Planning Report 02-3, May 2002. The $59.5B national estimate that turns Boehm's curve into a budget line.
- Futamura, Y. [*Partial Evaluation of Computation Process — An Approach to a Compiler-Compiler*](https://link.springer.com/article/10.1023/A:1010095604496) — Higher-Order and Symbolic Computation 12, 1999 (originally 1971). The formal statement that specializing a program to its known inputs is where runtime cost goes to die.
- Winand, M. [*Anatomy of an Index*](https://use-the-index-luke.com/sql/anatomy) — use-the-index-luke.com. The everyday proof of the same principle: an index is a design-time artifact that removes a runtime derivation.
- This blog: [Engineering is art and philosophy, grounded in economic law](https://blog.hackspree.com/#engineering-is-economics), [The economics of the dark factory](https://blog.hackspree.com/#dark-factory-economics), [Task automation economics: why an agent run is not automation](https://blog.hackspree.com/#task-automation-economics), [No solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs).

**The decision model at runtime (and why it is a waypoint)**

- Almeida, D. [*Introducing System One Models & Jev*](https://typesafe.ai/blog/introducing-system-one-models-and-jev) — TypeSafe AI, 15 September 2026. The primary source for the figures reproduced above: the published decision DAG, the workflow-eval Pareto plot (System One task accuracy vs. median latency), and the type-safety evaluation (wrong tool calls and type errors over 1,000 samples); plus pricing ($0.042/MTok input, output free, against a $0.20–$10 frontier band with output ~5x input), latency (70–500 ms vs. 3–329 s), the 193.6x faster / 444.6x cheaper home-page claim, the workflow-eval methodology (identical workflow for every model, reference probabilities from GPT-6 Astra and Fable 5.1), the Doom demo at ten queries per second for ~$7/hour, the Wikiracing demo (cardinality up to 255), the Kahneman *Thinking, Fast and Slow* naming, and the Jevons naming. TypeSafe's own "Nuance" notes flag the capabilities-team authorship, the OpenAI/Anthropic-biased reference, and that the type-error figure is guaranteed by construction rather than empirical.
- [*Jev AI Model (TypeSafe) — Typed System One Decisions*](https://jevmodel.org/) — third-party reference guide, updated 21 September 2026: API access, the $42-per-billion input price, `jev-1.13.0` aliases, the closed-weights status, and the explicit contrast with Laya as an independent open-weight System One model.
- Convai Innovations. [*Laya*](https://huggingface.co/convaiinnovations/laya) (Hugging Face model hub, three checkpoints under one repo) — 421M ModernBERT-large English, 322M mmBERT-base multilingual, and a typed-decisions checkpoint, Apache 2.0; the three decision primitives (`choice`, `score`, `noul`); ~33 ms p50 and 7.2 ms/question batched; the ~808 MB English checkpoint inside a 2.5 GB bundle; and the model card's own limits: base checkpoints at ~0.35 on typed-decisions against a 0.318 random baseline, 0.766 only after fine-tuning, ECE 0.466 → 0.081 after per-question-type temperature fitting, and high-cardinality degradation past ~20 options.
- Convai Innovations. [*Laya — 33ms Multilingual System 1 Decision Engine*](https://laya.convaiinnovations.com/) — the research write-up behind the card: RLCD with strictly proper scoring rules, the 51-language MASSIVE sweep (0.000 accuracy on Khmer at 0.952 mean confidence; confidence does not warn you when the tokenizer cannot read the script), the Banking77 comparison (Laya 0.425 vs. Jev 0.870 at 77 labels), and the founder's account of the March 2025 [SalesRLAgent](https://arxiv.org/abs/2503.23303) and September 2025 [schema-based decision framework](https://arxiv.org/abs/2510.01237) that preceded it.
- [*NandhaKishorM/laya*](https://github.com/NandhaKishorM/laya) — the SDK, router, and reproducible benchmark harnesses (Apache 2.0); [`laya-typed-decisions`](https://huggingface.co/convaiinnovations/laya-typed-decisions) is the checkpoint those benchmarks are run against.
- Iverson, L. [*Jev and Laya define a new model class: typed-decision heads that replace LLM calls for routing*](https://aimodelreport.com/articles/2026-09-21-jev-and-laya-introduce-a-new-model-class-non-autoregressive-decision-models-that/) — AI Model Report, 21 September 2026. Independent framing of the two releases, the 32.8–276 ms and $0.042/MTok comparisons, the zero-shot 0.362-vs-0.318-baseline caveat, and developer reports (5–18x faster than a frontier model on a command-safety classifier; 10–20x cheaper on email classification at slightly lower accuracy).

**LLMs at compile time: AutoML and the semantic feature layer**

- Hollmann, N., Müller, S., and Hutter, F. [*Large Language Models for Automated Data Science: Introducing CAAFE for Context-Aware Automated Feature Engineering*](https://arxiv.org/abs/2305.03403) — NeurIPS 2024. The closest published instance of the pattern: an LLM reads dataset context and proposes new features, which are then evaluated by cross-validation rather than trusted.
- Zhang, S. et al. [*AutoML-GPT: Automatic Machine Learning with GPT*](https://arxiv.org/abs/2305.02499) — LLM-driven pipeline construction across model, optimizer, and hyperparameter choices.
- Feurer, M. et al. [*Auto-sklearn 2.0*](https://arxiv.org/abs/2007.04074); [TPOT](https://epistasislab.github.io/tpot/); [H2O AutoML](https://docs.h2o.ai/h2o/latest-stable/h2o-docs/automl.html) — the syntactic search layer that the semantic layer feeds.
- Kanter, J. M. and Veeramachaneni, K. [*Deep Feature Synthesis*](https://featuretools.alteryx.com/en/stable/) — Featuretools; the classical attempt to automate feature construction by composing primitives over relational structure, and the clearest illustration of where syntactic search stops.

**Why a small tree beats a big network on a structured decision**

- Shwartz-Ziv, R. and Armon, A. [*Tabular Data: Deep Learning Is Not All You Need*](https://arxiv.org/abs/2106.03253) — XGBoost outperforms the deep tabular models across 11 datasets, and deep ensembles are needed to be competitive.
- Grinsztajn, L., Oyallon, E., and Varoquaux, G. [*Why do tree-based models still outperform deep learning on typical tabular data?*](https://arxiv.org/abs/2207.08815) — NeurIPS 2022 datasets-and-benchmarks track; the mechanism (robustness to uninformative features, axis-aligned splits) is what engineered features exploit.
- Prokhorenkova, L. et al. [*CatBoost: unbiased boosting with categorical features*](https://arxiv.org/abs/1706.09516) — [catboost.ai](https://catboost.ai/) / [model prediction API](https://catboost.ai/docs/en/concepts/python-reference_catboost_predict); [XGBoost model serialization](https://xgboost.readthedocs.io/en/stable/tutorials/saving_model.html); [LightGBM](https://lightgbm.readthedocs.io/en/latest/); [scikit-learn histogram gradient boosting](https://scikit-learn.org/stable/modules/ensemble.html). The point is not which library wins; it is that "freeze a few megabytes of trees and score in milliseconds on one core" is a supported, boring operation in all of them.
- Lundberg, S. and Lee, S.-I. [*A Unified Approach to Interpreting Model Predictions*](https://arxiv.org/abs/1705.07874) — SHAP; and Molnar, C. [*Interpretable Machine Learning*](https://christophm.github.io/interpretable-ml-book/) for what an explanation does and does not license you to claim.

**Artifacts, hashes, and the wall**

- [MLIR](https://mlir.llvm.org/), [Apache TVM](https://tvm.apache.org/), [ONNX](https://onnx.ai/) / [ONNX Runtime](https://onnxruntime.ai/), and the [ONNX Runtime Go bindings](https://github.com/yalue/onnxruntime_go) — compile a model graph once, execute it without a scientist in the loop, embed it in a static binary.
- [SLSA](https://slsa.dev/), [in-toto](https://in-toto.io/), [Reproducible Builds](https://reproducible-builds.org/), [Nix](https://nixos.org/), [Bazel](https://bazel.build/) — provenance, content addressing, and "refuse to run when the digest does not match," which is the mechanism the wall needs.
- [JSON Schema](https://json-schema.org/), [Great Expectations](https://greatexpectations.io/), [Pandera](https://pandera.readthedocs.io/), [dbt tests](https://docs.getdbt.com/docs/build/data-tests), [DVC](https://dvc.org/) — the feature contract as a data contract: declared, validated at build time, versioned at rest.
- Feast. [*Point-in-time joins*](https://docs.feast.dev/getting-started/concepts/point-in-time-joins); Airbnb's [Chronon](https://github.com/airbnb/chronon) — the mechanics of compiling features so that no fold can see the future.

**Determinism, calibration, and auditability**

- [*Defeating Nondeterminism in LLM Inference*](https://thinkingmachines.ai/blog/defeating-nondeterminism-in-llm-inference/) — Thinking Machines, 2025. Why the same prompt and the same weights still do not reproduce a decision, and what it would take to make them.
- Kadavath, S. et al. [*Language Models (Mostly) Know What They Know*](https://arxiv.org/abs/2205.14334); [*Does Alignment Tuning Really Break LLMs' Internal Confidence?*](https://arxiv.org/abs/2409.00352) — expressed versus actual confidence, which is the property a decision graph has to branch on.
- Ji, Z. et al. [*Survey of Hallucination in Natural Language Generation*](https://arxiv.org/abs/2202.03629) — the taxonomy of a failure mode that type-safe decision models structurally cannot have and frozen tree ensembles never had.
- Mitchell, M. et al. [*Model Cards for Model Reporting*](https://arxiv.org/abs/1810.03993); Pineau, J. [*Reproducibility Checklist*](https://www.cs.mcgill.ca/~jpineau/ReproducibilityChecklist.pdf) — documentation and reproducibility as preconditions for audit.
- Sculley, D. et al. [*Machine Learning: The High-Interest Credit Card of Technical Debt*](https://research.google/pubs/pub43146/) (2014) and [*Hidden Technical Debt in Machine Learning Systems*](https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf) (NeurIPS 2015); Breck, E. et al. [*The ML Test Score*](https://research.google/pubs/the-ml-test-score-a-rubric-for-ml-production-readiness-and-technical-debt-reduction/); Google's [Rules of ML](https://developers.google.com/machine-learning/guides/rules-of-ml); Kreuzberger, D. et al. [*MLOps: Overview, Definition, and Architecture*](https://arxiv.org/abs/2205.02302) — what the ML engineering literature already knows about whose cost lands where.
- [EU AI Act, Regulation (EU) 2024/1689](https://eur-lex.europa.eu/eli/reg/2024/1689/oj) — see the [AI Act Explorer](https://artificialintelligenceact.eu/ai-act-explorer/) for the logging, documentation, and human-oversight articles that make "we called a hosted model" an unsatisfying answer.

**Small, local, and specialized**

- Gerganov, G. et al. [llama.cpp](https://github.com/ggml-org/llama.cpp); Xu, M. et al. [*On-Device Language Models: A Comprehensive Review*](https://arxiv.org/abs/2409.00088); Sanh, V. et al. [*DistilBERT*](https://arxiv.org/abs/1910.01108) — running and shrinking models that never needed the cloud.
- This blog: [SSLM: Specialized Small Language Models](https://blog.hackspree.com/#specialized-small-language-models), [Nanbeige4.2-3B](https://blog.hackspree.com/#nanbeige4-2-3b-agentic-model), [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models), [On-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem), [Go Can Keep Structured LLM Runtimes Boring](https://blog.hackspree.com/#go-can-keep-structured-llm-runtimes-boring).

**Compile time in the enterprise: rules and decision models**

- This blog: [On Rule Engines — Automating Decision Models](https://blog.hackspree.com/#on-rule-engines-automating-decision-models) and [Five Patterns for Composite AI](https://blog.hackspree.com/#on-rule-engines-five-patterns) — the KU Leuven program on extracting DMN decision models from text with LLMs, and the composite-AI patterns that keep extraction at design time. See also Goossens, A., De Smedt, J., and Vanthienen, J. [*Extracting Decision Model and Notation Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667) — Expert Systems with Applications 211, 2023.

**Durable execution, replay, and events as truth**

- [Every workflow is an FSM. Not every FSM is a workflow.](https://blog.hackspree.com/#every-workflow-is-an-fsm), [Durable Daemons](https://blog.hackspree.com/#durable-daemons), [Stories from Events](https://blog.hackspree.com/#stories-from-events), [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering), [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents), [Harnessing Agentic AI Systems: A Pattern Language](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems), [SWE-Agent Economics](https://blog.hackspree.com/#why-i-focused-my-research-on-swe-agent-economics) — the same argument from the workflow, event-sourcing, and harness sides.

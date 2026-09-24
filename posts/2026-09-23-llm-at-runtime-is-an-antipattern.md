---
title: "LLM at Runtime Is an Antipattern"
date: 2026-09-23
slug: llm-at-runtime-is-an-antipattern
summary: "An LLM in the event loop converts a fixed cost into a variable cost, and gives up replayability, auditability, versioning, and determinism in the process. The fix is a hard wall: the LLM belongs at compile time, as the semantic layer of the AutoML stack, proposing candidate features that an engine compiles and a small tree model freezes. Jev and Laya are a real advance over generative LLMs at runtime — and still the wrong shape for the event loop, because a decision model is a smaller variable cost, not a fixed one."
tags: llm, runtime, compile-time, automl, semantic-feature-engineering, decision-models, system-one, jev, laya, jevons, catboost, tabular-ml, software-economics, brooks, boehm, determinism, auditability, antipattern, essay
---

Fred Brooks observed that the hardest part of software engineering is not building the thing, but building it in a way that survives change:

> "The hardest single part of building a software system is deciding precisely what to build... No other part of the work so cripples the resulting system if done wrong." — [No Silver Bullet — Essence and Accident in Software Engineering](https://www.cs.unc.edu/techreports/86-020.pdf) (1986)

Barry Boehm priced the delay. A defect caught at design time costs one unit; caught at test time, ten; caught in production, a hundred. Boehm and Basili's [Software Defect Reduction Top 10 List](https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf) opens with exactly that number — "Finding and fixing a software problem after delivery is often 100 times more expensive than finding and fixing it during the requirements and design phase" — and adds that 40 to 50 percent of project effort goes to avoidable rework.

Neither man was writing about AI. They were writing about a discipline that had learned, painfully, that **where you pay a cost determines everything about the system that results**. We are now building AI systems and forgetting that lesson at scale: an LLM in the event loop of production systems, priced per million tokens, running forever.

---

## The Shape We Keep Choosing

Today the shape is this:

```text
raw_event ──▶ LLM ──▶ action
             (per event, forever)
```

Every decision. Every event. Every tick. A network call to a nondeterministic, slow, expensive, unversionable model. It works. It demos. It ships. And then it bleeds.

One economic fact decides everything else about that shape: **an LLM in the event loop has a nonzero marginal cost per decision, forever.** Every call costs tokens. Every token costs money. Every dollar scales with traffic. There is no volume discount that takes the marginal cost to zero, no learning curve that bends it downward, no amortization that spreads it across future decisions. You have taken what should be a **fixed cost** — the design of a decision procedure — and turned it into a **variable cost** that compounds with every event the system processes.

Software economics has one rule for this: **convert variable costs into fixed costs as early as possible.** That is what compilation is (the [Futamura projections](https://link.springer.com/article/10.1023/A:1010095604496) are its purest statement). That is what caching is. That is what an [index](https://use-the-index-luke.com/sql/anatomy) is: a precomputed answer to a question you would otherwise re-derive on every scan. Brooks called the difficulty that survives good design "essential complexity"; an LLM in the event loop does not reduce it, it re-pays for it on every event in the most expensive currency available.

It is not flexibility. It is a subscription to your own architecture. The ML engineering literature named this failure mode years ago — [Machine Learning: The High-Interest Credit Card of Technical Debt](https://research.google/pubs/pub43146/) and [Hidden Technical Debt in Machine Learning Systems](https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf), whose CACE principle ("Changing Anything Changes Everything") is precisely what an unversioned model in the loop does to every downstream decision.

## What You Give Up

The cost is not only latency and dollars. It is the set of engineering properties you surrender, each with its own price.

- **Replayability.** You cannot reproduce yesterday's decision: the model no longer exists in the form it existed yesterday, and even with identical weights [inference is not bit-deterministic](https://thinkingmachines.ai/blog/defeating-nondeterminism-in-llm-inference/). Incident review becomes guesswork and regression becomes mystery — the same problem durable execution solved for code, and the same lesson this blog keeps reaching from the workflow side: [graphs describe intentions, logs describe executions](https://blog.hackspree.com/#every-workflow-is-an-fsm).
- **Auditability.** You cannot explain why the system acted, because the reasoning lives inside weights you do not own. The EU AI Act's [documentation and logging obligations](https://eur-lex.europa.eu/eli/reg/2024/1689/oj) assume you *can* record and explain behaviour.
- **Versionability.** You cannot hash "GPT-4-turbo-as-of-Tuesday," pin it, roll it back, or A/B two versions without two vendor contracts. [Model cards](https://arxiv.org/abs/1810.03993) and [reproducibility checklists](https://www.cs.mcgill.ca/~jpineau/ReproducibilityChecklist.pdf) only help if the artifact is yours.
- **Determinism.** The same input produces different outputs; tests are probabilistic, SLAs are guesses, incident response is firefighting instead of forensics.
- **Cost predictability.** Frontier pricing is a band, not a number — $0.20 to $10 per million input tokens, output roughly 5x that, [by TypeSafe's own comparison](https://typesafe.ai/blog/introducing-system-one-models-and-jev). Your CFO cannot model it and your pricing team cannot pass it through.
- **Locality.** Every decision is a network round trip, and every round trip is a point of failure. Air-gapped and edge deployments are off the table without something like [llama.cpp](https://github.com/ggml-org/llama.cpp) — which is why [on-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem) rather than a model choice.

These are the properties that make software engineering a discipline rather than a craft. Give them up and you are not building a system; you are building a demo that happens to be in production. And you are paying for it on every event, forever.

## Why It Feels Right Anyway

The antipattern survives because it maps to how we think about intelligence. We imagine the agent as a brain: the brain sees the event, thinks, and acts. So we reach for the most brain-like thing we have and wire it into the event loop.

But production agents are not brains; they are *systems*, and systems have a distinction the brain metaphor hides. **Some work is done once. Some work is done forever.**

Figuring out *what matters* about an event is design work: it happens once, and it is allowed to be slow, expensive, and messy. Deciding *what to do* about a specific event is operational work: it happens millions of times and must be fast, cheap, and deterministic. The antipattern collapses the two, re-deriving the design on every event — the same confusion of capital expenditure with operational expenditure that shows up in every generation. [An agent run is not automation](https://blog.hackspree.com/#task-automation-economics); automation is the point at which the marginal cost of the next unit falls to near zero.

A compiler would never do this. You do not re-parse your source every time you run your program. The whole discipline is built on the premise that design time and runtime are different places, and that artifacts move between them.

## The Waypoint: Decision Models

The first correction people reach for is the **decision model**, and it is a real improvement. Jev and Laya strip out token generation entirely: you send state and typed questions, and you get back typed answers — a choice from a fixed set, a score on a rubric, a calibrated yes/no probability. TypeSafe's pitch is "a frontier-intelligence function call: unstructured state in, typed probabilistic decisions out" ([Introducing System One Models & Jev](https://typesafe.ai/blog/introducing-system-one-models-and-jev), 15 September 2026); Convai's Laya is the Apache-2.0 [open-weight counterpart](https://huggingface.co/convaiinnovations/laya), answering in a single forward pass (~33 ms on GPU) and never emitting text, "[so] there is nothing to parse and nothing to hallucinate" — a category of failure that generative systems cannot structurally avoid ([survey](https://arxiv.org/abs/2202.03629)).

Jev's launch post is worth reading as an economic document, and TypeSafe is candid enough about the numbers to make them usable. The published workload is four **decision DAGs** — the shape a production workflow takes once you stop prompting for prose and start writing code around a model:

[![TypeSafe AI's simplest published workflow — a decision compute graph: class/urgency/theme fanned out in parallel, then deterministic code for renewal date, overage, churn likelihood level, then four branches into save/loop, billing/inquiry, inquiry, and user conflict; Jev sits in the decisions and ordinary code sits everywhere else](https://framerusercontent.com/images/ih1bFwZGYJxlnijbTuXx3f9NeM.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

Run the same workflow against frontier models, and the Pareto picture is stark — Jev alone at ~300 ms and ~78% agreement with the smartest models, against 1.4 s to 18 s for everything else:

[![TypeSafe workflow evals — System One task accuracy vs. median latency per query: Jev at ~296 ms and 78%, Fable 5.1 at ~1.43 s and 73%, GPT-5.6 Terra at ~2.19 s and 74%, GPT-6 Astra at ~15.6 s and 72%, Gemini 3.5 Pro at ~18.1 s and 64%](https://framerusercontent.com/images/z4Uu1YpJeEZPBSMTCMI0CN2PX0.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

The type-safety numbers — the ones that matter if a decision is buried in a dependency chain — are a different order of magnitude again:

[![TypeSafe type-safety eval — wrong tool calls and type errors over 1,000 samples: Jev 0% / 0% (schema matching guaranteed), GPT-5.6 Terra 1.0% wrong tool calls and 10% type errors, GPT-6 Astra 0.4% and 0%, with latency bands of ~289 ms vs ~1.6 s to ~16.1 s](https://framerusercontent.com/images/KEoJ6ZaJkOZG6mcjBsOlB3NCqek.png)](https://typesafe.ai/blog/introducing-system-one-models-and-jev)

Three things in that post are structural rather than promotional. **Type errors are an architectural property, not a model-quality property:** if the output space is fixed at request time and the model never writes strings, schema violations stop being a probability and become impossible — a compile-time guarantee smuggled into a runtime call. **Calibration is what makes a decision usable:** a model that can do a task 95% of the time but cannot say which 5% is not automatable, because the graph has no branch for it; RLCD trains honesty directly by making a strictly proper scoring rule the only way to maximise reward. And **the numbers are footnoted on the page:** the workflows were written by TypeSafe's own capabilities team, the reference probabilities are an average of GPT-6 Astra and Fable 5.1 (which biases toward OpenAI and Anthropic), the baseline LLMs run inside TypeSafe's own structured-output wrapper, and the type-error figure is "not empirical — schema matching is guaranteed." The [independent coverage](https://aimodelreport.com/articles/2026-09-21-jev-and-laya-introduce-a-new-model-class-non-autoregressive-decision-models-that/) flags the same caveats.

Now the part that matters more than the speedups.

**The demo is a variable cost wearing a costume.** TypeSafe's team wired Jev into Doom at ten queries per second and were pleased it cost "~$7/hour" — roughly $61,000 a year to play one game, scaling with every additional player and frame. That line was published as a celebration; it is the marginal-cost problem stated by the model's own authors.

**The naming is the tell.** Jev is named after William Stanley Jevons — "every order of magnitude drop in the cost of intelligence unlocks orders of magnitude more use cases." That is correct, and it is the strongest argument against putting intelligence in the event loop. Jevons' paradox says efficiency gains increase total consumption: cheaper per-call decisions do not shrink the variable-cost line on your P&L, they grow it, because now you call the model where you previously wrote an `if`.

**The waypoint still costs what a neural network costs.** Jev is a hosted, closed API with no weights, no parameter counts, and no self-hosting, versioned on someone else's release schedule. Laya is open, and its card shows the bill in kind: 421M (English) and 322M (multilingual) parameters, an ~808 MB English checkpoint, a GPU to hit 33 ms, a per-option token budget that degrades past ~20 options (0.425 on Banking77's 77 labels, where Jev scores 0.870), calibration error that starts at 0.466 and reaches 0.081 only after fitting a temperature per question type on your domain, and an English checkpoint that shreds non-Latin scripts *while staying confident* — 0.000 accuracy on Khmer at 0.952 mean confidence, so no confidence gate can protect you. The card's own verdict is the thesis of this post: "Treat Laya as a fast foundation model to specialize, not as an omniscient zero-shot oracle." It is all in Laya's [benchmark report](https://laya.convaiinnovations.com/), which is unusually honest about ceilings.

So: take the waypoint seriously, use it where it fits, and treat it as an on-ramp. The decision model is how you *discover* the decision procedure. It is not how you *own* it.

## The Missing Layer in AutoML

The obvious next move is AutoML, and it usually stops too early. [auto-sklearn](https://arxiv.org/abs/2007.04074), [TPOT](https://epistasislab.github.io/tpot/), and [H2O AutoML](https://docs.h2o.ai/h2o/latest-stable/h2o-docs/automl.html) automate the *syntactic* pipeline: imputation, encoding, scaling, model families, hyperparameters, ensembling. It is genuinely powerful and it cannot *understand* your event schemas.

It cannot look at a stream of `OrderPlaced`, `RefundRequested`, and `SupportTicketOpened` and say: *"the ratio of refund requests to purchases over a rolling 90-day window is probably predictive of churn."* That is a semantic step; it requires knowing what the events mean. AutoML searches. It does not propose. Even the most ambitious classical answer, [Deep Feature Synthesis](https://featuretools.alteryx.com/en/stable/), mechanically composes primitives over relational structure — more ambition than most pipelines have, and still syntactic.

That is the gap where an LLM belongs: not at runtime, but at compile time, as the **semantic layer of the AutoML stack**. It is no longer speculative. CAAFE ([Hollmann, Müller, Hutter](https://arxiv.org/abs/2305.03403), NeurIPS 2024) puts an LLM in the loop to read dataset context and propose new features, then evaluates them by cross-validation instead of trusting them; [AutoML-GPT](https://arxiv.org/abs/2305.02499) extends the idea to pipeline construction. The pattern is being published — mostly on the wrong side of the wall. This blog has argued the same shape from the enterprise rule side, where the LLM's job is [extracting decision models at design time](https://blog.hackspree.com/#on-rule-engines-automating-decision-models) rather than adjudicating at runtime.

## The Wall

The fix is a hard boundary between compile time and runtime.

```text
COMPILE TIME — slow, messy, expensive, human-reviewed
  events    ─┐
  actions   ─┼─▶ LLM ─▶ candidate    ─▶ engine ─▶ AutoML ─▶ frozen
  objective ┘  semantic  feature spec    compiles   trains    artifacts
               layer     (40-100)       PIT folds  prunes    + hashes
════════════════════════ HARD WALL ════════════════════════
RUNTIME — fast, deterministic, replayable, versioned
  artifacts ─▶ verify hashes ─▶ Vector() ─▶ tree model ─▶ action ─▶ tool
```

Everything to the left is allowed to be slow, expensive, and messy. Everything to the right is pure code and frozen data. The compile-time LLM reads the declared event schemas, the tool registry, and the objective, and proposes candidates a syntactic search would never generate:

```yaml
candidate_features:
  - name: entity_age_days
    kind: numeric
    from: EntityCreated.timestamp
  - name: update_velocity_30d
    kind: numeric
    from: EntityUpdated count over 30d window
  - name: signal_ratio_90d
    kind: numeric
    from: EntityObserved.signal / EntityCreated count over 90d
  # ... 40-100 candidates total
```

The engine then folds the event stream into a projection, generates a `Vector()` function, and enforces point-in-time correctness structurally — the guarantee feature platforms sell as [point-in-time joins](https://docs.feast.dev/getting-started/concepts/point-in-time-joins), and the reason [Chronon](https://github.com/airbnb/chronon) exists. AutoML trains on the full candidate set, ranks by permutation importance, measures temporal generalization, drops the bottom 60–80%, and freezes the survivors.

What crosses the wall is only declarative artifacts and their hashes:

```text
spec.yaml     → H1  → binds features to model
model.cbm     → H2  → binds weights to features
schemas.yaml  → H3  → binds events to features
actions.yaml  → H4  → binds decisions to tools
```

At runtime the daemon verifies `H1(spec) == model.spec_hash`, `H3(schemas) == spec.schema_hash`, and `H4(actions) == spec.action_hash` before it acts. Mismatch is a hard refusal: if an artifact is not in the frozen set, it cannot affect the decision. This is the discipline [reproducible builds](https://reproducible-builds.org/) and [SLSA](https://slsa.dev/) apply to binaries, applied to a decision procedure.

**This is a capital expenditure.** You pay once, at design time, for an artifact, and amortize it across every decision the system ever makes. The marginal cost of the millionth decision is zero.

And there is no partial wall. Every way the antipattern sneaks back in — an LLM call inside the fold, a decision model inside `Vector()`, a neural fallback for low-confidence cases — breaks replayability, auditability, versioning, and determinism at once, and re-adds a variable cost that survives the compile-time conversion. Every exception is a subscription you forgot you signed.

## Why a Small Frozen Model Wins

What replaces the decision model is a feature vector and a small traditional model: [CatBoost](https://catboost.ai/), [XGBoost](https://xgboost.readthedocs.io/en/stable/tutorials/saving_model.html), [LightGBM](https://lightgbm.readthedocs.io/en/latest/), a [scikit-learn histogram gradient booster](https://scikit-learn.org/stable/modules/ensemble.html), or whatever the search selects. Saving and freezing it is a one-liner; the artifact is a few megabytes, not a checkpoint directory.

The evidence that this is not a downgrade is old and unglamorous. On tabular data, Shwartz-Ziv and Armon's [Tabular Data: Deep Learning Is Not All You Need](https://arxiv.org/abs/2106.03253) shows XGBoost beating the deep tabular models across 11 datasets, and Grinsztajn, Oyallon, and Varoquaux's [Why do tree-based models still outperform deep learning on typical tabular data?](https://arxiv.org/abs/2207.08815) attributes the gap to trees' robustness to uninformative features and their axis-aligned splits — which is exactly what a hand-engineered feature vector looks like. CatBoost's contribution ([Prokhorenkova et al.](https://arxiv.org/abs/1706.09516)) was ordered boosting and native categorical handling, so the boring default also eats your `category` column without a bespoke encoder.

And it is legible in a way a neural decision head is not: feature importance, [SHAP values](https://arxiv.org/abs/1705.07874), and the caveats in [Molnar's Interpretable Machine Learning](https://christophm.github.io/interpretable-ml-book/) are enough to explain any single decision to a regulator, a customer, or yourself at 3 AM.

The economics are **economies of scope, not economies of scale**: you win by making the problem narrower, not the model bigger. A general-purpose classifier becomes a special-purpose one, and the narrower the problem, the smaller the model, the cheaper the inference, the more explainable the decision. This is the same reason a database index beats a full table scan, a compiled function beats an interpreter, and a cache beats a recomputation — specialization is the oldest economic move in computing.

It is also where the rest of the field is already heading. AI systems are getting the treatment distributed systems got a decade ago — versioning, hashing, replayability, hard boundaries between design and execution — which is what the MLOps literature ([Kreuzberger et al.](https://arxiv.org/abs/2205.02302)) and Google's [Rules of ML](https://developers.google.com/machine-learning/guides/rules-of-ml) and [ML Test Score](https://research.google/pubs/the-ml-test-score-a-rubric-for-ml-production-readiness-and-technical-debt-reduction/) have been arguing since before the current wave. Design-time tools are absorbing work that used to happen at runtime, because [verification was always the bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering). And small, specialized models keep winning at the edge of the stack — [specialization, not size, is the lever](https://blog.hackspree.com/#specialized-small-language-models), and the deployable unit is [the model plus the harness](https://blog.hackspree.com/#better-harnesses-smaller-models) that lifts shared difficulty out of the model.

## When the Wall Is Wrong

I want to be honest about the boundaries.

The wall is not universal. Runtime LLM calls or decision models are the right choice when decisions depend on **rare, nuanced, or unstructured signals** that cannot be summarized into features; when **event schemas change frequently**, making retraining expensive; when you need **zero-shot generalization** to situations the model has never seen — exactly what Laya's base checkpoints are honest about not providing; when you **cannot get reliable labels**; and when the product is a **conversational interface**, where a language model is not the antipattern but the thing itself.

There is also a cold start. Without labeled examples, supervised ML has nothing to learn from, and using a runtime model as a *labeler* — writing into the training set rather than the decision path — is legitimate. But those labels are only as good as the model's judgment, and the gap between expressed and actual confidence in language models is [documented](https://arxiv.org/abs/2205.14334) and [still open](https://arxiv.org/abs/2409.00352). Validate against human labels before trusting the artifact in production.

In those cases the flexibility is worth the price. Be honest about the price, though: you are trading speed, cost, determinism, and auditability for it. As with every other "obvious" engineering answer, [there are no solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs).

## What This Changed in My View

Two things moved while I was writing this, and both moved toward the wall.

The first is that the decision-model waypoint arrived faster and better than I expected. I had filed "put a classifier in the loop instead of an LLM" under obvious-but-unbuilt. TypeSafe shipped it, published the workflow evals, footnoted their own biases, and priced the calls; Laya shipped open weights and a limitations section that says out loud that the base model is near random on the task it is advertised for. That honesty is what makes the case *for the wall*, because the same post that shows Jev on a Pareto frontier also shows the $7/hour Doom line — a variable cost doing exactly what variable costs do.

The second is that the argument is mostly about **artifacts, not models**. I came in thinking the choice was "big model at runtime" versus "small model at compile time." The stronger framing is "a call you make forever" versus "an artifact you own." A compiled spec, a frozen tree ensemble, and four hashes are not a smaller AI system; they are a different kind of thing — something you can diff in a pull request, roll back, replay, audit, and hand to someone else. That is what [conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) looks like for a decision procedure, and it is why the cost argument and the engineering argument keep landing in the same place.

What did not move: the honest cases at the boundary. Conversational interfaces are not going away, and unlabelable decisions still need something flexible. The wall is a default, not a religion.

## The Point

Compile time is where the LLM lives: it reads your event schemas, your tool registry, and your objective, and proposes candidate features that a syntactic search would never generate. An engine compiles them into point-in-time-correct folds. Traditional ML trains, prunes, and freezes them. Runtime is where neither the LLM nor the decision model appears — the daemon loads the frozen artifacts, verifies their hashes, builds a feature vector, asks the tree model for an action, maps it to a tool, and publishes. That artifact runs anywhere a static binary runs, [ONNX Runtime](https://onnxruntime.ai/) or [its Go binding](https://github.com/yalue/onnxruntime_go) included, with the network unplugged.

Decision models like Jev and Laya are a meaningful improvement over runtime LLMs: no token generation, hundreds of milliseconds instead of seconds, guaranteed type safety, calibrated confidence. They are also still neural networks with per-call pricing, hundreds of megabytes of weights, a GPU to hit their stated latency, and a per-option budget that degrades as your label space grows. A waypoint, not a destination.

---

## A Call to the Community

We are at the beginning of a discipline. AI systems engineering does not have its *Design Patterns* yet. It does not have its *Site Reliability Engineering* yet. It does not have its *Choose Boring Technology* yet.

What it has is a lot of demos in production and a growing sense that something is wrong. The sense is right. The wrongness is economic as much as it is architectural.

So here is what I am asking:

- **Stop putting LLMs in the event loop.** Not because they are bad, but because the event loop is the wrong place for them. The right place is design time, as the semantic layer of your AutoML pipeline.
- **Start treating prompts as source code.** Version them, review them, hash them, compile them into artifacts. Do not let them run free at runtime, and do not let them become a subscription you cannot cancel.
- **Measure your traditional models.** Do not assume a 400M-parameter model beats a 2 MB tree on your task. Test it; you will be surprised more often than you expect.
- **Build the wall.** Draw the line between compile time and runtime in your architecture, enforce it with hashes, and refuse to run when the wall is broken — or when a variable cost sneaks past.
- **Write down your unit economics.** Know your cost per decision. Know how it scales. Know what it would look like if it were zero. Then ask yourself why it isn't.
- **Write down what you learn.** This discipline is being invented right now, in production systems, by people too busy shipping to write blog posts. If you have built one of these systems, write it up. The community needs your scars.

The next decade of AI systems will not be won by the teams with the biggest models. It will be won by the teams who understand that intelligence belongs at design time and execution belongs at runtime — and who build the wall between them.

Brooks told us there is no silver bullet. Boehm told us what the delay costs. We have known this for forty years.

Let's apply it.

---

## References

**The economics of where you pay**

- Brooks, F. P. [*No Silver Bullet — Essence and Accident in Software Engineering*](https://www.cs.unc.edu/techreports/86-020.pdf) — UNC tech report TR86-020, 1986 (later IEEE Computer 20(4), 1987). Source of "essence vs. accident" and the "hardest single part... is deciding precisely what to build" quote. See also [*The Mythical Man-Month*](https://archive.org/details/mythicalmanmonth00broo) (1975), and this blog's [Brooks series](https://blog.hackspree.com/#brooks-design-conceptual-integrity) and [Accidental Complexity Is the Only Complexity You Can Remove](https://blog.hackspree.com/#brooks-accidental-complexity).
- Boehm, B. and Basili, V. [*Software Defect Reduction Top 10 List*](https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf) — IEEE Computer 34(1), January 2001, pp. 135–137. Item 1 is the 100x-after-delivery figure; the same page gives the 40–50% avoidable-rework estimate. The 1 : 10 : 100 gradient is Boehm's cost-of-change curve from [*Software Engineering Economics*](https://www.informit.com/store/software-engineering-economics-9780138221225) (Prentice Hall, 1981).
- NIST. [*The Economic Impacts of Inadequate Infrastructure for Software Testing*](https://www.nist.gov/system/files/documents/director/planning/report02-3.pdf) — Planning Report 02-3, May 2002. The $59.5B estimate that turns Boehm's curve into a budget line.
- Futamura, Y. [*Partial Evaluation of Computation Process — An Approach to a Compiler-Compiler*](https://link.springer.com/article/10.1023/A:1010095604496) — Higher-Order and Symbolic Computation 12, 1999 (originally 1971). Specializing a program to its known inputs is where runtime cost goes to die.
- Winand, M. [*Anatomy of an Index*](https://use-the-index-luke.com/sql/anatomy). An index is a design-time artifact that removes a runtime derivation.
- This blog: [Engineering is art and philosophy, grounded in economic law](https://blog.hackspree.com/#engineering-is-economics), [The economics of the dark factory](https://blog.hackspree.com/#dark-factory-economics), [Task automation economics: why an agent run is not automation](https://blog.hackspree.com/#task-automation-economics), [No solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs).

**The decision model at runtime (and why it is a waypoint)**

- Almeida, D. [*Introducing System One Models & Jev*](https://typesafe.ai/blog/introducing-system-one-models-and-jev) — TypeSafe AI, 15 September 2026. Source of the three figures reproduced above (the published decision DAG, the workflow-eval Pareto plot, the type-safety evaluation), pricing ($0.042/MTok input, output free, against a $0.20–$10 frontier band with output ~5x input), latency (70–500 ms vs. 3–329 s), the 193.6x faster / 444.6x cheaper home-page claim, the workflow-eval methodology (identical workflow per model, reference probabilities averaged from GPT-6 Astra and Fable 5.1), the Doom demo at ten queries per second for ~$7/hour, the Wikiracing demo, and the Kahneman and Jevons naming. Its own "Nuance" notes disclose the capabilities-team authorship, the OpenAI/Anthropic-biased reference, and that type safety is guaranteed by construction rather than measured.
- [*Jev AI Model (TypeSafe) — Typed System One Decisions*](https://jevmodel.org/) — third-party guide, updated 21 September 2026: API access, the $42-per-billion input price, `jev-1.13.0` aliases, closed weights, and the explicit contrast with Laya.
- Convai Innovations. [*Laya*](https://huggingface.co/convaiinnovations/laya) — 421M ModernBERT-large English, 322M mmBERT-base multilingual, and a typed-decisions checkpoint under one Apache-2.0 repo; the `choice` / `score` / `noul` primitives; ~33 ms p50 and 7.2 ms per question batched; the ~808 MB English checkpoint; and the card's own limits (base checkpoints ~0.35 against a 0.318 random baseline; 0.766 only after fine-tuning; ECE 0.466 → 0.081 after per-question-type temperature fitting; degradation past ~20 options).
- Convai Innovations. [*Laya — 33ms Multilingual System 1 Decision Engine*](https://laya.convaiinnovations.com/) — RLCD and strictly proper scoring rules, the 51-language MASSIVE sweep (0.000 accuracy on Khmer at 0.952 mean confidence), the Banking77 comparison (0.425 vs. Jev's 0.870 at 77 labels), and the March 2025 [SalesRLAgent](https://arxiv.org/abs/2503.23303) and September 2025 [schema-based decision framework](https://arxiv.org/abs/2510.01237) behind it.
- [*NandhaKishorM/laya*](https://github.com/NandhaKishorM/laya) — SDK, router, and benchmark harnesses; [`laya-typed-decisions`](https://huggingface.co/convaiinnovations/laya-typed-decisions) is the checkpoint they run against.
- Iverson, L. [*Jev and Laya define a new model class: typed-decision heads that replace LLM calls for routing*](https://aimodelreport.com/articles/2026-09-21-jev-and-laya-introduce-a-new-model-class-non-autoregressive-decision-models-that/) — AI Model Report, 21 September 2026. Independent framing, the 32.8–276 ms and $0.042/MTok comparisons, the zero-shot 0.362-vs-0.318-baseline caveat, and developer reports (5–18x faster on a command-safety classifier; 10–20x cheaper on email classification).

**LLMs at compile time: AutoML and the semantic feature layer**

- Hollmann, N., Müller, S., and Hutter, F. [*CAAFE: Context-Aware Automated Feature Engineering*](https://arxiv.org/abs/2305.03403) — NeurIPS 2024. The closest published instance of the pattern: an LLM proposes features that cross-validation then judges.
- Zhang, S. et al. [*AutoML-GPT: Automatic Machine Learning with GPT*](https://arxiv.org/abs/2305.02499) — LLM-driven pipeline construction.
- Feurer, M. et al. [*Auto-sklearn 2.0*](https://arxiv.org/abs/2007.04074); [TPOT](https://epistasislab.github.io/tpot/); [H2O AutoML](https://docs.h2o.ai/h2o/latest-stable/h2o-docs/automl.html) — the syntactic search layer the semantic layer feeds.
- Kanter, J. M. and Veeramachaneni, K. [*Deep Feature Synthesis*](https://featuretools.alteryx.com/en/stable/) — Featuretools; the classical attempt to automate feature construction, and where syntactic search stops.

**Why a small tree beats a big network on a structured decision**

- Shwartz-Ziv, R. and Armon, A. [*Tabular Data: Deep Learning Is Not All You Need*](https://arxiv.org/abs/2106.03253) — XGBoost beats the deep tabular models across 11 datasets.
- Grinsztajn, L., Oyallon, E., and Varoquaux, G. [*Why do tree-based models still outperform deep learning on typical tabular data?*](https://arxiv.org/abs/2207.08815) — NeurIPS 2022; the mechanism is what engineered features exploit.
- Prokhorenkova, L. et al. [*CatBoost: unbiased boosting with categorical features*](https://arxiv.org/abs/1706.09516) — [catboost.ai](https://catboost.ai/) / [model prediction API](https://catboost.ai/docs/en/concepts/python-reference_catboost_predict); [XGBoost model serialization](https://xgboost.readthedocs.io/en/stable/tutorials/saving_model.html); [LightGBM](https://lightgbm.readthedocs.io/en/latest/); [scikit-learn histogram gradient boosting](https://scikit-learn.org/stable/modules/ensemble.html). "Freeze a few megabytes of trees and score in milliseconds on one core" is a supported, boring operation in all of them.
- Lundberg, S. and Lee, S.-I. [*A Unified Approach to Interpreting Model Predictions*](https://arxiv.org/abs/1705.07874) — SHAP; Molnar, C. [*Interpretable Machine Learning*](https://christophm.github.io/interpretable-ml-book/) for what an explanation licenses you to claim.

**Artifacts, hashes, and the wall**

- [ONNX](https://onnx.ai/) / [ONNX Runtime](https://onnxruntime.ai/) and the [ONNX Runtime Go bindings](https://github.com/yalue/onnxruntime_go) — run a compiled graph without a scientist in the loop, embedded in a static binary.
- [SLSA](https://slsa.dev/) and [Reproducible Builds](https://reproducible-builds.org/) — provenance and content addressing, and "refuse to run when the digest does not match," which is the mechanism the wall needs.
- [JSON Schema](https://json-schema.org/), [Great Expectations](https://greatexpectations.io/), [Pandera](https://pandera.readthedocs.io/), [dbt tests](https://docs.getdbt.com/docs/build/data-tests), [DVC](https://dvc.org/) — the feature contract as a data contract: declared, validated at build time, versioned at rest.
- Feast. [*Point-in-time joins*](https://docs.feast.dev/getting-started/concepts/point-in-time-joins); Airbnb's [Chronon](https://github.com/airbnb/chronon) — compiling features so no fold can see the future.

**Determinism, calibration, and auditability**

- [*Defeating Nondeterminism in LLM Inference*](https://thinkingmachines.ai/blog/defeating-nondeterminism-in-llm-inference/) — Thinking Machines. Why the same prompt and the same weights still do not reproduce a decision.
- Kadavath, S. et al. [*Language Models (Mostly) Know What They Know*](https://arxiv.org/abs/2205.14334); [*Does Alignment Tuning Really Break LLMs' Internal Confidence?*](https://arxiv.org/abs/2409.00352) — expressed versus actual confidence, the property a decision graph has to branch on.
- Ji, Z. et al. [*Survey of Hallucination in Natural Language Generation*](https://arxiv.org/abs/2202.03629) — a failure mode type-safe decision models structurally cannot have.
- Mitchell, M. et al. [*Model Cards for Model Reporting*](https://arxiv.org/abs/1810.03993); Pineau, J. [*Reproducibility Checklist*](https://www.cs.mcgill.ca/~jpineau/ReproducibilityChecklist.pdf) — documentation and reproducibility as preconditions for audit.
- Sculley, D. et al. [*Machine Learning: The High-Interest Credit Card of Technical Debt*](https://research.google/pubs/pub43146/) (2014); [*Hidden Technical Debt in Machine Learning Systems*](https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf) (NeurIPS 2015); Breck, E. et al. [*The ML Test Score*](https://research.google/pubs/the-ml-test-score-a-rubric-for-ml-production-readiness-and-technical-debt-reduction/); Google's [Rules of ML](https://developers.google.com/machine-learning/guides/rules-of-ml); Kreuzberger, D. et al. [*MLOps: Overview, Definition, and Architecture*](https://arxiv.org/abs/2205.02302).
- [EU AI Act, Regulation (EU) 2024/1689](https://eur-lex.europa.eu/eli/reg/2024/1689/oj) — see the [AI Act Explorer](https://artificialintelligenceact.eu/ai-act-explorer/) for the logging, documentation, and oversight articles that make "we called a hosted model" an unsatisfying answer.

**Small, local, and specialized**

- Gerganov, G. et al. [llama.cpp](https://github.com/ggml-org/llama.cpp); Xu, M. et al. [*On-Device Language Models: A Comprehensive Review*](https://arxiv.org/abs/2409.00088) — running models that never needed the cloud.
- This blog: [SSLM: Specialized Small Language Models](https://blog.hackspree.com/#specialized-small-language-models), [Nanbeige4.2-3B](https://blog.hackspree.com/#nanbeige4-2-3b-agentic-model), [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models), [On-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem), [Go Can Keep Structured LLM Runtimes Boring](https://blog.hackspree.com/#go-can-keep-structured-llm-runtimes-boring).

**Compile time in the enterprise: rules and decision models**

- This blog: [On Rule Engines — Automating Decision Models](https://blog.hackspree.com/#on-rule-engines-automating-decision-models) and [Five Patterns for Composite AI](https://blog.hackspree.com/#on-rule-engines-five-patterns) — the KU Leuven program on extracting DMN decision models from text, and the composite-AI patterns that keep extraction at design time. See also Goossens, A., De Smedt, J., and Vanthienen, J. [*Extracting DMN Models from Text Using Deep Learning Techniques*](https://doi.org/10.1016/j.eswa.2022.118667) — Expert Systems with Applications 211, 2023.

**Durable execution, replay, and events as truth**

- [Every workflow is an FSM. Not every FSM is a workflow.](https://blog.hackspree.com/#every-workflow-is-an-fsm), [Durable Daemons](https://blog.hackspree.com/#durable-daemons), [Stories from Events](https://blog.hackspree.com/#stories-from-events), [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering), [Agents Are Too Stochastic for Intuition](https://blog.hackspree.com/#data-driven-design-swe-agents), [Harnessing Agentic AI Systems: A Pattern Language](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems), [SWE-Agent Economics](https://blog.hackspree.com/#why-i-focused-my-research-on-swe-agent-economics).

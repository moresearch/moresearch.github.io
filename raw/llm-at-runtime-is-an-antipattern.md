---
title: "LLM at Runtime Is an Antipattern"
date: 2026-09-23
slug: llm-at-runtime-is-an-antipattern
summary: "An LLM in the event loop converts a fixed cost into a variable cost, and gives up replayability, auditability, versioning, and determinism in the process — and Jevons' paradox means efficiency will not save the bill: it lowers the unit price, elasticity raises the volume, and only a change in the shape of the cost escapes the rebound. DSPy already taught the right discipline — declare signatures, compose modules, compile against a metric — but what it compiles to is still executed by a model. Keep the principles and change the artifact: the LLM belongs at compile time, as the semantic layer of the AutoML stack, proposing candidate features that an engine compiles and a small tree model freezes; the soft labels come from a tabular foundation model (TabPFN) rather than a token-priced LLM. Jev and Laya are a real advance over generative LLMs at runtime — and still the wrong shape for the event loop, because a decision model is a smaller variable cost, not a fixed one."
tags: llm, runtime, compile-time, automl, semantic-feature-engineering, dspy, compound-ai-systems, prompt-compilation, tabpfn, tabular-foundation-models, knowledge-distillation, soft-labels, decision-models, system-one, jev, laya, jevons-paradox, rebound-effect, software-economics, catboost, tabular-ml, brooks, boehm, determinism, auditability, antipattern, essay
---

Fred Brooks observed that the hardest part of software engineering is not building the thing, but building it in a way that survives change:

> "The hardest single part of building a software system is deciding precisely what to build... No other part of the work so cripples the resulting system if done wrong." — [No Silver Bullet — Essence and Accident in Software Engineering](https://www.cs.unc.edu/techreports/86-020.pdf) (1986)

Barry Boehm priced the delay. A defect caught at design time costs one unit; caught at test time, ten; caught in production, a hundred. Boehm and Basili's [Software Defect Reduction Top 10 List](https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf) opens with exactly that number — "Finding and fixing a software problem after delivery is often 100 times more expensive than finding and fixing it during the requirements and design phase" — and adds that 40 to 50 percent of project effort goes to avoidable rework.

Neither man was writing about AI. They were writing about a discipline that had learned, painfully, that **where you pay a cost determines everything about the system that results**. We are now building AI systems and unlearning it in public: an LLM in the event loop, priced per million tokens, running forever.

---

## The Shape We Keep Choosing

Today the shape is this:

```text
raw_event ──▶ LLM ──▶ action
             (per event, forever)
```

Every decision. Every event. Every tick. A network call to a nondeterministic, slow, expensive, unversionable model. It works. It demos. It ships. And then it bleeds.

One economic fact decides the rest of this argument: **an LLM in the event loop has a nonzero marginal cost per decision, forever.** Every call costs tokens. Every token costs money. Every dollar scales with traffic. There is no volume discount that takes the marginal cost to zero, no learning curve that bends it downward, no amortization that spreads it across future decisions. You have taken what should be a **fixed cost** — the design of a decision procedure — and turned it into a **variable cost** that compounds with every event the system processes.

Software economics has one rule for this: **convert variable costs into fixed costs as early as possible.** That is what compilation is (the [Futamura projections](https://link.springer.com/article/10.1023/A:1010095604496) are its purest statement). That is what caching is. That is what an [index](https://use-the-index-luke.com/sql/anatomy) is: a precomputed answer to a question you would otherwise re-derive on every scan. Brooks called the difficulty that survives good design "essential complexity"; an LLM in the event loop does not reduce it, it re-pays for it on every event in the most expensive currency available.

The ML engineering literature named this failure mode years ago — [Machine Learning: The High-Interest Credit Card of Technical Debt](https://research.google/pubs/pub43146/) and [Hidden Technical Debt in Machine Learning Systems](https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf), whose CACE principle ("Changing Anything Changes Everything") is precisely what an unversioned model in the loop does to every downstream decision.

It is not flexibility. It is a subscription to your own architecture.

## What You Give Up

The cost is not only latency and dollars. It is the set of engineering properties you surrender, each with its own price.

- **Replayability.** You cannot reproduce yesterday's decision: the model no longer exists in the form it existed yesterday, and even with identical weights [inference is not bit-deterministic](https://thinkingmachines.ai/blog/defeating-nondeterminism-in-llm-inference/). Incident review becomes guesswork and regression becomes mystery — the same problem durable execution solved for code, and the same lesson this blog keeps reaching from the workflow side: [graphs describe intentions, logs describe executions](https://blog.hackspree.com/#every-workflow-is-an-fsm).
- **Auditability.** You cannot explain why the system acted, because the reasoning lives inside weights you do not own. The EU AI Act's [documentation and logging obligations](https://eur-lex.europa.eu/eli/reg/2024/1689/oj) assume you *can* record and explain behaviour.
- **Versionability.** You cannot hash "GPT-4-turbo-as-of-Tuesday," pin it, roll it back, or A/B two versions without two vendor contracts. [Model cards](https://arxiv.org/abs/1810.03993) and [reproducibility checklists](https://www.cs.mcgill.ca/~jpineau/ReproducibilityChecklist.pdf) only help if the artifact is yours.
- **Determinism.** The same input produces different outputs; tests are probabilistic, SLAs are guesses, incident response is firefighting instead of forensics.
- **Cost predictability.** Frontier pricing is a band, not a number — $0.20 to $10 per million input tokens, output roughly 5x that, [by TypeSafe's own comparison](https://typesafe.ai/blog/introducing-system-one-models-and-jev). Your CFO cannot model it and your pricing team cannot pass it through.
- **Locality.** Every decision is a network round trip, and every round trip is a point of failure. Air-gapped and edge deployments are off the table without something like [llama.cpp](https://github.com/ggml-org/llama.cpp) — which is why [on-device LLMs are a systems design problem](https://blog.hackspree.com/#on-device-llms-are-a-systems-design-problem) rather than a model choice.

These are the properties that make software engineering a discipline rather than a craft. Surrender them and you are not building a system; you are running a demo in production and paying for it on every event, forever.

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

Three consequences matter more than the speedups.

**The celebrated demo is the antipattern in miniature.** TypeSafe's team wired Jev into Doom at ten queries per second and were pleased it cost "~$7/hour" — roughly $61,000 a year to play one game, scaling with every additional player and frame. That line was published as a triumph; it is the marginal-cost problem stated by the model's own authors.

**The naming is the tell.** Jev is named after William Stanley Jevons, and TypeSafe's FAQ states the thesis outright: "every order of magnitude drop in the cost of intelligence unlocks orders of magnitude more use cases." That sentence is correct, it is the strongest argument against putting intelligence in the event loop, and it deserves more than a name-drop.

**The waypoint still costs what a neural network costs.** Jev is a hosted, closed API with no weights, no parameter counts, and no self-hosting, versioned on someone else's release schedule. Laya is open, and its card shows the bill in kind: 421M (English) and 322M (multilingual) parameters, an ~808 MB English checkpoint, a GPU to hit 33 ms, a per-option token budget that degrades past ~20 options (0.425 on Banking77's 77 labels, where Jev scores 0.870), calibration error that starts at 0.466 and reaches 0.081 only after fitting a temperature per question type on your domain, and an English checkpoint that shreds non-Latin scripts *while staying confident* — 0.000 accuracy on Khmer at 0.952 mean confidence, so no confidence gate can protect you. The card's own verdict is the thesis of this post: "Treat Laya as a fast foundation model to specialize, not as an omniscient zero-shot oracle." It is all in Laya's [benchmark report](https://laya.convaiinnovations.com/), which is unusually honest about ceilings.

Take the waypoint seriously, use it where it fits, and treat it as an on-ramp. The decision model is how you *discover* the decision procedure. It is not how you *own* it.

## The Jevons Trap

In 1865, William Stanley Jevons noticed something he found embarrassing: as British steam engines got dramatically more efficient — more work per ton of coal — the country did not burn less coal. It burned vastly more, because cheaper power made previously pointless things worth doing. Jevons called his own conclusion "paradoxical"; economists later named the mechanism the **rebound effect**. Efficiency lowers the *effective price* of a unit of service, and when demand for that service is responsive to price, total consumption rises even as unit cost falls. Whether the rebound is partial or complete is an empirical argument the energy-policy literature has been having for four decades ([Gillingham, Rapson, and Wagner](https://www.nber.org/papers/w20424)) — and the deciding variable is elasticity, not efficiency.

Four independent strands of AI research have now arrived at the same place, and none of them was written to sell a model:

- [Luccioni, Strubell, and Crawford](https://arxiv.org/abs/2501.16548) argue that rebound effects "undermine the assumption that improved technical efficiency alone will ensure net reductions" — as they put it, the trajectory hinges on business incentives and market structure, not on engineering.
- [Zhang and Zhang](https://arxiv.org/abs/2601.12339) model language models as an asset class and find that falling inference prices push firms toward *more* compute-intensive agent architectures, making aggregate compute demand super-elastic: a **structural Jevons paradox**.
- [Sharma's thermodynamic model of the cloud](https://arxiv.org/abs/2411.11540) isolates the actual driver, and it is not efficiency: it is *system growth*. Hyperscale platforms get more efficient and consume more in the same motion.
- And [Husom et al.](https://arxiv.org/abs/2608.17055) put the empirical point plainly — efficiency work "has not translated into reduced consumption" — which is why they reach for *prevention*: the cheapest unit of work is the one you never run.

The industry supplies its own confirmation, in public. TypeSafe named a model after Jevons, priced it at $0.042 per million input tokens against a frontier band of $0.20 to $10, and demoed ten model calls a second playing Doom at "~$7/hour" — reported as *lower than expected* by people who had just made intelligence cheap. That is the rebound, measured three days into a product launch.

So the conclusion is not "use cheaper inference." **Efficiency is not a strategy; decoupling is.** If the driver of consumption is growth, and the unit cost stays positive, then growth keeps buying more consumption: cheaper calls become more calls, the model lands where you previously wrote an `if`, and the bill you were avoiding arrives with more line items. The only move that escapes the trap is changing the *shape* of the cost — pay once, at compile time, for an artifact, and let the millionth decision cost nothing. At that point the rebound stops being a threat and becomes the business: more events, more decisions, more volume, flat spend, because what scales is usage rather than cost.

Be honest about the residual. The trap is relocated, not abolished — more decision procedures mean more compilations, more teacher runs, more labelled data, more design review. But those costs scale with the *number of decisions you have chosen to model*, a slow-moving number you control, instead of the *number of events the world sends you*, a fast-moving number your customers control. That difference between a capital expenditure and a subscription is the whole reason to draw the wall.

## The Missing Layer in AutoML

The obvious next move is AutoML, and it usually stops too early. [auto-sklearn](https://arxiv.org/abs/2007.04074), [TPOT](https://epistasislab.github.io/tpot/), and [H2O AutoML](https://docs.h2o.ai/h2o/latest-stable/h2o-docs/automl.html) automate the *syntactic* pipeline: imputation, encoding, scaling, model families, hyperparameters, ensembling. All genuinely powerful, and none of it *understands* your event schemas.

It cannot look at a stream of `OrderPlaced`, `RefundRequested`, and `SupportTicketOpened` and say: *"the ratio of refund requests to purchases over a rolling 90-day window is probably predictive of churn."* That is a semantic step; it requires knowing what the events mean. AutoML searches. It does not propose. Even the most ambitious classical answer, [Deep Feature Synthesis](https://featuretools.alteryx.com/en/stable/), mechanically composes primitives over relational structure — more ambition than most pipelines have, and still syntactic.

That is the gap where an LLM belongs: not at runtime, but at compile time, as the **semantic layer of the AutoML stack**. It is no longer speculative. CAAFE ([Hollmann, Müller, Hutter](https://arxiv.org/abs/2305.03403), NeurIPS 2024) puts an LLM in the loop to read dataset context and propose new features, then evaluates them by cross-validation instead of trusting them; [AutoML-GPT](https://arxiv.org/abs/2305.02499) extends the idea to pipeline construction. The pattern is being published — mostly on the wrong side of the wall. This blog has argued the same shape from the enterprise rule side, where the LLM's job is [extracting decision models at design time](https://blog.hackspree.com/#on-rule-engines-automating-decision-models) rather than adjudicating at runtime.

## The Better Teacher: TabPFN, Not the LLM

A compile-time LLM has one job in that pipeline — propose features, read the schemas, supply meaning — and one job it should not be given: producing the labels.

When the events are numeric and structured, the best available teacher is not a language model. It is a **tabular foundation model**, and the one to reach for is [TabPFN](https://github.com/PriorLabs/TabPFN). It is a prior-data fitted network: pretrained on synthetic tables drawn from a causal-model prior ([Hollmann et al.](https://arxiv.org/abs/2207.01848), [Nature 2025](https://doi.org/10.1038/s41586-024-08328-6)), then used as is, with no gradient descent and no hyperparameters. It reads your training set as *context* and returns a probability for every row in a single forward pass — which turns the compile step into a distillation ([Hinton et al.](https://arxiv.org/abs/1503.02531): soft targets carry more information per example than hard ones):

Then the numbers arrive. The strongest in-context tabular models take 151 to 1,275 ms per batch on a GPU, and a fraud scorer has under 2 ms ([Pocket Foundation Models](https://arxiv.org/pdf/2605.18654), 2026). The window for a runtime teacher closes long before the request arrives — which is exactly why the teacher belongs at compile time, and why the student has to be a tree.

```python
# compile time only — the teacher never ships
import numpy as np
from sklearn.model_selection import StratifiedKFold

soft = np.zeros(len(y_train))
for ctx_idx, hold_idx in StratifiedKFold(5).split(X_train, y_train):
    teacher = TabPFNClassifier()                     # in-context, no tuning
    teacher.fit(X_train[ctx_idx], y_train[ctx_idx])   # context: the other folds
    soft[hold_idx] = teacher.predict_proba(X_train[hold_idx])[:, 1]

student = CatBoostClassifier(loss_function="CrossEntropy")   # probability targets
student.fit(X_train, soft)               # match the distribution, not the label
student.save_model("model.cbm")          # this is what crosses the wall
```

That fold loop is not decoration; it is the whole trick. In-context teachers leak, and the paper measures how badly: on a 5-class dataset, a teacher scoring rows *outside* its context spreads probability across two or three classes (mean entropy around 0.6–0.9 nats), while the same teacher scoring rows *inside* its context puts more than 99.9% of the mass on the true class (entropy near 10⁻³ nats). At that point the soft targets are one-hot recall with extra steps, Hinton's loss collapses to hard-label cross-entropy at the wrong temperature, and the student learns nothing the labels did not already say — the paper's own summary is blunter: *skip out-of-fold labelling and ICL distillation produces students worse than hard-label training.* Stratified out-of-fold labelling is the fix (they use K=5; 3, 5, and 10 folds were indistinguishable, and the choice is a trade between labels per fold and coverage), because no row is ever labelled by a teacher whose context contained its label.

That is also the ground on which an LLM makes a poor teacher for CatBoost, and why the failure is not a bug you can patch. Six things decide a teacher, and a chat model fails five of them:

- **The teacher must be good at your data shape.** That precondition is what a tabular foundation model is *for*: default TabPFN-2.5 beats default XGBoost on 100% of small-to-medium classification datasets (≤10k points, 500 features) and 87% of larger ones up to 100k samples, and a TabPFN-3 forward pass beats 8-hour-tuned gradient-boosted baselines on tables up to 1M rows ([2.5](https://arxiv.org/abs/2511.08667), [3](https://arxiv.org/abs/2605.13986)). It is also small enough to be a laptop step: the v2 classifier is a [12.9–29 MB checkpoint](https://huggingface.co/Prior-Labs/TabPFN-v2-clf).
- **A teacher's value is the shape of its distribution, not its answer.** Distillation works because the teacher's second and third choices tell the student where the boundary is ambiguous. An LLM teacher gives you either a row it has effectively seen the label for — the collapse above, with no fold structure that can hide it, since the labelled examples *are* the task signal you need in the prompt — or a zero-shot guess whose "probabilities" are next-token scores over label strings, sensitive to naming, order, and formatting ([Liu et al.](https://arxiv.org/abs/2312.16702)) and [overconfident](https://arxiv.org/abs/2205.14334) by disposition.
- **It has to see rows as data.** The teacher's targets are only as good as its access to the table: LLMs tokenise numerals discontinuously ([xVal](https://arxiv.org/abs/2310.02989)), need rows serialised into a prompt, and remain [hard to make competitive](https://arxiv.org/abs/2510.17385) on tabular prediction even after structural-prior post-training. A tabular foundation model was trained on the thing itself.
- **It has to see all of the data.** Prompt-based teachers see tens to hundreds of rows plus a description, and each call sees a *different* handful — so the target distribution drifts from row to row instead of being a posterior over your table. TabPFN-2.5 handles [50k rows in context](https://arxiv.org/abs/2511.08667) and TabPFN-3 handles [1M](https://arxiv.org/abs/2605.13986), locally, on pinned weights.
- **The student inherits the teacher's rank, and nothing more.** The paper found that teacher rank by solo AUC on a held-out sample picks the best student-producing teacher across all four student families it tested, and that when the teacher itself trails a tuned CatBoost — on high-dimensional tables, in their results — distillation makes the student *worse* than the baseline. An LLM teacher sitting below a tuned CatBoost on your table is therefore not a neutral choice; it is a way to ship a worse model than the one you could have trained on hard labels in an afternoon.
- **Targets have to be reproducible.** The soft-label matrix is an input to the artifact that crosses the wall, so it needs a version, a hash, and a re-run that lands in the same place. Local weights with a pinned checkpoint and a frozen fold assignment give you that. A non-deterministic vendor service billed per row gives you a target set you can never regenerate and a lineage you can never defend.

None of this is hypothetical. The measurement is public: across 153 classification datasets (TALENT, OpenML-CC18, TabZilla, TabArena), distilling TabICLv2 into XGBoost reached 0.882 macro-mean AUC — **96.5% of the teacher's AUC at 1.9 ms on CPU**, a 38x to 860x latency reduction, with a statistically significant edge over a tuned CatBoost baseline on 51% of the datasets (Wilcoxon p=0.0008). Teacher choice turns out to be a one-decision problem: rank the candidates by solo AUC on a held-out sample and distil from the winner; the ranking survived distillation exactly, and for tree students a single strong teacher beat averaging several (adding a fourth teacher moved macro-mean AUC by 0.0006, and the weakest teacher subtracted). Prior Labs' TabPFN-2.5 release ships "a new distillation engine that converts TabPFN-2.5 into a compact MLP or tree ensemble"; the fusion literature gets [TabPFN-augmented GBDTs beating both components](https://arxiv.org/abs/2502.02672) across dataset sizes; [TabDistill](https://arxiv.org/abs/2511.05704) does the neural-student version. The full pipeline, including the out-of-fold labelling, is open sourced as [TabTune](https://github.com/Lexsi-Labs/TabTune).

Three things to be honest about. **The licence is a fixed cost, not a free lunch:** TabPFN-2 weights are Apache-2.0 with an attribution requirement, while the 2.5, 2.6, 3, and 3.5 weights are [non-commercial](https://github.com/PriorLabs/TabPFN) and want a commercial licence — a one-off line item, which is the shape this post keeps arguing for. **The caps move, so pin the teacher:** row and column limits went 50k to 100k to 1M within nine months ([model table](https://docs.priorlabs.ai/models)), and TabPFN's answers depend on which rows end up in context ([context sampling](https://arxiv.org/abs/2402.06971)), so version the teacher, freeze its context, and freeze the folds. And **a teacher cannot teach what it does not know:** the measured gain over a tuned CatBoost is +0.011 AUC on low-dimensional tables (≤21 features) and +0.001 above that, and it turns negative where the teacher itself loses — so distil only when the teacher wins on *your* table, then re-check the student's calibration instead of assuming it inherited the teacher's. The study's own caveats are worth carrying across: the evaluation is IID per dataset, so temporal shift, heavy missingness, and time-series structure are untested — which is precisely why your folds should be time-ordered rather than random — and its latency figures are per-query on one CPU core, a lower bound on the speedup rather than an end-to-end serving comparison.

Where the LLM remains the better teacher is where the table stops being a table: the free-text columns, the column semantics, and the question of which features should exist at all. That is CAAFE's job and the semantic layer's job. The division is clean — the LLM supplies meaning, the tabular foundation model supplies statistics, and neither crosses the wall.

## The DSPy Lesson: Compile the Program, Not the Prompt

A production AI system is not a model; it is a program of model calls, retrievals, and ordinary code — a [compound AI system](https://bair.berkeley.edu/blog/2024/02/18/compound-ai-systems/). The most principled tooling for building those is [DSPy](https://dspy.ai/) ([Khattab et al.](https://arxiv.org/abs/2310.03714)), and its principles are the ones this post has been arguing from a different direction:

- **Declare signatures, not prompts.** You state the inputs and outputs and stop hand-writing the strings that coax them out of a model ([signatures](https://dspy.ai/learn/programming/signatures/)).
- **Compose modules, not monoliths.** A compound system is a program of typed calls ([modules](https://dspy.ai/learn/programming/modules/)) that you can inspect, test, and reuse.
- **Compile against a metric.** Optimizers such as BootstrapFewShot, [MIPROv2](https://arxiv.org/abs/2406.11695), and [GEPA](https://arxiv.org/abs/2507.19457) search instructions and demonstrations on your behalf instead of leaving a human to guess at them ([optimization overview](https://dspy.ai/learn/optimization/overview/)).
- **Ship the compiled artifact.** The output of a compile is a program you save, version, and evaluate — not a prompt you keep tweaking by hand.

That is the right shape, and it stops one variable short. What DSPy compiles *to* is still a bundle of instructions and demonstrations that calls a language model on every event. It moves the *writing* to compile time and leaves the *inference* at runtime: the optimization run becomes a fixed cost, and the program it produces remains a subscription.

So keep the principles and change the target. Signatures become the event schema and the objective. Modules become the proposer, the folder, and the pruner. The metric stays a metric — temporal cross-validation instead of accuracy on a dev split, permutation importance as the pruner. And what the compiler emits is not a prompt bundle but `spec.yaml`, `model.cbm`, and their hashes. The optimizer can still be an LLM; that is precisely where an LLM earns its place in a compound system — as a design-time component of the compiler, never as a call on the event path.

DSPy tells you to compile the program. This post is asking for one more turn of the screw: compile it into something that cannot call a model.

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

Everything to the left is allowed to be slow, expensive, and messy. Everything to the right is pure code and frozen data. Laid out this way, it is [DSPy](https://dspy.ai/)'s discipline with a different output type: signatures become schemas, modules become the feature proposer, the folder, and the pruner, the metric-driven optimizer becomes the compiler — and what it emits is data instead of instructions. So the compile-time LLM reads the declared event schemas, the tool registry, and the objective, and proposes candidates a syntactic search would never generate:

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

There is also a cold start. Without labeled examples, supervised ML has nothing to learn from, and using a runtime model as a *labeler* — writing into the training set rather than the decision path — is legitimate. For structured tables, prefer [the tabular teacher](#the-better-teacher-tabpfn-not-the-llm) to a chat model, and treat LLM-generated labels as the fallback for text-heavy tables. Either way, those labels are only as good as the model's judgment, and the gap between expressed and actual confidence in language models is [documented](https://arxiv.org/abs/2205.14334) and [still open](https://arxiv.org/abs/2409.00352). Validate against human labels before trusting the artifact in production.

In those cases the flexibility is worth the price. Be honest about the price, though: you are trading speed, cost, determinism, and auditability for it. As with every other "obvious" engineering answer, [there are no solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs).

## What This Changed in My View

Two things moved while I was writing this, and both moved toward the wall.

The first is that the decision-model waypoint arrived faster and better than I expected. I had filed "put a classifier in the loop instead of an LLM" as an obvious idea nobody had shipped. TypeSafe shipped it, published the workflow evals, footnoted their own biases, and priced the calls; Laya shipped open weights and a limitations section that says out loud that the base model is near random on the task it is advertised for. That honesty is what makes the case *for the wall*: the same post that shows Jev on a Pareto frontier also shows the $7/hour Doom line, which is a variable cost doing exactly what variable costs do.

The second is that the argument is really about **artifacts, not models**. I came in thinking the choice was "big model at runtime" versus "small model at compile time." The sharper framing is "a call you make forever" versus "an artifact you own." A compiled spec, a frozen tree ensemble, and four hashes are not a smaller AI system; they are a different kind of thing — something you can diff in a pull request, roll back, replay, audit, and hand to someone else. That is what [conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) looks like for a decision procedure, and it is why the cost argument and the engineering argument keep arriving at the same place.

What did not move: the honest cases at the boundary. Conversational interfaces are not going away, and unlabelable decisions still need something flexible. The wall is a default, not a religion.

## The Point

Compile time is where the LLM lives: it reads your event schemas, your tool registry, and your objective, and proposes candidate features that a syntactic search would never generate. An engine compiles them into point-in-time-correct folds. Traditional ML trains, prunes, and freezes them. Runtime is where neither the LLM nor the decision model appears — the daemon loads the frozen artifacts, verifies their hashes, builds a feature vector, asks the tree model for an action, maps it to a tool, and publishes. That artifact runs anywhere a static binary runs, [ONNX Runtime](https://onnxruntime.ai/) or [its Go binding](https://github.com/yalue/onnxruntime_go) included, with the network unplugged.

Decision models like Jev and Laya are a meaningful improvement over runtime LLMs: no token generation, hundreds of milliseconds instead of seconds, guaranteed type safety, calibrated confidence. They are also still neural networks with per-call pricing, hundreds of megabytes of weights, a GPU to hit their stated latency, and a per-option budget that degrades as your label space grows. A waypoint, not a destination.

---

## A Call to the Community

We are at the beginning of a discipline. AI systems engineering does not have its *Design Patterns* yet. It does not have its *Site Reliability Engineering* yet. It does not have its *Choose Boring Technology* yet.

What it has is a lot of demos in production and a growing sense that something is wrong. The sense is right. The wrongness is economic as much as it is architectural.

Here is what I am asking:

- **Stop putting LLMs in the event loop.** Not because they are bad, but because the event loop is the wrong place for them. The right place is design time, as the semantic layer of your AutoML pipeline.
- **Borrow DSPy's principles, and aim them at compile time.** Declare signatures instead of writing prompts. Compose modules instead of one giant prompt. Let a metric-driven optimizer do the tuning — then make what it emits a feature contract and a frozen model, not an optimized prompt bundle.
- **Start treating prompts as source code.** Version them, review them, hash them, compile them into artifacts. Do not let them run free at runtime, and do not let them become a subscription you cannot cancel.
- **Measure your traditional models, and pick the teacher deliberately.** Do not assume a 400M-parameter model beats a 2 MB tree on your task. Test it; you will be surprised more often than you expect. And when you need soft labels rather than hard ones, distil them from a tabular foundation model before you rent an LLM for the job: the LLM is for meaning, the teacher is for statistics.
- **Build the wall.** Draw the line between compile time and runtime in your architecture, enforce it with hashes, and refuse to run when the wall is broken — or when a variable cost sneaks past.
- **Write down your unit economics.** Know your cost per decision. Know how it scales. Know what it would look like if it were zero. Then ask yourself why it isn't — and know which curve your growth rides: if the cost per decision is positive, every efficiency you buy will be spent back on volume. That is Jevons, not a forecast.
- **Write down what you learn.** This discipline is being invented right now, in production systems, by people too busy shipping to write blog posts. If you have built one of these systems, write it up. The community needs your scars.

The next decade of AI systems will not be won by the teams with the biggest models. It will be won by the teams who understand that intelligence belongs at design time and execution belongs at runtime — and who build the wall between them.

Brooks told us there is no silver bullet. Boehm told us what the delay costs. We have known this for forty years.

Let's apply it.

---

## References

**The economics of where you pay**

- Jevons, W. S. [*The Coal Question: An Inquiry Concerning the Progress of the Nation, and the Probable Exhaustion of Our Coal-Mines*](https://archive.org/details/bub_gb_gAAKAAAAIAAJ) — Macmillan, 1865 (later editions via the [Internet Archive](https://archive.org/details/in.ernet.dli.2015.224624)). The original observation: steam engines became far more efficient per ton of coal, and British coal consumption rose anyway. Jevons called it paradoxical; the modern name is the rebound effect.
- Gillingham, K., Rapson, D., and Wagner, G. [*The Rebound Effect and Energy Efficiency Policy*](https://www.nber.org/papers/w20424) — NBER Working Paper 20424 (2014); published in *Review of Environmental Economics and Policy* 10(1), 2016. The empirical literature on how much of an efficiency gain is spent back as additional demand, and why the answer is elasticities rather than engineering ratios.
- Sharma, P. [*The Jevons Paradox in Cloud Computing: A Thermodynamics Perspective*](https://arxiv.org/abs/2411.11540) — models the cloud as a thermodynamic system and finds that *system growth*, not efficiency, drives consumption; validated against Meta and Google data. The cleanest statement of the mechanism this section applies to inference.
- Luccioni, A. S., Strubell, E., and Crawford, K. [*From Efficiency Gains to Rebound Effects: The Problem of Jevons' Paradox in AI's Polarized Environmental Debate*](https://arxiv.org/abs/2501.16548) — argues that rebound effects "undermine the assumption that improved technical efficiency alone will ensure net reductions", and that incentives and governance decide the trajectory.
- Zhang, Y. and Zhang, T. [*The Economics of Digital Intelligence Capital: Endogenous Depreciation and the Structural Jevons Paradox*](https://arxiv.org/abs/2601.12339) — models LLMs as an asset class and shows falling inference prices pushing firms toward more compute-intensive architectures, making compute demand super-elastic (their phrasing), with the Red Queen effect on rivals' existing capital.
- Narayanan, R. P. and Pace, R. K. [*Will Neural Scaling Laws Activate Jevons' Paradox in AI Labor Markets? A Time-Varying Elasticity of Substitution Analysis*](https://arxiv.org/abs/2503.05816) — formal conditions under which cheaper and better AI translates into market-wide substitution, with elasticity of substitution above one as the threshold.
- Husom, E. J., Nylund, M. E., and Prillard, O. [*Wasted large language models: A life cycle thinking approach*](https://arxiv.org/abs/2608.17055) — efficiency work "has not translated into reduced consumption", and the waste hierarchy's answer is prevention: the cheapest call is the one you do not make.
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

**Compound AI systems, DSPy, and compiled programs**

- Zaharia, M. et al. [*The Shift from Models to Compound AI Systems*](https://bair.berkeley.edu/blog/2024/02/18/compound-ai-systems/) — Berkeley AI Research, February 2024. The framing this post assumes: the unit of production AI is a system of calls, retrievals, and code, not a model.
- Khattab, O. et al. [*DSPy: Compiling Declarative Language Model Calls into Self-Improving Pipelines*](https://arxiv.org/abs/2310.03714) — ICLR 2024, plus the [DSPy documentation](https://dspy.ai/) ([signatures](https://dspy.ai/learn/programming/signatures/), [modules](https://dspy.ai/learn/programming/modules/), [optimizers](https://dspy.ai/learn/optimization/overview/)). Signatures over prompt strings, programs over monoliths, metric-driven compilation over hand-tuning, and a compiled program you save — the four principles this post borrows and then re-aims at an artifact that cannot call a model.
- Opsahl-Ong, K. et al. [*Optimizing Instructions and Demonstrations for Multi-Stage Language Model Programs*](https://arxiv.org/abs/2406.11695) — MIPROv2; Agrawal, L. et al. [*GEPA: Reflective Prompt Evolution Can Outperform Reinforcement Learning*](https://arxiv.org/abs/2507.19457). The state of the art in prompt-side compilation, and the clearest illustration of what it still leaves at runtime.

**Why a small tree beats a big network on a structured decision**

- Shwartz-Ziv, R. and Armon, A. [*Tabular Data: Deep Learning Is Not All You Need*](https://arxiv.org/abs/2106.03253) — XGBoost beats the deep tabular models across 11 datasets.
- Grinsztajn, L., Oyallon, E., and Varoquaux, G. [*Why do tree-based models still outperform deep learning on typical tabular data?*](https://arxiv.org/abs/2207.08815) — NeurIPS 2022; the mechanism is what engineered features exploit.
- Prokhorenkova, L. et al. [*CatBoost: unbiased boosting with categorical features*](https://arxiv.org/abs/1706.09516) — [catboost.ai](https://catboost.ai/) / [model prediction API](https://catboost.ai/docs/en/concepts/python-reference_catboost_predict); [XGBoost model serialization](https://xgboost.readthedocs.io/en/stable/tutorials/saving_model.html); [LightGBM](https://lightgbm.readthedocs.io/en/latest/); [scikit-learn histogram gradient boosting](https://scikit-learn.org/stable/modules/ensemble.html). "Freeze a few megabytes of trees and score in milliseconds on one core" is a supported, boring operation in all of them.
- Lundberg, S. and Lee, S.-I. [*A Unified Approach to Interpreting Model Predictions*](https://arxiv.org/abs/1705.07874) — SHAP; Molnar, C. [*Interpretable Machine Learning*](https://christophm.github.io/interpretable-ml-book/) for what an explanation licenses you to claim.

**Tabular foundation models as teachers (TabPFN and distillation)**

- Tanna, A., Bouarour, N., Bouadi, M., Sankarapu, V., and Seth, P. [*Pocket Foundation Models: Distilling TFMs into CPU-Ready Gradient-Boosted Trees*](https://arxiv.org/pdf/2605.18654) (arXiv:2605.18654, May 2026) — the measurement this section rests on: across 153 classification datasets, distilling TabICLv2 into XGBoost reaches 0.882 macro-mean AUC (96.5% of the teacher's) at 1.9 ms on CPU, 38x–860x faster than the teachers, beating a tuned CatBoost baseline on 51% of datasets (Wilcoxon p=0.0008). The finding that matters most for the recipe is that out-of-fold teacher labelling is mandatory: an in-context teacher scoring rows inside its own context collapses to near-one-hot targets (mean entropy ~10⁻³ nats versus 0.6–0.9 out of context), leaving no inter-class structure to distil. Also: teacher rank transfers to student rank, gains concentrate on tables with ≤21 features (+0.011 AUC over tuned CatBoost, versus +0.001 above), multi-teacher averaging helps MLP students but is negligible for tree students, and distillation hurts where the teacher itself loses. Pipeline open sourced as [TabTune](https://github.com/Lexsi-Labs/TabTune).
- Qu, J., Holzmüller, D., Varoquaux, G., and Le Morvan, M. [*TabICLv2: A better, faster, scalable, and open tabular foundation model*](https://arxiv.org/abs/2602.11139) (2026) — the teacher used in the study above, and the strongest of the four it compared (ahead of TabPFNv2.6 and [LimiX](https://arxiv.org/abs/2509.03505)); the family is plural and open, which is the point.
- Hollmann, N. et al. [*TabPFN: A Transformer That Solves Small Tabular Classification Problems in a Second*](https://arxiv.org/abs/2207.01848) (ICLR 2023) and [*Accurate predictions on small data with a tabular foundation model*](https://doi.org/10.1038/s41586-024-08328-6) (Nature, 2025) — the prior-data fitted network: trained once on synthetic tables from a causal-model prior, then used in-context with no hyperparameters. TabPFN-2.5's report adds the production path this section leans on: ["a new distillation engine that converts TabPFN-2.5 into a compact MLP or tree ensemble"](https://arxiv.org/abs/2511.08667), plus the 100% win rate against default XGBoost on small-to-medium classification. [TabPFN-3](https://arxiv.org/abs/2605.13986) scales context to 1M rows and beats 8-hour-tuned GBDT baselines; [TabPFN-3.5](https://arxiv.org/abs/2609.17895) adds text-rich, grouped, and temporal tables.
- [Prior Labs model table](https://docs.priorlabs.ai/models) and [repository](https://github.com/PriorLabs/TabPFN) — the per-release row/column caps (50k × 2k for 2.5, 100k × 2k for 2.6, 1M × 2k for 3, 1M × 20k for 3.5) and the licence split: TabPFN-2 weights are Apache-2.0 with attribution, while 2.5, 2.6, 3, and 3.5 weights are non-commercial. Checkpoint sizes on [Hugging Face](https://huggingface.co/Prior-Labs/TabPFN-v2-clf) (12.9–29 MB for v2, ~213 MB for TabPFN-3, ~876 MB for TabPFN-3.5) are what makes the teacher a laptop step rather than a cluster step.
- Hinton, G., Vinyals, O., and Dean, J. [*Distilling the Knowledge in a Neural Network*](https://arxiv.org/abs/1503.02531) — why soft targets beat hard labels: the teacher's distribution encodes the relative plausibility of wrong answers, which is information a 0/1 label destroys.
- Jayawardhana, M., Renbo, Dooley, S. [*Transformers Boost the Performance of Decision Trees on Tabular Data across Sample Sizes*](https://arxiv.org/abs/2502.02672) — PFN-Boost and LLM-Boost: fused transformer-plus-GBDT models beat both standalone components on average, with [code released](https://github.com/MayukaJ/LLM-Boost). The closest published version of the recipe in this section.
- Dissanayake, P. and Dutta, S. [*TabDistill: Distilling Transformers into Neural Nets for Few-Shot Tabular Classification*](https://arxiv.org/abs/2511.05704) — distillation from tabular transformers into smaller students, with neural students; the GBDT student is the variant this post argues for.
- Ma, J., Thomas, V., Yu, G., and Caterini, A. [*In-Context Data Distillation with TabPFN*](https://arxiv.org/abs/2402.06971) — the other reading of "distillation" around TabPFN: shrinking the *context* to fit the teacher's limits. Read it as the caveat: above the row cap you are distilling a subsample, so the context is part of the compile step.
- Treerath, W. and Pittorino, F. [*Active In-Context Learning for Tabular Foundation Models*](https://arxiv.org/abs/2603.27385) — states the calibration property this section relies on ("TabPFN provide[s] calibrated probabilistic predictions via in-context learning") and benchmarks against CatBoost and XGBoost baselines; Ramalingam, M. [*Uncertainty-Aware Tabular Prediction: Evaluating VBLL-Enhanced TabPFN*](https://arxiv.org/abs/2509.10048) finds plain TabPFN better calibrated than a variational variant — useful when deciding whether the soft labels need recalibration.
- CatBoost. [*Classification: objectives and metrics*](https://catboost.ai/docs/en/concepts/loss-functions-classification) — `CrossEntropy` as an optimisation objective, which is what lets a tree ensemble be trained against probabilities instead of labels.
- Cai, P., Gao, Z., and Lian, W. [*Strengthening LLMs for Tabular Prediction with Structural Priors*](https://arxiv.org/abs/2510.17385); Hegselmann, S. et al. [*TabLLM: Few-shot Classification of Tabular Data with Large Language Models*](https://arxiv.org/abs/2210.10723); Liu, T., Wang, F., and Chen, M. [*Rethinking Tabular Data Understanding with Large Language Models*](https://arxiv.org/abs/2312.16702); Golkar, S. et al. [*xVal: A Continuous Numerical Tokenization for Scientific Language Models*](https://arxiv.org/abs/2310.02989) — the LLM-teacher side: serialisation, column-order sensitivity, discontinuous numeral tokenisation, and how much structural help an LLM needs to be competitive at all.

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

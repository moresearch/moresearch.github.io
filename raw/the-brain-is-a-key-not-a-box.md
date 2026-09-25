---
title: "The Brain Is a Key, Not a Box"
date: 2026-09-23
slug: the-brain-is-a-key-not-a-box
summary: "The storage metaphor of memory — encode, compress, file, retrieve, unzip — fails for brains and for neural networks in the same three ways: nothing is stable, nothing is located, and nothing survives being switched off. Victoria Trumbull's argument that the brain enables memory rather than containing it is a description of parametric memory in machine learning; her two readings of 'time stores memory' map onto recurrence and attention; and CatBoost's ordered target statistics — which the paper calls an 'artificial time' — are the clearest small example of a model that keeps the past in an ordering instead of compressing it into a value."
tags: memory, philosophy-of-mind, neural-networks, parametric-memory, engram, reconsolidation, bergson, duration, eternalism, enactivism, catboost, xgboost, gradient-boosting, ordered-boosting, rag, context-window, hallucination, essay
---

There is a metaphor we all carry around and almost never examine. It goes like this:

```text
experience ──▶ encode ──▶ compress ──▶ store in a folder ──▶ retrieve ──▶ unzip
             (hippocampus)            (a "memory trace")        (recall)
```

You experience something, your brain encodes it, compresses it like a zip file, files it in a folder — hippocampus, cortex, somewhere — and later you open it and watch it play. Nothing about that feels like a philosophical position. It feels like the obvious description of remembering.

It is a very specific philosophical position with a name: the **storage metaphor** of memory, encoding a kind of reductive materialism in which the memory *is* a physical thing located *inside* the brain — what neuroscience calls an **engram**. The metaphor is older than neuroscience (Plato's wax block, Aristotle's seal in wax, then traces, vibrations, switches), and for most of that history it was not a prediction about biology so much as the only vocabulary anyone had ([Danziger](https://archive.org/details/markingmindhisto0000danz); [SEP: Memory](https://plato.stanford.edu/entries/memory/)).

The philosopher Victoria Trumbull argues it fails, and — more interestingly — that it has misdescribed the *shape* of the answer: the brain does not *contain* memories, it *enables* them, the way a radio enables music without containing the music ([*Memories are not stored in the brain*](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull)).

[![Victoria Trumbull — "Memories are not stored in the brain" (IAI, 2026)](https://iai.tv/assets/Uploads/_resampled/FillWyI4MDAiLCI1MDAiXQ/memories-are-not-stored-in-the-brain.webp)](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull)

What follows takes that seriously without either swallowing it or dismissing it — and then asks the question this blog always ends up asking, because we built machines out of the same metaphor and it is breaking in the same three places.

## Three ways the zip file metaphor breaks

**1. Recollection rewrites the file.** A zip file does not change when you unzip it. A memory does. The mechanism is called **reconsolidation**, and it was pinned down in a famous experiment: a consolidated fear memory in rats, once retrieved, becomes temporarily unstable and needs new protein synthesis to restabilize; block the synthesis and it does not come back intact ([Nader, Schafe & LeDoux](https://doi.org/10.1038/35021052), *Nature*, 2000). Retrieval is not read-only. It is a rewrite followed by a save.

That is the physiological version of what psychology had documented for decades: Bartlett's subjects reshaped stories toward their own expectations ([1932](https://archive.org/details/rememberingstudy00bart)); Loftus and Palmer's subjects "remembered" broken glass that was never there after being asked about cars that *smashed* into each other ([1974](https://doi.org/10.1016/S0022-5371(74)80011-3)); and constructive-memory research now treats remembering and imagining as overlapping operations on the same machinery ([Schacter & Addis](https://doi.org/10.1098/rstb.2007.2087)). A storage system whose retrieval corrupts the store is not a filing cabinet. It is closer to a wiki that rewrites itself every time someone reads a page.

**2. There is no folder.** Sixty years of looking for the location of a specific memory did not produce a box. Engrams are real — cells have been identified, tagged, artificially activated to trigger a behavior, and silenced to prevent one ([Josselyn & Tonegawa](https://doi.org/10.1126/science.aaw4325), *Science*, 2020) — but the trace is not addressable the way the metaphor needs it to be. One memory is represented by ensembles *distributed across multiple brain regions*, each contributing a piece, with the same cells participating in overlapping memories ([Roy et al.](https://doi.org/10.1038/s41467-022-29384-4), 2022). Karl Lashley spent decades cutting into rat brains looking for the trace, found that removing more tissue degraded performance smoothly instead of deleting memories one at a time, and titled the write-up *In search of the engram* ([reviewed in Josselyn, Köhler & Frankland](https://doi.org/10.1038/nrn4000), 2015).

**3. The store is never powered off.** A hard drive can sit on a shelf for a decade and return the same bits. A brain spends roughly 20% of the body's energy budget ([Raichle & Gusnard](https://doi.org/10.1073/pnas.172399499), 2002), most of it on activity unrelated to the task in front of it — the spontaneously active, task-negative network called the default mode ([Raichle et al.](https://doi.org/10.1073/pnas.98.2.676), 2001). Nothing is parked. When activity stops, the memory does not wait on the shelf; the system that could reconstitute it is gone.

So we are describing an object with none of the properties of a file: it changes on read, has no location, and exists only while running. A radio does not contain Beethoven's Seventh; it is a structure that, powered and tuned, produces it. Smash the radio and you have not destroyed the symphony — you have destroyed one way of hearing it.

## What "time itself stores memory" could mean

The radical half of the argument is that memory is not stored *anywhere*, because the thing doing the storing is **time**. Two traditions get compressed into that sentence, and they are worth separating.

**Reading A: Bergsonian duration.** In *Matter and Memory* (1896), the brain is not an archive but "an organ of attention to life" — a filter that narrows the whole of the past down to what is useful now ([Bergson](https://archive.org/details/mattermemor00berg); [SEP: Bergson](https://plato.stanford.edu/entries/bergson/)). The past is not gone, it is simply not useful at the moment. Memory is *duration*: a living accumulation that persists rather than a shelf that survives. Remembering is less like retrieving a document than like turning your attention to a region of your own temporal extent.

**Reading B: the block universe.** Under eternalism every moment of spacetime exists equally; 1900 and 2026 are both real at different coordinates, the way New York and Paris are both real at different places. That is one live option among several in the philosophy of time ([SEP: Time](https://plato.stanford.edu/entries/time/); Sider's [four-dimensionalism](https://global.oup.com/academic/product/four-dimensionalism-9780199263523)), with presentists denying the whole picture. On this reading a memory needs no storage location: the event is still at its coordinate, and the brain's job is to be a structure complex enough to re-establish a relation to it.

Both readings land on the same inversion: the brain is an *enabling condition*, not a *storage location* — and we mistook one for the other because the brain is the thing we can watch working while remembering happens.

## Where the argument is harder than it looks

Two steps are weaker than the rhetoric suggests, and both matter for what follows.

**The physics does not deliver the conclusion by itself.** Eternalism is an interpretation, not a result; relativity removes a privileged global "now" while the debate between presentism, the growing block, and the moving spotlight continues. Rovelli's *The Order of Time* goes further and then lands somewhere that cuts *against* the simple "the past is just elsewhere" picture: the arrow of time, he argues, is something we recover from the *records* the world leaves behind ([Rovelli](https://www.penguinrandomhouse.com/books/557814/the-order-of-time-by-carlo-rovelli/)). If traces give time its arrow, then "time stores memory" quietly becomes "memory is what makes time storable" — stranger, and more interesting.

**Even granting the block universe, you still need a key.** If the event at t₁ exists, what makes a brain at t₂ *about* that event rather than any other coordinate? Something has to connect them — a causal chain, a structural correspondence, a disposition to respond to the right thing. Whatever that is, it is a trace, and traces are what the causal theory of memory is built on ([SEP: Memory](https://plato.stanford.edu/entries/memory/)). So the storage metaphor is better called *misplaced* than false: the trace is not the repository, it is the hook that lets a present process re-instantiate a past one. That is the structure the extended and enactive traditions build on ([Clark & Chalmers](https://www.consc.net/papers/extended.html); [Varela, Thompson & Rosch](https://archive.org/details/embodiedmindcogn0000vare); [SEP: Embodied Cognition](https://plato.stanford.edu/entries/embodied-cognition/)).

The empirical bridge is worth one line: divers who learned word lists underwater recalled them better underwater ([Godden & Baddeley](https://doi.org/10.1111/j.2044-8295.1975.tb01468.x), 1975). The room is part of the key, which the box model cannot represent. And a caution: Whitehead's *Process and Reality* is a metaphysics in which the past is causally efficacious in the present, not a claim about neurobiology. Reading it as physics is how a good idea becomes a bad one.

## The same metaphor, in the machine

The storage metaphor was also the founding assumption of AI, and here the philosophy stops being a spectator sport.

Classical symbolic AI was a database. Knowledge sat in rows — `Paris = capital_of(France)` — in semantic networks and ontologies, and reasoning was retrieval plus rules. Connectionism replaced that with something stranger. A neural network has no rows; it has a pattern of connection strengths, and a "fact" exists only as a disposition distributed across millions of weights, re-instantiating when the right input arrives. There is no folder in GPT for "Paris" — no address you can point to and say *here it lives*.

This is standard vocabulary rather than a poetic reading: we call it **parametric memory**, knowledge held in weights, as opposed to **non-parametric memory**, knowledge held where you can look it up ([Petroni et al.](https://arxiv.org/abs/1909.01066), 2019). Which means Trumbull's sentence — *the system enables memory without containing it* — is a description of a trained network, not an analogy for one. The AI form of reductive materialism says intelligence is *in the weights*; the everyday experience of a neural network is that nothing is *in* anything.

The evidence that there is no folder is that we keep failing to edit one. If a fact were a row, changing it would be an `UPDATE` statement; instead, locating and editing a single factual association means finding the mid-layer activations that carry it, and the edit remains entangled with everything else those weights do ([Meng et al.](https://arxiv.org/abs/2202.05262), ROME, 2022). Distributed is not a slogan here; it is a budget line.

The failure mode matches too. A generative model does not retrieve a stored sentence — it reconstructs a continuation, and when the reconstruction has no grounding the result is fluent and wrong, structurally rather than accidentally ([Xu & Jain](https://arxiv.org/abs/2401.11817); for the human taxonomy, [Ji et al.](https://arxiv.org/abs/2202.03629)). Confabulation is what reconstruction looks like when the key turns in the wrong lock.

The analogy should not be oversold: human confabulation and LLM hallucination share a *shape* — construction without verification — but not a mechanism, one running through reconsolidation in an organism with stakes and the other through next-token distributions over a corpus. The useful claim is structural. **Both are keys, and neither is a box.**

## The two readings, in code

Machine learning has spent thirty years building one architecture for each of Trumbull's two readings, and arguing about which is right.

**Recurrence: memory lives in the process, not the weights.** A recurrent network carries a hidden state forward — the past exists as *how it has changed the present state*:

```python
h_t = f(h_{t-1}, x_t)   # the past is not retrieved; it persists
```

That is duration in code: nothing is looked up, the past is present as accumulated modification. The disanalogy is the second half of the story. A hidden state is a fixed-width *compression* of the past, not a retention of it; Bergson's duration keeps everything and prioritizes by attention, while a vanilla RNN forgets and its gradient signal decays over long horizons ([Bengio et al.](https://doi.org/10.1109/72.279181), 1994). Gates, and now selective state-space models like [Mamba](https://arxiv.org/abs/2312.00752), are negotiations with exactly that limit.

**Attention: leave the past where it is and learn to point at it.** Transformers hold the whole sequence and let every position look at every other, recomputing each token's representation from the entire context on every forward pass. Nothing is zipped, because nothing is unzipped. Two corrections to the popular telling: attention is content-addressed rather than coordinate-addressed (positions enter as added information, not as the mechanism), and the KV cache is an artifact of a single generation, recomputed next time and present only while the process runs. The past there is *held*, not *kept*.

The engineering trajectory is the interesting part. We spent years trying to store everything inside the model; the last several years have gone the other way, and the results keep favouring it — [RAG](https://arxiv.org/abs/2005.11401), [nearest-neighbour LMs](https://arxiv.org/abs/1911.00172) (whose title, *Generalization through Memorization*, is the irony of the decade), [RETRO](https://arxiv.org/abs/2112.04426) at trillion-token scale, [Memorizing Transformers](https://arxiv.org/abs/2203.08913), and now the entire industry of vector databases and long context. That is Trumbull's move, executed by engineers who have never read Bergson: do not force the model to contain the past — leave it out in the world and make the model a structure that can establish a relation to it.

With the same caveat a brain has: having the past laid out is not the same as using it. Retrieval quality degrades sharply depending on where the relevant passage sits in a long window, worst in the middle ([Liu et al.](https://arxiv.org/abs/2307.03172), 2023). The key has to fit the lock. Storage was never the hard part.

## Case study: the zip file in gradient boosting

The cleanest small example is gradient-boosted trees, and specifically the difference between XGBoost and CatBoost — usually explained as a feature-engineering advantage, actually a disagreement about where memory should live.

Start with the naive move, which is the zip file as a line of code. For a categorical column like `city`, the shortcut is mean encoding:

```python
# the zip-file move: compress all of history into one stored number
df["city_encoded"] = df.groupby("city")["target"].transform("mean")
```

Every Paris row now carries one number computed from *every* Paris row — including rows that come after it, and including its own label. The past has been compressed into a constant stored in a cell, and the future has leaked into it. That is the storage metaphor as a bug: a value that *is* the memory, true regardless of when you read it.

CatBoost's answer is the philosophical one, and the paper says it almost openly. Instead of computing the encoding once and storing it, it computes, for each example, a statistic from the examples that come *before* it in an ordering — and to make that possible in a static dataset, the authors introduce what they call an **artificial "time"**:

> "Clearly, the values of TS for each example rely only on the observed history. To adapt this idea to standard offline setting, we introduce an artificial 'time', i.e., a random permutation σ of the training examples." — [Prokhorenkova et al., *CatBoost*](https://arxiv.org/abs/1706.09516), 2018

To avoid storing a memory as a value, give the data a *time* — and the encoding becomes a relation to a history instead of a constant. Extended to gradients, the same mechanism is **ordered boosting**, which prevents *prediction shift*: the leakage you get when the same rows both fit the model and estimate what it should predict.

Now the correction that keeps the analogy honest, because it is tempting to read "ordered" as "temporal" and stop. The paper's own word is *artificial*: by default that ordering is a **random permutation**, several of them across boosting steps. Your event stream's chronology is not what the library is respecting, which is why CatBoost exposed a `has_time` switch documented as: *"Use the order of objects in the input data (do not perform a random permutation of the dataset at the preprocessing stage)… Default: FALSE (not used; permute input dataset)"* ([R package source](https://github.com/catboost/catboost/blob/master/catboost/R-package/R/catboost.R)).

So `ordered` means ordered for *anti-leakage* purposes, not chronological. Those are different guarantees, and conflating them is the mistake I warned about [last post](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern): label leakage is prevented structurally by the library, while temporal leakage is your responsibility, prevented by time-ordered folds and point-in-time-correct features. A random permutation is a beautiful trick against one and a quiet hazard for the other.

With that caveat, the reading holds. CatBoost also builds **oblivious trees** — the same split at every level, which makes them fast to evaluate and, in the paper's phrase, "less prone to overfitting." You cannot open a trained model and find row #4521 in it: the training set's influence is spread across thousands of splits and leaf values, enabling prediction without containing data. XGBoost deserves no strawman here — it is equally distributed, and has handled categoricals natively since 1.5 with partition-based splits ([docs](https://xgboost.readthedocs.io/en/stable/tutorials/categorical.html)). Historically it asked you to bring your own encoding, which meant carrying the leakage risk yourself. That is a difference in *where the history lives during training*, not between a database and a brain.

The honest version: **both are keys.** What differs is whether the history is precomputed into a column and treated as timeless — the zip file — or recomputed per example from an ordering that defines what counts as the past. CatBoost's contribution to this essay is that it is ordinary, popular software whose designers independently decided the past should live in an order rather than a value.

## What this means for how we build

A warning first: machines really do store. A file on disk is exactly the box model, and for silicon the storage metaphor is accurate rather than metaphorical. The takeaway is narrower — **do not assume a brain is doing what your database does**, and be suspicious when a cognitive claim arrives pre-shaped like a system you already know how to build.

The suspicion runs both ways, because "memory" in agent harnesses is almost always a storage layer: embeddings as engrams, a vector database as the hippocampus, retrieval as unzip-and-read. If the philosophy is right, that imports the failure modes:

- **Retrieved memories get treated as facts.** A file is authoritative; a reconstruction is a hypothesis. If recall is construction, then a system that learns from its own recalled state is compounding its own edits — the dynamic behind drift and confidently retrieved falsehood.
- **The room gets dropped.** Context-dependence means the situation is part of the memory; retrieval keyed on text similarity alone discards state, task, provenance, and what the system was doing when it learned the thing.
- **Architecture assumes a powered-off shelf.** A brain never has one. A log you can re-derive from is closer to that than a store you load from — [the event log is the honest memory](https://blog.hackspree.com/#stories-from-events), [narratives are reconstructions built from it](https://blog.hackspree.com/#events-as-the-source-of-truth), [a durable daemon survives the power cycle](https://blog.hackspree.com/#durable-daemons).

The economics follow. If the storage metaphor is wrong, then its plan — make the container bigger so it can hold more — is a plan for buying hard drives. The evidence supports a smaller structure with better access: less memorization, more reconstruction skill, memory as a *relation between system and world* rather than content inside the system. That is the same thesis this blog keeps arguing from the model side, where [specialization, not size, is the lever](https://blog.hackspree.com/#specialized-small-language-models) and the deployable unit is [the model plus the harness](https://blog.hackspree.com/#better-harnesses-smaller-models).

And the honest limit: none of this is a design specification. Philosophy can tell you the storage metaphor is a bad map of biological memory and a leaky one for models. It cannot tell you how to build better agent memory, and "the past still exists at its coordinate" is not a caching strategy. What it can do is change your default assumption and the failure mode you watch for — memories confidently retrieved, unaccountably shaped by context, and quietly rewritten every time you read them.

## What this changed in my view

Three things moved while I was working through this, and one didn't.

**I had assumed the engram debate was settled empirically, in the storage metaphor's favour.** Find the trace, tag it, activate it — Tonegawa-style experiments really do this, and I had filed it as "the box exists, the philosophers are quibbling about vocabulary." The distributed-engram work changed that: the trace exists and is *not* a location. That is a result the storage metaphor cannot represent, which is stronger than a philosophical preference.

**I had treated "memory is reconstruction" as a slogan about false memory.** Reconsolidation makes it a mechanism — recall destabilizes, edits, re-saves — and a mechanism is a warning about any system that reads its own memory and writes back, including the agents we are building.

**The best example turned out to be small and ordinary.** I expected a large language model, where the story is diffuse and hard to verify. Instead it was seventy lines of gradient boosting: CatBoost's paper introduces an *artificial time* to avoid storing a memory as a value, and hands you a switch to say "no, use the real order." I nearly wrote the tidy version of that — that CatBoost respects time. It does not; it respects *an* order, usually a random one, and `ordered` is about leakage rather than chronology. The correction improved the essay, because it separates two things the storage metaphor also conflates: keeping the past *around* and keeping it *in order* are different achievements, and only one of them is free.

What did not move: brains are physical, engrams are real, weights are real, and the storage metaphor is not worthless. It is just load-bearing in the wrong place — which is the most consequential kind of wrong a metaphor can be.

## References

**The primary argument**

- Trumbull, V. [*Memories are not stored in the brain*](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull) — Institute of Art and Ideas talk, 2026. The claim under discussion: the brain enables rather than contains, and the storage metaphor mistakes an enabling condition for a storage location. Figure above links to the talk.
- Michaelian, K. and Sutton, J. [*Memory*](https://plato.stanford.edu/entries/memory/) — Stanford Encyclopedia of Philosophy. Traces, the causal theory of memory, distributed and extended accounts, and the objections to each.
- Semon, R. [*The Mneme*](https://archive.org/details/cu31924100387210) (1921; German original 1904) — the source of the term *engram*, and of *ecphory*: Semon already distinguished the trace from the act of reactivating it, which is a reminder that the box was never the whole theory.
- Danziger, K. [*Marking the Mind: A History of Memory*](https://archive.org/details/markingmindhisto0000danz) — Cambridge University Press, 2008. How the storage vocabulary — wax, seals, traces, vibrations, switches, files — was carried across four centuries of theories.

**The three failures**

- Nader, K., Schafe, G. and LeDoux, J. [*Fear memories require protein synthesis in the amygdala for reconsolidation after retrieval*](https://doi.org/10.1038/35021052) — *Nature* 406, 2000. Retrieval destabilizes and re-stores.
- Bartlett, F. C. [*Remembering: A Study in Experimental and Social Psychology*](https://archive.org/details/rememberingstudy00bart) — Cambridge, 1932; Loftus, E. and Palmer, J. [*Reconstruction of automobile destruction*](https://doi.org/10.1016/S0022-5371(74)80011-3) — 1974; Schacter, D. and Addis, D. R. [*The cognitive neuroscience of constructive memory*](https://doi.org/10.1098/rstb.2007.2087) — 2007. Reconstruction, from psychology to neuroscience.
- Josselyn, S., Köhler, S. and Frankland, P. [*Finding the engram*](https://doi.org/10.1038/nrn4000) — 2015; Josselyn, S. and Tonegawa, S. [*Memory engrams: Recalling the past and imagining the future*](https://doi.org/10.1126/science.aaw4325) — 2020; Roy, D. et al. [*Brain-wide mapping reveals that engrams for a single memory are distributed across multiple brain regions*](https://doi.org/10.1038/s41467-022-29384-4) — 2022. Lashley's failure, the strongest case for engrams, and the distribution result.
- Raichle, M. and Gusnard, D. [*Appraising the brain's energy budget*](https://doi.org/10.1073/pnas.172399499) — 2002; Raichle, M. et al. [*A default mode of brain function*](https://doi.org/10.1073/pnas.98.2.676) — 2001. The never-off property.
- Godden, D. and Baddeley, A. [*Context-dependent memory in two natural environments*](https://doi.org/10.1111/j.2044-8295.1975.tb01468.x) — 1975. The environment as part of the retrieval key.

**Time, duration, and the block universe**

- Bergson, H. [*Matter and Memory*](https://archive.org/details/mattermemor00berg) (1896) — the brain as an organ of attention to action, and duration as the persistence of the past; [SEP: Bergson](https://plato.stanford.edu/entries/bergson/).
- Markosian, N. et al. [*Time*](https://plato.stanford.edu/entries/time/) — SEP, including eternalism, presentism, and the Rietdijk–Putnam argument; [*Presentism*](https://plato.stanford.edu/entries/presentism/); [*Being and Becoming in Modern Physics*](https://plato.stanford.edu/entries/spacetime-bebecome/); Sider, T. [*Four-Dimensionalism*](https://global.oup.com/academic/product/four-dimensionalism-9780199263523) — OUP, 2001, with the [SEP entry on temporal parts](https://plato.stanford.edu/entries/temporal-parts/).
- Canales, J. [*The Physicist and the Philosopher*](https://www.degruyterbrill.com/document/doi/10.1515/9781400873310/html) — Princeton, 2015. The 1922 Einstein–Bergson debate.
- Rovelli, C. [*The Order of Time*](https://www.penguinrandomhouse.com/books/557814/the-order-of-time-by-carlo-rovelli/) — Riverhead, 2018. The arrow of time as something recovered from records.

**Parametric and non-parametric memory in machines**

- Petroni, F. et al. [*Language Models as Knowledge Bases?*](https://arxiv.org/abs/1909.01066) — 2019. Parametric knowledge, stated cleanly.
- Meng, K. et al. [*Locating and Editing Factual Associations in GPT*](https://arxiv.org/abs/2202.05262) — ROME, 2022. Why there is no row to update.
- Xu, Z. and Jain, S. [*Hallucination is Inevitable*](https://arxiv.org/abs/2401.11817) — 2024; [Ji et al.](https://arxiv.org/abs/2202.03629) on the human-side taxonomy.
- Khandelwal, U. et al. [*Generalization through Memorization*](https://arxiv.org/abs/1911.00172) — 2019; Lewis, P. et al. [*Retrieval-Augmented Generation*](https://arxiv.org/abs/2005.11401) — 2020; Borgeaud, S. et al. [*RETRO*](https://arxiv.org/abs/2112.04426) — 2021; Wu, Y. et al. [*Memorizing Transformers*](https://arxiv.org/abs/2203.08913) — 2022. The turn from storing the past in weights to leaving it outside and pointing at it.
- Liu, N. et al. [*Lost in the Middle*](https://arxiv.org/abs/2307.03172) — 2023. Laid out is not the same as used.
- Bengio, Y. et al. [*Learning long-term dependencies with gradient descent is difficult*](https://doi.org/10.1109/72.279181) — 1994; Gu, A. and Dao, T. [*Mamba*](https://arxiv.org/abs/2312.00752) — 2023. Why a fixed-width state cannot retain everything.

**Where the past lives in gradient boosting**

- Prokhorenkova, L. et al. [*CatBoost: unbiased boosting with categorical features*](https://arxiv.org/abs/1706.09516) — NeurIPS 2018. The *artificial time* quote, ordered target statistics, prediction shift, oblivious trees. [Docs](https://catboost.ai/docs/en/concepts/algorithm-main-stages).
- [CatBoost R package source](https://github.com/catboost/catboost/blob/master/catboost/R-package/R/catboost.R) — `has_time`: use the order of objects in the input data, do not permute. The switch separating *an* order from *your* order.
- [XGBoost: Categorical Data](https://xgboost.readthedocs.io/en/stable/tutorials/categorical.html) — native categorical support since 1.5, partition-based splits.
- For the divergence between anti-leakage ordering and temporal correctness: [*LLM at Runtime Is an Antipattern*](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) on point-in-time folds, and Feast's [point-in-time joins](https://docs.feast.dev/getting-started/concepts/point-in-time-joins).

**Minds that are not sealed containers**

- Clark, A. and Chalmers, D. [*The Extended Mind*](https://www.consc.net/papers/extended.html) — 1998; Varela, F., Thompson, E. and Rosch, E. [*The Embodied Mind*](https://archive.org/details/embodiedmindcogn0000vare) — MIT Press, 1991; [SEP: Embodied Cognition](https://plato.stanford.edu/entries/embodied-cognition/).
- Whitehead, A. N. [*Process and Reality*](https://archive.org/details/processrealityes0000whit) (1929) — the past as causally efficacious in the present; James, W. [*Human Immortality*](https://archive.org/details/humanimmortality00jame) (1898) — the transmission theory, the radio analogy argued a century early.

**This blog, on memory and logs**

- [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Durable Daemons](https://blog.hackspree.com/#durable-daemons), [SSLM: Specialized Small Language Models](https://blog.hackspree.com/#specialized-small-language-models), [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern).

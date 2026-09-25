---
title: "The Brain Is a Key, Not a Box"
date: 2026-09-23
slug: the-brain-is-a-key-not-a-box
summary: "The storage metaphor of memory — encode, compress, file, retrieve, unzip — fails for brains and it fails for neural networks in the same three ways: nothing is stable, nothing is located, and nothing survives being switched off. Victoria Trumbull's argument that the brain enables memory rather than containing it is a description of parametric memory in machine learning; her two readings of 'time stores memory' map onto recurrence and attention; and CatBoost's ordered target statistics — which the paper calls an 'artificial time' — are the clearest small example of a model that keeps the past in an ordering instead of compressing it into a value."
tags: memory, philosophy-of-mind, neural-networks, parametric-memory, engram, reconsolidation, bergson, duration, eternalism, enactivism, catboost, xgboost, gradient-boosting, ordered-boosting, rag, context-window, hallucination, essay
---

There is a metaphor we all carry around and almost never examine. It goes like this:

```text
experience ──▶ encode ──▶ compress ──▶ store in a folder ──▶ retrieve ──▶ unzip
             (hippocampus)            (a "memory trace")        (recall)
```

You experience something. Your brain encodes it, compresses it like a zip file, files it in a folder — hippocampus, cortex, somewhere — and later you open it, decompress it, and watch it play. Nothing about that feels like a philosophical position. It feels like the obvious description of what remembering is.

It is a very specific philosophical position, with a name. It is the **storage metaphor** of memory, and the theory it encodes is a kind of reductive materialism: the memory *is* a physical thing, located *inside* the brain, which neuroscientists call an **engram**. The metaphor is older than neuroscience — Plato's wax block in the *Theaetetus*, Aristotle's seal pressed into wax in *De Memoria*, and then the modern line of traces, traces-as-vibrations, traces-as-switches — and for most of that history it was not a prediction about biology so much as the only vocabulary anyone had ([Danziger](https://archive.org/details/markingmindhisto0000danz); [SEP: Memory](https://plato.stanford.edu/entries/memory/)).

The philosopher Victoria Trumbull argues the metaphor fails, and — more interestingly — that it has misdescribed the *shape* of the answer. Her talk is bluntly titled [*Memories are not stored in the brain*](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull), and her conclusion is that the brain does not *contain* memories; it *enables* them — the way a radio enables music without containing the music, or legs enable walking without containing "walks" inside them.

[![Victoria Trumbull — "Memories are not stored in the brain" (IAI, 2026)](https://iai.tv/assets/Uploads/_resampled/FillWyI4MDAiLCI1MDAiXQ/memories-are-not-stored-in-the-brain.webp)](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull)

What follows takes that claim seriously without either swallowing it or dismissing it. And then it does the thing this blog always ends up doing: it asks what happens when you build machines out of the same metaphor — because we did. We imported the storage model directly into AI, and it is breaking in exactly the same three places.

## Three ways the zip file metaphor breaks

**1. Recollection rewrites the file.** A zip file does not change when you unzip it. A memory does. The mechanism has a name — **reconsolidation** — and it was pinned down in a now-famous experiment: a consolidated fear memory in rats, when retrieved, becomes temporarily unstable and requires new protein synthesis to restabilize; block the synthesis and the memory does not come back intact ([Nader, Schafe & LeDoux](https://doi.org/10.1038/35021052), *Nature*, 2000). Retrieval is not read-only. It is a rewrite followed by a save.

That is the physiological version of what psychologists had been documenting for decades: Bartlett's subjects reshaped stories toward their own expectations ([*Remembering*, 1932](https://archive.org/details/rememberingstudy00bart)); Loftus and Palmer's subjects, asked how fast the cars were going when they *smashed* into each other, "remembered" broken glass that was never there ([Loftus & Palmer, 1974](https://doi.org/10.1016/S0022-5371(74)80011-3)); and the constructive-memory literature now treats remembering and imagining as overlapping operations on the same machinery ([Schacter & Addis](https://doi.org/10.1098/rstb.2007.2087), 2007). A storage system whose retrieval corrupts the store is not a filing cabinet. It is closer to a wiki that rewrites itself every time someone reads a page.

**2. There is no folder.** Sixty years of looking for the location of a specific memory did not produce a box. It produced something stranger. Engrams are real — engram cells have been identified, tagged, artificially activated to trigger a behavior, and silenced to prevent one ([Josselyn & Tonegawa](https://doi.org/10.1126/science.aaw4325), *Science*, 2020) — but the trace is not addressable the way the metaphor needs it to be. A single memory is represented by cell ensembles *distributed across multiple brain regions*, each contributing a piece ([Roy et al.](https://doi.org/10.1038/s41467-022-29384-4), *Nature Communications*, 2022), and the same cells participate in overlapping memories. The history is telling: Karl Lashley spent decades cutting into rat brains looking for the trace, found that removing more tissue degraded performance more or less smoothly instead of deleting memories one at a time, and summarized the failure in a paper whose title is the metaphor's tombstone — *In search of the engram* ([reviewed in Josselyn, Köhler & Frankland](https://doi.org/10.1038/nrn4000), 2015).

**3. The store is never powered off.** A hard drive can sit on a shelf for a decade with no electricity and return the same bits. A brain cannot: it spends roughly 20% of the body's energy budget ([Raichle & Gusnard](https://doi.org/10.1073/pnas.172399499), *PNAS*, 2002), most of it on activity that has nothing to do with the task in front of it — the spontaneously active, task-negative network now standardly called the default mode ([Raichle et al.](https://doi.org/10.1073/pnas.98.2.676), *PNAS*, 2001). Nothing is parked. When activity stops, the memory does not wait on the shelf; the system that could reconstitute it is gone.

Put together, the three failures describe an object with none of the properties of a file: it changes on read, has no location, and exists only while running. Which is why the radio analogy is not as silly as it sounds. A radio does not contain Beethoven's Seventh. It is a structure that, when powered and tuned, produces it. Smash the radio and you have not destroyed the symphony; you have destroyed one way of hearing it.

## What "time itself stores memory" could mean

The radical half of Trumbull's argument is that memory is not stored *anywhere*, because the thing doing the storing is **time**. That sounds mystical until you separate the two traditions being compressed into one sentence.

**Reading A: the Bergsonian one — the past persists in duration.** Bergson's *Matter and Memory* (1896) makes a distinction that ordinary psychology keeps collapsing: the brain is not an archive but an organ of *action* — "an organ of attention to life" — whose job is to filter the whole of the past down to what is useful for the present ([Bergson](https://archive.org/details/mattermemor00berg); [SEP: Bergson](https://plato.stanford.edu/entries/bergson/)). In his picture the past is not gone, it is simply not useful right now. Memory is *duration* — a living accumulation that persists rather than a shelf of records that survives — and remembering is less like retrieving a document than like turning your attention to a region of your own temporal extent.

**Reading B: the block-universe one — the past is still there.** Under eternalism, every moment of spacetime exists equally. 1900 and 2026 are both real, at different coordinates, the way New York and Paris are both real at different places. Minkowski's fusion of space and time is the standard physical backdrop, and the philosophy of time treats this as one live option among several: four-dimensionalism with temporal parts ([Sider](https://global.oup.com/academic/product/four-dimensionalism-9780199263523); [SEP: Temporal Parts](https://plato.stanford.edu/entries/temporal-parts/)), the Rietdijk–Putnam argument for the reality of other times, and the presentist opposition that denies the whole picture ([SEP: Time](https://plato.stanford.edu/entries/time/), [Presentism](https://plato.stanford.edu/entries/presentism/), [Being and Becoming in Modern Physics](https://plato.stanford.edu/entries/spacetime-bebecome/)).

On this reading a memory needs no storage location, because the original event is still at its own coordinate. The brain's job is not to keep a copy; it is to be a structure complex enough to re-establish a relation to that coordinate. Time does the storing. The brain is the key, not the box.

Both readings converge on the same inversion: the brain is an *enabling condition*, not a *storage location* — and we mistook one for the other because the brain is the thing we can watch working while remembering happens.

## Where the argument is harder than it looks

Two steps are weaker than the rhetoric suggests, and both matter for what follows.

**The physics does not deliver the conclusion by itself.** Eternalism is a live interpretation, not a settled result; relativity removes a privileged global "now" while the philosophy of physics continues to argue about presentism, the growing block, and the moving spotlight ([SEP: Time](https://plato.stanford.edu/entries/time/)). Rovelli's *The Order of Time* goes further, treating time as relational all the way down — and then lands somewhere that cuts *against* the simple "the past is just elsewhere" picture: the direction of time, he argues, is not a feature of the geometry but something we recover from the *records* the world leaves behind ([Rovelli](https://www.penguinrandomhouse.com/books/557814/the-order-of-time-by-carlo-rovelli/)). If traces are what give time its arrow, then "time stores memory" quietly inverts into "memory is what makes time storable" — a stranger claim than the block universe, and more interesting.

**Even granting the block universe, you still need a key.** If the event at t₁ exists, what makes a brain at t₂ *about* that event rather than about any other coordinate? Something has to connect them: a causal chain, a structural correspondence, a disposition to respond to the right thing. Whatever that something is, it is a trace — and traces are what the causal theory of memory is built on ([SEP: Memory](https://plato.stanford.edu/entries/memory/)). So the storage metaphor may not be *false* so much as *misplaced*: the trace is not the repository of memory, it is the hook that lets a present process re-instantiate a past one. That is the structure the extended and enactive traditions build on — cognition as something the whole organism does in a world over time ([Clark & Chalmers, 1998](https://www.consc.net/papers/extended.html); [Varela, Thompson & Rosch](https://archive.org/details/embodiedmindcogn0000vare); [SEP: Embodied Cognition](https://plato.stanford.edu/entries/embodied-cognition/)).

There is a nice empirical bridge. Divers who learned word lists underwater recalled them better underwater than on land, and vice versa ([Godden & Baddeley](https://doi.org/10.1111/j.2044-8295.1975.tb01468.x), 1975). The environment is part of the retrieval key. The storage metaphor predicts that the file is in the head and the room is irrelevant. The data says the room is a key.

And a caution for anyone smuggling Bergson or Whitehead in for aesthetics: *Process and Reality* (1929) is not a thesis about neurobiology but a metaphysics in which the past is *causally efficacious in the present* by being objectified into it ([Whitehead](https://archive.org/details/processrealityes0000whit)). It is an ontology of influence, not a claim that memories hang in spacetime like coats on a rack. Reading it as physics is how a good idea becomes a bad one.

## The same metaphor, in the machine

Here is where the philosophy stops being a spectator sport. The storage metaphor is not only how we think about brains; it was the founding assumption of AI, and the two fields imported it from each other.

Classical symbolic AI was a database. Knowledge was stored as rows — `Paris = capital_of(France)` — in semantic networks, frames, and ontologies, and reasoning was retrieval plus rule application. Memory was literally a table, and the engineering problem was how to file things in it.

Connectionism threw that away and replaced it with something much stranger. A neural network has no rows. What it has is a pattern of connection strengths, and a "fact" exists only as a disposition distributed across millions of weights, re-instantiating itself when the right input arrives. There is no folder in GPT for "Paris." There is no address you can point at and say: here is where that lives. The whole apparatus is closer to the radio than to the hard drive.

This is not a poetic reading of machine learning; it is the standard vocabulary. We call it **parametric memory** — knowledge held in weights, as opposed to **non-parametric memory**, knowledge held in something you can look up ([Petroni et al.](https://arxiv.org/abs/1909.01066), 2019). Which means Trumbull's sentence — *the system enables memory without containing it* — is a description of a trained network, not an analogy for one. Reductive materialism, in its AI form, says intelligence is *in the weights*; the everyday engineering experience of a neural network is that nothing is *in* anything.

The evidence that there is no folder is that we keep failing to edit one. If "Paris is the capital of France" were a row, changing it would be a `UPDATE` statement. Instead, locating and editing a single factual association requires finding the specific mid-layer activations that carry it, and even then the edit is entangled with everything else those weights do ([Meng et al.](https://arxiv.org/abs/2202.05262), ROME, 2022). Distributed is not a slogan here; it is a budget line.

And the failure mode matches too. Trumbull's point about reconstruction is why human recall confabulates. Machine learning has the same shape from the same cause: a generative model does not retrieve a stored sentence, it reconstructs a continuation from a distribution, and when the reconstruction has no grounding the result is fluent and wrong. The literature is now explicit that this is not a bug on the way to being patched — hallucination is a structural consequence of the setup ([Xu & Jain](https://arxiv.org/abs/2401.11817), 2024; for the human side of the taxonomy, [Ji et al.](https://arxiv.org/abs/2202.03629)). Confabulation is what reconstruction looks like when the key turns in the wrong lock.

I want to be careful here, because the analogy is easy to oversell. Human confabulation and LLM hallucination share a *shape* — construction without verification, fluency without grounding — but not a mechanism: one runs through reconsolidation and episodic memory in an organism with stakes, the other through next-token distributions over a corpus. The useful claim is structural, not anatomical. **Both are keys, and neither is a box.**

## The two readings, in code

Trumbull's two readings of "time stores memory" are not philosophy-only. Machine learning has spent thirty years building one architecture for each, and arguing about which is right.

**Recurrence: memory lives in the process, not in the weights.** In a recurrent network the model does not need to store the past. It carries a hidden state forward, and the past exists as *how it has changed the present state*:

```python
h_t = f(h_{t-1}, x_t)   # the past is not retrieved; it persists
```

That is Bergsonian duration, in code. Nothing is looked up. The past is present as accumulated modification, and the family resemblance to "the past is not gone, it is not useful right now" is not accidental — the architecture was designed to model exactly that.

The disanalogy matters, though, and it is the second half of the story. A hidden state is a *compression* of the past with a fixed width, not a retention of all of it. Bergson's duration retains everything and prioritizes by attention; a vanilla RNN forgets, and its gradient signal decays over long horizons ([vanishing gradients](https://doi.org/10.1109/72.279181), 1994). Every engineering response since — gates in an LSTM, and now selective state-space models like [Mamba](https://arxiv.org/abs/2312.00752) — is a negotiation with that fundamental difference: how much of the past can a fixed-size state carry before it must be summarized?

**Attention: leave the past where it is and learn to point at it.** Transformers made the opposite move. Instead of compressing history into weights or state, they hold the whole sequence and let every position look at every other position. Nothing is zipped, because nothing is unzipped: each token's representation is *recomputed* from the entire context on every forward pass.

Two corrections to the popular telling. Attention is not "going back 40 tokens"; it is content-addressed, addressing by similarity rather than by coordinate (positions enter as *added* information, not as the mechanism). And the KV cache is not memory in the storage sense: it is a derived artifact of a single generation, recomputed from scratch next time, present only while the process runs. The past there is *held*, not *kept*.

The engineering trajectory since is the interesting part. We started by trying to store everything inside the model — bigger weights, more parameters, more facts baked in at pretraining. We have spent the last several years going the other way, and the results keep coming back in favour of the other way: [RAG](https://arxiv.org/abs/2005.11401) (retrieve documents, put them in the context), [kNN-LM](https://arxiv.org/abs/1911.00172) (a datastore of representations you can query, whose paper's title is a lovely irony — *Generalization through Memorization*), [RETRO](https://arxiv.org/abs/2112.04426) (retrieval at trillion-token scale), [Memorizing Transformers](https://arxiv.org/abs/2203.08913) (attention over an explicit external memory), and now the entire industry of vector databases and long-context windows.

That is Trumbull's move, executed by engineers who have never read Bergson: *don't force the model to contain the past; leave the past out in the world where it already is, and make the model a structure that knows how to establish a relation to it.*

With the same caveat the brain has. Having the past laid out is not the same as accessing it well: retrieval quality degrades sharply depending on *where in the context* the relevant passage sits, with material in the middle of a long window used worst of all ([Liu et al.](https://arxiv.org/abs/2307.03172), *Lost in the Middle*, 2023). The key has to fit the lock. Storage was never the hard part.

## Case study: the zip file in gradient boosting

The cleanest small example of all of this in practice is a gradient-boosted tree ensemble — and specifically the difference between XGBoost and CatBoost, which is usually explained as a feature-engineering advantage and is actually a disagreement about where memory should live.

Start with the naive move, which is the zip file rendered as a line of code. If you have a categorical column like `city`, the shortcut is **mean encoding**: replace the category with the average target for that category.

```python
# the zip-file move: compress all of history into one stored number
df["city_encoded"] = df.groupby("city")["target"].transform("mean")
```

Every row for Paris now carries one number, computed from *every* Paris row in the dataset — including rows that come after it and including its own label. The past has been compressed into a constant and stored in a column, and the future has been leaked into it. That is the storage metaphor as a bug: a value that *is* the memory, sitting in a cell, true regardless of when you read it.

CatBoost's answer is the philosophical one, and it is stated almost openly in the paper. Instead of computing the encoding once and storing it, it computes, for each example, a statistic from the examples that come *before* it in an ordering — and to make that possible in a static dataset it introduces what the authors call an **artificial "time"**:

> "Clearly, the values of TS for each example rely only on the observed history. To adapt this idea to standard offline setting, we introduce an artificial 'time', i.e., a random permutation σ of the training examples." — [Prokhorenkova et al., *CatBoost*](https://arxiv.org/abs/1706.09516), 2018

Read that again, because it is the whole essay in one sentence from a machine-learning paper: to avoid storing a memory as a value, you give the data a *time*, and the encoding becomes a relation to a history instead of a constant. The same mechanism, extended to gradients, is what CatBoost calls **ordered boosting** — it prevents *prediction shift*, the target leakage you get when the model uses the same rows to both fit the data and estimate what it should predict.

And now the correction that makes the analogy honest, because it is tempting to read "ordered" as "temporal" and stop. The paper's own word for it is *artificial*: by default that ordering is a **random permutation**, and CatBoost uses several of them across boosting steps. Your event stream's real chronology is not what the library is respecting. If your data has a genuine order that matters — the events of your business — you have to tell it so, which is why older CatBoost interfaces exposed a `has_time` switch documented as: *"Use the order of objects in the input data (do not perform a random permutation of the dataset at the preprocessing stage)… Default: FALSE (not used; permute input dataset)"* ([catboost R package source](https://github.com/catboost/catboost/blob/master/catboost/R-package/R/catboost.R)).

So `ordered` in ordered boosting means *ordered for anti-leakage purposes*, not *chronological*. Those are different guarantees, and confusing them is exactly the mistake I spent [the last post](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) warning about: label leakage is prevented structurally by the library, whereas temporal leakage is your responsibility, and it is prevented by time-ordered folds and point-in-time-correct features. A random permutation is a beautiful trick against one and a subtle hazard for the other.

With that caveat in place, the philosophical reading holds up, and it is genuinely satisfying. The other structural choice CatBoost makes is to build **oblivious trees**: the same splitting criterion at every level, which makes them balanced, faster to evaluate, and — in the paper's phrase — "less prone to overfitting." You cannot open a trained CatBoost model and find row #4521 in it. The training set's influence is spread across thousands of splits and leaf values; the model *enables* the prediction without containing the data. It generalizes because it did not store; it learned a structure that can re-derive.

XGBoost is not the villain of this story and does not deserve the strawman. It is also a distributed structure with no rows and no folder; it has supported categorical data natively since version 1.5, splitting on partitions (`value ∈ categories`) rather than forcing everything into numbers ([XGBoost docs](https://xgboost.readthedocs.io/en/stable/tutorials/categorical.html)). Historically the library expected you to bring your own encoding, which meant carrying the leakage risk yourself — that is a difference in *where the history lives during training*, not a difference between a database and a brain.

Which is the honest version of the analogy: **both are keys**. What differs is whether the history is precomputed into a column and then treated as timeless — the zip file — or recomputed per example from an ordering that defines what counts as the past. CatBoost's contribution to this essay is that it is a working, popular, ordinary piece of software whose designers independently decided that the past should live in an order rather than in a value.

## What this means for how we build

A warning first, because this is where the essay could go wrong. Machines really do store. A file on disk is exactly the box model, and for silicon the storage metaphor is accurate rather than metaphorical. So the takeaway is not "import Trumbull's metaphysics into your data layer." It is narrower: **do not assume a brain is doing what your database does**, and be suspicious when a cognitive claim arrives pre-shaped like a system you already know how to build.

The suspicion runs the other way too, because "memory" in agent harnesses is almost always a storage layer: embeddings as engrams, a vector database as the hippocampus, retrieval-augmented generation as unzip-and-read. If the philosophy is right, that design imports the failure modes of the metaphor:

- **Retrieved memories get treated as facts.** A file is authoritative; a reconstruction is a hypothesis. If recall is construction — which it is — then any system that learns from its own recalled state is compounding its own edits. That is the dynamic behind drift, entrenchment, and confidently retrieved false memories.
- **The room gets dropped.** Context-dependence means the situation is part of the memory. Retrieval keyed on text similarity alone discards state, task, provenance, and what the system was doing when it learned the thing. It is the software equivalent of recalling underwater words better on land.
- **Architecture assumes a powered-off shelf.** A brain never has one: when the process stops, the memory is not waiting. A log you can re-derive from is closer to that than a store you load from — the argument this blog keeps making from the other direction: [the event log is the honest memory](https://blog.hackspree.com/#stories-from-events), [narratives are reconstructions built from it](https://blog.hackspree.com/#events-as-the-source-of-truth), [a durable daemon is a process that survives the power cycle](https://blog.hackspree.com/#durable-daemons). And the boosters make the same point about point-in-time correctness: [the artifact should be frozen, not re-derived on every event](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern).

Which brings the economics back around. If the storage metaphor is wrong, then the plan it implies — make the container bigger so it can hold more — is a plan for buying hard drives. The alternative the evidence supports is a smaller structure with better access: less memorization, more reconstruction skill, memory as a *relation between system and world* rather than content inside the system. That is exactly the trajectory of the last few years, and it is the same thesis this blog has been arguing from the model side — [specialization, not size, is the lever](https://blog.hackspree.com/#specialized-small-language-models), and the deployable unit is [the model plus the harness](https://blog.hackspree.com/#better-harnesses-smaller-models) that supplies what the model should not be carrying.

And the honest limit. None of this is a design specification. Philosophy can tell you the storage metaphor is a bad map of biological memory and a leaky one for models. It cannot tell you how to build better agent memory, and "the past still exists at its coordinate" is not a caching strategy. What it changes is the default assumption and the failure mode you watch for: memories that are confidently retrieved, unaccountably shaped by context, and quietly rewritten every time you read them.

## What this changed in my view

Four things moved while I was working through this, and one didn't.

**I had assumed the engram debate was settled empirically, and in the storage metaphor's favour.** Find the trace, tag it, activate it — Tonegawa-style experiments really do this, and I had filed it as "the box exists, the philosophers are arguing about vocabulary." The distributed-engram work changed that: the trace exists and is *not* a location. Cells across many regions contribute pieces, overlap across memories, and require the system's ongoing state to be reconstituted. That is a result the storage metaphor cannot represent — stronger than a philosophical preference.

**I had treated "memory is reconstruction" as a slogan about false memory.** Reconsolidation makes it a mechanism: recall destabilizes, edits, re-saves. As a slogan it is a warning about eyewitness testimony. As a mechanism it is a warning about any system that reads its own memory and writes back — including the agents we are building.

**I had the model side of the analogy backwards in an instructive way.** I expected the interesting example to be a large language model, where the story is diffuse and hard to verify. The cleaner example turned out to be seventy lines of gradient boosting — because CatBoost's paper says the quiet part out loud. It introduces an *artificial time* to avoid storing a memory as a value, and it hands you a switch to say "no, use the real order." A library telling you which ordering its memory depends on is the most concrete statement of Trumbull's thesis I have found in engineering.

**And I nearly wrote the tidy version of that.** My first draft of this section claimed CatBoost respects time. It does not; it respects *an* order, usually a random one, and `ordered` is about leakage rather than chronology. The correction made the essay better rather than weaker, because it separates two things the storage metaphor also conflates: keeping the past *around* and keeping it *in order* are different achievements, and only one of them is free.

What did not move: brains are physical, engrams are real, weights are real, and the storage metaphor is not worthless. It is just load-bearing in the wrong place — which is the most consequential kind of wrong a metaphor can be.

## References

**The primary argument**

- Trumbull, V. [*Memories are not stored in the brain*](https://iai.tv/video/meories-are-not-stored-in-the-brain-victoria-trumbull) — Institute of Art and Ideas talk, 2026. The claim under discussion: the brain enables rather than contains, and the storage metaphor mistakes an enabling condition for a storage location. Figure above links to the talk.
- Michaelian, K. and Sutton, J. [*Memory*](https://plato.stanford.edu/entries/memory/) — Stanford Encyclopedia of Philosophy. Traces, the causal theory of memory, distributed and extended accounts, and the objections to each.
- Semon, R. [*The Mneme*](https://archive.org/details/cu31924100387210) (1921; German original 1904) — the source of the term *engram*, and of *ecphory*: Semon already distinguished the trace from the act of reactivating it, which is a useful reminder that the box was never the whole theory.
- Danziger, K. [*Marking the Mind: A History of Memory*](https://archive.org/details/markingmindhisto0000danz) — Cambridge University Press, 2008. How the storage vocabulary — wax, seals, traces, vibrations, switches, files — was carried across four centuries of theories, and what each metaphor smuggled in.

**The three failures**

- Nader, K., Schafe, G. and LeDoux, J. [*Fear memories require protein synthesis in the amygdala for reconsolidation after retrieval*](https://doi.org/10.1038/35021052) — *Nature* 406, 2000. Retrieval destabilizes and re-stores: the rewrite-on-read result.
- Bartlett, F. C. [*Remembering: A Study in Experimental and Social Psychology*](https://archive.org/details/rememberingstudy00bart) — Cambridge University Press, 1932. Reconstruction in memory, decades before the neuroscience.
- Loftus, E. and Palmer, J. [*Reconstruction of automobile destruction: an example of the interaction between language and memory*](https://doi.org/10.1016/S0022-5371(74)80011-3) — *Journal of Verbal Learning and Verbal Behavior* 13(5), 1974.
- Schacter, D. and Addis, D. R. [*The cognitive neuroscience of constructive memory: remembering the past and imagining the future*](https://doi.org/10.1098/rstb.2007.2087) — *Philosophical Transactions of the Royal Society B* 362, 2007.
- Josselyn, S., Köhler, S. and Frankland, P. [*Finding the engram*](https://doi.org/10.1038/nrn4000) — *Nature Reviews Neuroscience* 16, 2015. The history of the search, Lashley's failure, and what engram cells actually are.
- Josselyn, S. and Tonegawa, S. [*Memory engrams: Recalling the past and imagining the future*](https://doi.org/10.1126/science.aaw4325) — *Science* 367, 2020. The strongest case that something trace-like exists, and how dynamic and distributed it is.
- Roy, D. et al. [*Brain-wide mapping reveals that engrams for a single memory are distributed across multiple brain regions*](https://doi.org/10.1038/s41467-022-29384-4) — *Nature Communications* 13, 2022.
- Raichle, M. and Gusnard, D. [*Appraising the brain's energy budget*](https://doi.org/10.1073/pnas.172399499) — *PNAS* 99(16), 2002; Raichle, M. et al. [*A default mode of brain function*](https://doi.org/10.1073/pnas.98.2.676) — *PNAS* 98(2), 2001. The never-off property.
- Godden, D. and Baddeley, A. [*Context-dependent memory in two natural environments: on land and underwater*](https://doi.org/10.1111/j.2044-8295.1975.tb01468.x) — *British Journal of Psychology* 66(3), 1975. The environment as part of the retrieval key.

**Time, duration, and the block universe**

- Bergson, H. [*Matter and Memory*](https://archive.org/details/mattermemor00berg) (1896; English translation 1911) — the brain as an organ of attention to action rather than a store, and duration as the persistence of the past; background in the [SEP entry on Bergson](https://plato.stanford.edu/entries/bergson/).
- Whitehead, A. N. [*Process and Reality*](https://archive.org/details/processrealityes0000whit) (1929; corrected edition 1979) — the past as causally efficacious in the present through objectification. Metaphysics, not neurobiology, and worth keeping that way.
- Markosian, N. et al. [*Time*](https://plato.stanford.edu/entries/time/) — SEP, including eternalism, presentism, and the Rietdijk–Putnam argument; see also [*Presentism*](https://plato.stanford.edu/entries/presentism/) and [*Being and Becoming in Modern Physics*](https://plato.stanford.edu/entries/spacetime-bebecome/).
- Sider, T. [*Four-Dimensionalism: An Ontology of Persistence and Time*](https://global.oup.com/academic/product/four-dimensionalism-9780199263523) — Oxford University Press, 2001; [SEP: Temporal Parts](https://plato.stanford.edu/entries/temporal-parts/) for the surrounding map.
- Canales, J. [*The Physicist and the Philosopher: Einstein, Bergson, and the Debate That Changed Our Understanding of Time*](https://www.degruyterbrill.com/document/doi/10.1515/9781400873310/html) — Princeton University Press, 2015. The 1922 Paris debate, and why the philosophical side of it lost so comprehensively that we stopped reading it.
- Rovelli, C. [*The Order of Time*](https://www.penguinrandomhouse.com/books/557814/the-order-of-time-by-carlo-rovelli/) — Riverhead, 2018. Time as relational, and the arrow of time as something recovered from records rather than baked into geometry.

**Parametric and non-parametric memory in machines**

- Petroni, F. et al. [*Language Models as Knowledge Bases?*](https://arxiv.org/abs/1909.01066) — 2019. The clean statement of parametric knowledge: facts recoverable from weights by prompting, with no lookup table anywhere.
- Meng, K. et al. [*Locating and Editing Factual Associations in GPT*](https://arxiv.org/abs/2202.05262) — ROME, 2022. Why editing "one fact" is hard: there is no row to update, only distributed circuitry that everything else also uses.
- Xu, Z. and Jain, S. [*Hallucination is Inevitable: An Innate Limitation of Large Language Models*](https://arxiv.org/abs/2401.11817) — 2024. Why reconstruction without grounding yields confabulation structurally rather than accidentally; [Ji et al.](https://arxiv.org/abs/2202.03629) for the human-side taxonomy of the same failure.
- Khandelwal, U. et al. [*Generalization through Memorization: Nearest Neighbor Language Models*](https://arxiv.org/abs/1911.00172) — 2019; Lewis, P. et al. [*Retrieval-Augmented Generation for Knowledge-Intensive NLP Tasks*](https://arxiv.org/abs/2005.11401) — 2020; Borgeaud, S. et al. [*Improving language models by retrieving from trillions of tokens*](https://arxiv.org/abs/2112.04426) — RETRO, 2021; Wu, Y. et al. [*Memorizing Transformers*](https://arxiv.org/abs/2203.08913) — 2022. The turn from "put the past in the weights" to "leave the past outside and learn to point at it".
- Liu, N. et al. [*Lost in the Middle: How Language Models Use Long Contexts*](https://arxiv.org/abs/2307.03172) — 2023. Having the past in the context is not the same as being able to use it.
- Bengio, Y., Simard, P. and Frasconi, P. [*Learning long-term dependencies with gradient descent is difficult*](https://doi.org/10.1109/72.279181) — *IEEE Transactions on Neural Networks* 5(2), 1994; Gu, A. and Dao, T. [*Mamba: Linear-Time Sequence Modeling with Selective State Spaces*](https://arxiv.org/abs/2312.00752) — 2023. Why a fixed-width state cannot retain everything, and what the successors do about it.

**The case study: where the past lives in gradient boosting**

- Prokhorenkova, L. et al. [*CatBoost: unbiased boosting with categorical features*](https://arxiv.org/abs/1706.09516) — NeurIPS 2018. The source of the *artificial time* quote, ordered target statistics, prediction shift, and oblivious trees ("the same splitting criterion is used across an entire level… balanced, less prone to overfitting, and allow speeding up execution at testing time"). [CatBoost documentation](https://catboost.ai/docs/en/concepts/algorithm-main-stages) for the implementation.
- [CatBoost R package source](https://github.com/catboost/catboost/blob/master/catboost/R-package/R/catboost.R) — the `has_time` parameter: "Use the order of objects in the input data (do not perform a random permutation of the dataset at the preprocessing stage)… Default: FALSE (not used; permute input dataset)". The switch that makes the distinction between *an* order and *your* order explicit.
- [XGBoost: Categorical Data](https://xgboost.readthedocs.io/en/stable/tutorials/categorical.html) — native categorical support since version 1.5, with partition-based splits (`value ∈ categories`) or one-hot encoding, and the option to let XGBoost handle the encoding rather than pre-encoding into numbers.
- For where the two guarantees diverge — anti-leakage ordering versus genuine temporal correctness — see [*LLM at Runtime Is an Antipattern*](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) on point-in-time folds, and Feast's [point-in-time joins](https://docs.feast.dev/getting-started/concepts/point-in-time-joins) for the operational version.

**Minds that are not sealed containers**

- Clark, A. and Chalmers, D. [*The Extended Mind*](https://www.consc.net/papers/extended.html) — *Analysis* 58(1), 1998. Cognition as something that can run partly outside the skull.
- Varela, F., Thompson, E. and Rosch, E. [*The Embodied Mind: Cognitive Science and Human Experience*](https://archive.org/details/embodiedmindcogn0000vare) — MIT Press, 1991. Enaction: mind as something an organism does in a world over time.
- [*Embodied Cognition*](https://plato.stanford.edu/entries/embodied-cognition/) — SEP. The philosophical landscape of embodiment, enaction, and their disagreements.
- James, W. [*Human Immortality: Two Supposed Objections to the Doctrine*](https://archive.org/details/humanimmortality00jame) — 1898. The transmission theory: the brain as transmitter rather than secreter — the radio analogy, argued in public, a century early.

**This blog, on memory and logs**

- [Make Stories from Events](https://blog.hackspree.com/#stories-from-events), [Events as the Source of Truth](https://blog.hackspree.com/#events-as-the-source-of-truth), [Durable Daemons](https://blog.hackspree.com/#durable-daemons), [Every workflow is an FSM](https://blog.hackspree.com/#every-workflow-is-an-fsm), [SSLM: Specialized Small Language Models](https://blog.hackspree.com/#specialized-small-language-models), [Better Harnesses, Smaller Models](https://blog.hackspree.com/#better-harnesses-smaller-models), [LLM at Runtime Is an Antipattern](https://blog.hackspree.com/#llm-at-runtime-is-an-antipattern) — the software-side versions of the same question: what is stored, what is re-derived, and what is only ever reconstructed.

---
title: LLMs Can't Jump
date: 2026-07-29
slug: llms-cant-jump
summary: Einstein described discovery as a cycle: induction, deduction, and a mysterious third thing — the intuitive jump from experience to axioms. Tom Zahavy's position paper argues that LLMs have mastered the first two and are structurally incapable of the third. The jump is abduction. Abduction is the search problem of design. And it requires something LLMs don't have: a body.
tags: ai, llm, abduction, scientific-discovery, deepmind, epistemology, design
---

Einstein wrote a letter to Maurice Solovine in 1952. In it, he drew a diagram of how discovery works. Sensory experience forms the base. From it, the mind makes an intuitive leap to axioms — general principles that are not logically derivable from the data. From those axioms, the mind deduces consequences. The consequences are tested against experience. The loop repeats. Einstein called the leap *Aufstieg* — ascent. He said it was the most mysterious part of the process, the part that cannot be mechanized.

Tom Zahavy, at Google DeepMind, has written a position paper arguing that Einstein was right. The paper is called *LLMs Can't Jump*. The jump is abduction — the generation of novel explanatory hypotheses from sparse or absent data. Zahavy's argument is that LLMs have mastered induction (statistical pattern matching) and are rapidly conquering deduction (formal proof). They lack the mechanism for the jump. They cannot do it. Their architecture prevents it.

[Watch: Tom Zahavy presents LLMs Can't Jump — The Abductive Gap in AI Discovery](https://www.youtube.com/watch?v=vaH2eweyuvs)

> Induction is interpolation within the known. Deduction is derivation from premises. Abduction is the leap to premises that do not yet exist. LLMs interpolate. They derive. They do not leap.

The case study is General Relativity. When Einstein began working on it in 1907, Newtonian mechanics was wildly successful. There was no crisis in the data. The only anomaly was Mercury's orbital precession — a tiny discrepancy that most physicists considered an measurement error. Einstein did not have a dataset that demanded a new theory of gravity. He had a thought experiment: a person falling freely in an elevator would not feel their own weight. From this embodied intuition — not from gradient descent over a loss function — he derived the equivalence principle, and from there, the field equations of General Relativity.

> Einstein did not compress data. He jumped from it. The jump was not an interpolation between known points. It was the creation of a new point — a new axis in the conceptual space. LLMs cannot create new axes. They can only navigate the ones we have already named.

Zahavy identifies the structural reasons. LLMs process tokens, not physical experience. "Gravity" is a statistical relationship between words, not a sensation. The model has never fallen. It has never felt weight disappear. It cannot perform the embodied thought experiment that gave Einstein his axioms. This is the symbol grounding problem, applied to discovery. You cannot reason about what you have not experienced. You cannot jump from a platform you have never stood on.

The prevailing theory of creativity in AI is compression — Schmidhuber's thesis that creativity is just efficient data compression. A sufficiently good compressor, given enough data, will produce general intelligence. Zahavy's counterargument is General Relativity. There was no data to compress. Newtonian mechanics fit almost all the data perfectly. The discovery came from a jump away from the data, not a better fit to it. Compression explains induction. It does not explain abduction. A compressor can find the shortest representation of what is. It cannot represent what is not yet.

> Compression tells you what the data contains. Abduction tells you what the data implies but does not contain. The first is interpolation. The second is invention. LLMs are compressors. The jump requires something else.

Zahavy's proposed solution is multimodal world models — AI systems that can interact with simulated physical environments, perform counterfactual interventions, and ground their representations in sensory experience. DeepMind's Genie architecture, which allows action-controllable interaction with generated worlds, is a step in this direction. The idea is that you cannot jump from tokens. You need a body — even a simulated one — that can fall, collide, and observe the consequences. The translation of simulation into formal axioms remains the critical bottleneck. But at least the simulation provides something to jump *from*.

> Einstein's elevator was a simulation running in his own sensorimotor cortex. He had fallen. He had felt acceleration. His body knew what weightlessness felt like before his mind could formalize it. An LLM has no elevator to fall in.

This connects to a deeper argument this blog has been making. Design is a search problem. The stack constrains the search. Style is the trained heuristic that makes the search tractable. But abduction — the jump — is the search problem at its most radical. It is a search that must create the space it searches. The designer who sees David in the marble is performing an abductive leap. The axioms are not in the data. They are in the mind's eye, projected onto the stone, tested by the chisel. Brooks called this seeing the design in the mind's eye before it is built. Einstein called it the intuitive leap from experience to axioms. Both were describing the same thing. Neither could mechanize it.

> Abduction is the search problem of design at its origin. It is the moment the search space itself is created. LLMs search spaces we define. They cannot define new ones. The jump is the act of definition. It is the one thing we have not automated. It may be the one thing we cannot.

---

**References:**

- Tom Zahavy. (2026). [LLMs Can't Jump](https://philsci-archive.pitt.edu/28024/). Google DeepMind. ICML 2026.
- Albert Einstein. (1952). Letter to Maurice Solovine. In *Letters to Solovine*, Philosophical Library.
- Jürgen Schmidhuber. (2009). Driven by Compression Progress. *arXiv:0812.4360*.
- Related: [On Finding David in the Marble](https://blog.hackspree.com/#finding-david-in-the-marble) — Design as search, the intuitive leap from form to marble.
- Related: [The Stack as a Design Abstraction](https://blog.hackspree.com/#the-stack-is-the-abstraction) — How constraints make search tractable, and why agents need them.

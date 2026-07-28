---
title: Finding David in the Marble — Iteration and The Design of Design
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo said he didn't carve David. He just chipped away everything that wasn't David. Brooks' The Design of Design argues the same thing about software: great design is not built. It is found. Through iteration. By removing what does not belong.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

In 1501, a 26-year-old sculptor was given a block of marble. It was not a fresh block. It had been sitting in a cathedral courtyard for thirty-five years. Two other sculptors had tried to carve it and abandoned it. The legs were roughed out. The block was considered flawed, unwieldy, and probably unusable. The sculptor was Michelangelo. The block became David.

> The marble was not special. The sculptor was not magical. The process was subtractive. Michelangelo saw what the statue could be and removed everything that wasn't it.

![Michelangelo's David — carved from a block of marble abandoned for 35 years, considered flawed. The process was not additive. It was subtractive.](/images/david-michelangelo.jpg)

Fred Brooks published *The Design of Design* in 2010, thirty-five years after *The Mythical Man-Month*. The later book is less famous. It is more important. Brooks' central argument: great design is an empirical process, not a rational one. The rational model — plan everything in advance, specify completely, build once — is a fantasy. It works for problems we have solved before. It fails for problems worth solving. The alternative is iteration. Build something. Learn what is wrong with it. Remove what is wrong. Build again. The design does not emerge from the plan. It emerges from the removal.

> The rational model says: design, then build. The empirical model says: build, then remove, then build again. The first requires omniscience. The second requires honesty. Only one of these is available to humans.

This is Michelangelo's method applied to software. He did not assemble David from pieces. He did not have a blueprint precise enough to hand to an assistant. He worked the stone directly. He removed material. At each stage, the emerging form told him what to remove next. The statue was already in the marble. His job was to find it. A software design is already in the problem. The designer's job is to find it — not by getting it right the first time, but by getting it wrong and removing the wrongness.

Brooks is explicit about this. In *The Design of Design*, he rejects the waterfall model as a description of how real design happens. He endorses Barry Boehm's spiral model — design, prototype, evaluate, redesign. Each iteration is a pass at the stone. Each pass removes what doesn't belong. The design gets simpler, not more complex. The final system has fewer parts than the intermediate versions. It is not that features were added. It is that features were removed.

> A system grows by accretion. A design grows by reduction. The first is assembly. The second is sculpture.

This is why Brooks' three principles of conceptual integrity are all negative constraints. **Propriety**: do not introduce what is immaterial. If a feature does not earn its place, remove it. **Orthogonality**: do not link what is independent. If two things are coupled, separate them. **Generality**: do not restrict what is inherent. If a limitation is artificial, eliminate it. Each principle is a chisel. Each application removes something. Propriety removes features. Orthogonality removes coupling. Generality removes constraints. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks' three principles are three chisels. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The statue emerges.

Michelangelo's David was not the first attempt on that block. Agostino di Duccio worked it for two years and quit. Antonio Rossellino tried and quit. The block sat exposed to weather for three and a half decades. Everyone who saw it saw a flawed block. Michelangelo saw David. This is what separates a designer from a builder. The builder sees what is there. The designer sees what could be there — and has the discipline to remove everything else.

Software projects have their own abandoned blocks. The legacy codebase that everyone says needs a rewrite. The module that three teams have tried to refactor and given up on. The architecture that accumulated so many exceptions it no longer has a ruling principle. The instinct is to start over. New language. New framework. Clean slate. Michelangelo did not start over. He took the abandoned block — the one with the roughed-out legs, the one everyone said was ruined — and he removed what didn't belong to the David he saw. The legacy codebase is not a ruined block. It is a block someone roughed out and abandoned. Your job is to find the design inside it and remove everything else.

> Starting over is the fantasy of the rational model. Finding the design in the existing system is the discipline of the empirical one. The marble is already on your desk. It has been for thirty-five years.

Iteration is not about getting it right. It is about getting it less wrong. Each cycle, you remove something. A feature that doesn't earn its place. A coupling that shouldn't exist. A constraint that serves no purpose. The codebase shrinks. The design clarifies. You are not building David. You are finding him. He was always in there. He was just buried under the features, the couplings, and the constraints that were not him.

## Design is a search problem

Herbert Simon, in *The Sciences of the Artificial* (1969), formalized what Michelangelo did intuitively. Design is a search problem. The design space is the set of all possible designs that satisfy the constraints — every statue that could be carved from the block, every architecture that could organize the code. The designer's job is to navigate this space. Each iteration is a move from one point in the space to a neighboring point. Each evaluation is a judgment of whether the new point is better than the old one.

> The design space is infinite. The constraints make it finite. The designer's style makes it navigable.

Brooks adopted Simon's framework directly in *The Design of Design*. The rational model has four components. **Goals**: what the design must achieve. **Utility function**: how you compare two designs to decide which is better. **Constraints**: boundaries the design must not cross. **Resources**: time, people, budget. The designer proposes a tentative design, evaluates it against the utility function, and iterates. This is a search through a combinatorial space. The goal is not optimality — the space is too large. The goal is *satisficing*, Simon's term for a design that is good enough given the constraints. You stop searching when the cost of further search exceeds the expected improvement.

Michelangelo's search space was defined by the marble block. The block's dimensions were the constraints. The existing rough cuts from di Duccio were constraints. The commission's requirements — a figure of David, biblical hero — were the goals. Michelangelo's utility function was his eye. He looked at the emerging form and judged whether it was closer to David or further away. Each strike of the chisel was a move in the search space. The search terminated when the form matched what he saw in his mind.

> Michelangelo did not enumerate the search space. He did not evaluate every possible statue in the block. His style — his accumulated judgment about what forms were beautiful, what proportions were right, what surface was finished — pruned the space to a single path. Style is the search heuristic.

This is Brooks' most undervalued insight. In a chapter called "Esthetics and Style in Technical Design," he argues that style is not decoration. It is the thing that makes search tractable. A designer without style evaluates every option. A designer with style knows which options are worth evaluating. Style is developed by studying other designers' work intentionally — not copying it, but understanding the choices they made and why. By making conscious judgments about what you like. By practicing until the judgments become instinct. By revising your own designs and noticing where your style was inconsistent. Style is not a personality. It is a trained search heuristic.

> Brooks: "Great designs come from great designers." Not from great processes. Not from great committees. From great designers — people whose style has pruned the search space so effectively that they find David where others see a flawed block.

Michelangelo developed his style by studying classical sculpture, dissecting cadavers to understand anatomy, and carving constantly from the age of thirteen. By twenty-six, when he took on the David commission, his style was sharp enough to see the statue in a block everyone else had abandoned. He did not search randomly. He did not evaluate every possible statue. His style told him where David was, and he removed everything else. A great software designer does the same. The style tells them which architecture to try first. Which abstraction to keep. Which feature to remove. The search space is too large for brute force. Style is what makes it small enough for a human mind.

> Style is the compression algorithm for the design space. It takes the infinite and makes it finite. It takes the navigable and makes it intuitive. Without it, you are enumerating possibilities until the budget runs out. With it, you are Michelangelo.

The most important thing Brooks said about design is not about process. It is about seeing. "The designer must see the design in the mind's eye before it is built." Michelangelo saw David in the abandoned block. The software designer must see the system in the problem before it is coded. This is not mysticism. It is the result of a trained search heuristic — style — operating on a constrained space. And the heuristic is developed through iteration: by building things, seeing what is wrong with them, removing what is wrong, and building again. You learn to see David by chipping away at enough blocks that weren't David.

> Taste is the residue of iteration. Style is the search heuristic that taste produces. The designer who has chipped enough marble sees the statue before the first strike of the chisel.

---

**References:**

- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- Fred Brooks. (1975). *The Mythical Man-Month*. Addison-Wesley.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment) — POLA, CoC, and conceptual integrity for agent-consumed software.
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux) — Conceptual integrity in operating system design.

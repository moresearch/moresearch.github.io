---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo did not carve David. He removed everything that was not David. Brooks' The Design of Design is a book about chisels — the ones you use on software, on constraints, on yourself. Great design is not built. It is found. Through iteration. By subtraction. By learning to see.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

The block had been waiting thirty-five years. Two sculptors had tried and walked away. The legs were roughed out, the proportions already compromised. The marble was considered ruined — too tall, too narrow, too damaged by weather and bad first cuts. In 1501 it was given to a twenty-six-year-old who looked at it and saw David.

> Michelangelo did not carve David. He removed everything that was not David. This is the first lesson of design. The form is already in the problem. Your job is to see it and set it free.

![Michelangelo's David — carved from a block of marble abandoned for 35 years. The process was subtractive. The form was already there.](/images/david-michelangelo.jpg)

Fred Brooks published *The Design of Design* in 2010. It is the book he spent thirty-five years learning to write. His argument is quiet and absolute: great design is an empirical act, not a rational one. The rational model — specify, architect, implement, ship — assumes you can see the end from the beginning. You cannot. Not for anything worth building. The alternative is iteration. Build. Learn what is wrong. Remove it. Build again. The design does not emerge from the plan. It emerges from the removal. Brooks called this Boehm's spiral, but the metaphor is older than software. It is the sculptor's rhythm. Strike. Assess. Strike again.

> The rational model requires omniscience. The empirical model requires honesty. One is a fantasy. The other is a discipline.

Brooks understood that conceptual integrity is the most important quality of a system. One design voice. Every part consistent with every other. But how do you achieve it? Not by adding. By removing. His three principles are all negative constraints, all chisels. **Propriety**: do not introduce what is immaterial. If a feature does not earn its place, strike it away. **Orthogonality**: do not link what is independent. If two things are coupled without cause, separate them. **Generality**: do not restrict what is inherent. If a limitation is artificial, free it. Each principle is a removal. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks gave us three chisels. Most of us use them to add. He meant them to take away.

Herbert Simon, whom Brooks drew from, said design is a search problem. The design space is every possible form the marble could take, every possible architecture the code could have. The space is infinite. You cannot enumerate it. You navigate it. Each iteration is a move. Each evaluation is a judgment: closer or further? Michelangelo's search space was bounded by the block — its height, its flaws, the existing cuts. His utility function was his eye. He looked at the emerging surface and knew whether David was nearer or farther. He did not evaluate every possible statue in the stone. He did not need to. His style — the accumulated judgment of a lifetime of looking and cutting — collapsed the infinite space into a single path.

> The design space is infinite. Constraints make it finite. Style makes it navigable.

Brooks devoted a chapter to style — "Esthetics and Style in Technical Design." It is the most overlooked chapter in the book. Style, he argued, is not personality. It is not decoration. It is the trained heuristic that tells you which way to move in the search space. A designer without style tries everything. A designer with style tries the right thing first. Style is acquired by studying other designers' work — not imitating it, but understanding the choices behind it. By making conscious judgments about what moves you and why. By practicing until judgment becomes instinct. By revising your own work and noticing where your hand wavered.

> "Great designs come from great designers." Not from process. Not from committees. From people whose style has pruned the infinite space into something a mind can hold.

Michelangelo developed his style by studying Greek and Roman sculpture. By dissecting cadavers in the morgue at Santo Spirito, learning how muscle lay over bone. By carving from the age of thirteen. By the time he stood before the abandoned block, he had cut enough marble to know what the marble wanted to become. He did not impose David on the stone. He listened to the stone and found David already inside, waiting to be released.

> You learn to see David by chipping away at enough blocks that were not David. Taste is the residue of iteration. Style is the search heuristic that taste produces.

The most important thing Brooks said is not about process. It is about seeing. The designer must see the design in the mind's eye before it is built. This is not mysticism. It is the result of a trained heuristic operating on a constrained space for long enough that the path becomes visible before the first step is taken. Michelangelo saw David in the ruined block. The software designer must see the system in the problem before the code is written. And the only way to learn to see is to build, to look honestly at what you built, to remove everything that isn't it, and to build again.

The marble is already on your desk. It has been for thirty-five years.

---

**References:**

- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

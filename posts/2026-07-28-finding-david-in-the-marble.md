---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo did not carve David. He removed everything that was not David. Brooks' The Design of Design is about the same thing: constraints as precondition, iteration as search, style as heuristic, and the loop where engineering and art become one motion.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

The block had been waiting thirty-five years.

It was quarried in 1464 for Agostino di Duccio. Too narrow for his figure. He roughed out the legs and quit. Antonio Rossellino tried a decade later, saw the veins in the marble, walked away. The stone sat in the courtyard of the Florence Cathedral for twenty-five more years. Rain. Sun. The Operai called it *lo gigante* — the giant — too large to move, too narrow to use, too compromised to trust.

In 1501 they gave it to a twenty-six-year-old. He looked at the narrowness and saw a figure who could stand within it. He looked at the flaws and saw features that could absorb them. He looked at what everyone called a ruined block and saw David.

> The constraints that defeated two sculptors made the third. The narrowness forced David's contrapposto — the weight shift, the coiled stillness. Without the flaw, there is no masterpiece.

![Michelangelo's David — carved from a block too narrow, too tall, abandoned for 35 years. The flaws defined him.](/images/david-michelangelo.jpg)

Fred Brooks published *The Design of Design* in 2010, the book he spent thirty-five years learning to write. Constraints, he argued, are not the enemy of design. They are the precondition. A blank page is paralysis. A constrained page is possibility. Great design is empirical, not rational — the rational model assumes omniscience. You do not have it. The alternative is iteration: build, learn what is wrong, remove it, build again. Brooks endorsed Boehm's spiral, but the metaphor is older than software. Strike. Assess. Strike again.

> The rational model requires omniscience. The empirical model requires honesty. One is a fantasy. The other is a discipline.

Brooks defined conceptual integrity as the most important quality of a system — one design voice, every part consistent with every other. How do you achieve it? Not by adding. By removing. His three principles are chisels. **Propriety**: do not introduce what is immaterial. **Orthogonality**: do not link what is independent. **Generality**: do not restrict what is inherent. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks gave us three chisels. Most of us use them to build. He meant them to take away.

Herbert Simon, in *The Sciences of the Artificial* (1969), formalized design as search. The design space is every possible form — infinite until constrained. The block's narrowness eliminated nine-tenths of what Michelangelo might have imagined. The rough cuts eliminated more. The veins eliminated more. What remained was not a ruined block. It was a sharply constrained corridor through the infinite, leading to David. Simon's term for this is *satisficing* — you stop searching not when you find the perfect design, but when the cost of further search exceeds the expected improvement. Constraints make satisficing possible.

Michelangelo's utility function was his eye. He did not enumerate. He did not evaluate every possibility. The constraints had already done most of the elimination. His style — the accumulated judgment of a lifetime — collapsed what remained into a single path.

> The design space is infinite. Constraints make it finite. Style makes it navigable.

Brooks devoted a chapter to style. Style, he argued, is not personality. It is the trained heuristic that tells you which way to move. A designer without style tries everything. A designer with style tries the right thing first. Style is acquired by studying other designers' work — understanding the choices, not imitating them. By making conscious judgments about what moves you. By practicing until judgment becomes instinct. Michelangelo developed his by dissecting cadavers at Santo Spirito, studying Greek proportion, carving from thirteen. By twenty-six he had cut enough marble to know what the marble wanted to become.

> "Great designs come from great designers." Not from process. Not from committees. From people whose style has pruned the infinite into something a mind can hold.

## The loop

Art and engineering are not similar disciplines. They are two moments in the same loop. In the first, you act on the system — build, measure, cut. In the second, you step back and judge — closer to David or further? Engineering without the art phase is accretion: systems that grow without shape, complexity no one intended. Art without the engineering phase is fantasy: beautiful on paper, impossible in code. The loop requires both.

> Strike. Assess. Strike. The sculptor's rhythm is the designer's rhythm. The only difference is the medium.

Michelangelo's loop: dissecting cadavers was engineering. Studying Polykleitos was engineering. Thirteen years of carving was engineering. Seeing David in the block was art. But the loop did not stop there. Each strike was engineering. Each assessment was art. For two years, the loop ran. The result was a thing with no unnecessary surface, no wrong angle. Conceptual integrity in stone. He was not an artist who happened to be good at engineering. He moved between the two moments so fluently they became one motion. Brooks called that fluency *style*.

A well-designed system has the same quality. A clean interface. A coherent abstraction. A module that does one thing completely. These are not engineering properties — no metric captures them. They are aesthetic properties, the residue of the art phase accumulated across enough iterations that the system feels inevitable. Trust is the aesthetic response to conceptual integrity. You trust a system whose parts agree, that does what it says and nothing else, from which someone removed everything that was not it.

> Engineering without the art phase produces systems that work and are hated. Art without the engineering phase produces designs that are loved and cannot be built. The loop is not optional. It is the discipline.

Brooks insisted that design must proceed from "one mind, or from a very small number of agreeing resonant minds." He was not talking about intelligence. He was talking about fluency in the loop. The designer who has built, removed, alternated enough times develops an eye. They see the system before it is built. They know which feature is immaterial. They know because they have recognized David enough times, in enough blocks of marble, in enough codebases.

> Every good engineer moves through both phases whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The marble releases what it was always holding.

## Open questions

**Is style teachable?** Brooks says yes — study others, practice, revise. But Michelangelo's style came from dissecting cadavers and carving from thirteen. Can you compress that into a curriculum? Or does style require the obsession that makes someone sleep beside a block of marble for two years?

**Does the loop work for teams?** Brooks says design must proceed from one mind or a very small number. But most software is built by teams. Does the loop scale? Or does conceptual integrity dilute with every additional mind in the loop?

**Can an AI coding agent develop style?** An agent can study every open-source repository ever written. It can practice instantly. But can it develop taste — the judgment that tells you this abstraction is right and that one is wrong? Or is taste the one thing that requires a human eye?

**Is satisficing still the right stopping condition?** Simon said stop when further search costs more than expected improvement. But software decays. The design that satisficed last year is today's technical debt. Is the loop ever really done, or do we just pause it?

**What is the equivalent of the abandoned block?** Michelangelo's block was a physical object with irrecoverable constraints. Software is mutable — you can always rewrite. Does that make constraints harder to accept? When everything can be changed, does anything feel like a constraint? And without constraints, where does the design come from?

---

**References:**

- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

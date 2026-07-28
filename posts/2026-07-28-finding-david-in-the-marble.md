---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo did not carve David. He removed everything that was not David. Brooks' The Design of Design is a book about chisels — the ones you use on software, on constraints, on yourself. Great design is not built. It is found. Through iteration. By subtraction. By learning to see.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

The block had been waiting thirty-five years.

It was quarried in 1464 for a sculptor named Agostino di Duccio. He worked it for two years and quit. The stone was too narrow for the figure he had planned. The proportions were wrong. He roughed out the legs and abandoned it. A decade later, Antonio Rossellino tried. He looked at the block — the awkward dimensions, the existing cuts, the veins running through the marble — and walked away. The stone sat exposed in the courtyard of the Florence Cathedral for another twenty-five years. Rain. Sun. The legs slowly weathering. The Operai, the committee in charge of the cathedral's artworks, considered it ruined. They referred to it as *lo gigante* — the giant — a block too large to move, too narrow to use, too compromised to trust.

In 1501, they gave it to a twenty-six-year-old. He looked at the narrowness and saw a figure who could stand within it. He looked at the flaws and saw features that could absorb them. He looked at what everyone called a ruined block and saw David.

> The constraints that defeated two sculptors made the third. They were not obstacles to the design. They were the design. The narrowness forced David's contrapposto — the weight shift, the tension, the coiled stillness that makes the statue a living thing. Without the flaw, there is no masterpiece.

![Michelangelo's David — carved from a block of marble too narrow, too tall, abandoned for 35 years. The flaws didn't prevent David. They defined him.](/images/david-michelangelo.jpg)

Fred Brooks published *The Design of Design* in 2010, the book he spent thirty-five years learning to write. His argument: constraints are not the enemy of design. They are the precondition. A blank page is paralysis. A constrained page is possibility. Great design is empirical, not rational. The rational model — specify, architect, implement, ship — assumes omniscience. You do not have it. Not for anything worth building. The alternative is iteration. Build. Learn what is wrong. Remove it. Build again. Brooks endorsed Boehm's spiral model, but the metaphor is older than software. It is the sculptor's rhythm. Strike. Assess. Strike again. Each strike answers to the stone.

> The rational model requires omniscience. The empirical model requires honesty. One is a fantasy. The other is a discipline.

Brooks understood that conceptual integrity is the most important quality of a system. One design voice. Every part consistent with every other. But how do you achieve it? Not by adding. By removing. His three principles are all negative constraints, all chisels. **Propriety**: do not introduce what is immaterial. If a feature does not earn its place, strike it away. **Orthogonality**: do not link what is independent. If two things are coupled without cause, separate them. **Generality**: do not restrict what is inherent. If a limitation is artificial, free it. Each principle is a removal. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks gave us three chisels. Most of us use them to add. He meant them to take away.

Herbert Simon, whom Brooks drew from, said design is a search problem. The design space is every possible form the marble could take, every possible architecture the code could have. The space is infinite — until you constrain it. The block's narrowness eliminated nine-tenths of the statues Michelangelo might have imagined. The rough cuts eliminated more. The veins in the marble eliminated still more. What remained was not a ruined block. It was a sharply constrained search space — a narrow corridor through the infinite, leading to David. The constraints did not limit Michelangelo. They guided him.

His utility function was his eye. He looked at the emerging surface and knew whether David was nearer or farther. He did not enumerate. He did not evaluate every possibility. He did not need to. The constraints had already done most of the elimination. His style — the accumulated judgment of a lifetime of looking and cutting — collapsed what remained into a single path.

> The design space is infinite. Constraints make it finite. Style makes it navigable.

Brooks devoted a chapter to style — "Esthetics and Style in Technical Design." It is the most overlooked chapter in the book. Style, he argued, is not personality. It is not decoration. It is the trained heuristic that tells you which way to move in the search space. A designer without style tries everything. A designer with style tries the right thing first. Style is acquired by studying other designers' work — not imitating it, but understanding the choices behind it. By making conscious judgments about what moves you and why. By practicing until judgment becomes instinct. By revising your own work and noticing where your hand wavered.

> "Great designs come from great designers." Not from process. Not from committees. From people whose style has pruned the infinite space into something a mind can hold.

Michelangelo developed his style by studying Greek and Roman sculpture. By dissecting cadavers in the morgue at Santo Spirito, learning how muscle lay over bone. By carving from the age of thirteen. By the time he stood before the abandoned block, he had cut enough marble to know what the marble wanted to become. He did not impose David on the stone. He listened to the stone and found David already inside, waiting to be released.

> You learn to see David by chipping away at enough blocks that were not David. Taste is the residue of iteration. Style is the search heuristic that taste produces.

The marble is already on your desk. It has been for thirty-five years.

## The loop

Art and engineering are not similar disciplines. They are phases in the same loop. The loop has two moments. In the first, you act on the system — you build, you measure, you cut. Call this engineering. In the second, you step back and judge — is it closer to David or further away? Call this art. Then you act again. Engineering without the art phase produces accretion: systems that grow without ever being shaped, features added and never removed, complexity that no one intended and no one can arrest. Art without the engineering phase produces fantasy: designs that are beautiful on paper and impossible in code. The loop requires both.

> Engineering and art are not two disciplines. They are two moments in one discipline. Strike. Assess. Strike. The sculptor's rhythm is the software designer's rhythm. The only difference is the medium.

Michelangelo dissected cadavers at Santo Spirito to understand how muscle lay over bone. This was engineering — measurement, analysis, craft. He studied the classical proportions of Polykleitos and Praxiteles. Engineering. He carved for thirteen years before touching a major commission. Engineering. Then he stood before the abandoned block and saw David. That was art — judgment, taste, the trained eye seeing what the stone already contained. But the art did not end there. He struck. He assessed. He struck again. Each strike was engineering. Each assessment was art. The loop ran for two years. The result was a thing with no unnecessary surface, no wrong angle, no part that does not belong to the whole. Conceptual integrity in stone.

> Michelangelo was not an artist who happened to be good at engineering. He moved between the two moments so fluently that they became one motion. That fluency is what Brooks called style.

Brooks titled the chapter "Esthetics and Style in Technical Design." He was not being poetic. He was being precise. A well-designed system has a quality that metrics cannot capture. A clean interface. A coherent abstraction. A module that does one thing completely. These are not engineering properties. They are aesthetic properties — the residue of the art phase, accumulated across enough iterations that the system feels inevitable. You look at it and sense that nothing is missing and nothing is extra. You trust it. Trust is the aesthetic response to conceptual integrity — and conceptual integrity is what the loop produces when both phases are executed with discipline.

This is what Brooks meant when he said great designs come from great designers. He was not talking about intelligence. He was talking about fluency in the loop. The designer who has built enough systems, removed enough wrongness, alternated between the engineering moment and the artistic moment enough times — that designer develops an eye. They see the system before it is built. They know which feature is immaterial. They know which coupling is unnatural. They know because they have seen David enough times to recognize him before the first strike.

> Every good engineer moves through both phases of the loop whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The design clarifies. The marble releases what it was always holding.

---

**References:**

- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

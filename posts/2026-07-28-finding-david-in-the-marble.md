---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo said he didn't carve David. He just removed everything that wasn't David. Plato would have called this remembering a Form. Aristotle would have called it actualizing a potential. Brooks called it design. All three were right.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

The block had been waiting thirty-five years.

It was quarried in 1464 and given to Agostino di Duccio, who worked it for two years before declaring it impossible. The stone was too narrow for any figure that could stand within it. He roughed out the legs and abandoned it. A decade later, Antonio Rossellino examined the block — the awkward dimensions, the gouges where di Duccio had cut, the veins running through the marble — and refused the commission. The block was hauled into the courtyard of the Florence Cathedral and left there. For twenty-five years it sat in the rain and sun, the roughed-out legs slowly weathering, a monument to two failures. The Operai, the cathedral's works committee, called it *lo gigante* — the giant. Too large to move. Too narrow to use. Too compromised to trust.

In 1501, with the block still occupying the courtyard and the Operai still paying storage, they gave it to a twenty-six-year-old sculptor who had recently returned from Rome. His name was Michelangelo Buonarroti. He had been carving since the age of thirteen. He had studied the Greek and Roman masters. He had dissected cadavers in the morgue at Santo Spirito, peeling back skin to understand how muscle lay over bone. He accepted the commission on August 16. He built a wooden shed around the block so no one could watch him work. He slept beside the marble. For two years he struck the stone in secret. When the shed came down, David stood where a ruined block had been.

> Michelangelo did not carve David. He removed everything that was not David. The form was already in the stone. His job was to set it free.

![Michelangelo's David — carved from a block too narrow, too tall, abandoned for 35 years. The flaws did not prevent the masterpiece. They defined it.](/images/david-michelangelo.jpg)

Plato would have understood this immediately. In the *Theory of Forms*, the true reality of a thing is not its physical instantiation but its ideal Form — perfect, eternal, unchanging. The David that stands in the Galleria dell'Accademia is a shadow of the Form of David that exists in the realm of ideas. Michelangelo did not invent David. He remembered him. The sculptor's eye, trained by years of looking at what others had carved and what nature had made, could perceive the Form through the flawed stone. His chisel merely removed the matter that obscured it.

Aristotle, Plato's student and critic, would have corrected his teacher. The Form does not exist in a separate realm. It exists *in* the marble, as a potential waiting to be actualized. The acorn is potentially an oak. The block is potentially David. The sculptor does not remember the Form — he actualizes the potential. The block's narrowness did not prevent David. It was the specific constraint that determined *which* David could emerge. A wider block would have produced a different statue, a different actualization. Without the flaw, there is no *this* David. The constraints are not obstacles to the form. They are the form's boundary conditions.

> Plato said the Form exists in the mind. Aristotle said the potential exists in the matter. Michelangelo proved both were right. The mind sees. The matter constrains. The chisel resolves.

Fred Brooks published *The Design of Design* in 2010, thirty-five years after *The Mythical Man-Month*. He spent those decades learning what Michelangelo knew at twenty-six. Constraints are not the enemy. They are the precondition. A blank page is paralysis — infinite possibility, no reason to choose. A constrained page is freedom — the boundaries tell you what you are looking for. Great design is empirical, not rational. The rational model — specify everything, then build — assumes you can see the Form without touching the stone. You cannot. Not for anything worth building. The alternative is the sculptor's rhythm. Strike. Step back. Assess. Strike again. Brooks endorsed Boehm's spiral model, but Boehm was describing what Michelangelo practiced. The design does not emerge from the plan. It emerges from the removal.

> The rational model requires omniscience. The empirical model requires honesty. Only one of these is available to humans.

Brooks defined *conceptual integrity* as the most important quality of a system — one design voice, every part consistent with every other. It is the software equivalent of the Platonic Form: the ideal system that exists in the designer's mind before a line of code is written. How do you achieve it? Brooks gave three principles, all negative, all chisels. **Propriety**: do not introduce what is immaterial. **Orthogonality**: do not link what is independent. **Generality**: do not restrict what is inherent. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks gave us three chisels. Most of us use them to build. He meant them to take away. The Form is revealed by subtraction, not addition.

Herbert Simon, in *The Sciences of the Artificial* (1969), gave this a third philosophical frame: design is *search*. The design space is every possible form the marble could take, every possible architecture the code could have. The space is infinite — until you constrain it. The block's narrowness eliminated nine-tenths of what Michelangelo might have imagined. The rough cuts eliminated more. The veins eliminated more. What remained was not a ruined block. It was a sharply constrained corridor through the infinite. Simon's term for when you stop is *satisficing* — you halt not when the design is perfect, but when the cost of further search exceeds the expected gain. Constraints make satisficing possible. Without them, you search forever.

Michelangelo could not explain how he knew where David was in the stone. He just knew. This is what Michael Polanyi called *tacit knowledge* — we know more than we can tell. A master sculptor cannot write a manual for seeing the statue in the block. A master designer cannot write a manual for seeing the architecture in the problem. The knowledge exists in the hands, in the eye, in the accumulated judgment of a lifetime. It can be transmitted only by practice — by striking the stone, assessing the result, striking again. Brooks understood this. His chapter on style is an argument that style is tacit knowledge made visible. A designer without style tries everything because they have no tacit knowledge to guide them. A designer with style tries the right thing first because their hands know what their mouth cannot say.

> Michelangelo could not tell you how he saw David. He could only show you. Style is the difference between knowing and telling. It is acquired the way Michelangelo acquired it: by cutting enough stone that the stone began to speak back.

## Why the loop?

Why does the loop exist at all? Why can't the designer see the Form, specify it completely, and hand the specification to a builder? Plato thought they could. The philosopher, after sufficient training in mathematics and dialectic, could perceive the Forms directly — no iteration required. The *Republic* describes the ideal ruler as one who has seen the Form of the Good and can therefore design the just city without trial and error. This is the rational model of design, and it is the oldest fantasy in Western philosophy.

It fails for one reason. We are not disembodied minds contemplating eternal Forms. We are embodied minds interacting with resistant matter. We cannot perceive the Form directly. We can only perceive the specific — this stone, this cut, this surface, this code. The Form is not something we see. It is something we infer, gradually, by acting on matter and observing how matter responds. Each strike of the chisel is a question asked of the stone. Each assessment is the stone's answer. The loop is not a method we chose. It is a constraint we are subject to. We iterate because we cannot do otherwise.

> The loop exists because we are embodied. We cannot see the Form. We can only ask the stone. Strike. Listen. Strike again. The loop is the cost of having a body.

This is why Aristotle broke with Plato. The Form is not in a separate realm. It is *in* the matter, as potential. You access it not by pure reason but by interacting with the thing. You carve the block to discover what it can become. You build the prototype to discover what the system should be. The interaction is not a detour on the way to the design. The interaction *is* the design process. Aristotle's word for this was *energeia* — actuality, the state of being-at-work. The Form is not perceived and then executed. It is discovered *through* execution. The loop is the mechanism of discovery.

> Plato's designer sees the Form and builds. Aristotle's designer builds and discovers the Form. The history of software engineering is the slow, painful victory of Aristotle over Plato.

This is the philosophical foundation for why art and engineering are not separate disciplines. They are two moments in one loop because the loop itself is the only way an embodied mind can access form. In the first moment, you act on the system — build, measure, cut. You are asking the matter a question. In the second moment, you step back and judge — closer to the Form or further away? You are listening to the answer. The two moments alternate because matter does not speak until it is struck, and you cannot strike intelligently until you have listened.

> The loop is not an aesthetic preference. It is an epistemic necessity. We iterate because we cannot perceive form except through its interaction with matter. Strike. Assess. Strike. This is what it means to design with a body.

Michelangelo's genius was not in either moment. It was in the fluency with which he moved between them — the speed of the loop. He struck. He listened. He struck again. The pause between action and assessment collapsed until the two became indistinguishable. Brooks called that fluency *style*. It is the closest an embodied mind can come to seeing the Form directly. The loop tightens with practice. The designer who has iterated enough no longer experiences the two moments as separate. They become one motion: strike-as-assessment, assessment-as-strike. The chisel knows where to go before the mind can say why.

A well-designed system has a quality no metric captures. A clean interface. A coherent abstraction. A module that does one thing completely. These are aesthetic properties — the residue of the loop, accumulated across enough iterations that the system feels inevitable. Trust is what you feel when a system's parts agree, when it does what it says and nothing else, when someone has struck and listened and struck again until nothing remains that is not it. Trust is the aesthetic response to conceptual integrity. Conceptual integrity is what the loop produces when both moments are disciplined. And the loop works because we are embodied — because we cannot see the Form except by asking the stone.

Brooks insisted that great design comes from "one mind, or from a very small number of agreeing resonant minds." He was not talking about intelligence. He was talking about fluency in the loop — the designer who has alternated enough times develops an eye. They see the system before it is built. They know which feature is immaterial. They know because they have asked the stone enough times, in enough blocks of marble, in enough codebases, that the stone answers before the question is fully formed.

> Every good engineer moves through the loop whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The marble releases what it was always holding. This is what it means to design with a body.

## Open questions

**Is style teachable?** Brooks says yes — study others, practice, revise. Polanyi says no — tacit knowledge cannot be transmitted, only acquired through experience. Michelangelo slept beside the marble for two years. Can you teach that in a workshop?

**Does the loop scale?** Brooks says design must come from one mind or a very small number. But most software is built by teams. Does conceptual integrity dilute with every additional mind? Or can a team learn to move through the loop together — one mind made of many?

**Can a machine develop style?** An AI coding agent can study every repository ever written. It can practice instantly. But tacit knowledge requires a body — hands that strike the stone, eyes that assess the surface. Can style exist without a body? Can taste?

**Is software too mutable for the loop?** Michelangelo's block had irrecoverable constraints. Software has none — you can always rewrite. When everything can be changed, does anything feel like a constraint? And without constraints, where does the Form come from?

**What is the Form of software?** Plato's Forms are eternal and unchanging. Software is neither. It rots. It accumulates debt. The design that satisficed last year is today's liability. Is there a Form that survives decay? Or is the loop never done — just paused, waiting for the next iteration?

---

**References:**

- Plato. *Theory of Forms*. *The Republic*, Books VI–VII.
- Aristotle. *Potentiality and Actuality*. *Metaphysics*, Book IX.
- Michael Polanyi. (1966). *The Tacit Dimension*. University of Chicago Press.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

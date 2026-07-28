---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Michelangelo didn't carve David. He removed everything that wasn't David. Plato called this remembering a Form. Aristotle called it actualizing a potential. Polanyi called it tacit knowledge. Brooks called it design. All four were describing the same loop.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

The block had been waiting thirty-five years.

It was quarried in 1464 and given to Agostino di Duccio, who worked it for two years before declaring it impossible. Too narrow. He roughed out the legs and quit. Antonio Rossellino examined it a decade later — the gouges, the veins — and refused. The block was hauled into the courtyard of the Florence Cathedral and left there. Twenty-five years of rain and sun. The Operai called it *lo gigante* — the giant. Too large to move. Too narrow to use. Too compromised to trust.

In 1501, with the block still occupying the courtyard and the Operai still paying storage, they gave it to a twenty-six-year-old who had recently returned from Rome. Michelangelo Buonarroti. He had been carving since thirteen. He had dissected cadavers at Santo Spirito to understand how muscle lay over bone. He built a wooden shed around the block so no one could watch. He slept beside the marble. For two years he struck the stone in secret. When the shed came down, David stood where a ruined block had been.

> Michelangelo did not carve David. He removed everything that was not David. The form was already in the stone. His job was to set it free.

![Michelangelo's David — carved from a block too narrow, too tall, abandoned for 35 years. The flaws did not prevent the masterpiece. They defined it.](/images/david-michelangelo.jpg)

Plato would have understood. In the *Theory of Forms*, true reality is not physical but ideal — perfect, eternal, unchanging. The David in the gallery is a shadow of the Form of David. Michelangelo did not invent him. He remembered him. The sculptor's eye, trained by years of looking, perceived the Form through the flawed stone. The chisel removed the matter that obscured it.

Aristotle, Plato's student, corrected his teacher. The Form is not elsewhere. It is *in* the marble, as potential. The acorn is potentially an oak. The block is potentially David. The sculptor actualizes what the matter already contains. The block's narrowness did not prevent David — it determined *which* David. A wider block, a different statue. Without the flaw, no *this* David. Constraints are not obstacles. They are the form's boundary conditions.

> Plato: the Form lives in the mind. Aristotle: the potential lives in the matter. Michelangelo proved both were right. The mind sees. The matter constrains. The chisel resolves.

Fred Brooks published *The Design of Design* in 2010, thirty-five years after *The Mythical Man-Month*. He spent those decades learning what Michelangelo knew at twenty-six. Constraints are not the enemy. They are the precondition. A blank page is paralysis. A constrained page is freedom — the boundaries tell you what to look for. Design is empirical, not rational. The rational model assumes you can see the Form without touching the stone. You cannot. Not for anything worth building. The alternative is the sculptor's rhythm: strike, step back, assess, strike again. Brooks endorsed Boehm's spiral, but Boehm was describing what Michelangelo practiced. The design emerges from the removal.

> The rational model requires omniscience. The empirical model requires honesty. Only one is humanly available.

Brooks defined *conceptual integrity* as the most important quality of a system — one design voice, every part consistent with every other. It is the Platonic Form of software. How do you achieve it? Brooks gave three chisels. **Propriety**: do not introduce what is immaterial. **Orthogonality**: do not link what is independent. **Generality**: do not restrict what is inherent. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> Brooks gave us three chisels. Most of us use them to build. He meant them to take away. The Form is revealed by subtraction.

Herbert Simon, in *The Sciences of the Artificial*, formalized design as *search*. The design space is infinite — every possible form, every possible architecture — until you constrain it. The block's narrowness eliminated nine-tenths of what Michelangelo might have imagined. The rough cuts eliminated more. The veins eliminated more. What remained was not ruin. It was a corridor. Simon called the stopping condition *satisficing*: you halt not at perfection, but when further search costs more than expected gain. Constraints make satisficing possible. Without them, you search forever.

Michelangelo could not explain how he knew where David was. He just knew. Polanyi called this *tacit knowledge*: we know more than we can tell. A master cannot write a manual for seeing the statue in the block. The knowledge lives in the hands, in the eye, accumulated across a lifetime of striking and assessing. Brooks' chapter on style is an argument that style is tacit knowledge made visible. A designer without style tries everything — no tacit knowledge to prune the space. A designer with style tries the right thing first — their hands know what their mouth cannot say. Michelangelo acquired his by cutting enough stone that the stone began to speak back.

> Michelangelo could not tell you how he saw David. He could only show you. Style is what you know that you cannot say. It is acquired the way he acquired it: by asking the stone until it answered.

## Why the loop?

Why can't the designer see the Form, specify it, and hand the plan to a builder? Plato thought they could — the philosopher, trained in dialectic, perceives the Forms directly. The *Republic* describes the ideal ruler who has seen the Form of the Good and designs the just city without trial and error. This is the rational model, and it is the oldest fantasy in Western thought.

It fails for one reason. We are not disembodied minds. We are minds embedded in matter, interacting with matter. We cannot perceive the Form directly — only the specific. This stone. This cut. This code. The Form is not something we see. It is something we *infer*, gradually, by acting on matter and observing how it responds. Each strike asks the stone a question. Each assessment hears the answer. The loop is not a method we chose. It is a constraint we are subject to. We iterate because we cannot do otherwise.

> The loop exists because we have bodies. We cannot see the Form. We can only ask the stone. Strike. Listen. Strike again. This is what it means to design in the flesh.

Aristotle broke with Plato here. The Form is in the matter as potential. You access it not by reason alone but by interaction. You carve to discover what the block can become. You prototype to discover what the system should be. The interaction *is* the design. Aristotle's word: *energeia* — actuality, being-at-work. The Form is not perceived then executed. It is discovered *through* execution.

> Plato's designer sees, then builds. Aristotle's designer builds, then sees. The history of software engineering is Aristotle's slow victory over Plato.

Art and engineering are not two disciplines. They are two moments in the one discipline that embodied minds can practice. In the first, you act — build, measure, cut. You ask the matter. In the second, you step back — closer to the Form or further? You hear the answer. The moments alternate because matter does not speak until struck, and you cannot strike intelligently until you have listened. Engineering without the art moment is accretion. Art without the engineering moment is fantasy. The loop is not optional. It is the only access we have to form.

Michelangelo's genius was not in either moment. It was in the speed of the loop. He struck. He listened. He struck again. The pause collapsed until the two became one motion — strike-as-assessment, assessment-as-strike. Brooks called that fluency *style*. It is what the loop produces when practiced for a lifetime. The chisel knows where to go before the mind can say why. The designer who has asked the stone enough times, in enough marble, in enough code, hears the answer before the question is fully formed.

> Every good engineer moves through the loop whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The marble releases what it was always holding.

## Open questions

**Is style teachable?** Brooks says yes. Polanyi says no — tacit knowledge cannot be transmitted, only acquired. Michelangelo slept beside the marble for two years. Can you teach that in a workshop? Or does the loop require obsession?

**Does the loop scale?** Brooks says design must proceed from one mind or very few. Most software is built by teams. Does conceptual integrity dilute with each additional mind? Can a team become one mind — moving through the loop together?

**Can a machine develop style?** An agent can study every repository ever written. But tacit knowledge requires a body — hands that strike, eyes that assess. Can style exist without flesh? Can taste?

**Is software too mutable for constraints?** Michelangelo's block had irrecoverable limits. Software has none — you can always rewrite. When everything can change, does anything feel like a constraint? And without constraints, how does the Form emerge?

**What is the Form of software?** Plato's Forms are eternal. Software rots. The design that satisficed last year is today's debt. Is the loop ever done? Or just paused — waiting for the next strike?

---

**References:**

- Plato. *Theory of Forms*. *The Republic*, Books VI–VII.
- Aristotle. *Potentiality and Actuality*. *Metaphysics*, Book IX.
- Michael Polanyi. (1966). *The Tacit Dimension*. University of Chicago Press.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

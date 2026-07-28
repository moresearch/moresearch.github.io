---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Finding David in the marble is a search problem. So is finding the design in the code. Plato called it remembering a Form. Aristotle called it actualizing a potential. Simon called it satisficing. Brooks called it iteration. All of them were describing the same loop — the one every software engineer learns to move through, or burns out fighting.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

You begin your career believing in the plan. You learn to gather requirements. You learn to draw architecture diagrams with clean boxes and confident arrows. You learn that a well-run project proceeds from specification to implementation to verification — each phase complete before the next begins. This is what you were taught. This is what the textbooks say. This is the rational model of software engineering, and it will take you years to unlearn it.

> Every engineer begins as a Platonist. We believe the Form exists in the mind, perfect and complete, and that building is merely the translation of thought into code. We are wrong, but it takes the stone to teach us.

The marble had been waiting thirty-five years. It was quarried in 1464 for Agostino di Duccio, who worked it for two years before declaring it impossible — too narrow for any figure that could stand. He roughed out the legs and quit. Rossellino examined it and refused. The block sat in the courtyard of the Florence Cathedral for a quarter century. Rain. Sun. The Operai called it *lo gigante* — the giant. Too large to move. Too narrow to use. Too compromised to trust. In 1501 they gave it to a twenty-six-year-old who built a wooden shed around it, slept beside it, and struck it in secret for two years. When the shed came down, David stood where a ruined block had been.

![Michelangelo's David — carved from a block too narrow, too tall, abandoned for 35 years. The flaws did not prevent the masterpiece. They defined it.](/images/david-michelangelo.jpg)

Plato would have recognized your younger self. In the *Theory of Forms*, the true reality of a thing is not its physical instantiation but its ideal version — perfect, eternal, accessible to the trained mind through reason alone. The philosopher, after sufficient study of mathematics and dialectic, can perceive the Forms directly. The *Republic* describes the ideal ruler as one who has seen the Form of the Good and can therefore design the just city without trial and error. No prototypes. No iterations. No listening to the stone. The plan is correct because the mind that produced it has touched the eternal. You believed this. You drew the architecture diagram and thought you had seen the system.

> Plato's designer perceives the Form and specifies it. The builder builds. This is the rational model. It is the oldest fantasy in Western thought, and it is the first thing the stone disproves.

Then you build. The architecture diagram meets the database. The clean boxes meet the edge cases. The confident arrows meet the legacy module no one mentioned. The system that was perfect in your mind emerges as something else — something that works, mostly, but doesn't feel like what you imagined. You cannot say exactly what is wrong. You only know it is not what you meant.

You have just discovered you are an embodied mind. You cannot perceive the Form directly — only the specific. This function. This latency. This error. The Form you thought you saw was a projection, a hallucination born of insufficient contact with the matter. The real Form is not in your mind. It is in the problem, as potential. You cannot think your way to it. You can only build, observe what is wrong, remove what is wrong, and build again.

> The stone is the first honest teacher. It does not care about your plan. It does not respect your diagram. It has veins and flaws and existing cuts that you did not put there. It will tell you what it can become. Your job is to listen.

Aristotle broke with Plato here. The Form is not in a separate realm of ideas. It is *in* the matter, as potential. The acorn is potentially an oak. The block is potentially David. The problem is potentially a system. The designer does not remember the Form. The designer *actualizes* the potential — by interacting with the thing, by carving to discover what the block can become, by prototyping to discover what the system should be. The interaction is not a detour on the way to the design. The interaction *is* the design. Aristotle's word was *energeia* — actuality, the state of being-at-work. The Form is not perceived and then executed. It is discovered through execution.

The block's narrowness didn't prevent David. It determined *which* David. The veins Rossellino feared became the contours of David's torso. Di Duccio's gouges became the space between his legs. Constraints are the form's boundary conditions. The legacy code is not an obstacle. It is the block — its constraints determine which system can emerge.

> Plato: the Form lives in the mind. Aristotle: the potential lives in the matter. The designer who has built enough systems knows both are right. The mind sees what the matter could become. The matter constrains what the mind can imagine. The chisel resolves.

Fred Brooks spent thirty-five years between *The Mythical Man-Month* and *The Design of Design* learning what every engineer learns: design is empirical, not rational. The rational model assumes omniscience. You don't have it. The alternative is the sculptor's rhythm — strike, step back, assess, strike again. Brooks endorsed Boehm's spiral, but Boehm was describing Michelangelo. The design emerges from the removal.

Brooks defined *conceptual integrity* as the Platonic Form of software — one design voice, every part consistent. How do you achieve it now that you know the Form was a hallucination? Three chisels. **Propriety**: do not introduce what is immaterial. **Orthogonality**: do not link what is independent. **Generality**: do not restrict what is inherent. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> The rational model adds until the spec is complete. The empirical model removes until the design is clear. The first produces systems that are full. The second produces systems that are finished.

Herbert Simon, in *The Sciences of the Artificial*, formalized this: design is *search*. The design space is infinite — every possible architecture — until constrained. The block's narrowness eliminated nine-tenths of what Michelangelo might have imagined. The legacy module eliminated half your architectures. The budget eliminated more. What remains is not restriction. It is direction. Simon called the stopping condition *satisficing* — halt not at perfection, but when further search costs more than expected gain. With experience, you recognize good enough sooner, because you've searched enough spaces to know what it looks like.

> The junior engineer searches the entire space. The senior engineer knows which regions are worth searching. The master enters the space already facing the right direction. That direction is style.

Michelangelo could not explain how he knew where David was. He just knew. Polanyi called this *tacit knowledge* — we know more than we can tell. A master cannot write a manual for seeing the architecture in the problem. The knowledge lives in the hands that typed enough code, the eye that read enough systems, the gut that tightens at the wrong abstraction. It is acquired one way: build, assess, remove, repeat. For years. No shortcut. No curriculum. Only the stone and the chisel and the willingness to sleep beside the marble until it speaks.

> Polanyi: we know more than we can tell. Brooks: style is tacit knowledge made visible. Michelangelo: the stone will tell you what it wants to become, but only if you strike it long enough to earn its trust.

## The loop

You are ten years into your career. You no longer believe in the plan. You have made peace with the stone. You have learned to strike, assess, strike again. And you have begun to notice something. The pause is shrinking. You no longer build a prototype, evaluate it, and decide what to change. You change as you build. Your fingers know this abstraction is wrong before your mind can say why. You see the system before it is written — not as a hallucination, but as a Form emerging *from* the problem, shaped by its constraints.

This is the loop. Why does it exist? Why can't you see the design, specify it, and be done?

Because you are an embodied mind. You do not have direct access to the Form. You have eyes that see this function, this latency, this error. You have hands that type this code, refactor this module, delete this feature. You have a mind that infers the general from the specific — but only *after* the specific has been produced. You cannot think the system into existence. You can only build something, look at it, and know whether it is closer to what it should be or further away. The building produces the specific. The looking produces the judgment. The judgment guides the next building. This is the loop. It is not a method. It is the epistemic condition of having a body.

> The loop exists because you cannot see the Form. You can only ask the stone. Strike. Listen. Strike again. Matter does not speak until struck. You cannot strike intelligently until you have listened. This is what it means to design in the flesh.

Art and engineering are not two disciplines. They are two moments in the one discipline available to embodied minds. First moment: act on the system — write, ship, cut. You produce the specific without which no judgment is possible. Engineering. Second moment: step back — is this closer to the Form or further? You hear what the matter tells you about the general. Art. Neither moment produces the design alone. Engineering without art is accretion — systems that grow without shape, the career your colleagues settled into. Art without engineering is fantasy — diagrams that never meet the database, the career you might have had. The loop requires both because knowledge requires both. You cannot judge what doesn't exist. You cannot improve what you don't judge.

> The loop is not a preference. It is an epistemic necessity. The engineer who never judges builds ruins. The artist who never builds produces fantasies. The loop is what separates a career from a vocation.

Michelangelo's genius was the speed of his loop. He struck and assessed so rapidly the two moments became one motion. Brooks called that fluency *style* — what the loop produces when practiced for a lifetime. The chisel knows where to go before the mind can say why. The engineer who has asked the codebase enough questions, across enough years, hears the answer before the question is formed. They see the system before it's built. They recognize David in the ruined block. Not because they're geniuses. Because the loop has become who they are.

> Every engineer moves through the loop whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The marble releases what it was always holding. This is what it means to design with a body. This is what it means to become an engineer.

## Open questions

**Is style teachable?** Brooks says yes — study, practice, revise. Polanyi says no — tacit knowledge can only be acquired, not transmitted. Michelangelo slept beside the marble for two years. Is there a curriculum for obsession?

**Does the loop scale?** Brooks says design must come from one mind or very few. Most software is built by teams. Can a team become one mind — searching the design space together? Or does conceptual integrity dilute with every additional voice?

**Can a machine enter the loop?** An agent can study every repository instantly. But tacit knowledge requires a body — hands that strike, eyes that assess, a gut that tightens. Can style exist without flesh? Can taste?

**Is software too unconstrained to search?** Michelangelo's block had irrecoverable limits that narrowed the search. Software has none — you can always rewrite. When everything can change, does anything constrain the search? Without constraints, does the design space collapse into indifference?

**What is the Form of software?** Plato's Forms are eternal and unchanging. Software rots. The design that satisficed last year is today's debt. Is the loop ever done? Or just paused — the chisel resting, the stone waiting for the next strike?

---

**References:**

- Plato. *Theory of Forms*. *The Republic*, Books VI–VII.
- Aristotle. *Potentiality and Actuality*. *Metaphysics*, Book IX.
- Michael Polanyi. (1966). *The Tacit Dimension*. University of Chicago Press.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

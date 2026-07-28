---
title: Finding David in the Marble
date: 2026-07-28
slug: finding-david-in-the-marble
summary: Every software engineer begins believing in the rational model — plan everything, then build. The marble teaches otherwise. Michelangelo, Plato, Aristotle, Polanyi, and Brooks all discovered the same loop. This is the story of that discovery.
tags: software-engineering, design-of-design, iteration, fred-brooks, conceptual-integrity, design-philosophy
---

You begin your career believing in the plan. You learn to gather requirements. You learn to draw architecture diagrams with clean boxes and confident arrows. You learn that a well-run project proceeds from specification to implementation to verification — each phase complete before the next begins. This is what you were taught. This is what the textbooks say. This is the rational model of software engineering, and it will take you years to unlearn it.

> Every engineer begins as a Platonist. We believe the Form exists in the mind, perfect and complete, and that building is merely the translation of thought into code. We are wrong, but it takes the stone to teach us.

The marble had been waiting thirty-five years. It was quarried in 1464 for Agostino di Duccio, who worked it for two years before declaring it impossible — too narrow for any figure that could stand. He roughed out the legs and quit. Rossellino examined it and refused. The block sat in the courtyard of the Florence Cathedral for a quarter century. Rain. Sun. The Operai called it *lo gigante* — the giant. Too large to move. Too narrow to use. Too compromised to trust. In 1501 they gave it to a twenty-six-year-old who built a wooden shed around it, slept beside it, and struck it in secret for two years. When the shed came down, David stood where a ruined block had been.

![Michelangelo's David — carved from a block too narrow, too tall, abandoned for 35 years. The flaws did not prevent the masterpiece. They defined it.](/images/david-michelangelo.jpg)

Plato would have recognized your younger self. In the *Theory of Forms*, the true reality of a thing is not its physical instantiation but its ideal version — perfect, eternal, accessible to the trained mind through reason alone. The philosopher, after sufficient study of mathematics and dialectic, can perceive the Forms directly. The *Republic* describes the ideal ruler as one who has seen the Form of the Good and can therefore design the just city without trial and error. No prototypes. No iterations. No listening to the stone. The plan is correct because the mind that produced it has touched the eternal. You believed this. You drew the architecture diagram and thought you had seen the system.

> Plato's designer perceives the Form and specifies it. The builder builds. This is the rational model. It is the oldest fantasy in Western thought, and it is the first thing the stone disproves.

Then you build. The architecture diagram meets the database. The clean boxes meet the edge cases. The confident arrows meet the legacy module that no one mentioned in the requirements meeting. The system that was perfect in your mind emerges from the code as something else — something that works, mostly, but doesn't feel like what you imagined. You cannot say exactly what is wrong. You only know it is not what you meant.

This is the moment you discover that you are an embodied mind. You cannot perceive the Form directly. You can only perceive the specific — this function, this latency, this error message, this user complaint. The Form is not something you saw. It was something you *projected*, a hallucination generated from insufficient interaction with the matter. The real Form — what the system should actually be — is not in your mind. It is in the problem, as a potential waiting to be discovered. You cannot reach it by thinking harder. You can only reach it by building something, observing what is wrong with it, removing what is wrong, and building again.

> The stone is the first honest teacher. It does not care about your plan. It does not respect your diagram. It has veins and flaws and existing cuts that you did not put there. It will tell you what it can become. Your job is to listen.

Aristotle broke with Plato here. The Form is not in a separate realm of ideas. It is *in* the matter, as potential. The acorn is potentially an oak. The block is potentially David. The problem is potentially a system. The designer does not remember the Form. The designer *actualizes* the potential — by interacting with the thing, by carving to discover what the block can become, by prototyping to discover what the system should be. The interaction is not a detour on the way to the design. The interaction *is* the design. Aristotle's word was *energeia* — actuality, the state of being-at-work. The Form is not perceived and then executed. It is discovered through execution.

The block's narrowness did not prevent David. It determined *which* David. A wider block would have produced a different statue — a different actualization of the same potential. The veins that Rossellino saw as fatal flaws became the contours of David's torso. The gouges from di Duccio's abandoned attempt became the space between David's legs. The constraints were not obstacles. They were the form's boundary conditions. In software, the legacy code is not an obstacle to the design. It is the block. Its constraints are not preventing the system from being what it should be. They are determining *which* system can emerge.

> Plato: the Form lives in the mind. Aristotle: the potential lives in the matter. The designer who has built enough systems knows both are right. The mind sees what the matter could become. The matter constrains what the mind can imagine. The chisel resolves.

Fred Brooks published *The Design of Design* in 2010, thirty-five years after *The Mythical Man-Month*. He spent those decades learning what every engineer learns: design is empirical, not rational. The rational model assumes omniscience. You do not have it. The alternative is the sculptor's rhythm. Strike. Step back. Assess. Strike again. Brooks endorsed Boehm's spiral model, but Boehm was describing what Michelangelo practiced. The design does not emerge from the plan. It emerges from the removal.

Brooks defined *conceptual integrity* as the most important quality of a system — one design voice, every part consistent with every other. It is the Platonic Form of software: the ideal system that haunted your first architecture diagram, the one that was perfect until it touched the stone. How do you achieve it now that you know the Form was a hallucination? Brooks gave three principles, all negative, all chisels. **Propriety**: do not introduce what is immaterial. **Orthogonality**: do not link what is independent. **Generality**: do not restrict what is inherent. Propriety removes the unnecessary. Orthogonality removes the tangled. Generality removes the arbitrary. The design becomes what it should be by becoming less of what it shouldn't.

> The rational model adds until the spec is complete. The empirical model removes until the design is clear. The first produces systems that are full. The second produces systems that are finished.

Herbert Simon, in *The Sciences of the Artificial*, gave this a third frame: design is *search*. The design space is infinite — every possible architecture, every possible abstraction — until you constrain it. The block's narrowness eliminated nine-tenths of what Michelangelo might have imagined. The legacy module eliminated half the architectures you might have drawn. The budget eliminated more. What remains is not restriction. It is direction. Simon called the stopping condition *satisficing* — you halt not at perfection, but when further search costs more than expected improvement. You stop when the system is good enough. The older you get, the earlier you recognize good enough, because you have searched enough spaces to know what good enough looks like.

> The junior engineer searches the entire space. The senior engineer knows which regions are worth searching. The master enters the space already facing the right direction. That direction is style.

Which brings you to the hardest lesson. Michelangelo could not explain how he knew where David was in the stone. He just knew. This is what Michael Polanyi called *tacit knowledge* — we know more than we can tell. A master cannot write a manual for seeing the architecture in the problem. The knowledge lives in the hands that have typed enough code, the eye that has read enough systems, the gut that tightens at the wrong abstraction. It is acquired only one way: by building, assessing, removing, and building again. For years. There is no shortcut. There is no curriculum. There is only the stone and the chisel and the willingness to sleep beside the marble until it speaks.

> Polanyi: we know more than we can tell. Brooks: style is tacit knowledge made visible. Michelangelo: the stone will tell you what it wants to become, but only if you strike it long enough to earn its trust.

## The loop

You are ten years into your career. You no longer believe in the plan. You have made peace with the stone. You have learned to strike, assess, strike again. And you have begun to notice something. The pause is shrinking. You no longer build a prototype, evaluate it, and decide what to change. You change as you build. Your fingers know this abstraction is wrong before your mind can say why. You see the system before it is written — not as a hallucination, but as a Form emerging *from* the problem, shaped by its constraints.

This is the loop. Why does it exist? Why can't you see the design, specify it, and be done?

Because you are an embodied mind. You do not have direct access to the Form. You have eyes that see this function, this latency, this error. You have hands that type this code, refactor this module, delete this feature. You have a mind that infers the general from the specific — but only *after* the specific has been produced. You cannot think the system into existence. You can only build something, look at it, and know whether it is closer to what it should be or further away. The building produces the specific. The looking produces the judgment. The judgment guides the next building. This is the loop. It is not a method. It is the epistemic condition of having a body.

> The loop exists because you cannot see the Form. You can only ask the stone. Strike. Listen. Strike again. Matter does not speak until struck. You cannot strike intelligently until you have listened. This is what it means to design in the flesh.

This is why art and engineering are not two disciplines. They are two moments in the one discipline available to embodied minds. In the first moment, you act — you write the function, you ship the feature, you strike the stone. You are producing the specific without which no judgment is possible. Call this engineering. In the second moment, you step back — is this closer to the Form or further away? You are listening to what the matter tells you about the general. Call this art. Neither moment produces the design alone. Engineering without the art moment produces accretion: systems that grow without shape, features added and never removed, the career you watched your colleagues settle into. Art without the engineering moment produces fantasy: diagrams that never meet the database, the career you might have had if you never built anything. The loop requires both because knowledge requires both. You cannot judge what does not exist. You cannot improve what you do not judge.

> The loop is not a preference. It is an epistemic necessity. The engineer who never judges is a builder of ruins. The artist who never builds is a producer of fantasies. The loop is what separates a career from a vocation.

> Art and engineering are not two disciplines. They are two moments in one discipline. The engineer who never steps back to judge is a builder of ruins. The artist who never strikes the stone is a producer of fantasies. The loop is not optional. It is what separates a career from a vocation.

Michelangelo's genius was in the speed of the loop. He struck and assessed so rapidly that the two moments became one motion — strike-as-assessment, assessment-as-strike. Brooks called that fluency *style*. It is what the loop produces when practiced for a lifetime. The chisel knows where to go before the mind can say why. The engineer who has asked the codebase enough questions, in enough codebases, across enough years, hears the answer before the question is fully formed. They see the system before it is built. They know which feature is immaterial. They recognize David in the ruined block. Not because they are geniuses. Because they have done the loop enough times that the loop has become who they are.

> Every engineer moves through the loop whether they know it or not. The ones who know it move faster. They strike with intention. They assess with honesty. They recognize David earlier. They remove less because they add less. The loop tightens. The marble releases what it was always holding. This is what it means to design with a body. This is what it means to become an engineer.

## Open questions

**Is style teachable?** Brooks says yes. Polanyi says no — tacit knowledge cannot be transmitted, only acquired. Michelangelo slept beside the marble for two years. Is there a curriculum for that? Or does the loop require an obsession that cannot be taught?

**Does the loop scale?** Brooks says design must come from one mind or very few. Most software is built by teams. Can a team become one mind — moving through the loop together? Or does conceptual integrity dilute with every additional voice?

**Can a machine enter the loop?** An agent can study every repository. It can practice instantly. But tacit knowledge requires a body — hands that strike, eyes that assess, a gut that tightens at the wrong abstraction. Can style exist without flesh?

**Is software too mutable for the loop to matter?** Michelangelo's block had irrecoverable constraints. Software has none — you can always rewrite. When everything can change, does anything feel like a constraint? And without constraints, how does the Form emerge? Does infinite mutability produce infinite mediocrity?

**What is the Form of software?** Plato's Forms are eternal. Software rots. The design that satisficed last year is today's debt. Is the loop ever done? Or is it just paused — the chisel resting, the stone waiting for the next strike?

---

**References:**

- Plato. *Theory of Forms*. *The Republic*, Books VI–VII.
- Aristotle. *Potentiality and Actuality*. *Metaphysics*, Book IX.
- Michael Polanyi. (1966). *The Tacit Dimension*. University of Chicago Press.
- Herbert Simon. (1969). *The Sciences of the Artificial*. MIT Press.
- Fred Brooks. (2010). *The Design of Design: Essays from a Computer Scientist*. Addison-Wesley.
- [Michelangelo's David](https://en.wikipedia.org/wiki/David_(Michelangelo)). 1501–1504. Galleria dell'Accademia, Florence.
- Related: [The Principle of Least Astonishment — for AI Coding Agents](https://blog.hackspree.com/#principle-of-least-astonishment).

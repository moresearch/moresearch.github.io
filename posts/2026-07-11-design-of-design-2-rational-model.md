---
title: Brooks on Software Design Series: the waterfall is wrong
date: 2026-07-11
slug: brooks-design-rational-model
summary: "The rational model says gather requirements, design, implement, test, ship. Brooks says it is wrong and harmful. You don't know the goal at the start. Neither does your client. The Gantt chart is a lie."
tags: design, fred-brooks, waterfall, rational-model, co-evolution
---

Parts 1-3: who designs (one mind, protected). Parts 4-5: how design works. Spoiler: not the way your project plan says. The project plan was wrong before it was printed.

The dominant model is wrong. Not slightly. "Earth is flat" wrong.

## Definitions

**Rational model (waterfall).** Gather requirements. Design. Implement. Test. Ship. Each phase finishes before the next. Proceeds logically from premises to conclusion. Makes sense on a Gantt chart. Has never once worked in reality. Yet we keep using it because it makes managers feel better.

Clean. Orderly. Wrong.

> "The Waterfall Model is wrong and harmful; we must outgrow it. What is wrong is that it is an essentially rational model, and for wicked problems, the rational model is simply the wrong model."

Not "sometimes inappropriate." *Wrong and harmful.* It demands decisions at maximum ignorance — the beginning — and forbids revisiting them. It's like ordering dessert before you've seen the menu, eaten the meal, or confirmed you're hungry. Then acting surprised when nobody wants the tiramisu. Then blaming the tiramisu.

The model assumes the designer knows the goal at the start. False. Laughably false. Anyone who has built anything real knows this. Anyone who hasn't nods along with the Gantt chart.

Herbert Simon: design as systematic search. Goals, utility functions, constraints. Find the optimum. Theory is beautiful. Simon won a Nobel Prize. His theory still doesn't survive contact with a real client who changes their mind after seeing the first prototype. The client is not irrational. The model is wrong about when knowledge arrives.

Practice: **nobody knows the goal.** Not the designer. Not the client. *"That's what I asked for, but that's not what I want."* Every designer has heard this. It means the process is working, not failing. A client cannot articulate needs until they see something. They can spot what's wrong instantly. They cannot describe what's right before anything exists. This is not a communication failure. This is how cognition works. Seeing is knowing. Speculating is guessing.

**Co-evolution.** Requirements and design change each other. Designing reveals new requirements. New requirements change the design. Cycle continues until both stabilize. Or until the budget runs out. Whichever comes first. Usually the budget.

> "Requirements and design co-evolve. The act of designing changes the designer's understanding of the problem. As the design emerges, the requirements change. This is not failure; it is discovery."

Requirements are not extracted like ore. They are produced through designing. Each iteration teaches. The process ends when further iteration yields diminishing returns. In practice, it ends when the PM says "we need to ship." This is also a form of diminishing returns. Just not the one Brooks had in mind.

Peter Naur, same era. **Theory building** (1985): a program is not its code. It is the theory its builders hold of the problem. Cannot be extracted upfront. Built through designing. When the last person who understands the system leaves, the theory leaves with them. The code remains. Nobody knows why it works. This is called "legacy." It is also called "most production systems."

Brooks and Naur: the real product is understanding. Understanding emerges through the work. Documentation is not understanding. Documentation is a fossil of understanding that was alive six months ago. The fossil is useful. It is not the animal.

---

[← Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 5 →](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

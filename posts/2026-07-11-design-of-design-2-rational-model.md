---
title: "Brooks on design, part 4: the waterfall is wrong"
date: 2026-07-11
slug: brooks-design-rational-model
summary: "The rational model says gather requirements, design, implement, test, ship. Brooks says it is wrong and harmful. You don't know the goal at the start."
tags: design, fred-brooks, waterfall, rational-model, co-evolution
---

Parts 1-3: who should design (one mind, protected). Parts 4-5: how design actually works.

The dominant model is wrong.

## Definitions

**Rational model (waterfall).** Gather requirements. Design. Implement. Test. Ship. Each phase finishes before the next begins. Proceeds logically from premises to conclusion.

Clean. Orderly. Wrong.

> "The Waterfall Model is wrong and harmful; we must outgrow it. What is wrong is that it is an essentially rational model, and for wicked problems, the rational model is simply the wrong model."

Not "sometimes inappropriate." *Wrong and harmful.* Why: it demands decisions at maximum ignorance — the beginning — and forbids revisiting them.

## The fatal flaw

The model assumes the designer knows the goal at the start. False.

Herbert Simon: design as systematic search. Goals, utility functions, constraints, resources. Find the optimum. Theory is beautiful.

Practice: **nobody knows the goal.** Not the designer. Not the client. Brooks: *"That's what I asked for, but that's not what I want."* A client cannot articulate needs until they see something. They can spot what's wrong. They cannot describe what's right before anything exists. Not a communication failure. A property of design.

**Co-evolution.** Requirements and design change each other. Designing reveals new requirements. New requirements change the design. Cycle continues until both stabilize.

> "Requirements and design co-evolve. The act of designing changes the designer's understanding of the problem. As the design emerges, the requirements change. This is not failure; it is discovery."

Requirements are not extracted. They are produced through designing. Each iteration teaches. The process ends when further iteration yields diminishing returns.

Peter Naur, same era, different tradition. **Theory building** (1985): a program is not its code. It is the theory its builders hold of the problem. Cannot be extracted upfront. Built through designing. Brooks and Naur: the real product is understanding. Understanding emerges through the work.

---

**Part 4 of 7.** [← Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 5 →](https://blog.hackspree.com/#brooks-design-empiricist-alternative)

**References:** Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

---
title: "Brooks on design, part 4: the waterfall model is wrong and harmful"
date: 2026-07-11
slug: brooks-design-rational-model
summary: "The rational model says: gather requirements, design, implement, test, ship. Brooks says it is wrong and harmful. You don't know the goal at the start, and the act of designing changes what you think the goal is."
tags: design, fred-brooks, rational-model, waterfall, co-evolution, peter-naur
---

Parts 1-3 established who should design: one mind, protected by the organization. Parts 4-5 are about how design actually works.

The dominant model is wrong.

## Definition: the rational model

**Rational model (waterfall).** Gather requirements. Produce a design. Implement the design. Test. Ship. Each phase finishes before the next begins. The design proceeds logically from known premises to a correct conclusion.

It is clean. It is orderly. It is wrong.

> "The Waterfall Model is wrong and harmful; we must outgrow it. What is wrong is that it is an essentially rational model, and for wicked problems, the rational model is simply the wrong model."

Not "sometimes inappropriate." Not "useful in certain domains." *Wrong and harmful.* The model demands decisions at the point of maximum ignorance — the beginning — and forbids revisiting them.

## The fatal flaw

The rational model assumes the designer knows the goal at the start. This is false.

In Herbert Simon's formulation, design is systematic search through a combinatorial space. You have goals, utility functions, constraints, and resources. You search for the optimal solution. The theory is beautiful.

The practice: **nobody knows the goal.** Not the designer. Not the client. Brooks captures this: *"That's what I asked for, but that's not what I want."* A client cannot articulate what they need until they see something. Once they see it, they can tell you what's wrong. They cannot tell you what's right before anything exists. This is not a communication failure. It is a property of design problems.

## Definition: co-evolution

**Co-evolution.** Requirements and design change each other. The act of designing reveals new requirements. New requirements change the design. This cycle continues until both stabilize.

> "Requirements and design co-evolve. The act of designing changes the designer's understanding of the problem. As the design emerges, the requirements change. This is not failure; it is discovery."

Requirements do not sit fully formed in the client's mind, waiting for extraction. They are *produced* through designing. Each iteration teaches the client something about what they actually need. The process ends when further iteration produces diminishing returns on new understanding.

Peter Naur reached the same conclusion. **Theory building** (1985): a program is not its code. It is the theory its builders hold of the problem it solves. That theory cannot be extracted upfront. It is built through designing. Brooks's co-evolution and Naur's theory-building are the same insight from different traditions. The real product of design is understanding. Understanding emerges through the work.

The next part describes the positive alternative: if we can't reason our way to a correct design, what do we do instead?

---

**This is part 4 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](https://blog.hackspree.com/#brooks-design-conceptual-integrity)
- [Part 2: Why one mind must rule the design](https://blog.hackspree.com/#brooks-design-one-mind-rule)
- [Part 3: How to protect designers from their organizations](https://blog.hackspree.com/#brooks-design-protecting-designer)
- [Part 5: Build, test, iterate — the empiricist method](https://blog.hackspree.com/#brooks-design-empiricist-alternative)
- [Part 6: Why experts design the wrong thing beautifully](https://blog.hackspree.com/#brooks-design-experts-divorce)
- [Part 7: Great designs come from great designers — not great processes](https://blog.hackspree.com/#brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

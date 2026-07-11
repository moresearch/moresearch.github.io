---
title: "Brooks on design, part 1: conceptual integrity"
date: 2026-07-11
slug: brooks-design-conceptual-integrity
summary: "Conceptual integrity: the system feels like one mind designed it. The most important property of any designed thing."
tags: design, fred-brooks, conceptual-integrity, orthogonality
---

In 2010, Fred Brooks published *The Design of Design*. It is his best book and his least read.

The argument: **conceptual integrity is the most important property of any designed system.** Everything else serves it.

## Definition

**Conceptual integrity.** The system feels like one mind designed it. The user forms one mental model. They can predict behavior in new situations.

Brooks defined it in *The Mythical Man-Month* (1975), sharpened it in *No Silver Bullet* (1986), and devoted *The Design of Design* (2010) to it:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — *The Mythical Man-Month*, 1975

The tradeoff: omit useful features to preserve coherence. Say no to good ideas that don't fit.

Three properties:

**Orthogonality.** One way to do each thing. No overlapping concepts. Non-redundancy.

**Propriety.** Nothing unnecessary.

> "The essential skill of the designer is saying no — repeatedly, to smart people with good arguments — and having the authority to make it stick. Every added feature is a subtraction from all existing features, for it adds complexity without corresponding benefit."

Feature cost is multiplicative. Each new thing makes every existing thing harder.

**Generality.** No arbitrary limits. Primitives compose to handle unanticipated cases. Generality without orthogonality is a kitchen sink. Orthogonality without generality is a toy.

Brooks was not alone. Knuth: programs should be readable by humans. Dijkstra: elegance prevents bugs — transparent structure hides nothing. Kernighan and Plauger catalogued the difference between clarity and muddle in *The Elements of Programming Style* (1974). Wirth designed Pascal as an exercise in integrity — one mind, one language. Lampson: the Alto's GUI worked because one designer decided what to leave out.

Brooks's contribution: identifying the structural precondition. One mind must control the design. In *Computer Architecture* (1997), he and Blaauw showed this governed even ISAs, where orthogonality was the explicit goal.

## Proof

> "Reims Cathedral has conceptual integrity. Built over eight generations of architects, each stuck to the original plan. The result is a unified work, coherent in every detail. Most cathedrals are not like this; their conflicting concepts produce architectural chaos."

Reims worked because the original architect's plan outlived him. Two centuries of successors submitted to it. It should not exist. That it does is evidence.

Most cathedrals: Gothic nave, Renaissance facade, Baroque chapel on Romanesque transept. Each generation had a vision. Nobody had authority. Chaos.

Software systems are cathedrals built over decades. Unix (Thompson and Ritchie, one resonant pair). Lisp (McCarthy, one mind). Go (Thompson, Pike, Griesemer). The Macintosh (Jobs). These are Reims. Thompson: "One of my most productive days was throwing away 1,000 lines of code."

---

**Part 1 of 7.** [Part 2 →](https://blog.hackspree.com/#brooks-design-one-mind-rule)

**References:** Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

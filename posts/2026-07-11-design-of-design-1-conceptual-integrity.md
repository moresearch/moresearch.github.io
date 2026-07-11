---
title: "Brooks on design, part 1: conceptual integrity — the most important property"
date: 2026-07-11
slug: brooks-design-conceptual-integrity
summary: "Conceptual integrity is the most important consideration in system design. Everything else — process, tools, team structure, schedules — serves this property, or should."
tags: design, fred-brooks, conceptual-integrity, orthogonality, reims
---

In 2010, thirty-five years after *The Mythical Man-Month*, Fred Brooks published *The Design of Design*. It was met with less fanfare than its predecessor. That is a mistake. It is the most distilled thinking on design from someone who spent six decades designing across five media: computer architecture, software, houses, books, and organizations.

The central argument: **conceptual integrity is the most important consideration in system design**. Everything else serves this property, or should.

## Definition

Conceptual integrity means the system feels like one mind designed it. The user forms a coherent mental model. They can predict how the system will behave in situations they haven't encountered.

Brooks first defined it in *The Mythical Man-Month* (1975), sharpened it in *No Silver Bullet* (1986), then devoted *The Design of Design* (2010) to it:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — *The Mythical Man-Month*, 1975

The tradeoff: omit useful features to preserve coherence. Most designers agree in principle. Most violate it the moment someone proposes a good feature that doesn't fit. Saying yes is easier than explaining why a good idea makes the whole worse.

Conceptual integrity has three properties:

**Orthogonality. One way to do each thing.** If two paths produce the same result, the user must learn both, choose between them, and absorb the inconsistencies. Orthogonality is not minimalism. It is non-redundancy.

**Propriety. Nothing unnecessary.** Everything included must earn its place.

> "The essential skill of the designer is saying no — repeatedly, to smart people with good arguments — and having the authority to make it stick. Every added feature is a subtraction from all existing features, for it adds complexity without corresponding benefit."

The cost of a feature is not additive. It is multiplicative. Each new thing makes every existing thing harder to find, learn, and use.

**Generality. No arbitrary limits.** The primitives must compose to handle cases the designer didn't anticipate. Generality without orthogonality is a kitchen sink. Orthogonality without generality is a toy.

Brooks was not alone. Knuth argued that the best programs are readable by humans, not just executable by machines. Dijkstra insisted that elegance prevents bugs: transparent structure leaves fewer places for errors to hide. Kernighan and Plauger catalogued the decisions that distinguish clarity from muddle in *The Elements of Programming Style* (1974). Wirth designed Pascal and Modula-2 as deliberate exercises in integrity — each language the vision of one mind. Lampson observed that the Alto's GUI succeeded because one designer decided what to include and what to leave out.

The whole generation agreed integrity matters. Brooks's contribution: identifying the *structural precondition*. One mind must control the design. In *Computer Architecture: Concepts and Evolution* (1997), he and Blaauw showed that this principle governed even instruction set architectures, where orthogonality of operations, data types, and addressing modes was the explicit design goal.

## Proof: the Reims Cathedral

> "Reims Cathedral has conceptual integrity. Built over eight generations of architects, each stuck to the original plan. The result is a unified work, coherent in every detail. Most cathedrals are not like this; their conflicting concepts produce architectural chaos — individually interesting, collectively incoherent."

Reims worked because the original architect's plan had authority that outlived him. Every successor, across two centuries, submitted to it. A unified work built by many hands. It should not exist. That it does is the evidence.

Most cathedrals are the other kind: Gothic nave, Renaissance facade, Baroque chapel on a Romanesque transept. Each generation had a vision. Nobody had authority to enforce continuity. The result is chaos — individually interesting, collectively incoherent.

Software systems are cathedrals built over decades. Unix (Thompson and Ritchie, one resonant pair). Lisp (McCarthy, one mind). Go (Thompson, Pike, Griesemer, a tight trio). The original Macintosh (Jobs channeling a singular vision). These are Reims. Everything else is a patchwork quilt. Thompson captured the ethic: "One of my most productive days was throwing away 1,000 lines of code." Saying no to your own complexity is the act that preserves integrity.

The next part takes up the practical question: if integrity requires one mind, how do you achieve that in an organization of many minds?

---

**This is part 1 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 2: Why one mind must rule the design](https://blog.hackspree.com/#brooks-design-one-mind-rule)
- [Part 3: How to protect designers from their organizations](https://blog.hackspree.com/#brooks-design-protecting-designer)
- [Part 4: The waterfall model is wrong and harmful](https://blog.hackspree.com/#brooks-design-rational-model)
- [Part 5: Build, test, iterate — the empiricist method](https://blog.hackspree.com/#brooks-design-empiricist-alternative)
- [Part 6: Why experts design the wrong thing beautifully](https://blog.hackspree.com/#brooks-design-experts-divorce)
- [Part 7: Great designs come from great designers — not great processes](https://blog.hackspree.com/#brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

---
title: "Brooks on design, part 1: conceptual integrity and the Reims Cathedral"
date: 2026-07-11
slug: brooks-design-conceptual-integrity
summary: "Fred Brooks' 'The Design of Design' argues that the most important property of any system is conceptual integrity — and that integrity requires a single mind, or at most two, controlling the design."
tags: design, fred-brooks, conceptual-integrity, software-architecture, collaboration
---

In 2010, thirty-five years after *The Mythical Man-Month*, Fred Brooks published *The Design of Design: Essays from a Computer Scientist*. It is a quieter book — less urgent, more reflective — and it was met with less fanfare. That is a mistake. It contains the most distilled thinking on what makes design work from someone who spent six decades designing across five media: computer architecture, software, houses, books, and organizations.

The central argument is that **conceptual integrity is the most important consideration in system design**. Everything else — process, tools, team structure, schedules — serves this property, or should.

## What conceptual integrity is

Brooks first defined the concept in *The Mythical Man-Month* (1975), sharpened it in *No Silver Bullet* (1986) and the *Anniversary Edition* (1995), then devoted *The Design of Design* (2010) to exploring it as a universal property of all designed things:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — *The Mythical Man-Month*, 1975

The sentence is a thesis statement for his career, and it turns on a hard tradeoff: omit useful features to preserve coherence. Most designers agree with this in principle. Most violate it in practice the moment someone proposes a good feature that doesn't fit — because saying yes is easier than explaining why a good idea makes the whole worse.

Brooks decomposes conceptual integrity into three properties:

**Orthogonality.** Concepts should not overlap. Each function should be achievable in exactly one way. Two paths to the same result force the user to learn both, decide which to use, and absorb the inconsistencies between them. Orthogonality is not minimalism. It is non-redundancy. A design can be rich and still orthogonal.

**Propriety.** The design should include only what is necessary, and what is included should be transparent.

> "The essential skill of the designer is saying no — repeatedly, to smart people with good arguments — and having the authority to make it stick. Every added feature is a subtraction from all existing features, for it adds complexity without corresponding benefit."

Every feature you add makes every other feature harder to find, learn, and use. The cost is not additive. It is multiplicative. Saying no is the designer's highest-leverage act, and it requires an organizational structure that backs the no.

**Generality.** Within its chosen scope, the design should be complete — no arbitrary limits the user must work around. A design with generality handles the cases the designer anticipated and the ones they didn't, because the primitives compose. Generality without orthogonality is a kitchen sink. Orthogonality without generality is a toy.

A system with all three properties feels like it came from one mind. The user forms a coherent mental model. They can predict how the system will behave in situations they haven't yet encountered. This is not an aesthetic preference. It is the difference between a tool and an obstacle course.

Brooks was not alone in valuing coherence. Knuth, his contemporary, argued that the best programs possess a "literate" quality — readable by humans, not just executable by machines. Dijkstra insisted that elegance was a practical property: an elegant program contains fewer bugs because its structure is transparent. Kernighan and Plauger catalogued the small decisions that distinguish clarity from muddle in *The Elements of Programming Style* (1974). The whole generation agreed that integrity matters. Brooks's contribution was identifying its *structural precondition*: one mind must control the design.

## The Reims Cathedral

Brooks uses architecture to make the point undeniable.

> "Reims Cathedral has conceptual integrity. Built over eight generations of architects, each stuck to the original plan. The result is a unified work, coherent in every detail. Most cathedrals are not like this; their conflicting concepts produce architectural chaos — individually interesting, collectively incoherent."

Reims worked because the original architect's plan had authority that outlived him. Every successor — across two centuries — submitted to it. The result is a unified work built by many hands. It should not exist. That it does is evidence for Brooks's thesis.

Most cathedrals are the other kind: Gothic nave, Renaissance facade, Baroque chapel grafted onto a Romanesque transept. Each generation had a vision and nobody had the authority to enforce continuity. The result is architectural chaos — individually interesting, collectively incoherent.

Software systems are cathedrals built over decades. Unix (Ken Thompson and Dennis Ritchie, one resonant pair), Lisp (John McCarthy, one mind), Go (Thompson, Pike, Griesemer, a tight trio), the original Macintosh (Jobs channeling a singular vision) — these are Reims. Everything else is a patchwork quilt. Thompson captured the Brooksian ethic perfectly: "One of my most productive days was throwing away 1,000 lines of code." Saying no to your own complexity is the act that preserves conceptual integrity.


The next part takes up the practical question: if integrity requires one mind, how do you achieve that in an organization of many minds?

---

**This is part 1 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 2: The one-mind rule](/posts/brooks-design-one-mind-rule)
- [Part 3: Protecting the designer](/posts/brooks-design-protecting-designer)
- [Part 4: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 5: The empiricist alternative](/posts/brooks-design-empiricist-alternative)
- [Part 6: How experts go wrong and the divorce of design](/posts/brooks-design-experts-divorce)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

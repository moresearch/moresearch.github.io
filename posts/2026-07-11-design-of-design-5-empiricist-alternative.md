---
title: "build, test, iterate"
date: 2026-07-11
slug: brooks-design-empiricist-alternative
summary: "If we can't think our way to a correct design, what do we do? Build, test, and iterate. Prototyping, constraints, and user models are the tools."
tags: design, fred-brooks, empiricism, prototyping, constraints
---

The rational model is wrong. What's the alternative?

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

## Definitions

**Empiricist method.** You cannot think your way to a correct design. Build, test, learn, iterate.

Not the Brooks of *The Mythical Man-Month*. Younger Brooks believed planning — "plan to throw one away" meant doing the rational process twice. Older Brooks: the first plan was never going to be right. No analysis would have fixed it. The only path runs through being wrong.

Contrast Wirth's **stepwise refinement** (1971): decompose, refine until trivial. Works when you already understand the problem. Brooks: you never do. Decomposition emerges through building and testing. Wirth's method: well-understood problems. Brooks's method: everything else.

Six steps: study domain → design (knowing it's wrong) → prototype → test with real users → iterate → build incrementally.

**Prototype.** A concrete version built to be tested and discarded. Not a draft. A question posed to reality.

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

Specifications are unfalsifiable. Only running code can be wrong in a way that teaches.

**Formal methods.** Proving programs correct by deduction. Works for small modules. Cannot scale.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

The rationalist dream — correct by construction — survives only in CS departments. Every other discipline abandoned it centuries ago. *No Silver Bullet* (1986): no breakthrough eliminates essential difficulty. *Computer Architecture* (1997): even ISAs evolved through trial and error. Hamming: "The purpose of computing is insight, not numbers." Designing is insight, not specifications.

## Constraints

**Constraints as friends.** No constraints = no criteria for excellence. Constraints make the problem solvable.

Infinite possibility paralyzes. Clear constraints — budget, schedule, weight, power — create a defined field. Creativity: elegant solutions within boundaries. Hoare: "Premature optimization is the root of all evil." Brooks goes further. Constraints are not deferred evils. They are the conditions that make design possible.

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

Clients confuse requirements with implementation. "Use React" is not a requirement. "Renders at 60fps" is. The how is the designer's problem.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

Every hour clarifying saves ten building the wrong solution.

Cautionary tale: helicopter project added "fly across the Atlantic" as a final requirement. Contradicted every constraint. But it was documented. The rational model treats all documented requirements as valid. No defense against absurdity. The helicopter was never built.

## User models

**User model.** Write down who the user is, what they know, what they need. It will be wrong. Wrong and precise beats vague.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

"The user is a domain expert who uses the command line daily" — everyone sees the assumption and challenges it. Write nothing: everyone fills in their own model, nobody disagrees. Explicit → testable → correctable. Empiricism applied to the most important unknown.

---

[← Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 6 →](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

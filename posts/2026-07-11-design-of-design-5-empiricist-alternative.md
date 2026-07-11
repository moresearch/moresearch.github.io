---
title: "Brooks on design, part 5: build, test, iterate — the empiricist method"
date: 2026-07-11
slug: brooks-design-empiricist-alternative
summary: "If we can't reason our way to a correct design, what do we do? Brooks: build, test, and iterate. Prototyping, constraints, and user models are the tools."
tags: design, fred-brooks, empiricism, prototyping, constraints, formal-methods
---

Part 4 argued that the rational model is wrong. This part gives the alternative.

## Definition: the empiricist method

**Empiricist method.** You cannot think your way to a correct design. You must build, test, learn, and iterate.

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

This is not the Brooks of *The Mythical Man-Month*. The younger Brooks believed in planning — "plan to throw one away" meant doing the rational process twice. The older Brooks knows better. The first plan was never going to be right. No amount of upfront analysis would have fixed it. The only path to a good design runs through being wrong, noticing, and trying again.

Contrast this with Niklaus Wirth's **stepwise refinement** (1971): decompose a problem into sub-problems, refine each until trivial to code. This works when you already understand the problem well enough to decompose it correctly. Brooks's response: you never do. The real decomposition emerges through building, testing, and discovering which boundaries were wrong. Wirth's method works for well-understood problems. Brooks's empiricism works for everything else.

The method has six steps:
1. **Understand the domain** — study users, context, constraints
2. **Design something** — knowing it will be wrong
3. **Build a prototype** — make it concrete enough to test
4. **Test with real users** — watch what they do, not what they say
5. **Iterate** — use what you learned
6. **Build incrementally** — grow the system in steps, testing at each step

## Definition: the prototype

**Prototype.** A concrete version of an idea, built to be tested and discarded. Not a draft of the final product. A question posed to reality.

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

A failed prototype is not waste. It is the fastest way to learn the requirements. Specifications cannot do this — they are unfalsifiable. Only a running system can be wrong in a way that teaches something.

## Definition: formal methods

**Formal methods.** Proving programs correct by deduction from axioms. Works for small, well-specified modules. Cannot scale to large, evolving systems.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

The rationalist dream — correct by construction — survives only in computer science departments. Every other design discipline abandoned it centuries ago. Brooks's entire body of work is an argument that software should join them. *No Silver Bullet* (1986): no breakthrough will eliminate essential difficulty. *Computer Architecture* (1997, with Blaauw): even instruction sets evolved through trial, error, and market selection.

Richard Hamming put it directly: "The purpose of computing is insight, not numbers." Brooks would add: the purpose of designing is insight, not specifications. You build to learn. The artifact is the byproduct.

## Definition: constraints as friends

**Constraints as friends.** A problem with no constraints has no criteria for excellence. Constraints make the problem solvable.

The designer without constraints faces infinite possibility. The designer with clear constraints — budget, schedule, weight, power — has a defined field. Creativity lies in finding elegant solutions within boundaries.

C.A.R. Hoare warned that "premature optimization is the root of all evil" — one constraint should not dominate too early. Brooks goes further. Constraints are not deferred evils. They are the conditions that make design possible. Without them: infinite search space, no way to judge. With them: the job goes from impossible to hard.

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

Clients confuse requirements with implementation. "Use React" is not a requirement. "Renders at 60fps on mobile" is. The how is the designer's problem. The what is the client's — and the client doesn't know the what until they see a candidate how.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

Every hour clarifying the problem saves ten hours building the wrong solution.

Brooks illustrates the alternative with a cautionary tale. A military helicopter project spent months negotiating requirements. At the final meeting, someone added: "It shall fly itself across the Atlantic." This contradicted every prior constraint. But it was now in the document. The rational model has no defense against a late-breaking absurdity. It treats all documented requirements as equally valid. The helicopter was never built.

## Definition: user model

**User model.** An explicit description of who the user is, what they know, what they need, what constrains them. Write it down. It will be wrong. Wrong and precise beats vague.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

When you write "the user is a domain expert who uses the command line daily," everyone sees the assumption and can challenge it. When you write nothing, everyone fills in their own model and nobody realizes they disagree. Once explicit, assumptions become testable. Once testable, correctable. This is empiricism applied to the most important unknown: who is it for?

---

**This is part 5 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](/posts/brooks-design-conceptual-integrity)
- [Part 2: Why one mind must rule the design](/posts/brooks-design-one-mind-rule)
- [Part 3: How to protect designers from their organizations](/posts/brooks-design-protecting-designer)
- [Part 4: The waterfall model is wrong and harmful](/posts/brooks-design-rational-model)
- [Part 6: Why experts design the wrong thing beautifully](/posts/brooks-design-experts-divorce)
- [Part 7: Great designs come from great designers — not great processes](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

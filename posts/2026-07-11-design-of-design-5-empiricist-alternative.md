---
title: "Brooks on design, part 5: build, test, iterate — the empiricist method"
date: 2026-07-11
slug: brooks-design-empiricist-alternative
summary: "Brooks declares himself a hard-core empiricist: we cannot think our way to a correct design. We must build, test, and iterate. Prototyping, constraints, and user models are the tools of the trade."
tags: design, fred-brooks, empiricism, prototyping, constraints, formal-methods
---

Part 4 argued that the rational model is wrong — you don't know the goal at the start, and requirements co-evolve with the design. This part gives the positive alternative: if we can't reason our way to a correct design, what do we do instead?

Brooks names his position directly:

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

This is not the Brooks of *The Mythical Man-Month*. The younger Brooks believed in planning — the famous "plan to throw one away" was about doing the rational process twice, as if better requirements gathering would close the gap. The older Brooks knows better. The first plan was never going to be right. No amount of upfront analysis would have made it right. The only path to a good design runs through being wrong, noticing, and trying again.

## Stepwise refinement vs. empiricism

Compare this with the dominant methodology of Brooks's early career: Niklaus Wirth's *stepwise refinement* (1971). Wirth argued that you decompose a problem into sub-problems, then refine each until it is trivial to code. The method works beautifully — when you already understand the problem well enough to decompose it correctly on the first try.

Brooks's response, implicit but clear: stepwise refinement is the rational model in miniature. It assumes the decomposition is known. It never is. The real decomposition emerges through building, testing, and discovering which boundaries were wrong. Wirth's method works for well-understood problems — the kind you've solved before. Brooks's empiricism works for everything else — which is most of what matters.

The empiricist method has six steps:

1. **Understand the problem domain** — study the users, the context, the constraints
2. **Design something** — produce a candidate design, knowing it will be wrong
3. **Build an early prototype** — make it concrete enough to test
4. **Test it with real users** — watch what they do, not what they say
5. **Iterate** — use what you learned to improve the design
6. **Build incrementally** — grow the system in steps, testing at each step

## The prototype is the pivot

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

The prototype is not a draft. It is a question posed to reality: does this work? Does this make sense? Does it solve the right problem? A failed prototype is not wasted effort — it is the fastest mechanism for learning what the requirements actually are. Specifications cannot do this because they are unfalsifiable. Only a running system can be wrong in a way that teaches something.

This is not agile methodology, though the philosophical roots overlap. It is a deeper claim about the nature of design knowledge. Design knowledge is empirical, not deductive.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

You cannot derive a good design from axioms. You must discover it through interaction with the problem, the users, and the constraints. This is true of buildings, bridges, airplanes, and software. The rationalist dream — design from first principles, correct by construction — survives only in computer science departments. Every other design discipline abandoned it centuries ago.

Brooks's entire body of work is an argument that software should join them. In *No Silver Bullet* (1986), he argued there is no single breakthrough that will eliminate the essential difficulty of software design. In *Computer Architecture: Concepts and Evolution* (1997, with Gerrit Blaauw), he showed that even computer instruction sets — among the most rigorously specified artifacts ever designed — evolved through trial, error, and market selection, not deduction from first principles.

Richard Hamming, another giant of that generation, put it in his own aphoristic way: "The purpose of computing is insight, not numbers." Brooks would add: the purpose of designing is insight, not specifications. You build to learn. The learning is the point. The artifact is the byproduct.

## Constraints are friends

One of Brooks's most counterintuitive insights: a problem with no constraints has no criteria for excellence. When anything is possible, nothing is good.

Constraints reduce the search space. They make the problem tractable. A designer without constraints faces infinite possibility and is paralyzed by it. A designer with clear constraints — budget, schedule, weight, power, compatibility, regulatory requirements — has a defined field. The creativity lies in finding an elegant solution within the boundaries, not in pretending the boundaries do not exist.

This inverts the conventional wisdom of Brooks's era. C.A.R. Hoare famously warned that "premature optimization is the root of all evil" — one constraint, performance, should not dominate design too early. Brooks goes further. Constraints are not deferred evils. They are the conditions that make design possible. Without them, you have infinite search space and no way to distinguish good from bad. With them, the designer's job goes from impossible to merely hard.

Brooks draws a practical line:

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

This is aimed at clients and stakeholders who confuse requirements with implementation. "Use React" is not a requirement. "Renders at 60fps on mobile devices" is. "Works offline" is. "Accessible to screen readers" is. The how is the designer's problem. The what is the client's — and the client often does not know the what until they see a candidate how.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

This is the designer's real job — not translating requirements into blueprints, but helping the client discover what the requirements are. Every hour spent clarifying the problem saves ten hours of building the wrong solution. Tools change. This does not.

Brooks illustrates the alternative with a cautionary tale. A military helicopter project had spent months carefully negotiating requirements. At the final meeting, someone added: "It shall be capable of flying itself across the Atlantic." This contradicted every prior constraint — weight, range, cost, everything. But it was now in the document. The rational model treats all documented requirements as equally valid. It has no mechanism for saying "this requirement, despite being written down, is absurd." The helicopter was never built.

## User models: better wrong than vague

One of Brooks's most practical recommendations sounds obvious and is almost never done: write down your explicit model of the user. Who are they? What do they know? What do they need? What are their constraints?

The model will be wrong. That is fine.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

When you write "the user is a domain expert who uses the command line daily," everyone can see the assumption and challenge it. When you write nothing, everyone fills in their own implicit model and nobody realizes they disagree. The user model makes assumptions explicit. Once explicit, they become testable. Once testable, they become correctable. This is the empiricist method applied to the most important unknown in any design: who it is for. UX designers arrived at personas independently. Brooks got there from engineering. Same logic, different path.

---

**This is part 5 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the Reims Cathedral](/posts/brooks-design-conceptual-integrity)
- [Part 2: The one-mind rule](/posts/brooks-design-one-mind-rule)
- [Part 3: Protecting the designer](/posts/brooks-design-protecting-designer)
- [Part 4: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 6: How experts go wrong and the divorce of design](/posts/brooks-design-experts-divorce)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

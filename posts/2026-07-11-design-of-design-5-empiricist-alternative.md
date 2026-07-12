---
title: Brooks on Software Design Series: build, test, iterate
date: 2026-07-11
slug: brooks-design-empiricist-alternative
summary: "If we can't think our way to a correct design, what do we do? Build, test, iterate. The scientific method, applied to the glorious mess of making software."
tags: design, fred-brooks, empiricism, prototyping, constraints
---

The rational model is wrong. What's the alternative? Science. Not computer science — actual science. Hypothesize, test, learn, repeat. You are not designing a system. You are running experiments on your own ignorance. The experiments will hurt. That's how you know they're working.

> "I am a hard-core empiricist, in design as in science. I don't believe we can think our way to a correct design; we must build, test, and iterate."

## Definitions

**Empiricist method.** You cannot think your way to a correct design. Build, test, learn, iterate. If this sounds obvious, ask why your team spent three months debating an architecture nobody has prototyped. The debate felt productive. It was not. It was comforting. Comfort is not progress.

Not the Brooks of *The Mythical Man-Month*. Younger Brooks believed planning — "plan to throw one away" meant doing the rational process twice. Older Brooks: the first plan was never going to be right. No analysis would have fixed it. The only path runs through being wrong. Intelligence is knowing the design is wrong. Wisdom is shipping anyway to find out why. Seniority is having done this enough times to stop arguing about it.

Contrast Wirth's **stepwise refinement** (1971): decompose, refine until trivial. Works when you already understand the problem. Brooks: you never do. Decomposition emerges through building and testing. Wirth's method: well-understood problems. Brooks's method: everything else. Which is most of what you get paid for. The easy problems were automated decades ago.

Six steps: study domain → design (knowing it's wrong) → prototype → test with real users → iterate → build incrementally. Notice "argue about it in Slack" is not a step. Notice "write a design doc and never revisit it" is not a step. Notice "get sign-off from seven stakeholders" is not a step. The steps are hard. That's why they're skipped.

**Prototype.** A concrete version built to be tested and discarded. Not a draft. Not "the MVP we'll refactor later." A question posed to reality. Reality answers. Reality is usually right. Reality is also usually impolite about it.

> "The prototype is the pivot of the design process. It makes ideas concrete and thereby falsifiable. A prototype that fails teaches more than a specification that pleases."

Specifications are unfalsifiable. Nobody ever looked at a spec and said "this won't work." They looked at a spec and said "looks good" — which is worse. "Looks good" means "I haven't found the problem yet." Only running code can be wrong in a way that teaches. Only a crash tells you where the bridge was weak. Specifications don't crash. That's the problem.

**Formal methods.** Proving programs correct by deduction. Works for small modules. Cannot scale. The mathematicians disagree. The mathematicians don't ship software on deadlines with changing requirements.

> "Formal methods — proving programs correct — represent rationalism's last stand in software. They work in principle for small, well-specified modules. They cannot scale to large, complex, evolving systems. No other design discipline even attempts formal correctness proofs. Architects do not prove buildings will stand; they build them and test them."

The rationalist dream — correct by construction — survives only in CS departments and grant proposals. Every other discipline abandoned it centuries ago. Bridge builders test. Aircraft engineers test. The people who build things that kill you if they fail? They test relentlessly. The people who build things that lose your data? They debate formal verification on Hacker News.

*No Silver Bullet* (1986): no breakthrough eliminates essential difficulty. *Computer Architecture* (1997): even ISAs evolved through trial and error. Hamming: "The purpose of computing is insight, not numbers." Designing is insight, not specifications. The spec is the fossil. The prototype is the living thing. Fossils are evidence. They are not alive.

## Constraints

**Constraints as friends.** No constraints = no criteria for excellence. Constraints make the problem solvable. "Build anything" is not a brief. It is a cry for help. It is also why your last "greenfield" project was harder than the legacy one.

Infinite possibility paralyzes. Clear constraints — budget, schedule, weight, power — create a defined field. Creativity: elegant solutions within boundaries. Hoare: "Premature optimization is the root of all evil." Brooks goes further. Constraints are not deferred evils. They are the conditions that make design possible. Without walls, you're not in a room. You're in a void. People in voids don't design. They drift.

> "When you specify something to be designed, tell what properties you need, not how they are to be achieved."

Clients confuse requirements with implementation. "Use React" is not a requirement. "Renders at 60fps" is. The how is the designer's problem. The what is the client's. Both need to learn which is which. This learning takes years. It cannot be shortcut by a requirements-gathering workshop.

> "The hardest part of design is deciding what to design. The chief service of a designer is helping clients discover what they really want."

Every hour clarifying saves ten building the wrong solution. This ratio is remarkably stable across industries, technologies, and decades. It is either a law of nature or evidence that we are all, collectively, terrible at knowing what we want until we see it. Either way, budget for clarification. It's cheaper than rework. Nobody budgets for it.

Cautionary tale: helicopter project added "fly across the Atlantic" as a final requirement. Contradicted every constraint. But it was documented. The rational model treats all documented requirements as valid. No defense against absurdity. The helicopter was never built. The requirement lived forever in the document. Some say it still flies there, crossing the Atlantic on PDF pages, untroubled by physics.

## User models

**User model.** Write down who the user is, what they know, what they need. It will be wrong. Wrong and precise beats vague. "The user is a domain expert who uses the command line daily" is wrong and useful. "The user wants a good experience" is correct and worthless. One you can test against. The other is a fortune cookie.

> "Better a precise model, even if wrong, than a vague one. A precise model exposes its assumptions and invites correction; a vague one is unfalsifiable and thus unhelpful."

Write nothing: everyone fills in their own model, nobody disagrees. Explicit → testable → correctable. Empiricism applied to the most important unknown. UX designers call these personas. Engineers call them "that thing we should have written down six months ago." Both are right. The persona would have prevented the rewrite.

---

[← Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 6 →](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).


Brooks's principles apply beyond software. Conceptual integrity, the one-mind rule, the empiricist method — these are engineering principles that hold across any designed system. The building, the organization, the codebase, the protocol. The medium changes. The principles don't. That is the definition of engineering: principles that hold across domains.


> The empiricist method is not about not planning. It is about not trusting the plan. The plan is a hypothesis. The prototype is the experiment. The experiment either confirms the hypothesis or teaches something the plan missed.

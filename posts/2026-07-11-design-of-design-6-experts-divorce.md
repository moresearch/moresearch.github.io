---
title: "Brooks on Software Design Series: why experts get it wrong"
date: 2026-07-11
slug: brooks-design-experts-divorce
summary: "Novice mistakes are easy to spot. Expert mistakes are comprehensively, systematically wrong. And the organization is designed to prevent anyone from noticing."
tags: design, fred-brooks, expertise, paradigm-trap, feedback
---

Design is empirical. Two forces undermine empiricism. One is cognitive. The other is organizational. Both are invisible to the people inside them. Both are obvious to everyone outside. This is the definition of a structural problem.

## Definitions

**Paradigm trap.** Deep expertise in one paradigm becomes a liability when the paradigm shifts. Intuition points systematically wrong. Your instincts become your enemy. Slowly. Without telling you. You feel smarter than ever.

Novices make technical mistakes. Wrong data structure. Missed edge case. Tests catch these. The system breaks visibly. The junior engineer learns something. Progress occurs. The mistake is the lesson.

Experts make a different mistake. Designs that are **comprehensively, systematically wrong.** Internally consistent. Well-executed. Every part fits. And it solves the wrong problem. Or uses assumptions from a previous era. Or optimizes for a constraint that no longer exists. The expert is the last to know because the expert validates the design against the expert's own assumptions. That is not a review. That is a mirror. Mirrors don't find bugs.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

Term from Thomas Kuhn, *The Structure of Scientific Revolutions* (1962). Old-paradigm practitioners are last to see the new one. The people who built mainframes didn't invent PCs. The monolith experts didn't lead microservices. The relational purists said NoSQL would never work. They were right about some things and wrong about others. The paradigm shifted anyway. It always does.

Paradigm shifts: mainframes → minicomputers → microservices → probabilistic systems. At each step, the previous generation's experts produced beautiful, coherent designs for the wrong problem. A novice's bug fails a test. An expert's paradigm error passes all tests — the tests were written within the same paradigm. The test suite is also wrong. Nobody wrote a test for "are we solving the right problem?" Nobody ever does. It's not in the test plan template.

Defense: empiricism. Test designs against reality. Prototype. Watch users. Iterate on evidence. Ask yourself: when did you last change a fundamental decision because of something a user did? If the answer is "never," you are not an expert. You are a paradigm with a pulse.

**Divorce of design.** Designers, builders, and users become separated. Each handoff loses knowledge, accountability, and feedback speed. The designer designs in a vacuum. The builder builds from a document. The user suffers both. Nobody connects the dots because the dots are in different departments that report to different VPs who are in different meetings.

> "The designer who does not build, and the builder who does not use, are both crippled."

Epistemology, not career advice. Don't build → don't know if design works. Don't use → don't know what "works" means. Each handoff is lossy compression. Information is lost. Accountability is diffused. Nobody is wrong. Everyone is slightly less right. The sum of partial correctness is not correctness. It is a bug report nobody fully owns.

Wright brothers: designed, built, flew. Bad decision showed up that afternoon. Loop: hours. Modern software: months. Designers write specs. Implementers write code. Users file bugs. Each step loses information. The bug report is a shadow of the experience. The code is a shadow of the spec. The spec is a shadow of the intent. By the time the user suffers, the original decision has been through three lossy encodings and nobody can trace it to its source. Nobody is accountable. The process worked. The product didn't.

The architect who never codes designs abstractions elegant on paper and unbuildable. The developer who never meets users builds features technically impressive and functionally useless. Each is doing their job as defined. The definition is wrong. The job is making something that works for the user. Everything else is overhead. Most of what we call process is overhead that has been institutionalized into job descriptions.

Fix: shorten the loops. Designers build. Builders use. Users in the room. If your users aren't in the room, you're not designing for them. You're designing for your idea of them. Your idea is wrong. The structures that prevent this benefit the people who could change them. That is why they persist. That is why every part of this series sounds obvious and almost no organization does any of it.

Connects to Part 3. The designer needs protection — and contact with reality. Protected without feedback is marooned. Marooned designers produce marooned designs. The organization calls it "vision." It is isolation. Isolation produces coherence with the wrong world. The design is beautiful. The user is baffled. Everyone did their job.

---

[← Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 3](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 7 →](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

---
title: "Brooks on design, part 6: why experts get it wrong"
date: 2026-07-11
slug: brooks-design-experts-divorce
summary: "Novice mistakes are easy to spot. Expert mistakes are comprehensively, systematically wrong. And the divorce of designers from builders and users makes both harder to catch."
tags: design, fred-brooks, expertise, paradigm-trap, feedback
---

Parts 4-5: design is empirical. Part 6: two forces that undermine empiricism.

## Definitions

**Paradigm trap.** Deep expertise in one paradigm becomes a liability when the paradigm shifts. Intuition points systematically wrong.

Novices make technical mistakes. Wrong data structure. Missed edge case. Tests catch these. The system breaks visibly.

Experts make a different mistake. Their designs are **comprehensively, systematically wrong.** Internally consistent. Well-executed. Every part fits. And it solves the wrong problem. Or uses assumptions from a previous era. Or optimizes for a constraint that no longer exists.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

Term from Thomas Kuhn, *The Structure of Scientific Revolutions* (1962). Old-paradigm practitioners are last to see the new one.

Paradigm shifts: mainframes → minicomputers → microservices → probabilistic systems. Expert intuition becomes misleading. The design is coherent, so nobody spots the error. A novice's bug fails a test. An expert's paradigm error passes all tests — the tests were written within the same paradigm.

Defense: empiricism. Test designs against reality. Prototype. Watch users. Iterate on evidence. Ask: when did you last change a fundamental decision because of something a user did?

**Divorce of design.** Designers, builders, and users become separated. Each handoff loses knowledge, accountability, and feedback speed.

> "The designer who does not build, and the builder who does not use, are both crippled."

Epistemology, not career advice. Don't build → don't know if design works. Don't use → don't know what "works" means.

Wright brothers: designed, built, flew. Bad decision showed up that afternoon. Modern software: months. Designers write specs. Implementers write code. Users file bugs. Each step loses information.

Architect who never codes: abstractions elegant on paper, unbuildable in practice. Developer who never meets users: features technically impressive, functionally useless. Each doing their job as defined. The definition is wrong.

Fix: shorten the loops. Designers build. Builders use. Users in the room. The structures that prevent this benefit the people who could change them.

Connects to Part 3. The designer needs protection — and contact with reality. Protected without feedback is marooned.

---

**Part 6 of 7.** [← Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 7 →](https://blog.hackspree.com/#brooks-design-great-designers)

**References:** Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

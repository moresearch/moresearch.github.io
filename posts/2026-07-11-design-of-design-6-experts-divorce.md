---
title: "Brooks on design, part 6: why experts design the wrong thing beautifully"
date: 2026-07-11
slug: brooks-design-experts-divorce
summary: "Brooks observes that novice mistakes are easy to spot and fix, but expert mistakes are comprehensively, systematically wrong. And the growing separation between designers, builders, and users makes both kinds of error harder to catch."
tags: design, fred-brooks, expertise, paradigm-trap, feedback-loops
---

Parts 4-5 established that design is empirical — you learn by building, testing, and iterating. This part examines two forces that undermine empiricism.

## Definition: the paradigm trap

**Paradigm trap.** An expert's deep experience in one paradigm becomes a liability when the paradigm shifts. Their intuition, built over years, now points systematically in the wrong direction.

Novices make technical mistakes. They forget edge cases. They pick the wrong data structure. These errors are visible and fixable. A test catches them. The system breaks in a way that points to the problem.

Experts make a different kind of mistake. They produce designs that are **comprehensively, systematically wrong**. The design is internally consistent. Every part fits. It is well-executed. And it solves the wrong problem. Or it solves the right problem using assumptions from a previous era. Or it optimizes for a constraint that no longer exists.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

The term comes from Thomas Kuhn's *The Structure of Scientific Revolutions* (1962). Kuhn argued that scientific progress happens through paradigm shifts — and the established practitioners of the old paradigm are the last to see the new one. Brooks applies the same logic to design.

When the paradigm shifts — mainframes to minicomputers, monoliths to microservices, deterministic to probabilistic systems — the expert's intuition becomes systematically misleading. And because the design is internally coherent, nobody spots the error until too late. A novice's bug fails a test. An expert's paradigm error passes all the tests. The tests were written within the same paradigm.

The defense is empiricism. An expert who tests designs against reality — who prototypes, watches users, iterates on evidence — discovers their assumptions are wrong before building an entire system on top of them. An expert who designs from first principles and never tests ships a beautiful, coherent, wrong system. Ask: when did you last change a fundamental design decision because of something a user did?

## Definition: the divorce of design

**Divorce of design.** Designers, builders, and users become separated. Each handoff loses knowledge, accountability, and feedback speed.

> "The designer who does not build, and the builder who does not use, are both crippled."

This is epistemology, not career advice. If you don't build, you don't know if your design works. If you don't use, you don't know what "works" means. Each handoff — designer to implementer to user — is a site where knowledge is lost.

When the designer is also the builder and the user, feedback is immediate. The Wright brothers designed, built, and flew their aircraft. A bad decision showed up in the workshop that afternoon. Modern software has stretched this loop to months. Designers write specifications. Implementers write code. Users file bugs. Each step loses information.

The architect who never codes designs abstractions elegant on paper and unbuildable in practice. The developer who never meets users builds features technically impressive and functionally useless. Each is doing their job as the organization defines it. The organization is wrong.

Brooks's prescription: shorten the loops. Designers should build. Builders should use. Users should be in the room. The structures that prevent this are design problems — and the hardest kind, because the people who need to solve them benefit from them.

This connects to Part 3. The designer needs protection, yes — but also contact with reality. A protected designer who never builds or talks to users is not protected. They are marooned.

---

**This is part 6 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](/posts/brooks-design-conceptual-integrity)
- [Part 2: Why one mind must rule the design](/posts/brooks-design-one-mind-rule)
- [Part 3: How to protect designers from their organizations](/posts/brooks-design-protecting-designer)
- [Part 4: The waterfall model is wrong and harmful](/posts/brooks-design-rational-model)
- [Part 5: Build, test, iterate — the empiricist method](/posts/brooks-design-empiricist-alternative)
- [Part 7: Great designs come from great designers — not great processes](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

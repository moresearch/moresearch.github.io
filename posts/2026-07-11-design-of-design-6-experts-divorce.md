---
title: "Brooks on design, part 6: why experts design the wrong thing beautifully"
date: 2026-07-11
slug: brooks-design-experts-divorce
summary: "Brooks observes that novice mistakes are easy to spot and fix, but expert mistakes are comprehensively, systematically wrong. And the growing separation between designers, builders, and users makes both kinds of error harder to catch."
tags: design, fred-brooks, expertise, paradigm-trap, feedback-loops
---

Parts 4 and 5 established that design is empirical — you learn what works by building, testing, and iterating. This part examines two forces that undermine empiricism: the paradigm trap that makes experts go wrong, and the organizational trend that separates designers from the feedback that would catch their errors.

## How expert designers go wrong

Novices make technical mistakes. They forget edge cases, choose the wrong data structure, build something that doesn't scale under load. These errors are visible and usually fixable. A code review catches them. A load test catches them. The system breaks in a way that points to the problem.

Experts make a different kind of mistake. They produce designs that are **comprehensively, systematically wrong**. The design is internally consistent. It is well-executed. Every part fits with every other part. And it solves the wrong problem. Or it solves the right problem using assumptions from a previous era. Or it optimizes for a constraint that no longer exists.

Brooks's explanation is the paradigm trap. The term is borrowed from Thomas Kuhn's *The Structure of Scientific Revolutions* (1962), which argued that scientific progress occurs not through accumulation but through paradigm shifts — and that the established practitioners of an old paradigm are often the last to see the new one. Brooks applies the same logic to design. An expert has deep experience in one paradigm — one way of seeing problems, one set of solutions, one intuition for what works. That depth is their strength.

> "The expert's very expertise becomes a liability when the paradigm shifts. The habits that served him well in one era mislead him systematically in the next. He is not making small errors; he is solving yesterday's problem with today's tools."

When the paradigm shifts — mainframes to minicomputers, monoliths to microservices, deterministic to probabilistic systems — the expert's intuition becomes systematically misleading. They are not making small mistakes. They are building the wrong thing, beautifully.

And because the design is internally coherent, nobody spots the error until it is too late. A novice's bug fails a test. An expert's paradigm error passes all the tests — the tests were written within the same paradigm. The system works as designed. It just doesn't solve the problem that actually exists.

The defense is the empiricist method. An expert who tests designs against reality — who prototypes, watches users, iterates on evidence — will discover that their assumptions are wrong before they've built an entire system on top of them. An expert who designs from first principles and never tests will ship a beautiful, coherent, wrong system. The industry has plenty of both kinds. You can tell which is which by asking: when did you last change your mind about a fundamental design decision because of something a user did?

## The divorce of design

The second force that undermines empiricism is organizational. Brooks identifies a trend that has only accelerated since 2010: the separation of designers from implementers, and implementers from users.

> "The designer who does not build, and the builder who does not use, are both crippled."

This is not career advice. It is epistemology. If you don't build, you don't know whether your design works. If you don't use, you don't know what "works" means. Each handoff — designer to implementer to user — is a site where knowledge is lost.

When the designer is also the builder and the user, feedback is immediate and brutal. The Wright brothers designed, built, and flew their aircraft. A bad design decision showed up in the workshop that afternoon, or worse, in the air. The feedback loop was measured in hours. Modern software has stretched this loop to weeks or months. Designers produce specifications. Implementers produce code. Users produce bug reports. Each handoff loses information, accountability, and the rapid feedback that drives empirical improvement.

The consequences are visible everywhere. The architect who never writes code designs abstractions that are elegant on paper and unbuildable in practice. The developer who never talks to users builds features that are technically impressive and functionally useless. Each is doing their job as the organization defines it. The organization is wrong.

Brooks's prescription is straightforward: shorten the loops. Designers should build. Builders should use. Users should be in the room. The organizational structures that prevent this are design problems in their own right — and they are the hardest kind, because the people who need to solve them are the people the structures benefit.

This connects back to Part 3's argument about protecting the designer. The designer needs protection from external pressure, yes — but also from isolation. A protected designer who never builds and never talks to users is not protected. They are marooned. The protection must include enforced contact with reality: prototypes, user tests, production incidents, the things that tell you your design is wrong before you've bet the company on it.

## The connection

The paradigm trap and the divorce of design are the same problem at different scales. The paradigm trap is what happens when an individual designer loses contact with reality — when their mental model becomes self-reinforcing and untested. The divorce of design is what happens when an organization systematically severs the feedback loops that would prevent this.

Fixing either requires the same thing: empiricism, enforced by structure. Experts who test. Designers who build. Builders who use. Feedback loops measured in hours, not months. The tools for this exist — prototyping, incremental development, user research, observability in production. What is missing, in most organizations, is the will to restructure around them.

---

**This is part 6 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the Reims Cathedral](/posts/brooks-design-conceptual-integrity)
- [Part 2: The one-mind rule](/posts/brooks-design-one-mind-rule)
- [Part 3: Protecting the designer](/posts/brooks-design-protecting-designer)
- [Part 4: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 5: The empiricist alternative](/posts/brooks-design-empiricist-alternative)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

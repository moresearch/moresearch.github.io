---
title: "Brooks on design, part 2: why one mind must rule the design"
date: 2026-07-11
slug: brooks-design-one-mind-rule
summary: "Conceptual integrity requires a single mind, or at most a resonant pair, controlling the design. Committees produce compromises, not coherence."
tags: design, fred-brooks, one-mind, collaboration, brooks-law
---

Part 1 defined conceptual integrity. This part asks: how do you achieve it?

The answer is simple and unpopular. **The design must proceed from one mind, or from a very small number of agreeing resonant minds.**

## Definition: the one-mind rule

**One-mind rule.** The conceptual design — the core abstractions, the primitives, their relationships, the mental model presented to the user — must be controlled by one person. At most, two in genuine resonance.

Brooks is explicit. Great designs are attributed to individuals or pairs. Not committees. Not teams. Not communities. The list of counterexamples is empty.

> "The Design of Design sharpens the earlier contention: the design must represent the vision of one designer or, at most, a pair."

In *The Mythical Man-Month* (1975), Brooks allowed that implementation could be done by teams. The architecture had to belong to one mind. By 2010, he sharpened further. The architecture itself can have at most two authors — and only if they are genuinely resonant.

> "Two people can serve as one mind only if they are in genuine resonance — finishing each other's thoughts, sharing a mental model so deeply that each knows what the other would decide. Three cannot."

**Resonant pair.** Two designers who share a mental model so completely that either can speak for the architecture. Brooks found this with Gerrit Blaauw on System/360. It is rare. Most pairs never achieve it.

## Definition: committee design

**Committee design.** Multiple stakeholders each add their requirements. The result accommodates everyone and satisfies no one.

> "Design by committee produces designs that offend no one and satisfy no one. Each member's wish list is accommodated, each objection smoothed over, until the result is a feature-laden compromise lacking any coherent vision."

The problem is structural. Each new mind adds assumptions. These must be reconciled. Reconciliation produces compromises. Each compromise chips away at integrity. The committee didn't make the design better. It made it safer. Safety is the enemy of coherence.

This is not an argument against all collaboration. Teams help with requirements elicitation (more perspectives, more edge cases). Teams help with exploring the design space (brainstorming, alternatives). Teams help with implementation (parallel work, code review). But the conceptual design must belong to one person.

## The cost of collaboration

Brooks quantifies three costs:

**Learning cost.** Each new person must learn the shared vision. If the vision lives in one head, transferring it to *n* people takes *n × l* effort. For ten people, the designer spends more time teaching than designing.

**Communication cost.** *n* people have *n(n−1)/2* communication paths. Five people: 10 paths. Ten: 45. A hundred: 4,950. Each path carries misunderstandings and negotiated compromises. Overhead grows quadratically. Throughput grows linearly.

**Change control cost.** More contributors means harder changes. More people to consult. More objections. More downstream effects. The design calcifies.

> "Adding manpower to a late software project makes it later." — *The Mythical Man-Month*, 1975

**Brooks's Law.** Adding people to a late project makes it later. The same n(n−1)/2 communication paths that doom late projects also doom committee architectures. Adding designers dilutes the design.

## Definitions: essential and accidental complexity

In *No Silver Bullet* (1986), Brooks distinguished two kinds of complexity:

**Essential complexity.** Inherent in the problem. Cannot be eliminated by better tools or methods.

**Accidental complexity.** Imposed by our tools and methods. Can be reduced.

The rational model treats all complexity as accidental — as if process could eliminate it. Brooks argued that essential complexity remains no matter what you do. Conceptual integrity is how you manage essential complexity: one voice in the design, even if a hundred hands build it.

## Why this is structural

The one-mind rule is not a preference. It is a requirement.

A design is a set of interdependent decisions. Each constrains the others. When different people make different decisions, the constraints conflict. The system acquires multiple personalities. The user suffers.

"Design by community" is incoherent. A community has many visions. The output is not a design. It is a negotiated settlement. Settlements can govern societies. They cannot produce coherent software.

Melvin Conway, Brooks's IBM contemporary, stated this as a law in 1968: "Organizations which design systems are constrained to produce designs which are copies of the communication structures of these organizations." Read alongside Brooks: if you want a coherent design, you need a coherent design organization. One mind. A committee's communication structure is a complete graph. Its design output will be, too.

The practical implication: every project needs a design owner. One person who reviews every interface, every abstraction, every user-visible decision. They can say no without escalation. Their job is conceptual integrity. Everything else follows.

The next part asks: how do you protect that person from the organization?

---

**This is part 2 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](/posts/brooks-design-conceptual-integrity)
- [Part 3: How to protect designers from their organizations](/posts/brooks-design-protecting-designer)
- [Part 4: The waterfall model is wrong and harmful](/posts/brooks-design-rational-model)
- [Part 5: Build, test, iterate — the empiricist method](/posts/brooks-design-empiricist-alternative)
- [Part 6: Why experts design the wrong thing beautifully](/posts/brooks-design-experts-divorce)
- [Part 7: Great designs come from great designers — not great processes](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

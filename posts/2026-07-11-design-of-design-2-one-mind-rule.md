---
title: "Brooks on Software Design Series: one mind"
date: 2026-07-11
slug: brooks-design-one-mind-rule
summary: "Conceptual integrity requires one mind — or at most a resonant pair. Committees produce compromises, not coherence."
tags: design, fred-brooks, one-mind, collaboration, brooks-law
---

Conceptual integrity requires one mind. How do you achieve that in an organization of many?

**The design must proceed from one mind, or from a very small number of agreeing resonant minds.**

## Definitions

**One-mind rule.** The conceptual design — core abstractions, primitives, relationships, user mental model — must be controlled by one person. At most, two in genuine resonance.

> "The Design of Design sharpens the earlier contention: the design must represent the vision of one designer or, at most, a pair."

In 1975, Brooks allowed implementation by teams if architecture belonged to one mind. By 2010: architecture itself can have at most two authors.

**Resonant pair.** Two designers who share a mental model so completely either can speak for the architecture.

> "Two people can serve as one mind only if they are in genuine resonance — finishing each other's thoughts, sharing a mental model so deeply that each knows what the other would decide. Three cannot."

Brooks found this with Gerrit Blaauw on System/360. Rare. Most pairs never achieve it.

**Committee design.** Multiple stakeholders add requirements. Result: accommodates everyone, satisfies no one.

> "Design by committee produces designs that offend no one and satisfy no one. Each member's wish list is accommodated, each objection smoothed over, until the result is a feature-laden compromise lacking any coherent vision."

Structural. Each new mind adds assumptions. Reconciliation produces compromises. Each compromise chips at integrity. The committee made it safer, not better. Safety is the enemy of coherence.

Not an argument against all collaboration. Teams help with requirements (more edge cases), design space exploration (brainstorming), and implementation (parallel work). The conceptual design must belong to one person.

## Costs

**Learning cost.** Transferring a vision to *n* people takes *n × l* effort. For ten people, more time teaching than designing.

**Communication cost.** *n* people = *n(n−1)/2* paths. Five: 10. Ten: 45. One hundred: 4,950. Overhead grows quadratically.

**Change control cost.** More contributors = harder changes. The design calcifies.

**Brooks's Law.** Adding people to a late project makes it later. Same math applies to design. Adding designers dilutes.

> "Adding manpower to a late software project makes it later." — *The Mythical Man-Month*, 1975

**Essential complexity.** Inherent in the problem. Cannot be eliminated.

**Accidental complexity.** Imposed by tools and methods. Can be reduced.

From *No Silver Bullet* (1986). The rational model treats all complexity as accidental. Brooks: essential complexity remains. Conceptual integrity manages it — one voice in the design, a hundred hands building.

The rule is not a preference. A design is interdependent decisions. Different people make different decisions — constraints conflict — system acquires multiple personalities — user suffers. "Design by community" is incoherent. Output is a negotiated settlement. Settlements govern societies. They cannot produce coherent software.

Conway's Law (1968): organizations produce designs that copy their communication structures. Read with Brooks: coherent design needs coherent design organization. One mind. A committee's communication graph is complete. Its output will be, too.

Implication: every project needs a design owner. Reviews every interface. Says no without escalation. Job: conceptual integrity.

---

[← Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 3 →](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

---
title: Brooks on Software Design Series: one mind
date: 2026-07-11
slug: brooks-design-one-mind-rule
summary: "Conceptual integrity requires one mind — or at most a resonant pair. Committees produce compromises, not coherence. Three people is already a negotiation."
tags: design, fred-brooks, one-mind, collaboration, brooks-law
---

Conceptual integrity requires one mind. How do you achieve that in an organization of many? Short answer: you don't. You fight a rearguard action and hope for the best.

**The design must proceed from one mind, or from a very small number of agreeing resonant minds.** Most organizations read this, nod, and schedule a cross-functional alignment workshop. The workshop produces a shared document. The document has no owner. Nothing improves. This is the industry in microcosm.

## Definitions

**One-mind rule.** The conceptual design — core abstractions, primitives, relationships, user mental model — must be controlled by one person. At most, two in genuine resonance.

> "The Design of Design sharpens the earlier contention: the design must represent the vision of one designer or, at most, a pair."

In 1975, Brooks allowed implementation by teams if architecture belonged to one mind. By 2010: architecture itself can have at most two authors. At this rate, by 2045 he'd argue for half a designer. The trend line is clear. The industry is moving the opposite direction.

**Resonant pair.** Two designers who share a mental model so completely either can speak for the architecture. Finishing each other's thoughts. Same taste. Same instincts. Same willingness to say no to the same things.

> "Two people can serve as one mind only if they are in genuine resonance — finishing each other's thoughts, sharing a mental model so deeply that each knows what the other would decide. Three cannot."

Brooks found this with Gerrit Blaauw on System/360. Rare. Most pairs are just two people who've learned which topics to avoid. Resonance is not collaboration. Resonance is one mind in two bodies. If you have to explain your decisions to your partner, you are not a resonant pair. You are coworkers.

**Committee design.** Multiple stakeholders add requirements. Result: accommodates everyone, satisfies no one. The platypus is nature's committee design. It works. Nobody would design it from scratch.

> "Design by committee produces designs that offend no one and satisfy no one. Each member's wish list is accommodated, each objection smoothed over, until the result is a feature-laden compromise lacking any coherent vision."

Structural. Each new mind adds assumptions. Reconciliation produces compromises. Each compromise chips at integrity. The committee made it safer, not better. Safety is the enemy of coherence. Never confuse "nobody objected" with "this is good." Nobody objects to mediocre food either. That's why most restaurants are forgettable.

Teams help with requirements (more edge cases), design space exploration (brainstorming), and implementation (parallel work). The conceptual design must belong to one person. Teams exist to execute, not to design. Meetings exist for reasons less clear.

## Costs

**Learning cost.** Transferring a vision to *n* people takes *n × l* effort. For ten people, more time teaching than designing. Onboarding takes six months and nobody questions it because everyone has forgotten it could be otherwise.

**Communication cost.** *n* people = *n(n−1)/2* paths. Five: 10. Ten: 45. One hundred: 4,950. Overhead grows quadratically. Throughput does not. Mathematics is cruel and indifferent to your standup cadence.

**Change control cost.** More contributors = harder changes. The design calcifies. A living vision becomes a frozen document nobody fully owns. Then someone says "we should refactor" and the cycle begins again, with a larger committee this time.

**Brooks's Law.** Adding people to a late project makes it later. Same math applies to design. Adding designers dilutes. Your manager's solution to a late project is more people. Your manager's solution is wrong. Your manager is applying a linear fix to a quadratic problem. Mathematics will win.

> "Adding manpower to a late software project makes it later." — *The Mythical Man-Month*, 1975

**Essential complexity.** Inherent in the problem. Cannot be eliminated. Like death, taxes, and npm dependencies.

**Accidental complexity.** Imposed by tools and methods. Can be reduced. This is what you spend most of your time fighting. The rest is meetings about complexity reduction strategies you'll never implement.

From *No Silver Bullet* (1986). The rational model treats all complexity as accidental. Brooks: essential complexity remains. Conceptual integrity manages it — one voice in the design, a hundred hands building. Without the one voice, the hundred hands build a hundred different things and call it a microservices architecture.

The rule is structural, not preferential. A design is interdependent decisions. Different people make different decisions — constraints conflict — system acquires multiple personalities — user suffers. "Design by community" is incoherent. Output is a negotiated settlement. Settlements govern societies. They cannot produce coherent software. The EU is not an API. If it were, it would have seventeen conflicting ways to authenticate.

Conway's Law (1968): organizations produce designs that copy their communication structures. Read with Brooks: coherent design needs coherent design organization. One mind. A committee's communication graph is complete. Its output will be, too. This is why your microservices map exactly to your org chart. It's not supposed to. It's supposed to map to the problem. The problem doesn't care about your reporting lines.

Every project needs a design owner. Reviews every interface. Says no without escalation. Job: conceptual integrity. If this person doesn't exist on your project, they are not you. Find them. Or become them. Either way, you're currently in trouble you haven't noticed yet.

---

[← Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 3 →](https://blog.hackspree.com/#brooks-design-protecting-designer) · [Part 4](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).


Brooks's principles apply beyond software. Conceptual integrity, the one-mind rule, the empiricist method — these are engineering principles that hold across any designed system. The building, the organization, the codebase, the protocol. The medium changes. The principles don't. That is the definition of engineering: principles that hold across domains.

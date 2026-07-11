---
title: "protect the designer"
date: 2026-07-11
slug: brooks-design-protecting-designer
summary: "If conceptual integrity requires one mind, you must protect that mind from the organization. Real authority, organizational backing, separation of architecture from implementation."
tags: design, fred-brooks, system-360, interface-design, architecture
---

One mind must control the design. How do you protect that mind?

Organizations generate feature requests, compatibility demands, and stakeholder preferences. Each is reasonable. Together, they destroy coherence. The designer who says no needs protection, or the no won't stick.

## Definitions

**Interface integrity.** The interface *is* the system for the user. It must have conceptual integrity above all else.

> "For the user interface, conceptual integrity is even more essential. The interface is the system for the user. If the interface has multiple personalities, the user must learn each one, must decide which to use when, and will be confused by their inconsistencies."

You can distribute the backend. You cannot distribute the user's mental model. The interface owner needs veto power. Real veto — not advisory. Advisory means the VP overrides the designer. VP wins. User loses.

**Architecture/implementation separation.** Small team defines *what*. Large team builds *how*. Roles distinct. Staffed differently.

> "The architecture must be separated from implementation. This was the key organizational insight of System/360: a small architecture team defines what the machine is; a large implementation team builds it. The architecture team must be protected; the implementation team must be coordinated."

Architecture team: protection from external pressure. Implementation team: coordination across contributors. Conflate them: nobody does either well.

## The System/360 pattern

> "On System/360, a small team — Brooks, Amdahl, Blaauw — controlled the architecture. We had the authority to say no. More importantly, we were protected from the organizational forces that dilute design: feature requests from field sales, compatibility demands, performance optimizations that compromise clean abstraction."

Two things mattered. Authority to say no. Protection to exercise it. Field sales wanted features. Customers demanded compatibility. Engineers wanted optimizations. Each reasonable alone. Together: destroyed coherence. The team had power — and organizational backing.

David Parnas reached the same conclusion. **Information hiding** (1972): modules conceal design decisions from each other. Cannot do this with a committee. Hiding needs one mind to decide what to expose and bury.

Every project needs a design owner. One person. Reviews every interface, abstraction, user-visible decision. Authority to say no without escalation. Hard role: hold the entire system in your head, have taste, say no repeatedly — and be trusted.

The industry sort of does this. Architects, tech leads, staff engineers. Rarely with enough separation to say no to the VP. We don't train, hire, or protect for this role. Then we wonder why systems feel like patchwork quilts.

> "Plan to throw one away; you will, anyhow." — *The Mythical Man-Month*, 1975

Accept the first version will be wrong. Bridges to Part 4.

> "The building of a design, indeed, is the forcing of the will of one upon the stuff of the world."

Design is not consensus. It is imposition. Coherence forced onto a medium with no opinion. An act of authority. Does your organization have the nerve?

---

[← Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 4 →](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).

---
title: "Brooks on design, part 3: how to protect designers from their organizations"
date: 2026-07-11
slug: brooks-design-protecting-designer
summary: "If conceptual integrity requires one mind, the practical question is how to protect that mind from organizational forces that dilute design. Brooks's answer: real authority, organizational backing, and a clean separation between architecture and implementation."
tags: design, fred-brooks, system-360, interface-design, architecture, parnas
---

Part 2 established that conceptual integrity requires one mind. The next question: how do you protect that designer from the organization?

Organizations generate feature requests, compatibility demands, performance optimizations, and stakeholder preferences. Each request is individually reasonable. Collectively, they destroy coherence. The designer who says no to all of them needs protection, or the no won't stick.

## Definition: interface integrity

**Interface integrity.** The user interface must have conceptual integrity because the interface *is* the system for the user.

> "For the user interface, conceptual integrity is even more essential. The interface is the system for the user. If the interface has multiple personalities, the user must learn each one, must decide which to use when, and will be confused by their inconsistencies."

Everything the user knows comes through the interface. If it lacks integrity, the system lacks it — however clean the internals. You can distribute the backend across a hundred services. You cannot distribute the user's mental model across multiple personalities.

The interface owner needs veto power. Real veto power — not advisory. Advisory authority means the VP who wants a feature overrides the designer who says it would fracture the experience. The VP wins. The user loses. The system joins the patchwork quilt.

## The System/360 pattern

> "On System/360, a small team — Brooks, Amdahl, Blaauw — controlled the architecture. We had the authority to say no. More importantly, we were protected from the organizational forces that dilute design: feature requests from field sales, compatibility demands, performance optimizations that compromise clean abstraction."

Two things mattered. The authority to say no. And the protection to exercise it without being overruled. Field sales wanted features. Existing customers demanded compatibility. Performance engineers wanted optimizations. Each request was reasonable alone. Together, they would have destroyed coherence. Brooks, Amdahl, and Blaauw had the power to say no — and the organizational backing to make it stick.

## Definition: architecture/implementation separation

**Architecture/implementation separation.** A small architecture team defines *what* the system is. A large implementation team builds *how* it works. The roles are distinct. They must be staffed differently.

> "The architecture must be separated from implementation. This was the key organizational insight of System/360: a small architecture team defines what the machine is; a large implementation team builds it. The architecture team must be protected; the implementation team must be coordinated."

The architecture team needs protection from external pressure. The implementation team needs coordination across many contributors. Conflating the two means nobody does either well.

David Parnas reached the same conclusion independently. His principle of **information hiding** (1972): modules should conceal design decisions from each other. You cannot do this with a committee. Hiding requires one mind to decide what to expose and what to bury. The interface between modules is a design decision. Making it coherent requires one designer.

## What this means

The one-mind rule is uncomfortable in an industry that valorizes collaboration. Brooks's point is not that teams are bad. It is that teams need a real design authority.

Every project needs a design owner. One person who reviews every interface, every abstraction, every user-visible decision. They can say no without escalation. This role is hard to fill. It requires someone who can hold the entire system in their head, who has taste, who is willing to say no repeatedly — and who the organization trusts enough to give real authority.

The industry sort of does this — architects, tech leads, staff engineers. But rarely with enough separation that the design owner can actually say no to the VP. We do not train for this role. We do not reward it in hiring. We do not protect it in structures. And then we wonder why most systems feel like patchwork quilts.

> "Plan to throw one away; you will, anyhow." — *The Mythical Man-Month*, 1975

The first version will be wrong. Accept it. Build it. Learn. Build the right one. This bridges to the next part: the argument that the rational model, which assumes you can get requirements right before building, is structurally wrong.

> "The building of a design, indeed, is the forcing of the will of one upon the stuff of the world."

A design is not a consensus. It is an imposition. The designer imposes coherence on a medium that has no opinion about coherence. Design is an act of authority. The question is whether your organization has the nerve to grant it.

---

**This is part 3 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity — the most important property](/posts/brooks-design-conceptual-integrity)
- [Part 2: Why one mind must rule the design](/posts/brooks-design-one-mind-rule)
- [Part 4: The waterfall model is wrong and harmful](/posts/brooks-design-rational-model)
- [Part 5: Build, test, iterate — the empiricist method](/posts/brooks-design-empiricist-alternative)
- [Part 6: Why experts design the wrong thing beautifully](/posts/brooks-design-experts-divorce)
- [Part 7: Great designs come from great designers — not great processes](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

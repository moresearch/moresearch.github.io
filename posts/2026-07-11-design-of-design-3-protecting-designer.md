---
title: "Brooks on design, part 3: protecting the designer"
date: 2026-07-11
slug: brooks-design-protecting-designer
summary: "If conceptual integrity requires one mind, the practical question is how to protect that mind from the organizational forces that dilute design. Brooks's answer: real authority, organizational backing, and a clean separation between architecture and implementation."
tags: design, fred-brooks, system-360, interface-design, architecture, parnas
---

Part 2 established that conceptual integrity requires one mind — or at most a resonant pair. The practical question that follows is uncomfortable: how do you protect that designer from the organization?

Organizations are not neutral. They generate feature requests, compatibility demands, performance optimizations, and the accumulated preferences of every stakeholder who has ever been in a meeting. Each request is individually reasonable. Collectively, they destroy coherence. The designer who says no to all of them needs protection, or the no won't stick.

## The interface is the system

The place this matters most is the user interface.

> "For the user interface, conceptual integrity is even more essential. The interface is the system for the user. If the interface has multiple personalities, the user must learn each one, must decide which to use when, and will be confused by their inconsistencies."

Everything the user knows about your software comes through the interface. If it lacks conceptual integrity, the system lacks it — however clean the internals. You can distribute the backend across a hundred services. You cannot distribute the user's mental model across multiple personalities and expect coherence.

This is why Brooks insists the interface must be "tightly controlled by one mind." The interface owner needs the authority to veto — to say no to the feature that makes sense for one backend team but would fracture the user's experience. That authority must be real. Advisory authority means the VP who wants a feature overrides the designer who says it would break coherence. The VP wins, the user loses, and the system joins the patchwork quilt.

## The System/360 model

Brooks's answer is rooted in experience. System/360 was one of the most influential computer architectures ever designed, and it succeeded in part because of how the design team was structured.

> "On System/360, a small team — Brooks, Amdahl, Blaauw — controlled the architecture. We had the authority to say no. More importantly, we were protected from the organizational forces that dilute design: feature requests from field sales, compatibility demands, performance optimizations that compromise clean abstraction."

The authority to say no mattered. The *protection* to exercise it without being overruled mattered more. Field sales wanted features for specific customers. Existing customers demanded backward compatibility. Performance engineers wanted to optimize the hot paths. Each request was individually reasonable. Collectively, they would have destroyed the machine's coherence. Brooks, Amdahl, and Blaauw had the power to say no — and the organizational backing to make it stick.

This is not autonomy for its own sake. It is autonomy within the right constraints: clear overriding objectives, a schedule with urgency, and the freedom to make design decisions within those bounds. The worst outcome — and the most common — is a designer with no real authority, overseen by a committee that can override any decision but takes responsibility for none. The designer gets the blame when the design fails. The committee never does.

## Architecture vs. implementation

> "The architecture must be separated from implementation. This was the key organizational insight of System/360: a small architecture team defines what the machine is; a large implementation team builds it. The architecture team must be protected; the implementation team must be coordinated. The roles are distinct and must be staffed differently."

This is Brooks's most practical organizational pattern. Separate the design authority from the build workforce. Give the design authority real power over the *what*. Give the build workforce the scale to execute the *how*. The architecture team needs protection from external pressure. The implementation team needs coordination across many contributors. The skills are different. The staffing is different. Conflating them — having the same people do both — means nobody does either well.

David Parnas, working independently in the same era, reached the same conclusion from a different angle. His principle of information hiding (1972) requires modules to conceal design decisions from each other. You cannot do this with a committee, because hiding requires a single mind to decide what to expose and what to bury. The interface between modules is a design decision. Making it coherent requires one designer, or the interfaces will reflect the organizational chart rather than the problem structure. Conway's Law in reverse.

## What this means for teams today

The one-mind rule is uncomfortable in an industry that valorizes collaboration. Brooks's point is not that teams are bad. It is that teams need a *real* design authority.

Every project of any complexity should have a single person, or a tight pair, who owns the conceptual design. That person reviews every interface, every abstraction, every user-visible decision. They can say no without escalation. They are not the team lead, the engineering manager, or the product manager — though they may wear those hats. They are the design owner. Their primary responsibility is conceptual integrity.

The industry sort of does this — architects, tech leads, staff engineers — but rarely with enough separation that the design owner can actually say no to the VP who wants a feature. We do not train for this role. We do not reward it in hiring. We do not protect it in organizational structures. And then we wonder why most systems feel like patchwork quilts.

This role is hard to fill. It requires someone who can hold the entire system in their head, who has taste, who is willing to say no repeatedly to smart people — and who the organization trusts enough to give real authority. Brooks spent six decades arguing that everything else is secondary.

> "Plan to throw one away; you will, anyhow." — *The Mythical Man-Month*, 1975

The younger Brooks already understood that the first version would be wrong. Accept it. Build it. Learn from it. Then build the right one. This is the bridge to the next part — the argument that the rational model of design, which assumes you can get the requirements right before you build, is not just optimistic but structurally wrong.

> "The building of a design, indeed, is the forcing of the will of one upon the stuff of the world."

A design is not a consensus. It is an imposition. The designer imposes coherence on a medium that has no opinion about coherence. This is uncomfortable language — "forcing," "will," "one upon the world" — and Brooks means it to be uncomfortable. Design is an act of authority. The question is whether your organization has the nerve to grant it.

---

**This is part 3 of a 7-part series on Fred Brooks' *The Design of Design*.**
- [Part 1: Conceptual integrity and the Reims Cathedral](/posts/brooks-design-conceptual-integrity)
- [Part 2: The one-mind rule](/posts/brooks-design-one-mind-rule)
- [Part 4: The rational model is wrong](/posts/brooks-design-rational-model)
- [Part 5: The empiricist alternative](/posts/brooks-design-empiricist-alternative)
- [Part 6: How experts go wrong and the divorce of design](/posts/brooks-design-experts-divorce)
- [Part 7: Growing great designers](/posts/brooks-design-great-designers)

**References:**
- Fred Brooks, *The Mythical Man-Month: Essays on Software Engineering*, Addison-Wesley, 1975. (Anniversary Edition with *No Silver Bullet*, 1995.)
- Fred Brooks, *No Silver Bullet: Essence and Accident in Software Engineering*, 1986.
- Fred Brooks and Gerrit A. Blaauw, *Computer Architecture: Concepts and Evolution*, Addison-Wesley, 1997.
- Fred Brooks, *The Design of Design: Essays from a Computer Scientist*, Addison-Wesley, 2010.

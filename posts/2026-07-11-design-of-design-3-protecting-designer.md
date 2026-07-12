---
title: Brooks on Software Design Series: protect the designer
date: 2026-07-11
slug: brooks-design-protecting-designer
summary: "If conceptual integrity requires one mind, you must protect that mind from the organization. Real authority, organizational backing. Or: how to keep your best designer from quitting."
tags: design, fred-brooks, system-360, interface-design, architecture
---

One mind must control the design. How do you protect that mind? The organization wants features. The designer wants coherence. One of them has to win. The organization has more people. The designer needs better defenses.

Organizations generate feature requests, compatibility demands, and stakeholder preferences. Each is reasonable alone. Together, they destroy coherence. The designer who says no needs protection, or the no won't stick. "No" costs political capital. "Yes" gets you promoted. Guess which one happens more.

## Definitions

**Interface integrity.** The interface *is* the system for the user. It must have conceptual integrity above all else. The backend can be a disaster. The user should never know. This is the entire job of API design.

> "For the user interface, conceptual integrity is even more essential. The interface is the system for the user. If the interface has multiple personalities, the user must learn each one, must decide which to use when, and will be confused by their inconsistencies."

You can distribute the backend across a hundred microservices on five clouds. You cannot distribute the user's mental model. Mental models don't shard. They don't load balance. They don't fail over. They just break.

The interface owner needs veto power. Real veto — not advisory. Advisory means the VP overrides the designer. VP wins. User loses. Feature ships. Nobody uses it. Everyone is confused. The VP moves to a new role. The designer inherits the mess. This is not hypothetical. This is Tuesday.

**Architecture/implementation separation.** Small team defines *what*. Large team builds *how*. Roles distinct. Staffed differently. The architecture team thinks. The implementation team does. Both are essential. Only one gets protected.

> "The architecture must be separated from implementation. This was the key organizational insight of System/360: a small architecture team defines what the machine is; a large implementation team builds it. The architecture team must be protected; the implementation team must be coordinated."

Architecture team: protection from external pressure. Implementation team: coordination across contributors. Conflate them: nobody does either well. This is your startup's "everyone is full-stack" model. It works until it doesn't. Then you hire architects and call it "maturing." Then the architects complain they have no authority. The cycle is predictable. Nobody reads Brooks. The cycle continues.

## The System/360 pattern

> "On System/360, a small team — Brooks, Amdahl, Blaauw — controlled the architecture. We had the authority to say no. More importantly, we were protected from the organizational forces that dilute design: feature requests from field sales, compatibility demands, performance optimizations that compromise clean abstraction."

Two things mattered. Authority to say no. Protection to exercise it. Field sales wanted features for customers. Engineering wanted optimizations. Customers demanded compatibility. Each reasonable alone. Together: destroyed coherence. The team had power — and organizational backing. Without backing, power is just a loudly stated opinion that gets overruled in the next steering committee.

David Parnas reached the same conclusion. **Information hiding** (1972): modules conceal design decisions from each other. Cannot do this with a committee. Hiding requires one mind to decide what to expose and bury. A committee exposes everything and hides nothing, which is also how it makes decisions. The meeting minutes are public. The reasoning is not.

Every project needs a design owner. One person. Reviews every interface, abstraction, user-visible decision. Authority to say no without escalation. Hard role: hold the entire system in your head, have taste, say no repeatedly — and be trusted. These people are underpaid relative to their value and over-stressed relative to their support. They know this. They stay anyway. That is the only reason your system still works.

The industry sort of does this. Architects, tech leads, staff engineers. Rarely with enough separation to say no to the VP. We don't train, hire, or protect for this role. Then we wonder why systems feel like patchwork quilts. We built the quilter's guild and asked why nobody weaves.

> "Plan to throw one away; you will, anyhow." — *The Mythical Man-Month*, 1975

Accept the first version will be wrong. If you're not embarrassed by version one, you shipped too late. Bridges to Part 4.

> "The building of a design, indeed, is the forcing of the will of one upon the stuff of the world."

Design is not consensus. It is imposition. Coherence forced onto a medium with no opinion. An act of authority. Does your organization have the nerve? Most don't. Most call a meeting. The meeting schedules a follow-up. Nothing is forced. Nothing coheres.

---

[← Part 2](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [Part 1](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [Part 4 →](https://blog.hackspree.com/#brooks-design-rational-model) · [Part 5](https://blog.hackspree.com/#brooks-design-empiricist-alternative) · [Part 6](https://blog.hackspree.com/#brooks-design-experts-divorce) · [Part 7](https://blog.hackspree.com/#brooks-design-great-designers)

Fred Brooks, *The Mythical Man-Month* (1975, Anniversary Ed. 1995), *No Silver Bullet* (1986), *The Design of Design* (2010). Brooks & Blaauw, *Computer Architecture: Concepts and Evolution* (1997).


Brooks's principles apply beyond software. Conceptual integrity, the one-mind rule, the empiricist method — these are engineering principles that hold across any designed system. The building, the organization, the codebase, the protocol. The medium changes. The principles don't. That is the definition of engineering: principles that hold across domains.


> Protecting the designer is not about giving them autonomy. It is about giving them authority. Autonomy without authority is frustration. Authority without protection is theater.

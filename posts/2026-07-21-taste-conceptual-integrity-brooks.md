---
title: Taste as Conceptual Integrity
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: What engineers call "taste" is what Fred Brooks called conceptual integrity. The Pentagon Wars shows what happens when it's absent.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

The sentence does not work, and it does not work for a precise reason. "Taste" belongs to the domain of aesthetics and personal preference. An inspector who makes structural judgments based on taste is not being overruled — they are being negligent. The category error is built into the word. When you call engineering judgment "taste," you reframe a question of correctness as a question of opinion, and opinions require no rebuttal. "That's just your taste" is a conversation-ender.

But the thing people are reaching for when they use the word is not imaginary. There is a real capacity involved in looking at a system and recognizing that something is wrong — that the parts no longer compose, that the design has accumulated contradictions, that the whole has lost its voice. Fred Brooks gave this property a name.

## The name

![The Design of Design — Fred Brooks's last book, his best, and his least read](/images/design-of-design.jpg)

Brooks called it **conceptual integrity**. He first articulated it in *The Mythical Man-Month* in 1975:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas."

Thirty-five years later, in *The Design of Design*, he sharpened the claim to its most concentrated form:

> "Most great works have been made by one mind. The exceptions have been made by two minds."

This is not an argument against collaboration in implementation. It is an observation about how coherence enters a designed thing. Design does not parallelize. Every additional mind introduces new assumptions. Assumptions must be reconciled. Reconciliation produces compromise. Each compromise chips at coherence. "Many hands make light work — Often. But many hands make more work — Always." You can divide the labor of building. You cannot divide the labor of deciding what the thing is.

Brooks's positive case is Reims Cathedral. Ground was broken in 1211. Structural work was completed by 1275, with decorative work continuing into the 1460s. Across those two and a half centuries, four master masons directed the construction: Jean d'Orbais, Jean-le-Loup, Gaucher of Reims, and Bernard de Soissons. Their names were inscribed in a labyrinth set into the nave floor — not merely as a signature, but as a public oath to a design larger than any single lifetime.

![Reims Cathedral — four architects, two and a half centuries, one coherent result](/images/reims-cathedral.jpg)

The coherence was not accidental. Reims was among the first buildings to use stones and materials of standardized sizes, which accelerated construction and reduced the opportunity for deviation between architects. The structural system — four-part rib vaults producing arcades of identical pillars rather than alternating designs — made the architectural rhythm self-enforcing. Once the pillar spacing and vault geometry were locked in, they dictated what followed. The original plan had mechanical authority, not just symbolic authority. Later architects could elaborate on the scheme. They could not alter it without breaking the structural logic.

The result is a building where you cannot identify the seams between architects. The west façade exhibits what architectural historians describe as "an unusual unity of style." The transept roses and the façade roses show evolving Gothic fashion across the centuries — bar tracery giving way to full Rayonnant — but the evolution reads as variation within a theme, not competing visions. The whole speaks with one voice. The first architect's plan had authority that outlived him. His successors submitted to it. Reims should not exist. That it does is the evidence.

## The counter-example

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle: 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

The Bradley Fighting Vehicle began as a light troop carrier. Over 17 years and $14 billion, successive stakeholders added requirements: more armor, more firepower, more troop capacity, larger contracts. Each request was defensible in isolation. The accumulation produced a vehicle that Sergeant Fanning, a character in the 1998 HBO film *The Pentagon Wars*, describes with the clarity of someone who has stopped being polite:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

Every clause names a stakeholder requirement that was reasonable on its own. Together they describe a machine that did nothing well because it was asked to do everything. Testing was manipulated to hide the failures. The officers responsible were promoted. Lt. Colonel James Burton — the one person in the story who acted as though coherence mattered — forced an honest live-fire test and was eventually forced into retirement.

The Bradley is the limiting case of "many good but independent and uncoordinated ideas." It was not a failure of intentions or competence. It was the structural consequence of a system in which nobody had the authority to say no — every stakeholder could add a requirement, and no single mind could reject one. The output was not a design. It was a negotiated settlement between competing interests.

## The structural cure

Brooks did not merely name the property. On IBM System/360, he, Gene Amdahl, and Gerrit Blaauw built the organizational structures that preserve it. The principles run through both books.

**One mind — at most two in genuine resonance.** Brooks and Blaauw achieved a state in which either could speak for the architecture. Brooks called it "genuine resonance" and added: "And two is indeed a magic number for collaborations; marriage was a brilliant invention and has a lot to be said for it." Three cannot do this. At three, the dynamic shifts from resonance to negotiation, and negotiation produces settlements, not coherence.

**Real veto power.** The architect must be able to say no — "repeatedly, to smart people with good arguments" — and have the decision hold. Advisory veto, where a VP can override the architect, is indistinguishable in practice from no veto at all. The VP approves the feature. The feature ships. The coherence degrades. The VP eventually moves on. The architect stays, holding the accumulated incoherence.

**The divorce of design from implementation.** A chapter of *The Design of Design* is devoted to the progressive separation, beginning around the 16th century, of designing from making. On System/360, a small architecture team defined *what* the system was. A large implementation team built *how* it worked. "The architecture team must be protected; the implementation team must be coordinated." Conflate the two and neither function is performed adequately. Architects drawn into implementation stop thinking about systemic coherence. Implementers asked to make architectural decisions optimize locally at the expense of the whole.

**Protection from organizational forces.** The architecture team was deliberately shielded from field sales, engineering, and customers — all of whom had legitimate demands that would have been reasonable to accommodate individually and fatal to accommodate collectively. The shielding was not a matter of status. It was a structural precondition for coherent design.

**Career paths that reward refusal.** Brooks implied this principle throughout both books but never stated it as bluntly as the evidence warrants. In most organizations, agreement builds political capital and refusal spends it. If the incentive structure penalizes the person who protects coherence, coherence will not be protected regardless of what the documentation claims. The Bradley program is the proof at the limit: the saboteurs were promoted. The honest officer was removed.

---

Ray's thought experiment succeeds because it exposes the gap between what "taste" can carry and what the role demands. "The building fell down after the inspector's taste was ignored" describes a negligent inspector, not an overruled professional. That is the problem with the word.

Reformulate with the right term: the Bradley became a death trap because no single mind had the authority to enforce the vehicle's conceptual integrity. That sentence names a role, a property, and a failure mode. An engineer can be held accountable for it.

What engineers call taste is conceptual integrity. Brooks gave us the term, the evidence, and the organizational principles. The remaining question is whether we have the nerve.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — Conceptual integrity as the most important consideration in system design.
- Brooks, F. P. (2010). *The Design of Design.* — "Most great works have been made by one mind. The exceptions have been made by two minds." "Many hands make light work — Often. But many hands make more work — Always." The divorce of design from implementation.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)

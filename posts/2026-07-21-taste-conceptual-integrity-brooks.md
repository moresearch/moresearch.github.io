---
title: Taste as Conceptual Integrity
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: What engineers call "taste" is what Fred Brooks called conceptual integrity. The Pentagon Wars shows what happens when it's absent.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

The sentence doesn't work, for a precise reason. "Taste" belongs to aesthetics and personal preference. An inspector who makes structural judgments based on taste is not being overruled — they are being negligent. When you call engineering judgment "taste," you reframe correctness as opinion, and opinions require no rebuttal. "That's just your taste" is a conversation-ender. But the thing people are groping toward when they invoke the word is not imaginary. Fred Brooks spent a career giving it a name.

![The Design of Design — Fred Brooks's final book, and his least read](/images/design-of-design.jpg)

Brooks called it **conceptual integrity**. In *The Mythical Man-Month* (1975) he wrote: "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." Thirty-five years later, in *The Design of Design*, he sharpened it further: "Most great works have been made by one mind. The exceptions have been made by two minds." This is not a preference for small teams. It is a structural observation about how coherence enters a designed thing. Design does not parallelize. Every additional mind introduces assumptions that must be reconciled. Reconciliation produces compromise. Compromise chips at coherence. "Many hands make light work — Often. But many hands make more work — Always." You can divide the labor of building. You cannot divide the labor of deciding what the thing is.

Brooks's positive case is Reims Cathedral. Ground was broken in 1211, structural work completed by 1275, decorative work continuing into the 1460s. Across those two and a half centuries, four master masons — Jean d'Orbais, Jean-le-Loup, Gaucher of Reims, Bernard de Soissons — directed construction. Their names were inscribed in a labyrinth in the nave floor: not a signature, but a public oath to a design larger than any single lifetime. The coherence was not accidental. Reims was among the first buildings to use standardized stones, which reduced opportunity for deviation between architects. The four-part rib vaults produced arcades of identical pillars — once the geometry was locked, later architects could elaborate but not alter. The plan had mechanical authority, not just symbolic authority. The result is a building with no visible seams between architects. Two centuries of evolving Gothic fashion read as variation within a theme, not competing visions. The first architect's plan had authority that outlived him. His successors submitted to it.

![Reims Cathedral — four architects, two and a half centuries, one coherent result](/images/reims-cathedral.jpg)

If Reims is what happens when the principle holds, the Bradley Fighting Vehicle is what happens when it doesn't.

![The Pentagon Wars (1998) — 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

The Bradley began as a light troop carrier. Over 17 years and $14 billion, successive stakeholders added requirements: more armor, more firepower, more troop capacity, larger contracts. Each defensible in isolation. The accumulation produced a vehicle that Sergeant Fanning, a character in the 1998 HBO film *The Pentagon Wars*, describes with the clarity of someone who has stopped being polite:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

Every clause names a stakeholder requirement that was reasonable alone. Together they describe a machine that did nothing well because it was asked to do everything. Testing was manipulated. The officers who rigged the results were promoted. Lt. Colonel James Burton — the one person who acted as though coherence mattered — forced an honest live-fire test and was eventually forced into retirement. The Bradley is the limiting case of "many good but independent and uncoordinated ideas." It was not a failure of intentions. It was the structural consequence of a system in which nobody had authority to say no. The output was not a design. It was a negotiated settlement.

Brooks did not merely name the property. On IBM System/360, he, Gene Amdahl, and Gerrit Blaauw built the organizational structures that preserve it.

**One mind — at most two in genuine resonance.** Brooks and Blaauw achieved a state in which either could speak for the architecture. "And two is indeed a magic number for collaborations; marriage was a brilliant invention and has a lot to be said for it." Three cannot do this. At three, resonance becomes negotiation, and negotiation produces settlements, not coherence.

**Real veto power.** The architect must be able to say no — "repeatedly, to smart people with good arguments" — and have it stick. Advisory veto, where a VP can override, is indistinguishable from no veto. The VP approves the feature. It ships. Coherence degrades. The VP moves on. The architect stays, holding the accumulated incoherence.

**The divorce of design from implementation.** A chapter of *The Design of Design* traces the progressive separation, beginning in the 16th century, of designing from making. On System/360, a small architecture team defined *what* the system was; a large implementation team built *how*. Conflate them and neither function works. Architects drawn into implementation stop thinking about coherence. Implementers asked to make architectural decisions optimize locally at the expense of the whole.

**Protection from organizational forces.** The architecture team was shielded from field sales, engineering, and customers — all of whom had legitimate demands, reasonable individually, fatal collectively. The shielding was a structural precondition, not a status symbol.

**Career paths that reward refusal.** Brooks implied this throughout both books but never stated it as bluntly as the evidence warrants. Agreement builds political capital. Refusal spends it. If the incentives punish coherence, coherence will not occur. The Bradley is the proof: the officers who rigged the tests were promoted. The officer who forced an honest evaluation was removed.

## Open questions for the agentic era

Coding agents produce tokens. They do not produce conceptual integrity. A function that compiles and passes tests can violate every design principle the surrounding codebase observes — not maliciously, but because coherence is not a statistical property. The agent has no model of the whole. It cannot acquire one from a context window.

This raises questions Brooks did not live to answer. If agents generate most code, does the one-mind rule still apply — or does it become impossible? Can an architect maintain coherence against a force that produces incoherence at machine speed? Does the architect's role become more critical (because the alternative is chaos at scale) or obsolete (because you cannot veto output faster than it arrives)? If the Bradley took 17 years to go wrong with humans in the loop, how long does it take with agents?

Brooks's principles — one mind, real veto, protected architects, separated design and implementation — assume a human author generating at human speed. The agent era breaks that assumption. The question is whether the principles survive the break.

---

Ray's experiment exposes the gap. Reformulate with the right term: the Bradley became a death trap because no single mind had the authority to enforce the vehicle's conceptual integrity. That sentence names a role, a property, and a failure mode. An engineer can be held accountable for it. What engineers are reaching for when they say "taste" is what Brooks called conceptual integrity. He gave us the term, the evidence, the principles. The remaining question is whether we have the nerve to use them.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — Conceptual integrity as the most important consideration in system design.
- Brooks, F. P. (2010). *The Design of Design.* — "Most great works have been made by one mind. The exceptions have been made by two minds." "Many hands make light work — Often. But many hands make more work — Always." The divorce of design from implementation.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)

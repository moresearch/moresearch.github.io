---
title: Taste as Conceptual Integrity
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: When engineers say someone has "taste," what they mean is that person can perceive conceptual integrity. Fred Brooks spent a career explaining why that's the most important property of any designed thing.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

The sentence doesn't work, for a precise reason. "Taste" sounds like personal preference — it belongs to the domain of font choices and wine. An inspector whose structural judgments are treated as taste isn't being overruled; they're being described as though their objections were aesthetic. But when a senior engineer looks at a proposed change and says "this doesn't feel right," they are not expressing a preference. They are perceiving something. The question is what.

Here is a definition: **taste is the ability to perceive conceptual integrity.** What the senior engineer detects — before they can articulate it, before they can produce evidence — is that the system has lost coherence. The parts no longer compose. The design has accumulated a contradiction. The whole no longer speaks with one voice. Taste is the faculty. Conceptual integrity is the property being perceived.

![The Design of Design — Fred Brooks's final book, and his least read](/images/design-of-design.jpg)

Fred Brooks defined **conceptual integrity** as the property of a system that feels like one mind designed it. In *The Mythical Man-Month* (1975) he wrote: "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." Thirty-five years later, in *The Design of Design*, he sharpened it: "Most great works have been made by one mind. The exceptions have been made by two minds." This is not a preference for small teams. It is a structural observation about how coherence enters a designed thing: through a single point of decision. Design does not parallelize. Every additional mind introduces assumptions that must be reconciled. Reconciliation produces compromise. Compromise chips at coherence. "Many hands make light work — Often. But many hands make more work — Always."

If taste is the ability to perceive this property, then the presence or absence of conceptual integrity is how you know whether someone's taste was listened to. Reims Cathedral is the positive case. Ground was broken in 1211, structural work completed by 1275, decorative work continuing into the 1460s. Across those two and a half centuries, four master masons directed construction; their names were inscribed in a labyrinth in the nave floor as a public oath to a design larger than any single lifetime. The coherence wasn't accidental. Reims was among the first buildings to use standardized stones; the four-part rib vaults produced arcades of identical pillars; once the geometry was locked, later architects could elaborate but not alter. The plan had mechanical authority, not just symbolic authority. The result is a building with no visible seams between architects. Two centuries of evolving Gothic fashion read as variation within a theme. Someone with taste set the constraints, and every successor submitted to them.

![Reims Cathedral — four architects, two and a half centuries, one coherent result](/images/reims-cathedral.jpg)

The Bradley Fighting Vehicle is what happens when nobody with taste has authority. The Bradley began as a light troop carrier. Over 17 years and $14 billion, successive stakeholders added requirements: more armor, more firepower, more troop capacity, larger contracts. Each defensible in isolation. The accumulation produced a vehicle that Sergeant Fanning, in the 1998 HBO film *The Pentagon Wars*, describes with the clarity of someone who has stopped being polite:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

![The Pentagon Wars (1998) — 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

Every clause in that sentence names a stakeholder requirement that was reasonable alone. Together they describe a machine that did nothing well because it was asked to do everything. Testing was manipulated. The officers who rigged the results were promoted. Lt. Colonel James Burton — the one person who acted as though coherence mattered — forced an honest live-fire test and was forced into retirement. The Bradley is the limiting case of "many good but independent and uncoordinated ideas." It was not a failure of intentions. It was what happens when taste has no veto.

Brooks did not merely name the property. On IBM System/360, he, Gene Amdahl, and Gerrit Blaauw built the organizational structures that give taste authority.

**One mind — at most two in genuine resonance.** Brooks and Blaauw achieved a state in which either could speak for the architecture. "And two is indeed a magic number for collaborations; marriage was a brilliant invention and has a lot to be said for it." Three cannot do this. At three, resonance becomes negotiation, and negotiation produces settlements, not coherence.

**Real veto power.** The architect must be able to say no — "repeatedly, to smart people with good arguments" — and have it stick. Advisory veto, where a VP can override, is indistinguishable from no veto. The VP approves the feature. It ships. Coherence degrades. The VP moves on. The architect stays, holding the accumulated incoherence.

**The divorce of design from implementation.** A chapter of *The Design of Design* traces the progressive separation, beginning in the 16th century, of designing from making. On System/360, a small architecture team defined *what* the system was; a large implementation team built *how*. Conflate them and neither function works. Architects drawn into implementation stop thinking about coherence. Implementers asked to make architectural decisions optimize locally at the expense of the whole.

**Protection from organizational forces.** The architecture team was shielded from field sales, engineering, and customers — all of whom had legitimate demands, reasonable individually, fatal collectively. The shielding was a structural precondition, not a status symbol.

**Career paths that reward refusal.** Agreement builds political capital. Refusal spends it. If the incentives punish the person who says no, taste will be present in the organization and powerless within it. The Bradley is the proof: the officers who rigged the tests were promoted. The officer who forced an honest evaluation was removed.

## Open questions

Brooks's principles assume a human generating at human speed. Coding agents break that assumption. They produce tokens, not coherence. A function that compiles and passes tests can violate every design principle the surrounding codebase observes — not maliciously, but because the agent has no model of the whole and statistical prediction does not recover design intent from a context window. Some questions this raises:

**Does taste become more valuable or impossible?** An architect can veto a human's output. An architect cannot veto faster than an agent generates. If taste requires holding the whole system in your head, and the whole system is now changing at machine speed, does the role still make sense?

**Can you automate taste?** If taste is the ability to perceive conceptual integrity, and conceptual integrity is a structural property of the whole, can that perception be automated? Or is taste the thing that survives when everything else is automated — the one human capacity the agent cannot replicate because it requires a model of the system the agent cannot form?

**Who has taste in an agentic team?** If the agent writes 500 lines and the human keeps 30, who wrote the feature? If the human's contribution is selecting which output coheres with the system, the job title "prompt engineer" is underselling the role by approximately the same margin as calling it "taste."

---

Ray's experiment works because it exposes the gap between what the word suggests and what the role demands. Reformulate with the right definition: taste is the ability to perceive conceptual integrity. That makes the Bradley sentence work: the Bradley became a death trap because the people with taste had no authority. It names a faculty, a property, and a failure mode. An engineer can be held accountable for it.

Brooks gave us the property, the evidence, and the organizational conditions. The open question is whether taste — defined this way, as the perception of coherence — survives the agent era, or whether it becomes the only thing that matters.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — Conceptual integrity as the most important consideration in system design.
- Brooks, F. P. (2010). *The Design of Design.* — "Most great works have been made by one mind. The exceptions have been made by two minds." "Many hands make light work — Often. But many hands make more work — Always." The divorce of design from implementation.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)

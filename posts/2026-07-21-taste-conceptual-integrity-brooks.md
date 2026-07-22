---
title: Taste as Conceptual Integrity
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks called the real thing conceptual integrity. The Pentagon Wars shows what happens when you don't have it.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

Ray's point is that the word "taste" cannot carry the weight of engineering accountability. He is right about this, and his example makes the problem visible immediately. An inspector who makes structural judgments based on taste is not being overruled — they are being negligent. The category error is built into the word itself. Taste belongs to aesthetics and personal preference. Applying it to engineering judgment reframes a question of correctness as a question of opinion, and opinions require no rebuttal. "That's just your taste" ends the conversation without addressing the evidence.

At the same time, the thing people are groping toward when they invoke the word is not imaginary. There is a real capacity involved in looking at a system and recognizing that something is wrong beneath the surface — that the parts no longer compose, that the design has accumulated contradictions, that the whole has lost its voice. Fred Brooks spent a career giving this property a name that can carry the weight Ray's experiment demands.

## Conceptual integrity

![The Design of Design — Fred Brooks's last book, his best, and his least read](/images/design-of-design.jpg)

Brooks called it **conceptual integrity**. He first articulated the idea in *The Mythical Man-Month* in 1975, in what remains the most important paragraph ever written about system design:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas."

He returned to the subject 35 years later in *The Design of Design*, his final book and his best, and sharpened the claim further. The chapter titled "Collaboration in Design" opens with a sentence that lands like a verdict on the entire history of human making:

> "Most great works have been made by one mind. The exceptions have been made by two minds."

This is not a preference for small teams. It is a structural observation about how coherence enters a designed thing. He elaborates:

> "Many hands make light work — Often. But many hands make more work — Always."

The reasoning is precise. Every additional mind added to a design introduces new assumptions. Those assumptions must be reconciled with the existing ones. Reconciliation produces compromise. Each compromise chips at coherence. The total work may be distributed across more people, but the total work to be done increases — and the quality of the result, measured as coherence, degrades with each additional author. Brooks was not arguing against collaboration in implementation. He was arguing that design does not parallelize. You can divide the labor of building. You cannot divide the labor of deciding what the thing is.

This is not a statement about aesthetics. It is a structural claim about what makes systems maintainable over time. In a system with conceptual integrity, every interface reflects the same design philosophy. Every component uses the same idioms. Each decision is consistent with the others because someone with authority enforced that consistency — repeatedly, against smart people with good arguments, as Brooks put it. When conceptual integrity is absent, the system's parts don't compose cleanly. Its behavior surprises its users in inconsistent ways. Its maintenance costs grow faster than linear with size, because each change must reconcile assumptions that were never reconciled during the design.

Brooks's canonical example of conceptual integrity sustained across time is Reims Cathedral. Eight generations of architects — roughly two centuries — each submitting to the constraints of the original plan. The first architect's design had authority that outlived him. Hundreds of hands working across eight lifetimes produced a building that feels as though a single mind designed it. Reims should not exist. That it does is the evidence that the principle works.

## The Pentagon Wars: what happens when authority is absent

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle: 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

The counter-example is the Bradley Fighting Vehicle, the subject of the 1998 HBO film *The Pentagon Wars*. The Bradley began as a light troop carrier. Over 17 years and $14 billion, successive stakeholder groups added requirements: more armor, more firepower, more troop capacity, larger contracts. Each request was defensible when considered alone. The accumulation produced a vehicle that Sergeant Fanning, a character who has watched the program deteriorate for years, describes with the clarity of someone who has stopped being polite:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

Every clause in that sentence names a stakeholder requirement that was reasonable in isolation. Together they describe a machine that does nothing well because it was asked to do everything. The film documents that testing was manipulated to conceal the vehicle's failures; that the officers responsible for the manipulation were promoted; and that Lt. Colonel James Burton — the officer who forced an honest live-fire test, the one person in the story who acted as though coherence mattered — was eventually forced into retirement.

The Bradley is the limiting case of what Brooks described as "many good but independent and uncoordinated ideas." It was not a failure of intentions. It was not incompetence. It was the structural consequence of a system in which nobody had the authority to say no — in which every stakeholder could add a requirement and no single mind could reject one. The output was not a design. It was a negotiated settlement between competing interests, and the settlement killed soldiers.

## The organizational conditions for coherence

Brooks did not stop at naming the property. On the IBM System/360 project, he, Gene Amdahl, and Gerrit Blaauw built the organizational structures that make conceptual integrity possible. The principles are documented across both *The Mythical Man-Month* and *The Design of Design*.

**One mind — at most two in genuine resonance.** Brooks and Blaauw achieved a working relationship in which either could speak for the architecture. Brooks described this state as "genuine resonance" and added what might be the most human line in either book: "And two is indeed a magic number for collaborations; marriage was a brilliant invention and has a lot to be said for it." Three cannot do this work. At three, the dynamic shifts from resonance to negotiation, and negotiation produces settlements, not coherence.

**Real veto power.** The architect must be able to say no and have the decision hold. Brooks addressed this directly in *The Mythical Man-Month*: the designer must have the authority to reject features — "repeatedly, to smart people with good arguments." Advisory veto, in which the architect can object but a superior can override the objection, is indistinguishable in practice from having no veto at all. The VP approves the feature. The feature ships. The coherence degrades. The VP eventually moves to a different organization. The architect remains, holding the accumulated incoherence.

**The divorce of design from implementation.** In *The Design of Design*, Brooks devoted a chapter to what he called "The Divorce of Design" — the progressive separation, beginning around the 16th century, of the act of designing from the act of making. On System/360, this took the form of a small architecture team that defined what the system was and a large implementation team that built how it worked. "The architecture team must be protected; the implementation team must be coordinated." Organizations that conflate these two functions discover that neither is performed adequately. The architects, drawn into implementation, stop thinking about systemic coherence. The implementers, asked to make architectural decisions, optimize locally at the expense of the whole.

**Protection from organizational forces.** The architecture team on System/360 was deliberately shielded from field sales (who wanted features for specific customers), from engineering (who wanted optimizations that would have compromised clean abstractions), and from customers (who demanded compatibility with existing systems). Every one of these demands was reasonable. Accommodating any significant fraction would have destroyed the system's conceptual integrity. The shielding was not a matter of status. It was a structural precondition for coherent design.

**Career paths that reward refusal.** This is the principle Brooks implied but never stated as bluntly as the evidence warrants. In most organizations, agreement builds political capital and refusal spends it. If the incentive structure penalizes the person who protects coherence, coherence will not be protected regardless of what the documentation claims. The Bradley program is the proof at the limit: the officers who rigged the tests were promoted. The officer who forced an honest evaluation was removed.

## The agent era

Coding agents, as they exist now, generate output with no regard for a system's conceptual integrity. This is not a defect. It is a consequence of how they function. They produce tokens that are statistically probable given their training distribution and the contents of their context window. They have no model of the system as a whole, and they cannot acquire one, because the coherence of a codebase across its entire surface area is not a statistical property that can be recovered from token prediction. It is a design property, and design properties must be imposed. They do not emerge.

An agent will produce a function that compiles, passes the tests it was given, and violates the design philosophy that every other function in the surrounding module respects. It will do this without awareness, because the design philosophy is not present in the text it was trained on in any form the model can extract. The agent's context window contains the text of nearby files. It does not contain the design rationale that produced them. It cannot contain it, because that rationale was never written down — it lived in the head of the person who had the authority to say no.

The implication is that the human role in software development converges toward the role Brooks described for the architect on System/360. The primary contribution becomes maintaining conceptual integrity: enforcing the design philosophy the agent cannot perceive, rejecting contributions that work locally but break coherence globally, and making the decisions about what the system will not do. Brooks argued that the one-mind rule was the hardest organizational principle to maintain. The agent era makes it harder, because the forces that erode conceptual integrity now operate at the speed of token generation. The architect is defending the system against a force that produces incoherence faster than any human can review it.

---

Ray's experiment succeeds because it exposes the gap between what the word "taste" can carry and what the role actually demands. "The building fell down after the inspector's taste was ignored" does not describe an overruled professional. It describes a negligent one. That is the entire problem.

But reformulate the sentence with the right term. The Bradley became a death trap because no single mind had the authority to enforce the vehicle's conceptual integrity. That sentence works. It names a role — someone with the authority to maintain coherence. It names a property — the system speaking with one voice across every component and every decision. And it names a failure mode — the property was absent because the role was absent, and reasonable people making reasonable requests, unchecked by anyone with the power to refuse them, produced a machine that did nothing well.

"Taste" cannot carry that weight. It was never designed to. Brooks gave us a term that can, and a set of organizational principles for making it real. The remaining question is whether we have the nerve to use them.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — The original formulation: conceptual integrity as the most important consideration in system design; the one-mind rule; the architect's authority to say no.
- Brooks, F. P. (2010). *The Design of Design.* — "Most great works have been made by one mind. The exceptions have been made by two minds." "Many hands make light work — Often. But many hands make more work — Always." The divorce of design from implementation; the collaboration chapter; the mature statement of a career's worth of thinking about coherence.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first) — What OpenBSD's approach to system coherence teaches about the authority to say no.

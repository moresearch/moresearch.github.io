---
title: "Taste" Is a Terrible Word for What We Mean
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks called the real thing conceptual integrity. The Pentagon Wars shows what happens when you don't have it.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

Ray's point is that the word "taste" cannot carry the weight of engineering accountability. He is right about this, and his choice of example makes the problem clear. An inspector who makes structural judgments based on taste is not being overruled — they are being negligent. The word belongs to aesthetics and personal preference. Applying it to engineering judgment reframes a question of correctness as a question of opinion, and opinions are easy to dismiss without evidence.

At the same time, the thing people are reaching for when they use the word is not imaginary. There is a real skill involved in looking at a system and recognizing that something is wrong — that the parts don't fit together, that the design has accumulated contradictions, that the whole no longer speaks with a single voice. Fred Brooks gave this property a name that can carry the weight Ray's thought experiment demands.

## Conceptual integrity

![The Design of Design — Fred Brooks's last book, his best, and his least read](/images/design-of-design.jpg)

In *The Design of Design*, Brooks defined conceptual integrity as the property of a system that feels like one mind designed it. He made this the central claim of his career:

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — Brooks, *The Mythical Man-Month*, 1975

This is not a statement about aesthetics. It is a structural claim about what makes systems maintainable over time. In a system with conceptual integrity, every interface reflects the same design philosophy. Every component uses the same idioms. Each decision is consistent with the others because someone with authority ensured that consistency. When conceptual integrity is absent, the system's parts don't compose cleanly. Its behavior surprises its users in inconsistent ways. Its maintenance costs grow faster than linear with size, because each change must reconcile assumptions that were never reconciled during design.

Brooks's canonical example of conceptual integrity sustained over time is Reims Cathedral, built across eight generations of architects — roughly two centuries — with each successive architect submitting to the constraints of the original plan. The result is a building that feels as though a single designer produced it, despite hundreds of hands working across eight lifetimes. The plan had authority that outlived its author.

## The Pentagon Wars: what happens when nobody has that authority

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle: 17 years, $14 billion, and a machine that did nothing well](/images/pentagon-wars.jpg)

The counter-example is the Bradley Fighting Vehicle, the subject of the 1998 HBO film *The Pentagon Wars*. The Bradley began as a light troop carrier. Over the course of 17 years and at a cost of $14 billion, successive groups of stakeholders added requirements. Armor advocates wanted more plating. Infantry commanders wanted more troop capacity. Generals wanted more firepower. Contractors wanted larger contracts. Each of these requests was defensible when considered in isolation. The accumulation produced a vehicle that Sergeant Fanning, a character in the film who has spent years watching the program deteriorate, describes this way:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

Every clause in that sentence represents a stakeholder requirement that was reasonable on its own. Together they describe a machine that does nothing well because it was asked to do everything. The film also documents that testing was manipulated to conceal the vehicle's failures, that the officers responsible for the manipulation were promoted, and that Lt. Colonel James Burton — the officer who forced an honest live-fire test — was eventually forced into retirement.

The Bradley program exhibits the structural failure that conceptual integrity is designed to prevent. It was not a case of bad intentions or individual incompetence. It was a case of reasonable people making reasonable requests in a system where nobody had the authority to say no. The result was not a coherent design but a negotiated settlement between stakeholders — and the settlement killed people.

## How Brooks organized against this failure mode

Brooks did not merely name the property. On the IBM System/360 project, he, Gene Amdahl, and Gerrit Blaauw implemented a set of organizational structures designed to preserve conceptual integrity against the forces that erode it.

The first is what Brooks called the **one-mind rule**: the conceptual design of a system must be controlled by one person, or at most two people in what he termed "genuine resonance" — a pair who share a mental model so completely that either can speak for the architecture. Brooks and Blaauw achieved this. Three people cannot, because at three the dynamic shifts from resonance to negotiation, and negotiation produces compromise rather than coherence.

The second is **real veto power**. The architect must be able to say no to a feature — "repeatedly, to smart people with good arguments," as Brooks put it — and have the decision stand. Advisory veto, where the architect can object but a vice president can override the objection, produces the same result as having no veto at all. The VP approves the feature, the feature ships, the user experience degrades, the VP eventually moves to a different role, and the architect is left with the accumulated incoherence.

The third is the **separation of architecture from implementation**. A small team defines what the system is — its interfaces, its abstractions, its guarantees. A larger team builds how the system works beneath those interfaces. "The architecture team must be protected; the implementation team must be coordinated." Organizations that conflate these two functions find that neither is done well. The architecture team, drawn into implementation details, stops thinking about coherence. The implementation team, asked to make architectural decisions, produces decisions that serve local convenience rather than system-wide consistency.

The fourth is **protection from organizational forces**. On System/360, the architecture team was deliberately shielded from field sales (who wanted features for specific customers), from engineering (who wanted optimizations that would have compromised clean abstractions), and from customers (who demanded compatibility with their existing systems). Each of these demands was reasonable. Accommodating any significant fraction of them would have destroyed the system's conceptual integrity. The shielding was not a matter of arrogance — it was a structural requirement for the design work to proceed coherently.

The fifth is harder to implement but follows from the previous four: **career paths that reward saying no**. In most organizations, agreeing to requests builds political capital and refusing them spends it. If the incentive structure penalizes the person who protects coherence, coherence will not be protected regardless of what the org chart says. The Bradley program is the limiting case: the officers who manipulated the tests were promoted, and the officer who forced an honest evaluation was forced out of the service.

## The agent era

Coding agents, as they exist today, produce output with no regard for a system's conceptual integrity. This is not a criticism of the agents — it is a statement about how they work. They generate tokens that are statistically probable given their training data and context window. They do not form a model of the system as a whole, and they cannot, because coherence across a codebase is not a statistical property that emerges from token prediction. It is a design property that must be imposed by someone who holds the whole system in their head.

An agent will write a function that compiles, passes the provided tests, and violates the design philosophy that every other function in the surrounding module observes. It will do this not because it is careless but because it has no way to know that the design philosophy exists. The agent's context window contains the text of nearby files, not the design rationale that produced them.

The implication is that the human role in software development shifts toward something that looks very much like Brooks's architect role. The primary contribution becomes maintaining the system's conceptual integrity — enforcing the design philosophy the agent cannot perceive, rejecting contributions that work individually but break coherence in combination, and making the decisions about what the system should not do. Brooks argued that the one-mind rule was the hardest organizational principle to sustain. The agent era makes it more important, because the forces that erode conceptual integrity now operate at the speed of token generation.

---

Ray's thought experiment works because it exposes the gap between what the word "taste" can carry and what the role actually demands. "The building fell down after the inspector's taste was ignored" describes a negligent inspector, not an overruled professional. That is the whole problem with the word.

But try the formulation with a different term: the Bradley became a death trap because nobody had the authority to enforce the vehicle's conceptual integrity. That sentence works. It identifies a role — someone with the authority to enforce coherence. It identifies a property — conceptual integrity, the system speaking with one voice. And it identifies a failure mode — the property was absent because the role was absent, and the result was a vehicle that did nothing well.

"Taste" was never the right word for that. Brooks gave us the right one.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (2010). *The Design of Design.* — Conceptual integrity, the one-mind rule, the separation of architecture from implementation, and the organizational protections required to sustain them.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — The original articulation of conceptual integrity as the most important consideration in system design.
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first) — What OpenBSD's approach to system coherence teaches about the authority to say no.

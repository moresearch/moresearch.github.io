---
title: "Taste" Is a Terrible Word for What We Mean
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks called the real thing conceptual integrity. The Pentagon Wars shows what happens when you don't have it.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

The sentence doesn't work. "Taste" means personal preference. An inspector who made structural judgments based on taste would be negligent, not overruled. The word can't carry engineering accountability because it was never meant to.

But the thing people are fumbling toward when they say "taste" is real. Fred Brooks gave it a name that can carry the weight.

## Conceptual integrity

![The Design of Design — Fred Brooks's last book, his best, and his least read](/images/design-of-design.jpg)

Brooks defined **conceptual integrity** as the property of a system that feels like one mind designed it. Not one team. Not one process. One mind. Every interface reflects the same design philosophy. Every component uses the same idioms. Every decision is consistent with every other decision.

> "I will contend that conceptual integrity is the most important consideration in system design. It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — Brooks, *The Mythical Man-Month*, 1975

This is not a preference. It is a structural claim about what makes systems coherent. A system without conceptual integrity is not ugly — it is disorganized. Its parts don't compose. Its behavior surprises its users. Its maintenance costs grow non-linearly because each change must reconcile conflicting assumptions.

Brooks's positive case is Reims Cathedral. Eight generations of architects, each submitting to the original plan. Two centuries. One coherent result. The plan had authority that outlived its author.

## What happens when nobody has that authority

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle: 17 years, $14 billion, satisfied every stakeholder and no user](/images/pentagon-wars.jpg)

The counter-example is the Bradley Fighting Vehicle, immortalized in the 1998 HBO film *The Pentagon Wars*. The Bradley began as a light troop carrier. Over 17 years and $14 billion, armor advocates, infantry commanders, generals, and contractors each added requirements. Every request was defensible in isolation. Accumulated, they produced a vehicle that the film's Sergeant Fanning describes with the precision of a man who has spent years watching the disaster unfold:

> "A troop transport that can't carry troops, a reconnaissance vehicle that's too conspicuous to do reconnaissance, and a quasi-tank that has less armor than a snowblower, but has enough ammo to take out half of D.C."

Each clause in that sentence is a stakeholder requirement that was reasonable on its own. Together they describe a machine that does nothing well because it was asked to do everything.

Testing was manipulated. Officers who rigged results were promoted. The whistleblower who forced an honest test — Lt. Colonel James Burton — was forced into retirement.

> This is what happens when conceptual integrity is absent. Not bad intentions. Not incompetence. Reasonable people making reasonable requests, with nobody authorized to say no. The system degrades into a negotiated settlement between stakeholders rather than a coherent design.

## Brooks's organizational protections

Brooks didn't just name the property. On System/360, he, Gene Amdahl, and Gerrit Blaauw implemented the organizational structure that preserves it:

- **One mind — at most two in genuine resonance.** Brooks and Blaauw could each speak for the architecture. Three people cannot do this. Three is a committee.

- **Veto power, not advisory input.** The architect must be able to say no and have it stick. Advisory veto means the VP overrides the architect. The feature ships. The user suffers. The VP moves on. The architect inherits the incoherence.

- **Architecture separated from implementation.** A small team defines *what* the system is. A large team builds *how*. "The architecture team must be protected; the implementation team must be coordinated." Conflate the two and neither function works.

- **Protection from organizational forces.** Field sales, engineering, and customers all had legitimate demands on System/360. The architecture team was shielded from them. Each demand was reasonable. Accommodating all of them would have destroyed the system's coherence.

- **Career paths that reward saying no.** In most organizations, "yes" gets promoted. Saying no costs political capital. If the incentive structure punishes coherence, coherence will not occur. *The Pentagon Wars* is the proof: the saboteurs were promoted, the honest officer was purged.

## Why this matters now

Coding agents produce output with no regard for the system's conceptual integrity. They generate tokens that are statistically probable, not architecturally coherent. An agent will write code that compiles and passes tests while violating every design principle the surrounding codebase observes — because the agent has no model of the whole. It can't. Coherence is not a statistical property.

The human's role is not to apply "taste" to the agent's output — as if the job were selecting the most aesthetically pleasing token sequence. The human's role is to maintain the system's conceptual integrity against a force that generates incoherence at machine speed. Brooks's one-mind rule was always hard. In the agent era, it becomes the primary thing humans contribute.

---

Ray's experiment proves his point. "The building fell down after the inspector's taste was ignored" is nonsense — it describes a negligent inspector, not an overruled one. But try this: "The Bradley became a death trap because nobody had the authority to enforce conceptual integrity." That sentence names a role, a property, and a failure mode. An engineer can be held accountable for it.

> "Taste" can't do that work. It was never supposed to.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (2010). *The Design of Design.* — Conceptual integrity, the one-mind rule, protecting the designer, architecture/implementation separation.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — The original articulation: "I will contend that conceptual integrity is the most important consideration in system design."
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first)

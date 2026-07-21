---
title: "Taste" Is a Terrible Word for What We Mean
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks had a better word: conceptual integrity. The Pentagon Wars shows what happens when you don't have it.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

He's right. "Taste" is a terrible word. It belongs to wine and font choices, not engineering judgment that determines whether a system stands or falls. When you call engineering judgment "taste," you make it sound like preference. And preference is easy to dismiss — "that's just your taste" is a conversation-ender. It requires no rebuttal because it treats the objection as opinion rather than evidence.

But the fact that "taste" can't carry the weight of accountability doesn't mean the thing people are reaching for doesn't exist. It means we've been using the wrong word. Fred Brooks had the right one.

## What people mean when they say "taste"

In *The Design of Design*, Brooks called it **conceptual integrity**: the system feels like one mind designed it. It is the most important property of any designed thing.

Conceptual integrity is not personal preference or aesthetics. It is the structural property that makes a system coherent rather than assembled — every interface reflecting the same design philosophy, every component using the same idioms, every decision consistent with every other decision because one mind had authority over the whole.

> "It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — Brooks, *The Mythical Man-Month*, 1975

Brooks's positive example is Reims Cathedral. Built over eight generations of architects, each stuck to the original plan. The result: a unified work, coherent in every detail, as if one mind designed it. The first architect's plan had authority that outlived him. Two centuries of successors submitted.

## The Pentagon Wars: what happens when nobody has authority

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle took 17 years, cost $14 billion, and pleased no one](/images/pentagon-wars.jpg)

The counter-example to Reims Cathedral is the Bradley Fighting Vehicle, immortalized in the 1998 HBO film *The Pentagon Wars*. The Bradley began as a light troop carrier. Over 17 years and $14 billion, it became — as the film's Colonel Smith describes it — "a bulky tank-like vehicle poorly suited to its original role."

What happened? Every stakeholder added requirements. Armor advocates wanted more plating. Infantry wanted more troop capacity. Generals wanted more firepower. Contractors wanted more budget. Not one of these requests was unreasonable on its own. Together, they destroyed the vehicle's coherence. The Bradley attempted to be a troop carrier, a reconnaissance vehicle, and a light tank simultaneously. It did none of these jobs well.

> This is design by committee: every stakeholder gets their feature, nobody has the authority to say no, and the result satisfies everyone's wish list and nobody's actual needs.

The film's protagonist, Lt. Colonel James Burton, discovers that testing was routinely manipulated to hide the vehicle's failures. Master Sergeant Dalton, who ran live-fire tests, admits he was ordered to rig the results. When Burton finally forces an honest test, the Bradley is destroyed — literally and reputationally. The postscript notes that a redesigned Bradley served effectively in the Gulf War, while "most of the officers involved earned promotions" and Burton was forced into retirement.

The Bradley is what Reims Cathedral would have been if each successive architect had insisted on adding a Baroque chapel, a Gothic spire, and a Renaissance facade to satisfy their patrons. The cathedral would have stood — and meant nothing. The Bradley stood — and soldiers died because of it.

## How Brooks said to prevent this

Brooks didn't just diagnose the problem. He prescribed specific organizational protections for the designers who maintain conceptual integrity.

**One mind — at most two in resonance.** The design must proceed from one person, or a pair who share a mental model so deeply either can speak for the architecture. Brooks found this with Gerrit Blaauw on System/360. Three cannot do it. Three is already a committee.

**Real veto power, not advisory.** The designer must be able to say no — "repeatedly, to smart people with good arguments" — and make it stick. Advisory veto means the VP overrides the designer. The VP wins. The user loses. The feature ships. Nobody uses it. The VP moves to a new role. The designer inherits the mess.

**Architecture separated from implementation.** A small team defines *what*. A large team builds *how*. Roles distinct, staffed differently. The architecture team needs protection from external pressure. The implementation team needs coordination. Conflate them: nobody does either well.

**Protection from organizational forces.** On System/360, Brooks, Amdahl, and Blaauw controlled the architecture. But they also had organizational backing to exercise that control. Field sales wanted features for customers. Engineering wanted optimizations. Customers demanded compatibility. Each request was reasonable alone. Together, they would have destroyed coherence. The team was shielded from those forces.

> "The architecture team must be protected; the implementation team must be coordinated." — Brooks, *The Design of Design*

**The designer needs a career path that rewards saying no.** In most organizations, "yes" gets you promoted. "No" costs political capital. If the incentive structure punishes coherence, coherence will not happen. The Pentagon Wars is the proof: the officers who destroyed the Bradley were promoted. The one who insisted on honesty was forced out.

## The bridge: taste as fast orientation

In the comments on Ray's post, Sean Cooper offered a nuanced take: taste is real but insufficient. It's "fast orientation" — the senior engineer's "this feels wrong" that compresses decades of pattern recognition into a moment. But it must be converted into evidence: name the odor, articulate the heuristic, demonstrate the risk.

This bridges Brooks and Ray. The senior engineer who says "this feels wrong" is doing what Brooks described as the designer's essential skill: seeing the whole and detecting when a part doesn't fit. But the feeling is not the argument. The feeling is the starting point. The argument must be: this breaks coherence in *this* specific way, with *these* specific consequences.

> Taste is fast orientation. Conceptual integrity is the slow, structural work of making the system coherent. You need both. You can only be held accountable for the second.

## The agent era makes this urgent

Agents have zero conceptual integrity. They generate tokens statistically, not coherently. An agent can produce code that compiles, passes tests, and violates every design principle the codebase follows — not maliciously, but because it has no model of the whole.

The human's role is to be the one mind. To enforce the design philosophy the agent can't see. To reject code that works but doesn't fit. To say no to good ideas that would break coherence. Brooks's one-mind rule was always important. In the agent era, it becomes the primary human contribution.

> Calling that role "taste" doesn't just undersell it. It makes it sound optional. Conceptual integrity is not optional. Systems without it accumulate complexity until they cannot be changed. Agents accelerate that accumulation because they produce code faster than anyone can review it for coherence.

---

Ray's experiment proves his point. "The building fell down after the inspector's taste was ignored" sounds absurd because it is. But "the Bradley became a death trap because nobody had the authority to enforce conceptual integrity" — that sentence works. It names a role, a skillset, and a failure mode. It is a sentence an engineer can be held accountable for.

Brooks gave us the framework. Reims Cathedral and the Pentagon Wars are the proof cases — one positive, one negative, same principle. The organizations that protect their designers produce Reims. The organizations that let committees accumulate requirements produce the Bradley. Same dynamics, different stakes, same outcome: coherence or chaos.

> "Taste" can't carry that weight. Conceptual integrity can.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (2010). *The Design of Design.* — Conceptual integrity, the one-mind rule, protecting the designer, architecture/implementation separation.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — "It is better to have a system omit certain anomalous features... than to have one that contains many good but independent and uncoordinated ideas."
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. — Based on Col. James G. Burton's book *The Pentagon Wars: Reformers Challenge the Old Guard.*
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first) — What OpenBSD teaches about coherence, discipline, and the authority to say no.

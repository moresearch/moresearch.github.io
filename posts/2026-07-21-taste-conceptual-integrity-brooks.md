---
title: "Taste" Is a Terrible Word for What We Mean
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks had a better word: conceptual integrity. The Pentagon Wars shows what happens when you don't have it.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

He's right. "Taste" belongs to wine and font choices, not to the engineering judgment that determines whether a system stands or falls. It makes accountability sound like preference, and preference is easy to dismiss: "that's just your taste" is a conversation-ender.

But the thing people are reaching for is real. They're just using the wrong word.

## Brooks had the right word

![The Design of Design — Fred Brooks's last book, his best, and his least read](/images/design-of-design.jpg)

In *The Design of Design*, Fred Brooks called it **conceptual integrity**: the system feels like one mind designed it. It's the most important property of any designed thing — and it is not aesthetics. It's the structural property that makes a system coherent rather than assembled.

> "It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — Brooks, *The Mythical Man-Month*, 1975

Brooks's positive example is Reims Cathedral. Eight generations of architects, each sticking to the original plan. Two centuries. One coherent result.

## The Pentagon Wars: what happens when nobody can say no

![The Pentagon Wars (1998) — the Bradley Fighting Vehicle: 17 years, $14 billion, and pleased no one](/images/pentagon-wars.jpg)

Now consider the Bradley Fighting Vehicle, immortalized in the 1998 HBO film *The Pentagon Wars*. It began as a light troop carrier. Over 17 years and $14 billion, every stakeholder added requirements — more armor, more firepower, more troop capacity, more budget. Each request was reasonable alone. Together, they produced a vehicle the film describes as "a bulky tank-like vehicle poorly suited to its original role."

The testing was rigged to hide the failures. The officers who manipulated results were promoted. The whistleblower, Lt. Colonel James Burton, was forced into retirement.

> This is design by committee: every stakeholder gets their feature, nobody has authority to say no, and the result satisfies everyone's wish list and nobody's actual needs. The Bradley is what Reims Cathedral would have been if each architect had added a Gothic spire, a Baroque chapel, and a Renaissance facade to please their patrons.

## How Brooks said to prevent this

Brooks didn't just diagnose the problem. On System/360, he, Amdahl, and Blaauw lived the solution:

- **One mind — at most two in resonance.** Three is already a committee.
- **Real veto power.** Advisory veto means the VP overrides the designer. The VP wins. The user loses. The feature ships. Nobody uses it.
- **Architecture separated from implementation.** Small team defines *what*. Large team builds *how*. Conflate them and neither works.
- **Protection from organizational forces.** Field sales, engineering, and customers all had legitimate demands. The architecture team was shielded from them — otherwise each reasonable request would have destroyed coherence.
- **Career paths that reward saying no.** "Yes" gets promoted. "No" costs political capital. If the incentives punish coherence, coherence won't happen. The Pentagon Wars is the proof.

> "The architecture team must be protected; the implementation team must be coordinated." — Brooks

## The bridge: taste as fast orientation

Sean Cooper, in the comments on Ray's post, got the nuance right: taste is "fast orientation" — the senior engineer's "this feels wrong" that compresses decades of pattern recognition into a moment. It's real. But it must be converted into evidence. Name the odor. Articulate the heuristic. Demonstrate the risk.

> Taste is fast orientation. Conceptual integrity is the slow, structural work of making the system coherent. You need both. You can only be held accountable for the second.

## The agent era makes this urgent

Agents have zero conceptual integrity. They generate tokens statistically. An agent writes code that compiles and passes tests while violating every design principle the codebase follows — not maliciously, but because it has no model of the whole.

The human's role is to be the one mind. To reject code that works but doesn't fit. To say no to good ideas that would break coherence. Brooks's one-mind rule was always important. In the agent era, it's the primary human contribution.

---

Ray's experiment is the test. "The building fell down after the inspector's taste was ignored" — absurd. "The Bradley became a death trap because nobody had the authority to enforce conceptual integrity" — that sentence lands. It names a role, a skillset, and a failure mode. An engineer can be held accountable for it.

> "Taste" can't carry that weight. Conceptual integrity can.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn.
- Brooks, F. P. (2010). *The Design of Design.* — Conceptual integrity, the one-mind rule, protecting the designer.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — "It is better to have a system omit certain anomalous features..."
- *The Pentagon Wars* (1998). Dir. Richard Benjamin. HBO. — Based on Col. James G. Burton's book.
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) · [one-mind rule](https://blog.hackspree.com/#brooks-design-one-mind-rule) · [protecting the designer](https://blog.hackspree.com/#brooks-design-protecting-designer)
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first)

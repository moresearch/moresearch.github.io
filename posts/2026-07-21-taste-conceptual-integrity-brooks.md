---
title: "Taste" Is a Terrible Word for What We Mean
date: 2026-07-21
slug: taste-conceptual-integrity-brooks
summary: Ray Myers is right that "taste" belittles engineering judgment. Fred Brooks had a better word: conceptual integrity. The difference matters more in the agent era than ever.
tags: design, fred-brooks, conceptual-integrity, taste, engineering-judgment, agents
---

Ray Myers posted a thought experiment on LinkedIn this week:

> "Imagine a phrase like 'The building fell down after the inspector's taste was ignored.' Is there any situation where you would feel accountable to heed someone's taste?"

He's right. "Taste" is a terrible word for what we mean. It belongs to wine, to font choices, to whether a photograph looks good in black and white. It does not belong to engineering judgment that determines whether a system stands or falls.

But Ray's experiment also reveals something deeper. The fact that "taste" can't carry the weight of engineering accountability doesn't mean the thing people are reaching for doesn't exist. It means we've been using the wrong word.

Fred Brooks had the right word.

## The thing people mean when they say "taste"

In *The Design of Design*, Brooks called it **conceptual integrity**: "the system feels like one mind designed it." It is the most important property of any designed system. Everything else serves it.

Conceptual integrity is not personal preference. It is not aesthetics. It is the structural property that makes a system coherent rather than assembled — where every interface reflects the same design philosophy, every component uses the same idioms, and every decision is consistent with every other decision because one mind (or one team acting as one mind) had authority over the whole.

> "It is better to have a system omit certain anomalous features and improvements, but to reflect one set of design ideas, than to have one that contains many good but independent and uncoordinated ideas." — Fred Brooks, *The Mythical Man-Month*, 1975

That is not taste. That is design discipline.

Brooks's proof is Reims Cathedral. Built over eight generations of architects, each stuck to the original plan. The result: a unified work, coherent in every detail, as if one mind designed it. Most cathedrals are not like this. Gothic nave, Renaissance facade, Baroque chapel on Romanesque transept — each generation had a vision, nobody had authority, and the result was chaos.

Software systems are cathedrals built over decades. The ones that last — Unix under Thompson and Ritchie, Go under Thompson, Pike, and Griesemer, the Macintosh under Jobs — have conceptual integrity. One mind, or one resonant pair, had the authority to say no. The ones that don't have conceptual integrity are Linux distributions assembled from a hundred projects with a hundred design philosophies — functional, full of life, and impossible to understand as a whole.

## Why the word matters

Ray's concern is not semantic. It's about power.

When you call engineering judgment "taste," you make it sound like preference. And preference is easy to dismiss. "That's just your taste" is a conversation-ender. It requires no rebuttal because it treats the objection as a matter of opinion rather than evidence.

This is dangerous in any engineering context. It's lethal in the context Ray is responding to: an industry where some voices argue that the human contribution to AI-assisted development is "taste" — as if the human's role is to look at the agent's output and decide whether they like it.

> "The building fell down after the inspector's taste was ignored" sounds absurd because it is absurd. The inspector doesn't have taste. The inspector has expertise, heuristics, evidence, and the authority to say no. So does the senior engineer reviewing agent-generated code.

Brooks understood this distinction. Conceptual integrity is not asserted. It is enforced. The designer must have the authority to say no — "repeatedly, to smart people with good arguments" — and make it stick. That authority is not granted by having better taste. It is granted by having a coherent design vision and the track record to justify it.

## What Sean Cooper got right

In the comments on Ray's post, Sean Cooper offered a nuanced take: taste is real but insufficient. It's "fast orientation" — the senior engineer's "this feels wrong" that compresses decades of pattern recognition into a single moment. But taste must be converted into evidence. Name the odor. Articulate the heuristic. Demonstrate the risk.

This is exactly the bridge between Brooks's conceptual integrity and Ray's accountability.

The senior engineer who says "this feels wrong" is drawing on the same pattern-matching system that Brooks described as the designer's essential skill: the ability to see the whole and detect when a part doesn't fit. But the feeling is not the argument. The feeling is the starting point for the argument. The argument must be: this breaks coherence in this specific way, with these specific consequences.

> Taste is fast orientation. Conceptual integrity is the slow, structural work of making the system coherent. You need both. You can only be held accountable for the second.

## The agent era makes this urgent

Agents have no conceptual integrity. Zero. They generate tokens statistically, not coherently. An agent can write a function that compiles and passes tests and violates every design principle the rest of the codebase follows — not because it's malicious, but because it has no model of the whole.

The human's role is not to apply "taste" to the agent's output. The human's role is to be the one mind that maintains conceptual integrity — to enforce the design philosophy the agent can't see, to reject code that works but doesn't fit, to say no to good ideas that would break coherence.

> Calling that role "taste" doesn't just undersell it. It makes it sound optional. Conceptual integrity is not optional. Systems without it accumulate complexity until they cannot be changed. The agent era accelerates that accumulation because agents produce code faster than humans can review it for coherence.

If we call it taste, we lose. If we call it conceptual integrity, we have a framework, a tradition, and a standard of accountability that stretches from Reims Cathedral to the Unix kernel.

## The right word

Ray's experiment proves his point. Nobody would feel accountable to heed someone's taste. The word doesn't carry the weight. It can't.

But the thing the industry is fumbling toward — the human contribution that survives when code generation is commoditized — is real. It's the ability to see the whole, to enforce coherence, to say no to features that don't fit, and to be held accountable when the system fails because those decisions were wrong.

Brooks called it conceptual integrity. The word is five syllables and not catchy. It will never trend on LinkedIn. But it has the property that matters most: it can carry the weight of accountability that "taste" cannot.

> "The system collapsed because the engineer with conceptual integrity was ignored." That sentence works. It names a role, a skillset, and a failure mode. It is a sentence an engineer can be held accountable for. "Taste" can't do that work. Conceptual integrity can.

---

**References:**

- Myers, R. (2026). ["Let's try an experiment to see if we want to embrace 'taste' as the term for engineering judgment."](https://www.linkedin.com/posts/cadrlife_lets-try-an-experiment-to-see-if-we-want-share-7485045944925720576-WlW5) LinkedIn. — The post this essay responds to.
- Brooks, F. P. (2010). *The Design of Design.* — The source of "conceptual integrity" as the most important property of any designed system.
- Brooks, F. P. (1975). *The Mythical Man-Month.* — Where Brooks first articulated the argument: it is better to omit features and preserve coherence than to include many good but uncoordinated ideas.
- Related: [Brooks on Software Design: conceptual integrity](https://blog.hackspree.com/#brooks-design-conceptual-integrity) — This blog's earlier treatment of Brooks's framework and Reims Cathedral as proof.
- Related: [Correctness First](https://blog.hackspree.com/#correctness-first) — What OpenBSD teaches about coherence, discipline, and the authority to say no.

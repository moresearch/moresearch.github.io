---
title: Two ways to design software
date: 2026-07-12
slug: two-ways-to-design
summary: "Tony Hoare: 'There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult.' The difficulty is economics. The simplicity costs more now. The complexity costs more later."
tags: hoare, simplicity, complexity, design, economics
---

Tony Hoare, in his 1980 Turing Award lecture, described two approaches to software design:

> "There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult."

The sentence is famous. It is usually quoted for its elegant symmetry — two ways, two outcomes, one harder than the other. It is less often quoted for its economic content. The economic content is the point. The first method is more difficult. Difficulty in engineering translates to cost. The simple design costs more to produce. The complex design costs less to produce and more to maintain. The trade-off is economic. The choice between them is a choice about when to pay.

## The simple way

A simple design is one where the structure is transparent. A reader can see what the system does and how it does it. The components have clear responsibilities. The interfaces are minimal. The interactions are predictable. The system has no unnecessary parts.

Producing such a design requires understanding the problem deeply enough to identify what is unnecessary. The understanding cannot be acquired by thinking. It must be acquired by building. The first version will be complex because the first version is where you learn what the problem actually is. The simple design is the second version, or the third, or the fifth. It is the version produced after you have built the complex one, understood it, and removed everything that wasn't earning its place. The removal is the difficulty. The removal requires judgment about what is essential and what is accidental. The judgment cannot be automated. It requires taste. Taste is expensive to develop. The expense is time spent building complex things and simplifying them. The expense is the tuition for the simple design.

Perlis said the same thing in different words: "Simplicity does not precede complexity, but follows it." The simple design follows the complex one. The time spent on the complex one is the cost of the simple one. The cost is invisible in the final product. The final product looks like it was designed simply from the start. It was not. It was simplified. The simplification was work. The work was expensive.

## The complicated way

A complicated design is one where the structure is opaque. The system works. It has been tested. It passes its tests. But a reader cannot easily see *why* it works. Components have overlapping responsibilities. Interfaces have grown extra parameters over time — "just add a flag" is the entropy mechanism of software. Interactions have edge cases that are handled but not documented. The system has many parts. The parts interact in ways that surprise even the people who built them.

Producing such a design is easy. You build what works. You add what's needed. You don't remove what isn't. The removal is the hard part. Skipping it makes the design complicated. The complexity is not malicious. It is the natural state of a system that has been changed by many people over many years, each adding what they needed and nobody removing what was no longer necessary. The removal requires knowing what is no longer necessary. The knowing is distributed across the team, the codebase, the incident history. It is Hayek's dispersed knowledge, applied to a single system. Centralizing it is hard. The difficulty is why the removal doesn't happen.

> "Inside every large program, there is a small program trying to get out." — Tony Hoare

The small program is the simple design that would have sufficed. It is buried under the accumulated weight of features that were added because they were easier to add than to integrate properly, workarounds that were applied because the root cause was too expensive to fix, abstractions that were generalized prematurely because generality felt like good design. The small program is still there. It is obscured. The obscuring material is the complexity. The complexity is the cost of decisions made under time pressure. The time pressure was economic. The decisions were economic. The complexity is economic debt.

## The economics of the choice

The choice between the two methods is not aesthetic. It is economic. The simple method costs more now. The complicated method costs more later. The choice is about the discount rate — how much you value the present relative to the future.

A team with a high discount rate — next sprint's features matter more than next year's maintainability — will choose the complicated method. The choice is rational given the incentive structure. The incentive structure rewards velocity now. It does not reward maintainability later. The engineer who spends two weeks simplifying a design that already works is less visibly productive than the engineer who ships two features in the same time. The simplification prevents future problems. The prevention is invisible. The features are visible. The visible gets rewarded. The invisible doesn't. The incentive structure produces complicated designs. The structure is the problem.

A team with a low discount rate — sustainability matters, the system will exist for years, the cost of future complexity is priced into present decisions — will choose the simple method. The choice is also rational. These teams are rare. They are rare because low discount rates require organizational stability — the same people maintaining the system they built, long enough to feel the cost of their own complexity. If you build a complicated system and leave before the complexity costs you, you benefited from the speed and didn't pay the maintenance cost. The cost was paid by your replacement. Your incentive was to build complicated. The incentive was structural. The structure produced the behavior.

## The false choice

The choice between simple and complicated is sometimes presented as a choice between elegance and pragmatism. Elegance is for academics. Pragmatism ships. This framing is wrong. The pragmatic choice is usually the simple one, if the time horizon is long enough. The complicated design is pragmatic only on a short horizon. On a long horizon, the complicated design is the most expensive choice you can make. The expense is deferred, compounded, invisible in the current sprint, undeniable in year five. The pragmatism that ignores the future is not pragmatism. It is myopia with a professional vocabulary.

Hoare's sentence is not a preference for elegance. It is a statement of economic fact. The simple design is more difficult — costs more now. The complicated design is easier — costs more later. Choose. The choice is yours. The structure of your incentives will make it for you if you don't make it consciously. Conscious choices are better than structural ones. Structural ones feel like they weren't choices at all. They were. The structure disguised them.

---

**References:**
- C.A.R. Hoare, "The Emperor's Old Clothes," Turing Award Lecture, *Communications of the ACM*, Vol. 24, No. 2, February 1981.
- Alan Perlis, "Epigrams on Programming," *SIGPLAN Notices*, 1982.
- Related posts: [Simplicity does not precede complexity, but follows it](https://blog.hackspree.com/#simplicity-follows-complexity), [No solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs), [Engineering is art and philosophy](https://blog.hackspree.com/#engineering-is-economics)

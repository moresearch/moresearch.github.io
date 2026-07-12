---
title: No solutions, only trade-offs
date: 2026-07-12
slug: no-solutions-only-tradeoffs
summary: "Thomas Sowell wrote: 'There are no solutions, only trade-offs.' This is the hardest sentence in economics. It is also the hardest sentence in software engineering. Every architecture decision is a trade. The engineer who looks for a solution is looking for something that doesn't exist."
tags: sowell, trade-offs, economics, software-design, architecture
---

Thomas Sowell, the economist who has probably done more than anyone since Milton Friedman to explain scarcity to the general public, wrote a sentence so compressed it can be missed:

> "There are no solutions, only trade-offs."

This is not cynicism. It is clarity. A solution is something that makes a problem go away. A trade-off is something that makes one problem smaller at the cost of making another problem larger. Most of what we call solutions are trade-offs where we like the outcome and prefer not to think about what we gave up. The preference is human. The trade is real. The thing we gave up didn't disappear. It became someone else's problem.

Software engineering is trade-offs all the way down. Microservices trade coordination complexity for deployment independence. Monoliths trade deployment independence for coordination simplicity. Type systems trade expressiveness for safety. Dynamic languages trade safety for velocity. Relational databases trade flexibility for consistency. NoSQL trades consistency for scale. Synchronous calls trade resilience for simplicity. Asynchronous messages trade simplicity for resilience. Every architecture decision is a bet that the thing you're gaining is worth the thing you're losing. The bet is economic. The economics are usually implicit. Making them explicit is the discipline.

## The trade-off you can't see

The dangerous trade-offs are the ones where the cost is invisible. You gain something now — a feature, a shortcut, a simpler implementation. You pay later — in complexity, in reduced velocity, in a rewrite. The gain is visible. The cost is deferred. Deferred costs are easy to ignore. Ignored costs compound. Compounding costs produce bankruptcy. The bankruptcy is the rewrite. The rewrite is the admission that the original trade-off was mispriced.

> "The first lesson of economics is scarcity: There is never enough of anything to satisfy all those who want it. The first lesson of politics is to disregard the first lesson of economics." — Thomas Sowell

The first lesson of software engineering management is also to disregard the first lesson of economics. Every unrealistic deadline is a refusal to accept the trade-off between time and quality. Every under-resourced project is a refusal to accept the trade-off between scope and resources. Every "do more with less" is a refusal to accept that more of one thing means less of another. The refusal is not neutral. It pushes the cost somewhere. Usually onto the people doing the work. Usually onto the maintainability of the system. Usually onto the future, where it will be someone else's problem.

## The trade-off you choose

The mature engineer accepts that every decision is a trade. They don't look for the solution. They look for the trade-off they can live with. The question is not "what is the right answer?" The question is "what are we willing to give up?"

A team choosing between a monolith and microservices is not choosing between good and bad. They are choosing between coordination overhead distributed across teams and coordination overhead concentrated in a single codebase. The overhead doesn't disappear. It changes form. The form change may be worth it. The overhead is still there. The team that thinks microservices eliminate coordination overhead will discover that coordination overhead in a distributed system manifests as API versioning conflicts, data inconsistency, deployment ordering dependencies, and the distributed monolith. The overhead was not eliminated. It was moved. The move may have been worth it. It was not free.

A team choosing between REST and NATS is not choosing between simple and complex. They are choosing between spatial coupling (REST: callers know callees) and semantic coupling (NATS: callers and callees agree on subjects). The coupling doesn't disappear. It changes form. The form change means different failure modes, different debugging tools, different operational practices. The choice is not about which is better. It is about which form of coupling the team is equipped to manage.

## The trade-off as discipline

Naming the trade-off is the discipline. Most architecture debates are arguments about which trade-off to make, conducted by people who haven't named what they're trading. "We should use microservices." "We should stick with the monolith." The debate is unresolvable because the trade is unstated. State the trade: "Microservices will reduce coordination overhead between teams at the cost of increased operational complexity." "The monolith will reduce operational complexity at the cost of increased coordination overhead within the codebase." Now the debate is about which cost the organization is better equipped to bear. That is a debate that can be resolved. It requires knowing the organization — its team structure, its operational maturity, its tolerance for distributed complexity. The knowledge is local. The trade-off is universal.

Sowell's sentence is a tool. Apply it to every decision. When someone says "the solution is X," ask: what are we trading? When someone says "we need to do Y," ask: what are we giving up? When you find yourself certain that Z is right, ask: what cost am I not seeing? The cost is there. The certainty is hiding it. The sentence pierces the certainty. The piercing is uncomfortable. The discomfort is productive.

> "No solutions, only trade-offs."

The sentence is small. The discipline is large. The sentence is the discipline. Apply it. The decisions will improve. The systems will survive longer. The rewrites will be postponed. The postponement is a trade-off too. Everything is.

---

**References:**
- Thomas Sowell, *A Conflict of Visions*, William Morrow, 1987.
- Thomas Sowell, *Basic Economics*, Basic Books, 2000.
- Related posts: [Engineering is art and philosophy, grounded in economic law](https://blog.hackspree.com/#engineering-is-economics), [On Scarcity](https://blog.hackspree.com/#scarcity), [No Free Lunch](https://blog.hackspree.com/#scarcity-and-software-economics)

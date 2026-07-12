---
title: Engineering is economics
date: 2026-07-12
slug: engineering-is-economics
summary: "Robbins: economics is choice under scarcity (1932). Engineering: the application of knowledge to solve problems within constraints. The constraints are always economic. Time, money, attention, complexity — all finite. All have alternative uses. Engineering without economics is mathematics with a deadline it ignores."
tags: engineering, economics, scarcity, software-design
---

Engineering is the application of knowledge to solve problems within constraints. Mathematics has no constraints. Physics has natural constraints. Engineering has *chosen* constraints — budget, schedule, weight, power, materials, attention, complexity. The constraints are what make it engineering and not science. Science discovers what is true. Engineering decides what to build given that truth and a budget.

The constraints are always, ultimately, economic. Time is scarce. Money is scarce. Attention is scarce. The complexity a system can absorb before becoming unmaintainable is scarce. Every decision to use a resource for one purpose is a decision not to use it for another. That is the definition of an economic choice.

> "The first lesson of economics is scarcity: There is never enough of anything to satisfy all those who want it. The first lesson of politics is to disregard the first lesson of economics." — Thomas Sowell

The first lesson of engineering is also scarcity. The first lesson of engineering management is to disregard the first lesson of engineering. Every unrealistic deadline, every under-resourced project, every "do more with less" mandate is politics disregarding economics. The budget is real. The deadline is real. The complexity budget is real. Pretending otherwise does not make the constraints disappear. It makes the failure invisible until the constraints are violated. The violation is called a death march. The death march was predictable. The economics predicted it.

## The Robbins definition, applied

In 1932, Lionel Robbins defined economics:

> "Economics is the science which studies human behaviour as a relationship between ends and scarce means which have alternative uses."

Replace "economics" with "engineering decision-making" and the sentence remains true. Every architectural choice is a relationship between ends (what you want the system to do) and scarce means (developer time, compute, cognitive capacity) which have alternative uses (different features, different architectures, different optimizations). The engineer who does not think economically is not an engineer. They are a mathematician who happens to write code. The mathematician asks: is this correct? The engineer asks: is this correct *and* worth what it costs?

Milton Friedman compressed the same insight:

> "There's no such thing as a free lunch."

Every feature has a cost. The cost is not just the time to build it. The cost is the time plus the complexity it adds plus the features you didn't build instead. The lunch appears free because the bill arrives later. The bill is the accumulated complexity that makes every subsequent feature take longer than the last. The bill is paid in reduced velocity, increased bugs, and the eventual rewrite. The rewrite is bankruptcy. The lunch was never free. The accounting was deferred.

## What economics gives engineering

Economics gives engineering three things: a vocabulary, a decision framework, and a theory of systems.

**The vocabulary.** Opportunity cost. Sunk cost. Marginal cost. Comparative advantage. Economies of scale. Net present value. Option value. Each term names a specific structure. Each structure appears in engineering decisions whether you name it or not. Naming it lets you reason about it. Not naming it lets it operate invisibly. Invisible economic forces produce worse engineering decisions than visible ones.

Opportunity cost is the largest cost in software engineering and the one that appears on no invoice. The features you didn't build are the cost of the features you did. Every hour spent on Feature A is an hour not spent on Feature B. The cost of A is not just the time to build it. It is the time plus the value of B. B was never built, so its value was never measured. The absence of measurement doesn't mean the cost wasn't real. It means the accounting was incomplete.

Sunk cost is the cost already incurred and unrecoverable. The three years spent on the monolith are sunk. They should not influence the decision to migrate. They do. The influence is irrational. The irrationality is human. Friedrich Hayek explained why no single mind can overcome it:

> "The knowledge of the circumstances of which we must make use never exists in concentrated or integrated form, but solely as the dispersed bits of incomplete and frequently contradictory knowledge which all the separate individuals possess."

The knowledge of what the system costs, what it should become, and what it would take to get there — it is dispersed across the team, the codebase, the incident history, the user feedback. No single person holds it all. The architect who pretends to is making decisions with incomplete information. The admission of incomplete information is the beginning of economic thinking. The pretense of complete information is the beginning of architectural hubris.

**The decision framework.** Robbins's four questions apply to any engineering choice. What is the end? What are the means? Are they scarce? Do they have alternative uses? If yes — and it always is — the decision is economic. Name the scarce resource. Name the alternative foregone. Calculate the trade. Make the choice.

Most engineering debates skip this. Two engineers debate two architectures. Neither names the constraint. The debate is unresolvable because the constraint is unstated. "We are optimizing for developer time, not compute." "We are optimizing for change velocity, not operational simplicity." State the constraint. The debate resolves. The constraint determines the answer. Name the constraint.

**The theory of systems.** Hayek's price system is a mechanism for communicating scarcity without requiring anyone to understand the whole. When the price of compute rises, users use less compute. They don't need to know why the price rose. The price communicates the scarcity. The behavior adapts. The system self-regulates. This is mechanism design before the term existed: design the signal, not the response. The signal communicates the scarcity. The agents adapt. The system finds equilibrium.

A well-designed software system does the same. API rate limits are prices. Queue depths are prices. Circuit breaker states are prices. Each communicates a scarcity — of compute, of downstream capacity, of healthy instances. Services adapt by backing off, retrying, routing elsewhere. No service knows why the scarcity exists. No service needs to. The signal communicates. The behavior adapts. The system self-regulates. This is economics implemented as infrastructure. The infrastructure is the mechanism. The mechanism communicates scarcity. Scarcity drives behavior.

## What happens without economics

Engineering without economics produces systems that are correct and useless. The architecture is elegant. The abstractions are clean. The system solves a problem nobody has, at a cost nobody calculated, with a complexity budget nobody tracked. The system is correct. It is also uneconomic. Uneconomic systems are not deployed, or are deployed and abandoned, or are deployed and maintained at a loss until the loss becomes undeniable. The loss was always deniable because the economics were never made explicit. Explicit economics can be debated. Implicit economics produce implicit failure.

Engineering without economics also produces systems that are popular and catastrophic. The features shipped on time. The architecture was compromised at every deadline. The complexity budget was overdrawn at every sprint. The system worked until it didn't. When it didn't, the rewrite cost more than the original build. The rewrite was necessary because the economics were ignored. The economics were ignored because the deadlines didn't include them. The deadlines were economic decisions stated as technical requirements. "Ship by Q3" is an economic decision. It means "we value time-to-market over complexity budget." Stating it that way forces the question: is the time-to-market premium worth the future complexity cost? Sometimes yes. Usually nobody asks.

## The engineers who knew

The great software engineers were economists before the field existed. Brooks: conceptual integrity is an argument about the diseconomy of scale in design teams. Parnas: information hiding is an argument about the option value of clean interfaces. Lehman: E-type systems must evolve, and evolution increases complexity unless work is done to reduce it — an argument about the depreciation of unmaintained structure. Christensen: the innovator's dilemma is an argument about capital allocation under asymmetric constraints. None of them used the economic vocabulary. All of them made economic arguments. The vocabulary existed. They didn't know it. The arguments were correct anyway.

Tony Hoare stated the economic tradeoff in design:

> "There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult."

Far more difficult means far more expensive upfront. Far less expensive over time. The economic choice is between paying now or paying later. Paying now is visible. Paying later is invisible. The invisible cost feels cheaper. It isn't. The accounting that makes it feel cheaper is wrong.

## The unifying claim

There is no engineering without economics because engineering is decision-making under constraints, and constraints are the subject matter of economics. Every engineering decision is an economic decision stated in technical language. The language obscures the economics. The economics determine the outcome. The engineer who learns the economic vocabulary makes better decisions. The engineer who doesn't makes economic decisions without knowing they are economic decisions. The decisions are still economic. They are just worse.

> "No solutions, only trade-offs." — Thomas Sowell

Engineering is the discipline of trade-offs. Economics is the science of trade-offs. The two are the same activity described in different vocabularies. Learn both vocabularies. Apply both to every decision. The decisions will improve. The systems will survive longer. The rewrite will be postponed. The rewrite is always coming. Economics postpones it. Ignorance accelerates it.

---

**References:**
- Lionel Robbins, *An Essay on the Nature and Significance of Economic Science*, Macmillan, 1932.
- Thomas Sowell, *Basic Economics*, Basic Books, 2000.
- Friedrich Hayek, "The Use of Knowledge in Society," *American Economic Review*, 1945.
- Frederick P. Brooks, Jr., *The Mythical Man-Month*, Addison-Wesley, 1975.
- David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, 1972.
- M.M. Lehman, "Programs, Life Cycles, and Laws of Software Evolution," *Proceedings of the IEEE*, 1980.
- Related posts: [On Scarcity](https://blog.hackspree.com/#scarcity), [Scarcity and Software Economics](https://blog.hackspree.com/#scarcity-and-software-economics), [Scarcity in Practice](https://blog.hackspree.com/#scarcity-in-practice)

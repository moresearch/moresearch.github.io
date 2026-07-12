---
title: On Scarcity
date: 2026-07-12
slug: scarcity
summary: "Robbins defined economics as choice under scarcity in 1932: ends, means, scarcity, alternative uses. Software engineering is economics by other means. Time, attention, complexity, and compute are the four scarcities that shape every architecture decision."
tags: scarcity, economics, software-engineering, robbins
series: scarcity
part: 1
---

In 1932, Lionel Robbins published *An Essay on the Nature and Significance of Economic Science*. On page 15, he wrote:

> "Economics is the science which studies human behaviour as a relationship between ends and scarce means which have alternative uses."

Four concepts. **Ends** — what you want. **Means** — what you have. **Scarcity** — means are finite. **Alternative uses** — means can be deployed in different ways, forcing choice. Every human activity involving these four conditions is economic. Robbins made this explicit:

> "Insofar as it deals with the influence of scarcity, any kind of human behaviour falls within the scope of Economic Generalisations. There are no limitations on the subject-matter of Economic Science save this."

There are no limitations on the subject-matter. Software engineering is human behaviour under conditions of scarcity. Time is scarce. Attention is scarce. Compute is scarce. Complexity budget is scarce. You have finite means — developer hours, cognitive capacity, hardware, money — and infinite ends — features to build, bugs to fix, systems to improve, customers to satisfy. The means have alternative uses. Every hour spent on one feature is an hour not spent on another. Every dollar spent on infrastructure is a dollar not spent on hiring. Every ounce of cognitive capacity spent on one problem is an ounce not available for another. This is economics. The subject-matter is software. The structure is scarcity.

## The four scarcities of software

**Time is scarce.** You have finite developer hours. Every feature you build is a feature you didn't build. Every refactor you do is a refactor you didn't do. Every meeting you attend is code you didn't write. The scarcity of time forces prioritization. Prioritization is an economic act. The framework for prioritization — what to build now, what to build later, what to never build — is capital budgeting applied to code. The budget is time. The investments are features. The returns are user value, revenue, reduced maintenance cost. The same mathematics that determines whether to build a factory determines whether to build a microservice. The logic is identical.

**Attention is scarce.** A developer can hold one complex problem in their head at a time. Two if exceptional. Three is impossible. The scarcity of attention is the binding constraint on software complexity. Brooks's Law — adding people to a late project makes it later — is a statement about attention scarcity. Each new person must learn the system. The teaching consumes the attention of those who already know it. Communication overhead grows quadratically. Attention per person shrinks. The project gets later. Brooks's argument for conceptual integrity — one mind controlling the design — is also about attention scarcity. The design must fit in one mind because only one mind can hold it. When the design exceeds one mind's capacity, it must be split. The split requires coordination. Coordination consumes attention. Attention consumed by coordination is not available for design. Design quality degrades.

**Complexity budget is scarce.** Lehman's Second Law: complexity increases unless work is done to reduce it. The work requires time and attention, both scarce. The complexity budget is the total complexity the system can absorb before becoming unmaintainable. Every feature adds to the budget. Every quick fix adds. Every workaround adds. The budget is finite. When exhausted, the system must be rewritten. The rewrite is expensive. The expense is the cost of exceeding the budget. The budget was always finite. The accounting ignored it.

Parnas's information hiding is an economic strategy. Hide volatile decisions behind stable interfaces. The interface is the investment. The hiding is the return. When the volatile decision changes, the change is contained within the module. The containment saves time, attention, and complexity budget. The investment costs more upfront. The return accrues over time — every change that doesn't propagate is a cost avoided. The net present value is positive if the decision is sufficiently volatile. Parnas didn't state it in economic terms. The economics are implicit. The economics are correct.

**Compute is scarce.** Moore's Law made compute cheap but not free. Cloud computing made it elastic but not infinite. Every algorithm choice trades compute for something else — development time, code simplicity, latency. The trade is economic. The price of compute determines which side is optimal. When compute was expensive, developers optimized for cycles. When compute became cheap, they optimized for developer time. The price changed. The optimal trade changed. The change was economic. The economics were invisible to the developers. They thought they were making technical decisions. They were making economic decisions with technical parameters.

## The Robbins test

Robbins gives a test for any decision. Ask four questions:

1. What is the end? (What are you trying to achieve?)
2. What are the means? (What resources do you have?)
3. Are the means scarce? (Could you use them for something else?)
4. Do they have alternative uses? (What are you giving up?)

If the answer to 3 and 4 is yes — and it always is — the decision is economic. Treat it as such. Name the scarce resource. Name the alternative foregone. Calculate the trade. Make the choice. The choice will be better for having been made explicitly.

Most software decisions are made without naming the scarcity. Two engineers debate two architectures. Neither names the constraint. The debate is unresolvable because the constraint is unstated. State the constraint. The debate resolves. This feature or that feature? State the scarce resource. This architecture or that architecture? State the trade. The Robbins test makes implicit economics explicit. Explicit economics are debatable. Implicit economics are invisible. Invisible economics produce worse decisions.

## The unifying principle

Resources are finite. Ends are infinite. Means have alternative uses. Every decision is a choice under scarcity. The choice has consequences. The consequences propagate. The propagation is the system's behavior. The behavior is emergent from the choices. The choices are economic. Software engineering is economics.

---

**This is part 1 of a 7-part series on scarcity and software.**
- [Part 2: On Games](https://blog.hackspree.com/#on-games)
- [Part 3: On Software Engineering Economics](https://blog.hackspree.com/#on-software-engineering-economics)
- [Part 4: On Games in Software](https://blog.hackspree.com/#on-games-in-software)
- [Part 5: On AI and Mechanism Design](https://blog.hackspree.com/#on-ai-and-mechanism-design)
- [Part 6: On Practice](https://blog.hackspree.com/#on-practice)
- [Part 7: The Catalog of Games](https://blog.hackspree.com/#the-catalog-of-games)

**References:**
- Lionel Robbins, *An Essay on the Nature and Significance of Economic Science*, Macmillan, 1932.
- Frederick P. Brooks, Jr., *The Mythical Man-Month*, Addison-Wesley, 1975.
- David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, 1972.
- M.M. Lehman, "Programs, Life Cycles, and Laws of Software Evolution," *Proceedings of the IEEE*, 1980.

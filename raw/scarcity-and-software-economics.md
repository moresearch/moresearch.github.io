---
title: No Free Lunch
date: 2026-07-12
slug: scarcity-and-software-economics
summary: "Barry Boehm named the field in 1981. This post defines every economic concept that applies to software — opportunity cost, sunk cost, NPV, option value, technical debt, build vs. buy — and shows how Brooks, Parnas, and Lehman were economists before the field existed."
tags: software-economics, boehm, opportunity-cost, technical-debt, npv
series: scarcity
part: 3
---

Barry Boehm published *Software Engineering Economics* in 1981. The book applied cost estimation, net present value, and decision analysis to software projects. It named a field that had been practiced without a name since the first programmer decided which feature to build first. The vocabulary is precise. Most engineers use it without knowing the definitions. The definitions matter.

Milton Friedman distilled economics to one sentence: "There's no such thing as a free lunch." Every feature has a cost. The cost is not just the time to build it. It is the time plus the complexity it adds plus the features you didn't build instead. The lunch appears free. The bill arrives in year three.

Every organization has a discount rate — the rate at which it discounts the future relative to the present. The discount rate is the most important number in your technical culture and the one nobody knows. A high discount rate means next quarter matters more than next year. Features ship. Refactors don't. Complexity accumulates. A low discount rate means sustainability matters. The refactor gets done. The feature waits. Neither rate is correct in the abstract. The correct rate depends on whether the company will exist in five years. The discount rate should be a conscious choice. It never is. It is set by the urgency of the nearest deadline. The deadline is a discount rate of nearly infinity. Infinity is too high.

## The terminology

**Opportunity cost.** The value of the best alternative foregone. Every hour on Feature A is an hour not on Feature B. The cost of Feature A is the time to build it plus the value of whatever you would have built instead. Opportunity cost is invisible. It appears on no invoice. It is the largest cost in software engineering. The features you didn't build are the cost of the features you did.

**Sunk cost.** A cost already incurred and unrecoverable. The three years on the monolith are sunk. They should not influence the migration decision. They do. The influence is irrational. The irrationality is human. The only defense is to externalize the decision to a process that doesn't know the sunk cost.

**Marginal cost.** The cost of one additional unit. The marginal cost of a user is near zero for software. The marginal cost of complexity is not zero — adding a feature to an already-complex system costs more than adding it to a simple one. Software has decreasing marginal cost of serving users and increasing marginal cost of adding features. The first makes software businesses attractive. The second makes them eventually unmaintainable.

**Comparative advantage.** Produce what you're relatively better at. Trade for the rest. Ricardo's logic, applied to microservices. A team better at both frontend and backend should specialize in whichever it has the greater relative advantage in. Absolute advantage doesn't matter. Comparative advantage does. Most teams don't know theirs. They assume they should build everything. They are wrong.

**Economies of scale.** Cost per unit decreases as volume increases. The monolith has economies of scale: one build pipeline, one deployment. Microservices lose these in exchange for independent deployability. The optimal number of services is where the marginal benefit of independence equals the marginal cost of lost scale. Few calculate this. Most guess. The guess is usually wrong.

**Diseconomies of scale.** Cost per unit increases as volume increases. Brooks's Law is a diseconomy: adding people increases output by less than the increase in coordination cost. Net effect negative. Diseconomies of scale are why organizations don't grow infinitely. At some size, internal coordination cost exceeds external transaction cost. Coase (1937): firms exist because internal transactions are cheaper than market transactions. The firm's boundary is where internal cost equals market cost. The service's boundary is where building equals buying.

**Net present value (NPV).** The current value of future cash flows, discounted by the time value of money. A refactoring costing $100K now and saving $20K/year for ten years has NPV dependent on the discount rate. At 5%, positive. At 15%, negative. The discount rate is the organization's preference for present over future. Organizations with high discount rates don't refactor. The rate is set by quarterly earnings pressure. The pressure is economic. The decision not to refactor is economic, stated in technical language.

**Option value.** The value of keeping a choice available. A clean interface has option value: change the implementation later without changing callers. The option costs more upfront. The option is valuable — being able to change without coordination. Financial options have a market price (Black-Scholes). Software options don't. They have an implicit value estimated by the architect. Good architects price options correctly. Bad architects don't know they're pricing options.

**Technical debt as economic debt.** Borrowing future productivity to increase present velocity. Principal: the cleanup work required. Interest: reduced velocity from uncleaned code. Interest compounds. Compounding interest eventually consumes all available velocity. The system becomes unchangeable. The debt must be repaid or defaulted on. Default is a rewrite. The rewrite is bankruptcy. The accounting that didn't track the debt was wrong.

**Cost of delay.** Revenue or value lost per unit time by not shipping. If Feature A generates $10K/week and takes 10 weeks, cost of delay is $100K. If Feature B generates $2K/week, cost of delay is $20K. Build A first. Most prioritization ignores cost of delay. They prioritize by effort, by intuition, by the loudest stakeholder. That framework is not a framework. It is ritualized negotiation.

**Build vs. buy.** Make-or-buy. Build if internal cost < market cost, adjusted for risk, control, and strategic value. Coase: the boundary of the firm is where internal cost equals market cost. Teams that build everything have not calculated the boundary. They have assumed it.

## The economists who didn't know they were economists

Tony Hoare observed: "There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult." The difficult method costs more upfront. The easy method costs more over time. The economics of software is the economics of choosing which cost to pay. Most choose the easy method. The easy method is why refactoring exists.


**Brooks.** Conceptual integrity is an economic argument. One mind controls the design because coordination cost exceeds the benefit of additional designers. The n(n−1)/2 communication paths make adding designers counterproductive. Brooks's Law is a statement about the diseconomy of scale in software teams. Brooks was doing economics without the vocabulary.

**Parnas.** Information hiding is an economic strategy. Invest in a stable interface. The return is reduced propagation of change. The investment costs more upfront. The return accrues over time. The net present value is positive if the decision is sufficiently volatile. Parnas was pricing options without Black-Scholes.

**Lehman.** E-type systems must evolve. Evolution increases complexity unless work is done to reduce it. The work costs time and attention — both scarce. The complexity budget is finite. The accounting that ignores it produces systems cheap to build and expensive to maintain. Lehman was doing cost accounting without the ledger.

**Christensen.** Incumbents are held captive by their customers. The captivity is rational under conditions of scarcity. The resources that could fund the disruption are allocated to sustaining innovations for existing customers. The allocation is optimal in the short run. Fatal in the long run. Christensen was describing capital allocation under asymmetric constraints without calling it that.

---

**This is part 3 of a 7-part series on scarcity and software.**
- [Part 1: On Scarcity](https://blog.hackspree.com/#scarcity)
- [Part 2: On Games](https://blog.hackspree.com/#scarcity-and-games)
- [Part 4: On Games in Software](https://blog.hackspree.com/#scarcity-and-software-games)
- [Part 5: On AI and Mechanism Design](https://blog.hackspree.com/#scarcity-and-mechanism-design)
- [Part 6: On Practice](https://blog.hackspree.com/#scarcity-in-practice)
- [Part 7: The Catalog of Games](https://blog.hackspree.com/#catalog-of-scarcity-games)

**References:**
- Barry W. Boehm, *Software Engineering Economics*, Prentice-Hall, 1981.
- R.H. Coase, "The Nature of the Firm," *Economica*, 1937.
- Frederick P. Brooks, Jr., *The Mythical Man-Month*, Addison-Wesley, 1975.
- David L. Parnas, "On the Criteria to Be Used in Decomposing Systems into Modules," *Communications of the ACM*, 1972.
- M.M. Lehman, "Programs, Life Cycles, and Laws of Software Evolution," *Proceedings of the IEEE*, 1980.
- Clayton M. Christensen, *The Innovator's Dilemma*, Harvard Business School Press, 1997.


Scarcity is the universal engineering constraint. Time, attention, compute, complexity — every engineering decision is made within a budget. The budget is economic. The engineer who doesn't track the budget makes decisions blind. The engineer who tracks it makes decisions with full knowledge of the trade-off. The trade-off is the decision. The budget is the constraint. Scarcity is the unifying principle.

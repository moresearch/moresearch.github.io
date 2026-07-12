---
title: "Game Theory Model: The Principal-Agent Problem"
date: 2026-07-12
slug: game-principal-agent
summary: "A principal hires an agent to do work. The agent has private information and their own objectives. The principal cannot perfectly monitor the agent's effort. The problem: how to design a contract that aligns the agent's incentives with the principal's goals. This is the model for every employment relationship, every outsourcing decision, and every API contract between teams."
tags: game-theory, principal-agent, moral-hazard, contract-theory, incentives
series: game-theory-models
part: 10
---

The principal-agent problem is the foundational model of contract theory. A principal wants a task performed. An agent can perform it. The agent's effort is unobservable — the principal sees the outcome but not the effort. The outcome depends on effort and on random factors beyond the agent's control. The principal must design a contract — a payment scheme contingent on observable outcomes — that incentivizes the agent to exert the desired level of effort.

The tension: the principal wants high effort at low cost. The agent wants high payment for low effort. If the principal pays a fixed wage, the agent exerts minimum effort — there's no incentive to do more. If the principal pays entirely based on outcome, the agent bears all the risk from random factors. The optimal contract balances insurance (protecting the agent from randomness) with incentives (rewarding effort). The balance is the contract. The contract is the mechanism.

The problem has two variants. Moral hazard: the agent takes hidden actions after the contract is signed. Adverse selection: the agent has hidden information before the contract is signed — the agent knows their own type, and the principal doesn't. Both are information asymmetries. Both require mechanism design to resolve.

## Interpretations from different branches

**Contract theory (Hart, Holmström, 2016 Nobel).** Oliver Hart and Bengt Holmström shared the 2016 Nobel for contract theory. Holmström's informativeness principle: any performance measure that provides information about effort should be included in the contract. If the measure is informative, including it reduces the agent's risk for a given level of incentives. Hart's incomplete contracts: real contracts cannot specify every contingency. When contracts are incomplete, the allocation of residual control rights — who decides what happens in unforeseen circumstances — determines outcomes. Ownership matters because ownership confers residual control.

**Corporate governance (Jensen and Meckling, 1976).** The separation of ownership and control in corporations is a principal-agent problem. Shareholders are principals. Managers are agents. Managers may pursue their own interests — empire-building, risk-aversion, short-term stock price — rather than shareholder value. The mechanisms: stock options (aligning incentives), boards of directors (monitoring), hostile takeovers (discipline). The mechanisms are imperfect. The imperfection is the cost of the agency relationship.

**Regulation (Laffont and Tirole, 1993 Nobel 2014).** Regulating a monopoly is a principal-agent problem. The regulator (principal) wants the monopoly (agent) to operate efficiently. The monopoly has private information about its costs. The regulator designs a pricing scheme that incentivizes cost reduction while preventing excessive pricing. The scheme is a contract. The contract is mechanism design applied to public utility regulation.

**Political science.** Voters are principals. Politicians are agents. The agency problem is accountability. Elections are the incentive mechanism — politicians who perform poorly are voted out. The mechanism is imperfect because voters have incomplete information, politicians control the flow of information, and election cycles are coarse. The imperfections are the subject of political economy.

## Software engineering interpretations

**Team and manager.** The manager (principal) wants the team (agent) to produce high-quality work. The manager cannot perfectly observe effort — code quality is partly effort, partly skill, partly the difficulty of the task. The contract: salary, performance review, promotion. The mechanisms: code review (monitoring), OKRs (outcome-based incentives), peer feedback (multi-source monitoring). Each mechanism reduces the information asymmetry.

**Outsourcing vendor management.** The company (principal) hires a vendor (agent) to build a system. The vendor has private information about their true costs, their true timeline, and the quality of their engineers. The contract: fixed-price or time-and-materials. Fixed-price transfers risk to the vendor but creates incentive to cut corners. Time-and-materials transfers risk to the company but creates incentive to inflate hours. The optimal contract balances risk-sharing with incentive alignment. The balance is the principal-agent problem in procurement form.

**Platform team and service teams.** The platform team (agent) provides infrastructure to service teams (principals). The service teams cannot observe the platform team's effort. The platform team may optimize for its own interests — interesting technical work, clean architecture — rather than service team needs. The contract: SLAs with penalties, internal billing (chargebacks), user satisfaction surveys. The mechanisms align the platform team's incentives with service team outcomes.

**Open-source maintainer and corporate user.** The maintainer (agent) produces a library. The corporation (principal) depends on it. The maintainer's effort is unobservable. The corporation cannot compel the maintainer to fix bugs or accept patches. The contract is social, not legal — reputation, sponsorship, contribution guidelines. The principal-agent problem in open source is acute because the mechanisms are weak. The weakness is why maintainers burn out.

## The alignment problem

The principal-agent problem is the alignment problem in economic form. Align the agent's incentives with the principal's goals. The alignment is never perfect because information is never perfect. The residual misalignment is the agency cost. The cost is irreducible. The mechanism designer's job is to minimize it. Minimize it by making outcomes observable, linking rewards to observables, and accepting that some misalignment will remain. The acceptance is realism. The minimization is engineering.

---

**References:**
- Bengt Holmström, "Moral Hazard and Observability," *Bell Journal of Economics*, 1979.
- Oliver Hart and Bengt Holmström, "The Theory of Contracts," 1987.
- Michael Jensen and William Meckling, "Theory of the Firm: Managerial Behavior, Agency Costs and Ownership Structure," *Journal of Financial Economics*, 1976.
- Related posts: [Design the Game](https://blog.hackspree.com/#scarcity-and-mechanism-design), [Mechanism Design](https://blog.hackspree.com/#game-mechanism-design)

---
title: Field Guide to Scarcity Games
date: 2026-07-12
slug: catalog-of-scarcity-games
summary: "An exhaustive taxonomy of game types, organized by cooperation, conflict, timing, information, symmetry, player count, determinism, strategy, state space, time horizon, and special classes. Every category has a concrete software example. A reference for recognizing game structures in architecture."
tags: game-theory, taxonomy, reference, software-engineering
series: scarcity
part: 7
---

Every software situation is a game. The game has a type. The type determines the appropriate analysis. This catalog classifies games along eleven independent dimensions. Each category includes a concrete software example. Use it to locate any situation in game-theoretic space. The location tells you what you're dealing with. What you're dealing with determines what you should do.

This catalog is a diagnostic tool. When a situation feels stuck — the migration isn't happening, the standard isn't being adopted, the teams keep breaking each other's APIs — locate it here. The name of the game tells you why it's stuck. The why tells you what to change. Most organizational interventions fail because they treat every stuck situation as a communication problem. Some are stag hunts. Some are chicken. Some are commons tragedies. The solution to a stag hunt is visible early adopters. The solution to chicken is a pre-committed rule. The solution to a commons tragedy is a governance mechanism. The solutions are different because the games are different. The catalog tells you which one you're in.

> "Simplicity does not precede complexity, but follows it." — Alan Perlis

The catalog that follows is a map of the complexity. Use it to find the simplicity on the other side.

## By cooperation

| Category | Description |
|---|---|
| **Cooperative** | Players can form binding agreements. — SLA with penalty clauses. The SLA is the binding agreement. |
| **Non-cooperative** | No binding agreements. — Teams without API contracts. Communication is cheap talk. |
| **Cheap talk** | Communication allowed but unenforceable. — "We promise not to break the API." Without testing, this is cheap talk. |
| **Bargaining** | Two-player cooperative game over surplus division. — Two teams negotiating shared infrastructure costs. |
| **Team problems** | Non-cooperative structure, identical payoffs. — Team sharing the same OKRs. Incentives aligned by design. |

## By conflict of interest

| Category | Description |
|---|---|
| **Zero-sum** | One's gain = another's loss. — Fixed headcount allocation. Every hire for Team A is one Team B doesn't get. |
| **Non-zero-sum** | Players can both gain or both lose. — API design. Both gain from clean contract. Both lose from broken one. |
| **Coordination** | No conflict. Players need to align. — Choosing a shared logging format. Everyone wants the same thing. |
| **Mixed-motive** | Both conflict and cooperation incentives. — Most real situations. Shared goals but private priorities. |

## By timing

| Category | Description |
|---|---|
| **Simultaneous** | Players move without observing others. — Teams making independent tech choices in the same quarter. |
| **Sequential** | Players move in turn. — Deployment ordering. Stackelberg: first-mover advantage. |
| **Repeated** | Same game multiple times. — Sprint planning. Repetition enables reputation and reciprocity. |
| **One-shot** | Played exactly once. — A rewrite decision. No repetition. |
| **Stochastic** | State evolves probabilistically. — Incident response. Alerts arrive randomly. |

## By information

| Category | Description |
|---|---|
| **Complete information** | All players know all payoffs. — Open-source. Everyone sees the code, issues, priorities. |
| **Incomplete information** | Private information. Bayesian games. — Team proposing rewrite has private info about true motivation. |
| **Perfect information** | All previous moves known. — Monolith. Every module observes every other module's state. |
| **Imperfect information** | Some moves unknown. — Microservices. Service A can't observe Service B's internal state. |
| **Symmetric information** | Same uncertainty. — Both teams uncertain about new CTO's priorities. |
| **Asymmetric information** | Different knowledge. — Senior engineer knows legacy weaknesses. New hire doesn't. |

## By symmetry

| Category | Description |
|---|---|
| **Symmetric** | Payoffs depend only on strategies, not identity. — Two identical microservices with the same SLA. |
| **Asymmetric** | Changing identities changes payoffs. — Frontend vs. backend team. Different constraints, different strategies. |

## By player count

| Category | Description |
|---|---|
| **1-player** | Decision theory. No strategic interaction. — Choosing an algorithm. Nature is the constraint. |
| **2-player** | Classic case. — Two teams negotiating an API contract. |
| **N-player** | Three or more. Coalitions possible. — Organization with many teams. Alliances form. Politics emerges. |
| **Large/Many-player** | Continuum. Atomic vs. non-atomic. — Open-source ecosystem. No single contributor changes equilibrium. |
| **Mean field** | Players interact through state distribution. — Microservices at scale. Each service interacts with aggregate behavior. |

## By determinism

| Category | Description |
|---|---|
| **Deterministic** | No chance elements. — Deterministic deployment pipeline. |
| **Stochastic** | Some moves by nature/chance. — System with probabilistic failures. |
| **Games of chance** | All moves by one player and chance. — A/B testing. Nature randomizes users. |

## By strategy type

| Category | Description |
|---|---|
| **Pure strategy** | Single deterministic action. — Always use the same database for every service. |
| **Mixed strategy** | Randomize over actions. — Random on-call assignment. Prevents predictable exploitation. |

## By state/action space

| Category | Description |
|---|---|
| **Finite** | Finite actions and states. — Choosing between exactly three database options. |
| **Infinite/Continuous** | Continuous spaces. — Allocating compute resources. Continuous variable. |
| **Discrete-time** | Decisions at discrete intervals. — Sprint planning. Every two weeks. |
| **Continuous-time** | Decisions at every instant. — Real-time auto-scaling. Differential equations. |

## By time horizon

| Category | Description |
|---|---|
| **Finite horizon** | Known number of periods. — Project with fixed deadline and milestones. |
| **Infinite horizon** | No predetermined end. — Ongoing system maintenance. Cooperation sustainable. |
| **Discounted** | Future payoffs weighted less. — Quarterly planning. This quarter > next year. |

## Special game classes

| Category | Description |
|---|---|
| **Signaling** | Informed player acts to reveal info. — Detailed RFC signals competence and seriousness. |
| **Screening** | Uninformed player moves first. — Requiring prototype before architecture review. |
| **Stackelberg** | Leader-follower. Leader first. — Platform team sets API standard. Services build against it. |
| **Pursuit-evasion** | Zero-sum differential. — Intrusion detection. Attacker evades. Defender pursues. |
| **Mechanism design** | Reverse game theory. — SLAs, contract testing, deployment gates, code review requirements. |
| **Global games** | Noisy private signals of underlying state. — Teams deciding on new tech based on private maturity signals. |
| **Combinatorial** | Finite, deterministic, perfect-info, 2P, zero-sum. — Automated theorem proving. Compiler optimization. |
| **Evolutionary** | Fit strategies survive. Replicator dynamics. — Architecture patterns that persist. Microservices resist re-monolithing. |
| **Partizan** | Moves differ per player. — Frontend and backend teams have different available moves. |
| **Impartial** | Moves depend only on position, not player. — Identical services with identical deployment options. |

---

**This is part 7 of a 7-part series on scarcity and software.**
- [Part 1: On Scarcity](https://blog.hackspree.com/#scarcity)
- [Part 2: On Games](https://blog.hackspree.com/#scarcity-and-games)
- [Part 3: On Software Engineering Economics](https://blog.hackspree.com/#scarcity-and-software-economics)
- [Part 4: On Games in Software](https://blog.hackspree.com/#scarcity-and-software-games)
- [Part 5: On AI and Mechanism Design](https://blog.hackspree.com/#scarcity-and-mechanism-design)
- [Part 6: On Practice](https://blog.hackspree.com/#scarcity-in-practice)


Scarcity is the universal engineering constraint. Time, attention, compute, complexity — every engineering decision is made within a budget. The budget is economic. The engineer who doesn't track the budget makes decisions blind. The engineer who tracks it makes decisions with full knowledge of the trade-off. The trade-off is the decision. The budget is the constraint. Scarcity is the unifying principle.

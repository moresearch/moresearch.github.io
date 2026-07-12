---
title: "Game Theory Model: The Tragedy of the Commons"
date: 2026-07-12
slug: game-tragedy-of-commons
summary: "Garrett Hardin described it in 1968: a shared resource, individually rational use, collective ruin. The Tragedy of the Commons is a multi-player Prisoner's Dilemma with a shared resource. It explains overloaded CI pipelines, degraded staging environments, and every shared service that nobody maintains. Elinor Ostrom won a Nobel proving it's not inevitable."
tags: game-theory, tragedy-of-commons, hardin, ostrom, common-pool-resources
series: game-theory-models
part: 7
---

Garrett Hardin published "The Tragedy of the Commons" in *Science* in 1968. The argument: a pasture open to all herders. Each herder gains the full benefit of adding an animal — more milk, more meat. Each herder bears only a fraction of the cost — the pasture is shared, the overgrazing is distributed. The individually rational choice: add another animal. The collective outcome: the pasture is destroyed. Everyone loses.

> "Therein is the tragedy. Each man is locked into a system that compels him to increase his herd without limit — in a world that is limited. Ruin is the destination toward which all men rush, each pursuing his own best interest in a society that believes in the freedom of the commons."

The tragedy is a multi-player Prisoner's Dilemma with a renewable resource. The resource has a carrying capacity. Below capacity, grazing is sustainable. Above capacity, the resource degrades. Each player's marginal benefit of adding an animal is private. The marginal cost is shared. The private benefit exceeds the private cost until the resource collapses. The collapse is predictable. The predictability doesn't prevent it.

## Interpretations from different branches

**Classical game theory.** The Tragedy is a social dilemma. The Nash equilibrium is overuse. The equilibrium is Pareto-suboptimal — everyone would be better off with restraint. The dilemma is structurally identical to the Prisoner's Dilemma scaled to N players. The larger the N, the smaller each player's share of the cost of their own overuse. The tragedy intensifies with group size.

**Institutional economics (Ostrom, 1990 Nobel 2009).** Elinor Ostrom challenged Hardin's conclusion that commons are inevitably tragic. She studied real communities that managed common-pool resources successfully for centuries — Swiss grazing pastures, Japanese forests, Spanish irrigation systems. She identified eight design principles for sustainable commons: clearly defined boundaries, proportional equivalence between benefits and costs, collective choice arrangements, monitoring, graduated sanctions, conflict resolution mechanisms, minimal recognition of rights, and nested enterprises. The principles are mechanism design for commons. The mechanisms work. The tragedy is not inevitable.

**Environmental economics.** Climate change is the ultimate tragedy of the commons. The atmosphere is a shared resource. Carbon emissions benefit the emitter. The cost is shared globally. The individually rational choice: emit. The collectively optimal choice: reduce. The coordination problem is planet-scale. The mechanisms — carbon taxes, cap-and-trade, international agreements — are Ostrom principles scaled to nations. The scaling is hard because enforcement across sovereign states is weak. The weakness is the tragedy.

**Digital commons.** Open-source software is a commons. The code is a shared resource. Contributors maintain it. Users consume it. The individually rational choice: use without contributing. The collectively optimal choice: everyone contributes. The tragedy: maintainer burnout, abandoned projects, the internet running on a single developer's unpaid labor. The Ostrom principles apply: clear governance, graduated sanctions (from bug reports to commit access), collective choice (RFC processes). The mechanisms exist. They are fragile.

## Software engineering interpretations

**The shared CI pipeline.** Every team adds tests to the shared CI pipeline. Each test benefits the team that added it. The cost — longer build times — is shared by all teams. The individually rational choice: add tests. The collectively optimal choice: add only high-value tests. The tragedy: the pipeline takes 45 minutes. Everyone suffers. The mechanism: a test budget per team, periodic culling, a requirement that new tests justify their existence with historical failure catch rate. Ostrom's proportional equivalence: the cost a team imposes on the commons must be proportional to the benefit they derive.

**The staging environment.** Staging is a shared resource. Every team wants to use it for integration testing, demos, and load testing. Overuse degrades it. The individually rational choice: use staging whenever you need it. The tragedy: staging is unreliable, nobody trusts it, everyone builds their own staging-like environment. The mechanism: a booking calendar, dedicated demo environments, tiered access with SLAs. Ostrom's clearly defined boundaries: who can use staging, for what, when.

**The monolith as commons.** The monolith's codebase is a shared resource. Every team adds code. The cost of complexity is shared. The individually rational choice: add the feature in the simplest way for your team. The tragedy: the monolith becomes unmaintainable. The mechanism: module ownership, interface contracts, automated complexity budgets per module. Ostrom's monitoring: visibility into who added what complexity, and what it cost.

**The shared database.** Multiple services read and write the same database. Each service optimizes its queries for its own use case. The cost — contention, locking, schema complexity — is shared. The tragedy: the database is the bottleneck, nobody can change their schema without breaking others. The mechanism: one service owns the database, all access goes through its API. Ostrom's clearly defined boundaries: the database is not a commons. It is property. The property has an owner.

## Resolving the commons

Hardin was wrong that tragedy is inevitable. Ostrom was right that it can be governed. The governance requires: clear boundaries, proportional costs, collective decision-making, monitoring, graduated sanctions, conflict resolution, recognized rights, and layered organization. These are not optional. They are the design principles. They apply to pastures, fisheries, CI pipelines, staging environments, monoliths, and shared databases. The commons is everywhere. The principles are the same. The implementation is local. The failure to implement is the tragedy.

---

**References:**
- Garrett Hardin, "The Tragedy of the Commons," *Science*, 1968.
- Elinor Ostrom, *Governing the Commons*, Cambridge University Press, 1990.
- Elinor Ostrom, "Beyond Markets and States: Polycentric Governance of Complex Economic Systems," Nobel Prize Lecture, 2009.
- Related posts: [Scarcity and Software Games](https://blog.hackspree.com/#scarcity-and-software-games), [Field Guide to Scarcity Games](https://blog.hackspree.com/#catalog-of-scarcity-games)

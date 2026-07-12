---
title: "Game Theory Model: The Stag Hunt"
date: 2026-07-12
slug: game-stag-hunt
summary: "Jean-Jacques Rousseau described it in 1755: two hunters must decide whether to hunt a stag together or take a rabbit alone. The stag is worth more. The rabbit is certain. If either defects, the stag escapes. The Stag Hunt models coordination under trust — and explains network effects, standard adoption, and why good architectures fail to spread."
tags: game-theory, stag-hunt, coordination, trust, network-effects
series: game-theory-models
part: 2
---

Jean-Jacques Rousseau described the Stag Hunt in his *Discourse on Inequality* (1755):

> "If it was a matter of hunting a deer, everyone well realized that he must remain faithfully at his post; but if a hare happened to pass within reach of one of them, we cannot doubt that he would have gone off in pursuit without scruple."

Two hunters. They can hunt a stag together or hunt rabbits alone. The stag requires both to cooperate. If both hunt the stag, they share a large reward. If one defects to hunt rabbits, the defector gets a rabbit — less than half a stag but guaranteed — and the cooperator gets nothing. If both hunt rabbits, each gets a rabbit.

The payoff matrix:

| | Hunt Stag | Hunt Rabbit |
|---|---|---|
| **Hunt Stag** | (10, 10) | (0, 5) |
| **Hunt Rabbit** | (5, 0) | (5, 5) |

There are two Nash equilibria: both hunt stag and both hunt rabbit. Both are stable. Neither player can unilaterally improve by switching. The stag equilibrium is Pareto-superior. The rabbit equilibrium is risk-dominant — if you're uncertain what the other will do, rabbit is safer. The game is about trust and coordination, not conflict. The challenge is moving from the rabbit equilibrium to the stag equilibrium.

## Interpretations from different branches

**Classical game theory.** The stag hunt has two pure-strategy Nash equilibria. Equilibrium selection — which one the players end up in — is not determined by the payoff structure alone. It depends on beliefs, expectations, and focal points. Schelling's concept of the focal point was developed partly to explain equilibrium selection in coordination games.

**Risk dominance (Harsanyi and Selten).** The rabbit equilibrium risk-dominates the stag equilibrium if the expected payoff of hunting rabbit, given uncertainty about the other player, exceeds the expected payoff of hunting stag. Formally: (5+5)/2 > (10+0)/2 → 5 > 5, so neither risk-dominates. In the classic stag hunt, the equilibria are payoff-equivalent under uniform uncertainty. But if the stag is worth more — say 20 — then (20+0)/2 = 10 > 5, and stag risk-dominates. The size of the prize determines whether coordination on the ambitious outcome is rational under uncertainty.

**Evolutionary game theory.** In a population playing stag hunts, which equilibrium is selected depends on the initial proportion of stag hunters. If enough players hunt stag, the stag equilibrium is reached. There is a critical threshold. Below the threshold, the population converges to rabbits. The threshold is the tipping point. The dynamics are the replicator equation.

**Network economics.** The stag hunt models adoption of technologies with network effects. A communication standard is a stag. Everyone benefits if everyone uses it. But if you're the first adopter, you bear the switching cost with no guarantee others will follow. The critical mass problem is the stag hunt in economic form.

## Software engineering interpretations

**Logging standardization.** SRE proposes a standard logging library. If all 12 services adopt it, logs become queryable across services. If only some adopt, the adopters get no benefit — their logs are standardized but they can't query across services. If nobody adopts, nothing changes. The stag is the universal query. The rabbit is keeping your own format. The critical mass problem: who goes first?

**API gateway adoption.** The platform team builds an API gateway. Each service team must decide whether to route through it. If all use it, the organization gets unified auth, rate limiting, and monitoring. If some bypass it, the bypassers move faster (no gateway latency, no configuration overhead) and the adopters get partial benefit. The stag is the unified gateway. The rabbit is direct access. The critical mass is the number of services needed for the gateway's benefits to be self-sustaining.

**Design system adoption.** The design team ships a component library. Each frontend team chooses whether to use it. If all use it, the product gets visual consistency and shared maintenance. If some don't, the non-adopters ship faster and the adopters maintain the library for everyone. The stag is consistency. The rabbit is velocity. The library needs critical mass to survive.

## Resolving the stag hunt

The stag hunt is resolved by mechanisms that raise confidence that others will cooperate. Visible early adopters signal commitment. Public commitments — "we will adopt the standard by Q3" — reduce uncertainty. Sequential adoption with increasing benefits — each new adopter increases the value for remaining players — creates the tipping dynamic. Schelling's focal points — obvious choices that everyone expects everyone else to make — select the equilibrium. "We'll all use the most popular format on npm." The format is a focal point. The popularity is the signal.

---

**References:**
- Jean-Jacques Rousseau, *Discourse on the Origin and Basis of Inequality Among Men*, 1755.
- Brian Skyrms, *The Stag Hunt and the Evolution of Social Structure*, Cambridge University Press, 2004.
- Thomas Schelling, *The Strategy of Conflict*, Harvard University Press, 1960.
- Related posts: [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games), [Scarcity and Software Games](https://blog.hackspree.com/#scarcity-and-software-games)

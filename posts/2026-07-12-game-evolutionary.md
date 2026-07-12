---
title: Evolutionary Game Theory
date: 2026-07-12
slug: game-evolutionary
summary: "Maynard Smith asked: what if players don't choose strategies rationally but inherit them genetically? Evolutionary game theory replaces rational choice with replicator dynamics. Strategies that perform well reproduce. Strategies that perform poorly die out. The Evolutionarily Stable Strategy cannot be invaded by mutants. This is why microservices survived and monoliths didn't."
tags: game-theory, evolutionary, maynard-smith, replicator-dynamics, ESS
series: game-theory-models
part: 12
---

John Maynard Smith published *Evolution and the Theory of Games* in 1982. He asked: what if players don't choose strategies rationally but inherit them genetically? The question launched evolutionary game theory. The key concepts: replicator dynamics and the Evolutionarily Stable Strategy (ESS).

Replicator dynamics describe how the proportions of different strategies in a population change over time. Strategies that earn above-average payoffs grow. Strategies that earn below-average payoffs shrink. The growth rate is proportional to the difference between the strategy's payoff and the population average. The dynamics are deterministic. Given the initial proportions and the payoff matrix, the trajectory is determined.

An Evolutionarily Stable Strategy is a strategy that, if adopted by the entire population, cannot be invaded by any small group of mutants playing a different strategy. The ESS is a refinement of Nash equilibrium for biological contexts. Every ESS is a Nash equilibrium. Not every Nash equilibrium is an ESS. The ESS requires that if a mutant strategy appears, it does worse against the existing population than the existing strategy does. The condition is stricter than Nash. The strictness is appropriate for evolution, where stability means resistance to invasion, not just mutual best response.

## Interpretations from different branches

**Biology (Maynard Smith, 1982).** The hawk-dove game models animal conflict. Hawks fight for resources. Doves display but retreat if attacked. Hawk vs. Hawk: one wins, one is injured. Hawk vs. Dove: Hawk wins. Dove vs. Dove: both share. The ESS is a mixed population — some proportion of Hawks, some of Doves. The proportion depends on the value of the resource relative to the cost of injury. The prediction matches observed behavior in multiple species. Evolution is a game. The game has an equilibrium. The equilibrium is the observed behavior.

**Economics (evolutionary game theory, Nelson and Winter, 1982).** Firms don't optimize. They follow routines — organizational habits inherited from the past. Routines that produce profits survive. Routines that produce losses are replaced. The market is the selection environment. The replicator is the firm's growth rate. The ESS is the industry equilibrium. The dynamics explain why industries converge on similar practices. The practices are not optimal in any absolute sense. They are stable against invasion by alternatives.

**Anthropology (Boyd and Richerson, 1985).** Culture evolves through imitation and social learning. Cultural variants — beliefs, practices, technologies — are strategies. Successful variants are copied. Unsuccessful variants are abandoned. The replicator dynamics are cultural transmission. The ESS is a cultural equilibrium. The dynamics explain why some cultural practices persist despite being individually costly — they are stable against invasion by alternatives, even if alternatives would be better for individuals.

**Computer science (genetic algorithms, classifier systems).** Evolutionary computation uses replicator dynamics to solve optimization problems. A population of candidate solutions. Fitness evaluation. Selection. Recombination. Mutation. Repeat. The algorithm is replicator dynamics implemented in code. The solutions evolve. The evolution finds optima that gradient-based methods miss. The search is stochastic. The convergence is evolutionary.

## Software engineering interpretations

**Architecture pattern evolution.** Microservices emerged as a mutant strategy. The monolith was the incumbent. Early microservices adopters demonstrated advantages — independent deployability, team autonomy, fault isolation. Other teams observed the success and adopted. The replicator dynamics: adoption rate proportional to observed advantage. The ESS: microservices are now the default for new systems. The monolith couldn't resist the invasion. Not because microservices are universally better. Because they are stable against re-invasion by the monolith pattern. Once adopted, the switching cost back to monolith is high. The switching cost is the invasion barrier.

**Language and framework adoption.** React emerged as a mutant in a jQuery-dominant ecosystem. Early adopters demonstrated advantages — component model, virtual DOM, unidirectional data flow. Adoption accelerated as the advantages became visible. jQuery declined. The replicator dynamics: framework popularity follows relative fitness. The fitness landscape changes as tooling, community, and hiring markets co-evolve. Today's ESS may not be tomorrow's. The dynamics continue.

**Process evolution.** Agile emerged as a mutant in a waterfall-dominant ecosystem. Early adopters demonstrated faster delivery, better responsiveness to change. Adoption accelerated. Waterfall declined — not eliminated, but no longer the default. The replicator dynamics: process fitness is measured by organizational survival. Organizations that adopted agile survived at higher rates. The surviving organizations shaped the hiring market. The hiring market shaped the training pipeline. The pipeline reinforced agile. The equilibrium is self-reinforcing.

**The persistence of patterns.** Some architecture patterns persist despite being widely criticized. The distributed monolith — microservices with tight coupling — persists because it is an ESS in certain environments. A team that inherits a distributed monolith can't unilaterally refactor to clean boundaries. The refactoring requires coordination across teams. The coordination is a public good. The individually rational choice is to work within the existing architecture. The equilibrium is suboptimal but stable. Stability doesn't mean goodness. It means resistance to change.

## The evolutionary lens

Evolutionary game theory gives the software engineer a lens: what strategies are growing in the population? What strategies are stable against invasion? The dominant architecture, the dominant language, the dominant process — they are not necessarily optimal. They are evolutionary equilibria. They persist because they resist invasion. Understanding the invasion barrier — switching cost, coordination cost, network effects — explains why they persist. Changing them requires lowering the barrier or raising the fitness of the alternative. The lowering is mechanism design. The raising is engineering. Both are evolutionary.

---

**References:**
- John Maynard Smith, *Evolution and the Theory of Games*, Cambridge University Press, 1982.
- John Maynard Smith and George Price, "The Logic of Animal Conflict," *Nature*, 1973.
- Richard Nelson and Sidney Winter, *An Evolutionary Theory of Economic Change*, Harvard University Press, 1982.
- Related posts: [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games), [Field Guide to Scarcity Games](https://blog.hackspree.com/#catalog-of-scarcity-games)

---
title: The Ultimatum Game
date: 2026-07-12
slug: game-ultimatum
summary: "One player proposes how to split $100. The other accepts or rejects. If they reject, both get nothing. The Nash equilibrium: proposer offers $1, responder accepts. Real humans: proposers offer $30-$50, responders reject offers below $20. The gap between theory and behavior is a window into fairness, emotion, and what rationality actually means."
tags: game-theory, ultimatum-game, fairness, behavioral-economics, rationality
series: game-theory-models
part: 5
---

The Ultimatum Game was introduced by Werner Güth in 1982. It is the simplest game that tests fairness. Two players. The proposer receives $100 and must offer some portion to the responder. The responder can accept — in which case both get the proposed split — or reject — in which case both get nothing. The game is played once. The players are anonymous.

The Nash equilibrium: the proposer offers the minimum possible amount ($1), and the responder accepts. Why? Because the responder, facing a choice between $1 and $0, rationally prefers $1. Knowing this, the proposer offers $1. The outcome is (99, 1). The prediction is clear. The prediction is wrong.

In thousands of replications across dozens of cultures, proposers offer 30-50% of the pie. Offers below 20% are rejected roughly half the time. The rejection rate increases as the offer decreases. The Nash equilibrium predicts 0% rejection at any positive offer. The prediction fails. The failure is systematic. The system is human.

## Interpretations from different branches

**Classical game theory.** The Nash equilibrium is (proposer offers ε, responder accepts any positive offer). The prediction relies on the assumption that players care only about their own monetary payoff. The assumption is false. The falsity is informative.

**Behavioral economics (Kahneman, Thaler).** Responders reject low offers because they value fairness. The rejection is costly — they give up money — but the cost is worth it to punish unfair behavior. The proposer anticipates this and offers a fair split to avoid rejection. The fairness norm constrains the equilibrium. The constraint is not in the payoff matrix. It is in the players' heads.

**Neuroeconomics (Sanfey et al., 2003).** fMRI studies show that receiving an unfair offer activates the anterior insula — a brain region associated with disgust — and the dorsolateral prefrontal cortex — associated with cognitive control. The brain treats unfairness as physically aversive. The rejection is not a calculated choice. It is an emotional response. The emotion is disgust. The disgust overrides the rational calculation. The override is visible in the brain.

**Cross-cultural studies (Henrich et al., 2001).** The Ultimatum Game has been played in 15 small-scale societies. The results vary enormously. In some societies, proposers offer as little as 15% and responders accept. In others, proposers offer more than 50% — hyper-fair — and responders reject both low *and* high offers. The variation tracks cultural norms about sharing, gift-giving, and market integration. The game reveals culture. The culture varies. The variation is systematic.

**Evolutionary psychology.** Fairness norms evolved in small groups where reputation mattered. The one-shot anonymous Ultimatum Game is evolutionarily novel. The brain applies reputation-based reasoning to an anonymous situation. The misapplication produces rejection of unfair offers. The brain is not wrong. The environment is novel. The adaptation is for a different world.

## Software engineering interpretations

**Salary negotiation.** The offer is the salary. The rejection is walking away. If the offer is too low, the candidate rejects — even though a low salary is better than no salary. The rejection is the Ultimatum Game. The candidate is the responder. The company is the proposer. Companies that lowball lose good candidates not because the candidates can't use the money but because the offer signals undervaluation. The signal is the information. The information is worth more than the salary difference.

**Resource allocation between teams.** The infrastructure budget must be split. The platform team proposes an allocation. The service teams can accept or escalate. Escalation costs time and political capital — both parties lose. The Nash equilibrium: platform proposes a minimal allocation, service teams accept. The observed behavior: platform proposes a fair split to avoid escalation. The fairness norm constrains the outcome.

**Review assignment.** A tech lead assigns code reviews. If the assignment is visibly unfair — one person gets all the difficult reviews — the overloaded reviewer may refuse, forcing the lead to rebalance. The refusal costs the reviewer social capital. The cost is worth it to punish the unfairness. The lead, anticipating refusal, assigns fairly. The fairness is strategic.

**The takeaway.** The Ultimatum Game teaches that humans are not purely self-interested payoff maximizers. They value fairness. They will pay to punish unfairness. The willingness to pay is a constraint on any system that allocates resources among humans. Algorithms that produce technically optimal but visibly unfair allocations will be rejected. The rejection is rational — if your model of rationality includes fairness in the utility function. The model should.

---

**References:**
- Werner Güth, Rolf Schmittberger, Bernd Schwarze, "An Experimental Analysis of Ultimatum Bargaining," *Journal of Economic Behavior & Organization*, 1982.
- Daniel Kahneman, Jack Knetsch, Richard Thaler, "Fairness and the Assumptions of Economics," *Journal of Business*, 1986.
- Joseph Henrich et al., "In Search of Homo Economicus," *American Economic Review*, 2001.
- Related posts: [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games)

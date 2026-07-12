---
title: "Game Theory Model: Repeated Games and the Folk Theorem"
date: 2026-07-12
slug: game-repeated-games
summary: "Robert Aumann proved that in infinitely repeated games, cooperation can be sustained as an equilibrium. The shadow of the future disciplines the present. The Folk Theorem says: any feasible, individually rational payoff can be an equilibrium if the game is repeated long enough. This is why teams that stay together build trust. The trust is mathematics."
tags: game-theory, repeated-games, folk-theorem, aumann, cooperation
series: game-theory-models
part: 11
---

A one-shot Prisoner's Dilemma produces defection. A repeated Prisoner's Dilemma can produce cooperation. The difference is the future. In a one-shot game, there is no tomorrow in which defection can be punished. In a repeated game, today's defection costs you tomorrow's cooperation. If tomorrow matters enough — if the discount rate is low enough — cooperation becomes the rational strategy.

The Folk Theorem, so called because it was known informally before it was formally proved, states: in an infinitely repeated game with sufficiently patient players, any feasible and individually rational payoff vector can be sustained as a Nash equilibrium. "Feasible" means the payoffs can be achieved by some combination of strategies. "Individually rational" means each player gets at least their minimax payoff — the worst the other players can impose on them. The theorem says: if the game goes on long enough, almost any outcome is possible as an equilibrium. The future is a mechanism for sustaining norms.

The mechanism is the threat of punishment. If I cooperate today, you cooperate tomorrow. If I defect today, you defect tomorrow — and the day after, and the day after. The punishment must be credible. It must be in your interest to carry it out once defection has occurred. The credibility of the punishment is the constraint on what equilibria can be sustained. The Grim Trigger — cooperate until the first defection, then defect forever — is the harshest credible punishment. Tit-for-Tat — start with cooperation, then mirror the other's last move — is gentler and more robust.

## Interpretations from different branches

**Game theory (Aumann, 1959 Nobel 2005).** Aumann's contribution was the formal analysis of repeated games with incomplete information. If players have private information, the repetition allows them to learn about each other. The learning changes the equilibrium. The Folk Theorem extends to games with imperfect monitoring — players observe noisy signals of each other's actions. The extension is technical. The implication is practical: even when you can't perfectly observe what others did, repetition enables cooperation.

**Political science (Axelrod, 1984).** Robert Axelrod's tournaments showed that Tit-for-Tat is a robust strategy. It is nice (never defects first), retaliatory (punishes defection immediately), forgiving (returns to cooperation if the other does), and clear (easy to recognize). These four properties make it effective in repeated interactions. The clarity is crucial — if the other player can't figure out your strategy, they can't adapt to it. Clarity is strategic.

**International relations (Keohane, 1984).** International cooperation is a repeated game among states. Trade agreements, arms control treaties, environmental protocols — these are equilibria sustained by the shadow of the future. The mechanism: if you violate the treaty today, we withdraw cooperation tomorrow. The mechanism works when the future is valued. It fails when states discount the future heavily — when leaders face elections, when regimes are unstable, when the long-term benefits of cooperation accrue to successors rather than incumbents. The discount rate is political. The politics determine whether cooperation is sustainable.

**Organizational behavior.** Company culture is a repeated-game equilibrium. Norms of collaboration, knowledge-sharing, and mutual support are sustained by the expectation of continued interaction. When turnover is high, the shadow of the future shortens. Cooperation declines. The decline is attributed to "bad culture." The culture is a symptom. The cause is the shortened horizon. Fix the horizon. The culture follows.

## Software engineering interpretations

**Team continuity.** A stable team is a repeated game. Members expect to work together indefinitely. Cooperation is an equilibrium — helping a colleague today is repaid tomorrow. A team with high turnover is a sequence of short-horizon games. Cooperation is fragile — the new person hasn't built the history, the departing person won't face the consequences of defection. The Folk Theorem predicts that stable teams will have more cooperation. The prediction is correct. The mechanism is the horizon.

**Inter-service API stability.** Two services with a long history of mutual dependence are in a repeated game. Breaking the API today costs you in future coordination. The equilibrium is stability. A service consumed by many anonymous clients is in a one-shot game with each. Breaking the API harms each client individually but none enough to punish. The equilibrium is instability — the provider changes the API when convenient. The difference is the horizon. The horizon is structural.

**Code review reciprocity.** Engineers who review each other's code are in a repeated game. Reviewing thoroughly today earns thorough reviews tomorrow. Skimming today earns skimmed reviews tomorrow. The equilibrium is a norm of thoroughness — if the team is stable and the horizon is long. The norm emerges without being mandated. The mandate is unnecessary. The horizon is sufficient.

**Cross-team collaboration.** Two teams that expect to work together for years develop informal cooperation — shared understanding, mutual accommodation, the benefit of the doubt. Two teams thrown together for a single project have no shadow of the future. Cooperation must be formalized — detailed specs, explicit contracts, escalation paths. The formality is the substitute for the missing horizon. The formality is costly. The horizon was free.

## The design implication

If you want cooperation, lengthen the horizon. Stable teams. Long-term ownership. Continuity of relationships. If you can't lengthen the horizon, simulate it: automated contract testing makes defection immediately visible, creating a repeated-game payoff structure. SLAs with penalties bring the future cost of defection into the present. The mechanisms are substitutes for the missing horizon. The best mechanism is the horizon itself. The horizon is free.

---

**References:**
- Robert Aumann, "Acceptance Speech," Nobel Prize, 2005.
- Robert Axelrod, *The Evolution of Cooperation*, Basic Books, 1984.
- Drew Fudenberg and Eric Maskin, "The Folk Theorem in Repeated Games with Discounting or with Incomplete Information," *Econometrica*, 1986.
- Related posts: [Cooperation is logical](https://blog.hackspree.com/#cooperation-is-logical), [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games)

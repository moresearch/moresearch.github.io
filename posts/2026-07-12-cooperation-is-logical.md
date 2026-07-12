---
title: In a single encounter, confrontation is logical
date: 2026-07-12
slug: cooperation-is-logical
summary: "Robert Aumann won the Nobel Prize for proving that cooperation is rational in repeated games. 'In a single encounter, confrontation is the logical move; but when the interaction will occur repeatedly, cooperation is the logical behavior.' This is why teams that stay together build trust. The trust is mathematics."
tags: aumann, game-theory, cooperation, repeated-games, teams
---

Robert Aumann shared the 2005 Nobel Prize in Economics with Thomas Schelling. His contribution was the mathematics of repeated games. The insight is compressed to a sentence:

> "In a single encounter, confrontation is the logical move; but when the interaction will occur repeatedly, cooperation is the logical behavior."

In a single encounter — a one-shot game — the rational strategy is to defect. There is no future in which the other player can punish you. There is no shadow of tomorrow to discipline today. You take what you can get. They would do the same. Both of you know this. Both defect. Both get the suboptimal outcome. The Prisoner's Dilemma is a one-shot game. The dilemma is real.

In a repeated encounter — an iterated game with no known end — the calculus changes. Defection today costs you cooperation tomorrow. The other player will remember. They will punish you in future rounds. The future rounds matter because the game continues indefinitely. The shadow of the future disciplines the present. Cooperation becomes rational. Not because the players are virtuous. Because the mathematics changed. The payoff structure changed. The change is in the repetition. The repetition is the mechanism.

## Teams as repeated games

Software teams are repeated games. The same people work together sprint after sprint, quarter after quarter, year after year. The game has no predetermined end. The shadow of the future is long. This is why teams that stay together develop trust. The trust is not a personality trait. It is an equilibrium. Each person has learned that cooperation produces better long-run outcomes than defection. Each person expects others to cooperate because others have learned the same thing. The expectation is rational. The cooperation is stable.

Teams that churn cannot sustain this equilibrium. When people leave and join frequently, the effective horizon shortens. The new person doesn't have a history of cooperation with the existing team. The existing team doesn't know if the new person will cooperate. The uncertainty reduces the expected value of cooperation. The shadow of the future is shorter because the future with this person is uncertain. The shorter shadow produces less cooperation. Less cooperation produces worse outcomes. The churn is costly in ways that appear on no balance sheet. The cost is real. The mathematics predicts it.

Aumann's insight is that cooperation doesn't require central enforcement. It doesn't require a manager mandating collaboration. It doesn't require HR programs or team-building exercises. It emerges from the structure of the interaction. Repeated interaction with no known end produces cooperation as an equilibrium. The structure is the mechanism. The mechanism produces the behavior. The behavior looks like culture. It is mathematics.

## API contracts as repeated games

The same logic applies to inter-service communication. Service A depends on Service B's API. If the interaction is one-shot — A calls B once and never again — B has no incentive to maintain a stable API. B can change the API whenever it wants. A's dependence is A's problem. The interaction is one-shot. Defection is rational.

If A and B will interact repeatedly — A will call B's API every day for years — the calculus changes. B knows that breaking the API today costs B in the future. A will be angry. A will escalate. A might build their own version of B. The future cost of breaking the API exceeds the present benefit. B maintains the API. The stability is not because B is considerate. It is because the game is repeated. The repetition changes the payoff.

This is why internal APIs between teams that have worked together for years are more stable than external APIs consumed by strangers. The internal teams are in a repeated game. The external consumers are in a one-shot game from the provider's perspective. The provider doesn't feel the future cost of breaking the API because the future cost is diffused across thousands of anonymous consumers. The consumers can't coordinate to punish the provider. The coordination problem prevents the repeated-game equilibrium from forming. The provider defects — changes the API, deprecates the endpoint, raises the price. The consumers suffer individually. The suffering is aggregate but uncoordinated. The provider doesn't feel it. The one-shot structure produces the defection. The structure is the problem.

## Mechanism design for repeated games

If the natural structure produces one-shot interactions where repeated interactions would produce better outcomes, change the structure. This is mechanism design. Automated contract testing changes API interactions from one-shot to repeated. Every build runs the contract tests. Every breaking change is immediately visible. The visibility creates a repeated-game payoff structure. The provider can't defect invisibly. The defection is detected. The detection has a cost — the build breaks, the provider must fix it. The cost is immediate. The immediacy simulates repetition. The simulation changes the behavior.

SLAs with penalty clauses do the same. The penalty is the future cost of defection, brought forward to the present. The provider who breaks the SLA pays now. The payment is the shadow of the future, compressed into a contract. The contract is a mechanism for making one-shot interactions behave like repeated ones. The mechanism substitutes for the missing future.

Code review is a repeated game. The author and the reviewer will interact again. The author who ignores feedback today will receive less helpful feedback tomorrow. The reviewer who is needlessly harsh today will find their reviews ignored tomorrow. The mutual expectation of future interaction disciplines present behavior. The discipline is automatic. It doesn't require rules. It requires continuity.

## The half-life of trust

Aumann's mathematics implies that trust has a half-life. It decays when the future becomes uncertain. A reorg that shuffles teams resets the repeated-game equilibrium. The new teams have no history of cooperation. They must rebuild it. The rebuilding takes time. During the rebuilding, cooperation is suboptimal. The system performs worse. The reorg's cost includes the lost cooperation during the rebuilding period. Nobody accounts for this cost. The cost is real.

A layoff that cuts a team in half shortens the shadow of the future for everyone who remains. The remaining people now know the game can end unexpectedly. The unexpected end converts an infinite-horizon game into an uncertain-horizon game. Uncertain horizons produce less cooperation than infinite horizons. The layoff's cost includes the reduced cooperation among survivors. Nobody accounts for this cost. The cost is real.

A team that knows it will be disbanded in six months is in a finite-horizon game. Finite-horizon games unravel from the end. In the final sprint, defection is rational — there is no future to punish it. In the second-to-last sprint, defection is rational because defection in the final sprint is already expected. The logic propagates backward. By induction, cooperation collapses in the first sprint. The collapse is mathematical. The team's morale didn't fail. The structure changed. The structure produced the outcome.

The Hayekian manager understands this. They preserve team continuity not because "culture matters" but because continuity is the structural precondition for cooperation. The structure produces the behavior. Change the structure. The behavior changes. The change is predictable. Aumann gave us the mathematics. The mathematics is clear. Most organizations ignore it.

---

**References:**
- Robert Aumann, "Acceptance Speech," Nobel Prize in Economics, 2005.
- Robert Aumann and Michael Maschler, *Repeated Games with Incomplete Information*, MIT Press, 1995.
- Anatol Rapoport and Albert Chammah, *Prisoner's Dilemma*, University of Michigan Press, 1965.
- Related posts: [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games), [On Scarcity](https://blog.hackspree.com/#scarcity), [Design the Game](https://blog.hackspree.com/#scarcity-and-mechanism-design)


Engineering is the discipline of building things that work within constraints. Every topic on this blog — operating systems, AI models, trading infrastructure, research labs, innovation economics — is examined through the lens of systems design. The lens is engineering. The method is: understand the constraints, design within them, verify the design works, iterate. The domain provides the specifics. The method is universal.

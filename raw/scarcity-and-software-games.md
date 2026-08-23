---
title: The Architecture Is a Game
date: 2026-07-12
slug: scarcity-and-software-games
summary: "Every software situation is a game. The service boundary game, the microservices migration game, the deployment chicken game, the logging stag hunt — each has a game-theoretic structure. Recognizing it tells you what to do."
tags: game-theory, software-architecture, microservices, coordination
series: scarcity
part: 4
---

The most dangerous games are the ones you don't know you're playing. Nobody told you that logging format choice was a stag hunt. Nobody announced that the rewrite proposal was a signaling game. The games are invisible. Their outcomes are visible — the broken APIs, the fragmented standards, the stalled migrations. The outcomes are blamed on individuals. "Team B should have communicated better." "The SRE team should have enforced the standard." But the individuals were playing rationally given the game they were in. Blaming the player is easier than recognizing the game. Recognizing the game is easier than changing it. Change the game.

Software engineering is a multiplayer game. The players are teams, services, organizations, companies. Each player makes choices under conditions of scarcity. Each player's outcome depends on the choices of others. Recognizing the game you're in is the first step to playing it well.

## The service boundary game

Thomas Schelling observed that "what makes many agreements enforceable is only the recognition of future opportunities for agreement that will be eliminated if mutual trust is not created and maintained, and whose value outweighs the momentary gain from cheating in the present instance." API contracts are agreements. Automated testing makes cheating visible. Visibility eliminates the momentary gain. The contract becomes enforceable because the future cost of breaking it exceeds the present benefit. This is mechanism design as Schelling described it before the term existed.

Team A builds Service A. Team B builds Service B. Service A depends on Service B. Team B changes Service B's API without telling Team A. Service A breaks. Team A is angry. Team B is surprised.

This is a coordination game with asymmetric information. Team B didn't know what Team A depended on. Team A assumed Team B wouldn't change the API without notice. The assumptions were incompatible. The failure is a Nash equilibrium — neither team can unilaterally improve their outcome given what the other is doing. Team A can't make Team B communicate better. Team B can't make Team A depend on fewer things. The equilibrium is suboptimal. The solution is mechanism design: automated contract testing that makes breaking changes immediately visible. Visibility changes the payoff. The equilibrium shifts.

## The microservices migration game

The monolith works but is increasingly expensive to change. Each team wants to extract their service. Extraction requires coordination with other teams. Coordination is costly. The cost falls on the team doing the extraction. The benefit accrues to all teams.

This is a public goods game. Each team would benefit if everyone extracted. Each team would prefer someone else pay the coordination cost. Individually rational: wait. Collectively optimal: coordinate. The gap is the Prisoner's Dilemma scaled to organizational architecture. The dilemma produced the distributed monolith — services extracted without clean boundaries, communicating through a complete call graph, with all the monolith's coupling and all the network's latency. Nobody wanted this. The game produced it.

## The deployment chicken game

Two teams both want to deploy Friday afternoon. Both know Friday deploys are risky — if something breaks, on-call spends the weekend fixing it. Both want their feature in before the weekend. If both deploy and nothing breaks, both win. If both deploy and something breaks, both lose. If one deploys and one waits, the deployer wins and the waiter deploys Monday.

Both deploying is the crash outcome. The organizational rule "no Friday deploys" is mechanism design. It removes the choice. Before the rule: Friday deploys. After: Monday deploys. The mechanism changed the equilibrium.

## The staging environment battle of the sexes

Team A wants staging for integration testing. Team B wants staging for customer demos. Both want staging. Both prefer any coordinated solution to constant conflict. Both prefer their own preferred time.

Battle of the sexes. Two Nash equilibria: Team A's schedule or Team B's schedule. The solution is a coordination mechanism — a booking calendar, a dedicated demo environment, a policy. The mechanism selects an equilibrium. Before: conflict. After: coordination. The mechanism worked because it made defection visible.

## The logging library stag hunt

Twelve services. Each uses its own logging format. SRE proposes a standard library. If all adopt, logs become queryable across services. If some adopt and others don't, adopters get no benefit — their logs are standardized but they still can't query across services. If nobody adopts, nothing changes.

Stag hunt. The stag is cross-service observability. The rabbit is keeping your own format. The stag hunt succeeds when early adopters reach critical mass. Once enough services adopt, the benefit to remaining services exceeds switching cost. The equilibrium tips. The tipping point is a property of network effects.

## The rewrite signaling game

Team A proposes rewriting a legacy service. The proposal is costly to evaluate — architecture reviews, specs, stakeholder meetings. Team A has private information: genuine belief the rewrite is necessary, or boredom with legacy work.

Signaling game. A costly signal separates sincere from insincere. Requiring a working prototype before the architecture review is a costly signal. Sincere teams pay the cost. Insincere teams won't. The signal screens. Pay the price. Get the review.

## The regression test commons game

The regression test suite is a common-pool resource. Everyone benefits from tests. Everyone benefits from fast tests. Adding tests makes the suite slower. The cost is shared. The benefit of your test is yours alone.

Individually rational: add tests. Collectively optimal: add only high-value tests. The commons is overgrazed. The suite grows. Build times increase. Hardin described the tragedy in 1968. Your CI pipeline is living it. The solution: a test budget per service, periodic culling of low-value tests, a rule requiring historical failure catch rate to justify additions. The mechanism is the institutional response to a tragedy of the commons.

## The open-source game

Everyone benefits from open-source software. Contributing costs time. Using costs nothing. Individually rational: use without contributing. Collectively optimal: everyone contributes. Public goods game. Resolved by reputation, corporate sponsorship, intrinsic motivation. The free-rider problem is managed, never solved. GitHub sponsorships, open-source foundations, corporate OSPOs — the institutions exist because the game exists. The game exists because scarcity exists.

## The platform pricing Stackelberg game

A platform team sets API pricing for internal services. Service teams respond by choosing how much to consume. The platform team moves first. Service teams move second.

Stackelberg game. The leader (platform) chooses price anticipating the followers' (services') responses. Set price too high: services build their own. Too low: platform is underfunded. The optimal price is where marginal cost of providing the service equals the marginal value to consumers. The calculation is economic. The implementation is an internal API with metered billing. The billing is mechanism design.

## Recognize the game

Every situation is a game. The game has a structure. The structure determines the likely outcome. If you don't like the outcome, change the game. Mechanism design is the tool for changing games. Automated contract testing changes the service boundary game. "No Friday deploys" changes the deployment chicken game. A booking calendar changes the staging environment game. A test budget changes the regression test commons game. Each mechanism changes the payoffs. Changed payoffs change behavior. Changed behavior is the point.

---

**This is part 4 of a 7-part series on scarcity and software.**
- [Part 1: On Scarcity](https://blog.hackspree.com/#scarcity)
- [Part 2: On Games](https://blog.hackspree.com/#scarcity-and-games)
- [Part 3: On Software Engineering Economics](https://blog.hackspree.com/#scarcity-and-software-economics)
- [Part 5: On AI and Mechanism Design](https://blog.hackspree.com/#scarcity-and-mechanism-design)
- [Part 6: On Practice](https://blog.hackspree.com/#scarcity-in-practice)
- [Part 7: The Catalog of Games](https://blog.hackspree.com/#catalog-of-scarcity-games)

**References:**
- Garrett Hardin, "The Tragedy of the Commons," *Science*, 1968.
- John Maynard Smith, *Evolution and the Theory of Games*, Cambridge University Press, 1982.
- Related posts: [Henney's Microservices](https://blog.hackspree.com/#kevlin-henney), [NATS pub/sub beats REST](https://blog.hackspree.com/#nats-pubsub-microservices)


Scarcity is the universal engineering constraint. Time, attention, compute, complexity — every engineering decision is made within a budget. The budget is economic. The engineer who doesn't track the budget makes decisions blind. The engineer who tracks it makes decisions with full knowledge of the trade-off. The trade-off is the decision. The budget is the constraint. Scarcity is the unifying principle.

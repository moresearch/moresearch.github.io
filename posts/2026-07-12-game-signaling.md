---
title: Signaling Games
date: 2026-07-12
slug: game-signaling
summary: "One player has private information. They take an action that may reveal it. The action is a signal. A signal works if it's costly enough that only the 'good' type would send it. Michael Spence won a Nobel for proving that education can be a pure signal — it doesn't teach skills, it proves you had the skills to get in. In software: the RFC that's too detailed, the rewrite proposal that comes with a prototype."
tags: game-theory, signaling, spence, private-information, separating-equilibrium
series: game-theory-models
part: 6
---

A signaling game has two players. The sender has private information about their type. The sender takes an action — sends a signal. The receiver observes the signal and responds. The sender's payoff depends on the receiver's response and the sender's type. The receiver's payoff depends on the sender's type and the receiver's response.

The classic model is Michael Spence's job market signaling (1973). A worker knows their own productivity (high or low). The employer doesn't. The worker can acquire education — a costly signal. If education is more costly for low-productivity workers than for high-productivity workers, there exists a separating equilibrium: high types acquire education, low types don't, employers pay high wages to educated workers and low wages to uneducated workers. The education doesn't need to increase productivity. It only needs to be differentially costly. The differential cost is what makes the signal informative.

## Interpretations from different branches

**Economics (Spence, 1973 Nobel 2001).** The separating equilibrium exists when the single-crossing property holds — the marginal cost of signaling differs across types. Education signals productivity not because education teaches skills but because it is harder for low-productivity people to acquire. The hardship is the signal. The credential is the evidence of hardship.

**Biology (Zahavi's handicap principle, 1975).** The peacock's tail is a signal of genetic fitness. The tail is costly — it requires energy, attracts predators, impedes movement. Only a genetically fit peacock can afford such a handicap. The cost is the signal. The tail doesn't help the peacock survive. It proves the peacock *can* survive despite the tail. The handicap is honest because it's too expensive to fake.

**Political science (Fearon, 1994).** Audience costs in international relations are signaling games. A leader who makes a public threat — "we will retaliate" — pays an audience cost if they back down. The cost is domestic political damage. Only a leader who is genuinely committed can afford to make the threat. The threat is credible because backing down is costly. The cost is the signal.

**Computer science (mechanism design).** Screening is the reverse of signaling. In screening, the uninformed party moves first, offering a menu of contracts. Each type self-selects into the contract designed for it. The menu separates types by making each contract optimal for the intended type. Screening is mechanism design. Signaling is sender-driven. Screening is receiver-driven.

## Software engineering interpretations

**The rewrite proposal.** A team proposes rewriting a legacy service. The proposal is cheap talk — anyone can propose. The architecture review must decide whether the team genuinely believes the rewrite is necessary or is bored and wants greenfield work. A prototype requirement is a screening mechanism. A sincere team will build the prototype. An insincere team won't. The prototype is costly. The cost separates types.

**The detailed RFC.** Writing a thorough RFC — with trade-off analysis, migration plans, and risk assessment — is costly. It takes days. The cost signals seriousness. A team that writes a one-paragraph proposal is signaling low investment. The architecture review discounts the proposal accordingly. The RFC's length is not about information transfer. It is about signaling commitment.

**Open-source contributions as job market signaling.** A developer contributes to a well-known open-source project. The contribution is publicly visible. It signals skill to potential employers. The contribution is costly — it takes time outside work. The cost signals passion and competence. The signal works because it's harder to fake than a resume bullet point. The code doesn't lie. The commit history is the credential.

**The production incident response.** How an engineer handles a production incident signals competence to the entire organization. The signal is costly — incidents are high-stress, high-visibility, and occur at inconvenient times. The engineer who stays calm, diagnoses systematically, and communicates clearly is signaling a type that cannot be faked under pressure. The pressure is the cost. The calm is the signal.

## Separating vs. pooling equilibria

A signaling game has two kinds of equilibria. In a separating equilibrium, different types send different signals. The receiver can infer the type from the signal. In a pooling equilibrium, all types send the same signal. The receiver learns nothing. Which equilibrium emerges depends on the cost structure and the prior distribution of types. If the cost of the signal is too low for all types, everyone sends it — credential inflation. If too high, nobody does. The sweet spot is where the cost is differentially burdensome.

In software: if writing a prototype is too easy, everyone writes one. The prototype stops separating types. If it's too hard, nobody writes one. The review process loses the signal. The optimal screening cost is calibrated to the distribution of sincere and insincere proposers. The calibration is mechanism design.

---

**References:**
- Michael Spence, "Job Market Signaling," *Quarterly Journal of Economics*, 1973.
- Amotz Zahavi, "Mate Selection — A Selection for a Handicap," *Journal of Theoretical Biology*, 1975.
- James D. Fearon, "Domestic Political Audiences and the Escalation of International Disputes," *American Political Science Review*, 1994.
- Related posts: [Scarcity and Games](https://blog.hackspree.com/#scarcity-and-games), [Scarcity and Software Games](https://blog.hackspree.com/#scarcity-and-software-games)

---
title: Prediction Markets
date: 2026-07-12
slug: fi-prediction-markets
summary: "A prediction market lets people bet on the outcome of future events. The market price is a probability estimate. The functional origin is 16th-century papal conclave betting. The modern incarnation is Polymarket, which correctly called the 2024 U.S. election while polls were uncertain. Markets aggregate dispersed knowledge. The aggregation is the product."
tags: defi, prediction-markets, polymarket, information, hayek
series: crypto-fi-instruments
part: 9
---

A prediction market is a market where participants trade contracts that pay out based on the outcome of a future event. A contract that pays $1 if a candidate wins an election and $0 if they lose. If the contract trades at $0.60, the market's implied probability of the candidate winning is 60%. The price is a probability. The market is a forecasting tool.

The key insight, formalized by Friedrich Hayek in "The Use of Knowledge in Society" (1945), is that markets aggregate dispersed information. No single individual knows the true probability of an event. Each individual holds fragments of relevant knowledge — a pollster has survey data, a journalist has sources, a trader has a model, a local observer has on-the-ground impressions. The market price aggregates these fragments into a single number. The number is more accurate than any individual's estimate. The accuracy is the market's product.

## The functional origin: papal conclave betting

Betting on papal elections was common in 16th-century Rome. Cardinals gathered in conclave. The outside world speculated on the outcome. Bookmakers offered odds. The odds fluctuated as news leaked from the conclave — a cardinal was seen visiting another's cell, a delegation arrived with a message from a foreign monarch. The betting markets aggregated the leaks into a probability estimate. The estimate was often more accurate than the assessments of diplomats and ambassadors, who had access to more formal information but lacked the market's ability to weigh competing signals.

The tradition continued. Betting on elections was widespread in the United States in the 19th and early 20th centuries, conducted through organized exchanges and informal bookmaking. The markets were suppressed by anti-gambling laws in the mid-20th century. They reemerged in the 1980s with the Iowa Electronic Markets, a small-scale academic prediction market for U.S. elections. The IEM demonstrated that prediction markets could forecast elections more accurately than polls. The demonstration was academic. The adoption was limited by regulatory constraints.

## Polymarket and the 2024 election

Polymarket, launched in 2020, is a decentralized prediction market built on Polygon. Users deposit USDC and trade shares in event outcomes. The market resolves when an oracle — currently UMA's optimistic oracle — reports the outcome. The oracle is the source of truth. The trust in the oracle is the trust in the market.

The 2024 U.S. presidential election was Polymarket's breakthrough moment. The market correctly called the election outcome while traditional polls showed a statistical tie. The market's implied probability moved sharply in the final weeks before the election, aggregating signals that polls were missing or misweighting. The market's accuracy generated mainstream attention. Trading volume exceeded $1 billion. The market became a news source in its own right — journalists cited Polymarket odds alongside polling averages.

The success validated the Hayekian thesis: markets aggregate knowledge that no single institution possesses. Polls ask people what they think. Markets ask people what they'll bet on. The difference is skin in the game. Skin in the game improves accuracy. The improvement is the efficiency of markets.

## The mechanism

Prediction markets use a simple mechanism. A binary outcome market has two tokens: YES and NO. Each pays $1 if correct. The tokens trade freely. The price of YES is the market's probability estimate. A trader who believes the true probability is higher than the market price buys YES. A trader who believes it's lower buys NO. The trading moves the price toward the traders' collective estimate. At resolution, the correct token is redeemable for $1. The incorrect token is worthless.

The mechanism is incentive-compatible. Traders profit from correcting the market's errors. The profit motive drives information into prices. The information is the traders' private knowledge. The private knowledge becomes public through the price. The price is the public good. The market produces the public good as a byproduct of private profit-seeking. The byproduct is the innovation.

## The reference

Robin Hanson, "Shall We Vote on Values, But Bet on Beliefs?" (2013). Hanson is the leading academic advocate for prediction markets. His proposal: futarchy — a form of government where elected officials set goals and prediction markets determine which policies will achieve them. The proposal is radical. The underlying logic is mainstream: markets aggregate information better than committees. Polymarket is the partial implementation of Hanson's vision. The implementation is for elections, not policy. The extension to policy is the next frontier.

## The engineering connection

Prediction markets are information aggregation mechanisms. The price is the signal. The traders are the sensors. The market is the aggregator. This is the same architecture as a distributed monitoring system: each node observes a fragment of the system state, reports its observation, and the aggregator produces a unified view. The market does this without a central aggregator. The price does the aggregation. The price is the emergent output of millions of independent observations, weighted by conviction (position size).

The oracle that resolves the market is a trust-minimized data feed — the same pattern as an API contract. The oracle reports the outcome. The market settles. If the oracle is wrong, the market produces wrong outcomes. The oracle's reliability is the system's reliability. Engineering reliable oracles — UMA's optimistic oracle, Chainlink's decentralized oracle networks — is a distributed systems problem. The problem is: how do you get truthful reports from potentially adversarial reporters? The solution is mechanism design: stake, challenge periods, economic incentives for honesty. The mechanism is the same as staking in proof-of-stake. The engineering is the same. The domain is different.

Prediction markets also teach a software engineering lesson about decision-making. A team estimating a project timeline is a prediction market without prices. Each person has private information. The information is aggregated through discussion, not through trading. Discussion is a worse aggregator than prices because discussion rewards confidence, not accuracy. The person who speaks most confidently shapes the estimate. The person who knows the truth but speaks quietly is ignored. A prediction market would surface the quiet person's knowledge through their willingness to bet on it. The bet is the signal. The price is the aggregation. The engineering team that replaces estimation meetings with internal prediction markets would produce better estimates. No team does this. The technology exists. The adoption is zero.

---

**References:**
- Robin Hanson, "Shall We Vote on Values, But Bet on Beliefs?" 2013.
- Friedrich Hayek, "The Use of Knowledge in Society," *American Economic Review*, 1945.
- Polymarket, "Polymarket Documentation."
- Related posts: [The knowledge is dispersed](https://blog.hackspree.com/#dispersed-knowledge), [On Scarcity](https://blog.hackspree.com/#scarcity)

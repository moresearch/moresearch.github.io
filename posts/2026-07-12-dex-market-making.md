---
title: Market Making
date: 2026-07-12
slug: dex-trading-market-making
summary: "A market maker posts bids and asks and stands ready to trade. The profit is the spread. The risk is adverse selection — informed traders who trade against you when the price is about to move. The Avellaneda-Stoikov model formalized this trade-off in 2008. It is now the standard model for algorithmic market making in both traditional and crypto markets."
tags: dex, market-making, avellaneda-stoikov, adverse-selection, inventory
series: dex-trading
part: 6
---

A market maker provides liquidity. They post a bid — a price at which they will buy. They post an ask — a price at which they will sell. They earn the spread between the bid and the ask. They lose when the price moves against their inventory. The market maker's art is setting the bid and ask such that spread income exceeds inventory losses over time.

The market maker is the counterparty to every impatient trader. The trader who must buy now crosses the spread and pays the ask. The trader who must sell now crosses the spread and receives the bid. The market maker absorbs the order flow imbalance. When there are more buyers than sellers, the market maker accumulates a short position — selling to buyers, hoping the price doesn't rise too far before sellers arrive. When there are more sellers than buyers, the market maker accumulates a long position. The inventory is the risk. The spread is the compensation for bearing it.

## The functional origin: the jobber

The market maker's role emerged on the London Stock Exchange in the 19th century. The jobber was a member who made markets in specific securities. The jobber quoted two prices — the bid and the ask — and stood ready to trade at those prices with any broker who approached. The jobber did not deal with the public. The jobber dealt only with brokers. The brokers dealt with the public. The jobber's profit was the spread. The jobber's risk was adverse selection: a broker with better information about the security's true value would trade against the jobber's quote, profiting at the jobber's expense.

The jobber's defense was the spread. A wider spread compensated for higher adverse selection risk. A narrower spread attracted more order flow. The optimal spread balanced the two. The jobber who set spreads too wide lost business. The jobber who set spreads too narrow lost money to informed traders. The balance was learned through experience. The learning was trial and error. The errors were expensive. The experience was the jobber's edge.

The electronic market maker of today is the jobber, automated. The spread is set by an algorithm. The algorithm reads the order book, estimates adverse selection risk, tracks inventory, and adjusts quotes in microseconds. The algorithm is faster than any human jobber. It is also less judgmental. The human jobber could sense when a broker was informed — the broker was nervous, the broker was eager, the broker asked for an unusually large size. The algorithm sees only the order flow. The signal is statistical. The edge is quantitative. The quantification is the subject of this post.

## The Avellaneda-Stoikov model

Marco Avellaneda and Sasha Stoikov published "High-Frequency Trading in a Limit Order Book" in 2008. The paper formalized the market maker's problem as stochastic optimal control. The market maker chooses bid and ask quotes to maximize expected utility of terminal wealth, subject to inventory risk and adverse selection. The solution is a closed-form approximation: the optimal bid and ask are functions of the market maker's current inventory, the volatility of the asset, the market maker's risk aversion, and the intensity of order arrival.

The key insight: the market maker should skew quotes away from their inventory. If the market maker is long — has bought more than they've sold — they should lower their bid (less willing to buy more) and lower their ask (more willing to sell). The skew reduces inventory risk. If the market maker is short, they should raise both bid and ask. The spread widens with volatility and risk aversion. The spread narrows with order arrival intensity — more competition means tighter spreads.

The model is implemented in production market-making systems. It is the standard. It is taught in quantitative finance programs. It is adapted for crypto with modifications: gas costs replace exchange fees, AMM curves replace order book depth, and the discrete block time of blockchains replaces the continuous time of traditional markets. The adaptations are engineering. The core model is the same. The model works. The work is in the calibration.

## Market making in crypto

Crypto market making has three forms:

**CEX market making.** Traditional order book market making on centralized exchanges. The market maker runs a server colocated with the exchange, streams order book updates, adjusts quotes in microseconds. The model is Avellaneda-Stoikov with exchange-specific calibrations. The competition is intense. The margins are thin. The thin margins are the evidence of competition.

**AMM liquidity provision.** Providing liquidity to a constant-function market maker. The LP deposits tokens into a pool. The pool's formula sets the price. The LP earns fees from trades. The LP bears impermanent loss. The LP's problem is similar to the market maker's: earn fee income while managing inventory risk. The difference: the AMM LP doesn't set the spread. The formula sets the spread. The LP only chooses the range (in Uniswap V3) and the amount. The LP's optimization is passive — choose parameters, deposit, wait. The active management is the subject of concentrated liquidity strategies.

**On-chain order book market making.** Running a market-making bot on a DEX with an on-chain order book (Serum, dYdX, Hyperliquid). The bot posts bids and asks as transactions. Each quote update costs gas. The gas cost constrains the update frequency. The constraint creates a trade-off: update frequently for tighter spreads (higher gas cost, more accurate quotes) or update infrequently to save gas (wider spreads to protect against adverse selection). The trade-off is the gas cost of decentralization. The gas cost is paid by the market maker. The cost is passed to traders in wider spreads. The wider spreads are the price of trustlessness.

## The reference

Marco Avellaneda and Sasha Stoikov, "High-Frequency Trading in a Limit Order Book," *Quantitative Finance*, 2008. The foundational paper on algorithmic market making. It is 25 pages. It assumes knowledge of stochastic calculus. The key result — the optimal quote skew formula — is implementable in a few lines of code. The implementation is the easy part. The calibration — estimating volatility, order arrival intensity, adverse selection — is the hard part. The calibration is the edge. The edge is the market maker's competitive advantage. The advantage erodes as competitors adopt the same model. The erosion is the subject of the arms race.

---

**References:**
- Marco Avellaneda and Sasha Stoikov, "High-Frequency Trading in a Limit Order Book," *Quantitative Finance*, 2008.
- Larry Harris, *Trading and Exchanges*, Oxford University Press, 2003.
- Related posts: [AMMs](https://blog.hackspree.com/#dex-trading-amms), [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

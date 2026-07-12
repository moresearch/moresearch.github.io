---
title: Perpetual Futures
date: 2026-07-12
slug: fi-perpetuals
summary: "The perpetual futures contract — a futures contract with no expiry — is the most traded financial instrument in crypto. It was invented by Robert Shiller in 1992 for real estate indices. It was adapted by BitMEX in 2016 for Bitcoin. The funding rate mechanism that keeps the contract price close to the spot price is a work of economic engineering."
tags: defi, perpetuals, futures, funding-rate, bitmex
series: crypto-fi-instruments
part: 3
---

The perpetual futures contract is a futures contract with no expiration date. Unlike a traditional futures contract — which settles on a specific date, at which point the contract converges to the spot price — a perpetual contract never settles. The position can be held indefinitely. The mechanism that keeps the perpetual price close to the spot price is the funding rate: a periodic payment between long and short positions. If the perpetual trades above spot, longs pay shorts. The payment incentivizes selling the perpetual and buying spot, which pushes the perpetual price down toward spot. If the perpetual trades below spot, shorts pay longs. The payment incentivizes the reverse. The funding rate is the mechanism. The mechanism replaces expiration as the convergence force.

The perpetual is the most traded instrument in crypto. Daily volumes regularly exceed $100 billion across centralized exchanges. The volume exceeds spot trading volume by a factor of 3-5×. The perpetual is the dominant instrument for speculation, hedging, and leveraged trading. Its dominance is a function of its design: no expiration means no roll costs, no settlement dates to manage, no term structure to model. The perpetual is the simplest possible futures contract. The simplicity drove adoption.

## The functional origin: Japanese rice futures

The first futures contracts were traded on the Dōjima Rice Exchange in Osaka, Japan, in the 17th century. Rice was the currency of feudal Japan — samurai were paid in rice stipends, taxes were collected in rice, wealth was measured in rice. The price of rice fluctuated with the harvest. Rice futures allowed merchants and samurai to lock in prices in advance. The contracts had expiration dates. At expiration, the contract converged to the spot price of rice in Osaka.

The Dōjima exchange was the first organized futures market in the world. It predated the Chicago Board of Trade by 150 years. It had standardized contracts, a clearinghouse, and mark-to-market settlement. The infrastructure was sophisticated. The principle — a contract for future delivery at a price agreed today — was the foundation of all derivatives markets that followed.

The perpetual futures contract eliminates the expiration that the Dōjima exchange institutionalized. The elimination is an innovation. The innovation changes the market structure: without expiration, the cost of maintaining a position is reduced to the funding rate. The funding rate is paid periodically — typically every 8 hours. The rate is determined by the market, not by a counterparty. The rate is transparent. The transparency enables algorithmic trading strategies that would be impractical with traditional futures.

## The funding rate mechanism

The funding rate is the economic engine of the perpetual. It is calculated as the difference between the perpetual price and the spot price, scaled by a funding interval. The formula: funding rate = (perpetual price - spot price) / spot price / funding intervals per day. If the perpetual trades at a 0.1% premium, longs pay 0.1% per 8-hour period to shorts. The payment is direct — no intermediary, no clearinghouse fee. The payment is enforced by the exchange's smart contract or matching engine.

The funding rate serves three functions. It anchors the perpetual price to spot. It indicates market sentiment — positive funding means the market is bullish (longs are paying to maintain their positions), negative funding means bearish. It generates yield for market-neutral strategies — a trader who buys spot and shorts the perpetual earns the funding rate without directional exposure. The yield is the basis trade. The basis trade is the subject of quantitative strategy research.

During bull markets, funding rates can reach extreme levels — 0.1% per 8 hours, which annualizes to over 100%. The extreme rates reflect extreme demand for leverage. The demand is self-limiting: high funding rates attract basis traders who short the perpetual and buy spot, earning the funding premium while pushing the perpetual price toward spot. The basis traders are the mechanism for market efficiency. The mechanism works. The speed at which it works depends on the capital available for basis trading.

## The reference

Robert Shiller, "Measuring Asset Values for Cash Settlement in Derivative Markets: Hedonic Repeated Measures Indices and Perpetual Futures" (1992). Shiller proposed the perpetual futures contract as a mechanism for creating derivative markets on illiquid assets — specifically, real estate indices. The contract would have no expiration and would use a "dividend" payment (what we now call the funding rate) to keep the contract price close to the index value. Shiller's proposal was academic. BitMEX's implementation in 2016 was commercial. The academic proposal became a $100 billion daily market. The market is larger than Shiller's wildest expectation. The mechanism is exactly as he described it.

---

**References:**
- Robert Shiller, "Measuring Asset Values for Cash Settlement in Derivative Markets," 1992.
- BitMEX, "Perpetual Contracts," BitMEX Documentation, 2016.
- Related posts: [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage), [Market Making](https://blog.hackspree.com/#dex-trading-market-making)

---
title: CEX-DEX Arbitrage
date: 2026-07-12
slug: dex-trading-cex-dex-arb
summary: "CEX-DEX arbitrage is the dominant strategy in crypto. A token trades at different prices on Binance and Uniswap. The arbitrageur buys on one, sells on the other. $233M extracted in 18 months by 19 searchers. Three of them captured 75%. This post explains the strategy, the infrastructure, and why the concentration is structural."
tags: dex, cex-dex, arbitrage, searchers, consolidation
series: dex-trading
part: 11
---

CEX-DEX arbitrage is the most empirically significant trading strategy in crypto. A token trades on a centralized exchange (Binance, Coinbase, Kraken) and a decentralized exchange (Uniswap, Curve, SushiSwap). The prices differ. The arbitrageur buys on the cheaper venue, sells on the dearer, pockets the spread. The trade is the mechanism that enforces the Law of One Price across the CEX-DEX boundary.

The strategy is not atomic. The CEX leg and the DEX leg execute on different systems, at different speeds, with different failure modes. The CEX leg requires an account, API keys, and compliance with the exchange's rules. The DEX leg requires a wallet, gas, and smart contract execution. The two legs cannot be bundled into a single transaction. The arbitrageur bears execution risk.

The AFT 2025 study of CEX-DEX arbitrage on Ethereum, covering August 2023 to March 2025, provides the definitive empirical picture:

| Metric | Value |
|---|---|
| Total value extracted | $233.8 million |
| Total arbitrage transactions | ~7.2 million |
| Number of major searchers | 19 |
| Top 3 searcher share | ~75% of volume and value |
| Daily transaction growth | 7.2× from Q3 2023 to Q1 2025 |
| Dominant searchers | Wintermute, SCP, Kayle |

The concentration is extreme. Three searchers capture three-quarters of the value. The concentration has been increasing over time. The increase is structural. The structure is the infrastructure cost.

## The infrastructure

CEX-DEX arbitrage requires infrastructure on both sides of the trade.

**CEX side.** Low-latency API access to Binance, Coinbase, Kraken. Colocation with the exchange's matching engine for minimum latency. Multiple accounts to avoid rate limits. Inventory of tokens on each exchange to enable immediate execution — you can't wait for a deposit to clear. The inventory is capital at risk. The capital must be managed across exchanges.

**DEX side.** Direct node access to the blockchain for mempool monitoring. Integration with block builders (Flashbots, MEV-Boost) for reliable transaction inclusion. Smart contracts optimized for gas efficiency. Flash loan integration for capital-free execution when the arbitrage is on-chain atomic. Gas price monitoring and dynamic fee adjustment.

**The bridge between them.** A system that monitors prices on both CEXs and DEXs in real time, identifies arbitrage opportunities, calculates profitability net of gas and fees, and executes the trade. The system must operate at the speed of the faster venue — if the CEX price updates before the DEX transaction confirms, the opportunity disappears. The window is milliseconds. The system that sees the window first, calculates profitability first, and executes first, wins. The system that wins consistently captures most of the value. The winner is the one with the best infrastructure.

## The consolidation

The AFT study documented a trend: searcher-builder vertical integration is deepening. Integrated searchers operate at lower margins, sometimes negative net profit, subsidized by revenue sharing with affiliated builders. The subsidy allows them to win arbitrages that independent searchers cannot profitably compete for. The independent searchers are driven out. The concentration increases.

The mechanism: a builder operates a block construction service. The builder's affiliated searcher submits arbitrage bundles to the builder. The builder includes the bundle and shares the MEV revenue with the searcher. The searcher can bid more aggressively for arbitrage opportunities because part of the revenue comes back through the builder relationship. The independent searcher has no such arrangement. The independent searcher must bid less to remain profitable. The independent searcher wins fewer arbitrages. The independent searcher's volume declines. The decline feeds back: lower volume means less data for model calibration, means worse execution, means even lower volume. The feedback loop drives consolidation.

This is the same dynamic that occurred in traditional equity markets, documented by Lewis in *Flash Boys*. The high-frequency trading firms that invested in infrastructure — colocation, private fiber, microwave networks — captured an increasing share of trading volume. The firms that didn't invest were driven out. The market consolidated around a few players. The players extracted rents. The rents were paid by investors in the form of wider effective spreads. The dynamic is repeating in crypto. The names are different. The structure is the same.

## The reference

"The CEX-DEX Arbitrage Landscape on Ethereum" (AFT 2025) is the definitive empirical study. It covers 18 months, 7.2 million transactions, $233.8 million in extracted value. It documents the consolidation trend, the searcher-builder vertical integration, and the implications for market efficiency and decentralization. It is the quantitative foundation for any discussion of arbitrage in crypto. The numbers are the story. The story is consolidation. The consolidation is accelerating.

---

**References:**
- AFT 2025, "The CEX-DEX Arbitrage Landscape on Ethereum: 2023–2025."
- Michael Lewis, *Flash Boys*, W.W. Norton, 2014.
- Related posts: [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage), [MEV](https://blog.hackspree.com/#dex-trading-mev), [Latency](https://blog.hackspree.com/#dex-trading-latency)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

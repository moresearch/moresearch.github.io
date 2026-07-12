---
title: Statistical Arbitrage
date: 2026-07-12
slug: dex-trading-stat-arb
summary: "Statistical arbitrage is not about a single price discrepancy. It is about a statistical edge across many trades. Pairs trading, cointegration, mean reversion. The math is from the 1980s. The execution is now on-chain. The principles haven't changed. The speed has."
tags: dex, stat-arb, pairs-trading, cointegration, mean-reversion
series: dex-trading
part: 7
---

Statistical arbitrage is the exploitation of statistical mispricings. Unlike pure arbitrage — buy cheap, sell dear, simultaneously — stat arb involves risk. The mispricing may persist. It may widen before it corrects. The stat arb trader bets that it will correct, on average, over many trades. The edge is small per trade. The volume makes it profitable.

The strategy emerged from the quantitative trading revolution of the 1980s. Gerry Bamberger and Nunzio Tartaglia at Morgan Stanley developed the first pairs trading strategies. The idea: identify two stocks that historically move together. When they diverge — one rises, the other falls — the spread between them has widened beyond its historical range. Buy the underperformer. Short the overperformer. Wait for convergence. The trade is market-neutral — you are not betting on the direction of the market, only on the relationship between the two stocks. Market-neutral strategies attracted capital because they were uncorrelated with the market. The lack of correlation was the selling point.

## The functional origin: pairs trading

Pairs trading is the simplest stat arb strategy. Two assets with a historical relationship. A spread between them that mean-reverts. The trader identifies the relationship, waits for divergence, enters the trade, exits on convergence. The relationship can be economic — two companies in the same industry, two tokens on the same blockchain — or statistical — two assets whose prices are cointegrated.

Cointegration, developed by Clive Granger (Nobel Prize, 2003) and Robert Engle (Nobel Prize, 2003), is the statistical property that makes pairs trading work. Two time series are cointegrated if each is non-stationary (its statistical properties change over time) but a linear combination of them is stationary (mean-reverts). The classic example: a drunk and her dog on a leash. Both walk randomly. The distance between them — the leash length — is stationary. The dog wanders. The drunk wanders. The leash pulls them back together. The leash is the cointegrating relationship.

In finance: two stocks in the same sector. Each follows a random walk. Their price ratio mean-reverts. The ratio is the leash. The trader buys when the ratio is low relative to its historical average — the underperformer is cheap relative to the outperformer. The trader sells when the ratio reverts. The trade is profitable if the cointegrating relationship persists. The relationship can break. The break is the risk.

In crypto: ETH and a liquid staking derivative like stETH. The two should trade at parity — 1 stETH = 1 ETH. They occasionally diverge — stETH trades at a discount during market stress, when holders want to exit staked positions quickly. The divergence is the opportunity. The trader buys stETH at a discount, waits for convergence, sells at parity. The trade is directional in the pair but market-neutral in dollar terms. The risk: the discount widens further, the trader's capital is locked, the convergence takes longer than the trader can afford. The risk is real. The risk is managed by position sizing and stop-losses.

## Statistical arbitrage in crypto

Crypto stat arb has advantages over traditional stat arb. The data is public and real-time. Every trade on a DEX is on-chain. Every price on a CEX is streamed. The data quality is higher than in traditional markets, where dark pools and off-exchange trading obscure true volumes. The crypto market structure — fragmented across hundreds of venues, each with its own liquidity profile — creates more opportunities for statistical mispricing than the consolidated equity markets. The fragmentation is the opportunity.

The strategies:

**Cross-DEX pairs.** The same token pair trades on multiple DEXs. The prices are usually close. When they diverge, the divergence is an opportunity. The trader buys on the cheaper DEX, sells on the dearer. The trade is atomic if the DEXs are on the same chain. It is non-atomic if they are on different chains — the trader bears bridge latency risk.

**Liquid staking derivative arbitrage.** stETH/ETH, rETH/ETH, cbETH/ETH. The derivatives should trade at or near parity with the underlying. They don't always. The discount widens during market stress. The trader accumulates at a discount, waits for convergence, redeems. The trade has a natural exit: the derivative can be redeemed for the underlying after the unstaking period. The redemption is the convergence guarantee. The guarantee makes the trade lower-risk than traditional pairs trading. The lower risk attracts capital. The capital compresses the spread.

**Mean reversion in AMM pools.** AMM pools exhibit mean-reverting behavior around the external market price. When a pool's price diverges from the CEX price, arbitrageurs trade against the pool to bring it back. The mean reversion is not guaranteed — it depends on arbitrageurs acting. But the arbitrageurs are reliable because the trade is profitable. The stat arb trader can front-run the arbitrageurs: enter when the divergence appears, exit when the arbitrageurs correct it. The trade is a bet on arbitrageurs doing their job. The bet is well-founded. The arbitrageurs are reliable.

## The reference

Andrew Lo, *Adaptive Markets: Financial Evolution at the Speed of Thought* (2017). Lo's book is not specifically about statistical arbitrage. It is about the broader framework that makes stat arb intelligible: markets are not perfectly efficient. They are adaptive systems populated by boundedly rational agents competing for profits. The competition produces efficiency in the long run. The transition produces profit opportunities for those who can identify and exploit them faster than others. Stat arb is the mechanism of the transition. The stat arb trader is the agent of efficiency. The efficiency, once achieved, eliminates the stat arb trader's edge. The trader must find new edges. The search is continuous. The search is the subject of this series.

---

**References:**
- Clive Granger, "Investigating Causal Relations by Econometric Models and Cross-Spectral Methods," *Econometrica*, 1969.
- Robert Engle and Clive Granger, "Co-Integration and Error Correction: Representation, Estimation, and Testing," *Econometrica*, 1987.
- Andrew Lo, *Adaptive Markets: Financial Evolution at the Speed of Thought*, Princeton University Press, 2017.
- Related posts: [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage), [Market Making](https://blog.hackspree.com/#dex-trading-market-making)

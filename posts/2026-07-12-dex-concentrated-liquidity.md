---
title: Concentrated Liquidity and JIT Attacks
date: 2026-07-12
slug: dex-trading-concentrated-liquidity
summary: "Uniswap V3 let LPs concentrate liquidity in a price range. Capital efficiency multiplied. Passive LPs were replaced by active managers. And a new form of MEV emerged: just-in-time liquidity, where an attacker provides liquidity for exactly one block, earns the fees, and exits. The LP game changed forever."
tags: dex, uniswap-v3, concentrated-liquidity, jit, mev
series: dex-trading
part: 10
---

Uniswap V2 required liquidity providers to deposit tokens across the entire price curve — from zero to infinity. Most of that liquidity was never used. In a stablecoin pair like USDC/DAI, the price rarely deviates from 1:1. Liquidity deposited at a price of 1:2 or 2:1 sat idle, earning no fees, providing no benefit. The capital was wasted. The waste was the cost of passive liquidity provision.

Uniswap V3, launched in May 2021, changed this. LPs choose a price range. Their liquidity is only active within that range. Within the range, the LP earns fees proportional to their share of the active liquidity. Outside the range, the LP earns nothing. The capital efficiency gain is dramatic: an LP who provides liquidity in a narrow range around the current price can achieve the same depth as a V2 LP with a fraction of the capital. Alternatively, an LP can provide far greater depth around the current price with the same capital — reducing slippage for traders and earning more fees.

The trade-off: the LP must actively manage the range. If the price moves outside the range, the LP's position becomes inactive. The LP must withdraw and redeposit in a new range. Each reposition costs gas. The LP who manages aggressively earns higher fees but pays higher gas. The LP who manages passively earns lower fees but pays lower gas. The optimal frequency depends on fee income, gas costs, and price volatility. The optimization is quantitative. The LP is now a market maker.

## The functional origin: tick sizes and decimalization

The concept of discrete price levels is not new. Traditional exchanges have tick sizes — the minimum price increment. A stock might trade in penny increments. The tick size constrains where orders can be placed. Decimalization — the switch from fractions (1/8, 1/16) to decimals (0.01) — reduced tick sizes in U.S. equity markets in 2001. The reduction narrowed spreads. The narrower spreads reduced market maker profits. The reduced profits drove consolidation among market makers.

Uniswap V3's tick-based liquidity is the crypto equivalent. The continuous price curve of V2 is discretized into ticks. LPs provide liquidity between ticks. The ticks are the price grid. The grid is finer than traditional tick sizes — 1 basis point (0.01%) per tick. The fineness enables precise range selection. The precision enables capital efficiency. The efficiency is the innovation.

## Just-in-time liquidity

Concentrated liquidity enabled a new form of MEV: just-in-time (JIT) liquidity. A searcher sees a large pending swap in the mempool. The searcher deposits concentrated liquidity at the exact tick range the swap will traverse, in the same block as the swap. The swap executes, traversing the searcher's liquidity. The searcher earns the fees — typically the majority of the swap's fee, since the searcher provided most of the active liquidity in that range. The searcher withdraws the liquidity in the same block. The searcher earns swap fees with near-zero inventory risk — the position existed for a single block.

The victim: passive LPs whose fee income is diluted. The swap would have earned them fees. Instead, the JIT LP captured those fees. The passive LP provided liquidity continuously, paid gas to deposit and withdraw, bore impermanent loss — and earned less because a JIT LP front-ran their fee income. The JIT LP extracted the most profitable slice of the fee stream: the large swaps that would have generated the highest fees per unit of liquidity. The passive LP got the residual — small swaps, unpredictable swaps, swaps too small for JIT to be profitable. The passive LP's returns declined. The decline is structural. The structure is the concentration of MEV extraction.

## The response

The JIT problem is an instance of a broader issue: concentrated liquidity enables capital efficiency, but capital efficiency enables MEV extraction at the expense of passive LPs. The solutions are architectural:

**Fee structures that penalize short-term liquidity.** If swap fees accrued linearly over time — the longer your liquidity is active, the larger your share of fees — JIT LPs would earn less per block. The accrual mechanism would favor long-term LPs. Uniswap V4, currently in development, is expected to introduce more flexible fee structures.

**Batch auctions that eliminate ordering.** As with sandwich attacks, batch auctions eliminate the ordering within a block that enables JIT. All swaps in a batch execute at a uniform clearing price. All LPs in the batch earn proportional fees. There is no "before" and "after" within the batch. There is no JIT.

**Passive LP vaults.** Protocols like Arrakis and Gamma manage concentrated liquidity positions on behalf of passive LPs. The vault rebalances ranges, compounds fees, and optimizes for fee income net of gas costs. The vault is a market maker operated by an algorithm. The algorithm competes with JIT searchers on equal footing — both are automated, both are fast, both can reposition in a single block. The vault democratizes active LP management. The democratization is partial. The infrastructure cost remains. The infrastructure is the barrier.

## The reference

Hayden Adams, Noah Zinsmeister, Dan Robinson, "Uniswap v3 Core" (2021). The whitepaper that introduced concentrated liquidity. It is 8 pages. The key insight — LPs choose a price range — is stated in the first paragraph. The rest is implementation. The implementation is now the dominant AMM design, copied by nearly every DEX launched since 2021. The copying is the evidence of the idea's power. The idea was simple. The implications — JIT, active management, LP stratification — are still unfolding.

---

**References:**
- Hayden Adams, Noah Zinsmeister, Dan Robinson, "Uniswap v3 Core," 2021.
- Arrakis Finance, "Concentrated Liquidity Vaults," Arrakis Documentation.
- Related posts: [AMMs](https://blog.hackspree.com/#dex-trading-amms), [Market Making](https://blog.hackspree.com/#dex-trading-market-making), [MEV](https://blog.hackspree.com/#dex-trading-mev)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

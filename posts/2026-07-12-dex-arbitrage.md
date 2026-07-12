---
title: Arbitrage
date: 2026-07-12
slug: dex-trading-arbitrage
summary: "Arbitrage is the oldest strategy in finance: buy cheap, sell dear, simultaneously. In crypto, flash loans made arbitrage available to anyone with code. The strategy is simple. The execution is not. The game is speed, and speed is infrastructure."
tags: dex, arbitrage, flash-loans, cex-dex, mev
series: dex-trading
part: 3
---

Arbitrage is the purchase and sale of the same asset in different markets to profit from price differences. It is the oldest trading strategy. It is also the purest: if two markets price the same thing differently, an arbitrageur can buy in the cheaper market and sell in the more expensive, pocketing the difference. The trade is riskless in theory — both legs execute at known prices, the profit is locked at execution. In practice, execution risk, latency, and competition make arbitrage a technological arms race.

The classical definition comes from Gustave de Molinari, a 19th-century French economist, but the practice is older than the definition. Medieval merchants arbitraged between markets separated by geography — buying spices in Venice, selling in Bruges. The internet collapsed geography. The price difference between New York and London narrowed to milliseconds. Crypto collapsed it further. The "markets" are now protocols on the same blockchain. The geography is the mempool. The distance is latency.

## The functional origin: the Law of One Price

The Law of One Price states that identical goods should trade at identical prices in efficient markets. If they don't, arbitrageurs will buy the cheaper and sell the dearer until the prices converge. The arbitrageur's profit is the market's mechanism for enforcing the law. The law is not a law of nature. It is an equilibrium condition. The equilibrium is maintained by arbitrageurs. The arbitrageurs are paid for maintaining it. The payment is the spread.

The Law of One Price was formulated by William Stanley Jevons in *The Theory of Political Economy* (1871). Jevons observed that "in the same open market, at any one moment, there cannot be two prices for the same kind of article." The observation was empirical. The mechanism was arbitrage. The mechanism is the same today. The markets are now AMM pools and CEX order books. The Law of One Price is enforced by bots that execute in milliseconds. The bots are paid in extracted spread. The spread is the market's payment for maintaining efficiency.

## Arbitrage in crypto

Crypto arbitrage takes several forms. The simplest is **pool arbitrage**: the same token pair trades at different prices in two AMM pools on the same chain. Buy in the cheaper pool, sell in the dearer pool, profit. The trade must be atomic — both legs in the same transaction — because the price will move after the first leg. A flash loan enables atomic execution. Borrow a large amount of token A. Execute the buy in pool 1. Execute the sell in pool 2. Repay the loan plus fee. Keep the profit. If any leg fails, the entire transaction reverts. The loan is never drawn. The risk is zero. The cost is gas + flash loan fee.

**CEX-DEX arbitrage**: a token trades at a different price on Binance than on Uniswap. The arbitrageur must execute on both venues simultaneously. This is not atomic — the CEX leg and the DEX leg are separate transactions on separate systems. The arbitrageur bears execution risk: one leg succeeds, the other fails. The risk is managed by speed — the faster you detect and execute, the less likely the price moves against you. The speed is the edge.

**Cross-chain arbitrage**: the same token trades on two different blockchains. Arbitrage requires bridging — sending the token from one chain to the other. The bridge takes time. During the bridge, the price may move. The arbitrageur bears bridge latency risk. The risk is managed by maintaining inventory on both chains and netting flows rather than bridging each trade.

**Triangular arbitrage**: three tokens. ETH → USDC → BTC → ETH. The product of the three exchange rates should equal one. If it doesn't, there is a triangular arbitrage. The strategy is common in forex. It is also common in crypto. It requires no external price feed. It is pure math on the chain.

## The empirical reality

The CEX-DEX arbitrage market on Ethereum, measured from August 2023 to March 2025, extracted $233.8 million. Three searchers captured 75% of the volume. The daily transaction count grew 7.2× over the period. The concentration is extreme. The concentration is structural. The infrastructure cost — colocation with sequencers, low-latency mempool access, private relay connections — creates barriers to entry. The barriers produce concentration. The concentration raises the question: is the Law of One Price being enforced by a competitive market or by a cartel of the fastest?

The AFT 2025 study of CEX-DEX arbitrage documented a shift: searcher-builder vertical integration is deepening. Integrated searchers operate at lower margins, sometimes negative net profit, subsidized by builder revenue sharing. The subsidy drives out independent searchers. The concentration increases. The market becomes less competitive. The Law of One Price is still enforced. The enforcers are fewer. The fewer enforcers extract more rent. The rent is paid by traders in the form of wider effective spreads. The market is efficient. The efficiency is expensive.

## The reference

Robert Shiller's *Irrational Exuberance* (2000) is not about arbitrage. It is about why arbitrage fails to enforce the Law of One Price at scale. Shiller documented persistent mispricings — the dot-com bubble, the housing bubble — that arbitrageurs could not or did not correct. The limits of arbitrage are the subject of Shleifer and Vishny's "The Limits of Arbitrage" (1997): arbitrage is risky, capital is constrained, and arbitrageurs who trade against bubbles can be wiped out before the bubble corrects. The limits are real in traditional markets. In crypto, the limits are different — not capital constraints but latency constraints, gas constraints, and infrastructure constraints. The constraints shape the market structure. The structure is the subject of the posts that follow.

---

**References:**
- William Stanley Jevons, *The Theory of Political Economy*, 1871.
- Andrei Shleifer and Robert Vishny, "The Limits of Arbitrage," *Journal of Finance*, 1997.
- AFT 2025, "CEX-DEX Arbitrage on Ethereum: 2023–2025."
- Related posts: [The Order Book](https://blog.hackspree.com/#dex-trading-order-book), [Automated Market Makers](https://blog.hackspree.com/#dex-trading-amms)

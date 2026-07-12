---
title: Algorithmic trading in crypto
date: 2026-07-12
slug: algorithmic-trading-crypto
summary: "Crypto markets are the most algorithmically traded markets in history. CEX-DEX arbitrage extracted $233M in 18 months. MEV is an infrastructure arms race. AI agents now achieve Sharpe ratios above 2.0. Latency is the ultimate edge. The game is being played at machine speed. Understanding it is the first step to not being someone else's exit liquidity."
tags: crypto, algorithmic-trading, mev, arbitrage, defi, latency
---

Crypto markets are the most algorithmically traded markets in history. The data is public. The execution is programmable. The settlement is near-instant. The barriers to entry are two orders of magnitude lower than traditional finance. The result is an arms race where latency, strategy, and infrastructure determine who extracts value and who provides it.

In the 18 months from August 2023 to March 2025, 19 major searchers extracted $233 million from CEX-DEX arbitrage on Ethereum alone. Three searchers captured 75% of the volume. The daily transaction count grew 7.2×. This is not a niche activity. This is the dominant mode of professional trading in crypto.

## The strategies

**CEX-DEX arbitrage.** The bread and butter. A token trades at a different price on Binance than on Uniswap. The arbitrageur buys on the cheaper venue, sells on the more expensive, pockets the spread. The trade must execute atomically — both legs must succeed or neither — or the arbitrageur is exposed to inventory risk. Flash loans solve this: borrow millions in a single transaction, execute both legs, repay the loan, keep the profit. If either leg fails, the entire transaction reverts. The loan is never drawn. The risk is zero. The barrier is speed. The fastest bot wins.

**Statistical arbitrage.** Not a single price discrepancy but a statistical edge across many assets over many trades. Pairs trading: two tokens that historically move together diverge. Buy the underperformer. Short the overperformer. Wait for convergence. Mean reversion: a token spikes on news, the spike is overdone, it reverts. The model identifies the deviation. The bot executes. The edge is small per trade. The volume makes it profitable.

**Market making.** Provide liquidity on both sides of the order book. Earn the spread. Manage inventory — too much inventory, you're exposed to price moves. Too little, you earn no spread. The market maker's edge is the bid-ask spread. The risk is adverse selection — informed traders trade against you when the price is about to move. The market maker who can't distinguish informed from uninformed flow loses. The market maker who can, wins. This is the oldest trade in finance. Crypto makes it programmable.

**MEV extraction.** Maximal Extractable Value. The profit that can be extracted by ordering, including, or excluding transactions within a block. Front-running: see a large buy order in the mempool, buy the same token first, sell after the large order moves the price. Sandwich attacks: buy before the victim, let the victim move the price, sell after. The victim pays a higher price. The attacker pockets the difference. MEV extraction is zero-sum. The extractor gains what the user loses. The extraction is algorithmic. The victim is anyone whose transaction waits in the public mempool.

**Liquidity sniping.** A new token launches. A liquidity pool is created. The sniper's bot detects the pool creation, buys tokens in the same block, and sells after the initial buying pressure drives the price up. The sniper's edge is speed — being first to the pool. The victim is the retail trader who buys after the snipe. The sniper's exit is the retail trader's entry.

**JIT (just-in-time) liquidity.** Provide liquidity for exactly one block — the block containing a large swap. The JIT LP sees the pending swap in the mempool, deposits liquidity, collects the swap fees, and withdraws in the same block. The LP earns fees with zero inventory risk. The strategy requires atomically bundling deposit, swap, and withdraw in one transaction. Flashbots and similar MEV infrastructure enable it.

## Latency is the edge

On fast-finality chains — Ethereum L2s like Arbitrum, Base, ZKsync — block times are sub-second. Priority fees are largely ignored because block builders can't reliably order by fee within that window. The winning strategy is not the highest fee. It is the lowest latency. The bot whose transaction reaches the sequencer first wins the arbitrage. The difference between winning and losing is measured in milliseconds.

A June 2025 study of L2 MEV found that over 80% of reverted transactions on L2s are swap transactions from MEV bots. Bots spam duplicate transactions — rather than paying higher fees, they flood the network with copies, betting that at least one will land first. The spam is economically rational because latency, not fees, determines ordering. The spam degrades the network for everyone else. The degradation is an externality. The externality is unpriced.

Geography matters. The bot running in the same datacenter as the sequencer has a latency advantage measured in single-digit milliseconds. The bot running on a home connection has no chance. Colocation is the moat. The moat is expensive. The expense concentrates the game among professionals. The concentration is visible in the data: three searchers, 75% of the volume.

## AI agents are entering

Multi-agent AI systems are now trading crypto with institutional-grade performance. Nex-T1, a 25-agent LLM system powered by GPT-4 Turbo with RAG, achieved a Sharpe ratio of 2.34 over its test period — outperforming single-agent baselines by 65% while cutting maximum drawdown by nearly half. The agents are organized into specialized teams: Research agents scan on-chain data, news, and social sentiment. Risk Management agents evaluate position sizing and portfolio exposure. Execution agents route trades across venues with sub-second latency. Governance agents monitor compliance and override anomalous decisions.

The architecture is the story. Each agent has a narrow responsibility. They communicate through structured interfaces. The system is modular — swap the execution agent without changing the research agent. The parallels to microservices architecture are exact. The difference is that these services trade money. The latency budget is milliseconds. The cost of a bug is not a 500 error. It is a position that moves against you while your agent is still thinking.

The agents use the same primitives as human traders — order books, AMM pools, lending protocols, bridges — but they operate at speeds humans cannot match. The human trader researches for hours. The agent ingests the entire on-chain state in seconds. The human trader monitors a few assets. The agent monitors thousands. The human trader sleeps. The agent doesn't. The asymmetry is structural. The structure favors the agent.

## The infrastructure stack

Algorithmic trading in crypto requires infrastructure. The stack:

| Layer | Component |
|---|---|
| **Data** | Real-time price feeds from CEXs (Binance, Coinbase, Kraken) and DEXs (Uniswap, Curve, pools on every chain). On-chain data — mempool monitoring, event logs, state diffs. Off-chain data — news, sentiment, order book depth. |
| **Execution** | Direct node access for low-latency transaction submission. MEV infrastructure — Flashbots, MEV-Boost, private relays — to avoid front-running and extract MEV. Smart contract wallets for programmatic execution. Flash loan contracts for atomic arbitrage. |
| **Strategy** | The model. Statistical, rules-based, or ML-driven. Backtested on historical data. Validated out-of-sample. Paper-traded before going live. The model is the edge. The edge erodes as competitors replicate it. The erosion is the reason for continuous research. |
| **Risk management** | Position limits. Exposure limits. Drawdown limits. Circuit breakers that halt trading when losses exceed thresholds. The risk system must be independent of the strategy system — a bug in the strategy should not prevent the risk system from closing positions. |
| **Monitoring** | Real-time P&L dashboards. Alerting on anomalous behavior — unexpected position sizes, unusual trade frequencies, sudden drawdowns. The monitoring system is the operator's window into a system that moves faster than any human can follow. |

## The game

Algorithmic trading in crypto is a multiplayer game with asymmetric information, asymmetric infrastructure, and asymmetric speed. The players with the best data, the lowest latency, and the most sophisticated models extract value from the players with worse data, higher latency, and simpler models. The extraction is the game. The game is zero-sum in the short run. In the long run, the infrastructure improves for everyone, the edges compress, and the extraction migrates to new venues, new assets, new strategies. The migration is continuous. The game never ends.

The individual trader entering this game with a Python script and a Binance API key is not a player. They are the liquidity. The extraction targets them. The latency gap ensures they will lose. The only rational response for the non-professional is passive: buy and hold, provide liquidity through automated vaults, or delegate to professional market makers who have the infrastructure. The active game is for professionals. The professional game is played at machine speed. The machines are getting faster.

---

**References:**
- AFT 2025, "CEX-DEX Arbitrage on Ethereum: 2023-2025."
- "First-Spammed, First-Served: MEV Extraction on Fast-Finality Blockchains," June 2025.
- Nexis-AI, "Nex-T1: Multi-Agent Framework for Autonomous DeFi Trading," October 2025.
- "RediSwap: MEV Redistribution at the Application Layer," 2024.
- Related posts: [Scarcity Rules Everything](https://blog.hackspree.com/#scarcity), [Design the Game](https://blog.hackspree.com/#scarcity-and-mechanism-design)


Engineering is the discipline of building things that work within constraints. Every topic on this blog — operating systems, AI models, trading infrastructure, research labs, innovation economics — is examined through the lens of systems design. The lens is engineering. The method is: understand the constraints, design within them, verify the design works, iterate. The domain provides the specifics. The method is universal.

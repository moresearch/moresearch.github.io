---
title: Sandwich Attacks
date: 2026-07-12
slug: dex-trading-sandwich
summary: "A sandwich attack is the most profitable form of MEV. The attacker buys before the victim and sells after, pocketing the price difference. The victim pays more. The attacker extracts the difference. The sandwich is the dark forest made visible. Understanding it is the only defense against it."
tags: dex, sandwich-attacks, mev, front-running, amm
series: dex-trading
part: 9
---

A sandwich attack is a three-transaction sequence. The attacker sees a pending swap in the mempool — the victim's transaction. The attacker submits a buy transaction for the same token with a higher gas price, getting included first. The attacker's buy moves the AMM price up. The victim's swap executes at the higher price — the victim pays more. The attacker submits a sell transaction immediately after, selling the tokens bought in the first transaction at the elevated price. The attacker profits. The victim loses. The difference is the sandwich.

The name is apt. The victim's transaction is the filling. The attacker's two transactions are the bread. The bread surrounds the filling. The filling is consumed. The consumer is the attacker. The consumed is the victim.

## The mechanics

In a constant-product AMM (xy = k), a swap changes the ratio of tokens in the pool, which changes the price. The larger the swap relative to the pool's liquidity, the larger the price impact. The sandwich attacker exploits the price impact of the victim's trade. The attacker buys before the victim, moving the price up. The victim buys at the inflated price, moving the price further up. The attacker sells at the post-victim price, profiting from the difference between their entry price and exit price.

The attacker's profit is bounded by the victim's slippage tolerance. The victim sets a maximum acceptable price — the slippage limit — when submitting the swap. If the price moves beyond the limit, the transaction reverts. The attacker must ensure the sandwich keeps the price within the victim's slippage tolerance. If it doesn't, the victim's transaction reverts, and the attacker is left holding tokens bought at an elevated price with no victim to sell to. The attacker loses. The slippage limit is the victim's defense. Setting it low reduces the sandwichable spread. Setting it too low causes the swap to revert on normal price movement.

The attacker also bears gas costs for three transactions. On Ethereum L1, gas can be expensive. The sandwich is only profitable if the extracted value exceeds 3× gas cost. On L2s, gas is cheaper. The lower cost makes smaller sandwiches profitable. The lower cost increases sandwich frequency. The increased frequency is documented in the data: over 80% of reverted transactions on L2s are MEV bots, many of them failed sandwich attempts.

## The functional origin: front-running in traditional markets

Front-running is as old as markets. A broker receives a large client order. The broker knows the order will move the price. The broker buys for their own account before executing the client order, then sells after the price moves. The broker profits at the client's expense. The practice is illegal in most regulated markets. The illegality is enforced by surveillance and prosecution.

In crypto, front-running is not illegal. It is profitable. The enforcement mechanism is not law. It is code. The victim's only defense is slippage limits and private transaction submission. The private submission — through Flashbots or similar relays — hides the transaction from the public mempool. The sandwich attacker cannot see it. The sandwich attack requires visibility. Eliminate the visibility. Eliminate the attack.

## The defense

**Slippage limits.** The most basic defense. Set the maximum acceptable price movement low. The lower the slippage tolerance, the less profit available to the sandwich attacker. The trade-off: too low, and the transaction reverts on normal volatility. The optimal slippage depends on the pool's liquidity, the trade size, and the current volatility. Wallets like MetaMask now suggest optimal slippage based on recent pool behavior. The suggestion is algorithmic. The algorithm is a defense.

**Private transaction submission.** Flashbots Protect and similar services route transactions directly to block builders, bypassing the public mempool. The sandwich attacker cannot see the transaction. The attack requires visibility. Eliminate visibility. Eliminate attack. The cost: the transaction may take slightly longer to include — private relays have different inclusion guarantees than the public mempool. The trade-off is speed vs. safety. For non-urgent transactions, private submission is the correct choice.

**Batch auctions.** Protocols like CoWSwap and 1inch Fusion aggregate orders and execute them in batches at a uniform clearing price. All orders in the batch receive the same price. There is no ordering within the batch. There is no sandwich. The batch auction is the architectural solution to sandwich attacks. The architecture eliminates the attack vector. The vector is eliminated by design, not by defense.

## The reference

Phil Daian et al., "Flash Boys 2.0" (2019). The paper documented sandwich attacks on Ethereum DEXs and measured their prevalence and profitability. It named the phenomenon. The name became the field. The field is MEV. The sandwich is the most visible form of MEV. The visibility is the subject of this post. The invisibility is the subject of the next.

---

**References:**
- Phil Daian et al., "Flash Boys 2.0," 2019.
- Flashbots, "MEV-Boost and Sandwich Attacks," Flashbots Documentation.
- CoWSwap, "Batch Auctions," CoWSwap Documentation.
- Related posts: [MEV](https://blog.hackspree.com/#dex-trading-mev), [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage)

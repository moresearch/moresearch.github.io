---
title: Liquid Staking
date: 2026-07-12
slug: fi-liquid-staking
summary: "Liquid staking solves the capital efficiency problem of proof-of-stake: you stake ETH to secure the network, and you get a liquid token (stETH) that you can use in DeFi while your ETH is locked. The functional origin is the depository receipt — a claim on a deposited asset that trades freely. Lido now holds over 30% of all staked ETH."
tags: defi, liquid-staking, lido, steth, proof-of-stake
series: crypto-fi-instruments
part: 6
---

Proof-of-stake requires validators to lock capital. On Ethereum, a validator must stake 32 ETH. The staked ETH is locked — it cannot be transferred, traded, or used as collateral. The lockup secures the network: validators who misbehave lose their stake. The lockup also creates an opportunity cost. Staked ETH cannot earn yield in DeFi. The opportunity cost is the return the staker could earn by deploying the ETH elsewhere.

Liquid staking eliminates the opportunity cost. A liquid staking protocol accepts ETH deposits, stakes them with validators, and issues a liquid token — stETH for Lido, rETH for Rocket Pool, cbETH for Coinbase — representing the deposited ETH plus accrued staking rewards. The liquid token can be traded, lent, or used as collateral in DeFi. The staker earns staking rewards plus whatever DeFi yield they generate with the liquid token. The capital does double duty. The efficiency is the innovation.

## The functional origin: depository receipts

A depository receipt is a financial instrument that represents ownership of an underlying asset held in custody. The most famous example is the American Depositary Receipt (ADR), introduced by J.P. Morgan in 1927. A U.S. bank holds shares of a foreign company in custody. It issues ADRs that trade on U.S. exchanges. The ADR holder receives dividends, votes, and price exposure without directly owning the foreign shares. The ADR solves the problem of cross-border investment: the foreign shares never leave their home market, but the ADR trades freely in the U.S.

The first depository receipt predates the ADR by centuries. The Dutch East India Company, founded in 1602, issued negotiable share certificates that could be transferred without altering the company's share register. The certificates were depository receipts in function if not in name. The principle — a claim on a deposited asset that trades independently — is the same.

Liquid staking tokens are depository receipts for staked ETH. The ETH is deposited with validators. The stETH is issued to the depositor. The stETH accrues staking rewards through a rebasing mechanism — the balance of stETH in the holder's wallet increases daily to reflect staking rewards. The stETH trades on AMMs, is accepted as collateral on lending protocols, and can be deployed in yield farming strategies. The deposited ETH secures the network. The stETH circulates in DeFi. The separation of security provision from capital deployment is the innovation.

## The risk

Liquid staking concentrates stake. Lido holds over 30% of all staked ETH. The concentration is a systemic risk to Ethereum: if Lido's validators collude or are compromised, the network's security is threatened. Lido distributes stake across multiple independent node operators to mitigate this risk. The distribution is a governance mechanism. The mechanism is imperfect. The imperfection is the subject of ongoing protocol development — distributed validator technology (DVT), stake capping, and validator set rotation.

The secondary risk: stETH can depeg from ETH. During market stress, stETH holders may want to exit to ETH directly rather than wait for the unstaking period. The selling pressure pushes stETH below its 1:1 peg. The discount creates an arbitrage opportunity — buy stETH at a discount, redeem for ETH after the unstaking period, profit. The arbitrage requires capital and patience. The capital must be locked during the unstaking period. The patience is tested during extended drawdowns. The discount is the market's price for immediate liquidity. The price fluctuates with market stress.

## The reference

J.P. Morgan, "American Depositary Receipts," 1927. Morgan's innovation was financial infrastructure: the ADR created a mechanism for U.S. investors to hold foreign equities without navigating foreign custody, settlement, and currency conversion. The ADR was a bridge between national financial systems. Liquid staking tokens are a bridge between the staking layer and the DeFi layer. The bridge is the infrastructure. The infrastructure enables capital to flow between layers. The flow is the efficiency. The efficiency is the value.

---

**References:**
- Lido, "Lido: Ethereum Liquid Staking," Lido Documentation.
- Rocket Pool, "Rocket Pool: Decentralised Ethereum Liquid Staking," Rocket Pool Documentation.
- Ethereum Foundation, "Proof of Stake," Ethereum Documentation.
- Related posts: [Stablecoins](https://blog.hackspree.com/#fi-stablecoins), [AMMs](https://blog.hackspree.com/#dex-trading-amms)

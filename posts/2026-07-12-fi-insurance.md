---
title: Insurance Protocols
date: 2026-07-12
slug: fi-insurance
summary: "Insurance began at Lloyd's coffee house in 1688, where merchants pooled risk to protect their ships. DeFi insurance protocols — Nexus Mutual, Unslashed, InsurAce — automate risk pooling on-chain. Members contribute capital to a mutual. Claims are assessed by members. The model is the same as Lloyd's. The execution is code."
tags: defi, insurance, nexus-mutual, lloyds, risk
series: crypto-fi-instruments
part: 10
---

Insurance is the oldest risk management tool. A group of people faces a common risk — shipwreck, fire, death. Each contributes to a pool. When one suffers a loss, the pool compensates them. The contributions are premiums. The payouts are claims. The pool is the mutual. The mutual is the insurance company.

The first modern insurance market was Lloyd's of London, which began at Edward Lloyd's coffee house on Tower Street in 1688. Merchants, shipowners, and underwriters gathered at Lloyd's to share news of shipping, negotiate insurance contracts, and pool risk. A shipowner seeking insurance would pass a slip around the coffee house. Underwriters would sign their names under the risk they were willing to accept, each taking a fraction. The slip was the contract. The signature was the commitment. The coffee house was the exchange.

The Lloyd's model persists today. Lloyd's is not an insurance company. It is a market where syndicates of underwriters compete to accept risk. The syndicates are backed by capital providers — "Names" — who put up their personal wealth as collateral. The Names earn premiums in good years and lose their fortunes in bad ones. The unlimited liability of Names was the mechanism that aligned incentives: the underwriter had skin in the game.

## DeFi insurance

DeFi insurance protocols apply the mutual model to crypto risks. The risks: smart contract bugs, oracle manipulation, stablecoin depegs, exchange hacks. The protocols: Nexus Mutual, InsurAce, Unslashed. The mechanism: members deposit capital into a mutual. Members purchase coverage against specific risks. When a claim is filed, members vote on whether to pay it. The vote is the claims assessment. The assessment is decentralized.

Nexus Mutual, launched in 2019, is the largest DeFi insurance protocol. It has paid claims for the UST depeg, the FTX collapse, and multiple protocol hacks. The claims process is governed by NXM token holders, who stake their tokens on the outcome of claims assessments. If a claim is valid and the assessors correctly vote to pay, they earn rewards. If a claim is fraudulent and assessors vote to deny, they also earn rewards. If assessors vote incorrectly — paying a fraudulent claim or denying a valid one — they lose their stake. The staking mechanism aligns incentives. The alignment is the mechanism design.

The challenge: claims assessment requires expertise. A smart contract bug may have been exploited by an attacker. Was the loss caused by a bug in the covered protocol or by user error? The distinction requires technical analysis. The analysis is provided by claims assessors. The assessors are compensated for their work. The compensation attracts expertise. The expertise is the product.

## The risk

Mutual insurance is capital-constrained. The capital in the mutual must cover all potential claims. If claims exceed capital, the mutual is insolvent. The insolvency risk is managed by risk-based pricing: higher-risk protocols pay higher premiums. The premium is the price of risk. The price is set by the market — by the willingness of capital providers to accept the risk at a given premium. The market for risk is the innovation. The innovation is the same as Lloyd's in 1688. The venue is different. The principle is the same.

## The reference

Lloyd's of London, *A History of Lloyd's* (various editions). The history of Lloyd's is the history of insurance. The coffee house. The slip. The Name. The syndicate. The innovation of Lloyd's was the market for risk — bringing together those who had risk and those who would accept it for a price. DeFi insurance protocols are recreating that market on-chain. The market is the same. The technology is different. The principles — pooling, diversification, skin in the game — are unchanged.

---

**References:**
- Lloyd's of London, historical archives.
- Nexus Mutual, "Nexus Mutual Documentation."
- Hugh Eaves, "The History of Insurance," Lloyd's Library.
- Related posts: [Design the Game](https://blog.hackspree.com/#scarcity-and-mechanism-design), [On Scarcity](https://blog.hackspree.com/#scarcity)

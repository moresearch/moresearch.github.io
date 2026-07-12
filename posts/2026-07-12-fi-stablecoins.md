---
title: Stablecoins
date: 2026-07-12
slug: fi-stablecoins
summary: "A stablecoin is a token that holds its value relative to a reference asset. The idea is older than crypto: David Ricardo proposed a gold-backed currency in 1816. The implementation is newer: fiat-backed (USDT), crypto-overcollateralized (DAI), and algorithmic (UST, which failed). The story of stablecoins is the story of money itself, repeatedly reinvented."
tags: defi, stablecoins, dai, usdt, monetary-history
series: crypto-fi-instruments
part: 1
---

The desire for stable value is as old as money. Every monetary innovation — coinage, paper currency, the gold standard, fiat money — was an attempt to create a medium of exchange whose value was predictable. Stablecoins are the latest iteration. They are tokens designed to maintain a peg to a reference asset, typically the U.S. dollar. They are the bridge between the volatile world of crypto assets and the stable world of everyday commerce.

The three types of stablecoins — fiat-backed, crypto-overcollateralized, and algorithmic — each have a different mechanism for maintaining the peg. Each mechanism has a different failure mode. Each failure mode has been demonstrated in production. The demonstrations were expensive. The lessons are public.

## The functional origin: David Ricardo and the gold standard

In 1816, David Ricardo published *Proposals for an Economical and Secure Currency*. Ricardo, the greatest economist of his generation, argued that the Bank of England should stop issuing gold coins and instead issue paper notes fully backed by gold bullion. The public would hold notes, not coins. The notes would be redeemable for gold at a fixed rate. The gold would sit in the Bank's vaults. The system would be more efficient — paper is cheaper to produce and easier to transport than gold — while maintaining the stability of the gold standard.

Ricardo's proposal was not adopted in his lifetime. It was adopted after his death, in the Bank Charter Act of 1844. The act centralized note issuance in the Bank of England and required new notes to be fully backed by gold. The gold standard, in this form, persisted until 1931. The principle — a paper claim on a reserve asset, redeemable at a fixed rate — is the principle of the fiat-backed stablecoin. Tether (USDT) is a Ricardo note. Circle (USDC) is a Ricardo note. The notes are issued on a blockchain instead of on paper. The reserve is held in a bank account instead of a vault. The principle is the same.

The failure mode is also the same. If the reserve is not fully backing the notes — if the issuer issues more notes than it holds reserves — the peg is vulnerable to a run. The Bank of England suspended convertibility in 1797, during the Napoleonic Wars, and again in 1914, during World War I. Tether has never suspended convertibility, but it has never been fully audited. The opacity is the vulnerability. The vulnerability is the subject of ongoing regulatory attention.

## Crypto-overcollateralized: DAI

DAI, launched by MakerDAO in 2017, is a stablecoin backed by crypto assets rather than fiat reserves. The mechanism: users deposit crypto collateral (ETH, WBTC, other approved assets) into a vault. They mint DAI against the collateral. The collateralization ratio must exceed a minimum — typically 150%. If the value of the collateral falls below the liquidation threshold, the vault is liquidated: the collateral is sold, the DAI is repaid, and the remaining collateral is returned to the user. The overcollateralization absorbs price volatility. The liquidation mechanism enforces the peg.

DAI is the first decentralized stablecoin that achieved scale. It survived the March 2020 crash — when ETH fell 50% in a day, Maker's liquidation system failed to keep up, and DAI briefly traded above its peg — and emerged with improved mechanisms. It survived the Luna collapse in May 2022 — when UST imploded and algorithmic stablecoins were discredited — and DAI's overcollateralized model was vindicated by contrast. DAI is now the dominant decentralized stablecoin, with a market cap exceeding $5 billion.

The innovation: a stablecoin backed by volatile assets, stabilized by overcollateralization and automated liquidation. The risk: a black swan event that crashes collateral faster than the liquidation system can respond. The risk is managed by collateral diversification, liquidation parameter tuning, and an emergency shutdown mechanism. The risk is not eliminated. It is priced.

## Algorithmic: UST and the lesson

TerraUSD (UST), launched in 2020, was an algorithmic stablecoin. It had no reserves. It maintained its peg through a seigniorage mechanism: UST could be redeemed for $1 worth of LUNA, Terra's governance token, at any time. If UST traded below $1, arbitrageurs would buy UST and redeem it for LUNA, profiting from the difference and pushing UST back to $1. If UST traded above $1, arbitrageurs would mint UST with LUNA and sell UST, pushing it back down. The mechanism relied on LUNA having value. In May 2022, confidence in LUNA collapsed. UST lost its peg. The death spiral: UST below peg → arbitrageurs redeem for LUNA → LUNA supply expands → LUNA price falls → confidence falls further → more redemptions. Within days, $40 billion in market value was destroyed.

The UST collapse was the most significant event in stablecoin history. It demonstrated that algorithmic stablecoins without exogenous collateral are vulnerable to death spirals. The vulnerability is not a bug. It is a property of the mechanism. The mechanism works when confidence holds. Confidence is reflexive — it holds when people believe it will hold. When belief cracks, the mechanism accelerates the collapse. The acceleration is the death spiral. The death spiral is a feature of the design.

The lesson: stablecoins require backing. The backing can be fiat (USDT, USDC), crypto overcollateralized (DAI), or a basket of assets. It cannot be pure belief. Belief is not capital. Capital is the foundation. The foundation matters when the market turns.

## The reference

Friedrich Hayek, *The Denationalisation of Money* (1976). Hayek argued that the government monopoly on currency issuance should be abolished and replaced with competing private currencies. Each issuer would maintain its currency's value by promising redeemability and managing supply. The market would select the currencies that maintained stable purchasing power. Hayek's vision was dismissed as utopian. Stablecoins are the partial realization of that vision. Competing private currencies. Redeemability as the mechanism. Market selection as the arbiter. UST failed the market test. DAI passed it. The market is the selector. The selection is ongoing.

---

**References:**
- David Ricardo, *Proposals for an Economical and Secure Currency*, 1816.
- Friedrich Hayek, *The Denationalisation of Money*, Institute of Economic Affairs, 1976.
- MakerDAO, "The Maker Protocol: MakerDAO's Multi-Collateral Dai System," 2019.
- Related posts: [AMMs](https://blog.hackspree.com/#dex-trading-amms), [Scarcity Rules Everything](https://blog.hackspree.com/#scarcity)

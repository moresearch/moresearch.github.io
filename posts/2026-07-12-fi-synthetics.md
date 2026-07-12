---
title: Synthetic Assets
date: 2026-07-12
slug: fi-synthetics
summary: "A synthetic asset is a token that tracks the price of something without requiring custody of that thing. Synthetix lets you trade synthetic gold, synthetic Apple stock, synthetic Bitcoin — all on-chain, collateralized by crypto. The functional origin is the contract for difference (CFD), invented in London in the 1950s. The execution is on Ethereum."
tags: defi, synthetics, synthetix, cfds, derivatives
series: crypto-fi-instruments
part: 7
---

A synthetic asset is a financial instrument that simulates the payoff of another asset without requiring ownership of that asset. A synthetic S&P 500 token tracks the price of the S&P 500 index without the holder owning any of the underlying stocks. The synthetic is created by collateralization: a user deposits crypto as collateral and mints the synthetic token. The token's value is maintained by an oracle that reports the price of the underlying asset. If the synthetic token's price diverges from the underlying, arbitrageurs trade the divergence back to parity.

Synthetix, launched in 2018, is the dominant synthetic asset protocol. Users deposit SNX (the Synthetix governance token) as collateral and mint synthetic assets — synths — that track currencies, commodities, cryptocurrencies, and equities. The synths trade on Synthetix's own exchange. The exchange uses a pooled collateral model: all synths are backed by the collective SNX collateral, not by individual positions. The pooling enables infinite liquidity — any synth can be traded against any other synth at the oracle price, with zero slippage. The zero slippage is the key innovation. The key innovation is funded by SNX stakers, who bear the risk of the pooled collateral.

## The functional origin: contracts for difference

The contract for difference (CFD) was invented in London in the 1950s by hedge funds seeking to trade equities on margin without triggering stamp duty — a tax on share transactions. The CFD is an agreement between two parties to exchange the difference between the opening and closing price of an underlying asset. The buyer doesn't own the asset. The seller doesn't deliver it. They settle the price difference in cash.

CFDs became popular among retail traders in the 1990s, offered by spread-betting firms like IG Index. The appeal: leveraged exposure to any asset class without owning the underlying, without paying stamp duty, and without the complexity of futures or options. The risk: the leverage amplifies losses. The risk materialized in 2015 when the Swiss National Bank unpegged the franc from the euro. EUR/CHF collapsed. Retail CFD traders were wiped out. Several brokers went bankrupt.

Synthetic assets are CFDs on a blockchain. The collateral is on-chain. The settlement is automatic. The counterparty is the protocol, not a broker. The protocol's solvency depends on the collateralization ratio. If the ratio falls below the threshold, the protocol must liquidate positions or mint new tokens. The mechanism is the same as MakerDAO's DAI, generalized to any asset.

## The pooled collateral model

Synthetix's pooled collateral model is the key architectural difference from MakerDAO. In MakerDAO, each vault is independent — if one vault is undercollateralized, only that vault is liquidated. In Synthetix, all synths are backed by a single pool of SNX collateral. The pool absorbs gains and losses collectively. If the value of outstanding synths exceeds the value of the collateral pool, the protocol is insolvent. The insolvency risk is collective. The collective risk requires a higher collateralization ratio — Synthetix targets 400-500%, compared to MakerDAO's 150%.

The pooled model enables infinite liquidity. Because all synths are fungible against the collateral pool, any synth can be traded for any other synth at the oracle price. There is no order book. There is no AMM curve. There is no slippage. The trade is executed against the pool at the oracle price plus a fee. The fee goes to SNX stakers. The model is elegant. The elegance is the attraction. The attraction is counterbalanced by the complexity of managing a pooled collateral system with exposure to dozens of synthetic assets.

## The frontier

Synthetic assets enable on-chain exposure to off-chain assets. The oracle reports the price of Apple stock. The protocol mints sAAPL. The user trades sAAPL on-chain. The exposure is synthetic. The custody is unnecessary. The border between traditional finance and DeFi is the oracle. The oracle is the bridge. The bridge enables a future where any asset can be traded on-chain without the asset ever touching a blockchain. The future is synthetic. The synthesis is the product.

## The reference

Synthetix, "Synthetix Litepaper" (2018). The original description of the pooled collateral model and the infinite liquidity exchange. The litepaper is 12 pages. The key insight — pooled collateral enables zero-slippage trading between any synthetic assets — is on page 3. The rest is implementation. The implementation is now managing billions in synthetic asset value. The value is the proof of the concept.

---

**References:**
- Synthetix, "Synthetix Litepaper," 2018.
- IG Group, "Contracts for Difference," IG Documentation.
- Related posts: [Stablecoins](https://blog.hackspree.com/#fi-stablecoins), [Perpetual Futures](https://blog.hackspree.com/#fi-perpetuals)

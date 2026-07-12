---
title: Yield Farming
date: 2026-07-12
slug: fi-yield-farming
summary: "Yield farming is the practice of deploying capital across DeFi protocols to earn returns. The functional origin is sharecropping — providing land (capital) in exchange for a share of the harvest (yield). The innovation is liquidity mining — protocols paying users in governance tokens to bootstrap liquidity. The economics are incentive design. The risks are impermanent loss, smart contract bugs, and the inevitable decline of unsustainable yields."
tags: defi, yield-farming, liquidity-mining, incentives, compound
series: crypto-fi-instruments
part: 5
---

Yield farming is the practice of deploying capital across DeFi protocols to earn the highest available return. The farmer deposits tokens into lending pools, AMM liquidity pools, or options vaults. The protocols pay interest, trading fees, or incentive rewards. The farmer monitors returns and reallocates capital as yields change. The activity is called farming because the capital is the seed and the yield is the harvest.

Yield farming emerged in the summer of 2020 — "DeFi Summer" — when Compound introduced liquidity mining. Compound distributed COMP governance tokens to users who borrowed and lent on the protocol. The distribution turned lending from a low-yield activity into a high-yield one. Users borrowed assets they didn't need, lent them back, and earned COMP on both sides. The activity was economically circular — it didn't increase lending efficiency, it harvested tokens. The harvest was profitable while COMP traded at high valuations. The valuations were sustained by the expectation of future protocol revenue. The expectation was speculative. The speculation was the yield.

## The functional origin: sharecropping

Sharecropping is an agricultural arrangement where a landowner provides land to a farmer in exchange for a share of the crop. The arrangement emerged after the American Civil War, when former slaves had labor but no land, and former plantation owners had land but no labor. The sharecropper worked the land. The landowner provided the capital — land, tools, seed. The harvest was split. The split was the return on capital.

The arrangement was exploitative in practice — landowners often manipulated accounts to keep sharecroppers in debt — but the economic logic was sound: capital and labor combine to produce output. The capital provider earns a return proportional to the capital's contribution. The labor provider earns a return proportional to the labor's contribution. The split is the subject of bargaining.

Yield farming is sharecropping on blockchain rails. The farmer provides capital — tokens deposited into a protocol. The protocol provides the "land" — the smart contract infrastructure that generates fees. The harvest is the fees plus the incentive tokens. The split is determined by the protocol's parameters, not by a landowner's accounting. The transparency eliminates the exploitation vector. The automation eliminates the bargaining. The code is the landlord. The code is neutral.

## Liquidity mining

Liquidity mining is the distribution of protocol governance tokens to users who provide liquidity. The mechanism: the protocol allocates a percentage of its token supply to liquidity providers. The tokens are distributed pro-rata based on each provider's share of the pool. The tokens have value if the protocol has value. The value of the tokens is the incentive to provide liquidity.

Liquidity mining is an incentive design problem. The protocol wants to attract liquidity — deeper pools mean lower slippage, which attracts traders, which generates fees, which attracts more liquidity. The incentive tokens are the subsidy. The subsidy is paid to early liquidity providers to overcome the chicken-and-egg problem: no one wants to provide liquidity to an empty pool, and no one wants to trade in a pool with no liquidity. The subsidy breaks the equilibrium. The subsidy costs the protocol dilution of its token. The dilution is the cost of bootstrapping.

The problem: liquidity mining attracts mercenary capital. When yields are high, capital floods in. When yields decline — because the token price falls, or the subsidy ends, or a competing protocol offers higher yields — capital floods out. The mercenary capital is not loyal. The loyalty must be earned by the protocol's fundamentals: fee generation, user growth, sustainable economics. The protocols that achieve fundamentals survive the end of the subsidy. The protocols that don't die when the subsidies stop. The death is the market's verdict on the protocol's underlying value.

## The strategies

**Simple lending.** Deposit stablecoins into Aave or Compound. Earn the lending interest rate. The rate fluctuates with utilization. The yield is modest — typically 2-10% APY for stablecoins. The risk is low — the protocols are battle-tested, the collateral is overcollateralized. Simple lending is the baseline. The baseline is the risk-free rate of DeFi.

**LP farming.** Provide liquidity to an AMM pool. Earn trading fees plus incentive tokens. The yield is higher than lending. The risk is higher — impermanent loss, smart contract risk, incentive token price risk. LP farming is the most common yield farming strategy. The commonness is evidence of the risk-return trade-off. The trade-off is the farmer's decision.

**Leveraged farming (yield looping).** Deposit collateral. Borrow against it. Deposit the borrowed funds. Repeat. The leverage multiplies the yield and the risk. If the lending rate exceeds the farming yield, the position loses money. If the collateral value falls, the position is liquidated. Leveraged farming is the riskiest yield strategy. The riskiest strategy attracts the most sophisticated farmers. The sophistication is the barrier to entry. The barrier protects the yields of those who cross it.

**Auto-compounding.** Protocols like Yearn and Beefy automatically compound yields — harvesting rewards, selling them for the underlying asset, redepositing. The compounding increases APY. The automation saves gas. The protocol takes a fee. The fee is the price of convenience. The convenience is the product.

## The reference

Vitalik Buterin, "On Liquidity Mining" (2020). Buterin argued that liquidity mining is only sustainable if the distributed tokens confer genuine governance rights over a protocol that generates genuine revenue. Otherwise, liquidity mining is a wealth transfer from late buyers of the token to early farmers. The wealth transfer is zero-sum. The zero-sum game ends when the music stops. The music stopped for many protocols in 2022. The protocols that survived had fundamentals. The fundamentals were revenue. The revenue was from fees. The fees were from users. The users were real.

---

**References:**
- Vitalik Buterin, "On Liquidity Mining," 2020.
- Compound, "Compound Governance," 2020.
- Yearn Finance, "Yearn Vaults," Yearn Documentation.
- Related posts: [Lending Protocols](https://blog.hackspree.com/#fi-lending), [AMMs](https://blog.hackspree.com/#dex-trading-amms)

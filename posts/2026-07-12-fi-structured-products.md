---
title: Structured Products
date: 2026-07-12
slug: fi-structured-products
summary: "A structured product packages multiple financial instruments into a single product with a predefined payoff structure. Principal-protected notes, yield enhancement, range-bound strategies. The functional origin is the structured note — invented in the 1980s to give retail investors access to derivatives strategies. DeFi structured products automate these strategies on-chain."
tags: defi, structured-products, options, yield, ribeye
series: crypto-fi-instruments
part: 11
---

A structured product is a packaged investment strategy. It combines multiple financial instruments — typically a bond and one or more derivatives — to produce a predefined payoff profile. The investor deposits capital. The product returns the capital plus a return that depends on the performance of an underlying asset, subject to conditions. If the conditions are met, the investor earns an enhanced yield. If they are not, the investor may earn nothing, lose some principal, or receive a minimum guaranteed return.

The simplest example is a principal-protected note: the investor deposits $100. The product invests $95 in a zero-coupon bond that matures at $100 in one year. The remaining $5 buys a call option on the S&P 500. If the S&P rises, the investor participates in the upside. If it falls, the investor gets their $100 back at maturity. The principal is protected. The upside is capped by the cost of the option. The structure is a bond plus a call option. The structure is a structured product.

## The functional origin: structured notes

Structured notes were developed by investment banks in the 1980s and 1990s to provide retail investors with access to derivatives strategies without requiring them to understand or trade derivatives directly. The bank packaged the derivatives into a note — a debt instrument — and sold it through brokers. The investor bought the note. The bank managed the underlying strategy. The bank earned fees. The investor earned a return linked to the strategy's performance.

The market grew rapidly. By the 2000s, structured notes were a multi-trillion-dollar market. The products ranged from simple (principal-protected equity notes) to incomprehensible (tranched CDOs-squared). The complexity was the vulnerability. When the underlying assets — subprime mortgages — defaulted in 2008, the structured products collapsed. The collapse was the financial crisis. The crisis was the lesson: structured products are only as safe as their underlying assets and the transparency of their structure.

## DeFi structured products

DeFi structured products automate the same strategies on-chain with full transparency. The types:

**Covered call vaults.** Ribbon Finance and Thetanuts sell out-of-the-money call options against deposited assets. The premiums generate yield. If the underlying asset stays below the strike price, the options expire worthless and the depositor keeps the premium plus the asset. If the asset rises above the strike, the depositor's asset is called away — they receive the strike price instead of the asset. The upside is capped. The premium is the compensation for the capped upside. The vault automates the option selling and compounding. The depositor earns yield without managing strikes or expiries.

**Principal-protected products.** Cega and Friktion offer principal-protected notes on crypto assets. The investor deposits stablecoins. The vault invests most in lending protocols (the bond) and uses the remainder to buy options (the upside). If the options pay off, the investor earns enhanced yield. If they don't, the investor gets their principal back. The structure is transparent — all positions are on-chain. The transparency is the improvement over traditional structured notes, where the underlying positions were opaque.

**Range-bound strategies.** The investor bets that an asset will stay within a price range. If it does, they earn yield. If it breaks out, they may lose principal or have their asset converted. The strategy is implemented by selling a strangle — a call and a put at different strikes. The premium is the yield. The conversion is the risk.

**Basis trade vaults.** The vault executes the cash-and-carry trade — buying spot and shorting perpetual futures — to earn the funding rate. The trade is market-neutral. The yield is the funding rate. The vault automates position management, rebalancing, and compounding.

## The reference

Peter Bernstein, *Against the Gods: The Remarkable Story of Risk* (1996). Bernstein's book is a history of risk management from the Renaissance to modern finance. The central argument: the quantification of risk — probability theory, statistics, derivatives pricing — is the defining intellectual achievement of modern capitalism. Structured products are the application of that achievement to retail investment products. DeFi structured products are the next iteration. The math is the same. The execution is on-chain. The transparency is the improvement.

---

**References:**
- Peter Bernstein, *Against the Gods*, Wiley, 1996.
- Ribbon Finance, "Ribbon Finance Documentation."
- Cega, "Cega: Structured Products for DeFi," Cega Documentation.
- Related posts: [Options and DeFi Derivatives](https://blog.hackspree.com/#fi-options), [Perpetual Futures](https://blog.hackspree.com/#fi-perpetuals)

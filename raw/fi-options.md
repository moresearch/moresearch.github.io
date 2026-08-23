---
title: Options and DeFi Derivatives
date: 2026-07-12
slug: fi-options
summary: "Thales of Miletus traded the first recorded option in the 6th century BC — a bet on the olive harvest. The Black-Scholes formula made options mathematically tractable in 1973. DeFi options protocols — Hegic, Opyn, Ribbon — are automating option writing, selling, and settlement on-chain. The instrument is 2,600 years old. The execution is new."
tags: defi, options, black-scholes, derivatives, hegic
series: crypto-fi-instruments
part: 4
---

Aristotle, in the *Politics*, tells the story of Thales of Miletus. Thales was a philosopher. His critics mocked him for his poverty, arguing that philosophy was useless because it couldn't make money. Thales, using his knowledge of astronomy, predicted a particularly abundant olive harvest. During the winter, when demand for olive presses was low, he paid small deposits to reserve the use of all the olive presses in Miletus and Chios for the following autumn. When the harvest arrived — abundant, as he predicted — demand for olive presses surged. Thales sold his reservations at a premium. He made a fortune.

The transaction was a call option. Thales paid a premium for the right, but not the obligation, to use the olive presses at a future date. If the harvest had been poor, he would have let the option expire. His loss would have been limited to the premium. The structure — limited downside, asymmetric upside, premium paid upfront — is the structure of every option contract since. Thales invented the option. He also proved that philosophy could make money. The two achievements are connected.

## The functional origin: Black-Scholes

The modern options market was created by a formula. Fischer Black and Myron Scholes published "The Pricing of Options and Corporate Liabilities" in 1973. The Black-Scholes formula calculates the theoretical price of a European call option as a function of the underlying asset price, the strike price, the time to expiration, the risk-free interest rate, and the asset's volatility. The formula assumes continuous trading, no transaction costs, and log-normal price distributions. The assumptions are false. The formula is still useful. The usefulness derives from the formula's insight: an option can be replicated by dynamically hedging a position in the underlying asset. The replication argument is the foundation of all derivatives pricing.

The Chicago Board Options Exchange opened in April 1973, the same year Black-Scholes was published. The coincidence was not planned. The coincidence was catalytic. Traders had a formula for pricing options. The formula gave them confidence to trade. The trading created liquidity. The liquidity attracted more traders. The virtuous cycle created the modern options market, which now trades trillions in notional value annually.

## DeFi options

DeFi options protocols are automating the options market on-chain. The approaches:

**Order book options.** Hegic, Opyn, and Lyra use on-chain order books or request-for-quote systems for options trading. Buyers and sellers match on-chain. Settlement is automated. The challenge: options are complex instruments with many parameters (strike, expiry, type). The order book for any specific option is thin. The thinness produces wide spreads. The spreads limit adoption.

**Automated options vaults.** Ribbon Finance and ThetaNuts sell options automatically on behalf of depositors. A depositor deposits ETH into a covered call vault. The vault sells call options against the ETH, collecting premiums. The premiums are distributed to depositors. The vault automates the option selling strategy. The depositor earns yield without managing strikes, expirations, or greeks. The automation is the product. The product is yield.

**Structured products.** Cega and Friktion combine options into structured products — principal-protected notes, yield enhancement, volatility trading. The investor deposits capital. The protocol executes a strategy involving multiple options positions. The strategy generates yield in most market conditions and loses principal in tail events. The risk is the tail. The tail is priced into the premium. The pricing is the challenge.

**Perpetual options.** Paradigm and Panoptic are developing perpetual options — options with no expiry, analogous to perpetual futures. The mechanism is a funding rate between option buyers and sellers. The perpetual option eliminates expiration management. The perpetual option is the next frontier.

## The reference

Nassim Nicholas Taleb, *Dynamic Hedging: Managing Vanilla and Exotic Options* (1997). Taleb was a options market maker before he was an author. His book is the practitioner's guide to the reality of options trading — the greeks, the hedging, the tail risks that Black-Scholes assumes away. Taleb's later book, *The Black Swan* (2007), is about the consequences of those tail risks. The DeFi options market is still young. The tail risks have not yet materialized at scale. When they do, Taleb's framework will be the guide to understanding them. The framework is ready. The market is not.

## The engineering connection

An option is a function from an underlying price to a payoff. Black-Scholes is the algorithm that prices that function. DeFi options vaults automate the execution of that algorithm on-chain. The automation is software engineering applied to financial engineering. The vault deposits collateral, sells options, collects premiums, reinvests — the same loop as a CI/CD pipeline: trigger, execute, verify, repeat. The domain is different. The control flow is the same.

The DeFi options stack also illustrates a recurring engineering pattern: complexity compression. Black-Scholes is a partial differential equation. A covered call vault presents it as "deposit ETH, earn yield." The vault compresses the complexity of options pricing, Greeks management, and position monitoring into a single user action. The compression is the product. The product is an abstraction. The abstraction is good if it hides the right things and bad if it hides the wrong things. The vault that hides tail risk from the user is a bad abstraction. The vault that surfaces tail risk is a good one. The engineer's judgment is knowing which is which.

---

**References:**
- Aristotle, *Politics*, Book I, Chapter XI (Thales and the olive presses).
- Fischer Black and Myron Scholes, "The Pricing of Options and Corporate Liabilities," *Journal of Political Economy*, 1973.
- Nassim Nicholas Taleb, *Dynamic Hedging*, Wiley, 1997.
- Related posts: [Perpetual Futures](https://blog.hackspree.com/#fi-perpetuals), [Market Making](https://blog.hackspree.com/#dex-trading-market-making)

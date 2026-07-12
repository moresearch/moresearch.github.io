---
title: Flash Loans
date: 2026-07-12
slug: dex-trading-flash-loans
summary: "A flash loan lets you borrow millions with zero collateral, as long as you repay in the same transaction. It is the innovation that made atomic arbitrage possible at scale. It is also the weapon of choice for protocol exploits. The same mechanism that enforces the Law of One Price can drain a lending protocol in 12 seconds."
tags: dex, flash-loans, arbitrage, atomicity, defi
series: dex-trading
part: 5
---

A flash loan is an uncollateralized loan that must be borrowed and repaid within a single blockchain transaction. If the loan is not repaid by the end of the transaction, the entire transaction reverts. The loan is never actually disbursed. The borrower pays a fee. The lender earns the fee with zero credit risk — the atomicity of the transaction guarantees either full repayment or full reversion. There is no partial state. There is no default. The mechanism is pure atomicity.

Flash loans were introduced by Marble Protocol in 2018 and popularized by Aave and dYdX in 2019-2020. The innovation was not the concept — uncollateralized intraday credit has existed in traditional finance for centuries. The innovation was the enforcement mechanism: in traditional finance, uncollateralized credit requires trust, legal contracts, and recourse to courts. In DeFi, it requires a single line of Solidity: `require(repayment >= loan, "not repaid")`. The enforcement is code. The code is the court.

## The functional origin: intraday credit

In traditional finance, brokers extend intraday credit to clients. A hedge fund buys $10 million of stock in the morning and sells it in the afternoon. The broker lends the $10 million for a few hours. The credit is uncollateralized — the broker trusts the fund to settle by end of day. If the fund doesn't settle, the broker has legal recourse. The recourse is slow and expensive. The trust is the cost.

The continuous linked settlement (CLS) system in foreign exchange, launched in 2002, addressed a similar problem: FX trades settle in different currencies at different times, creating settlement risk — the risk that one party pays but the other doesn't. CLS uses payment-versus-payment (PvP) settlement: both legs settle simultaneously or neither settles. The simultaneity eliminates settlement risk. The mechanism is atomicity. Flash loans are CLS for DeFi — atomic settlement without a central clearinghouse. The blockchain is the clearinghouse. The atomicity is the mechanism.

## How flash loans enable arbitrage

Arbitrage requires capital. The price difference between two pools might be 0.5%. To extract meaningful profit, the arbitrageur needs to trade large amounts. Large amounts require capital. The capital has an opportunity cost — it could be deployed elsewhere. The opportunity cost reduces the net return on arbitrage. Flash loans eliminate the capital requirement. The arbitrageur borrows the capital, executes both legs, repays the loan, and keeps the profit. The capital is never at risk. The opportunity cost is zero. The arbitrageur's only cost is gas + flash loan fee. The fee is typically 0.09% on Aave. The spread must exceed the fee to be profitable. The spread that doesn't exceed the fee persists. The spread that does is extracted.

The flash loan democratized arbitrage — in theory. Anyone with a smart contract can borrow millions and execute an arbitrage. In practice, the democratization was limited by the same latency and infrastructure constraints that concentrate all MEV extraction. The flash loan solves the capital problem. It doesn't solve the speed problem. Speed still wins. Speed requires infrastructure. Infrastructure requires capital. The democratization is partial. The partial democratization is the state of the market.

## Flash loans as attack vectors

The same atomicity that enables arbitrage enables attacks. A flash loan can be used to manipulate the price of a governance token, borrow against the inflated collateral, drain the lending protocol, and repay the flash loan — all in one transaction. The attack requires no capital. The attacker pays only gas + flash loan fee. If the attack succeeds, the profit can be millions. The protocol is left with bad debt. The attacker is untraceable. The transaction was atomic. The exploitation was instantaneous.

The most famous flash loan attacks: bZx (February 2020, $1M), Harvest Finance (October 2020, $34M), Cream Finance (October 2021, $130M), Beanstalk (April 2022, $182M). Each attack used flash loans to amass voting power or manipulate prices. Each attack was atomic. Each attack exploited the gap between the protocol's economic assumptions and the reality of atomic composability: the protocol assumed that accumulating a controlling stake required capital. Flash loans made the assumption false. The assumption was the vulnerability.

The flash loan is a tool. It is neutral. It enables arbitrage that enforces the Law of One Price. It also enables attacks that destroy protocols. The tool is the same. The use determines the outcome. The outcome is a function of the protocol's design. The design must account for atomic composability. Most protocols don't. The ones that don't are exploited. The ones that do survive. The selection is evolutionary. The evolution is the market.

## The reference

Kaihua Qin, Liyi Zhou, and Arthur Gervais, "Quantifying Blockchain Extractable Value: How Dark is the Forest?" (2021). This paper quantified the prevalence of MEV extraction on Ethereum, including flash loan-based arbitrage and attacks. The paper documented the concentration of extraction, the profitability of different strategies, and the arms race dynamics. It is the quantitative complement to Daian et al.'s "Flash Boys 2.0." The numbers are the evidence. The evidence supports the narrative. The narrative is: the dark forest is real, it is concentrated, and it is accelerating.

---

**References:**
- Aave, "Flash Loans," Aave Documentation.
- Kaihua Qin, Liyi Zhou, Arthur Gervais, "Quantifying Blockchain Extractable Value: How Dark is the Forest?" 2021.
- Marble Protocol, "Flash Lending," 2018.
- Related posts: [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage), [MEV](https://blog.hackspree.com/#dex-trading-mev)

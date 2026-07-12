---
title: Maximal Extractable Value
date: 2026-07-12
slug: dex-trading-mev
summary: "MEV is the profit extractable from ordering, including, or excluding transactions within a block. It is the dark forest of crypto — an adversarial environment where every transaction in the mempool is prey. Phil Daian named it in 2019. The name stuck. The phenomenon is the economic layer of blockchain consensus."
tags: dex, mev, flashbots, front-running, sandwich-attacks
series: dex-trading
part: 4
---

Maximal Extractable Value — MEV — is the profit that can be extracted by controlling transaction ordering within a block. The term was coined by Phil Daian and colleagues in the 2019 paper "Flash Boys 2.0." The name is a deliberate reference to Michael Lewis's *Flash Boys*, which documented the high-frequency trading arms race in traditional equity markets. The arms race migrated to crypto. The stakes are the same. The speed is higher.

MEV exists because blockchains are not instant. Transactions are broadcast to a mempool — a waiting room where they sit until a block producer includes them in a block. While in the mempool, transactions are visible. Their contents are known. The block producer can choose which transactions to include, in what order. The choice is economic. Transactions with higher fees are more attractive to include. Transactions that create profit opportunities — arbitrage, liquidation, sandwich attacks — can be front-run, back-run, or sandwiched. The profit is the MEV. The extractor is the searcher. The victim is the user whose transaction created the opportunity.

## The functional origin: Flash Boys

Michael Lewis's *Flash Boys: A Wall Street Revolt* (2014) told the story of the high-frequency trading revolution in U.S. equity markets. The key discovery: the physical distance between exchanges created arbitrage opportunities that could be exploited by traders with faster connections. The traders built microwave towers, laid fiber-optic cable in the straightest possible lines, and colocated servers in exchange data centers. The speed advantage was measured in microseconds. The profit was measured in billions.

The crypto parallel is exact. The mempool is the digital equivalent of the physical distance between exchanges. Transactions in the mempool are visible. A searcher who can read the mempool faster, simulate the transaction's effect, and submit a profitable counterpart transaction before the original is included — that searcher extracts the MEV. The speed advantage is measured in milliseconds. The profit is measured in hundreds of millions.

The difference: in traditional markets, the arms race was infrastructure — microwave towers, fiber routes, exchange colocation. In crypto, the arms race is also infrastructure — but the infrastructure is mempool access, private relay connections, and integration with block builders. The traditional HFT firm needed a microwave license. The crypto MEV searcher needs a connection to a block builder. The barrier is different. The dynamic is the same.

## The forms of MEV

**Front-running.** A searcher sees a large buy order in the mempool. The searcher submits their own buy order for the same token with a higher gas price, getting included before the victim. The searcher's buy moves the price up. The victim's buy executes at the higher price. The searcher sells immediately after, pocketing the difference. The victim paid more. The searcher extracted the difference. The extraction is front-running.

**Sandwich attacks.** A variant. The searcher buys before the victim and sells after. The victim's trade is sandwiched between the searcher's two trades. The victim buys at an artificially elevated price. The searcher profits from both legs. The sandwich is the most profitable form of MEV for liquid token pairs with active mempools.

**Back-running.** A searcher sees a trade that will move the price. The searcher submits a trade after the victim's trade, profiting from the price movement. Back-running is less profitable than front-running — the price has already moved — but less risky — the searcher doesn't need to predict the direction, only react to it. Liquidation of undercollateralized loans is a form of back-running: the searcher sees the price update that triggers the liquidation and submits the liquidation transaction immediately after.

**Just-in-time liquidity.** A searcher sees a large swap in the mempool. They deposit liquidity into the pool just before the swap executes, earn the swap fees, and withdraw immediately after. The LP earns fees with zero inventory risk. JIT liquidity is a form of MEV extraction that benefits the extractor at the expense of passive LPs, whose fee income is diluted.

## The infrastructure: Flashbots

Flashbots was launched in 2020 by Phil Daian, Stephane Gosselin, and colleagues. It is a research and development organization focused on MEV. Its primary product: MEV-Boost, a middleware that separates block building from block proposing. Block builders construct blocks. Block proposers — validators — choose which block to propose. MEV-Boost allows validators to auction their block space to builders. Builders compete to offer the most valuable block. Validators earn the MEV. The auction democratizes MEV access — instead of requiring every validator to run MEV extraction infrastructure, they can outsource to builders and capture the value through competition.

Flashbots also operates a private relay. Searchers submit bundles — groups of transactions that must be executed atomically and in order — to the relay. The relay forwards bundles to builders. The bundles are not broadcast to the public mempool. The privacy prevents front-running of the searcher's own transactions. The relay is the infrastructure that makes MEV extraction possible without the searcher's strategies being copied or front-run.

The Flashbots model has been criticized for centralizing block building. A small number of builders dominate the market. The builders integrate vertically with searchers. The integration concentrates MEV extraction. The concentration creates a new class of intermediary between users and validators. The intermediary extracts rent. The rent is paid by users. The concentration is the subject of ongoing research and regulatory attention.

## The reference

Phil Daian et al., "Flash Boys 2.0: Frontrunning, Transaction Reordering, and Consensus Instability in Decentralized Exchanges" (2019). The paper that named MEV. It documented the prevalence of front-running and sandwich attacks on Ethereum DEXs, measured the extracted value, and proposed architectural responses. The paper is the foundation of MEV research. Every subsequent paper in the field cites it. The phenomenon it documented has grown by orders of magnitude. The growth is the subject of the posts that follow.

---

**References:**
- Phil Daian et al., "Flash Boys 2.0," 2019.
- Michael Lewis, *Flash Boys: A Wall Street Revolt*, W.W. Norton, 2014.
- Flashbots, [docs.flashbots.net](https://docs.flashbots.net).
- Related posts: [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage), [The Order Book](https://blog.hackspree.com/#dex-trading-order-book)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

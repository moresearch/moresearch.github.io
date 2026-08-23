---
title: The Order Book
date: 2026-07-12
slug: dex-trading-order-book
summary: "Every trade begins in an order book. Bids, asks, spread, depth, slippage. The order book is the oldest market microstructure and the foundation of all trading. Understanding it is the prerequisite for everything else: arbitrage, market making, MEV. The book is the game board. The moves happen on it."
tags: dex, trading, order-book, market-microstructure, crypto
series: dex-trading
part: 1
---

Every financial market, whether a 17th-century Amsterdam commodities exchange or a 2026 Solana DEX, rests on the same primitive: a place where buyers and sellers state what they want and at what price. The order book is that place. It is the oldest market microstructure. It is the foundation of every trading strategy that follows.

An order book is a list of bids and asks. A bid is an offer to buy: "I will buy 10 ETH at $3,000 each." An ask is an offer to sell: "I will sell 10 ETH at $3,010 each." The difference between the highest bid and the lowest ask is the spread. The spread is the cost of immediacy — the premium you pay to trade now rather than wait. The spread exists because someone must provide liquidity. The liquidity provider posts bids and asks and waits. The liquidity taker crosses the spread to trade immediately.

The order book is organized by price level. Bids are sorted descending — the highest bid is the best bid. Asks are sorted ascending — the lowest ask is the best ask. The best bid and best ask together form the top of the book. Below them, deeper in the book, lie larger orders at worse prices. The depth of the book at each price level determines how much you can trade before the price moves against you. This is slippage. A deep book absorbs large orders without significant price movement. A shallow book moves sharply on modest volume.

## The functional origin

The order book emerged from the physical trading floors of early modern Europe. The Amsterdam Stock Exchange, founded in 1602 by the Dutch East India Company, was the first permanent market for securities. Traders gathered in a courtyard. They shouted bids and offers. A clerk recorded the trades. The shouting was the order book. The clerk was the exchange.

The London Stock Exchange, founded in 1801, formalized the process. Members — jobbers and brokers — operated under rules. Jobbers made markets: they quoted two prices, a bid and an ask, and stood ready to trade at those prices. Brokers represented clients and executed against jobbers' quotes. The jobber's quote was the order book in miniature. The spread between the bid and the ask was the jobber's profit. The jobber who quoted too wide a spread lost business to other jobbers quoting narrower spreads. The jobber who quoted too narrow a spread absorbed adverse selection — informed traders traded against them when the price was about to move. The jobber's art was balancing spread income against adverse selection losses. That art is now algorithmic. The art is market making. The algorithms are the subject of later posts in this series.

The electronic order book emerged in the 1980s. NASDAQ introduced the Small Order Execution System in 1984, allowing automated execution of small orders. The London Stock Exchange introduced SEAQ, a screen-based quote system, in 1986. The electronic order book eliminated the trading floor. The shouting was replaced by a data structure. The data structure was an order book in memory. The market makers were now algorithms posting quotes to that data structure. The transition from floor to screen was the transition from art to engineering. The engineering is what we trade on today.

## The order book in crypto

Crypto exchanges — both centralized (Binance, Coinbase, Kraken) and decentralized (Serum, dYdX, Hyperliquid) — use electronic order books. The data structure is the same as NASDAQ's. The differences are in settlement and custody.

On a centralized exchange, the order book is a database managed by the exchange. You deposit funds. The exchange credits your account. You place orders. The exchange matches them. You withdraw funds. The exchange is the custodian. The exchange is the counterparty. You trust the exchange. The trust has been violated many times. The violations are the reason for DEXs.

On a decentralized exchange with an order book, the order book is on-chain. Orders are transactions. Matching is performed by the chain's validators or by an off-chain matching engine that settles on-chain. Serum, on Solana, uses an on-chain order book. The chain maintains the order book state. The validators execute matching logic. dYdX, on its own Cosmos chain, uses an off-chain order book with on-chain settlement — the matching engine runs on dYdX's validators, trades settle to the chain. The model is hybrid. The order book is off-chain for speed. The settlement is on-chain for trustlessness.

## The order book's information content

The order book is not just a matching engine. It is an information source. The shape of the book — the distribution of orders across price levels — reveals market sentiment. A book skewed heavily to the bid side suggests buying pressure. A book skewed to the ask side suggests selling pressure. A balanced book suggests equilibrium. The information is probabilistic. It is also actionable. Market-making algorithms read the book to set their quotes. Arbitrage algorithms read the book to detect cross-venue discrepancies. MEV searchers read the book to identify profitable extraction opportunities. The book is the input. The strategy is the output. The book is the game board. The moves happen on it.

Larry Harris's *Trading and Exchanges: Market Microstructure for Practitioners* (2003) is the canonical reference on order book mechanics. Harris was the SEC's chief economist. His book explains every aspect of market microstructure: order types, priority rules, trading costs, market maker obligations, dealer markets vs. auction markets, transparency, fragmentation. It is 600 pages. It is written for practitioners. Every concept in this series — arbitrage, market making, MEV — is built on the microstructure Harris describes. The microstructure changed from floor to screen. The principles didn't. The principles apply to crypto. The crypto application is the subject of the posts that follow.

---

**References:**
- Larry Harris, *Trading and Exchanges: Market Microstructure for Practitioners*, Oxford University Press, 2003.
- Maureen O'Hara, *Market Microstructure Theory*, Blackwell, 1995.
- Related posts: [Algorithmic trading in crypto](https://blog.hackspree.com/#algorithmic-trading-crypto)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

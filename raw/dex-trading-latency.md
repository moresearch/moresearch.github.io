---
title: Latency and Infrastructure
date: 2026-07-12
slug: dex-trading-latency
summary: "In crypto trading, latency is the edge. The bot with the fastest mempool access, the lowest-latency path to the sequencer, the colocated server wins the arbitrage. The difference is milliseconds. The investment in infrastructure is the rent paid for speed. The rent is extracted from slower traders."
tags: dex, latency, infrastructure, mempool, colocation
series: dex-trading
part: 8
---

Latency is the single most important competitive variable in algorithmic trading. Strategy matters. Risk management matters. Capital matters. But given equivalent strategy and risk management and capital, the faster participant wins. The win is measured in milliseconds. The investment to achieve that speed is measured in millions.

The latency stack for crypto trading has four layers. Each layer is a source of delay. Each layer is an opportunity for optimization.

**Network latency.** The time for a packet to travel from the trader's server to the exchange's matching engine or the blockchain's sequencer. Minimized by physical proximity — colocating in the same datacenter as the exchange or sequencer — and by network topology — choosing the shortest path, avoiding congested routes, using dedicated fiber. The speed of light in fiber is approximately 200,000 km/s. Over 1,000 km, that is 5 milliseconds. Over 100 km, 0.5 milliseconds. The difference between a server in Frankfurt and a server in London trading on a Frankfurt-based exchange is 5 milliseconds. In that 5 milliseconds, the price can move. The edge is gone. The trade loses.

**Mempool latency.** On blockchains, transactions are not sent directly to the block producer. They are broadcast to a mempool — a peer-to-peer network of nodes. The time between submitting a transaction and it being visible to the block producer depends on the node's position in the network topology, the number of hops to the producer, and the propagation delay at each hop. A trader who runs their own node, connected directly to high-staked validators or sequencers, sees transactions before a trader relying on public RPC endpoints. The private node is the latency edge. The edge is purchased through infrastructure.

**Execution latency.** The time between the transaction arriving at the exchange or sequencer and being included in a block or matched. On centralized exchanges, this is the matching engine's processing time — typically microseconds. On blockchains, this is the block time — the interval between blocks. On Ethereum L1, block time is 12 seconds. On Solana, 400 milliseconds. On Arbitrum, sub-second. The block time is the minimum latency for on-chain execution. Faster chains enable faster strategies. The chain choice is a latency decision.

**State latency.** The time between a state change occurring and the trader's system being aware of it. A trade on Uniswap changes the pool's price. The change is included in a block. The block is propagated through the network. The trader's node receives the block, updates its state, and triggers the strategy. The delay between the trade occurring and the strategy reacting is state latency. Minimized by running a full node, subscribing to block production directly, and processing state updates in parallel with strategy evaluation. The full node is expensive to run. The expense is the cost of being fast.

## The functional origin: the transatlantic cable

The latency arms race in finance began in the 19th century. The first transatlantic telegraph cable, laid in 1866, reduced communication time between London and New York from weeks (by ship) to minutes (by telegraph). The cable was used for arbitrage: prices in London and New York could be compared in near-real-time. The trader who received the cable first could trade before the information was widely known. The cable was the latency edge.

In the 1980s, fiber-optic cables replaced copper. In the 2000s, microwave networks replaced fiber for the most latency-sensitive routes — microwaves travel through air at the speed of light, which is faster than light through glass fiber. In the 2010s, laser networks in space were proposed for intercontinental routes. The arms race is continuous. The technology changes. The principle doesn't: faster information means faster reaction, faster reaction means profitable trades, profitable trades pay for the infrastructure. The infrastructure is the rent. The rent is extracted from slower participants.

Michael Lewis's *Flash Boys* (2014) documented the modern latency arms race in U.S. equity markets. The key discovery: a new fiber route between Chicago and New York, laid as straight as physically possible through mountains and under rivers, reduced latency by 3 milliseconds. The route cost $300 million. The route was built by a trading firm. The firm's name was Spread Networks. The name was the strategy. The spread was the profit. The network was the means.

## The crypto latency stack today

Crypto latency infrastructure is evolving along the same trajectory as traditional finance, compressed into a decade instead of a century. The major developments:

**Private relays.** Flashbots and similar services provide private transaction submission, bypassing the public mempool. The private relay reduces mempool latency to near-zero — the transaction goes directly to the builder. The privacy also prevents front-running of the trader's own transactions. The relay is the infrastructure. The fee is the cost.

**Sequencer colocation.** On L2 rollups, the sequencer orders transactions. Colocating with the sequencer provides minimum network latency. The colocation is offered by some L2s as a paid service. The service is the latency edge. The edge is purchased.

**Validator connections.** On proof-of-stake chains, validators propose blocks. A trader with direct connections to high-staked validators can submit transactions that are included in the next block with higher probability. The connections are built through relationships, infrastructure sharing, and direct payments. The connections are the edge. The edge is relational.

**Chain-specific optimization.** Each chain has different latency characteristics. Solana's Gulf Stream forwards transactions to validators before the current block is finalized. Arbitrum's sequencer orders transactions on a first-come, first-served basis. The optimal strategy varies by chain. The trader who understands the chain-specific latency model has an edge over the trader who treats all chains as equivalent. The understanding is the edge. The edge is informational.

---

**References:**
- Michael Lewis, *Flash Boys: A Wall Street Revolt*, W.W. Norton, 2014.
- Phil Daian et al., "Flash Boys 2.0," 2019.
- "First-Spammed, First-Served: MEV Extraction on Fast-Finality Blockchains," June 2025.
- Related posts: [MEV](https://blog.hackspree.com/#dex-trading-mev), [Arbitrage](https://blog.hackspree.com/#dex-trading-arbitrage)


Trading infrastructure is distributed systems engineering. The order book, the AMM, the matching engine, the relay — each is a component in a latency-critical distributed system. The engineering constraints are the same as any real-time system: throughput, latency, reliability, correctness under concurrency. The domain is finance. The engineering is systems.

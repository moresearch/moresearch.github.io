---
title: "NATS JetStream vs Kafka: Streaming Without the Partition Tax"
date: 2026-07-16
slug: nats-jetstream-vs-kafka-streaming-without-partitions
summary: Kafka's partition model is its superpower and its shackle — it gives you ordering guarantees at the cost of fixed parallelism and painful rebalancing. NATS JetStream takes the opposite approach: messages are stored in a single append-only stream, and consumers read independently with their own cursors. With Orbit.go's partitioned consumer groups, you get Kafka-style key-based ordering without Kafka's partition management overhead. This post explains the architectural difference and why it matters for teams that need streaming but don't want to operate Kafka.
tags: nats, jetstream, kafka, streaming, partitioned-consumers, distributed-systems, orbit
---

Kafka is the default answer for event streaming. It earned that position honestly — it solves a hard problem with a clear architecture. But operating Kafka is itself a hard problem. The partition is the unit of parallelism, and that design choice cascades into every operational concern: rebalancing, key distribution, retention, and scaling.

NATS JetStream approaches streaming from the opposite direction. Instead of partitioning the log for parallelism, it separates storage (the stream) from consumption (the consumer). Multiple consumers read the same stream independently, each at its own pace, with its own filter and acknowledgement state. Parallelism comes from concurrent consumers, not from partitions. Ordering comes from consumer configuration, not from partition assignment.

[Orbit.go's partitioned consumer groups](https://nats.io/blog/orbit-partitioned-consumer-groups/) — a pure client-side library — bridge the remaining gap, bringing Kafka-style key-based ordering to JetStream without requiring stream partitioning. The result is streaming that's simpler to operate and more flexible to scale, with the ordering guarantees Kafka users expect.

## The Kafka partition model

Kafka's fundamental abstraction is the partitioned log. Each topic is split into N partitions. Each partition is an ordered, immutable sequence of messages. Producers write to partitions based on a partitioning key. Consumers read from partitions in order. The partition is the unit of parallelism — you can have at most one consumer per partition (in the same consumer group) and at most as many active consumers as there are partitions.

This model provides strong guarantees:

- **Ordering per partition:** Messages with the same key land in the same partition and are consumed in order. Ordering is guaranteed within a partition, not across partitions.
- **Parallelism via partition count:** To increase throughput, increase the partition count. More partitions means more concurrent consumers.
- **Durability via replication:** Each partition has a leader and N followers. The leader handles reads and writes. Followers replicate the log. If the leader fails, a follower takes over.

These guarantees come with operational constraints:

- **Fixed parallelism ceiling.** Partition count is set at topic creation and is hard to change. Increasing partitions is possible but disruptive — it changes the key-to-partition mapping, breaking ordering for keys that move to new partitions. Decreasing partitions is effectively impossible without deleting and recreating the topic. You must provision for peak parallelism at topic creation, paying for capacity you may not need for months.

- **Rebalancing is expensive.** When a consumer joins or leaves a consumer group, Kafka triggers a rebalance: partitions are reassigned across the new set of consumers. During rebalance, consumption pauses. For large consumer groups with many partitions, rebalancing can take seconds to minutes. The cooperative rebalance protocol (KIP-429) mitigates this but doesn't eliminate it. Every deployment, every scale-up, every crash triggers a rebalance.

- **Hot partitions are undivisible.** If a single key generates disproportionate traffic (a "hot" partition), you can't split that partition without breaking ordering. The partition is atomic. You can't subdivide it. You can't parallelize within it. The partition is the bottleneck, and you can't make it wider.

- **Operational complexity.** Kafka requires ZooKeeper (pre-3.3) or KRaft (3.3+). It requires careful JVM tuning. It requires disk provisioning for partitioned logs with configurable retention. It requires monitoring of under-replicated partitions, consumer lag, and broker resource utilization. Running Kafka well is a specialized skill. Running Kafka at scale is a team.

## The JetStream model: one stream, many consumers

JetStream inverts the relationship between storage and consumption. A **stream** is a named append-only log that captures messages on one or more subjects. A **consumer** is a named view into a stream — a cursor with a filter, an acknowledgement policy, and a delivery mode.

```go
// One stream captures all order events
js.AddStream(&nats.StreamConfig{
    Name:     "ORDERS",
    Subjects: []string{"orders.>"},
    Storage:  nats.FileStorage,
    Replicas: 3,
})

// Multiple consumers read independently
// Consumer A: fulfillment, reading from the beginning
js.AddConsumer("ORDERS", &nats.ConsumerConfig{
    Durable:       "fulfillment",
    FilterSubject: "orders.created",
    DeliverPolicy: nats.DeliverAllPolicy,
    AckPolicy:     nats.AckExplicitPolicy,
})

// Consumer B: analytics, only new messages
js.AddConsumer("ORDERS", &nats.ConsumerConfig{
    Durable:       "analytics",
    FilterSubject: "orders.>",
    DeliverPolicy: nats.DeliverNewPolicy,
})

// Consumer C: fraud detection, last per subject
js.AddConsumer("ORDERS", &nats.ConsumerConfig{
    Durable:       "fraud-check",
    FilterSubject: "orders.payment.*",
    DeliverPolicy: nats.DeliverLastPerSubjectPolicy,
})
```

Notice what's missing: no partition count. No key-to-partition mapping. No partition assignment. The stream is a single log. Consumers are independent views. Each consumer has its own cursor. Each consumer acknowledges independently. The stream doesn't care how many consumers exist or what they're doing — it just appends messages.

Parallelism works differently in this model. A single consumer can be read by multiple instances concurrently. NATS distributes messages across instances that are pulling from the same consumer — no rebalancing, no partition assignment, no pause in consumption when an instance joins or leaves:

```go
// Instance 1, 2, and 3 all pull from the same consumer
sub, _ := js.PullSubscribe("orders.created", "fulfillment")
for {
    msgs, _ := sub.Fetch(10)
    for _, msg := range msgs {
        process(msg)
        msg.Ack()
    }
}
```

The three instances collectively consume from `fulfillment`. Each `Fetch(10)` returns up to 10 messages. NATS distributes messages across the active pullers. If Instance 1 crashes, Instances 2 and 3 continue — their pull requests are now served faster because there are fewer pullers competing. No rebalance. No partition reassignment. No pause.

This is elastic by default. You deploy more instances, they start pulling, throughput increases. You deploy fewer, throughput decreases. No topic configuration to update. No partition count to pre-provision. The parallelism is dynamic, not static.

## The ordering trade-off

JetStream's elastic model has a trade-off: ordering is not guaranteed across concurrent pulls. If Instance 1 pulls messages 1-10 and Instance 2 pulls messages 11-20 simultaneously, and Instance 2 processes faster, message 11 might be acknowledged before message 1. Messages are stored in order. They may be processed out of order.

For many workloads, this is fine. If each message is independent — a notification, a metric, a log entry — processing order doesn't matter. But for workloads that require per-key ordering (all events for Customer A processed in order), it's a real constraint.

The traditional solution is to set `MaxAckPending` to 1, which serializes all processing through a single message in flight at a time. This guarantees order but kills throughput. You've traded parallelism for ordering, and you're paying for it on every message, even ones with different keys that could safely be processed in parallel.

This is exactly the gap that Orbit.go's partitioned consumer groups fill.

## Partitioned consumer groups: Kafka semantics, JetStream simplicity

[Orbit.go](https://github.com/synadia-io/orbit.go) implements what Jean-Noël Moyne describes as "functionally equivalent to what Apache Kafka calls 'consumer groups' and how they implement partitioning" — entirely on the client side.

The key insight: most real-world ordering requirements are per-key, not global. You need Customer A's events processed in order. You don't need all customers' events processed in order. Kafka enforces this via partitions — same key → same partition → same consumer → ordered processing. Orbit.go enforces this via subject token hashing on top of JetStream consumers — same key → same member → ordered processing.

### Static partitioned consumer groups

Static groups require the stream to have a partition number as the first subject token (achievable via a stream subject transform at ingest). The library maps member names to partition numbers using consistent hashing:

```go
// Stream subjects include partition number
// orders.{partition}.created, orders.{partition}.paid, ...

group := orbit.CreateStaticGroup("order-processors", streamConfig, memberNames)
group.Join("fulfillment-member", consumerConfig, func(msg jetstream.Msg) {
    // Messages for this member's partitions arrive in order
    processInOrder(msg)
    msg.Ack()
})
```

Guarantees:
- Each partition is handled by exactly one member at a time
- Messages within a partition are processed in order
- Multiple partitions can be processed in parallel by different members
- If a member instance crashes, NATS 2.11's pinned consumer feature ensures the replacement instance picks up where the old one left off

Static groups are faster and use fewer server resources, but membership is fixed at creation. No adding members at runtime. The trade-off is latency and resource efficiency for operational flexibility.

### Elastic partitioned consumer groups

Elastic groups work on any existing stream — no partition token required. The library materializes the group as a work queue stream that sources from the original, inserting partition numbers during sourcing:

```go
group := orbit.CreateElasticGroup("order-processors", sourceStream, maxMembers)
group.AddMember("fulfillment")   // Add at runtime
group.AddMember("audit")         // Add at runtime
group.DropMember("fulfillment")  // Remove at runtime

group.Join("fulfillment", consumerConfig, func(msg jetstream.Msg) {
    processInOrder(msg)
    msg.Ack()
})
```

The work queue stream holds copies of messages, so consumption lag can be monitored by checking the work queue stream size. You can cap the work queue stream size; if it hits the limit, sourcing pauses briefly. This prevents unbounded memory consumption from a work queue that outpaces its consumers.

Elastic groups use more server resources and add slight latency (the materialization step), but you get runtime elasticity. Add members when load increases. Drop members when load decreases. No partition reassignment. No key-to-partition remapping. No rebalancing pause.

## Side by side: Kafka vs JetStream with partitioned consumers

| Concern | Kafka | JetStream + Orbit.go |
|---------|-------|----------------------|
| **Unit of storage** | Partitioned topic | Single stream |
| **Unit of parallelism** | Partition (fixed at creation) | Consumer instances (dynamic) |
| **Ordering** | Per partition | Per consumer (or per member in a group) |
| **Adding parallelism** | Increase partitions (disruptive) | Deploy more instances (elastic) |
| **Rebalancing** | Pauses consumption | No pause (pull-based distribution) |
| **Hot partition** | Undivisible bottleneck | Elastic member can be dedicated |
| **Retention** | Time or size per partition | Time, size, or count per stream |
| **Replication** | Per partition (leader/follower) | Per stream (Raft across cluster nodes) |
| **Operational complexity** | ZooKeeper/KRaft, JVM tuning, partition monitoring | Single Go binary, streams and consumers |
| **Key-based ordering** | Built-in (partition by key) | Via Orbit.go (hash by subject token) |

The fundamental difference: Kafka builds ordering into the storage layer (partitions). JetStream builds ordering into the consumption layer (consumers and consumer groups). The JetStream approach is more flexible because consumers can be reconfigured without touching the stream. The Kafka approach provides stronger guarantees because the ordering is physically enforced by the log structure. Whether one is better depends on whether you value operational flexibility or storage-level guarantees more.

## When Kafka still makes sense

Kafka's partition model is not a mistake. It's the right design for workloads where:

**Ordering is truly global.** If every message in a topic must be processed in exact append order (not just per-key order), Kafka partitions give you that — one partition, one consumer. JetStream can do this with `MaxAckPending: 1` but at the same throughput cost.

**You need the Kafka ecosystem.** Kafka Connect provides a rich set of source and sink connectors. Kafka Streams provides a sophisticated stream processing library with exactly-once semantics, stateful operations (joins, aggregations, windows), and an interactive query API. ksqlDB provides SQL over streams. The ecosystem is deep and mature. If your architecture depends on these tools, Kafka is the right choice.

**Your team already operates Kafka well.** Running Kafka is a skill. If your team has invested in that skill and the operational burden is manageable, the migration cost to NATS may not be justified. Don't migrate because NATS is simpler. Migrate because the operational burden of Kafka is a meaningful drag on your team's velocity.

**You need compacted topics.** Kafka's log compaction retains the latest value for each key, enabling table-like semantics for changelogs. JetStream doesn't have a direct equivalent (though `DiscardNewPerSubject` with `MaxMsgsPerSubject: 1` provides a rough approximation for single-message-per-subject use cases).

For greenfield systems that need streaming but don't already depend on the Kafka ecosystem, JetStream with Orbit.go partitioned consumer groups addresses the same ordering requirements with significantly less operational overhead. You get key-based ordering. You get parallel consumption. You don't get partition management, rebalance pauses, or ZooKeeper. For most teams, that's a good trade.

---

**References:**
- [Client-side Partitioned Consumer Groups for JetStream](https://nats.io/blog/orbit-partitioned-consumer-groups/) — Jean-Noël Moyne, NATS Blog
- [Orbit.go on GitHub](https://github.com/synadia-io/orbit.go)
- [NATS JetStream Documentation](https://docs.nats.io/nats-concepts/jetstream)
- [Kafka Documentation — Topic Partitioning](https://kafka.apache.org/documentation/#intro_concepts_and_terms)

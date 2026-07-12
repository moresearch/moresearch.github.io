---
title: NATS pub/sub beats REST for microservices
date: 2026-07-12
slug: nats-pubsub-microservices
summary: "REST couples services by address. NATS couples them by subject. One is a phone call where you must know the number. The other is a broadcast where you only need the frequency. Modularity lives in the difference."
tags: nats, jetstream, pubsub, microservices, golang, rest, modularity
---

REST is the default for microservices communication. It is the default because it is familiar, not because it is good. Every service exposes HTTP endpoints. Every caller knows every callee's address. The system is a web of explicit dependencies, each hardcoded in configuration or service discovery.

NATS inverts this. Services publish to subjects. Services subscribe to subjects. No service knows any other service exists. The subject is the interface. The publisher has no idea who is listening. The subscriber has no idea who is publishing. They agree on the subject name and the message schema. Nothing else.

This is not a preference. It is an architectural difference with measurable consequences for coupling, scalability, and resilience. This post explains what NATS is, how JetStream extends it, why pub/sub produces better modularity than REST, and when each approach makes sense.

## What NATS is

NATS is a messaging system. It is not a message broker in the RabbitMQ sense. It is closer to a network switch for messages. Clients connect to NATS servers. Servers route messages between clients based on subject subscriptions. The core is small, fast, and does one thing: move messages from publishers to subscribers with minimal latency and maximal throughput.

The protocol is text-based and trivial. A client sends `PUB <subject> <size>\r\n<payload>\r\n`. A client sends `SUB <subject> <sid>\r\n`. The server matches subjects to subscribers and delivers the message. That is the core. No persistence. No acknowledgements. No transactions. Just publish, subscribe, deliver.

### Subjects and wildcards

Subjects are hierarchical tokens separated by dots: `orders.created`, `inventory.updated.us-east`, `payment.processed.visa`. Subscribers can use wildcards:

- `*` matches one token: `orders.*` matches `orders.created` and `orders.cancelled`, not `orders.payment.authorized`
- `>` matches all remaining tokens: `orders.>` matches `orders.created`, `orders.payment.authorized`, `orders.items.returned.refund`

Wildcards let subscribers express interest in categories of events without knowing every specific subject. A logging service subscribes to `>.>` and receives everything. An inventory service subscribes to `orders.*` and receives only order lifecycle events. The subject namespace is the API. The wildcards are the subscriptions. The routing is automatic.

### Spatial decoupling

In REST, Service A must know Service B's address. When B moves, A's configuration must change. When B scales to multiple instances, a load balancer must be configured. The load balancer must know all instance addresses. Health checks must be written. The address is the coupling.

In NATS, Service A publishes to `orders.created`. Service B subscribes to `orders.created`. Neither knows the other's address. Neither knows the other exists. The NATS server — or cluster of servers — routes the message. If B moves to a different machine, region, or cloud, nothing changes. The subject is the only address. The subject does not change.

### Queue groups: scale without ceremony

NATS queue groups distribute messages across multiple subscribers without a load balancer:

```go
// Three instances, one queue group, automatic load distribution
nc.QueueSubscribe("orders.process", "order-workers", func(m *nats.Msg) {
    processOrder(m)
})
```

That is it. No load balancer. No health checks. No instance registry. NATS distributes messages round-robin across queue group members. If an instance crashes, NATS stops routing to it. If a new instance starts, NATS includes it. The publisher never knew how many instances existed. The publisher still doesn't. The subject is the interface. The instances are an implementation detail.

### Clustering and superclusters

A single NATS server can handle millions of messages per second. For scale beyond one machine, NATS clusters route messages between servers. For scale beyond one datacenter, gateways connect clusters into superclusters.

**Clustering (routes):** Servers within a cluster form a full mesh. Each server connects to every other server. Routes use a dedicated port. Subscription interest is gossiped between servers — a server only forwards messages to peers that have matching subscribers. For N nodes, N(N−1)/2 connections. A three-node cluster: three connections. A five-node cluster: ten. The protocol is designed for this. The overhead is minimal.

**Gateways (superclusters):** Gateways connect entire clusters. Three clusters of five nodes each: a full node-to-node mesh would require 105 connections. With gateways, each node connects to one node in each remote cluster — 30 connections. Interest propagation is cluster-scoped. A message published in Cluster A is only forwarded to Cluster B if Cluster B has expressed interest in that subject. Optimistic forwarding on first message, suppressed thereafter.

**Leaf nodes:** For edge deployments — IoT, retail, remote offices — leaf nodes extend a cluster across security boundaries without requiring bidirectional connectivity. Local clients authenticate locally. The leaf connection acts as a NATS client to the hub, exporting subjects the edge can publish and importing subjects the edge can subscribe to. Local traffic stays local. Remote traffic traverses the leaf. Queue semantics are preserved: local subscribers are preferred before forwarding across the leaf.

This architecture scales from a developer's laptop (`nats-server -js`) to a planet-scale messaging fabric. The subject namespace is global. The routing is automatic. The addressing is semantic, not spatial.

## What JetStream is

Core NATS is fire-and-forget. If a subscriber is offline, the message is lost. For systems that need persistence, replay, or guaranteed delivery, NATS provides JetStream — a persistence layer built into the NATS server.

JetStream decouples storage from consumption. **Streams** store messages. **Consumers** read from streams. This separation is the key architectural insight. Multiple consumers can read the same stream independently, each at its own pace, with its own acknowledgement state, filtering by subject, starting from different points in the log.

### Streams: the storage layer

A stream is a named, append-only log that captures messages published on one or more subjects:

```go
js, _ := jetstream.New(nc)

stream, _ := js.CreateStream(ctx, jetstream.StreamConfig{
    Name:      "ORDERS",
    Subjects:  []string{"orders.>"},
    Storage:   nats.FileStorage,
    Replicas:  3,                       // Raft across 3 nodes
    MaxAge:    7 * 24 * time.Hour,      // Retain 7 days
    MaxBytes:  10 * 1024 * 1024 * 1024, // 10 GB ceiling
})
```

Stream configuration determines retention policy:

- **LimitsPolicy** (default): keep messages until count, bytes, or age limits are reached. This is for event sourcing, replay, or audit trails.
- **WorkQueuePolicy**: delete each message after it's been acknowledged by *any* consumer. This is for job queues — process once, discard.
- **InterestPolicy**: delete messages only after *all* consumers have acknowledged them. This is for fan-out where multiple services process the same event.

Storage can be file-backed (survives restart) or memory-backed (faster, lost on restart). Replication uses Raft across 1, 3, or 5 cluster nodes. Odd numbers for quorum. A three-replica stream tolerates one node failure. A five-replica stream tolerates two.

### Consumers: the read cursors

A consumer is a named view into a stream. Multiple consumers read independently. Each has its own position, filter, and acknowledgement state.

**Pull consumers** are the default for work queues. The client explicitly requests messages:

```go
consumer, _ := stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
    Durable:       "order-processor",
    FilterSubject: "orders.created",
    AckPolicy:     jetstream.AckExplicitPolicy,
    AckWait:       30 * time.Second,
})

// Fetch a batch of 10 messages
batch, _ := consumer.Fetch(10)
for msg := range batch.Messages() {
    process(msg)
    msg.Ack()       // Explicit acknowledgement
    msg.AckSync()   // Double-ACK: wait for server confirmation
}
```

**Push consumers** are for low-latency streaming to a single instance. The server pushes messages to a delivery subject:

```go
consumer, _ := stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
    Name:          "order-streamer",
    DeliverPolicy: jetstream.DeliverNewPolicy,
    AckPolicy:     jetstream.AckExplicitPolicy,
})

cc, _ := consumer.Consume(func(msg jetstream.Msg) {
    process(msg)
    msg.Ack()
})
```

Consumer replay policies determine where to start reading:
- `all`: replay every message from the beginning
- `last`: start with the last message, then follow live
- `new`: only messages arriving after subscription
- `by_start_sequence`: start from a specific sequence number
- `by_start_time`: start from messages at or after a timestamp
- `last_per_subject`: deliver the last message for each unique subject

### Exactly-once semantics

JetStream provides exactly-once through two complementary mechanisms.

**Publish-side deduplication.** Set the `Nats-Msg-Id` header when publishing. Within the stream's duplicate window (default 2 minutes, configurable), messages with the same ID are silently discarded:

```go
msg := nats.NewMsg("orders.created")
msg.Header.Set("Nats-Msg-Id", orderID)
msg.Data = payload
js.PublishMsg(ctx, msg)
```

If the publisher crashes before receiving the publish acknowledgement and re-publishes on restart, the duplicate is suppressed. The publisher retries safely. The stream stores exactly one copy.

For infinite deduplication beyond the time window, use `DiscardNewPerSubject` with `MaxMessagesPerSubject = 1`. Publishing to the same subject with an existing message fails. This behaves like a SQL `INSERT` with a unique constraint on the subject.

**Consume-side double acknowledgement.** `AckSync()` sends the ACK with a reply subject and blocks until the server confirms receipt. If the ACK is lost and the consumer crashes before the confirmation, the message is redelivered and must be processed idempotently:

```go
func process(msg jetstream.Msg) error {
    if alreadyProcessed(msg) {
        msg.AckSync() // Still ACK — was processed, just not acknowledged
        return nil
    }
    if err := doWork(msg); err != nil {
        msg.Nak() // Return to queue for retry
        return err
    }
    msg.AckSync() // Guaranteed: server confirmed receipt of ACK
    return nil
}
```

Publish-side deduplication + double-acknowledged consumption + idempotent processing = true end-to-end exactly-once. Messages are never lost. Messages are never duplicated. The guarantee is as strong as any message broker provides and stronger than most.

### Temporal decoupling

REST fails when the callee is unavailable. The caller must retry, back off, circuit-break, or fail. Each retry loop is hand-rolled. Each circuit breaker has slightly different thresholds. Reliability is the sum of individually implemented strategies that were never tested together.

JetStream provides temporal decoupling by default:

```go
// Publisher: fire and forget
js.Publish(ctx, "orders.created", event)

// Consumer: receives when ready, even hours later
consumer.Consume(func(msg jetstream.Msg) {
    process(msg)
    msg.Ack()
})
```

The publisher publishes and moves on. If no subscriber is online, the message waits in the stream. If the subscriber crashes mid-processing, the message is redelivered after `AckWait` expires. The publisher doesn't retry. The subscriber doesn't need to be online when the message is published. The infrastructure handles reliability. The services handle business logic.

## The REST comparison

Here is an order service calling an inventory service over REST in Go:

```go
func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    body, _ := json.Marshal(ReserveRequest{
        ProductID: order.ProductID,
        Quantity:  order.Quantity,
    })
    resp, err := http.Post(
        "http://inventory-service:8080/api/reserve", // knows address
        "application/json",
        bytes.NewReader(body),
    )
    if err != nil {
        return fmt.Errorf("inventory unreachable: %w", err) // knows availability
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 { // knows success semantics
        return fmt.Errorf("reservation failed: %s", resp.Status)
    }
    return s.db.Insert(ctx, order)
}
```

The order service knows the inventory service's URL, endpoint path, request format, success semantics, and availability requirement. If the inventory service moves, the order service must be reconfigured. If the inventory service is down, the order fails. If the inventory service's API changes, the order service must be updated. Five types of shared knowledge. Five dimensions of coupling. This is one dependency. A typical microservices system has dozens.

Here is the same workflow over NATS:

```go
func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    event, _ := json.Marshal(OrderRequested{
        OrderID:   order.ID,
        ProductID: order.ProductID,
        Quantity:  order.Quantity,
    })
    // Fire and forget: publish to the stream
    js.Publish(ctx, "orders.created", event)

    // Request-reply: ask inventory, don't know who answers
    msg, err := nc.Request("inventory.reserve", event, 5*time.Second)
    if err != nil {
        return fmt.Errorf("reservation failed: %w", err)
    }
    var result ReserveResult
    json.Unmarshal(msg.Data, &result)
    if !result.Success {
        return fmt.Errorf("reservation denied: %s", result.Reason)
    }
    return s.db.Insert(ctx, order)
}

// Inventory service: subscribers don't know publishers exist
func (s *InventoryService) Run() {
    nc.Subscribe("inventory.reserve", func(m *nats.Msg) {
        var req OrderRequested
        json.Unmarshal(m.Data, &req)
        err := s.reserveStock(req.ProductID, req.Quantity)
        result := ReserveResult{Success: err == nil}
        if err != nil { result.Reason = err.Error() }
        reply, _ := json.Marshal(result)
        nc.Publish(m.Reply, reply)
    })
}
```

The order service knows two subjects: `orders.created` and `inventory.reserve`. It does not know the inventory service exists. It does not know how many instances are listening. It does not know where they are deployed. If the inventory service moves — nothing changes. If three more inventory instances start — nothing changes. If the inventory service's internal implementation is rewritten — nothing changes, as long as the subject and schema are preserved. The subject is the contract. The subject is the only coupling.

## What coupling actually means

Henney defined coupling as shared knowledge between components. REST maximizes shared knowledge. The caller knows the callee's address, API shape, availability window, response semantics, error taxonomy, authentication mechanism, and rate limits. Each is a dimension of coupling. When the callee changes any of these, the caller must adapt or break.

NATS minimizes shared knowledge. Publisher and subscriber agree on the subject name and the message schema. Neither knows the other exists. Neither knows how many of the other exist. Neither knows where the other is deployed. Neither knows whether the other is currently online. Subject name and schema: that is the total shared knowledge. Everything else is hidden.

This is Parnas's criterion applied to inter-service communication. Hide the volatile decisions — deployment location, instance count, availability status, internal API changes — behind a stable interface. The subject is the stable interface. The service is the volatile implementation. The subject name doesn't change when the implementation scales, moves, or restarts. The coupling is minimized. The modularity is real.

## When REST still makes sense

REST is not universally wrong. It is appropriate in specific conditions:

**External APIs at the system boundary.** Mobile apps, browsers, third-party integrations speak HTTP. The edge of the system faces outward. The interior of the system should not. Expose REST at the boundary. Use NATS internally. The boundary is where HTTP belongs. The interior is where it doesn't. External clients need synchronous request-reply with standard protocols and well-known ports. Internal services need decoupling and resilience. The two requirements are different. The two protocols should be different.

**Genuinely stable dependencies.** If Service A will always depend on Service B, and B's API is provably stable, and B's deployment will never change in ways that affect A — REST is fine. This condition is rarer than architects assume. Most dependencies that appear stable at design time are unstable at year three. The REST interface hardcodes assumptions that time will invalidate. The NATS subject survives the invalidation because it doesn't encode the assumptions.

For everything else — internal service communication, event-driven workflows, systems that scale, systems that survive partial failure, systems where the dependency graph will evolve — NATS pub/sub backed by JetStream is not just better. It is what REST pretends to be: modular, decoupled, and resilient. REST shares knowledge by default. NATS hides knowledge by default. Modularity lives in what services don't know about each other. NATS lets them know less. That is the definition of better architecture.

---

**References:**
- [NATS Documentation](https://docs.nats.io/) — core concepts, JetStream, clustering
- [NATS Go Client](https://github.com/nats-io/nats.go) — `nats.go` with JetStream v2 API
- [NATS by Example](https://natsbyexample.com/) — runnable patterns across clients
- Related posts: [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Henney's Microservices](https://blog.hackspree.com/#kevlin-henney), [I, Pencil](https://blog.hackspree.com/#i-pencil)

---
title: "NATS pub/sub beats REST for microservices"
date: 2026-07-12
slug: nats-pubsub-microservices
summary: "REST couples services by address. NATS couples them by subject. One is a phone call where you must know the number. The other is a broadcast where you only need to know the topic. Modularity lives in the difference."
tags: nats, pubsub, microservices, golang, rest, modularity
---

REST is the default microservices communication pattern. It is the default because it is familiar, not because it is good. Every service exposes HTTP endpoints. Every caller knows every callee's address. The system is a web of explicit dependencies, each one hardcoded in configuration or service discovery. When Service A calls Service B via REST, A knows that B exists, where B lives, what B's API looks like, and whether B is currently available. That is not modularity. That is coupling with extra steps.

NATS pub/sub inverts this. Services publish to subjects. Services subscribe to subjects. No service knows any other service exists. The subject is the interface. The publisher has no idea who is listening. The subscriber has no idea who is publishing. They agree on the subject name and the message format. Nothing else. This is spatial decoupling. This is temporal decoupling. This is modularity.

## The REST way

Here is an order service calling an inventory service over REST in Go:

```go
// Order service: needs to reserve inventory
func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    // Step 1: Reserve inventory via REST
    body, _ := json.Marshal(ReserveRequest{
        ProductID: order.ProductID,
        Quantity:  order.Quantity,
    })
    resp, err := http.Post(
        "http://inventory-service:8080/api/reserve",  // knows address
        "application/json",
        bytes.NewReader(body),
    )
    if err != nil {
        return fmt.Errorf("inventory service unreachable: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("reservation failed: %s", resp.Status)
    }
    // Step 2: Save order
    return s.db.Insert(ctx, order)
}
```

The order service knows the inventory service's URL. It knows the endpoint path. It knows the request format. It knows that a 200 means success. It knows the service must be reachable at call time. If the inventory service moves, the order service must be reconfigured. If the inventory service is down, the order fails. If the inventory service's API changes, the order service must be updated. The coupling is total.

Scale this across fifty services. Each knows the addresses of the others it depends on. Each retry loop is hand-rolled. Each circuit breaker is hand-rolled. Each service discovery integration is hand-rolled. The system is a distributed monolith with HTTP as the in-process communication bus. The network is the new call stack. The latency is the new overhead. The modularity did not improve. The deployment got harder.

## The NATS way

Same system, pub/sub:

```go
// Order service: publishes an event, doesn't know who listens
func (s *OrderService) CreateOrder(ctx context.Context, order Order) error {
    // Step 1: Publish intent
    event, _ := json.Marshal(OrderRequested{
        OrderID:   order.ID,
        ProductID: order.ProductID,
        Quantity:  order.Quantity,
    })
    nc.Publish("orders.requested", event)

    // Step 2: Wait for inventory to respond (request-reply)
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

// Inventory service: subscribes, doesn't know who publishes
func (s *InventoryService) Run() {
    nc.Subscribe("inventory.reserve", func(m *nats.Msg) {
        var req OrderRequested
        json.Unmarshal(m.Data, &req)

        err := s.reserveStock(req.ProductID, req.Quantity)
        result := ReserveResult{Success: err == nil}
        if err != nil {
            result.Reason = err.Error()
        }

        reply, _ := json.Marshal(result)
        nc.Publish(m.Reply, reply)
    })
}
```

The order service knows two subjects: `orders.requested` and `inventory.reserve`. It does not know the inventory service exists. It does not know how many instances are listening. It does not know where they are deployed. The inventory service knows one subject: `inventory.reserve`. It does not know the order service exists. If the inventory service moves to a different machine, different region, different cloud — nothing changes. The subject is the contract. The subject is the only coupling.

The difference is not syntax. It is architecture. REST couples spatially — the caller must know the callee's address. NATS couples semantically — the publisher and subscriber agree on the meaning of a subject, not on each other's existence. One is a phone call where you must have the number. The other is a radio broadcast where you only need the frequency. Anyone can tune in. Anyone can transmit. The transmitter and receiver don't know each other. The frequency is the interface.

## Queue groups: scale without coordination

REST scaling requires a load balancer. The load balancer must know the addresses of all instances. Health checks must be configured. The load balancer is a single point of failure. It is also a single point of configuration drift. Three teams, three load balancers, three different configurations, three different failure modes.

NATS queue groups do this with no additional infrastructure:

```go
// Three instances of inventory service, all in the same queue group
nc.QueueSubscribe("inventory.reserve", "inventory-workers", func(m *nats.Msg) {
    // Only one instance receives each message
    processReservation(m)
})
```

That's it. No load balancer. No health checks. No service discovery. NATS distributes messages across the queue group. If an instance dies, NATS stops sending it messages. If a new instance starts, NATS includes it. The subject name didn't change. The publisher didn't change. The queue group is an implementation detail hidden behind the subject. This is Parnas's information hiding applied to deployment topology. The subject is the stable interface. The instances are the volatile implementation. Adding instances doesn't change the interface. Removing instances doesn't change the interface. The publisher never knows. The publisher never needs to know.

## Temporal decoupling: survive the downtime

REST fails when the callee is unavailable. The caller must retry, back off, circuit-break, or fail. Each of these is hand-rolled in every service. Each retry loop is slightly different. Each circuit breaker has slightly different thresholds. The system's reliability is the sum of individually implemented retry strategies, none of which were tested together.

NATS provides temporal decoupling out of the box. If the subscriber is offline, the message can be queued in JetStream:

```go
// Publisher: fire and forget with persistence
js.Publish(ctx, "orders.requested", event)

// Subscriber: receives when ready, even hours later
consumer.Consume(func(msg jetstream.Msg) {
    processOrder(msg)
    msg.Ack()
})
```

The publisher publishes. The subscriber consumes when it's online. If no subscriber is online, the message waits. If the subscriber crashes mid-processing, the message is redelivered. The publisher doesn't retry. The publisher doesn't circuit-break. The publisher publishes and moves on. The infrastructure handles the reliability. The services handle the business logic. This is the separation of concerns that REST promises and never delivers.

## What coupling actually means

Henney defined coupling as shared knowledge between components. REST maximizes shared knowledge. The caller knows the callee's address, API, availability, response format, error codes, authentication method, rate limits. Each of these is knowledge. Each is coupling. When the callee changes any of them, the caller must change or break. The deployment is independent. The knowledge is shared. The coupling is real.

NATS minimizes shared knowledge. The publisher knows the subject name and the message schema. The subscriber knows the subject name and the message schema. Neither knows the other exists. Neither knows how many of the other exist. Neither knows where the other is deployed. Neither knows whether the other is currently online. The subject name and the schema: that is the total shared knowledge. Everything else is hidden. This is Parnas's criterion, applied to inter-service communication. Hide the volatile decisions — deployment location, instance count, availability, internal API changes — behind a stable interface. The subject is the stable interface. The service is the volatile implementation. The interface doesn't change when the implementation scales, moves, or restarts. The coupling is minimized. The modularity is real.

## When REST makes sense

REST is not universally wrong. It is appropriate when:

**The client needs an immediate response.** Request-reply over NATS works for this, but REST is simpler for simple cases. If one service genuinely needs a synchronous answer from another, and the dependency is stable, REST is fine. The issue is not REST. The issue is using REST as the default for everything.

**The API is externally consumed.** External clients — mobile apps, web browsers, third-party integrations — don't speak NATS. They speak HTTP. The edge of the system is REST. The interior of the system should not be. Expose REST at the boundary. Use NATS internally. The boundary is where HTTP belongs. The interior is where it doesn't.

**The dependency is genuinely stable.** If Service A will always need Service B, and B's API will never change, and B will never move, and B will never be replaced — REST is fine. This is a rare set of conditions. Most service dependencies are less stable than architects assume. The REST interface hardcodes assumptions that will be invalidated. The NATS subject survives the invalidation because it doesn't encode the assumptions.

For everything else — internal service communication, event-driven workflows, systems that must scale, systems that must survive partial failure — NATS pub/sub is not just better. It is what REST pretends to be: modular, decoupled, and resilient to change. The difference is not in the syntax. It is in the architecture of knowledge. REST shares knowledge by default. NATS hides knowledge by default. Modularity lives in what you don't know about the other service. NATS lets you know less. That is the definition of better modularity.

---

**References:**
- [NATS Go Client](https://github.com/nats-io/nats.go) — official Go client with JetStream v2 API
- [NATS by Example](https://natsbyexample.com/) — runnable cross-client examples
- Related posts: [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Henney's Microservices](https://blog.hackspree.com/#kevlin-henney), [Software dark factories](https://blog.hackspree.com/#software-dark-factories)

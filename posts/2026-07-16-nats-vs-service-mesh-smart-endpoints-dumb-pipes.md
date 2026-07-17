---
title: "Smart Endpoints, Dumb Pipes: Why NATS Replaces the Service Mesh"
date: 2026-07-16
slug: nats-vs-service-mesh-smart-endpoints-dumb-pipes
summary: The service mesh was supposed to solve microservices communication. Instead, it became the problem — many teams spend more time managing Istio than building features. "Smart endpoints and dumb pipes" is not a slogan but the third generation of an idea that has been winning arguments since 1984, and two diagrams from the NATS blog make the whole case visually — a mesh architecture and the same system on NATS. This post traces the principle's lineage, reads the two diagrams, defines what "dumb" actually means, and walks capability by capability (routing, load balancing, observability, security, error handling) through how NATS provides what the mesh promised, without the mesh.
tags: nats, microservices, service-mesh, istio, architecture, distributed-systems, history
---

When microservices hit a certain scale — Chanaka Fernando puts the threshold around 25 services — three problems become unavoidable: inter-service communication, observability, and failure handling. The conventional prescription for the last five years has been a service mesh. And for the last five years, teams have been learning the hard way that the cure is often worse than the disease.

Fernando, in his [NATS blog post on building scalable microservices](https://nats.io/blog/building-scalable-microservices-with-nats/), makes an observation that should stop every architect mid-scroll: "There are more teams struggling to manage microservices with Service Meshes than who succeeded with it."

This is not a hot take. It's a field report. And the alternative he proposes is not a better service mesh. It's a return to first principles — a principle old enough that his post needs only two diagrams to make the case. This post reads both.

## The forgotten principle, and where it comes from

The original microservices formulation — the one in [Lewis and Fowler's 2014 article](https://martinfowler.com/articles/microservices.html), the one that actually shipped — had a principle: **smart endpoints and dumb pipes**. The intelligence lives in the services. The pipe between them is simple. It routes messages. It doesn't transform, authenticate per-request, rate-limit, circuit-break, retry, or observe. Those are endpoint responsibilities. The pipe is reliable, fast, and boring.

The phrase is from 2014, but the idea is the third generation of an argument that has been winning since before microservices existed:

- **1984, networks.** Saltzer, Reed, and Clark's [End-to-End Arguments in System Design](https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf) made the foundational case: functions like reliability, ordering, and correctness checks can only be *completely* implemented at the endpoints, so implementing them inside the network is at best an optimization and at worst wasted complexity. TCP/IP is this argument deployed at planetary scale — the internet outlived every "intelligent network" the telecom industry proposed because IP routers do almost nothing.
- **1978, operating systems.** Doug McIlroy's Unix pipes, described in the [Bell System Technical Journal](https://archive.org/details/bstj57-6-1899), connect programs through a byte stream that has no opinions. The intelligence is in `grep` and `sort`, never in `|`.
- **2014, distributed applications.** Lewis and Fowler's target was the Enterprise Service Bus — Jim Webber's "Erroneous Spaghetti Box" — which put transformation, routing logic, and business rules *into the pipe*. Successful microservice teams did the opposite.

Each generation relearned the same result: **systems compose and scale when the intermediary is boring.** Then the industry forgot it again — twice. First with the ESB, which the microservices movement explicitly rebelled against. Then with the service mesh, which reintroduced per-request intelligence in the middle of every call, this time as a sidecar instead of a broker. Fernando's observation is blunt: most people "seem to forget this idea when designing microservices platforms."

NATS is a return to the original principle. It is a dumb pipe in the best sense: fast, reliable, simple. It routes messages by subject. It doesn't inspect payloads. It doesn't enforce retry policies. It doesn't generate traces unless you opt in. It does one thing — move messages from publishers to subscribers — and it does it at millions per second.

## Figure 1: the service mesh promise

The service mesh pitch is elegant. Microservices need to communicate. That communication needs routing, retries, timeouts, circuit breaking, load balancing, observability, and security. Rather than embedding that logic in every service, extract it into a sidecar proxy. Deploy Envoy alongside every pod. Control it centrally with Istio. The data plane moves the bytes. The control plane moves the config. The service code stays clean.

![Service mesh architecture: services with sidecar proxies forming a data plane, managed by a separate control plane](images/scalable-microservices-nats-1.png)

*Service mesh architecture. Figure from [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) by Chanaka Fernando (nats.io).*

Read this diagram the way you'd review a design doc, and count what's on the request path that isn't your code. Every service has a proxy bolted to it. Every call from service A to service B traverses A's sidecar and B's sidecar — two full L7 proxies per hop, each parsing HTTP, evaluating routing rules, checking policy, minting telemetry. Above them sits the control plane, a separate distributed system whose job is to configure the first distributed system: service discovery, certificate authority, routing configuration, policy distribution.

Now apply the end-to-end test from 1984: which of these functions is *completed* in the middle? Retries in the sidecar can't know whether a request is safe to retry — that's business knowledge. Circuit breaking in the sidecar can't know which failures matter — that's business knowledge. Even mTLS between sidecars secures the hop, not the request. The mesh implements, in the pipe, partial versions of functions the endpoints must implement anyway to be correct. That is precisely the redundancy Saltzer, Reed, and Clark warned about — except now it ships as a fleet of Envoys and a quarterly upgrade treadmill.

If you are Google, running hundreds of thousands of services on a unified infrastructure, this can still make sense. The operational cost of the mesh is amortized across an enormous fleet. The marginal cost of adding one more service is near zero. The control plane team is a separate organization from the service teams. The abstraction earns its keep.

## The service mesh reality

If you are not Google — if you are a team of 15 engineers running 40 services on a single Kubernetes cluster — the economics invert. The mesh is now a significant fraction of your operational surface area. Every upgrade is a negotiation with the mesh. Every debugging session starts with "is it the mesh?" Every new hire needs to understand Envoy config, Istio CRDs, and why there's a proxy between two services in the same namespace.

The specific failure modes are well-documented:

**Operational complexity.** Istio alone has dozens of CRDs — VirtualService, DestinationRule, Gateway, ServiceEntry, WorkloadEntry, PeerAuthentication, RequestAuthentication, AuthorizationPolicy, EnvoyFilter, WasmPlugin, Telemetry, ProxyConfig. Each represents a knob you didn't ask for but now must understand because it defaults in ways that affect your traffic.

**Debugging indirection.** When Service A calls Service B and gets a 503, the error could be in A's code, A's Envoy config, B's Envoy config, B's code, the Istio ingress gateway, the network policy, or mTLS certificate rotation. The sidecar adds two more places for things to go wrong, and the control plane adds configuration that can be wrong in ways the data plane silently enforces.

**Upgrade coupling.** Istio releases roughly quarterly. Each release deprecates APIs. The upgrade from 1.12 to 1.13 might change the default mesh policy. The upgrade from 1.18 to 1.20 might remove an API group you depended on. Your services are decoupled from each other but coupled to the mesh's release cycle.

**Performance overhead.** Every request traverses two Envoy proxies — the caller's sidecar and the callee's sidecar. Each adds latency. Under load, the sidecars consume CPU and memory proportional to the number of concurrent connections. The mesh is not free. You pay for it on every request, even between services in the same pod.

**Knowledge requirements.** A developer writing a service that calls another service needs to understand: the service's language, the service's business logic, Kubernetes networking, Envoy configuration, Istio routing rules, mTLS certificate management, and the mesh's observability stack. The mesh was supposed to hide complexity from developers. It created a new kind of complexity and made it everyone's problem.

Even Istio's own architecture has been in flux — the project has undergone significant rearchitecting, moving from a multi-component control plane (Pilot, Citadel, Galley, Mixer) to istiod (a monolith that bundles them all) and back toward separating concerns. When the tooling designed to solve complexity keeps changing its architecture to manage its own complexity, something is wrong.

The diagram shows an architecture where the pipe got smart. And smart pipes have a property the diagram can't show but every operator knows: when something breaks, the number of places to look is the number of boxes. Here, most of the boxes aren't yours.

## Figure 2: the same system on NATS

![Microservices communicating through a central NATS cluster instead of point-to-point connections with sidecars](images/scalable-microservices-nats-2.png)

*Inter-service communication with NATS. Figure from [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) by Chanaka Fernando (nats.io).*

Same services. The sidecars are gone. The control plane is gone. In the middle there is one thing: a NATS cluster that routes messages by subject and does nothing else to them.

What the second diagram deletes is instructive, but what it *keeps* is the real point. The mesh existed to provide discovery, load balancing, routing, and decoupling. Those needs don't disappear — they get satisfied structurally instead of by middleware:

- **Discovery** collapses into subscription. A service that subscribes to `orders.created` is discovered, by definition, by anything that publishes to that subject. There is no registry to sync because interest *is* the registry.
- **Load balancing** collapses into [queue groups](https://docs.nats.io/nats-concepts/core-nats/queue): N subscribers in the same group split the subject's traffic, competing-consumer style, with no balancer tier and no health-check config.
- **Routing** collapses into the subject namespace. `payments.processed.visa` is both the address and the meaning; wildcards give you the routing table you'd otherwise write as CRDs.
- **Decoupling** is the default rather than an aspiration: publishers don't hold connections to consumers, don't know their count, and don't fail when one of them redeploys.

The pipe stayed dumb — payload-agnostic, business-logic-free — and the coordination problems got absorbed into the *shape* of the system rather than into configuration.

## What "dumb" actually means

The principle is routinely misread as "use no infrastructure" or "the broker must be featureless." Neither is right, and the second diagram isn't claiming it. NATS clusters, does TLS and decentralized auth, and with JetStream will happily persist and replay streams. Dumb doesn't mean minimal. It means the pipe operates below the application's semantics:

1. **Payload-agnostic.** The pipe never parses your message body. The moment the middle needs to understand your schema, it's coupled to every service's release cycle.
2. **No business branching.** "If the order is over $10k, route to fraud review" is an endpoint decision. The pipe routes on the address (the subject), never on the content.
3. **Failure semantics live at the edge.** The pipe may redeliver; only the endpoint knows what a retry *means*. Idempotency, compensation, and circuit breaking are written where the business knowledge is.

Here's the whole principle as a code review question: **when a business rule changes, does the diff land in a service or in the middleware's config?** If routing rules, retry policies, and transformations keep accumulating in the middle, your endpoints are getting dumber and your pipe is getting smarter — and you are rebuilding the ESB, whatever the box is labeled.

A smart endpoint on a dumb pipe fits in a screenful of Go:

```go
// Smart endpoint: owns validation, idempotency, and reply semantics.
// The pipe's entire contribution is the subject and the bytes.
nc, _ := nats.Connect(nats.DefaultURL)

// Queue group = load balancing with zero middleware.
nc.QueueSubscribe("orders.create", "order-workers", func(m *nats.Msg) {
    var o Order
    if err := json.Unmarshal(m.Data, &o); err != nil {
        m.Respond(errReply("bad request")) // endpoint decides error semantics
        return
    }
    if seen(o.ID) {                        // endpoint owns idempotency
        m.Respond(okReply(o.ID))
        return
    }
    m.Respond(process(o))                  // endpoint owns the business
})
```

Everything interesting happens inside the handler. The [NATS services framework](https://docs.nats.io/using-nats/developer/services) adds discovery and per-endpoint stats to exactly this pattern — as a library in the endpoint, which is where the principle says that intelligence belongs.

## What the service mesh provides, and how NATS provides it differently

Let's go capability by capability.

### Routing

**Service mesh:** The mesh routes HTTP requests by hostname and path. A VirtualService defines match rules: `prefix: /api/users` routes to the `user-service` subset `v2`. A DestinationRule defines load balancing policy: `LEAST_REQUEST` with outlier detection. The routing is explicit, granular, and managed by the control plane. Changes propagate through xDS to every Envoy in the mesh.

**NATS:** Routing is by subject. A service subscribes to `orders.created`. Another publishes to `orders.created`. The NATS server routes the message. No routing rules to configure. No destination rules to manage. No control plane to synchronize. The subject *is* the route. Wildcards (`orders.*`, `orders.>`) handle categories of interest. The routing topology is implicit in the subject namespace — change the namespace, and you change the routing. No CRDs.

The key insight: subject-based routing is self-describing. The subject `payment.processed.visa.us-east` tells you what it is, where it came from, and what it's about. An HTTP path `/api/v2/payments/processed?provider=visa&region=us-east` requires the consumer to parse query parameters to understand the same thing. Subjects carry semantics. URLs carry hierarchy. One is designed for machine routing. The other is designed for human navigation.

### Load balancing

**Service mesh:** Envoy distributes requests across service instances based on configured policy — round-robin, least request, ring hash, random, maglev. Health checks determine pool membership. Outlier detection ejects unhealthy instances. Circuit breaking caps pending requests. All of this is configured in DestinationRules.

**NATS:** Queue groups. Multiple subscribers join the same queue group. NATS distributes messages round-robin across group members. If an instance disconnects, NATS stops routing to it. If a new instance connects, NATS includes it automatically:

```go
// Three instances, one queue group, no configuration
nc.QueueSubscribe("orders.process", "order-workers", func(m *nats.Msg) {
    processOrder(m)
})
```

No health checks. No pool configuration. No outlier detection thresholds. The NATS server knows which clients are connected. It distributes messages across the connected set. Disconnection is the health check. Subscription is the pool membership. The protocol is the configuration.

For request-reply patterns, queue groups provide natural load balancing — the reply goes to one instance, not all of them. The caller doesn't know how many instances exist. It sends a request and gets one reply. The distribution is transparent, efficient, and zero-config.

### Observability

**Service mesh:** The mesh generates telemetry from every proxy — request latency, response codes, connection counts, retry rates. This is powerful but noisy. Every sidecar generates metrics. Every request generates spans. The signal-to-noise ratio depends on careful configuration of sampling rates, aggregation, and retention. Most teams end up collecting everything and using almost none of it.

**NATS:** The server exports metrics and can be monitored with Prometheus. Trace context is propagated through message headers — but it's opt-in. You set `OTEL_EXPORTER_OTLP_ENDPOINT` and traces flow. You don't set it, and there's zero overhead. The observability surface is smaller because the pipe is simpler. Fewer moving parts generate fewer metrics. The metrics that exist are more meaningful because they're not buried in proxy-sidecar noise.

The NATS approach to distributed tracing is particularly clean. The hub and spoke inject and extract W3C trace context through NATS message headers:

```go
func InjectTraceContext(ctx context.Context, msg *nats.Msg) {
    otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(msg.Header))
}

func ExtractTraceContext(msg *nats.Msg) context.Context {
    return otel.GetTextMapPropagator().Extract(context.Background(), propagation.HeaderCarrier(msg.Header))
}
```

No sidecar. No proxy. No daemonset. The trace context travels with the message through standard propagation. The application controls when and how it observes. The infrastructure doesn't observe on your behalf. Smart endpoints, dumb pipes.

### Security

**Service mesh:** mTLS everywhere. Citadel (or istiod) issues certificates to every sidecar. Every request is authenticated at the transport layer. AuthorizationPolicy CRDs define who can call whom. SPIFFE identities tie services to cryptographic identities. The security is strong but complex — certificate rotation, trust domain configuration, and policy debugging are nontrivial.

**NATS:** Authentication and authorization are built into the protocol. No overlay. No sidecar. The connection is authenticated. The subjects are authorized. JWT-based user credentials carry scoped permissions:

```go
claims := jwt.NewUserClaims(userPub)
claims.Pub.Allow = []string{
    "heartbeat." + id,
    "joke.response." + id + ".>",
}
claims.Sub.Allow = []string{
    "_INBOX.>",
    "joke.request." + id,
}
```

The NATS server enforces these at connection time and on every publish/subscribe. If a client doesn't have permission to publish on a subject, the server rejects the message before it touches any subscriber. The security boundary is the subject, not the network. Encryption can be enabled at the transport layer, but authorization is at the application layer — where it belongs.

The decentralized model matters for scale. In a service mesh, the API gateway is often the centralized choke point for north-south traffic, while mTLS secures east-west. In NATS, the account is the security boundary, and every connection authenticates directly to the server — no gateway, no sidecar, no proxy.

Note that this doesn't violate the dumb-pipe principle: connection-time authentication and subject authorization operate below the application's semantics, like TLS on a socket. The pipe still never parses a payload or evaluates a business rule.

### Error handling

**Service mesh:** The mesh handles retries, timeouts, and circuit breaking at the proxy level. This keeps retry logic out of application code but creates a new problem: the application doesn't know what the proxy is doing. A retry that succeeds on the third attempt looks like a single fast request to the application. A circuit breaker that opens looks like a 503. The application loses visibility into the transport's behavior.

**NATS:** Error handling uses the messaging primitives directly. For fire-and-forget patterns, JetStream provides persistence and redelivery. The message waits in the stream until a consumer acknowledges it. If the consumer crashes, the message is redelivered. The publisher doesn't retry. The subscriber doesn't need to be online. The infrastructure handles reliability — the service handles business logic.

For request-reply, the caller sets a timeout and gets either a response or a timeout error. That's the entire error model. No retry storms. No circuit-breaker cascades. No proxy injecting 503s into a healthy connection. The error is either "here's the reply" or "nobody answered."

## The architectural inversion

A service mesh is infrastructure-heavy and application-light. NATS is infrastructure-light and application-heavy. Both solve the same problems. They differ in where the complexity lives.

In the service mesh model, you buy complexity once (the mesh) and hope it pays off across many services. In the NATS model, you avoid the complexity entirely and let each service handle its own communication through a simple, reliable pipe.

The service mesh bet is that centralizing communication logic in a proxy layer is worth the operational cost. The NATS bet is that the operational cost of the proxy layer exceeds the benefit — and that a messaging system with the right primitives (subjects, queue groups, JetStream, JWTs) eliminates the need for the proxy layer entirely.

The NATS bet is winning. Not because NATS is a better service mesh. Because NATS is a better architecture for the problem the service mesh was trying to solve. Figure 1 and Figure 2 provide the same capabilities; the difference is where the intelligence sits. In Figure 1 it's smeared across sidecars and a control plane that nobody's service owns. In Figure 2 it's concentrated in the endpoints, connected by a pipe too simple to be wrong in interesting ways. Forty years of systems history — end-to-end arguments, Unix pipes, the ESB's rise and fall — keep returning the same verdict on that choice.

## One caveat the diagrams earn

Location transparency has a known failure mode: pretending the network isn't there. Waldo, Wyant, Wollrath, and Kendall's 1994 paper [A Note on Distributed Computing](https://scholar.harvard.edu/files/waldo/files/waldo-94.pdf) demolished systems that hid remoteness behind local-looking interfaces. A dumb pipe doesn't repeal that — messages still get lost, reordered, and delivered twice. The difference is honest accounting: NATS's core delivery guarantee is at-most-once, stated plainly, and anything stronger (JetStream acknowledgments, exactly-once processing windows) is an explicit endpoint opt-in rather than a middlebox promise. Dumb pipes don't make distribution easy. They make it *visible*, which is the only foundation correctness can be built on.

## When a service mesh still makes sense

NATS doesn't replace every service mesh use case. It doesn't do traffic splitting for canary deployments (e.g., "send 10% of traffic to the v2 deployment"). It doesn't do fault injection for chaos engineering. It doesn't do request-level header manipulation. These are HTTP-specific features that a messaging system shouldn't replicate.

But ask yourself: how many of those features do you actually use? And of those, how many are compensating for the fact that your services are coupled at the HTTP layer? If you're using traffic splitting to test a new service version, would you need it if services communicated through subjects that don't change? If you're using fault injection, would you need it if your services were already designed for asynchronous delivery with redelivery and idempotency?

The service mesh solves problems created by the communication model it layers over. NATS changes the communication model. Many of the problems the mesh solves become non-problems.

## What to read first

The argument for smart-endpoints-dumb-pipes has been around since the beginning of microservices — and, under other names, since 1984. What's new is the evidence that the service mesh failed to deliver on its promise for most organizations, and that a simpler alternative — a proper messaging system — works at scale with a fraction of the operational burden. Fernando's [original post](https://nats.io/blog/building-scalable-microservices-with-nats/) makes this case clearly. His book, *Designing Microservices Platforms with NATS*, extends it with practical examples.

The next post in this series examines the concrete evidence: Sophotech's migration from RabbitMQ to NATS, where p99 latency dropped 3.75x and ops time fell from several hours a week to under one. The theory is sound. The numbers back it up.

## References

1. Chanaka Fernando, [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) — source of both figures.
2. James Lewis and Martin Fowler, [Microservices](https://martinfowler.com/articles/microservices.html) (2014) — origin of "smart endpoints and dumb pipes."
3. J.H. Saltzer, D.P. Reed, D.D. Clark, [End-to-End Arguments in System Design](https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf) (1984).
4. M.D. McIlroy et al., [UNIX Time-Sharing System: Foreword](https://archive.org/details/bstj57-6-1899), Bell System Technical Journal (1978) — the pipe philosophy.
5. Jim Waldo et al., [A Note on Distributed Computing](https://scholar.harvard.edu/files/waldo/files/waldo-94.pdf) (1994) — the limits of location transparency.
6. Gregor Hohpe and Bobby Woolf, [Enterprise Integration Patterns](https://www.enterpriseintegrationpatterns.com/) (2003) — the messaging-pattern vocabulary the ESB era misapplied.
7. [NATS documentation: Subject-Based Messaging](https://docs.nats.io/nats-concepts/subjects).
8. [NATS documentation: Queue Groups](https://docs.nats.io/nats-concepts/core-nats/queue).
9. [NATS documentation: Services Framework](https://docs.nats.io/using-nats/developer/services).
10. [NATS documentation: Security](https://docs.nats.io/running-a-nats-service/configuration/securing_nats).
11. [NATS documentation: Clustering](https://docs.nats.io/running-a-nats-service/configuration/clustering).
12. [Istio documentation: Architecture](https://istio.io/latest/docs/ops/deployment/architecture/) — the mesh design on its own terms.
13. Chanaka Fernando, *Designing Microservices Platforms with NATS* (Packt, 2021).

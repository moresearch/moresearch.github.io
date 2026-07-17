---
title: "Smart Endpoints and Dumb Pipes: The Whole Argument in Two Diagrams"
date: 2026-07-16
slug: nats-smart-endpoints-dumb-pipes-two-diagrams
summary: '"Smart endpoints and dumb pipes" is the most quoted and least obeyed principle in microservices. This post traces where the phrase actually comes from — the end-to-end argument of 1984, Unix pipes, the ESB backlash — and then reads two diagrams from the NATS blog side by side: a service mesh architecture and the same system on NATS. The difference between them is not a vendor preference. It is a decision about where intelligence is allowed to live.'
tags: nats, microservices, architecture, distributed-systems, history
---

Every architecture argument eventually compresses into a picture. For "smart endpoints and dumb pipes," two pictures from Chanaka Fernando's post on the NATS blog, [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/), do the whole job. The first shows a service mesh. The second shows the same microservices talking over NATS. Everything this post argues is visible in the space between them.

But diagrams only persuade if you know what to look for. So before reading them, it's worth establishing where the principle comes from — because "smart endpoints and dumb pipes" is not a 2014 slogan. It's the third generation of an idea that has been winning arguments since 1984.

## Where the phrase comes from

The phrase itself is from James Lewis and Martin Fowler's 2014 article [Microservices](https://martinfowler.com/articles/microservices.html), where it appears as one of nine characteristics of the architecture. Their target was explicit: the Enterprise Service Bus. ESBs of the 2000s put message transformation, routing logic, business rules, and protocol mediation *into the pipe*. Lewis and Fowler observed that successful microservice teams did the opposite — they kept application intelligence in the services and used infrastructure that was "dumb (dumb as in acts as a message router only)."

The idea is older than the phrase. Saltzer, Reed, and Clark's 1984 paper [End-to-End Arguments in System Design](https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf) made the foundational case: functions like reliability, ordering, and correctness checks can only be *completely* implemented at the endpoints of a communication system, so implementing them inside the network is at best an optimization and at worst wasted complexity. TCP/IP is the end-to-end argument deployed at planetary scale — the reason the internet outlived every "intelligent network" the telecom industry proposed is that IP routers do almost nothing.

Unix got there independently. Doug McIlroy's pipes, described in the [Bell System Technical Journal in 1978](https://archive.org/details/bstj57-6-1899), connect programs through a byte stream that has no opinions. The intelligence is in `grep` and `sort`, never in `|`. Kernighan and Pike's *The UNIX Programming Environment* canonized it: expect the output of every program to become the input of another, and keep the connective tissue trivial.

So the lineage runs: end-to-end argument (networks) → pipes (operating systems) → smart endpoints and dumb pipes (distributed applications). Each generation relearned the same result: **systems compose and scale when the intermediary is boring.**

Then the industry forgot it again — twice. First with the ESB, which the microservices movement explicitly rebelled against. Then with the service mesh, which reintroduced per-request intelligence in the middle of every call, this time as a sidecar instead of a broker. Fernando's observation in the NATS post is blunt: most people "seem to forget this idea when designing microservices platforms."

Which brings us to the diagrams.

## Figure 1: the service mesh

![Service mesh architecture: services with sidecar proxies forming a data plane, managed by a separate control plane](images/scalable-microservices-nats-1.png)

*Service mesh architecture. Figure from [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) by Chanaka Fernando (nats.io).*

Read this diagram the way you'd review a design doc, and count what's on the request path that isn't your code.

Every service has a proxy bolted to it. The proxies form the *data plane*: every call from service A to service B traverses A's sidecar and B's sidecar — two full L7 proxies per hop, each parsing HTTP, evaluating routing rules, checking policy, minting telemetry. Above them sits the *control plane*, a separate distributed system whose job is to configure the first distributed system: service discovery, certificate authority, routing configuration, policy distribution.

Now apply the end-to-end test from 1984: which of these functions is *completed* in the middle? Retries in the sidecar can't know whether a request is safe to retry — that's business knowledge. Circuit breaking in the sidecar can't know which failures matter — that's business knowledge. Even mTLS between sidecars secures the hop, not the request; end-to-end authenticity still needs endpoint participation. The mesh implements, in the pipe, partial versions of functions that the endpoints must implement anyway to be correct. That is precisely the redundancy Saltzer, Reed, and Clark warned about — except now it ships as a fleet of Envoys and a quarterly upgrade treadmill.

The diagram shows an architecture where the pipe got smart. And smart pipes have a property the diagram can't show but every operator knows: when something breaks, the number of places to look is the number of boxes. Here, most of the boxes aren't yours.

## Figure 2: the same system on NATS

![Microservices communicating through a central NATS cluster instead of point-to-point connections with sidecars](images/scalable-microservices-nats-2.png)

*Inter-service communication with NATS. Figure from [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) by Chanaka Fernando (nats.io).*

Same services. The sidecars are gone. The control plane is gone. In the middle there is one thing: a NATS cluster that routes messages by subject and does nothing else to them.

What the second diagram deletes is instructive, but what it *keeps* is the real point. The mesh existed to provide discovery, load balancing, routing, and decoupling. Those needs don't disappear — they get satisfied structurally instead of by middleware:

- **Discovery** collapses into subscription. A service that subscribes to `orders.created` is discovered, by definition, by anything that publishes to that subject. There is no registry to sync because interest *is* the registry.
- **Load balancing** collapses into [queue groups](https://docs.nats.io/nats-concepts/core-nats/queue): N subscribers in the same group split the subject's traffic, competing-consumer style, with no balancer tier and no health-check config — a subscriber that dies simply stops competing.
- **Routing** collapses into the subject namespace. `payments.processed.visa` is both the address and the meaning; wildcards give you the routing table you'd otherwise write as CRDs.
- **Decoupling** is the default rather than an aspiration: publishers don't hold connections to consumers, don't know their count, and don't fail when one of them redeploys.

The pipe stayed dumb — payload-agnostic, business-logic-free — and the coordination problems got absorbed into the *shape* of the system rather than into configuration.

## What "dumb" actually means

The principle is routinely misread as "use no infrastructure" or "the broker must be featureless." Neither is right, and the second diagram isn't claiming it. NATS clusters, does TLS and [decentralized auth](https://docs.nats.io/running-a-nats-service/configuration/securing_nats/auth_intro/jwt), and with JetStream will happily persist and replay streams. Dumb doesn't mean minimal. It means the pipe operates below the application's semantics:

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

## One caveat the diagrams earn

Location transparency has a known failure mode: pretending the network isn't there. Waldo, Wyant, Wollrath, and Kendall's 1994 paper [A Note on Distributed Computing](https://scholar.harvard.edu/files/waldo/files/waldo-94.pdf) demolished systems that hid remoteness behind local-looking interfaces. A dumb pipe doesn't repeal that — messages still get lost, reordered, and delivered twice. The difference is honest accounting: NATS's core delivery guarantee is at-most-once, stated plainly, and anything stronger (JetStream acknowledgments, exactly-once processing windows) is an explicit endpoint opt-in rather than a middlebox promise. Dumb pipes don't make distribution easy. They make it *visible*, which is the only foundation correctness can be built on.

## The two diagrams, restated

Figure 1 and Figure 2 provide the same capabilities. The difference is where the intelligence sits: in Figure 1 it's smeared across sidecars and a control plane that nobody's service owns; in Figure 2 it's concentrated in the endpoints, connected by a pipe too simple to be wrong in interesting ways. Forty years of systems history — end-to-end arguments, Unix pipes, the ESB's rise and fall — keep returning the same verdict on that choice.

For the operational head-to-head — capability by capability against Istio, and the honest cases where a mesh still wins — see the companion post: [Smart Endpoints, Dumb Pipes: Why NATS Replaces the Service Mesh](#nats-vs-service-mesh-smart-endpoints-dumb-pipes).

## References

1. Chanaka Fernando, [Building Scalable Microservices with NATS](https://nats.io/blog/building-scalable-microservices-with-nats/) — source of both figures.
2. James Lewis and Martin Fowler, [Microservices](https://martinfowler.com/articles/microservices.html) (2014) — origin of "smart endpoints and dumb pipes."
3. J.H. Saltzer, D.P. Reed, D.D. Clark, [End-to-End Arguments in System Design](https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf) (1984).
4. M.D. McIlroy et al., [UNIX Time-Sharing System: Foreword](https://archive.org/details/bstj57-6-1899), Bell System Technical Journal (1978) — the pipe philosophy.
5. Brian Kernighan and Rob Pike, *The UNIX Programming Environment* (1984).
6. Jim Waldo et al., [A Note on Distributed Computing](https://scholar.harvard.edu/files/waldo/files/waldo-94.pdf) (1994) — the limits of location transparency.
7. Gregor Hohpe and Bobby Woolf, [Enterprise Integration Patterns](https://www.enterpriseintegrationpatterns.com/) (2003) — the messaging-pattern vocabulary the ESB era misapplied.
8. [NATS documentation: Subject-Based Messaging](https://docs.nats.io/nats-concepts/subjects).
9. [NATS documentation: Queue Groups](https://docs.nats.io/nats-concepts/core-nats/queue).
10. [NATS documentation: Services Framework](https://docs.nats.io/using-nats/developer/services).
11. [Istio documentation: Architecture](https://istio.io/latest/docs/ops/deployment/architecture/) — the mesh design on its own terms.
12. Chanaka Fernando, *Designing Microservices Platforms with NATS* (Packt, 2021).

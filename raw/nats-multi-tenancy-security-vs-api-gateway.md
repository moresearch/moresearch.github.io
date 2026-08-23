---
title: "NATS Multi-Tenancy and Security: JWT Auth vs the API Gateway"
date: 2026-07-16
slug: nats-multi-tenancy-security-vs-api-gateway
summary: The API gateway is the traditional choke point for microservices security — every request passes through it, every auth decision is centralized, and every service behind it trusts the gateway implicitly. NATS inverts this with decentralized JWT-based authentication and authorization: every connection carries its own credentials, the server enforces subject-level permissions on every publish and subscribe, and accounts provide tenant isolation without a gateway. Using the Vitrifi case study as a reference, this post explains how NATS security works, how it compares to the gateway model, and when a gateway still makes sense.
tags: nats, security, jwt, multi-tenancy, api-gateway, authentication, authorization, distributed-systems
---

The API gateway is the standard answer to microservices security. External traffic hits the gateway. The gateway authenticates, authorizes, rate-limits, and routes. Services behind the gateway trust that incoming requests have been vetted. The network enforces the boundary: everything outside the gateway is untrusted, everything inside is trusted. The security model is perimeter-based.

NATS takes the opposite approach. Every connection authenticates directly to the NATS server — no proxy, no sidecar, no gateway. Authorization is enforced at the subject level on every publish and subscribe. Multi-tenancy is implemented through accounts that provide cryptographic isolation between tenants sharing the same NATS infrastructure. The security model is credential-based.

The Vitrifi case study on the [NATS blog](https://nats.io/blog/nats-vitrifi-secret-weapon/) demonstrates this model in production: a SaaS workflow automation platform using NATS accounts with JWT-based authentication to ensure "each tenant's data and processing remain entirely separate while sharing infrastructure." No API gateway per tenant. No network segmentation per tenant. Cryptographic isolation at the messaging layer.

## The API gateway model

In the traditional model, security is enforced at the edge:

```
External Client → [API Gateway] → Service A → Service B → Service C
                      ↑
            Auth, TLS, rate limiting,
            routing, API keys, JWT validation
```

The gateway does everything: terminates TLS, validates API keys or JWTs, enforces rate limits, applies routing rules, transforms requests, logs access. Services behind the gateway trust that the gateway has done its job. Service A doesn't authenticate Service B because both are inside the perimeter. The perimeter is the security boundary.

This model has real advantages:

**Centralized policy.** Security rules are in one place. Adding a new rate limit, rotating an API key, or changing an authentication provider happens at the gateway. Services don't change.

**Mature ecosystem.** API gateways are a mature product category. Kong, Apigee, AWS API Gateway, Envoy-based gateways — they all provide roughly the same feature set with battle-tested implementations. The patterns are documented. The operations are understood.

**Separation of concerns.** Security teams own the gateway. Development teams own the services. The security boundary is also an organizational boundary. Security doesn't need to understand the services' internals. Developers don't need to implement authentication.

And it has real disadvantages:

**The gateway is a choke point.** Every request passes through it. When it's down, everything is down. When it's slow, everything is slow. High availability requires multiple gateway instances behind a load balancer, which adds another layer of infrastructure.

**The gateway is a single point of policy.** A misconfigured routing rule affects every service it routes to. A bug in the rate limiter affects every client. The blast radius of gateway misconfiguration is the entire system.

**East-west traffic is unprotected.** Traffic between services behind the gateway is implicitly trusted. If Service A is compromised, it can call Service B without additional authentication. The perimeter model assumes the interior is safe. When it isn't — and it eventually isn't — lateral movement is unconstrained.

**The gateway doesn't understand application semantics.** It sees HTTP methods and paths. It doesn't see business operations. A rule that says "Service A can POST to `/api/orders`" doesn't distinguish between "create an order" and "cancel an order." The authorization is coarse-grained because the gateway operates at the protocol layer, not the application layer.

The service mesh partially addresses the east-west problem with mTLS between services, but it adds the operational complexity discussed in the first post of this series. And it still doesn't solve the application-semantics problem — mTLS authenticates the service identity, not the operation.

## The NATS model: decentralized credential-based security

In NATS, there is no gateway. Every client — whether it's an external API service or an internal worker — authenticates directly to the NATS server:

```
Service A ──► NATS Server ──► Service B
   │              │               │
   └── JWT ───────┘               │
   └── Subject permissions ───────┘
```

Authentication happens at connection time. Each connection presents credentials (JWT, nkey, or token). The server validates them against the configured trust chain. Once authenticated, every publish and subscribe is authorized against the user's JWT claims:

```go
claims := jwt.NewUserClaims(userPub)
claims.Pub.Allow = []string{
    "heartbeat.service-a",
    "orders.response.service-a.>",
}
claims.Sub.Allow = []string{
    "_INBOX.>",
    "orders.request.service-a",
}
```

The user can publish on `heartbeat.service-a` and `orders.response.service-a.>`. They cannot publish on `orders.response.service-b.>` or `heartbeat.service-b`. They can subscribe to `_INBOX.>` (required for request-reply) and `orders.request.service-a` (their own requests). They cannot subscribe to `orders.request.service-b`. The NATS server enforces this at the protocol layer. If the client sends a `PUB` on an unauthorized subject, the server rejects it before the message touches any subscriber.

This is Parnas's information hiding applied to security. The user's permissions are scoped to what they need — their own heartbeat, their own requests, their own responses. They don't need to know other services exist. They don't have permission to interact with them. The security boundary is the subject, and the enforcement is cryptographic.

### The trust chain

NATS JWT security uses a three-tier trust chain: operator → account → user.

```
Operator JWT
  └── Account JWT
        ├── User JWT (hub)
        └── User JWT (spoke)
```

The operator is the trust anchor. It signs account JWTs. Account JWTs define the security boundary — JetStream limits, import/export policies, user claim signing keys. User JWTs carry scoped subject permissions and are signed by the account key.

This chain means:

- The operator can revoke an entire account by expiring its JWT
- An account can revoke a user by expiring their JWT
- A user's permissions are bounded by what the account allows — even a compromised account key can't escalate beyond the operator's grant
- No user in Account A can access subjects in Account B unless there's an explicit import/export between the accounts

The trust chain is decentralized. There's no central authentication service that must be online for connections to be established. The NATS server validates JWTs locally against the operator's public key. The JWT itself carries the claims. The server only needs the operator JWT and the account JWT in its resolver. Connections are authenticated pairwise between client and server, not through a centralized gateway.

### Accounts: the multi-tenancy primitive

NATS accounts are the key abstraction for multi-tenancy. An account is a security boundary. Subjects within an account are isolated from subjects in other accounts. Services in Account A cannot see subjects in Account B unless Account B explicitly exports them and Account A explicitly imports them.

This maps directly to SaaS multi-tenancy. Each tenant gets an account. Each tenant's services are users within that account. Tenant A's `orders.created` subject is different from Tenant B's `orders.created` subject — they share a subject name but are isolated by account. No data leaks between tenants because the server won't route messages across accounts without explicit import/export.

Vitrifi's platform uses this model in production: "As a SaaS platform, tenant isolation is achieved through NATS accounts combined with JWT-based authentication, ensuring each tenant's data and processing remain entirely separate while sharing infrastructure."

The operational benefit is significant. You don't need a separate NATS cluster per tenant. You don't need network segmentation (VPCs, subnets, security groups) per tenant. You don't need an API gateway per tenant. Tenants share the same NATS servers, the same network, the same infrastructure — and cannot see each other's traffic because the server enforces isolation at the protocol layer. Adding a tenant is creating a new account and issuing user JWTs. Removing a tenant is expiring the account JWT.

## Leaf nodes: security at the edge

NATS leaf nodes extend the security model to edge deployments without requiring VPNs or network-level trust:

```
                    ┌──────────────┐
Edge Location       │   NATS Hub   │
                    │  (cluster)   │
┌──────────┐        │              │
│ Leaf     │◄──────►│  Account A   │
│ Node     │  Leaf  │  Account B   │
│ (remote) │  conn  │  Account C   │
└──────────┘        └──────────────┘
```

A leaf node is a NATS server running at the edge (a retail store, a factory, an IoT gateway). It connects to the hub cluster over a single outbound connection — no inbound firewall rules required. The leaf authenticates as a NATS client. Local clients connect to the leaf. The leaf exports subjects that the edge can publish and imports subjects the edge can subscribe to.

The security properties are:

- **Local traffic stays local.** Services at the edge communicate through the leaf without traversing the WAN. A leaf node is a full NATS server — it routes messages locally for clients connected to it.
- **Remote traffic is scoped.** The hub only sees subjects the leaf exports. The leaf only receives subjects it imports. The leaf cannot subscribe to subjects it hasn't been granted access to.
- **Authentication is local to the edge.** Edge clients authenticate to the leaf. The leaf authenticates to the hub. Credentials don't leave the edge. The hub trusts the leaf, not the edge clients.
- **No VPN required.** The leaf connection is a single outbound TLS connection. No site-to-site VPN. No network-level trust. The security boundary is the NATS connection, not the network.

This is fundamentally different from the API gateway model, where edge traffic must reach the gateway — which is in the cloud, behind a load balancer, accessible from the internet. With leaf nodes, the edge has its own messaging infrastructure. It operates autonomously when the WAN is down. It synchronizes when the WAN is up. The security model works the same way in both cases.

## When the API gateway still makes sense

NATS JWT security doesn't replace every API gateway use case:

**External clients that don't speak NATS.** Mobile apps, browsers, and third-party webhooks speak HTTP. They need an HTTP endpoint. An API gateway or a thin NATS-to-HTTP bridge at the boundary is still necessary. The gateway becomes the edge translation layer — it terminates HTTP, validates API keys, and publishes to NATS subjects. But it doesn't need to handle internal service-to-service traffic. Its scope shrinks to the system boundary, where it belongs.

**API key management.** NATS supports JWT, nkey, and token authentication. It doesn't have a built-in API key management system with developer portals, key rotation APIs, and usage analytics. If you need those, an API gateway in front of the NATS boundary handles the developer-facing API management, while NATS handles the internal security.

**Request transformation and enrichment.** NATS doesn't inspect or modify message payloads. If you need to add headers, rewrite URLs, or transform request bodies based on client identity, you need something above the messaging layer. The gateway can do this at the edge, publishing enriched messages to NATS subjects.

**DDoS protection and WAF.** NATS doesn't provide web application firewall features or volumetric DDoS protection. Those are edge concerns that belong at the network boundary, not the messaging layer. A CDN or WAF in front of the gateway handles these.

The pattern is: gateway at the boundary for protocol translation and edge security, NATS internally for service-to-service communication and tenant isolation. The gateway doesn't need to be a full API management platform. It can be a thin HTTP-to-NATS bridge — validate the JWT, extract the tenant, publish to the tenant-scoped subject. The heavy lifting of authentication, authorization, and isolation is done by NATS.

## The Vitrifi pattern in practice

Vitrifi's architecture demonstrates this model end-to-end. The platform has two major sections: a Content Management System (where users design workflows) and a Core section (where workflows execute). NATS sits between them.

When workflows are published from the CMS, they're transformed into immutable objects and persisted using [NATS as a key-value store](https://nats.io/blog/nats-vitrifi-secret-weapon/). The KV store provides "strong consistency within clusters and eventual consistency across distributed deployments" — replacing what would traditionally be a database with operational features (revision history, cross-cluster replication, automatic expiration) built into the messaging infrastructure.

Workflow state updates flow through JetStream with pull consumers for high-volume messages and push consumers for discrete actions that need immediate handling (pausing or canceling a workflow). Exactly-once delivery is critical for the Trigger Server, which converts incoming messages into workflow initiation commands — achieved through "message de-duplication, delivery tracking and acknowledgement mechanisms."

Multi-cloud resilience is a natural property of the architecture: "if the cloud deployment fails, the private datacenter continues seamlessly because the platform is fundamentally asynchronous — all events flow through NATS." There's no failover to configure. No active-passive gateway pair to manage. The messaging fabric spans clouds. The services connect wherever they're running. Resilience is a property of decoupling, not a feature of a load balancer.

The security model is consistent across all of this — workflows in one tenant's account cannot see or affect workflows in another tenant's account. The CMS publishes to the tenant's subjects. The Core section subscribes to the tenant's subjects. The isolation is cryptographic, not network-based. Adding a tenant is provisioning an account. Removing a tenant is expiring the account JWT. The operational surface area for tenant management is minimal because the security model is built into the messaging infrastructure from the start.

## Security that scales with the system

The API gateway model scales security at the cost of centralization — bigger gateways, more gateway instances, more complex gateway configuration. The NATS model scales security at the cost of credential management — more JWTs, more accounts, more import/export policies. The difference is that credential management can be automated (issue JWTs at deploy time, expire them at decommission time, rotate them on a schedule) while centralized policy management becomes a bottleneck as the number of services and tenants grows.

A system with 100 services and 50 tenants behind an API gateway has one place where security policy is defined and one place where it can be wrong. A system with the same scale on NATS has 100 user JWTs, 50 accounts, and a handful of import/export policies — but each piece is small, independently verifiable, and automatically enforced by the server. The gateway concentrates risk. NATS distributes it. For systems that need to grow beyond what a single team can carefully manage, distribution beats concentration.

---

**References:**
- [NATS.io: The Secret Weapon Behind Vitrifi's Workflow Automation Platform](https://nats.io/blog/nats-vitrifi-secret-weapon/)
- [NATS Documentation — JWTs and Security](https://docs.nats.io/running-a-nats-service/configuration/securing_nats)
- [NATS Documentation — Leaf Nodes](https://docs.nats.io/running-a-nats-service/configuration/leafnodes)
- [NATS Documentation — Accounts](https://docs.nats.io/running-a-nats-service/configuration/securing_nats/accounts)

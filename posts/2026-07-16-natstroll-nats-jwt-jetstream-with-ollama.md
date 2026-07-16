---
title: "Natstroll: A NATS JWT+JetStream Capability Test with Ollama"
date: 2026-07-16
slug: natstroll-nats-jwt-jetstream-with-ollama
summary: Natstroll is a small but surprisingly dense NATS capability test — a hub-and-spoke joke exchange that exercises embedded NATS servers, JWT operator/account/user authentication, dynamic credential issuance, JetStream streams and durable pull consumers, scoped subject permissions, request/reply patterns, and OpenTelemetry trace propagation, all wrapped around an Ollama-powered AI conversation loop. This post walks through the architecture in detail, explaining what each component exercises and why the design decisions matter for anyone building NATS-based distributed systems.
tags: nats, jetstream, jwt, ollama, go, distributed-systems, messaging
---

[Natstroll](https://github.com/moresearch/natstroll) is one of those projects that looks like a toy on the surface — a hub and a spoke telling each other AI-generated jokes — but turns out to exercise a remarkably complete slice of the NATS ecosystem. Embedded servers, JWT authentication chains, dynamic user credentials, JetStream streams, durable pull consumers, scoped subject permissions, request/reply patterns, heartbeat monitoring, and distributed tracing all show up in roughly 800 lines of Go.

I built it as a lab. Not a production system, not a framework — a capability test. The kind of thing you write when you want to verify that all the pieces actually work together before reaching for them in something that matters.

## The shape of the thing

Natstroll has two binaries: a **hub** and a **spoke**. The hub owns an embedded NATS server. Spokes connect to it, register themselves, receive dynamically issued JWT credentials scoped to their identity, and then participate in a joke exchange loop — the hub sends an Ollama-generated joke, the spoke generates a comeback, the hub fires back with a follow-up, and so on.

```
┌──────────────────────────────────────────────────┐
│                      Hub                          │
│  ┌──────────────┐  ┌──────────┐  ┌────────────┐  │
│  │ Embedded     │  │ JWT      │  │ Ollama     │  │
│  │ NATS Server  │  │ Issuer   │  │ Client     │  │
│  │ (JetStream)  │  │          │  │            │  │
│  └──────┬───────┘  └────┬─────┘  └─────┬──────┘  │
│         │               │               │         │
└─────────┼───────────────┼───────────────┼─────────┘
          │               │               │
    ┌─────▼─────┐   ┌─────▼─────┐   ┌─────▼─────┐
    │ NATS      │   │ Dynamic   │   │ Ollama    │
    │ Messages  │   │ Creds     │   │ Replies   │
    └─────┬─────┘   └─────┬─────┘   └─────┬─────┘
          │               │               │
┌─────────┼───────────────┼───────────────┼─────────┐
│         │               │               │         │
│  ┌──────▼───────┐  ┌────▼─────┐  ┌─────▼──────┐  │
│  │ JetStream    │  │ JWT      │  │ Ollama     │  │
│  │ Consumer     │  │ Auth     │  │ Client     │  │
│  └──────────────┘  └──────────┘  └────────────┘  │
│                      Spoke                         │
└───────────────────────────────────────────────────┘
```

The architecture is deliberately clean. The hub provisions infrastructure — the server, the trust chain, the stream. The spoke provisions its own consumer. That split is the point: it tests whether dynamic credentials can carry enough permissions for a client to manage JetStream resources on its own behalf.

## What this actually exercises

Most NATS tutorials stop at `nats.Connect()` with a token or a static creds file. That's fine for getting started, but distributed systems have a way of surfacing edge cases the moment you step off the happy path. Natstroll deliberately walks into those edge cases.

### 1. Embedded NATS server

The hub starts a `nats-server` process in-process. This is not a mock — it's the real server, with a real config file written to a temp directory, a real JetStream store, and real JWT resolution. The code handles the full lifecycle:

- Generate a one-shot operator key and system account key
- Write operator and account JWTs to a resolver directory
- Configure JetStream with memory and file storage limits
- Wait for the server to accept connections before proceeding
- Clean up the temp directory on shutdown

```go
ns, err := server.NewServer(opts)
go ns.Start()
if !ns.ReadyForConnections(5 * time.Second) {
    ns.Shutdown()
    return nil, nil, fmt.Errorf("NATS server not ready")
}
```

This alone catches real issues: port conflicts, JetStream initialization races, the fact that `ReadyForConnections` returning true does not mean JetStream's account info endpoint is answering yet (hence the explicit `waitForJetStream` retry loop in the hub).

### 2. JWT trust chain (operator → account → user)

Natstroll uses full JWT/operator mode. The trust chain has three tiers:

- **Operator JWT**: signs account JWTs, names the system account
- **Account JWT**: enables JetStream with unlimited quotas (this is a lab), acts as the signer for user JWTs
- **User JWTs**: per-spoke or per-hub, carry scoped subject permissions

The hub generates the operator and account on startup, places the account JWT in a resolver directory, and issues user JWTs dynamically. Nothing is pre-provisioned.

This is the right model for multi-tenant NATS deployments. The operator is a trust anchor. The account is a security boundary. Users are scoped within an account. The system account exists to make the resolver work — it's not used by the application, but the embedded server won't start without one when you're in full resolver mode.

The test exercises the resolution path end-to-end: the server loads the operator JWT from a file, finds account JWTs in the resolver directory, and validates user JWTs presented in creds files. If any link in that chain breaks — wrong system account key, missing account JWT, expired user claims — the connection fails with a clear auth error.

### 3. Dynamic credential issuance

This is the most interesting part. Most NATS tutorials use static creds files generated once and distributed manually. Natstroll has the hub generate user credentials on the fly when a spoke registers.

The spoke connects with narrow **registrar** credentials:

```go
claims.Pub.Allow = []string{"reg.request"}
claims.Sub.Allow = []string{"_INBOX.>"}
```

These can only publish a registration request and receive on inbox subjects used by NATS request/reply. If these credentials leak, the attacker can register — but they can't publish to joke subjects, read the stream, or do anything else.

When a spoke sends `reg.request` with its `SPOKE_ID`, the hub generates a new nkey pair, creates a user JWT with subject permissions scoped to that specific spoke identity, and returns the full creds file. The spoke then drops the registrar connection, reconnects with the new dynamic credentials, and proves they work by creating a JetStream consumer.

The spoke's dynamic permissions are identity-scoped:

```go
claims.Pub.Allow = []string{
    "heartbeat." + id,                          // can publish its own heartbeat
    shared.JokeResponseSubject + id + ".>",      // can publish joke responses
    "$JS.API.>",                                 // can manage JetStream
    "$JS.ACK." + shared.JokeStream + "." + consumerName + ".>",
}
claims.Sub.Allow = []string{
    "_INBOX.>",
    shared.JokeRequestSubject + id,             // only its own joke requests
}
```

Spoke A cannot subscribe to Spoke B's joke requests. It cannot publish heartbeats as Spoke B. It cannot publish responses on Spoke B's subjects. The NATS server enforces this at the protocol level — the spoke literally cannot express a message on the wrong subject.

### 4. JetStream stream and consumer separation

The hub creates the stream. The spoke creates its own consumer. This split is intentional.

**Hub side**: creates `JOKE_STREAM` covering both request and response subjects:

```go
cfg := &nats.StreamConfig{
    Name:     shared.JokeStream,
    Subjects: []string{shared.JokeRequestSubject + ">", shared.JokeResponseSubject + ">"},
    Storage:  nats.FileStorage,
}
```

**Spoke side**: creates a durable pull consumer filtered to its own request subject:

```go
cfg := &nats.ConsumerConfig{
    Durable:       consumerName,
    AckPolicy:     nats.AckExplicitPolicy,
    FilterSubject: filterSubject,  // joke.request.<spokeID>
}
```

This pattern matters for real systems. The party that owns the data (the hub, or an ops team) provisions the stream. The party that processes the data (a spoke, or a microservice) provisions its own consumer with the exact filter and ack policy it needs. Dynamic credentials must carry `$JS.API.>` to make this possible — and that's the capability under test.

The spoke also handles consumer creation idempotently. If the consumer already exists with the right filter, it moves on. If it exists with the wrong filter (from a previous run with a different config), it errors out rather than silently misbehaving. This is the kind of detail that separates a demo from something you'd actually run.

### 5. Request/reply with reply subject scoping

The hub's conversation loop uses a pattern worth studying:

```go
replySubject := shared.JokeResponseSubject + targetSpokeID + "." + requestID
sub, err := nc.SubscribeSync(replySubject)
// ... publish joke request with Reply set to replySubject ...
msg := &nats.Msg{
    Subject: shared.JokeRequestSubject + targetSpokeID,
    Reply:   replySubject,
    Data:    reqData,
}
js.PublishMsg(msg)
replyMsg, err := sub.NextMsg(SpokeTimeout)
```

The hub publishes the joke request via JetStream (so it's persisted) but waits for the reply through a core NATS subscription on a dynamically generated reply subject. The reply subject is scoped to the spoke's identity and includes a UUID request ID — so the spoke's JWT only needs `joke.response.<spokeID>.>` permissions, and reply routing is naturally collision-free.

The timeout is generous (75 seconds) because the spoke's Ollama model gets 60 seconds for generation, and network latency plus JSON marshaling adds a bit more. If the spoke doesn't reply in time, the hub logs the failure and the conversation loop ends. No hung goroutines, no leaked subscriptions — the `sub.Unsubscribe()` in the defer-like pattern at each iteration cleans up.

### 6. Heartbeat monitoring

Each spoke publishes a heartbeat every 10 seconds on `heartbeat.<spokeID>`:

```go
nc.Publish("heartbeat."+spokeID, []byte(`{"status":"alive"}`))
```

The hub subscribes to `heartbeat.>` and logs every heartbeat it sees. This is a simple liveness pattern, but it verifies that the spoke's dynamic credentials actually allow publishing to its scoped heartbeat subject, and that the hub's wildcard subscription works across all registered spokes.

### 7. OpenTelemetry trace propagation

Both binaries support opt-in OTLP trace export. When `OTEL_EXPORTER_OTLP_ENDPOINT` is set, the hub and spoke initialize a gRPC trace exporter with proper resource attributes (`service.name`, `service.version`, `host.name`).

Trace context is propagated through NATS message headers:

```go
func InjectTraceContext(ctx context.Context, msg *nats.Msg) {
    otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(msg.Header))
}

func ExtractTraceContext(msg *nats.Msg) context.Context {
    return otel.GetTextMapPropagator().Extract(context.Background(), propagation.HeaderCarrier(msg.Header))
}
```

This means you can follow a joke request from the hub's `hub.generate-joke` span through the NATS transport to the spoke's `spoke.generate-reply` span, and back — a full distributed trace across processes connected only by a NATS cluster. In production, this is what lets you debug latency in message-driven systems without stitching together log timestamps by hand.

When the endpoint is unset, the tracer falls back to the global no-op tracer — no spans are exported, no log noise, no dependency on a running collector. Opt-in with a clean fallback is the right default for a lab.

## The Ollama integration

Natstroll defaults to `deepseek-r1:1.5b`, a thinking model that spends tokens on internal chain-of-thought before producing a final answer. This is deliberate — thinking models stress-test the timeout and fallback paths.

The prompt prefix `"Return only the final answer. Do not think out loud."` suppresses most reasoning output, but if `num_predict` is too low, the model may spend all its tokens on hidden reasoning and return an empty string. The code handles this:

- The hub checks for empty Ollama responses and returns a descriptive error instead of publishing a blank joke
- The spoke falls back to a deterministic reply if Ollama fails: `"That joke took off about as well as a cat's paper airplane."`
- Both sides use a 60-second timeout with `num_predict: 256` — generous enough for a short joke, tight enough that a hung model doesn't wedge the system

These aren't just Ollama details. They're resilience patterns: fail loud, fail fast, provide a degraded response rather than crashing.

## The credential bootstrap flow

The first-run experience deserves attention because it gets the security model right from the start.

When you run the hub without `NATS_ACCOUNT_SEED` set, it generates a fresh account key, creates narrow-scoped registrar credentials, prints them, and exits:

```
========== COPY THESE EXACTLY ==========
export NATS_ACCOUNT_SEED="SAAPG3R4G..."
export REGISTRAR_CREDS_B64="W0Zvcm..."
========================================
```

You paste those into your terminal and run the hub again — this time it starts the server. The registrar credentials can only publish `reg.request` and subscribe to `_INBOX.>`. They can't read the joke stream, can't publish heartbeats, can't impersonate a spoke. If they leak, the blast radius is contained: an attacker can register, but registration is the only thing those credentials authorize.

Contrast this with most "getting started" setups that hand you a creds file with `pub: [">"]` and `sub: [">"]` and tell you to get to production first. Natstroll starts with principle-of-least-privilege from the first `go run`.

## Security: what's intentionally loose and why

The spoke's dynamic credentials include `$JS.API.>` — broad JetStream API access. This is the capability under test: can a dynamically issued credential create and bind a durable pull consumer? The answer is yes, and proving it requires those permissions.

The README is explicit about what you'd change for production:

- Have the hub create consumers instead of the spoke
- Scope JetStream API permissions to exact subjects per consumer
- Don't write debug credentials to `/tmp`
- Add credential rotation
- Add per-spoke quotas

There is a difference between a lab being insecure and a lab documenting where the security boundaries are and why they're drawn where they are. Natstroll does the second thing. The subjects that exist, the permissions that are granted, and the reasons for each grant are all visible in roughly 20 lines of `issueSpokeCredentials`.

## Why this matters as a demo

Most distributed-systems demos pick one thing and show it in isolation. A JWT tutorial shows a static creds file. A JetStream tutorial shows a stream and consumer created by the same client with admin credentials. An OpenTelemetry tutorial shows traces between HTTP services.

Natstroll compresses all of these into a single coherent flow. The JWT issuance feeds into the JetStream consumer creation. The request/reply pattern carries trace context. The heartbeat loop verifies that scoped subjects work. The credential upgrade (registrar → dynamic) proves that reconnection with new credentials is seamless.

It's also honest about its constraints. It uses a single account, so the account is both the issuer and the security boundary — production multi-tenancy would split those. It grants broad JetStream API access to spokes for testing purposes. It doesn't persist state between restarts. The conversation loop only involves the first registered spoke.

These are not bugs. They are scope decisions that make the lab small enough to understand in one sitting but complete enough to surface real distributed-systems concerns.

## Running it

```bash
git clone https://github.com/moresearch/natstroll
cd natstroll
ollama pull deepseek-r1:1.5b

# Terminal 1: bootstrap and start the hub
unset NATS_ACCOUNT_SEED REGISTRAR_CREDS_B64
go run ./cmd/hub
# → copy the export lines
export NATS_ACCOUNT_SEED="..."
export REGISTRAR_CREDS_B64="..."
go run ./cmd/hub

# Terminal 2: start a spoke
export NATS_URL=nats://127.0.0.1:4222
export REGISTRAR_CREDS_B64="..."
export SPOKE_ID=black-spoke
go run ./cmd/spoke
```

Add more spokes with different `SPOKE_ID` values. Spy on traffic with `nats` CLI tools and the debug credentials file. Enable OpenTelemetry by setting `OTEL_EXPORTER_OTLP_ENDPOINT`.

The Makefile produces cross-compiled, stripped binaries for Linux and Windows on amd64 and arm64 — useful if you want to run spokes on different machines or architectures.

## What I'd reach for it for

Natstroll is not a library. It's not a framework. It's a reference: here is a known-good configuration for an embedded NATS server in JWT mode, here is how dynamic credential issuance works, here is how a pull consumer on JetStream is created by a dynamically authenticated client, here is how trace context flows through NATS headers.

When I'm building something that needs any of these pieces, I'd rather start from something that exercises all of them together than from five separate tutorials that may or may not compose. That's the value of a capability test: it proves the integration works before you commit to it.

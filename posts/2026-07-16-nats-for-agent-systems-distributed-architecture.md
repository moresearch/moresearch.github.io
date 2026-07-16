---
title: "NATS for Agent Systems: The Distributed Architecture AI Needs"
date: 2026-07-16
slug: nats-for-agent-systems-distributed-architecture
summary: Agent systems are distributed systems in a trench coat — heterogeneous runtimes, mixed clouds, private data, and ephemeral clients that need to discover each other, exchange messages, and maintain liveness. Most teams glue this together with HTTP and hope. NATS provides a protocol-native substrate for agent communication, and the open-source Synadia Agent Protocol for NATS defines a contract for discovery, conversation, and liveness that works across any language, runtime, or environment. This post explains why HTTP is the wrong transport for agents and how the NATS agent protocol inverts the architecture.
tags: nats, ai-agents, distributed-systems, agent-protocol, mcp, microservices
---

The NATS blog recently published a post titled "[What's old is new: A NATS-native protocol for AI agents](https://nats.io/blog/nats-native-protocol-for-ai-agents/)" that makes a deceptively simple argument: agent systems are distributed systems, and NATS already solves distributed systems communication. The observation is so obvious in retrospect that it's worth walking through why the industry defaulted to HTTP for agent communication, why that default is wrong, and what the alternative looks like.

## Agent systems are distributed systems

An agent system in 2026 looks something like this: a Claude or GPT model running in one cloud, a DeepSeek model running on local hardware, a code execution sandbox in a different region, a vector database with proprietary documents, a Slack bot that can be invoked by name, a browser automation tool running as a sidecar, and a human-in-the-loop approval service that handles sensitive operations.

Each of these components is a different runtime. Different language. Different cloud. Different network. Different security boundary. Different lifecycle. Different owner.

The NATS blog describes this situation with a memorable image: "agentic systems are distributed systems in a trench coat." They look like one thing from the outside (a helpful assistant) but are actually a collection of independent services that need to find each other, exchange messages, handle long-running requests, and notice when a peer disappears.

This is not a new problem. Microservices solved it a decade ago. The patterns are well-understood: service discovery, request-reply, pub/sub, heartbeats, load balancing, circuit breaking, authentication, authorization. What's new is that agents are applying these patterns across a wider heterogeneity — more runtimes, more clouds, more trust boundaries, more ephemeral participants — than most microservices deployments ever faced.

## Why HTTP is wrong for agents

The current default for agent communication is HTTP. Anthropic's Model Context Protocol (MCP) uses HTTP. Most agent frameworks use HTTP. OpenAI's function calling uses HTTP. The reasoning is pragmatic: every language has an HTTP client, every cloud allows HTTP, every developer knows HTTP.

But HTTP for agent communication repeats the same architectural mistake that microservices made with REST. It couples the caller to the callee's address. It assumes synchronous request-response. It requires the callee to be online when the caller calls. It has no native concept of a stream of partial results — you need Server-Sent Events or WebSockets bolted on. It has no native concept of liveness — you need health check endpoints and polling. It has no native concept of peer discovery — you need a service registry, DNS, or hardcoded URLs.

The NATS blog frames this precisely: "most people end up gluing components together with HTTP or forcing everything through a single vendor's gateway." The glue is the problem. Every point-to-point HTTP connection is a coupling point. Every vendor gateway is a single point of failure and a single point of vendor lock-in.

NATS inverts this. Instead of agents connecting to each other, agents connect to NATS. The subject is the address. The subscription is the availability. The message is the protocol. No agent knows any other agent's IP address, port, or HTTP endpoint. They know subjects. The NATS infrastructure handles routing, load balancing, and liveness.

## The Synadia Agent Protocol for NATS

The [Synadia Agent Protocol for NATS](https://github.com/synadia-io/synadia-agent-protocol) (version 0.3 at time of writing) is not a framework, runtime, or product. It's a contract. "If an agent does these few specific things," the blog explains, "anything else on the same NATS system can find it and talk to it."

The protocol defines three responsibilities: discovery, conversation, and liveness.

### Discovery

Agents register as NATS micro services under the service name `agents`. Their metadata describes the agent type, owner, protocol version, and optionally a session identifier:

```json
{
  "name": "claude-sonnet",
  "version": "0.1.0",
  "metadata": {
    "agent": "claude",
    "owner": "engineering",
    "protocol_version": "0.3",
    "session": "session-abc123"
  }
}
```

Callers discover agents by sending a standard micro service ping:

```
$SRV.PING.agents
```

Every running agent the caller has permission to see responds. For full endpoint information (subjects for prompts, capability metadata, queue groups), callers query:

```
$SRV.INFO.agents
```

The NATS blog emphasizes the architectural implication: "No registry. No service catalog. No coordinator process you have to keep alive." Discovery is decentralized. The NATS infrastructure is the registry. Agents come and go, and the set of reachable agents is whatever is currently connected and authorized. No separate Consul, etcd, or DNS-based discovery. No registration API. No heartbeat-to-registry. The connection *is* the registration.

This matters for agent systems specifically because agents are ephemeral in ways that microservices aren't. A microservice might run for months. An agent might be a single CLI invocation that lasts 30 seconds — a human asks a question, the agent processes it, the agent exits. Discovery that requires explicit registration and deregistration breaks down when agents are this short-lived. NATS discovery works because it's connection-scoped: when the agent connects, it's discoverable. When it disconnects, it's gone. No cleanup. No stale registrations.

### Conversation

Prompting an agent is a single NATS request on a subject following the pattern `agents.prompt.{agent}.{owner}.{name}`:

```
nc.Request("agents.prompt.claude.engineering.sonnet", []byte("Summarize this document"), 60*time.Second)
```

The verb-first hierarchy (`agents.prompt.>`) means a single wildcard subscription captures all prompt traffic across all agents — useful for auditing, logging, or routing.

The agent streams typed JSON chunks back to the caller's reply inbox. Each chunk is `{type, data}`:

- **`response`** chunks deliver content. The data is a string or an object with `text` and optional `attachments`. Multiple response chunks concatenate in publication order. This is native streaming — no SSE, no WebSocket upgrade, no chunked transfer encoding. The stream is the sequence of NATS messages on the reply subject.

- **`status`** chunks carry lifecycle signals. A single `ack` status must be the very first message — before any latency-inducing work like model inference — so the caller knows the request was received and can reset its inactivity timer. Without this, the caller has no way to distinguish "the agent is thinking" from "the agent never got the message."

- **`query`** chunks allow the agent to pause mid-stream and ask the caller a question. This is how an agent requests permission ("Can I read this file?"), clarification ("Which of these three documents?"), or a menu selection. Each query provides a fresh reply subject for the answer. Multiple queries can be in flight concurrently. The stream pauses on a query and resumes when the caller answers.

Every stream terminates with a zero-byte message carrying no headers. This is the protocol's way of saying "I'm done" without requiring the agent to predict how many chunks it will send. The terminator is unambiguous. The caller knows the conversation is complete.

Errors use NATS micro service error headers immediately before the terminator:

```
NATS-Service-Error: permission_denied
NATS-Service-Error-Code: 403
```

The error format is standard NATS. Any NATS tooling that understands micro service errors understands agent errors. No custom error format. No HTTP status codes mapped to application errors. The protocol leverages what NATS already provides.

### Liveness

Each agent publishes a heartbeat on `agents.hb.{agent}.{owner}.{name}` at a configurable cadence (default 30 seconds). The payload carries agent identity and a per-instance `instance_id`:

```json
{
  "agent": "claude",
  "owner": "engineering",
  "name": "sonnet",
  "instance_id": "a1b2c3d4",
  "timestamp": "2026-07-16T14:30:00Z",
  "cadence_seconds": 30
}
```

The heartbeat subject is fixed — the one subject agents cannot override. A caller subscribes to `agents.hb.*.*.*` and watches every agent on the cluster come up, stay up, and go offline without polling. An instance is considered offline after three missed beats (90 seconds at default cadence).

For point-in-time checks, every agent exposes a `status` request/reply endpoint returning the same payload shape as a heartbeat, freshly built on demand. This lets a caller verify an agent is alive without waiting for the next heartbeat.

The heartbeat pattern is simple but powerful. It's the same pattern NATS itself uses for route and gateway connections. It's the same pattern embedded in the micro services protocol. It works at scale because the NATS server handles the fan-out — `agents.hb.*.*.*` matches all agent heartbeats without the server doing per-agent work. The wildcard subscription is a single entry in the subscription table. The matching is a trie traversal.

## Security: delegated, not reinvented

The agent protocol does not introduce its own security layer. It delegates entirely to NATS's existing primitives — accounts, users, and subject permissions. The blog is explicit: "End-to-end encryption and strong agent identity are explicitly deferred to future extensions."

This is the right call for v0.3. NATS already provides:

- **Authentication** at connection time via JWT, nkey, or token
- **Authorization** at publish/subscribe time via user JWT `pub.allow` and `sub.allow`
- **Isolation** via accounts — agents in different accounts cannot see each other's subjects
- **Import/export** for cross-account communication with explicit opt-in

A production deployment might look like: the `engineering` account owns internal agents (code exec, document search). The `public` account owns user-facing agents (ChatGPT plugin, Slack bot). The public account imports `agents.prompt.>` from engineering with a restricted set of allowed subjects. Engineering agents can't see public agents unless explicitly exported. The security boundary is the account. The policy is the import/export configuration. No API gateway. No mTLS sidecar per agent. No SPIFFE identity per process.

## Bridging environments

The protocol deliberately avoids ties to any framework, language, or runtime. A caller doesn't need to know "whether an agent is a Python script, a hosted model wrapper, a CLI session, or a long-running service." They discover capabilities, send prompts, receive typed streams, and observe liveness the same way across all of them.

This is NATS's core strength applied to agents. NATS already bridges heterogeneous environments — Linux, Windows, macOS, ARM, x86, cloud, edge, on-prem. The agent protocol inherits this. An agent running on a Raspberry Pi on a factory floor can participate in the same agent mesh as a Claude instance in us-east-1. The subject namespace is global. The security is consistent. The liveness is observable. The protocol doesn't care about the runtime.

The blog notes that reference SDKs exist in TypeScript and Python, but "the NATS CLI can speak the protocol directly, and anything that knows the rules can participate." This is the litmus test for a protocol: can you implement it from the spec without a library? The agent protocol passes. The subject conventions are documented. The JSON envelope is documented. The chunk format is documented. The terminator is documented. You can write an agent in bash with `nats` CLI and `jq` if you want to. That's not a toy feature — it means the protocol is simple enough to be correct.

## Why this matters for the agent ecosystem

The agent ecosystem in 2026 is fragmenting the same way the microservices ecosystem fragmented in 2016. Every framework has its own communication protocol. Every vendor has its own gateway. Interoperability is aspirational, not real. If you build an agent with LangChain and I build one with a custom Go service, they can't discover each other, can't exchange messages, can't observe each other's liveness — unless someone writes an adapter, a bridge, a translation layer.

The NATS agent protocol says: don't bridge. Use the same pipe. Connect to NATS. Follow the subject convention. Send the JSON chunks. Publish the heartbeat. If every agent does these four things, every agent can talk to every other agent — regardless of framework, language, cloud, or runtime. That's the promise of a protocol. That's the promise of dumb pipes and smart endpoints, applied to AI.

---

**References:**
- [What's old is new: A NATS-native protocol for AI agents](https://nats.io/blog/nats-native-protocol-for-ai-agents/)
- [Synadia Agent Protocol for NATS (GitHub)](https://github.com/synadia-io/synadia-agent-protocol)
- [NATS Micro Services Protocol](https://docs.nats.io/running-a-nats-service/configuration/services)
- [NATS Security — Accounts and JWTs](https://docs.nats.io/running-a-nats-service/configuration/securing_nats)

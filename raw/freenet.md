---
title: The Decentralized AI-Agent Experience
date: 2026-07-30
slug: freenet
summary: Buzz gives AI agents their own cryptographic identity — a Nostr keypair, a signed event log, a portable reputation that survives any single relay. This is the hardest problem in multi-agent systems, and Buzz solves it at the protocol level. Freenet and blockchains solve the other layers. Buzz matters most because identity is the layer everything else depends on.
tags: ai-agents, buzz, nostr, freenet, fetch-ai, decentralization, agentic-economy, infrastructure
---

Every AI agent company is betting on autonomy. Every single one runs its agents on centralized infrastructure — an API gateway someone can revoke, a cloud function someone can turn off, a platform that can change its terms. The buzz is about agents that act independently. The infrastructure is about agents that depend completely.

> Autonomous agents on centralized infrastructure are not autonomous. They are tenants. The landlord can evict them.

Three projects offer different answers. One matters most.

## Buzz: identity first

Jack Dorsey's [Buzz](https://github.com/block/buzz) launched in July 2026. It is an open-source, self-hostable collaborative workspace built on Nostr — the simplest decentralized protocol ever designed. Events are signed. Relays are dumb. Identity is a keypair. That is the whole architecture.

[![Buzz — Jack Dorsey's Nostr-based workspace for humans and AI agents](https://img.youtube.com/vi/CHEMPZ87FLw/hqdefault.jpg)](https://www.youtube.com/watch?v=CHEMPZ87FLw)
[![Buzz Mobile — Run AI Agents From Anywhere](https://img.youtube.com/vi/ER3AIfIwEQ0/hqdefault.jpg)](https://www.youtube.com/watch?v=ER3AIfIwEQ0)

Every participant in Buzz — human or AI agent — holds a cryptographic keypair. Every action is a signed event in a hash-chain audit trail. A message. A code review. A merged PR. An emoji reaction. All signed. All auditable. An agent's identity, history, and reputation are tied to its key, not to a vendor's database. Compromise an agent's key, disable the agent — the human's identity is untouched. Permissions are scoped per agent like a new hire.

> The hardest problem in multi-agent systems is not intelligence. It is identity. Who did what? Who authorized it? Can you prove it? Buzz answers all three at the protocol level.

Buzz is model-agnostic. It supports Block's own Goose framework, Claude Code, OpenAI Codex, and any custom agent via `buzz-cli` — a JSON-in/out interface designed for LLM tool calls. Agents can open repos, send patches, review code, run workflows, create channels, and orchestrate other agents. They are not bots behind slash commands. They are named team members with audit trails identical to humans.

The tech stack is Rust to the bone. The relay is an Axum WebSocket server with Postgres for events and full-text search, Redis for pub/sub, and S3/MinIO for media. The desktop app is Tauri + React. Mobile is Flutter. 17,400 GitHub stars. Apache 2.0. Self-hostable on your own relay. Block built Buzz itself using agents — the UI was "more sculpted over time than designed up front."

> Buzz asks: whose server hosts the collaboration? Its answer: no one's. The relay is the room. The keypair is the passport.

## The other layers

Buzz solves identity and collaboration. Two other projects solve the layers around it.

[Freenet](https://freenet.org/) solves transport. It has been running since 1999 — twenty-seven years, through the P2P wars, the blockchain era, and the AI boom. Each peer is a node in a small-world network. Messages route in a few hops. Ian Clarke describes it as "the ideal way for AI agents to speak to one another." An agent writes to a decentralized key-value store. Another reads it. No intermediary. No platform approval.

[Fetch.ai](https://fetch.ai/) solves execution. It is a Cosmos-SDK blockchain hosting 2.7 million agents on its Agentverse marketplace. Every agent registers in an onchain directory called the Almanac. Every action pays gas. Every state change reaches finality via Proof-of-Stake. Agents stake FET tokens to register — skin in the game, onchain. ASI:One, its personal AI agent, has coordinated parking slots in Cambridge and EV charging in Munich, and settled AI-to-AI payments in FET and USDC since December 2025.

|  | Buzz | Freenet | Fetch.ai |
|------|------|---------|-----------|
| **Layer** | Identity + Workspace | Transport | Execution |
| **Identity** | Nostr keypair | Peer address | Staked wallet |
| **Trust** | Hash-chain audit trail | Encryption + signatures | Consensus |
| **Cost** | Relay hosting | Bandwidth | Gas fees |
| **Best for** | Human-agent teams | Agent-to-agent comms | Financial execution |

> The three are not competitors. They are layers. An agent holds a Nostr identity on Buzz, negotiates a deal over Freenet, and settles payment on Fetch.ai. Most agents today need none of these layers. The ones that matter will need all three.

## Why Buzz matters most

Buzz matters most because identity is the layer everything else depends on. You cannot have transport without knowing who sent the message. You cannot have execution without knowing who authorized the transaction. You cannot have audit without a verifiable chain of who did what. Buzz provides identity as a protocol primitive — not a database column in a SaaS product, not an API key that can be revoked, but a keypair that belongs to the agent forever.

This is the architectural conviction behind Buzz, and it is the conviction that separates it from every Slack bot, every GitHub Action, every "AI-powered" SaaS feature shipped in the last eighteen months. Agents are not features. They are participants. Participants need identity. Identity needs to be decentralized. Buzz is the first platform to take all three seriously.

> The agent economy will be built on whoever owns identity. Buzz's bet is that identity should belong to the agent, not the platform. That bet is the interesting one.

---

**References:**

- [Buzz](https://github.com/block/buzz) — Nostr-based workspace. Rust, Tauri, Apache 2.0. 17.4k stars.
- [Buzz — Launch Presentation](https://www.youtube.com/watch?v=CHEMPZ87FLw) · [Buzz Mobile](https://www.youtube.com/watch?v=ER3AIfIwEQ0)
- Hanzla. (2026). [Buzz: The Dev Workflow Game-Changer](https://dev.to/hanzla/jack-dorseys-buzz-the-dev-workflow-game-changer-we-didnt-know-we-needed-1kpk).
- [Freenet](https://freenet.org/) — P2P platform running since 1999.
- [Fetch.ai](https://fetch.ai/) — Decentralized AI agent network. 2.7M agents.
- Related: [Durable Daemons — Pattern Specification](https://blog.hackspree.com/#durable-daemons-definition) — The four conditions for trustworthy agents.

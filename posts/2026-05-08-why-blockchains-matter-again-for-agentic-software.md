---
title: Blockchains for Agentic Software
date: 2026-05-08
slug: why-blockchains-matter-again-for-agentic-software
summary: Blockchains become interesting again when coding agents need programmable, auditable, and explicitly economic coordination instead of opaque platform rules.
tags: blockchain, agents, research
---

I did not focus on blockchain in my dissertation because I wanted to rehash crypto slogans. I focused on it because blockchains make economic rules programmable, transparent, and auditable.

That matters a lot more in an agentic world than it did in the earlier platform era.

In [my dissertation on SWE-Agent Economics and SWEChain-SDK](https://repositorio.ufu.br/bitstream/123456789/47974/3/SoftwareEngineeringAgent.pdf), I argue that decentralized SWE-Agent outsourcing markets are worth studying precisely because centralized platforms hide too much of the mechanism. They hide the ordering logic, the settlement logic, the admission rules, and often the event history needed for serious analysis.

If agents are going to negotiate, specialize, and compete in software markets, I want those rules visible.

## Why blockchain was the right research substrate

The dissertation frames a blockchain not as branding, but as an implementation of an economic mechanism. That distinction mattered a lot to me.

I was interested in:

- explicit allocation rules,
- transparent bids and payments,
- auditable state transitions,
- reproducible experiments under fixed conditions.

A blockchain-style substrate is useful there because it makes state change legible. Every bid, allocation, payment, and artifact trail can be logged as part of the system rather than reconstructed later from scattered dashboards.

That is exactly why one of the central contributions of the dissertation is **SWEChain-SDK**, a local-first blockchain-native SDK for economic network simulations of decentralized SWE-Agent markets.

## Why this gets more relevant in the agentic era

The stronger agents become, the more we need clean answers to questions like:

1. who can submit work,
2. who is allowed to bid,
3. how selection happens,
4. how settlement happens,
5. what evidence counts as completion,
6. how disputes or failures are inspected afterward.

Traditional software systems can answer those questions too, but they often do so in an opaque way. Blockchains are interesting here because they make those policies first-class and programmable.

That is why I focused on them in research. I was less interested in speculation and more interested in mechanism visibility.

## Go is a natural language for the experimental surface

Even if the settlement substrate is blockchain-based, the surrounding tooling still benefits from straightforward systems code. A local-first SDK needs CLIs, dashboards, bridges, and deterministic utilities. That is where Go fits very naturally.

A small Go surface makes the policy layer easier to reason about:

```go
package settlement

// AuctionResult is the minimal state needed to settle a finished task.
type AuctionResult struct {
	TaskID      string
	WinnerID    string
	PriceCents  int64
	ArtifactRef string
}

// Ledger abstracts the settlement backend behind one explicit call.
type Ledger interface {
	Settle(result AuctionResult) error
}

func Finalize(ledger Ledger, result AuctionResult) error {
	// Keep the settlement path obvious so it is easy to audit.
	return ledger.Settle(result)
}
```

Again, the point is not complexity. The point is clarity. If agentic systems are going to rely on explicit mechanisms, the code around those mechanisms should be boring, auditable, and testable.

## Why I think this topic is still underrated

A lot of agent discussion still assumes coordination will be solved inside application logic alone. I think that misses the opportunity.

Once agents are meaningful economic actors, infrastructure matters. Settlement matters. Logging matters. The ability to replay and inspect the exact rule path matters. That is why a blockchain-native SDK felt like a useful research artifact rather than a gimmick.

The dissertation made that case because I wanted a platform where decentralized SWE-Agent markets could be studied under controlled, paired experiments. If you want to compare policies seriously, you need the mechanism to be part of the experiment, not an invisible dependency.

## The real reason I cared

I focused on this because I think agentic software engineering will eventually force us to choose between opaque coordination and explicit coordination. My bet is that explicit coordination wins, especially in high-stakes systems where trust, incentives, and auditability matter.

That is why blockchains matter again in this context. Not because they make agents magical, but because they make the rules legible.

## Source

- [Mohamed A. Fouad, *Software Engineering Agent Economics: A Blockchain Software Development Kit for Economic Network Simulations*](https://repositorio.ufu.br/bitstream/123456789/47974/3/SoftwareEngineeringAgent.pdf)


Game design and systems thinking share a common structure: a set of rules, agents acting within those rules, and emergent behavior that no individual agent intended. The same structure appears in market design, protocol design, and software architecture. The engineer who studies games learns to see the rules behind the behavior. The behavior that looks like chaos is often the equilibrium of a system whose rules you haven't discovered yet. The discovery is the engineering.

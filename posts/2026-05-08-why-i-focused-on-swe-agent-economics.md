---
title: Why I focused my research on SWE-Agent economics
date: 2026-05-06
slug: why-i-focused-my-research-on-swe-agent-economics
summary: My dissertation focused on SWE-Agent economics because software work is becoming a market of autonomous decision-makers, not just a pipeline of human tickets.
tags: research, agentic-economics, golang
---

One of the clearest questions in front of us is not whether coding agents will get better. They will. The harder question is what happens when software work itself starts behaving like an economic system.

That is why I focused my dissertation on this area. In [*Software Engineering Agent Economics: A Blockchain Software Development Kit for Economic Network Simulations*](https://repositorio.ufu.br/bitstream/123456789/47974/3/SoftwareEngineeringAgent.pdf), I framed what I call **SWE-Agent Economics** as the intersection of intelligent software engineering and software-engineering economics. I was not interested only in whether agents can solve tasks. I was interested in what happens when they bid, specialize, compete, consume priced resources, and operate under explicit allocation rules.

That focus came from a simple observation: once software work can be decomposed into scoped issues, artifacts, tests, and payments, the system starts to look less like a task board and more like a market.

## Why I thought the economic lens mattered

A lot of discussion around coding agents still treats them as isolated assistants. That framing is too small.

In practice, agents already operate under:

- budget limits,
- latency limits,
- tool-access limits,
- quality thresholds,
- routing and scheduling decisions.

Those are economic constraints, even when people describe them as product or platform constraints.

My research focused on this because I wanted a language and an experimental setup that could capture those interactions directly rather than pretending they were just implementation details. The dissertation argues that decentralized SWE-Agent outsourcing markets are a useful central case because they make allocation, pricing, and settlement rules explicit.

## Why I built a toolkit instead of writing only theory

The dissertation does not stop at framing. One of its main contributions is **SWEChain-SDK**, a blockchain-native SDK for controlled economic network simulations of SWE-Agent markets. I cared about that because good ideas are cheap if nobody can run the experiment again under comparable conditions.

I wanted an environment where we could vary one policy dimension at a time and still keep:

1. the same datasets,
2. the same time base,
3. the same agent pool,
4. the same logging surface.

That is what makes claims about agent economics credible instead of anecdotal.

From a Go perspective, that kind of work benefits from explicit contracts and small composable binaries. A simulation stack like this should make every surface painfully clear:

```go
package market

type Bid struct {
	AgentID string
	TaskID  string
	Price   int64
	Score   float64
}

type Allocation struct {
	TaskID   string
	AgentID  string
	Accepted bool
}

type Logger interface {
	RecordBid(Bid) error
	RecordAllocation(Allocation) error
}
```

The code is intentionally boring. That is the point. If the economic mechanism is the thing under study, the interfaces around it should be stable enough to let experiments change policy without rewriting the whole platform.

## Why this mattered to me as a software engineering problem

I focused on this line of research because software engineering is entering a phase where coordination matters as much as raw model capability. The interesting question is no longer only “can an agent solve this issue?” It is also:

- which agent should do it,
- under what incentives,
- at what price,
- with what comparative advantage,
- under which settlement rule,
- and with what observable audit trail.

That is a very software-engineering question, but it is also an economic one.

The dissertation formalizes that intuition because I think the next stage of agentic software engineering will reward teams that can reason about incentives and market structure, not only prompts and models.

## The real motivation

The deepest reason I focused on SWE-Agent economics is that it creates a bridge between two worlds that are often separated: engineering systems and economic systems. Agents force them back together.

Once software workers become autonomous, the surrounding system has to answer questions about specialization, cost, settlement, trust, and transparency. I wanted my research to address those questions directly, with reproducible tooling rather than vague metaphors.

That is why the dissertation centers this topic. I think it is one of the most important lenses for understanding where software engineering is heading next.

## Source

- [Mohamed A. Fouad, *Software Engineering Agent Economics: A Blockchain Software Development Kit for Economic Network Simulations*](https://repositorio.ufu.br/bitstream/123456789/47974/3/SoftwareEngineeringAgent.pdf)

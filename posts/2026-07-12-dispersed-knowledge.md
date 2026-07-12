---
title: The knowledge is dispersed
date: 2026-07-12
slug: dispersed-knowledge
summary: "Hayek wrote in 1945: 'The knowledge of the circumstances of which we must make use never exists in concentrated form, but solely as the dispersed bits of incomplete and frequently contradictory knowledge which all the separate individuals possess.' This is the most important sentence in economics for software engineers."
tags: hayek, knowledge, distributed-systems, teams, architecture
---

In 1945, Friedrich Hayek published "The Use of Knowledge in Society." The paper is about economics. It is also about software architecture. It is also about why your team's decisions are worse than you think, and why that's not their fault.

> "The knowledge of the circumstances of which we must make use never exists in concentrated or integrated form, but solely as the dispersed bits of incomplete and frequently contradictory knowledge which all the separate individuals possess."

The knowledge is dispersed. Not concentrated. Not integrated. Dispersed. Across individuals. Each individual holds fragments. The fragments are incomplete — no one has the whole picture. The fragments are frequently contradictory — different people know different things that can't simultaneously be true. The contradiction is not a bug. It is a property of the system.

Hayek was writing about economic planning. His target was the idea that a central planner could allocate resources efficiently by gathering all relevant information and computing the optimal allocation. The information cannot be gathered because it doesn't exist in gather-able form. It exists as local knowledge — the farmer who knows which field drains poorly, the factory manager who knows which machine is about to break, the trader who knows which supplier is becoming unreliable. This knowledge cannot be transmitted to a central planner. It is tacit. It is local. It is constantly changing. The planner's model is always out of date. The market processes information through prices without requiring anyone to understand the whole. The price of copper rises. Users of copper use less copper. Nobody needs to know why the price rose. The price communicates the scarcity. The behavior adapts.

Software teams are Hayekian systems. The knowledge of what the system needs, what it costs, what will break, and what users actually do is dispersed across the team, the codebase, the incident history, the support tickets, and the production metrics. No single person holds it all. The architect who designs the system in advance is the central planner. Their model is out of date. The knowledge they need doesn't exist in the form they need it. It exists as the backend engineer who knows the database migration will take six weeks, the frontend engineer who knows the design system is being rewritten, the SRE who knows the load balancer can't handle the new traffic pattern, the PM who knows the requirements are changing next quarter. None of this knowledge appears in the architecture document. The architecture document is the plan. The plan is wrong.

## The price system of software

Hayek's insight is that markets solve the knowledge problem through prices. Prices are signals. They communicate scarcity without requiring anyone to understand the whole. The price of lumber rises because of a supply disruption in Canada. A furniture maker in Texas uses less lumber. The furniture maker doesn't need to know about the supply disruption. The price tells them everything they need to know: lumber is more expensive. Use less.

Software systems need prices. They don't need literal money. They need signals that communicate scarcity without requiring the consumer to understand the cause. API rate limits are prices. When the rate limit triggers, the caller backs off. The caller doesn't need to know that the service is overloaded because a downstream database is slow. The 429 tells them: this service is scarce right now. Try again later. Queue depths are prices. When the queue grows, the producer slows down. The producer doesn't need to know that the consumer is processing a batch of large messages. The queue depth tells them: consumption is scarce. Circuit breaker states are prices. When the circuit opens, traffic routes elsewhere. The router doesn't need to know why the service is failing. The open circuit tells them: this service is unavailable. Route around it.

These are Hayekian mechanisms. They communicate dispersed knowledge through signals. The signals are local. The response is local. The system adapts without central coordination. The adaptation is the intelligence of the system. The intelligence is not in any component. It is in the signals between components.

## What this means for teams

The knowledge dispersion has consequences for how teams should be structured. A team that makes decisions without consulting the people who hold the relevant local knowledge will make worse decisions. The knowledge exists. It is in the team. It cannot be extracted by a planning process. It must be surfaced by a decision process that involves the people who have it.

Conway's Law is a Hayekian observation. The system mirrors the communication structure because the communication structure determines what knowledge flows where. If the backend team and the frontend team don't talk, the API will be designed without knowledge of how the frontend actually uses it. The API will be clean and wrong. The knowledge of what the frontend needs existed in the frontend team. It didn't flow to the API designers. The API was designed without it. The design is worse than it could have been. The knowledge was dispersed. The communication structure didn't connect it.

The solution is not to eliminate the dispersion. The dispersion is irreducible. The solution is to design mechanisms that surface local knowledge at the point of decision. Code review surfaces the knowledge of the reviewer. Incident postmortems surface the knowledge of the responder. User research surfaces the knowledge of the user. Each mechanism connects dispersed knowledge to a decision that would otherwise be made without it. The mechanisms are not free. They cost time and attention — both scarce. The Hayekian engineer designs mechanisms that surface the most valuable knowledge at the lowest cost. The design is economic. The economics are Hayekian.

## The architect as market designer

The architect who understands Hayek stops trying to design the system. They design the mechanisms by which the system designs itself. Stable interfaces are the prices. Services are the market participants. API contracts are the property rights. Automated testing is the enforcement mechanism. The architect sets the rules. The system evolves within them. The evolution produces outcomes the architect didn't anticipate. The outcomes are better than the architect could have designed because they incorporate knowledge the architect didn't have. The knowledge was dispersed. The market aggregated it.

> "The price system is a mechanism for communicating information. The most significant fact about this system is the economy of knowledge with which it operates." — Friedrich Hayek

The economy of knowledge. The system operates without anyone needing to know the whole. The knowing is local. The responding is local. The coordinating is emergent. This is the architecture of systems that survive complexity. The complexity is in the world. The knowledge of it is dispersed. The system processes the dispersion without centralizing it. The processing is the architecture.

---

**References:**
- Friedrich Hayek, "The Use of Knowledge in Society," *American Economic Review*, Vol. 35, No. 4, September 1945.
- Related posts: [On Scarcity](https://blog.hackspree.com/#scarcity), [I, Pencil](https://blog.hackspree.com/#i-pencil), [Engineering is art and philosophy](https://blog.hackspree.com/#engineering-is-economics)


Engineering is the discipline of building things that work within constraints. Every topic on this blog — operating systems, AI models, trading infrastructure, research labs, innovation economics — is examined through the lens of systems design. The lens is engineering. The method is: understand the constraints, design within them, verify the design works, iterate. The domain provides the specifics. The method is universal.


> The knowledge is dispersed. No single person knows how to build the system. The architect who pretends to is making decisions with incomplete information. The market aggregates the dispersed knowledge through prices. The architecture must aggregate it through interfaces.

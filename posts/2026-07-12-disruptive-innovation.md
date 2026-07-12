---
title: "the electric light was not a better candle"
date: 2026-07-12
slug: disruptive-innovation
summary: "Edison's light bulb did not emerge from gas companies improving their mantles. It emerged from a different axis of value. Disruptive innovation explains why great firms fail — and why most software 'revolutions' are just better candles."
tags: disruptive-innovation, christensen, software-architecture, evolution, design
---

The electric light did not come from the continuous improvement of candles. It did not come from gas companies making better mantles. It did not come from the incremental optimization of existing lighting technology along the dimensions that existing customers valued — brightness, fuel efficiency, cost per lumen-hour.

It came from a different axis of value entirely. Cleanliness. Safety. Convenience. No soot on the walls. No open flame. A simple switch. The early electric light was worse than gas on almost every dimension the gas companies cared about. The bulbs were fragile. They lasted hours, not weeks. There was no distribution infrastructure. The cost was higher. The established customers — factories, street lighting, wealthy households — had no use for a dim, unreliable bulb that required new wiring.

> "Generally, disruptive innovations were technologically straightforward, consisting of off-the-shelf components put together in a product architecture that was often simpler than prior approaches. They offered less of what customers in established markets wanted and so could rarely be initially employed there. They offered a different package of attributes valued only in emerging markets remote from, and unimportant to, the mainstream." — Clayton Christensen, *The Innovator's Dilemma*, 1997

The gas companies did not fail because they were bad at innovation. They failed because they were *good* at it. They improved the mantle. They improved the burner. They optimized for their best customers. They did everything correctly. They were destroyed anyway. This is the innovator's dilemma.

## The theory

Christensen's central insight is a paradox: the practices that make great firms successful — listening to customers, investing in higher-margin products, pursuing sustaining innovation — are the same practices that cause them to fail when a disruptive technology emerges.

> "The way decisions get made in successful organizations sows the seeds of eventual failure."

> "It was as if the leading firms were held captive by their customers."

Incumbents are embedded in a value network. Their customers, suppliers, and investors define what "better" means. The gas companies' customers wanted brighter, cheaper, more reliable light. The gas companies delivered. The electric light companies sold something different — not better light by the existing metrics, but a different kind of value along a different axis. The incumbents had no framework for evaluating it because their customers didn't want it. Their customers were right. The electric light was worse for them. The customers who wanted it didn't exist yet.

> "Products based on disruptive technologies are typically cheaper, simpler, smaller, and more convenient to use."

Cheaper. Simpler. Smaller. More convenient. Not better on the metrics the mainstream values. Better on a different set of metrics for a different set of users. The new users are non-consumers — people who couldn't or didn't use the existing technology. The electric light initially served niches where gas couldn't go. Then it improved. Then it invaded the mainstream from below. By the time the gas companies recognized the threat, the disruption was complete. By 1885, Edison held 75% of the U.S. electric lamp market. The gas mantle was irrelevant. The best gas lighting technology ever made was irrelevant. The sustaining innovations were irrelevant. The axis of value had shifted underneath them.

## The pattern in software

The innovator's dilemma repeats in software with the regularity of Lehman's Laws. Incumbents optimize for existing customers. Disruptors enter with something simpler, cheaper, and worse on the metrics that incumbents care about. Incumbents dismiss the threat. The disruptor improves along a different trajectory. The axis of value shifts. The incumbents are displaced.

### Microservices: was it disruptive?

The microservices movement claimed to be a disruptive innovation. It was mostly a sustaining innovation dressed in disruption's clothes.

The monolith incumbents — large enterprise applications — were being improved along existing axes: faster deployment, better monitoring, more features. Microservices entered with a different value proposition: independent deployability, team autonomy, fault isolation. But most microservices adoptions were not disruptions. They were existing teams taking existing systems and splitting them by existing boundaries — database tables, team structures, Conway's Law made physical. The value proposition didn't change. The deployment got harder. The coupling remained. The complexity moved from in-process to on-network.

True disruption in software architecture would look different. It would serve non-consumers — applications that couldn't be built at all under the old model, users who couldn't afford the old model, problems too small to justify a monolith. Serverless was closer. Dark factories may be closer still. Most microservices migrations were large organizations optimizing for their existing customers — the internal teams — along dimensions the teams already valued. That's sustaining innovation. That's a better gas mantle. The electric light is something else.

### Dark factories: the real disruption?

The dark factory model — specs in, software out, no human-written code, no human-reviewed code — is worse than traditional development on almost every metric incumbents value. The code is less predictable. The quality is unproven. The tooling is immature. The enterprise customers don't want it because it doesn't solve their current problem. Their current problem is shipping features faster within existing architectures. The dark factory doesn't help with that. It solves a different problem: software production for problems where the specs are stable, the patterns are known, and the cost of human development exceeds the value.

That market — the non-consumers of custom software — is invisible to incumbents because their value network doesn't see it. Their customers don't ask for it. Their margins don't support it. Their processes don't allow it. The dark factory, if it follows Christensen's pattern, will improve along its own trajectory — better specs, better validation, better trust — and eventually invade the mainstream from below. The incumbents will have spent decades improving their gas mantles. The electric light will already be installed.

### Agentic coding: better candle or electric light?

> "Disruptive innovations are typically cheaper, simpler, smaller, and more convenient to use."

AI coding assistants — Copilot, Cursor, Claude Code — are initially sustaining innovations. They make existing developers faster at existing tasks along existing dimensions. Better candle. But the trajectory matters. As agents move from assistance to autonomy — from code review (Level 3) to spec-driven development (Level 4) to the dark factory (Level 5) — the axis of value may shift. The metric stops being "developer productivity" and becomes "software production without developers." That is a different axis. That is a different market. That is the electric light.

Most incumbent software organizations will not adopt this. Their best customers — internal product teams — don't want it. Their processes don't support it. Their career ladders don't reward it. Their managers can't measure it. They are held captive by their customers, exactly as Christensen described. They will improve their gas mantles — better linting, better CI/CD, better code review, better agile ceremonies — while the disruption takes root in markets they don't see.

## The connection to Lehman

Lehman's Laws describe the dynamics of E-type software evolution. Systems must change or die. Complexity increases unless you fight it. The process is self-regulating. The laws apply within a paradigm. Disruption is what happens when the paradigm itself changes.

A monolith following Lehman's Laws will evolve, grow in complexity, require increasing maintenance effort, and eventually become uneconomical to sustain. The rational response is a sustaining innovation: refactor, modularize, perhaps migrate to microservices. The disruptive response is a different architecture that changes the axis of evolution entirely. The dark factory doesn't evolve the way a monolith evolves. It evolves by improving specification quality, not by managing codebase complexity. The evolution problem moves from the code to the spec. Lehman's Laws still apply — E-type systems must change or die — but what constitutes an "E-type system" changes. The spec becomes the evolving system. The code becomes a generated artifact. The complexity moves from implementation to specification. The laws didn't change. The substrate did.

## The connection to Parnas

Parnas argued that modules should hide design decisions likely to change. The gas companies could have modularized their business: separate the light production from the light distribution, hide the fuel source behind a stable interface, make the mantle interchangeable with the bulb. If you can switch from gas to electric without changing the customer's experience, you have survived the disruption. The interface is the business. The implementation is the fuel source. The customer doesn't care how the light is made. They care that the light is there when they flip the switch.

Most software organizations don't modularize at this level. They modularize the code — services, libraries, APIs — but not the business architecture. The decision about *what kind of software development to use* is not hidden behind a stable interface. It is embedded in the organizational structure, the hiring pipeline, the promotion ladder, the budgeting process. When the disruption comes, there is no interface to swap out. The organization is the implementation. The implementation is coupled to the paradigm. When the paradigm shifts, the organization cannot follow.

## The connection to Brooks

Brooks argued that conceptual integrity requires one mind. The Reims Cathedral had it. Most cathedrals don't. Most software systems don't. Disruptive innovations almost always come from small teams with conceptual integrity — Edison's lab, the original Macintosh team, the StrongDM dark factory team of three. The incumbent, with its committees and its stakeholders and its sustaining-innovation processes, cannot produce conceptual integrity at the paradigm level. It can only optimize within the paradigm.

> "The central question in how to improve the software art, centers, as it always has, on people." — Brooks

The people who build the disruptive system are not the people who optimized the incumbent system. They don't share the incumbents' value network. They don't listen to the incumbents' customers. They build for a different market along a different axis. The incumbent cannot follow them — not because the technology is hard, but because the organizational structure prevents it. The structure is optimized for the existing paradigm. The structure is the paradigm. Changing the structure means changing who has power, who gets promoted, who decides what "better" means. Disruption is not a technology problem. It is an organizational problem that technology enables.

## The candle-makers were right

> "Blindly following the maxim that managers should keep close to their customers can be a fatal mistake."

The gas companies were right to improve their mantles. Their customers wanted better mantles. Their shareholders rewarded better mantles. Their engineers were trained to make better mantles. The electric light was irrelevant to them until it wasn't. By the time it was relevant, it was too late.

The same logic applies to software. The organizations that are best at monoliths will not lead the transition to microservices. The organizations that are best at manual coding will not lead the transition to dark factories. The organizations that are best at sustaining innovation will not lead the disruption. They are held captive by their customers. Their customers are right — for now. The disruption is happening somewhere else, for someone else, along an axis they don't measure.

The electric light was not a better candle. It was a different thing. Most of what we call innovation in software is better candles. The interesting question is: what is the electric light? And who is building it while we improve our mantles?

---

**References:**
- Clayton M. Christensen, *The Innovator's Dilemma: When New Technologies Cause Great Firms to Fail*, Harvard Business School Press, 1997.
- Related posts on this blog: [Brooks on Software Design series](https://blog.hackspree.com/#brooks-design-conceptual-integrity), [Lehman's Software Evolution](https://blog.hackspree.com/#lehmans-laws), [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Software dark factories](https://blog.hackspree.com/#software-dark-factories), [Henney's Microservices](https://blog.hackspree.com/#kevlin-henney)

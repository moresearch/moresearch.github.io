---
title: "Lewis & Fowler's Microservices: Everyone Copied the Boxes, Nobody Read the Cautions"
date: 2026-07-16
slug: fowlers-microservices-twelve-years-on
summary: 'The March 2014 article by James Lewis and Martin Fowler is the most-cited document in the history of microservices — and one of the most selectively read. Re-read twelve years later, it is barely an architecture paper at all: six of its nine characteristics are organizational claims, and it ends with four cautions that predicted nearly every microservices failure since. This post walks the article as written — the definition, the nine characteristics as one argument, and the ending everyone skipped.'
tags: microservices, architecture, fowler, conways-law, distributed-systems, history
---

On March 25, 2014, [martinfowler.com published "Microservices"](https://martinfowler.com/articles/microservices.html) by James Lewis and Martin Fowler. It did for the term what few documents ever do for an architecture: it fixed the definition. Every conference talk, vendor pitch, and migration proposal since has leaned on it, usually via one sentence — "a suite of small services, each running in its own process and communicating with lightweight mechanisms."

Twelve years on, the article deserves the thing it almost never gets: a full reading. Because read whole, it is a stranger and better document than its reputation. Most of it is not about technology. And its final section — the one with the cautions — reads today like a post-mortem written in advance.

## The definition, and the sentence before it

The famous definition lists the traits: small services, own process, lightweight communication, built around business capabilities, independently deployable by automated machinery, minimal centralized management, polyglot freedom. Teams tattooed that list onto migration decks and started counting services.

But the article's most load-bearing sentence is quieter. Lewis and Fowler define the thing being decomposed: "We consider an application to be a social construction that binds together a code base, group of functionality, and body of funding." Code, function, *funding*. The unit of decomposition was never the process — it was the socio-technical unit that owns the process. Miss that sentence and the rest of the article reads as a deployment pattern. Catch it and the article reads as what it actually is: an organizational design argument with technical consequences.

## Nine characteristics, one argument

The article is structured as nine "common characteristics." Listed flat, they look like a checklist. Grouped, they make a single argument — and the grouping is revealing, because only three of the nine are primarily technical.

**The technical core:**

*Componentization via services.* A component is "a unit of software that is independently replaceable and upgradeable." Libraries componentize in-process; services componentize across processes. The service buys you enforced encapsulation — most languages can't stop an in-process caller from reaching around a Published Interface, but a remote boundary can — and independent deployment. The cost, stated plainly in 2014 and relearned expensively ever since: "remote calls are more expensive than in-process calls," so APIs must be coarser-grained, and moving behavior between components becomes a negotiation instead of a refactor.

*Smart endpoints and dumb pipes.* The anti-ESB principle: domain logic lives in the services; the thing between them just moves messages. Jim Webber's rendering of ESB — "Erroneous Spaghetti Box" — makes the target unmistakable, and Ian Robinson's "be of the web, not behind the web" names the alternative. This is the characteristic the industry would forget fastest; I've traced its lineage back to the 1984 end-to-end argument and forward through the service-mesh relapse in [a companion post](#nats-vs-service-mesh-smart-endpoints-dumb-pipes).

*Decentralized data management.* One database per service, conceptual models allowed to differ, Domain-Driven Design's Bounded Context as the boundary-drawing tool. And the hard part, stated without anesthesia: microservices "emphasize transactionless coordination between services" — eventual consistency, compensating operations, and the business judgment call that the approach "is worth it as long as the cost of fixing mistakes is less than the cost of lost business under greater consistency." That is an economics sentence, not an engineering one. It should have disqualified more adoptions than it did.

**The organizational majority — six of nine:**

*Organized around business capabilities* is Conway's law applied deliberately. The article quotes Melvin Conway's 1968 paper: any organization "will produce a design whose structure is a copy of the organization's communication structure." Layered teams (UI, server, DBA) produce layered architectures where every feature crosses three org boundaries. Cross-functional teams owning a business capability produce services shaped like the business. Note the year of the source: 1968, the same year the [NATO conference named software engineering](#loop-engineering). The two founding observations of the field — feedback loops and org-structure mirroring — are the same age.

*Products not projects* imports Amazon's "you build it, you run it." A team owns its service for the service's lifetime, in "day-to-day contact with how their software behaves in production." No handoff to maintenance, no disbanding at ship.

*Decentralized governance* replaces enforced standards documents with "battle-tested code as libraries" — internal open source, Tolerant Readers, consumer-driven contracts. Netflix, not the enterprise architecture review board, is the named model.

*Infrastructure automation* is the precondition dressed as a characteristic: continuous delivery, automated pipelines, "one of the aims of CD is to make deployment boring." The article is explicit that the teams doing microservices well had this *first*.

*Design for failure* — circuit breakers, bulkheads, timeouts, the Simian Army breaking production on purpose during business hours, dashboards showing business metrics, not just request rates. The sentence that stings: applications must "tolerate the failure of services," which is "an extra complexity compared to a monolithic design." An admission of cost, in the pro-microservices founding document.

*Evolutionary design* drives modularity "through the pattern of change" — things that change together live together — and expects "many services to be scrapped rather than evolved." Services as cattle applies to the services themselves.

Read as one argument: **the technical characteristics exist to make the organizational ones enforceable.** Service boundaries are Conway's law with a compiler. Independent deployment is team autonomy with a pipeline. The database-per-service rule is a bounded context with a firewall. The industry read the article left to right and adopted the mechanisms; the argument runs right to left, from the org design to the mechanisms that protect it.

## The ending everyone skipped

The article closes with a section titled "Are Microservices the Future?" — and the answer given is not yes. It is "cautious optimism," followed by four warnings that map with uncomfortable precision onto the next decade of failed migrations.

**One: it was too early to know.** "The true consequences of your architectural decisions are only evident several years after you made them." In 2014 there were no old microservice systems. The authors said, in their own founding document, that the evidence wasn't in.

**Two: boundaries are hard, and services make boundary mistakes expensive.** In a monolith, a wrong module boundary is a refactor. Across services, "refactoring is much harder than with in-process libraries" — interface coordination, backward-compatibility shims, cross-team negotiation. Get the decomposition wrong and the architecture punishes you for exactly as long as you keep it.

**Three: complexity is conserved.** If your components don't compose cleanly, "all you are doing is shifting complexity from inside a component to the connections between components" — moving it "to a place that's less explicit and harder to control." Every team that traded a debuggable monolith for an undebuggable distributed call graph rediscovered this sentence the hard way, usually without knowing it existed.

**Four: the skill confound.** New techniques are adopted first by stronger teams, so early results overstate what the median team will get. And then the bluntest line in the article: "A poor team will always create a poor system." Microservices don't upgrade the team; they amplify it, in whichever direction it already points.

Fowler spent the following two years expanding these cautions into a small canon — [Monolith First](https://martinfowler.com/bliki/MonolithFirst.html), [Microservice Prerequisites](https://martinfowler.com/bliki/MicroservicePrerequisites.html), [Microservice Trade-Offs](https://martinfowler.com/articles/microservice-trade-offs.html) — essentially footnoting his own article with "we meant the warnings too." The prerequisites piece is three lines long in essence: rapid provisioning, basic monitoring, rapid deployment. Most organizations that failed at microservices failed the prerequisites, not the architecture.

## The 2026 scorecard

What held up, what didn't, from twelve years out:

**Held up:** Bounded contexts as the decomposition tool — the single most durable idea in the article. Conway's law as a design input rather than a lament; the "inverse Conway maneuver" and *Team Topologies* are this characteristic grown into a discipline. "You build it, you run it," which matured into SRE practice and platform engineering. Design for failure — circuit breakers and chaos engineering went from Netflix exotica to table stakes. And the monolith-first caution, vindicated so many times it became the default advice, including the high-profile cases of teams consolidating services back into modular monoliths and publishing the cost savings.

**Aged badly:** The industry's reading, more than the article. "Small" got operationalized as a size contest — hundreds of nanoservices at organizations with a dozen engineers — when the article's unit was the business capability and the funded team. Polyglot freedom became polyglot sprawl. And smart-endpoints-dumb-pipes was inverted wholesale by the sidecar era: the service mesh put routing rules, retry policy, and traffic intelligence back into the pipe, rebuilding the Erroneous Spaghetti Box out of Envoy proxies. The [case that messaging infrastructure should have stayed dumb](#nats-vs-service-mesh-smart-endpoints-dumb-pipes) is, at bottom, just characteristic four of this article, re-argued against its newest violation.

**The unresolved part:** the skill confound never resolved, because it can't. Twelve years of retrospectives — successes at organizations with elite platform teams, failures at organizations without them — are exactly the distribution the article predicted from selection effects alone. We still don't have clean evidence about what microservices do for the *median* team, and the founding document told us we wouldn't.

## How to read it now

Read it as an organizational design paper wearing an architecture paper's clothes. The test it implies is not "how many services do you have" but three questions in order: Do your service boundaries match funded, long-lived, cross-functional teams? Could each of those teams deploy independently today, with boring deployments, tolerating the others' failures? And have you priced the loss of transactions and in-process refactoring against what independence buys you?

Answer those honestly and the article has done its job — whichever architecture you end up with. That was always the strange virtue of the founding document of microservices: it is one of the few manifestos in software history that argues against its own cargo cult, in the text, from day one. The boxes were copied a million times. The sentences are still waiting.

## References

1. James Lewis and Martin Fowler, [Microservices](https://martinfowler.com/articles/microservices.html) (2014) — the article itself.
2. Melvin Conway, [How Do Committees Invent?](https://www.melconway.com/Home/Committees_Paper.html), Datamation (1968) — source of Conway's law.
3. Martin Fowler, [Monolith First](https://martinfowler.com/bliki/MonolithFirst.html) (2015).
4. Martin Fowler, [Microservice Prerequisites](https://martinfowler.com/bliki/MicroservicePrerequisites.html) (2014).
5. Martin Fowler, [Microservice Trade-Offs](https://martinfowler.com/articles/microservice-trade-offs.html) (2015).
6. Eric Evans, *Domain-Driven Design* (2003) — source of Bounded Context.
7. Sam Newman, *Building Microservices*, 2nd ed. (O'Reilly, 2021).
8. Michael Nygard, *Release It!*, 2nd ed. (Pragmatic Bookshelf, 2018) — circuit breakers, bulkheads, timeouts.
9. Werner Vogels, [A Conversation with Werner Vogels](https://queue.acm.org/detail.cfm?id=1142065), ACM Queue (2006) — "you build it, you run it."
10. Netflix Tech Blog, [The Netflix Simian Army](https://netflixtechblog.com/the-netflix-simian-army-16e57fbab116) (2011).
11. Ian Robinson, [Consumer-Driven Contracts](https://martinfowler.com/articles/consumerDrivenContracts.html) (2006).
12. Matthew Skelton and Manuel Pais, *Team Topologies* (IT Revolution, 2019) — Conway's law operationalized.
13. J.H. Saltzer, D.P. Reed, D.D. Clark, [End-to-End Arguments in System Design](https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf) (1984) — the pre-history of dumb pipes.

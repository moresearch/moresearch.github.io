---
title: The Unix philosophy is the only software engineering theory that works
date: 2026-07-12
slug: unix-philosophy
summary: "McIlroy wrote it in 1978. Kernighan and Pike explained it in 1984. Raymond codified it in 2003. Microservices rediscovered it in 2014. Nobody read the books. The pipes are the same. The mistakes are the same."
tags: unix, philosophy, pipes, tools, composition, microservices, software-design
---

In 1978, Doug McIlroy — the inventor of the Unix pipe — wrote down the Unix philosophy in the Bell System Technical Journal. It consisted of four directives:

1. Make each program do one thing well. To do a new job, build afresh rather than complicate old programs by adding new features.
2. Expect the output of every program to become the input to another, as yet unknown, program. Don't clutter output with extraneous information. Avoid stringently columnar or binary input formats. Don't insist on interactive input.
3. Design and build software, even operating systems, to be tried early, ideally within weeks. Don't hesitate to throw away the clumsy parts and rebuild them.
4. Use tools in preference to unskilled help to lighten a programming task, even if you have to detour to build the tools and expect to throw some of them out after you've finished using them.

He later condensed it to three lines that became famous:

> "This is the Unix philosophy: Write programs that do one thing and do it well. Write programs to work together. Write programs to handle text streams, because that is a universal interface."

That's it. The entire theory of software composition, stated in 1978, in three sentences. Everything that has happened in software architecture since — microservices, serverless, event-driven systems, hexagonal architecture, domain-driven design, bounded contexts, API gateways, service meshes — is rediscovery of these three sentences in new terminology. The terminology changes. The principles don't. The books are still on the shelf. The ignorance is expensive.

## The books you should have read

**Kernighan and Pike, *The Unix Programming Environment* (1984).** The canonical text. They wrote:

> "What makes it effective is the approach to programming, a philosophy of using the computer. At its heart is the idea that the power of a system comes more from the relationships among programs than from the programs themselves. Many UNIX programs do quite trivial things in isolation, but, combined with other programs, become general and useful tools."

The power is in the relationships, not the programs. The programs are trivial. The composition is where the capability lives. This is the sentence that should be printed above every microservices whiteboard. It is also the sentence most microservices architectures violate, because they focus on the services — what each service does, what database it owns, what team builds it — and ignore the relationships. The relationships are the system. The services are components. Components are easy. Relationships are hard.

**Eric Raymond, *The Art of Unix Programming* (2003).** Raymond codified seventeen rules. The ones that matter for this argument:

- **Rule of Modularity:** Build programs from simple, cleanly-connected parts.
- **Rule of Composition:** Programs must communicate easily with other programs.
- **Rule of Separation:** Separate mechanism from policy.
- **Rule of Parsimony:** Write small programs. Easy to replace when wrong.
- **Rule of Transparency:** Make operation visible and discoverable.
- **Rule of Silence:** Don't print unnecessary output. Let other programs decide what matters.

Each rule maps directly to a microservices principle with the names changed. Modularity → bounded contexts. Composition → API contracts. Separation → business logic vs. infrastructure. Parsimony → small services, easy to rewrite. Transparency → observability. Silence → don't log everything, emit meaningful events. Raymond wrote this in 2003, before microservices existed as a term. He was describing Unix. He was also describing microservices. He didn't know it. The microservices pioneers didn't know it either. They thought they were inventing something. They were rediscovering pipes.

**McIlroy's later reflection.** Years after the original paper, McIlroy watched Linux grow and said:

> "Everything was small. My heart sinks for Linux when I see the size of it. We used to sit around in the Unix Room saying, 'What can we throw out? Why is there this option?' It's often because there is some deficiency in the basic design — you didn't really hit the right design point. Instead of adding an option, think about what was forcing you to add that option."

The option is the symptom. The design deficiency is the cause. Adding the option papers over the deficiency without fixing it. The program grows. The option count grows. The composability shrinks. McIlroy's test: when you're tempted to add a feature, ask what deficiency in the existing design made the feature necessary. Fix the design. Don't add the feature. This test applies to microservices boundaries. When you're tempted to add a new endpoint to a service, ask what deficiency in the existing API contracts made the new endpoint necessary. Fix the contract. Don't add the endpoint. The endpoint is the option. The contract deficiency is the design flaw.

## Why pipes work: the engineering fundamentals

The Unix pipe is not a metaphor. It is a specific engineering construct with specific properties that make composition possible.

**Uniform interface.** Every program reads from `stdin` and writes to `stdout`. The interface is identical regardless of what the program does. `grep` reads text and writes text. `sort` reads text and writes text. `wc` reads text and writes text. They can be connected in any order, in any combination, because the interface is uniform. This is the opposite of REST microservices, where every service defines its own endpoints, its own request format, its own response format, its own error semantics. The interface is not uniform. Composition requires adapters. Adapters are coupling. Coupling is the thing pipes eliminate.

**Separation of mechanism and policy.** `sort` sorts. It does not know what it is sorting. It does not know why. It does not know what will happen to the sorted output. It sorts. That is the mechanism. The policy — what data to sort, what to do with the sorted result — is determined by the programs upstream and downstream. This is the Rule of Separation applied to data processing. In microservices: the service provides the mechanism (process an order, reserve inventory). The orchestration layer provides the policy (which services to call in which order under which conditions). If the service encodes policy, it cannot be reused in a different policy context. The mechanism is coupled to the policy. The coupling prevents composition.

**Composability without coordination.** `grep` was written before `sort` knew about it. `sort` was written before `wc` knew about it. None of them were designed to work together. They work together because they all obey the same interface contract. They can be composed into pipelines the original authors never imagined. This is the property that microservices promise and rarely deliver: composition without coordination. Services that were designed by different teams at different times, communicating through stable interfaces, composed into workflows that nobody designed in advance. The promise is real. The delivery is rare because the interfaces are not uniform. Every service defines its own contract. Every composition requires a new adapter. The adapters accumulate. The system becomes a collection of adapters connecting services that were supposed to be directly composable.

**Filter thinking.** McIlroy's second directive: "Expect the output of every program to become the input to another, as yet unknown, program." Every program is a filter — it transforms an input stream into an output stream. It doesn't know where the input came from. It doesn't know where the output is going. It transforms. This is the purest form of Parnas's information hiding: the program hides everything about itself except the transformation it performs. The caller doesn't know the algorithm. The caller doesn't know the implementation language. The caller knows the transformation. The transformation is the interface. Everything else is hidden.

## Where pipes fail

The Unix philosophy is not universal. It has known failure modes. Understanding them is as important as understanding the successes.

**Text as universal interface breaks at scale.** Text is universal. Text is also unstructured. Every program that receives text must parse it. Every parsing step is an opportunity for error, inconsistency, and performance cost. When the data has structure — nested objects, typed fields, relationships — text streams force every consumer to reconstruct the structure from its flattened representation. This is the argument for typed interfaces, for gRPC over REST, for Avro over JSON. Text is the universal interface for simple data. It is the wrong interface for complex data. The Unix philosophy doesn't tell you where the boundary is. Experience does. The boundary moves with the complexity of the data.

**State management is externalized.** Pipes connect stateless programs. Each program reads, transforms, writes. State is managed outside the pipeline — in files, databases, or the shell's variables. When the processing requires state that spans multiple pipeline stages — a running total, a windowed aggregation, a session — the stateless model breaks. You either pass the state through the pipe as additional data (cluttering the output, violating the Rule of Silence) or you externalize it (breaking the pipeline model). Modern stream processing systems — Kafka Streams, Flink, Spark Streaming — are essentially pipelines with built-in state management. They fix the failure mode at the cost of increased complexity. The tradeoff is inevitable. The Unix philosophers knew it. They never claimed pipes solved everything.

**Error handling is ad-hoc.** A pipeline of ten programs. The eighth program fails. What happens? The shell reports the exit code of the last program in the pipeline. Unless you use `set -o pipefail`, the failure of the eighth program is invisible. The pipeline continues. The output is partial. Nobody knows something went wrong. This is a design choice, not an oversight. Unix errs on the side of simplicity: programs should do their job and exit. Error handling is the caller's responsibility. But when the caller is a pipeline, the caller is distributed across multiple programs, none of which know about each other. Distributed error handling is hard. Microservices rediscovered this the hard way. Distributed sagas, compensating transactions, dead letter queues — these are the modern equivalents of `pipefail`. The terminology changed. The problem is the same.

**The composition model is linear.** Pipes compose programs sequentially. The output of A goes to the input of B. This is powerful for linear data processing. It is weak for systems where the data flow is a graph — fan-out, fan-in, conditionals, loops, feedback. You can build graph processing in shell, but the shell fights you. The composition model is not general. It is linear. Most workflows are not linear. The Unix philosophy works brilliantly for the subset of problems that are linear data transformations. For everything else, you need a different composition model. Microservices with message brokers — NATS, Kafka, RabbitMQ — implement graph composition. The broker is the compositor. The services are the components. The graph is the architecture. The Unix philosophy didn't fail. It was extended. The extension was necessary.

## Microservices: the old is the new new

In 2014, the term "microservices" entered the mainstream. The defining characteristics: small, focused services. Communication through uniform interfaces. Independent deployability. Composition into workflows. Decentralized data management. Design around business capabilities.

This is McIlroy's Unix philosophy, restated for distributed systems. Small, focused services → programs that do one thing well. Uniform interfaces → text streams as universal interface. Composition into workflows → pipes connecting programs. Independent deployability → programs that don't know about each other. Decentralized data management → each program manages its own state.

The microservices pioneers were not copying Unix. They were independently rediscovering the same principles at a different scale. The scale changed. The principles didn't. The oversight is that the Unix philosophers already documented the failure modes — and the microservices pioneers walked into every one of them.

**The uniform interface failure.** Unix has a uniform interface: text streams. Every program reads and writes the same format. Microservices initially attempted the same: REST with JSON. But JSON is not a uniform interface when every service defines its own schema, its own endpoint structure, its own error format. The interface is HTTP. The contract is ad-hoc. The uniformity is at the transport layer. The diversity is at the application layer. The diversity is where the coupling lives. Microservices adopted the transport uniformity of Unix without the semantic uniformity. The result is services that can talk to each other but can't understand each other without adapters. The adapters are the coupling. The coupling is what Unix pipes eliminated.

**The composition failure.** Unix composes programs with pipes: `A | B | C`. The composition is linear, immediate, and visible. Microservices compose with orchestration: Service A calls Service B, which calls Service C, with retries, timeouts, circuit breakers, and dead letter queues at each step. The composition is a distributed graph with failure modes at every edge. The complexity of composition in microservices is orders of magnitude higher than in Unix pipelines. The principles are the same. The implementation is harder. The failure to acknowledge the increased difficulty is why microservices projects fail.

**The state management failure.** Unix programs are stateless. State lives in the filesystem. Microservices are stateful. Each service owns its database. State management is decentralized. Decentralized state management is a hard problem — distributed transactions, eventual consistency, saga patterns, CQRS, event sourcing. Unix didn't have this problem because Unix programs didn't own state. Microservices created the problem by making services own state, then spent a decade inventing patterns to solve it. The patterns work. They are also complex. The complexity is inherent. It cannot be eliminated by better tooling. It can only be managed by accepting that decentralized state is expensive and choosing which services truly need it.

**The debugging failure.** When a Unix pipeline fails, you can run each program in isolation with the same input and see where the output diverges. The pipeline is reproducible. Microservices are not. A failure in a distributed workflow involves network timeouts, retry policies, eventual consistency windows, and state spread across multiple databases. Reproducing the failure requires reproducing the entire distributed state. This is hard. This is why observability tooling for microservices is a multi-billion-dollar industry. The tooling exists because the problem is hard. The problem is hard because the composition model is distributed. The composition model is distributed because each service owns its state. The state ownership is the root cause. The root cause was a design choice. The design choice had consequences nobody predicted — except McIlroy, who designed Unix with stateless programs, and Pike, who built Go with channels instead of shared state, and the Unix philosophers generally, who understood that state is the enemy of composability.

**The size failure.** McIlroy's lament about Linux — "my heart sinks when I see the size of it" — applies directly to microservices. A service that does one thing is small. A service that does one thing plus error handling, retry logic, circuit breaking, authentication, authorization, logging, metrics, tracing, configuration management, and service discovery is not small. The infrastructure concerns colonize the service. The service grows. The growth is not in business logic. It is in infrastructure. The infrastructure should be external to the service — in the platform, the mesh, the gateway. But it leaks in. It leaks in because the uniform interface that should handle these concerns — the service mesh, the API gateway — is not as uniform as Unix text streams. Every service has its own configuration, its own policies, its own exceptions. The uniformity is incomplete. The incompleteness is where the complexity accumulates.

## Why history repeats

The Unix philosophers wrote their books between 1978 and 2003. The microservices pioneers wrote their blog posts between 2012 and 2016. The gap is roughly a decade. The principles are the same. The terminology changed. Why didn't the microservices pioneers cite McIlroy, Kernighan and Pike, and Raymond?

Because software engineering does not read its own history. Architects study buildings. Composers study scores. Software engineers study the framework documentation for the current version. The old books are on the shelf. The old papers are in the Bell System Technical Journal. The old principles are correct. Nobody reads them. Every generation rediscovers composition, modularity, and information hiding and gives them new names. The names change. The principles don't.

The cost of this amnesia is not theoretical. It is measured in failed microservices migrations. Teams that split the monolith by database table instead of by bounded context. Teams that built a distributed system with more coupling than the monolith it replaced. Teams that discovered that twelve services with a complete call graph is not an architecture — it is a monolith with network latency. Teams that learned about distributed state management the hard way, in production, at 3am, when the saga pattern failed and the compensating transaction didn't compensate.

Every one of these failures was avoidable. Not by reading the microservices literature. By reading the Unix literature. The principles are older. The principles are clearer. The principles were stated in three sentences in 1978. The sentences are still correct. The sentences are still unread.

> "Those days are dead and gone and the eulogy was delivered by Perl." — Rob Pike

Pike's eulogy was premature. The philosophy didn't die. It moved up the stack. Pipes became channels. Text streams became typed interfaces. Small programs became small services. The shell became the orchestrator. The toolbox became the service catalog. The principles survived. The implementation changed. The failure modes are the same. The books are still on the shelf. Read them.

---

**References:**
- Doug McIlroy, "Unix Time-Sharing System: Forward," *Bell System Technical Journal*, Vol. 57, No. 6, July-August 1978.
- Brian Kernighan and Rob Pike, *The Unix Programming Environment*, Prentice-Hall, 1984.
- Eric S. Raymond, *The Art of Unix Programming*, Addison-Wesley, 2003.
- Peter H. Salus, *A Quarter-Century of Unix*, Addison-Wesley, 1994.
- Rob Pike, "Simplicity is Complicated," dotGo, 2015.
- Related posts: [Henney's Microservices](https://blog.hackspree.com/#kevlin-henney), [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Git is a Unix tool](https://blog.hackspree.com/#git-unix-philosophy), [NATS pub/sub beats REST](https://blog.hackspree.com/#nats-pubsub-microservices)

---
title: Durable Daemons
date: 2026-07-25
slug: durable-daemons
summary: This post coins the term "durable daemons" — AI agents that persist across sessions, accumulate state, act autonomously, and survive process boundaries. We define the term, trace its lineage from Maxwell's demon through BSD's Beastie to the always-on agent literature, and argue that durable execution runtimes (DBOS, Temporal) are the missing infrastructure layer. The economic case is not subtle: a daemon that forgets its commitments costs more than one that doesn't.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Richard Stallman's ITS operating system at MIT had a program called DAEMON — a background process that watched for new files, woke up when it found them, and performed an action. The name came from Maxwell's demon, a thought experiment in thermodynamics: an imaginary being that sits between two chambers, observes molecules, and selectively opens a gate. The demon does not create energy. It uses information to direct it. The operating system daemon operates the same way. It waits. It observes. It acts on what it observes. It is persistent, stateful, and autonomous within its narrow domain.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to.

## Coining the term: what is a durable daemon?

We are introducing the term **durable daemon** in this post. It needs a precise definition, because without one the phrase is just vibe. Here it is.

> A **durable daemon** is an AI agent that satisfies four conditions:
>
> 1. **Persistence.** It runs across sessions, not as a one-shot request-response loop. Its process may restart, but its identity and accumulated state survive.
> 2. **Stateful memory.** Its future behavior depends on durable state accumulated across earlier interactions — not just a vector database of past chats, but task ledgers, commitments, permissions, provenance records, and trigger conditions.
> 3. **Autonomous action.** It acts on its state without being prompted. It maintains queues, watches conditions, fires triggers, makes commitments, and discharges them. It is not a tool the user invokes. It is a process that invokes itself when its state crosses a threshold.
> 4. **Crash-proof execution.** Its workflows survive process death, machine reboot, and database failover. Completed steps are never re-executed. In-flight steps resume from the last checkpoint. Its decisions leave an audit trail by construction.
>
> Conditions 1–3 describe what a daemon *is*. Condition 4 is what makes it *durable*. An AI agent that satisfies 1–3 but not 4 is an always-on agent. An AI agent that satisfies all four is a durable daemon.

Why coin a new term? Because the thing we are building does not have a name. "Always-on agent" (from Ding et al.) captures persistence and state but says nothing about crash recovery or auditability. "AI agent" is too broad — it includes stateless chatbots. "Daemon" alone is too narrow — it means a Unix background process with a PID file and a config in `/etc`. We need a term that captures the synthesis: the Unix daemon's persistence and autonomy, married to the AI agent's learned state and reasoning, grounded on the durable execution runtime that keeps it alive. That term is *durable daemon*.

The rest of this post justifies the definition. We trace the lineage — Maxwell's demon, BSD's Beastie, the always-on agent survey — and then we argue that durable execution is the missing layer that turns an always-on agent into a durable daemon.

Almost fifty years later, the daemon metaphor is re-emerging in a new substrate. In June 2026, a survey by Ding, Nannapaneni, Liu, and Zhang — [*Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents*](https://arxiv.org/abs/2606.30306) — mapped 435 papers into the first systematic taxonomy of what happens when LLM-based agents stop being one-shot tools and become persistent, stateful processes. The paper does not use the word *daemon*. It does not need to. The thing it describes is a daemon. (I covered the paper's framework in detail [here](https://blog.hackspree.com/#always-on-agents).)

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for the cover of The Design and Implementation of the 4.3BSD Operating System. The trident symbolizes the fork(2) system call. The tennis shoes are unexplained but perfect.](/images/bsd-daemon-medium.gif)

> An always-on agent is a daemon with memory. A durable daemon is an always-on agent that cannot be killed by a deploy. The memory is the difference between a daemon and a service. The durability is the difference between a daemon and a reliable one.

## The BSD daemon as cultural artifact

The daemon concept acquired its cultural form through BSD Unix, but the story is weirder and better than most people know. It involves a locksmith, a Pixar founder, and a near-loss of the intellectual property to "a certain large company."

We love BSD over Linux for one reason that matters here — we being the hackers, the daemon-writers, the ones who still read man pages: BSD is designed as a whole system, one source tree, one team, one coherent design, and its daemons are first-class citizens of that design, not afterthoughts assembled from separate projects by a distribution maintainer.

### The etymology

The word *daemon* was chosen deliberately. It is not *demon* — an evil spirit. It is the ancient Greek δαίμων (*daimōn*) — a deified being, a guardian spirit, an intermediary between mortals and gods. Socrates spoke of his daimonion as an inner voice that warned him when he was about to make a mistake. The operating system daemon operates the same way: it watches, it waits, and it acts in the user's interest without the user needing to command it. Marshall Kirk McKusick, the copyright holder of the BSD Daemon and a key figure in BSD's history, is emphatic about the distinction. The daemon is a helper, not a threat. The pitchfork is not a weapon. It is the trident of Poseidon — and in the Unix context, it symbolizes the `fork(2)` system call, the mechanism by which a daemon creates a child process to do its work.

### The three artists

The BSD Daemon has no single creator. It was drawn by three different artists across twelve years, each building on what came before.

**Phil Foglio (1976).** The BSD Daemon was first drawn by the comic artist Phil Foglio — not for an operating system project, but as payment for a favor. Mike O'Brien, a developer working as a bonded locksmith, cracked a wall safe in Foglio's Chicago apartment after a roommate disappeared without leaving the combination. In return, Foglio agreed to draw T-shirt artwork. O'Brien provided Polaroid photos of a PDP-11 running UNIX and some visual puns: pipes, daemons, forks, `/dev/null`. Foglio produced a drawing of four cheerful red daemons with tridents, climbing on and falling off water pipes in front of a PDP-11. It debuted at the first national UNIX meeting in Urbana, Illinois. Bell Labs bought dozens of the T-shirts. The original artwork was lost after being sent to Digital Equipment Corporation — a casualty of corporate mailrooms.

**John Lasseter (1984 and 1988).** John Lasseter — who would later found Pixar and direct *Toy Story*, *A Bug's Life*, and *Cars* — was working at Lucasfilm in 1984 when he drew a greyscale BSD Daemon for the cover of the *Unix System Manager's Manual* published by USENIX for 4.2BSD. Four years later, in 1988, he produced the definitive version: the iconic red, horned daemon with the trident and tennis shoes, drawn for the cover of *The Design and Implementation of the 4.3BSD Operating System*, co-authored by Marshall Kirk McKusick. This is the Beastie the world knows. The tennis shoes remain unexplained. Lasseter has never publicly addressed them.

**Marshall Kirk McKusick (steward, 1988–present).** McKusick did not draw the daemon, but he is the reason it survived. He obtained the copyright from Lasseter in exchange for payment for the drawings and has held it ever since. In the early 1990s, he nearly lost the rights to "a certain large company" — he has never named them publicly — because he failed to show due diligence in protecting the mark. The experience turned him into a careful steward. He requires advance permission for any quantity reproduction. He restricts use to BSD-related contexts. He refuses to place the daemon under a Creative Commons license, preferring to personally ensure it is used appropriately. The daemon survived because McKusick refused to let it become clip art.

### Why FreeBSD adopted it

FreeBSD used Lasseter's 1988 daemon as both logo and mascot from the project's founding in 1993 until 2005, when it introduced a new scalable wordmark logo. Beastie remained the official project mascot. FreeBSD is the only major BSD variant that still actively uses the daemon as part of its identity — NetBSD and OpenBSD eventually adopted abstract logos (a flag and Puffy the blowfish, respectively).

The reason FreeBSD kept the daemon is not nostalgia. It is that the daemon captures something true about what an operating system is. A daemon is not a service. A service answers requests and forgets them. A daemon *maintains* something — a queue, a schedule, a watch on a directory, a log of events — and acts when its internal state crosses a threshold. The difference is persistence across individual requests. A daemon's future behavior depends on its accumulated state. This is the BSD design philosophy in miniature: the system is not a collection of tools called by the user. It is a set of persistent, autonomous processes that maintain the system's state and act on the user's behalf whether the user is watching or not.

This is also the paper's definition of an always-on agent: a system "whose future behavior depends on durable state accumulated across earlier interactions." An AI daemon is an always-on agent. Not a chatbot. Not a tool you invoke. A process that persists, remembers, and acts.

## The paper, briefly: conditions 1–3 are mapped

I wrote about the Ding et al. survey [at length in July](https://blog.hackspree.com/#always-on-agents). The short version: across 435 papers, the research concentrates on accumulating and retrieving state. Governance, recovery, and forgetting are neglected. The authors provide six diagnostic axes for reasoning about agent state — Authority, Scope, Mutability, Provenance, Recoverability, Actionability — and a lifecycle (write, validate, organize, retrieve, act, update, forget, audit, rollback). The headline finding is that we are good at making agents remember and bad at governing what they remember. The governance gap is where the risk lives.

For this post, the point is simpler. The paper maps conditions 1–3 without using the word *daemon*. The six diagnostic axes are a daemon design checklist. The state lifecycle is what every daemon must do with its memory. And the governance gap — the neglected work of auditing, recovering, and forgetting — is exactly the work that condition 4 exists to handle. The paper describes what must be governed. Durable execution provides the substrate on which governance can be built.

## Durable execution: condition 4, the missing layer

The paper identifies the problems — it maps conditions 1–3 thoroughly. It does not prescribe the runtime for condition 4. But there is an emerging class of infrastructure that directly provides it — not from the AI research community, but from the database and distributed systems community. It is called *durable execution*.

The idea is simple. When a function executes as a durable workflow, every step is checkpointed to a durable store — typically Postgres. If the process crashes mid-workflow, the system recovers by reading the last successful checkpoint and resuming from that point. Completed steps are never re-executed. The workflow survives process death, machine reboot, and database failover.

Two systems define the category. **Temporal**, founded in 2019 and battle-tested at DoorDash, Snap, Stripe, and Uber, provides a centralized orchestration server with polyglot SDKs and the operational maturity that comes from years of production use at scale. **DBOS**, founded in 2023 by database pioneer Mike Stonebraker, takes the opposite approach: an open-source library that checkpoints directly to Postgres tables with no orchestration server, no task queue, no sidecar. Step latency is 1–2 milliseconds — a single Postgres write. The philosophy: "Postgres is the operating system."

> Durable execution solves the recoverability axis directly. It provides the audit trail as a byproduct of checkpointing. It makes rollback tractable by making every state transition explicit and replayable. It is not the whole answer to the governance gap. It is the foundation on which the rest of the governance stack can be built.

### The economic case

An AI daemon that books a meeting, sends an email, updates a CRM record, and files a Jira ticket is executing a multi-step workflow with external side effects. If the process crashes after step two, what happens? Does the email get sent twice? Does the meeting get booked but never appear in the CRM? Does the agent start over from the beginning and double-book the meeting?

These are not hypothetical questions. Databases solved them decades ago with transactions. Workflow engines solved them with checkpointing and compensation. AI daemons, as of mid-2026, have largely not adopted either.

Consider an agent that processes customer onboarding: verify identity, create account, provision access, send welcome email, schedule follow-up call. If the agent crashes after provisioning access but before scheduling the call, the customer has access but no follow-up. The cost of the dropped step is a churned customer — or an engineer spending two hours manually reconciling the state, which costs as much as several thousand agent invocations. The durable execution layer pays for itself the first time it prevents a dropped step.

The economics compound at scale. A sales AI daemon maintaining a pipeline of two hundred deals, each with commitments it made to prospects — "I'll send you the pricing sheet by Tuesday," "I'll check with engineering and get back to you" — is a daemon with two hundred concurrent workflows, each with state that must survive process restarts. Without durable execution, every deploy, every crash, every autoscaling event is a risk that commitments will be dropped. The daemon becomes unreliable not because the model is bad but because the runtime is fragile. The trust cost is unquantifiable but real.

### Audit as a byproduct

The second economic advantage is less obvious but potentially larger: auditability. When every step of an AI daemon's execution is checkpointed to Postgres, the audit trail exists by construction. You can query why the daemon sent a particular email — trace the workflow, inspect the state at each step, determine what it knew when it made the decision. In regulated industries — finance, healthcare, legal, insurance — auditability is not optional. Building that audit trail from scratch is expensive. Getting it as a byproduct of the execution model is close to free.

### Which one?

DBOS is the library answer: add a dependency, annotate your functions, own your data by default. Temporal is the platform answer: polyglot workers, massive fan-outs, month-long workflows, Stripe-scale maturity. Both are better than the status quo — agents running as stateless request-response loops with ad-hoc persistence bolted on. The choice between them matters less than the choice to satisfy condition 4 at all. Without it, you have an always-on agent. With it, you have a durable daemon.

> A durable daemon that forgets its commitments is worse than no daemon at all. The user who cannot trust the daemon to remember what it promised will stop delegating to it. Condition 4 — crash-proof execution — is not a luxury. It is what separates a daemon you can rely on from one you cannot.

## The durable daemon layer

The first generation of AI agents were stateless: prompt in, response out, no memory between calls. The second generation added retrieval: find relevant documents, ground responses in provided context. The third generation — the one the Ding et al. survey maps — added persistence: the agent maintains state across interactions, learns from it, acts on it. Those are conditions 1–3. They describe an always-on agent. They do not describe a durable daemon.

The fourth generation adds condition 4: durability. State survives process boundaries. Decisions leave audit trails. Commitments are recoverable. Failures are compensatable. The agent does not just persist — it cannot be killed by a deploy, a crash, or an autoscaling event. It is crash-proof. When all four conditions hold, we have a durable daemon.

Each generation reduces the distance between what an agent is and what a Unix daemon has always been. Beastie was persistent, watchful, autonomous. But his state was a configuration file and a connection table. The durable daemon's state is everything it has ever seen, inferred, committed to, or been told. The durability problem is qualitatively different. Durable execution — DBOS, Temporal, and whatever comes next — is the first infrastructure layer built to handle it. It is condition 4, made real.

Beastie, in his tennis shoes and grin, has been running in the background since 1976. The durable daemons are larger, smarter, and stranger. They need the definition. They need the foundations. This post gives them the first. DBOS and Temporal give them the second.

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- Earlier: [Always-on agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents) — Detailed coverage of the survey's six axes, governance gap, and state lifecycle.
- [DBOS Documentation](https://docs.dbos.dev/) — Durable execution as a Postgres-backed library. Founded by Mike Stonebraker.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution as a platform. Battle-tested at DoorDash, Snap, Stripe, Uber.
- McKusick, M.K. [Beastie's Home Page](http://www.mckusick.com/beastie/index.html) — Copyright holder and steward of the BSD Daemon since 1988.
- [FreeBSD Foundation: Interview with Beastie, the BSD Daemon](https://freebsdfoundation.org/blog/freebsd-day-interview-with-beastie-the-bsd-deamon/) — The definitive history of the mascot.
- [BSD Daemon (Wikipedia)](https://en.wikipedia.org/wiki/BSD_Daemon) — The three artists (Foglio, Lasseter, McKusick), the etymology, and the near-loss of the copyright.
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux) — The design philosophy behind the operating systems that gave daemons their cultural form.
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) — A desktop AI agent that is, in effect, an AI daemon running on the user's machine.

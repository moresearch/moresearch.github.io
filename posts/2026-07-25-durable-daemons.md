---
title: Durable Daemons
date: 2026-07-25
slug: durable-daemons
summary: Always-on AI agents are the daemons of the AI era — persistent, stateful, autonomous. A new survey of 435 papers maps the emerging discipline of agent durability. The research concentrates on accumulation and retrieval; governance, recovery, and forgetting are neglected. Meanwhile, durable execution runtimes like DBOS and Temporal offer a path to daemons that survive crashes, audit their decisions, and earn their keep. The economic case for durability is not subtle: a daemon that forgets its commitments costs more than one that does not.
tags: daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Richard Stallman's ITS operating system at MIT had a program called DAEMON — a background process that watched for new files, woke up when it found them, and performed an action. The name came from Maxwell's demon, a thought experiment in thermodynamics: an imaginary being that sits between two chambers, observes molecules, and selectively opens a gate. The demon does not create energy. It uses information to direct it. The operating system daemon operates the same way. It waits. It observes. It acts on what it observes. It is persistent, stateful, and autonomous within its narrow domain.

Almost fifty years later, the daemon metaphor is re-emerging in a new substrate. In June 2026, a survey by Ding, Nannapaneni, Liu, and Zhang — [*Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents*](https://arxiv.org/abs/2606.30306) — mapped 435 papers into the first systematic taxonomy of what happens when LLM-based agents stop being one-shot tools and become persistent, stateful processes. The paper does not use the word *daemon*. It does not need to. The thing it describes is a daemon.

![The BSD Daemon — Beastie, drawn by John Lasseter and refined by Marshall Kirk McKusick. A small red creature with a pitchfork, tennis shoes, and a grin: not malevolent, but helpful, watchful, and slightly mischievous. The mascot that gave daemons their cultural form.](/images/bsd-daemon.jpg)

> Always-on agents are daemons with memory. The memory is the difference. A traditional daemon's state is configuration and connection state — small, explicit, bounded. An always-on agent's state is everything it has ever seen, inferred, decided, or been told. The surface area for error expands with the state. The economic consequences expand with it.

## The BSD daemon as cultural artifact

The daemon concept acquired its cultural form through BSD Unix. The BSD Daemon — Beastie, originally drawn by John Lasseter in 1976 and later refined by Marshall Kirk McKusick — is a small red creature with a pitchfork, tennis shoes, and a grin. He is not malevolent. He is helpful, watchful, and slightly mischievous. He lives in the background of the system, doing things the user does not see, until the user needs them done. He is the mascot of FreeBSD, NetBSD, and OpenBSD — the operating systems that kept the daemon tradition alive while Linux took the world.

The pitchfork is the telling detail. It is not a weapon. It is a tool for digging — into process state, into system internals, into the details the user should not have to think about. A daemon *maintains* something. It is not a service. A service answers requests and forgets them. A daemon maintains a queue, a schedule, a watch on a directory, a log of events — and acts when its internal state crosses a threshold. The difference is persistence across individual requests. A daemon's future behavior depends on its accumulated state.

This is also the paper's definition of an always-on agent: a system "whose future behavior depends on durable state accumulated across earlier interactions." The definition is deceptively simple. It says nothing about LLMs, about agents, about tool use. It says: if the system's next action depends on what it remembers from before, it is always-on. The memory is what matters.

The BSD daemon's design philosophy is instructive. In BSD, daemons are first-class citizens of the operating system. They are developed in the same source tree as the kernel, the C library, and the userland tools. Their configuration files share a consistent syntax. Their manual pages are maintained as part of the source. Their init scripts, signal handlers, and logging conventions are standardized across the system. The daemon is not bolted on. It is part of the architecture. The system was designed with the assumption that background processes would be persistent, reliable, and governable.

This is the standard that always-on AI agents must meet. They are not bolted-on chatbots. They are daemons. They persist. They maintain state. They act on triggers. And they need the same architectural care that BSD gave its daemons forty years ago — with one critical addition. A BSD daemon's state was small and explicit: a configuration file, a connection table, a queue. An AI agent's state is everything it has ever seen, inferred, committed to, or been told. The durability problem is qualitatively different.

## The six axes of durable state

The paper decomposes persistent agent state along six diagnostic axes. Each asks a question that a traditional daemon's designer would recognize — but the answers for AI agents are different in ways that matter.

**Authority**: Who controls the state? A traditional daemon's state is controlled by its administrator — configuration files, init scripts, signal handlers. An always-on agent's state is accumulated from interactions with users, other agents, and the environment. The agent may be able to modify its own memory. The question of who has authority over what the agent remembers is not settled. It may not even be answerable in general — different state items may need different authority models.

**Scope**: What does the state cover? A cron daemon knows about scheduled jobs. An HTTP daemon knows about connections and request queues. An always-on agent may know about the user's calendar, their email history, their code repositories, their Slack channels, their browser tabs, and the contents of every conversation they have had with the agent. The scope is unbounded by default. Containing it requires explicit design.

**Mutability**: Can the state change, and how? Traditional daemon state is mostly immutable configuration and ephemeral runtime state. Agent state crosses every category: immutable audit logs, slowly-evolving user preferences, rapidly-updating task ledgers, permissions that escalate and revoke, commitments that bind future behavior. Each category needs different mutation rules. Most current systems treat all agent memory as a uniform append-only log. That is simple. It is also wrong for anything that needs to be forgotten, corrected, or revoked.

**Provenance**: Where did the state come from? In a traditional daemon, this is rarely asked — the config file was written by the sysadmin, the connection state was established by a handshake. In an always-on agent, state can arrive from the user, from the model's training data, from web searches, from other agents, from tool outputs, from the agent's own inferences. A piece of state that "the user's birthday is March 3rd" might have been stated by the user, inferred from an email, scraped from a social media profile, hallucinated by the model, or copied from another agent that got it wrong. Without provenance, the agent cannot assess the reliability of its own memory.

**Recoverability**: Can the state be restored after loss? Traditional daemons have crash recovery: replay the log, reload the configuration, reconnect to peers. Always-on agents accumulate state across months or years of interaction. If that state is lost — a corrupted vector database, a dropped partition, a migration that fails silently — the agent does not just lose data. It loses *context*. It forgets who the user is, what they were working on, what commitments it made. The paper notes that recovery is one of the least-studied dimensions. The database literature has fifty years of work on durability and crash recovery. The agent literature barely cites it.

**Actionability**: Does the state drive behavior? This is the axis that separates a journal from a daemon. A journal records. A daemon acts. If the agent remembers that the user prefers afternoon meetings, does it actually avoid scheduling morning calls? If it remembers a commitment to review a document by Friday, does it remind the user, or does it just store the fact inertly? Actionability is what makes memory consequential. Most current agent memory systems are better at storing than at acting — the retrieval pipeline is decoupled from the decision pipeline, and the connection between "I know X" and "I will do Y because of X" is weak.

> The six axes are a design checklist. For each piece of state an agent holds, you should be able to answer: who controls it, what it covers, how it changes, where it came from, how to get it back, and what it causes the agent to do. If you cannot answer all six, you do not understand the agent's memory architecture. If the agent's developers cannot answer all six, neither does the agent.

## The governance gap

The paper's central finding is a distributional skew. Across the 435-work corpus, research concentrates on *accumulating* and *retrieving* state. Governance, recovery, and forgetting are comparatively neglected. The field knows how to give agents bigger memories and better retrieval. It does not know how to govern what those memories contain, recover them when they are lost, or remove them when they should be forgotten.

This is not surprising. Accumulation and retrieval are the features that make demos impressive — the agent that remembers your preferences, the agent that finds the relevant document from six months ago. Governance, recovery, and forgetting are the properties that make systems safe and maintainable over years of operation. They are invisible in a demo. They are catastrophic in production.

The authors introduce the **Always-On Evaluation Protocol (AOEP-v0)** as a corrective. Instead of scoring agents on answer quality — did the agent retrieve the right fact, did it produce the correct response — AOEP scores agents on state management obligations: did the agent correctly mutate state when new information arrived, did it recover gracefully when state was corrupted, did it forget something when instructed to forget it, did it maintain the provenance chain that would allow an auditor to determine why it acted as it did. This is a shift from evaluating agents as question-answerers to evaluating them as state-managers. It is the difference between testing a program and testing a daemon.

The governance citations in the paper tell their own story: *Memaudit* for auditing agent memory, *OverlayGovernance* for governance frameworks, *IntentGovernedAuth* for intent-based authorization, *SkillContainment* for constraining what agents can do with what they know, *Bureaucracy* for organizational oversight structures, *RegulatoryUI* for making governance visible to regulators. Governance is not a single mechanism. It is a stack — audit trails, permission models, containment boundaries, oversight processes, regulatory interfaces. Each layer needs to be designed. Most current agent deployments have none of them.

## The lifecycle of durable state

The paper traces state through a lifecycle: **write, validate, organize, retrieve, act, update, forget, audit, rollback**. Each stage is a research area in itself. Taken together, they describe what it means for an agent to be a good steward of its own memory.

Write is where state enters the system. The challenge is not storage — vector databases and key-value stores are well-understood — but *selection*. The agent sees more than it should remember. Deciding what to keep is a compression problem with an objective function that is unclear. Keep everything and the retrieval quality degrades. Keep too little and the agent loses context that matters. The paper surveys mechanisms from MemGPT's operating-system-inspired paging to HippoRAG's hippocampal-inspired indexing to Memory-t1's learned selection policies. None are solved problems.

Validate is where state is checked for correctness, consistency, and freshness. A fact that was true in March may be false in July. A preference the user stated once may no longer hold. A commitment the agent made may have been implicitly superseded by later events. Validation requires the agent to maintain not just facts but *confidence*, *expiry*, and *dependency* metadata. Few systems do.

Forget is where state is removed. The paper's citation cluster on forgetting — *Coforgetting*, *EditingForgetting*, *WhatShouldLLMsForget*, *RecallForgettingBenchmarkingLongTerm*, *ControlPlaneForgetting*, *MachineUnlearningSurvey* — signals that forgetting is hard for the same reason retrieval is hard: state is entangled. If the agent inferred fact B from fact A, and fact A is later deleted or corrected, what happens to fact B? Should it be deleted too, or should it be re-derived from other sources? The problem is the same one that databases solve with CASCADE deletes and materialized views. Agents do not have those primitives yet.

Audit is where state is examined after the fact to determine why the agent acted as it did. This requires provenance chains that survive the agent's own memory updates — if the agent can rewrite its own memory, it can destroy the evidence of why it did something. Immutable memory, proposed in some of the cited work, is a partial answer. It does not solve the problem of what the auditor does with the immutability when the thing that was remembered was wrong.

Rollback is where state is reversed. An agent that sends an email, then learns the email was based on incorrect information, needs to be able to not just send a correction but *understand what other decisions depended on the incorrect information*. This is a causal reasoning problem that current agents are not equipped to solve.

> The lifecycle reveals that durable state is not a feature. It is a discipline. Each stage constrains every other stage. An agent that cannot forget should not remember indiscriminately. An agent that cannot audit should not act autonomously. An agent that cannot roll back should not commit.

## Durable execution: what the daemon layer needs

The paper identifies the problems. It surveys the taxonomies. It does not prescribe the runtime. But there is an emerging class of infrastructure that directly addresses the durability requirements the paper lays out. It comes from an unexpected direction — not from the AI research community, but from the database and distributed systems community. It is called *durable execution*, and it is the missing layer between the always-on agent and the reliable daemon.

### The idea

Durable execution is a simple idea with deep consequences. When a function executes as a durable workflow, every step is checkpointed to a durable store — typically Postgres. If the process crashes mid-workflow, the system recovers by reading the last successful checkpoint and resuming from that point. Completed steps are never re-executed. In-flight steps are retried with exactly-once semantics within the system's control boundary. The workflow survives process death, machine reboot, and database failover.

Two systems define the category. **Temporal**, founded in 2019 and battle-tested at DoorDash, Snap, Stripe, and Uber, provides a centralized orchestration server with a separate worker tier. Workflows are defined in code. The server checkpoints state to a persistence backend and dispatches activities to workers. It adds infrastructure — a Temporal server cluster plus Cassandra or Postgres — but it provides polyglot SDKs (Go, Java, Python, .NET, TypeScript, PHP, Ruby) and the operational maturity that comes from years of production use at scale.

**DBOS**, founded in 2023 by database pioneer Mike Stonebraker, takes the opposite approach. It is an open-source library — not a server — that checkpoints workflow state directly to Postgres tables. There is no orchestration server, no task queue, no sidecar. You annotate functions as workflows and steps. The library handles the rest. Step latency is 1–2 milliseconds — a single Postgres write — versus the tens to hundreds of milliseconds required for a round-trip through Temporal's central server. The operational footprint is Postgres, which most organizations already run. The philosophy is stated plainly: "Postgres is the operating system."

> Durable execution solves the recoverability axis directly. It provides the audit trail as a byproduct of checkpointing. It makes rollback tractable by making every state transition explicit and replayable. It is not the whole answer to the governance gap. It is the foundation on which the rest of the governance stack can be built.

### Why it matters for AI agents

An always-on agent that books a meeting, sends an email, updates a CRM record, and files a Jira ticket is executing a multi-step workflow with external side effects. If the process crashes after step two, what happens? Does the email get sent twice? Does the meeting get booked but never appear in the CRM? Does the agent recover and continue from step three, or does it start over from the beginning and double-book the meeting?

These are not hypothetical questions. They are the standard failure modes of any long-running process that touches external systems. Databases solved them decades ago with transactions. Workflow engines solved them with checkpointing and compensation. AI agents, as of mid-2026, have largely not adopted either.

The economic case is straightforward. Consider an agent that processes customer onboarding: verify identity, create account, provision access, send welcome email, schedule follow-up call. If the agent crashes after provisioning access but before scheduling the call, the customer has access but no follow-up. The cost of the dropped step is a churned customer — or an engineer spending two hours manually reconciling the state, which costs the same as several thousand agent invocations. The durable execution layer pays for itself the first time it prevents a dropped step.

The economics compound when agents operate autonomously at scale. A sales agent that maintains a pipeline of two hundred deals, each with a sequence of commitments it made to prospects — "I'll send you the pricing sheet by Tuesday," "I'll check with engineering and get back to you" — is a daemon with two hundred concurrent workflows, each with state that must survive process restarts. Without durable execution, every deploy, every crash, every autoscaling event is a risk that commitments will be dropped. The agent becomes unreliable not because the model is bad but because the runtime is fragile. The business loses deals. The trust cost is unquantifiable but real.

### The audit argument

There is a second economic advantage that is less obvious but potentially larger: auditability. When every step of an agent's execution is checkpointed to Postgres, the audit trail exists by construction. You can query why the agent sent a particular email: trace the workflow, inspect the state at each step, determine what the agent knew when it made the decision. You can answer the question "did the agent correctly mutate state when new information arrived" — one of the AOEP-v0 evaluation criteria — by inspecting the checkpoint history rather than by trusting the agent's self-report.

This matters economically because auditability is not optional in regulated industries. Finance, healthcare, legal, insurance — any domain where an agent's decisions have compliance implications — require the ability to reconstruct what happened and why. Building that audit trail from scratch is expensive. Getting it as a byproduct of the execution model is free, or close to it.

DBOS's architecture makes this particularly natural. Because workflow state lives in your Postgres database, you own it. You can query it with SQL. You can join it against your application tables. You can build dashboards, alerts, and compliance reports using tools you already have. Temporal stores state in its own server; you can query it through Temporal's visibility APIs, but the data lives in Temporal's infrastructure, not yours. Both approaches work. The difference is data sovereignty — and for organizations that care about where their agent's memory lives, that difference is material.

### Which one?

The choice between DBOS and Temporal maps to a broader architectural question: should the durable execution layer be infrastructure you operate or a library you import?

DBOS is the library answer. If you already run Postgres, you add a dependency and a handful of annotations. The step latency is low enough for interactive use. You own your data by default. The trade-off is a smaller ecosystem and fewer language SDKs. It is the right choice for teams that want durability as a capability of their application rather than as a platform they adopt.

Temporal is the platform answer. If you need polyglot workers, massive fan-outs, month-long workflows with signals and queries, or the operational maturity of a system that has been running at Stripe scale, Temporal is the proven choice. The trade-off is operational complexity: you now operate a Temporal cluster alongside your application.

Both are better than the status quo, which is agents running as stateless request-response loops with ad-hoc persistence bolted on. The paper's governance gap exists partly because the agent community has not adopted the durability primitives that the database and distributed systems communities have already built. DBOS and Temporal are the bridge.

> A daemon that forgets its commitments is worse than no daemon at all. The user who cannot trust the agent to remember what it promised will stop delegating to it. The economic value of an always-on agent is proportional to the durability of its state. Unreliable memory is a negative feature — it creates work rather than removing it.

## What this means for the AI era

The traditional daemon had a narrow contract. It did one thing. Its state was small and explicit. Its failure modes were understood — a crashed daemon could be restarted, a misconfigured daemon could be reconfigured, a daemon with a memory leak would eventually exhaust its heap and be killed by the OOM killer. The daemon was persistent, but its persistence was bounded.

The always-on agent breaks those bounds. Its state is large, implicit, and learned. Its failure modes are not fully understood. It can fail by remembering the wrong things, by forgetting the right things, by acting on stale information, by leaking state across user boundaries, by being poisoned through its memory, by drifting its behavior as its accumulated state changes its implicit model of the user. These are not hypothetical failures. They are the failure modes of any system that couples persistent state with learned behavior. The paper is a map of where those failure modes live and a survey of the research that is beginning to address them.

The connection to operating systems is not accidental. The paper connects always-on agents to databases, distributed systems, formal methods, capability security, and machine unlearning — the disciplines that have spent decades thinking about what it means for a system to maintain state correctly over time. The agent community is rediscovering problems that those disciplines have already solved, are solving differently, or have proven are unsolvable in general. The survey makes these connections explicit. It is an invitation for the agent community to stop reinventing durability from scratch.

Durable execution runtimes are the practical manifestation of that invitation. They take the ideas from databases and distributed systems — checkpointing, exactly-once semantics, write-ahead logging, crash recovery — and package them as infrastructure that an agent developer can use without a PhD in distributed systems. They are not the whole answer. They do not solve the authority question, or the provenance question, or the forgetting question. But they solve the recoverability question, and they make the audit question tractable, and they provide a foundation on which the rest of the governance stack can be built. That is enough to make them the most important infrastructure decision an agent team will make.

For the practitioner, the framework is now complete. When you deploy an agent that persists across sessions — whether it is a desktop agent like OpenWorker, a coding agent with workspace memory, or a customer-support agent with user history — you are deploying a daemon. The six axes give you a vocabulary for reasoning about what your daemon remembers. The lifecycle gives you a checklist for what your daemon needs to do with its state. The governance gap tells you what you are probably neglecting. And the durable execution layer — DBOS, Temporal, or whatever comes next — gives you the runtime guarantee that your daemon will survive the crashes, restarts, and failures that every long-running process eventually faces.

For the field, the paper marks a transition. The first generation of AI agents were stateless: prompt in, response out, no memory between calls. The second generation added retrieval: the agent could look things up, find relevant documents, ground its responses in provided context. The third generation — the one the paper surveys — adds persistence: the agent maintains state across interactions, learns from that state, and acts on it. The fourth generation will add durability: the agent's state survives process boundaries, its decisions leave audit trails, its commitments are recoverable, its failures are compensatable. Each generation reduces the distance between what an agent is and what a daemon has always been.

The BSD Daemon's pitchfork is not a weapon. It is a tool for digging — into processes, into state, into the internals of the system. The always-on agents surveyed in this paper are doing the same kind of digging, but into a different substrate. They dig into interaction histories, accumulated preferences, learned patterns, inferred relationships. The digging is the work. The durability is what makes it possible. The governance is what keeps it from becoming dangerous. And the execution layer — the durable runtime beneath the agent — is what keeps it from being erased by a process restart.

Beastie, in his tennis shoes and grin, has been running in the background since 1976. The new daemons are larger, smarter, and stranger. They need better foundations than the ones we have been giving them.

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- [DBOS Documentation](https://docs.dbos.dev/) — Durable execution as a Postgres-backed library. Founded by Mike Stonebraker.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution as a platform. Battle-tested at DoorDash, Snap, Stripe, Uber.
- McKusick, M.K., et al. The BSD Daemon. [FreeBSD Copyright and Trademark](https://www.freebsd.org/copyright/daemon/).
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux) — The design philosophy behind the operating systems that gave daemons their cultural form.
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer) — A desktop AI agent that is, in effect, a durable daemon running on the user's machine.

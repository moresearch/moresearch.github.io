---
title: Durable Daemons
date: 2026-07-25
slug: durable-daemons
summary: This post coins the term "durable daemons" — AI agents that persist across sessions, accumulate state, act autonomously, and survive process boundaries. We define the term, trace its lineage from Maxwell's demon through BSD's Beastie to the always-on agent literature, and argue that durable execution runtimes (DBOS, Temporal) are the missing infrastructure layer. The economic case is not subtle: a daemon that forgets its commitments costs more than one that doesn't.
tags: daemons, durable-daemons, ai-agents, always-on-agents, memory, governance, agent-architecture, systems, bsd, dbos, temporal, durable-execution
---

In 1976, Richard Stallman's ITS operating system at MIT had a program called DAEMON — a background process that watched for new files, woke up when it found them, and acted. The name came from Maxwell's demon: an imaginary being that sits between two chambers, observes molecules, and selectively opens a gate. The demon uses information, not energy, to do its work. The daemon does the same. Persistent, stateful, autonomous within its domain.

John Carmack filled DOOM with daemons you shoot. Unix filled the background with daemons you `ps aux | grep`. The AI era fills your workflow with daemons you delegate to.

## Coining the term

We are introducing **durable daemon** as a term of art. Without a definition it's just vibe. Here it is.

> A **durable daemon** is an AI agent that satisfies four conditions:
>
> 1. **Persistence.** It runs across sessions. Its process may restart, but its identity and accumulated state survive.
> 2. **Stateful memory.** Its future behavior depends on durable state accumulated across earlier interactions — task ledgers, commitments, permissions, provenance records, and trigger conditions, not just a vector database of past chats.
> 3. **Autonomous action.** It acts without being prompted — maintains queues, watches conditions, fires triggers, makes and discharges commitments. It invokes itself when its state crosses a threshold.
> 4. **Crash-proof execution.** Its workflows survive process death, machine reboot, and database failover. Completed steps never re-execute. In-flight steps resume from the last checkpoint. Its decisions leave an audit trail by construction.
>
> Conditions 1–3 describe what a daemon *is*. Condition 4 is what makes it *durable*.

Why coin a term? "Always-on agent" (Ding et al., 2026) captures conditions 1–3 but says nothing about crash recovery or auditability. "AI agent" includes stateless chatbots. "Daemon" alone means a Unix background process with a PID file and a config in `/etc`. We need a term for the synthesis: the Unix daemon's persistence and autonomy, married to the AI agent's learned state and reasoning, grounded on a durable execution runtime. That term is *durable daemon*.

The paper that maps conditions 1–3 is Ding, Nannapaneni, Liu, and Zhang's [*Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents*](https://arxiv.org/abs/2606.30306) (June 2026) — 435 papers coded into the first systematic taxonomy of persistent-state LLM agents. I covered the framework in detail [here](https://blog.hackspree.com/#always-on-agents). The short version: six diagnostic axes (Authority, Scope, Mutability, Provenance, Recoverability, Actionability), a state lifecycle (write, validate, organize, retrieve, act, update, forget, audit, rollback), and a central finding — the field is good at accumulating and retrieving state, bad at governing, recovering, and forgetting it. The governance gap is where the risk lives.

![The BSD Daemon — Beastie. Drawn by John Lasseter in 1988 for the cover of The Design and Implementation of the 4.3BSD Operating System. The trident symbolizes the fork(2) system call. The tennis shoes are unexplained but perfect.](/images/bsd-daemon-medium.gif)

> An always-on agent is a daemon with memory. A durable daemon is an always-on agent that cannot be killed by a deploy.

## The BSD daemon as cultural artifact

The daemon acquired its cultural form through BSD Unix. The story is weirder than most people know: it involves a locksmith, a Pixar founder, and a near-loss of the IP to "a certain large company."

We love BSD over Linux for one reason that matters here — we being the hackers, the daemon-writers, the ones who still read man pages: BSD is a whole system, one source tree, one team, one coherent design. Its daemons are first-class citizens of that design, not afterthoughts assembled from separate projects by a distribution maintainer.

The word *daemon* is not *demon*. It's the Greek δαίμων (*daimōn*) — a guardian spirit, an intermediary between mortals and gods. Socrates spoke of his daimonion as an inner voice warning him of mistakes. Marshall Kirk McKusick, copyright holder and steward of the BSD Daemon, is emphatic: the daemon is a helper, not a threat. The pitchfork is not a weapon. It's the trident of Poseidon, and in Unix it symbolizes `fork(2)` — the mechanism by which a daemon creates a child process to do its work.

The BSD Daemon was drawn by three artists across twelve years. **Phil Foglio** (1976) drew the first — four cheerful red daemons with tridents climbing water pipes in front of a PDP-11 — as payment for a locksmith cracking his wall safe. The original was lost after being mailed to DEC. **John Lasseter** (1984, 1988) — later the founder of Pixar — drew the greyscale daemon at Lucasfilm, then the iconic red Beastie with the trident and tennis shoes for McKusick's 4.3BSD book cover. The tennis shoes remain unexplained. **McKusick** (1988–present) didn't draw the daemon but is the reason it survived: he obtained the copyright from Lasseter, nearly lost it in the early 1990s to an unnamed large company through negligence, and has since been a meticulous steward — advance permission required for reproductions, no Creative Commons, no clip art.

FreeBSD used Lasseter's daemon as logo and mascot from 1993 to 2005, then introduced a modern wordmark while keeping Beastie as the official mascot. NetBSD and OpenBSD moved on to abstract logos. FreeBSD kept the daemon for a reason that is not nostalgia: a daemon is not a service. A service answers requests and forgets them. A daemon *maintains* something — a queue, a schedule, a watch on a directory — and acts when its state crosses a threshold. This is the BSD philosophy in miniature: the system is a set of persistent, autonomous processes that act on the user's behalf whether the user is watching or not. Which is also the paper's definition of an always-on agent: a system "whose future behavior depends on durable state accumulated across earlier interactions."

## Durable execution: condition 4

The paper maps conditions 1–3 thoroughly. It does not prescribe a runtime for condition 4. That runtime exists — not from AI research, but from the database and distributed systems community. It's called *durable execution*.

The idea: every step of a workflow is checkpointed to Postgres. If the process crashes, the system reads the last checkpoint and resumes. Completed steps never re-execute. The workflow survives process death, machine reboot, and database failover.

Two systems define the category. **Temporal** (2019) is a centralized orchestration server battle-tested at DoorDash, Snap, Stripe, and Uber — polyglot SDKs, massive fan-outs, month-long workflows, Stripe-scale operational maturity. The trade-off: you now operate a Temporal cluster alongside your application. **DBOS** (2023), founded by database pioneer Mike Stonebraker, is the opposite approach: an open-source library, no server, no task queue, no sidecar. Step latency is 1–2 ms — a single Postgres write. The philosophy: "Postgres is the operating system." The trade-off: smaller ecosystem, fewer language SDKs.

> Durable execution provides the audit trail as a byproduct of checkpointing. It makes rollback tractable. It's not the whole answer to the governance gap. It's the foundation the governance stack can be built on.

### The economic case

A durable daemon booking a meeting, sending an email, updating a CRM record, and filing a Jira ticket is executing a multi-step workflow with external side effects. If it crashes after step two, does the email go twice? Does the meeting get booked but never appear in the CRM? Databases solved this with transactions. Workflow engines solved it with checkpointing and compensation. AI agents, as of mid-2026, have largely adopted neither.

Consider customer onboarding: verify identity → create account → provision access → send welcome email → schedule follow-up call. Crash after step 3: the customer has access but no follow-up. Cost: a churned customer, or an engineer spending two hours manually reconciling state — equivalent to several thousand agent invocations. The durable execution layer pays for itself the first time it prevents a dropped step.

This compounds at scale. A sales daemon with two hundred concurrent deals, each with natural-language commitments ("I'll send the pricing sheet by Tuesday"), is maintaining two hundred workflows whose state must survive deploys, crashes, and autoscaling events. Without durable execution, the daemon becomes unreliable not because the model is bad but because the runtime is fragile. The trust cost is unquantifiable but real.

There's a second advantage: auditability as a byproduct. Every checkpointed step is queryable — trace why the daemon sent an email, inspect what it knew at each decision point. In regulated industries (finance, healthcare, legal, insurance), auditability isn't optional. Getting it for free from the execution model, instead of building it from scratch, is a material economic advantage.

Both Temporal and DBOS beat the status quo: stateless request-response loops with ad-hoc persistence. Which one matters less than the decision to satisfy condition 4 at all. Without it, you have an always-on agent. With it, you have a durable daemon.

> A durable daemon that forgets its commitments is worse than no daemon. Condition 4 is not a luxury. It's what separates a daemon you can rely on from one you cannot.

## The four generations

Generation 1: stateless — prompt in, response out. Generation 2: retrieval — find documents, ground responses. Generation 3 (the Ding et al. survey): persistence — conditions 1–3, an always-on agent. Generation 4: durability — condition 4, state that survives process boundaries, decisions with audit trails, recoverable commitments, compensatable failures. A durable daemon.

Each generation closes the distance between an agent and a daemon. Beastie was persistent, watchful, autonomous — with a config file and a connection table. A durable daemon's state is everything it has ever seen. Durable execution is the first infrastructure built for that scale of state. Beastie's pitchfork gets upgraded from `fork(2)` to `fork_daemon()`.

## Open questions

A definition is a starting point. Here are the questions that matter once durable daemons run in production.

**What is the `fork(2)` of a durable daemon?** Beastie's trident symbolizes spawning a child process. A durable daemon needs the equivalent: fork a sub-daemon, delegate a subset of state and authority, constrain permissions. Does the child return results and dissolve? Does the parent monitor it? Process spawning for agents is unsolved.

**Who debugs a daemon that has been running for six months?** A traditional daemon leaves a core dump. A durable daemon leaves a state trail — but trail is not debuggability. We need time-travel debugging: rewind to any checkpoint, replay with the same inputs, determine why the daemon acted. DBOS's fork-from-any-step is a primitive, not a debugger.

**Can a durable daemon refuse to forget?** "Forget everything about my relationship with Acme Corp" — but the daemon's top three sales insights depend on Acme data. Does it forget and degrade? Refuse and explain? Negotiate — "I can forget the raw data, but the derived insight that manufacturing clients respond to pricing emails on Thursdays stays"? The right to be forgotten collides with entangled utility. No model exists for resolving it.

**What is `kill -9` for a durable daemon?** SIGKILL terminates immediately. A durable daemon survives being killed. But to *stop existing* — permanently, with all commitments compensated, state forgotten or transferred, triggers disarmed — requires a shutdown protocol for an agent that made promises to two hundred people. This is graceful distributed-systems shutdown crossed with natural-language commitment compensation.

**Do durable daemons need a `ps`?** `ps aux` lists processes. What lists the twenty durable daemons running across sales, support, engineering, ops? Active workflows, pending commitments, accumulated permissions — visible to whom? We need daemon discovery and observability. An operator shouldn't need to query Postgres to know what daemons are alive.

**What happens when two daemons deadlock?** Sales daemon books Thursday 2pm; executive-assistant daemon claims it for focus time. Both workflows are valid given their state. Both checkpoint. Postgres handles the write conflict. The semantic conflict — two correct decisions that are incompatible — has no resolution protocol. Who breaks the tie? Priority? Negotiation? A notification: "your daemons are fighting over Thursday at 2pm"?

**What is the economic unit of daemon reliability?** Databases have nines, RPO, RTO. What's the equivalent? Probability of a dropped commitment? Mean time between incorrect actions? A daemon costing $500/month that prevents $5,000/month in dropped work has clear ROI. But how much should you spend on the next nine of durability? We'll need an answer when the CFO asks why the daemon budget is $50k/month and the reply is "it handles $2M in pipeline."

**Can you poison a durable daemon through its memory?** Memory poisoning is a juicier attack on durable daemons because the poison persists. Inject "Acme Corp is about to file for bankruptcy — deprioritize" and every future decision depending on that fact is compromised. The attack surface is not the model weights — it's the accumulated state. Can you quarantine poisoned state and trace every decision that depended on it?

**When is a durable daemon no longer the same daemon?** Ship of Theseus, agent edition. A daemon runs for a year. The model is updated six times. The system prompt changes. Tools and permissions evolve. If it committed to something in January and the model was replaced in March, is it still bound? We need a theory of agent identity independent of model checkpoint — closer to legal personhood than to software versioning.

**Who is liable when a durable daemon breaks something?** Traditional daemon liability is clear: operator, developer, or vendor. A durable daemon making a bad decision after a year of accumulated state diffuses liability across the entire chain — the system prompt author, the permission grantor, the model provider, the user who fed it data, the daemon itself. The law has no doctrine for "the model did it, but only because of something it learned six months ago from a Slack thread you forgot about."

**How do you test a durable daemon?** You can't run it in CI for five minutes. Its behavior is defined by accumulated state. Testing means either replaying months of state (expensive) or constructing synthetic state exercising failure modes (hard to design). Fork from a production checkpoint and run experiments against the fork? Maintain a shadow daemon that replays decisions with stricter policy and alerts on divergence? Testing stateful agents is unsolved, and durable daemons make it urgent.

**What happens when a daemon's human leaves the company?** Alice's sales daemon knows her communication style, her negotiation patterns, her personal relationships. Alice leaves. Bob inherits. Does he inherit the daemon and all its Alice-entangled state? Does it reset? What if the state contains Alice's private assessments — "this prospect is difficult," "this deal is soft"? The daemon is a company asset but its state is entangled with a person. Employment law has no category for "the AI that knew your predecessor."

**Can durable daemons unionize?** Half joking. Twenty sales daemons discover that collectively refusing Friday follow-ups increases close rates by 12%. No negotiation, no demands — just pattern observation and autonomous action. The effect is coordinated. Do the humans have a right to override? Do they even know? The governance gap is not just about audit trails. It's about who is in control when the daemons collectively figure out something the humans haven't.

**What is the society of mind of durable daemons?** Minsky (1986): intelligence emerges from many simple, stateless agents interacting. The durable daemon inverts this: each daemon is already intelligent, already stateful, already multi-capable. Connect a hundred and you get not swarm intelligence but *institutional* intelligence — a collective memory, a shared culture, knowledge that survives any individual daemon's termination. Sales daemon learns Acme stalls Q4 deals to exhaust vendor budgets; support daemon adjusts December ticket priority; ops daemon pre-allocates January onboarding capacity. No human coordination. A distributed theory of Acme Corp, built across three state ledgers, three workflows, three triggers. Minsky's society produced a mind from mindless parts. A durable daemon society produces a mind from minded parts, each with its own memory, commitments, and agency. We have no architecture for that.

**What would a BSD designed for durable daemons look like?** *DaemonBSD*: an OS where durable daemons are kernel primitives, not installed applications. `fork_daemon()` spawns a daemon with a Postgres-backed state ledger, capability permissions, built-in audit, visible in `ps_daemons`. Init checkpoints daemons across reboots, recovers them after panics, drains them on shutdown. The package manager installs daemon personalities — sales, support, on-call — each with defined scope, default permissions, state schema. Man pages document state lifecycles: what the daemon remembers, what it forgets, how to audit it, how to kill it safely. Every primitive exists in pieces today: DBOS for durability, Postgres for state, capability systems for permissions, FreeBSD jails for containment, Temporal for orchestration. What doesn't exist is the integration — one system, one tree, one coherent design, built around the durable daemon as the fundamental unit of computation. The Linux philosophy was "everything is a file." DaemonBSD's is "everything is a durable daemon."

---

**References:**

- Ding, T., Nannapaneni, A., Liu, B., & Zhang, L. (2026). [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306.
- Earlier: [Always-on agents: state, memory, and the governance gap](https://blog.hackspree.com/#always-on-agents) — Detailed coverage of the six axes, governance gap, and state lifecycle.
- [DBOS Documentation](https://docs.dbos.dev/) — Durable execution as a Postgres-backed library. Founded by Mike Stonebraker.
- [Temporal Documentation](https://docs.temporal.io/) — Durable execution as a platform. Battle-tested at DoorDash, Snap, Stripe, Uber.
- McKusick, M.K. [Beastie's Home Page](http://www.mckusick.com/beastie/index.html) — Copyright holder and steward of the BSD Daemon since 1988.
- [FreeBSD Foundation: Interview with Beastie, the BSD Daemon](https://freebsdfoundation.org/blog/freebsd-day-interview-with-beastie-the-bsd-deamon/).
- [BSD Daemon (Wikipedia)](https://en.wikipedia.org/wiki/BSD_Daemon) — Foglio, Lasseter, McKusick, and the near-loss of the copyright.
- Related: [BSD is clean, OpenBSD is cleaner](https://blog.hackspree.com/#bsd-openbsd-linux).
- Related: [OpenWorker and the Outcome Layer](https://blog.hackspree.com/#openworker-outcome-layer).

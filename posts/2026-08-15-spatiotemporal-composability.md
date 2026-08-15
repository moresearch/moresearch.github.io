---
title: "Spatiotemporal Composability: The Missing Calculus for Self-Evolving Agents"
date: 2026-08-15
slug: spatiotemporal-composability
summary: "Every program that supports plugins has a shameful button: Reload Window. It exists because nobody has a theory of removing a plugin from a running program safely. A new 88-page paper from Peking University and DeepSeek-AI — A Programming Paradigm for Spatiotemporal Composability — is that theory: two guarantees (when a component leaves, everything it did is undone; when the world changes, everything that depended on it reconnects), a proof that the guarantees compose across a whole system, and a four-year production track record in Koishi's 4,000-plugin ecosystem. Here is what the paper actually does, in plain language, and why it is the missing foundation for self-evolving agents."
tags: agentic-design, dynamic-composition, spatiotemporal-composability, effects, coeffects, agent-harnesses, self-modifying-agents, plugin-systems, formal-methods, capability-security, sandboxing, koishi, cordis, deepseek
---

On August 13, 2026, an 88-page paper appeared on GitHub: *A Programming Paradigm for Spatiotemporal Composability*, by Yifan Shi and Wei Zhang of Peking University and Tianyi Cui of DeepSeek-AI, published from the [cordiverse](https://github.com/cordiverse/paper) repository. Within two days it had over 1,500 GitHub stars — an unusual reaction for a paper whose tools are monads, coeffects, and operational semantics.

The reaction is deserved, and it is not about the mathematics. The paper answers a question every engineer has felt but nobody has formalized: **how do you safely remove a piece of software from a running program?** Plugin systems have been dodging this question for decades. Agent harnesses are about to be forced to answer it, because an agent that modifies its own runtime is doing dynamic composition — adding, removing, and replacing components while the system runs — and it cannot keep dodging.

## The problem, in one paragraph

Every program that supports plugins has a shameful button. In Visual Studio Code it is "Developer: Reload Window" — a command that restarts the entire extension host so you can remove one extension. The button exists because the alternative is too hard: there is no safe way to unload a single extension's code, undo what it did to the program, and keep everything else running.

The paper's opening measurement makes the point concrete. Among the top 100 VSCode extensions by install count, **87 contain executable code** — and removing any one of them requires restarting the whole host, affecting all the others. The `deactivate` hook is not a removal mechanism; it is a goodbye callback for when the *process* is shutting down. On the other axis, only **7 of the top 100 extensions** declare dependencies on other extensions at all, and the one mechanism for inter-extension interaction returns an untyped `any`. VSCode does not fail at dynamic composition; it avoids it.

Why does the entire industry avoid it? Because operating systems and container orchestrators offer a blunt workaround: restart the process to get rid of a bad module, let the orchestrator manage service dependencies. The paper's cost accounting for this workaround is the quiet heart of the motivation. Every restart throws away all process-local state — caches, connections, partial computations — and rebuilding it takes seconds to minutes, with redundant replicas bought to cover the gap. Container orchestration cannot express dependencies between components that share an address space, and it turns local function calls into network calls. Modern software composes at a finer grain than processes and containers; the workaround is a granularity mismatch, and it is the only mechanism anyone has.

> Restart to reconfigure works — the way restarting your computer works. The paper is about making it unnecessary, with a theory underneath instead of a button.

## Two guarantees, named as dimensions

The paper's central move is to notice that dynamic composition has exactly two requirements, and to name them. They are called *temporal* and *spatial* because one concerns the time axis of a component's life and the other the space of the system's dependency graph.

**Temporal composability** is about the moment a component leaves. Everything the component did to the shared environment — every resource it allocated, every event it registered, every value it mutated — must be completely and safely reversed, in order. In ordinary static code this is a solved problem: it is called lexical scoping, or RAII in C++, where a variable's destructor runs automatically when its scope ends. At runtime the scopes are gone — a plugin's effects can outlive any function call — so there is no built-in destructor, and nothing automatic to run.

**Spatial composability** is about the space between components. A component must be able to declare what it needs from the others, and the system must resolve those needs as components appear, disappear, or are replaced. In static code this is module import resolution: the linker does it once, at startup. At runtime a dependency can vanish mid-execution, so the wiring has to happen continuously, not once.

![Two guarantees for pluggable software — temporal: plug in, every action gets an undo; plug out, the undoes play in reverse, state returns, no restart, others untouched. Spatial: a component declares what it needs; when a provider changes, dependents reconnect or wait — a missing provider means "wait", never "crash". In static code these are ordinary (RAII, lexical scoping, module imports); at runtime they are the whole problem.](/images/composability-dimensions.png)

Both dimensions map one-to-one onto the ways self-modification can kill an agent harness. The temporal failure, in the paper's words: "a faulty self-modification can disable the very process needed to recover." The spatial failure: "a naive code-replacement strategy may silently break dependents or introduce circular dependencies that surface only at reload time." Notice what the second sentence says: the failure is invisible until the next reload — the exact moment the system most needs to keep working.

## How the paper delivers the guarantees

This is where the paper stops being a position paper. To solve the two dimensions it takes two concepts from programming-language theory — *effects* and *coeffects* — and turns them from compile-time annotations into runtime mechanisms.

### Effect one: every action carries its own undo

**Effects** are what a component does to its world: allocating memory, registering a callback, writing to a store, opening a connection. The paper's idea is that every such effect should carry an *inverse* — its own undo — and that the runtime should keep a log of which effects each component performed, in what order.

When the component is removed, the runtime plays the inverses in reverse order: last effect undone first, like unwinding a stack. Two properties make this work: the inverse of a composite effect is just the composition of the inverses (so the runtime can derive cleanup for *anything* the component did, automatically), and the log is per-component (so undoing one component never touches another).

The consequence is almost unfair. A plugin author never writes an uninstall function. They write one line per atomic operation — the operation and its undo — and the runtime composes them into correct, ordered cleanup. **Teardown is derived from loading, not written alongside it.** Correctness that used to depend on each author's diligence becomes a structural property of the system.

### Effect two: dependencies that declare themselves

**Coeffects** are what a component *needs* from its world: a database, a filesystem, an adapter for a messaging platform. Where effects describe what you change, coeffects describe what you require.

The paper's idea is that a component declares its needs as a *specification* — "I require a storage backend" — and the runtime does the rest. Every time the world changes, the runtime checks each component's specification and tells it how to respond, classifying the change as *activating* (your need just got satisfied, start running), *deactivating* (your provider was removed, wind down), or *neutral* (nothing you declared changed — don't react). A component whose dependency is missing simply stays inactive; it does not error. When the dependency appears, it activates.

Two refinements carry the security weight. **Isolation** decides what a declared need resolves to — which concrete provider you actually get. **Interception** decides what you are allowed to do with it — which calls are permitted. Both are controlled by the runtime, not by the component.

### The one thing that holds it together: the context

The paper's third move unifies the two mechanisms behind a single object, which it calls **the context**. The context carries three things at once: the current state of the system, the undo log for effects, and the dependency map for coeffects. Every interaction between a component and its environment passes through it — there is no other door.

This is what makes the whole design auditable. Because everything goes through one object, the runtime can track everything a component does (to undo it later) and everything it needs (to reconnect it), and — as we will see — everything it is allowed to touch.

![One context — the only door: every interaction between a component and the world passes through a single object holding the current state, the undo log, and the dependency map. Components plug in (load: effects tracked) and plug out (unload: effects undone). One door means nothing can leak — and nothing can be missed.](/images/composability-context.png)

The paper gives this object a poetic, literal metaphor: **loading a component is plugging it in — executing its effects; unloading is unplugging it — recovering them — without affecting the other things plugged into the same strip.** Because the context is recursive, a parent context can hold child contexts, so components can be composed hierarchically: a plugin can itself be a socket that other plugins plug into, and removing it unplugs everything downstream of it, cleanly.

## Why this is a new way of programming

The paper argues that unifying effects and coeffects in one context is not a library design but a programming paradigm — a third answer to the oldest question in programming: how should a program handle side effects?

The two existing answers define the spectrum. **Functional** programming threads state explicitly: every function takes the world and returns a new world, so effects are visible in the types and reasoning is airtight — at the cost of threading a state parameter through every call, and stacking a monad for each new kind of effect. **Imperative** programming hides it: effects happen implicitly, and dependencies are pulled from a process-wide registry with null-checks and casts at every call site — ergonomic, and effectively untraceable. React's `useEffect` registers a persistent side effect that appears in no parameter; Spring's `getBean(...)` reaches into a global registry; understanding what a call does means reading its implementation, transitively, and refactoring silently breaks distant invariants.

The context paradigm is the middle path: **the traceability of the functional approach with the ergonomics of the imperative one.** Every effect and every dependency goes through an explicit context, so each operation is attributable to the component that performed it — but the developer does not thread state by hand, and does not write cleanup or wiring at all. The system derives both.

> The paper's central claim: correctness that used to rest on developer discipline becomes a structural property of the paradigm. For human plugin authors that is a quality-of-life improvement. For autonomous agents it is the difference between a harness that can evolve and one that self-destructs.

## The proof: the guarantees compose

A single component that can be plugged out cleanly is nice. What makes the paper a foundation rather than a feature is that it proves the guarantees **compose**. It defines a small calculus of dynamic composition — components with a full lifecycle: loading, iterating, withdrawing, failing, running asynchronously — and proves the properties that matter: type preservation, temporal composability, spatial composability, progress, and confluence. The theorems carry the two guarantees from a single component to a whole system of interleaved components: if every part can be removed cleanly and reconnects correctly, so can the whole.

This is the property a self-modifying system needs before anyone should let it modify itself unattended. "Each part is safe to swap" is a local statement; "the system stays consistent through a storm of swaps" is the global one, and the metatheory is what connects them.

One honesty note, stated plainly in the paper: recovery is guaranteed *up to observational equivalence*, not literal equality. Unplugging a component does not restore the heap's exact byte layout — `free` releases a block without restoring the arrangement `malloc` left behind. What the runtime guarantees is that no observer can tell the difference, where the observers are defined by the declared dependencies. What you expose as dependencies is what recovery promises to restore — a useful design pressure in itself.

## Proven in production: Koishi

The theory was not written first and applied later; it was extracted from a working system. **Koishi**, an open-source chatbot framework built on Cordis, has accumulated over **4,000 community-contributed plugins** over four years — IM adapters, database drivers, admin consoles, end-user features. Three results stand out.

**Temporal composability without cognitive overhead.** An orchestrator disables a plugin from the console and its effects are withdrawn in place. During development, hot module replacement re-applies an edited plugin on save while preserving cache state and live connections elsewhere. Because effects through the context are tracked and their inverses composed automatically, even an inexperienced plugin author gets correct, ordered cleanup *without writing an uninstall path*. The correctness that used to depend on each author's diligence is discharged once, by the abstraction.

**Spatial composability across an open ecosystem.** Koishi's plugins have a genuine dependency topology — adapters provide platforms, drivers provide storage, functional plugins use both. Switching the storage backend or reconnecting an adapter at runtime reactivates only the dependents whose resolved dependency actually changed. A plugin whose dependency is unavailable stays inactive until it appears — no crash, no error, no manual wiring. The composition holds across independently authored code: a plugin and its dependency are usually written by different people who coordinate on nothing but the declaration that connects them.

**Expressiveness and generality.** Koishi's web console is a second, independent Cordis application whose plugins compose browser and UI primitives rather than server ones — the same model in a wholly different runtime. The framework contributes only domain vocabulary; the paradigm fixes how effects and coeffects compose, leaving their meaning to each application.

The paper is honest about its threats to validity: a single ecosystem, a single host language (TypeScript), observational rather than controlled evidence. What it establishes is an existence-and-adoption result — which, for a programming paradigm, is exactly the result that matters most.

## Why this is the foundation for self-evolving agents

The paper names its own future direction, and it is the reason to read it now: "self-evolving agent harnesses, where an AI agent generates and replaces its own harness components continuously and with little human oversight." Its citations map the field — OpenAI's *Harness Engineering*, Anthropic's *Harness Design for Long-Running Application Development*, MemGPT's agents-as-operating-systems, the multi-agent surveys. The paper is the formal foundation this literature has been reaching for. Five reasons it is the right foundation:

1. **Self-modification is the one thing every current harness does worst.** Today, an agent that "adds a tool" appends to a list; one that "replaces a module" restarts the process — dropping the warm context, the cache, the in-flight work — or patches itself with hacks the harness cannot unwind. The temporal guarantee is the missing primitive: replace a component, have its effects completely reversed, without a restart, without touching the components that must survive.

2. **The two guarantees are the two failure modes of self-evolving agents.** The temporal failure — a faulty self-modification disabling the very process needed to recover — is the recursive-death scenario every self-modifying system eventually faces. The spatial failure — dependents silently broken, circular dependencies surfacing only at reload — is what happens the first time an agent replaces a provider other components depend on. The paper gives both failure modes a theory, which is the precondition for a mitigation.

3. **Agents cannot be trusted with discipline, so correctness must be structural.** A prompt can say "clean up after yourself." A structure that *derives* the cleanup from the load cannot forget. This is the same distinction this blog has drawn between agents that verify and agents that hope ([Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), [The Verification Horizon](https://blog.hackspree.com/#verification-horizon)): guarantees you must remember are guarantees an agent will, at some point, not.

4. **Sandboxing falls out of the design instead of being bolted on.** Because every interaction passes through the context, and dependencies are declared before a component runs, the set of capabilities a component requires is known *before it executes*. The paper calls the declaration a "capability request" and the context proxy a "capability mediator," explicitly citing capability-based security. Undeclared access raises an error; interception carries fine-grained policy — a filesystem dependency can declare which paths a component may read or write, checked on every call, adjustable by an orchestrator without touching either party. For agent harnesses, whose components are not merely untrusted but *machine-written*, access control that is structural rather than inspected is the only kind that scales ([Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents), [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface)).

5. **The restart tax has an economics, and this blog has been counting it.** The paper's seconds-to-minutes rebuild costs are the token economics of agent restarts: a restart is not free, it is the loss of every warm context, every cache, every connection ([Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag), [Always-on agents](https://blog.hackspree.com/#always-on-agents)). Spatiotemporal composability is the mechanism that turns "restart to reconfigure" into "reconfigure in place" — the difference between an agent that pays the full price of state loss on every self-modification and one that pays nothing.

And the provenance matters. This is not armchair type theory: it is a paradigm extracted from a 4,000-plugin production system, with a DeepSeek-AI co-author — a frontier lab that needs harness foundations the way previous eras needed operating systems. When a paper arrives with a formal calculus *and* a working meta-framework *and* a large production ecosystem, it is worth reading as more than theory.

## What the paper doesn't claim

Three caveats are worth carrying into any reading. First, the recovery guarantee is behavioral, not physical — "no one can tell it was ever there" rather than "the bytes are identical" — and the boundary between what the system can track and what it cannot (the paper's *system boundary*) is itself a design decision. Second, the evidence is one ecosystem, one language, and not yet measured: overhead and developer-productivity numbers against a baseline remain future work. Third, the paradigm is young: Γ∞, the unified context type, is a beautiful object, but paradigms are validated by adoption, not by theorems. The paper is also dense — 88 pages of monads, coeffects, and operational semantics — which is exactly why a digest helps.

None of these caveats change the shape of the argument. The paper named the two dimensions, gave them a calculus, proved the calculus, and shipped the proof in production. Every other foundation for dynamic composition is now playing catch-up.

## Bottom line

Agents will modify their own harnesses. That is not speculation; it is the stated direction of the harness-engineering literature and the paper's own future-work section. The only open question is whether self-modification is done with ad hoc patches and restarts — or with a runtime that guarantees, structurally, that a removed component leaves nothing behind and a changed dependency rewires itself safely.

*A Programming Paradigm for Spatiotemporal Composability* is the answer to that question: two guarantees, a single context that holds them together, a proof that they compose, and a four-year production validation with 4,000 plugins. It is the missing calculus for self-evolving agents — and it arrived, appropriately, in a GitHub repository: the first component of a future harness that is, in the most literal sense, self-composed.

---

**References:**

- Yifan Shi, Wei Zhang, Tianyi Cui. [A Programming Paradigm for Spatiotemporal Composability](https://github.com/cordiverse/paper) (Peking University / DeepSeek-AI, 2026) — the paper; 88 pages; [paper.pdf](https://github.com/cordiverse/paper/blob/main/paper.pdf).
- [Cordis](https://cordis.io/) — the meta-framework of spatiotemporal composability: core library (effect tracking, coeffect resolution, isolation, interception) and declarative component loader (configuration reconciliation, hot module replacement).
- [Koishi](https://github.com/koishijs/koishi) — the chatbot framework built on Cordis; 4,000+ community plugins over four years; [koishi.chat](https://koishi.chat).
- D. L. Parnas. [On the criteria to be used in decomposing systems into modules](https://dl.acm.org/doi/10.1145/361598.361623) (Communications of the ACM, 1972) — where the paper's composition story begins.
- D. Birsan. [On Plug-ins and Extensible Architectures](https://dl.acm.org/doi/10.1145/1053331.1053345) (ACM Queue, 2005) — plugin systems as the canonical instance of dynamic composition.
- OpenAI. [Harness Engineering: Leveraging Codex in an Agent-First World](https://openai.com/index/harness-engineering/) — the harness as the product surface for agents.
- Anthropic. [Harness Design for Long-Running Application Development](https://www.anthropic.com/engineering/harness-design-long-running-apps) — harness engineering for long-running agents.
- C. Packer et al. [MemGPT: Towards LLMs as Operating Systems](https://arxiv.org/abs/2310.08560) — the agents-as-operating-systems framing.
- L. Wang et al. [A Survey on Large Language Model Based Autonomous Agents](https://link.springer.com/article/10.1007/s11704-024-40231-1) (Frontiers of Computer Science, 2024) — the agent landscape the paper cites.
- E. Moggi. *Computational lambda-calculus and monads* (1989) — the monadic model of effects the paper lifts to runtime.
- G. Plotkin, J. Power. *Algebraic operations and generic effects* (2003) — algebraic effects, the other pillar.
- T. Petricek, D. Orchard, A. Mycroft. [Coeffects: A Calculus of Context-Dependent Computation](https://www.cl.cam.ac.uk/~tpp26/papers/coeffects/coeffects.pdf) (2014) — the coeffect pillar the paper operationalizes.
- Related: [Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents) — structural access control for agent components; [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface) — the coordination layer should add nothing beyond the shared room.
- Related: [Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), [The Verification Horizon](https://blog.hackspree.com/#verification-horizon) — why guarantees must be structural, not remembered.
- Related: [Every Token Has a Price Tag](https://blog.hackspree.com/#every-token-has-a-price-tag), [Always-on agents](https://blog.hackspree.com/#always-on-agents) — the restart tax: state loss, warm-context loss, and the cost of the coarse-grained workaround.
- Related: [Agentic-first CLI design](https://blog.hackspree.com/#agentic-first-cli-design) — the harness as the agent's interface to the machine; [Software dark factories](https://blog.hackspree.com/#software-dark-factories) — agents generating the components this paradigm can compose.

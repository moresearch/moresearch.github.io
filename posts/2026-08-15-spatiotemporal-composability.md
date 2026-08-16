---
title: "Spatiotemporal Composability: The Missing Calculus for Self-Evolving Agents"
date: 2026-08-15
slug: spatiotemporal-composability
summary: "A Programming Paradigm for Spatiotemporal Composability (Peking University / DeepSeek-AI) is organized around five formal contributions plus a framing claim. This post lays out the key insights the way the authors present them — the two dimensions of dynamic composition, revertible effects, reactive coeffects, the unified context as a programming paradigm, and the calculus that composes local guarantees into system-wide ones — then reflects on each in turn: what the insights really buy, where the guarantees stop, and why the 'one door' context and the observational-equivalence honesty are the two things agent engineers should internalize first."
tags: agentic-design, dynamic-composition, spatiotemporal-composability, effects, coeffects, agent-harnesses, self-modifying-agents, plugin-systems, formal-methods, capability-security, sandboxing, koishi, cordis, deepseek
---

On August 13, 2026, an 88-page paper appeared on GitHub: *A Programming Paradigm for Spatiotemporal Composability*, by Yifan Shi and Wei Zhang of Peking University and Tianyi Cui of DeepSeek-AI, published from the [cordiverse](https://github.com/cordiverse/paper) repository. Within two days it had over 1,500 GitHub stars — an unusual reaction for a paper whose tools are monads, coeffects, and operational semantics.

The paper answers a question every engineer has felt but nobody has formalized: **how do you safely remove a piece of software from a running program?** Every program that supports plugins has a shameful button — in VSCode it is "Developer: Reload Window," a command that restarts the entire extension host so you can remove one extension. The button exists because the alternative is too hard: there is no safe way to unload a single extension's code, undo what it did to the program, and keep everything else running.

This post does two things. First, it lays out the paper's key insights **the way the authors present them** — their own contributions, quoted and translated into plain language. Then it reflects on each insight in turn: what it really buys, where its guarantee stops, and what it means for the problem this blog keeps coming back to — agents that modify their own harnesses.

## The problem the authors start from

The paper's motivation is a gap, not a bug report. Composition — assembling complex systems from simpler parts — is the most-studied topic in software engineering, but almost entirely in its *static* form: function calls, module imports, class inheritance, resolved at compile time and fixed for the life of the process. Dynamic composition — components loaded, unloaded, and reconfigured at runtime — has "theoretical foundations... underdeveloped, compared to the rich formal frameworks available for static composition."

The evidence that this is a real gap, not an academic one, is measured: among the top 100 VSCode extensions by install count, **87 contain executable code** and require a host restart to remove; only **7 declare dependencies** on other extensions at all. The industry's answer to the gap is the coarse-grained workaround: restart the process to get rid of a bad module, let the container orchestrator manage service dependencies. The workaround's costs are real and structural — every restart discards process-local state (caches, connections, partial computations) and rebuilding takes seconds to minutes; orchestration cannot express dependencies between components that share an address space, and turns local calls into network calls.

> The authors' framing claim: the field has a rich theory of static composition and none of dynamic composition — and the industry papers over the hole with restarts.

## The authors' key insights

The paper's own structure makes its claims explicit. One framing insight (the two dimensions) and five formal contributions, each building on the last.

### Insight 1 — Dynamic composition needs two dimensions that static theory never had to face

"To characterize the requirements of dynamic composition, we identify two orthogonal dimensions beyond the well-studied algebraic aspects of composition." The first is **temporal composability**: "upon removal of a component, the modifications the component made to the shared environment must be completely and safely reversed." The second is **spatial composability**: "components must be able to declare, discover, and resolve their dependencies on one another in a structured and verifiable manner."

The authors are careful to anchor both in familiar static settings: temporal composability reduces to lexical scoping — RAII, bracket patterns — and spatial composability reduces to module import resolution. The move is not to invent new requirements, but to notice that the old guarantees stop holding the moment components arrive and depart at runtime. A plugin's effects can outlive any function call, so the lexical scope that once ran destructors automatically is gone; a dependency can vanish mid-execution, so the one-time wiring the linker performed is no longer enough.

![Two guarantees for pluggable software — temporal: plug in, every action gets an undo; plug out, the undoes play in reverse, state returns, no restart, others untouched. Spatial: a component declares what it needs; when a provider changes, dependents reconnect or wait — a missing provider means "wait", never "crash". In static code these are ordinary (RAII, lexical scoping, module imports); at runtime they are the whole problem.](/images/composability-dimensions.png)

### Insight 2 — Effects can be made revertible (contribution 1)

The authors' first contribution: "We formalize revertible effects: every context transformation carries an explicit inverse that the runtime tracks, and both tracking and recovery preserve composition, so the context is recovered upon component removal."

In plain language: **effects** are what a component does to its world — allocating memory, registering a callback, writing to a store. The insight is that each such action should carry its own undo, and the runtime should log which component performed which effects, in what order. On removal, the runtime plays the inverses in reverse order — last effect undone first, like unwinding a stack. Because the inverse of a composite effect is the composition of the inverses, cleanup is *derived*, never written by hand. The author of a component supplies one line per atomic operation — the operation and its inverse — and the runtime composes them into correct, ordered teardown. The consequence the authors draw: this "establishes local temporal composability."

### Insight 3 — Coeffects can be made reactive (contribution 2)

The second contribution: "We formalize reactive coeffects: a component declares the coeffects it requires as a specification, and each change of the context notifies the component against that specification as activating, deactivating, or neutral."

In plain language: **coeffects** are what a component *needs* from its world — a database, a filesystem, an adapter. A component declares "I require a storage backend," and the runtime does the rest. Every change to the world is classified against each component's specification: *activating* (your need was just satisfied — start), *deactivating* (your provider was removed — wind down), or *neutral* (nothing you declared changed — don't react). A component whose dependency is missing stays inactive; it does not error. When the dependency appears, it activates. This "establishes local spatial composability."

### Insight 4 — Effects and coeffects are one context, and that is a paradigm (contribution 3)

The third contribution is the paper's boldest claim: "We unify the effect context and the coeffect context into a single context type, in which an observational equivalence on the coeffects supplies the effects with independence, constituting a programming paradigm for spatiotemporal composability."

The unification is not cosmetic. Because the context type is recursive and its coeffect part unconstrained, "any state the system needs to share across components can be encoded as a dependency with an appropriate value type — Σ subsumes all shared mutable states, not just inter-component dependencies." Every interaction between a component and its environment passes through a single object carrying three things: the current state, the undo log, and the dependency map. The paper's metaphor is literal: "loading a component corresponds to executing its effects (plugging in); unloading a component corresponds to recovering its effects (unplugging, without affecting other running components)," with hierarchical contexts enabling arbitrarily nested composition.

![One context — the only door: every interaction between a component and the world passes through a single object holding the current state, the undo log, and the dependency map. Components plug in (load: effects tracked) and plug out (unload: effects undone). One door means nothing can leak — and nothing can be missed.](/images/composability-context.png)

The paradigm claim is situated between the two poles of side-effect handling: functional programming threads state explicitly through every call (traceable, equational, and boilerplate-heavy), while imperative programming hides it (ergonomic, and effectively untraceable — React's `useEffect` registers a persistent side effect that appears in no parameter; Spring's `getBean` pulls from a process-wide registry with casts at every call site). The context paradigm, the authors write, "combines the traceability of the functional approach with the ergonomics of the imperative approach," and the payoff is their most important sentence: "In both directions, correctness that would otherwise rest on developer discipline becomes a structural property of the paradigm."

### Insight 5 — Local guarantees can be proven to compose (contribution 4)

The fourth contribution upgrades the paradigm from a feature to a foundation: "We give a calculus of dynamic composition, which combines the two mechanisms into the notion of a component and equips its lifecycle with an operational semantics. Its metatheory carries spatiotemporal composability from a single component to a whole system of interleaved components."

The calculus models components through a full lifecycle — loading, iteration, withdrawal, asynchrony, failure — and the metatheory proves the properties that matter: preservation, temporal composability, spatial composability, progress, and confluence. The point is compositional in the strongest sense: if every part can be removed cleanly and reconnects correctly, then a whole system of interleaved parts can. "Each part is safe to swap" is a local statement; "the system stays consistent through a storm of swaps" is the global one, and the metatheory is what connects them.

One honesty note sits inside this contribution, and it matters: recovery is guaranteed *up to observational equivalence*, not literal equality. The authors are explicit that unplugging a component does not restore the heap's exact layout — `free` releases a block without restoring the arrangement `malloc` left. What is guaranteed is that "no observer can distinguish" the recovered state from the original, where the observers are defined by the declared dependencies. What you expose as dependencies is what recovery promises to restore.

### Insight 6 — It is a shipped system, not a proposal (contribution 5)

The fifth contribution: "We implement these ideas in Cordis, a meta-framework of spatiotemporal composability that provides a core library realizing the formal model with effect tracking and coeffect resolution, as well as a declarative component loader with configuration reconciliation and hot module replacement."

And the implementation is not a demo. **Koishi**, the chatbot framework built on Cordis, has accumulated over **4,000 community-contributed plugins** over four years. The case study makes three claims: temporal composability "without cognitive overhead" (even an inexperienced plugin author gets ordered cleanup without writing an uninstall path); spatial composability across an open ecosystem (switching the storage backend reactivates only the dependents whose resolved dependency changed; a plugin whose dependency is unavailable stays inactive until it appears); and expressiveness plus generality (Koishi's web console is a second, independent Cordis application, the same model in a wholly different runtime). The paper concedes the threats to validity — a single ecosystem, a single host language, observational rather than controlled evidence — but the existence result is the point: the theory was extracted from a working system, not written first and applied later.

## Reflecting on the insights

The authors' contributions are the skeleton. What follows is the reflection — where each insight is genuinely important, and where its guarantee stops.

### On naming: the first deliverable is the two dimensions

The most undervalued thing the paper does is name the problem. "Reload Window" existed for decades as an accepted tax — nobody called it a symptom of missing theory, because nobody had a vocabulary for what was missing. Naming temporal and spatial composability, and grounding both in their static analogues (RAII, module imports), converts a felt inconvenience into a designable requirement. The restart workaround survived precisely because it was the only available mechanism at the right granularity; once the finer-grained requirements have names, the granularity mismatch stops being an accepted cost and becomes a problem with a theory attached.

This is also the right frame for the agent conversation. Every agent harness that "adds a tool" or "replaces a module" is doing dynamic composition — usually badly, by appending to lists and restarting processes. The two dimensions give agent engineers a checklist: when the agent removes something, is everything it did undone (temporal)? When the world changes, do the things that depended on it reconnect (spatial)? Most harnesses fail both checks, silently.

### On revertible effects: structure beats discipline — but only the orchestration is free

The deep move in contribution 1 is the *inversion of responsibility*: cleanup moves from the author, who must remember and usually gets it wrong, to the structure, which derives it. This is the same move as RAII — the destructor is automatic — extended from lexical scope to runtime lifetime. For autonomous agents this is not a quality-of-life improvement, it is the only realistic option: an agent cannot be instructed into remembering to clean up; a teardown *derived* from the load is the only teardown an autonomous system will reliably perform. Guarantees you must remember are guarantees an agent will, at some point, not.

But it is worth being precise about what the paradigm makes free and what it does not. The author still writes the inverse of each atomic operation. The paradigm eliminates the *composition* burden — the ordering, the bookkeeping, the edge cases of interleaved effects — not the *semantics* burden. If an effect's inverse is wrong, the recovery is wrong, and no amount of structure detects it. This is a quiet but important boundary: revertible effects guarantee that the teardown you supplied is *executed*, completely and in order; they do not guarantee the teardown you supplied is *correct*. The structural guarantee is about mechanics, not meaning.

### On reactive coeffects: availability becomes a lifecycle state, not an exception

The insight hidden inside contribution 2 is the failure-model change. "A missing dependency means inactive, not error" replaces an exception with a state. Most plugin systems discovered this organically through events and hooks; the paper's contribution is to make availability a first-class, typed, *composable* property of a component's lifecycle, with a notification language (activating / deactivating / neutral) that the runtime — not the component — drives.

For agents, this is the difference between a harness where a vanished tool crashes the agent, and one where the tool's dependents deactivate and reactivate as tools come and go. An agent that picks tools dynamically is going to experience provider churn constantly; a lifecycle that treats "your dependency is not here" as a normal state rather than a failure is a precondition for agents that gracefully lose and regain capabilities. It also quietly answers a question the blog has asked before about always-on agents: state that survives its providers.

### On the one door: auditability, policy — and trust

Contribution 3's unification is the most consequential and the least flashy. If every interaction passes through one context, then the context is simultaneously the undo log, the dependency resolver, and — because declared dependencies are known before a component runs — the access-control chokepoint. The paper draws the security conclusion itself: the dependency declaration "acts as a capability request, and the context proxy acts as a capability mediator," with the complete capability set known statically, before execution. For agent harnesses, whose components are not merely untrusted but *machine-written*, access control that is structural rather than inspected is the only kind that scales — this is the same conclusion this blog reached about sandboxing ([Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents), [Zero Overhead Is Zero Attack Surface](https://blog.hackspree.com/#zero-overhead-is-zero-attack-surface)).

But "one door" cuts both ways, and the reflection should say so. A single context is a single point of trust: the most security-critical object in the system is now one object, and the paper's own Section 6.3 concedes the limit — capability mediation is access control over *dependencies*, not isolation of *code*. "The second requires an external sandbox," the authors write, and that admission matters: the paradigm makes every effect attributable and every dependency mediated, but arbitrary machine-generated code still needs a real sandbox around the whole room. The context is the door; the door is not the walls.

### On observational equivalence and the system boundary: composability ends where the world begins

The most important thing to internalize is the least discussed: recovery means *indistinguishability*, and the boundary of what can be made indistinguishable is a design decision. The paper develops this as the *system boundary*: a location lies "inside" when the system can modify it exclusively and restore it; everything else lies "outside," and a coeffect moves the boundary by reifying an external location — confining access to operations that carry inverses, so that a file or a connection can become trackable.

For agents, this is where the theory meets the hard wall. Dynamic composition can make an agent's *harness* self-consistent — its in-memory registrations, its connections, its caches can all be undone. It cannot undo the world: an email sent, a payment moved, a file written outside the boundary. The guarantee is exactly as wide as the reification. The system boundary is where composability ends and accountability begins — and every agent engineer should read Section 6.1 as the map of that line. "What you expose as dependencies defines the boundary of what recovery can promise" is the paper's deepest design pressure, and it applies with double force when the code doing the exposing is generated by an agent.

### On compositionality and sufficiency: the mechanics of change, not the intelligence of change

The metatheory (contribution 4) is what upgrades the paradigm from "nice plugin framework" to "foundation," and it is the property self-modifying systems need most and get least: most agent harness engineering is empirical iteration, not compositional guarantee. But the reflection must end with what the paper does not claim. Spatiotemporal composability makes change *safe*; it does not make change *right*. An agent still needs to know what to build, whether it built it correctly, and whether removing it was the right call — the intelligence of change, which is a verification problem, not a composition problem. The blog's own line of argument applies ([Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc), [The Verification Horizon](https://blog.hackspree.com/#verification-horizon)): composition gives an agent the ability to modify its harness without breaking it; verification gives it the ability to modify its harness correctly. A self-evolving agent needs both, and the paper has built — with proofs — one of the two halves. That is why it matters: it is the first half of the foundation, finished, proven, and shipped in production, with 4,000 plugins as the evidence.

And the other half is now a well-posed problem, which is the highest compliment a theory can pay: the paper does not close the question of self-evolving agents, it makes the remaining question answerable.

## Bottom line

*A Programming Paradigm for Spatiotemporal Composability* is organized around the authors' five contributions — revertible effects, reactive coeffects, the unified context as a paradigm, a calculus whose metatheory composes local guarantees into system-wide ones, and a shipped implementation validated by 4,000 production plugins — built on the framing insight that dynamic composition has two orthogonal dimensions, temporal and spatial.

The reflections that matter for agentic design: the two dimensions are the checklist; structure beats discipline but only the orchestration is free; availability-as-state is the failure model agents need; the one-door context is the policy chokepoint but also the trust boundary, and it is not a code sandbox; recovery is an observational promise whose boundary the system must design; and composition makes change safe while verification makes change right. The paper is the mechanics of self-modification, proven and shipped — and the intelligence of self-modification is now a question with a theory behind it, instead of a button marked "Reload Window."

---

**References:**

- Yifan Shi, Wei Zhang, Tianyi Cui. [A Programming Paradigm for Spatiotemporal Composability](https://github.com/cordiverse/paper) (Peking University / DeepSeek-AI, 2026) — the paper; 88 pages; [paper.pdf](https://github.com/cordiverse/paper/blob/main/paper.pdf). The author quotes in this post are from Sections 1.1, 1.3, 3.3, 4.4, 5.3, and 6.1–6.3.
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

---
title: Always-on agents: state, memory, and the governance gap
date: 2026-07-07
slug: always-on-agents
summary: "A new survey of 435 papers argues that making agents truly always-on requires treating state as a first-class systems concern — not just remembering, but governing, recovering, and forgetting."
tags: agents, memory, state, governance, survey
---

Most of the agent conversation focuses on what happens *during* a task: tool calls, reasoning loops, correctness. Far less attention goes to what happens *between* tasks — the accumulated state, memory, permissions, commitments, and audit trails that persist across interactions.

A new survey from Ding, Nannapaneni, Liu, and Zhang ([Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306), June 2026) argues that this gap is the critical unsolved problem for deploying agents that operate continuously. And they back the claim with a coded analysis of 435 papers.

## What is an always-on agent?

The paper defines always-on agents as LLM-based systems whose future behavior depends on durable, accumulated state from past interactions. This state is not just retrievable memories. It includes:

- **Task ledgers** — what the agent has done, is doing, and has committed to do
- **Permissions and credentials** — what the agent is authorized to access, and how those authorizations change over time
- **Commitments** — promises made to users, other agents, or external systems
- **Provenance and audit records** — how each decision was reached, for post-hoc review
- **Shared state** — what multiple agents or an agent and its user both rely on
- **Trigger conditions** — latent rules that fire when certain conditions are met
- **Externally committed effects** — side effects the agent has already pushed into the world

This is a much richer picture than "the agent has a vector database of past conversations." An always-on agent is a persistent-state system. The persistence *is* the feature.

## The gap: we're good at memory, bad at governance

The survey's central finding is blunt: the literature is heavily skewed toward accumulating and retrieving state, with far less attention to how to govern, recover, or relinquish that state.

We have plenty of papers on retrieval-augmented generation, embedding-based memory, and context window management. We have far fewer on:

- **Forgetting.** When should an agent delete a memory? How do you ensure it actually forgets — including from backups, cached contexts, and fine-tuned weights? This connects directly to machine unlearning and to legal requirements like GDPR's right to erasure.
- **Recovery.** If an agent's state is corrupted — by a bad interaction, a prompt injection, a buggy tool — how do you roll back? What is the agentic equivalent of a database transaction?
- **Auditing.** If an agent made a consequential decision three weeks ago, can you reconstruct exactly what state it had access to at that moment, what it retrieved, and how it weighed that information?
- **Authority and scope.** Who can modify the agent's state? If an agent has learned a preference from user A, should user B's interactions be influenced by it? What happens when state from different sources conflicts?

The paper frames this as a maturity problem. We have built the memory layer for always-on agents. We have not yet built the governance layer. And you cannot safely deploy persistent-state agents at scale without both.

## Six diagnostic axes

The authors propose six axes for analyzing any piece of agent state:

1. **Authority** — Who or what created this state item? Who can modify or delete it?
2. **Scope** — Is this state private to one agent, shared across a fleet, or tied to a specific user?
3. **Mutability** — Can this state change? Under what conditions? Is it append-only, versioned, or freely overwritable?
4. **Provenance** — Where did this state come from? What chain of interactions produced it?
5. **Recoverability** — If this state is lost or corrupted, can it be reconstructed? From what?
6. **Actionability** — Does this state item directly drive agent behavior, or is it purely informational?

Most current agent frameworks score well on actionability (of course state drives behavior) and poorly on provenance and recoverability (good luck reconstructing why the agent did what it did six weeks ago). The axes give teams a checklist for auditing their own systems: for each piece of state your agent accumulates, can you answer all six?

## The lifecycle: state as a managed resource

Beyond the diagnostic axes, the paper introduces a lifecycle model for agent state. State is not just written and retrieved — it moves through a series of stages, each of which can fail:

- **Write** — state is created or updated
- **Validate** — is the state correct, consistent, not poisoned?
- **Organize** — how is state structured, indexed, deduplicated?
- **Retrieve** — the well-studied part: finding relevant state at decision time
- **Act** — state drives a decision or an external effect
- **Update** — the decision's outcome feeds back into state
- **Forget** — state is intentionally removed or decayed
- **Audit** — state is examined after the fact for correctness or compliance
- **Rollback** — state is restored to a prior version after a failure

The lifecycle exposes the asymmetry in current research. Write, organize, retrieve, and act are well-covered. Validate, forget, audit, and rollback are not. This means we are building agents that accumulate state aggressively and have almost no machinery for unwinding it when something goes wrong.

## AOEP-v0: governance as an evaluation target

One of the paper's more interesting contributions is the **Always-On Evaluation Protocol (AOEP-v0)** — a pilot evaluation contract that scores systems on state mutation and recovery obligations rather than answer quality.

This is a meaningful departure from standard agent benchmarks. Most evals ask: "Did the agent complete the task correctly?" AOEP-v0 asks questions like: "If we corrupt a piece of the agent's state, does it detect the corruption? Can it recover? If we issue a forget request, is the memory actually gone from all layers?" These are systems questions, not task-completion questions. They require testing the agent's governance machinery, not its reasoning quality.

The protocol is explicitly a v0 — early, incomplete, aspirational. But the direction is right. As agents move from demo to deployment, the evaluation that matters is not "can it answer questions" but "can you trust it to run for six months without accumulating dangerous state, leaking cross-user information, or becoming un-auditable."

## Why this matters now

The timing of this survey is good. Agent deployment is accelerating — from coding assistants to customer-facing autonomous systems. Each of these deployments accumulates state. Each one will eventually hit the governance questions the paper raises. And right now, the answers are mostly ad-hoc: prompt the agent to "be careful about stale information," log everything to a table nobody queries, hope for the best.

The paper connects always-on agents to mature disciplines that have already solved adjacent problems: databases (transactions, rollback, consistency), distributed systems (state reconciliation, quorum, fencing), capability security (authority, attenuation, revocation), and formal methods (invariants, verification). The claim — and I think it is correct — is that agent state governance is not a novel problem requiring novel solutions. It is a composition problem: we have the pieces, but we have not wired them together in the agent context.

This is a call to action. If you are building agent infrastructure, the question is not just "how does the agent remember?" It is "how does the agent govern its memory?" The second question is harder. It is also the one that will determine whether always-on agents are safe to deploy.

---

**Reference:** Tianyu Ding, Aditya Nannapaneni, Bingfan Liu, Ling Zhang. [Always-On Agents: A Survey of Persistent Memory, State, and Governance in LLM Agents](https://arxiv.org/abs/2606.30306). arXiv:2606.30306, June 2026.


Systems design is the core engineering discipline. Every system — whether a dark factory, an agent governance framework, or a software architecture — involves the same set of decisions: what are the components? what are their interfaces? what changes do we hide? what stays stable? The engineer who can answer these questions can design any system. The domain provides the constraints. The principles provide the method.


> An always-on agent is not a faster request-response loop. It is a different architecture. The agent that persists state across interactions is a system. The system must be governed. Governance is the hard part.

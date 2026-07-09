---
title: "Why DOT files work for dark factories"
date: 2026-07-08
slug: dot-files-dark-factories
summary: "Two independent dark factory tools — Kilroy and Mammoth — both chose Graphviz DOT as their workflow format. The reasons are practical, not theoretical: DOT is tiny, readable, visualizable, diffable, and editable by both humans and agents."
tags: dark-factory, dot, graphviz, agents, pipelines, workflow
---

Two dark factory tools landed on the same format independently: Graphviz DOT. Kilroy uses it. Mammoth uses it. Neither project seems to have coordinated on the choice. The convergence is worth understanding.

But first, what these tools actually do.

## Kilroy: the heavy factory

Kilroy is the heavier of the two — opinionated, factory-scale, closer to the full dark factory vision. Its workflow starts with English requirements: you describe what you want, and Kilroy converts that prose into a Graphviz DOT pipeline. It then validates the graph structure, and runs each node with coding agents operating in isolated Git worktrees. This is not a thin wrapper around an LLM call — it is a production mechanism.

The isolation matters. Each agent invocation runs in its own Git worktree, so agents cannot interfere with each other's file state. This is exactly the kind of engineering discipline the [task automation factory](/posts/task-automation-economics) paper calls for: execution that is replayable, inspectable, and isolated.

Kilroy records everything. Typed run events and artifacts go into CXDB — a structured event store for run history. Git branches hold the code history. The split is deliberate: Git tracks the code, CXDB tracks the run. You can resume a run from logs, from CXDB, or from a run branch. If something fails, you are not starting from scratch. You are rewinding to a known state and continuing.

This is the factory model made concrete: specification (English requirements) → pipeline (DOT) → isolated execution (Git worktrees) → event recording (CXDB) → resumability (logs, CXDB, run branch). Each stage is auditable. Each stage is replayable. Each stage leaves a record.

## Mammoth: the clean runner

Mammoth takes a different approach — lighter, more direct, but equally committed to DOT as the workflow format. It runs DOT-based LLM agent pipelines, validates graphs, supports checkpointing, and includes an HTTP server mode. Its architecture spans DOT parsing, run persistence, a web UI, a TUI, an MCP server, and a CLI — all built around an external tracker execution library.

The web UI is where Mammoth differentiates itself. It includes a spec builder, a DOT editor, and a pipeline runner — all in one interface. You can build the specification, edit the pipeline graph, and execute it without leaving the browser. This lowers the barrier to entry. Kilroy expects you to operate a factory. Mammoth gives you a control panel.

The checkpointing support means Mammoth pipelines can pause, save state, and resume — not as a side effect of Git branches and event stores, but as a first-class feature of the runner. This is a different philosophy from Kilroy's "logs + CXDB + run branch" model, but it solves the same problem: long-running agent pipelines need to survive failure.

## Why DOT?

Two independent projects, same format choice. The reasons are practical, not theoretical. Here is why DOT wins for dark factory workflows:

**DOT is tiny.** The entire DOT language fits in a README. Nodes, edges, attributes, subgraphs. That is basically it. You can learn the syntax in ten minutes. For a format that both humans and agents need to read and write, this minimal surface area is a feature. Every additional syntax element is something an agent can hallucinate.

**DOT is readable.** A DOT file is a list of relationships. `A -> B [label="depends on"]`. You can read it without tooling. You can reason about the graph structure by looking at the text. This matters when you are debugging a pipeline at 2am and do not want to fire up a graph visualizer to understand the control flow.

**DOT is visualizable.** Run it through Graphviz and you get a rendered graph. The same file that drives execution also produces documentation. No translation step, no sync problem — the pipeline definition *is* the diagram. For teams adopting dark factories, this collapses two artifacts (workflow spec + architecture diagram) into one.

**DOT already means graph.** The semantics align with the problem. A dark factory pipeline is a directed graph of tasks with dependencies. DOT is a language for describing directed graphs with attributes. The impedance mismatch is zero. You are not forcing a workflow concept into a format designed for something else — DOT was built for exactly this kind of structure.

**DOT works well in Git diffs.** Each edge is typically one line. Adding a node, reordering dependencies, inserting a new pipeline stage — these produce clean, readable diffs. Compare this to JSON or YAML, where a single added node can cascade through indentation changes across the entire file. DOT diffs tell you what changed at a glance.

**DOT lets humans and agents edit the same workflow.** This is the killer feature for dark factories. An agent can generate a DOT pipeline from a specification. A human can review it, edit it, add an edge, remove a node — all in the same format, with the same tooling, in the same file. There is no round-trip through a GUI. There is no intermediate representation that only one side understands. The DOT file is the shared artifact.

## The format is the interface

Dark factories need a format that sits at the boundary between specification and execution. The format is what the human reviews. It is what the agent generates. It is what the runner executes. It is what goes into version control. It is what you look at when something fails.

Most workflow formats are optimized for one of these audiences. YAML is optimized for machines (and debatably so). GUI-based workflow builders are optimized for humans (at the cost of diffability and agent-writability). Custom DSLs are optimized for a specific tool (at the cost of portability and learnability).

DOT is unusual because it is not optimized for any particular audience — it is just small enough, old enough, and general enough that it works for all of them. That is the definition of a good interface: it gets out of the way.

The convergence of Kilroy and Mammoth on DOT is not a coincidence. It is a signal. When two independent dark factory implementations both reach for the same format, the format itself is part of the design space. If you are building tooling in this area, the question is not "should I use DOT?" It is "is there a reason *not* to?"

---

**References:**
- [Kilroy](https://github.com/kilroy/) — English-to-DOT pipeline generation, Git worktree isolation, CXDB event recording
- [Mammoth](https://github.com/mammoth/) — DOT-based pipeline runner with web UI, checkpointing, MCP server
- [Graphviz DOT language](https://graphviz.org/doc/info/lang.html)

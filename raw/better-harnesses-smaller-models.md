---
title: "Better Harnesses, Smaller Models: 90% Cheaper Agents via Automated Harness Adaptation"
date: 2026-08-26
slug: better-harnesses-smaller-models
summary: "A close reading of the CMU paper (arXiv:2607.08938): the failure-mode→adaptation framework for harnesses, the meta-agent harness optimizer, why repetitive tasks and capable SLMs respond best, and what the 'model+harness pair' framing means for the economics of agent deployment."
tags: [slm, small-language-models, harness-engineering, harness-optimization, meta-agents, agents, cost-efficiency, gemma, qwen, model-swap, automated-design, amortization]
---

[Better Harnesses, Smaller Models: Building 90% Cheaper Agents via Automated Harness Adaptation](https://arxiv.org/abs/2607.08938) (Chenyang Yang, Xinran Zhao, Tongshuang Wu, Christian Kästner, Carnegie Mellon University, arXiv:2607.08938, July 2026) is the paper this blog has been circling for months: it takes the harness — the software around the model — and treats it as the primary lever for making agents affordable. The title says it: keep the cheap model, fix the harness, and recover frontier-LLM performance at 4% of the cost. Companion code: [github.com/malusamayo/migration-analysis](https://github.com/malusamayo/migration-analysis).

The setup is the null result every deployment engineer already believes. Take a budget-approval agent — collecting, reviewing, and communicating budget requests across departments. With a frontier model (`gemini-3.1-pro`), the generic harness hits 97.3% accuracy at $0.22 per query. Swap in a small open-weight model (`gemma-4-26b-a4b`, 4B active parameters) with no other changes, and accuracy collapses to 75.0%. That is the standard narrative: *SLMs just aren't good enough for agentic work.* The paper's bet is that the collapse is a harness artifact, not a model property. After automated harness adaptation — a step-by-step workflow skeleton in the system prompt, a filtered tool set, and an anti-loop hook — the same SLM hits **98.3%** on the same task, beating the frontier model, at 8% of the cost.

## The reframing: difficulty is shared, so lift it into the harness

The paper's core claim is almost economic before it is technical: **much of the difficulty of a routine business task is shared across instances.** Every budget request in a company goes through the same policy lookup, the same pricing table, the same reserve-account handling. That shared structure is what a frontier LLM privately reconstructs on every single run, token by token, at frontier prices. An SLM does the same reconstruction worse.

The fix: move that shared difficulty out of the model and into the harness — as tailored instructions, custom tools, and deterministic orchestration logic — once, offline, and amortize it over every instance. "The key insight is that much of the task difficulty is shared across instances and can be lifted from the model into the harness."

This is the same argument this blog has made about [harness engineering best practices](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) and [task-harness engineering](https://blog.hackspree.com/#task-harness-engineering): the model is the part that burns money at runtime; everything stable about the task should be encoded where it is cheap. What this paper adds is the *automation* — nobody hand-engineers the harness, a meta-agent searches for it — and the *empirics*: seven tasks, three SLM families, 21 model-task pairs, with and without adaptation.

## The framework: failure modes → adaptation strategies

Before automating anything, the authors build a diagnostic framework that maps *why* an agent fails to *what* the harness can do about it. Failures are indexed by the capability they implicate:

| Failure mode | What it looks like | Harness adaptations that address it |
|---|---|---|
| **tool-use** | wrong tool, malformed calls, bad composition across tools | demonstrate examples; wrap common sequences into simpler tools; filter the tool set; adapt schemas |
| **instruction-following** | violates task requirements, formats, workflow rules | reinforce constraints in context; enforce them programmatically via hooks |
| **knowledge** | missing domain/environment/commonsense assumptions | externalize knowledge into instructions; encode it into tool implementations |
| **long-context** | forgets earlier instructions, repeats failures, misses observations | reveal/compress/prune context; reduce observations; split into sub-agents |
| **planning/reasoning** | no viable decomposition, premature answers, no backtracking | provide plan skeletons; encode plans in tools; orchestrate agent flows |

Adaptation strategies are grouped by which harness component they change: **context** (add: examples, plan skeletons, externalized knowledge; manage: progressive reveal, compression, pruning), **tools** (create: wrap sequences, enforce constraints, encode knowledge; manage: filter, retune schemas), and **agent loops** (instrument code: deterministic checks, hooks, guardrails; orchestrate agents: multi-agent topologies).

The framework is deliberately size-agnostic — capabilities live on a spectrum, and today's SLM is tomorrow's mid-tier model. But the authors note the mismatch the framework exists for is most common exactly where models are small: frontier models are often strong enough that a generic harness suffices; SLMs exhibit capability gaps that make specialized harnesses the difference between 5% and 99% accuracy.

## The optimizer: harness search as a meta-agent loop

The design space is the [OpenHands Software Agent SDK](https://arxiv.org/abs/2511.03690): system prompts, skills, dynamic contexts, primitive and custom tools, hooks triggered on tool use, context management (external file systems, condensers), and sub-agents. The optimizer is a **meta-agent** (`gemini-3.1-pro-preview`) running an evolutionary loop, in the spirit of [self-improving agent workflows](https://blog.hackspree.com/#self-improving-agent-workflows) and the [Darwin Gödel machine](https://arxiv.org/abs/2505.22954):

1. **Sample and evaluate.** Maintain a pool of harnesses with validation scores; select candidates GEPA-style from the Pareto front — a harness that is optimal on at least one training instance — and run it on a sampled batch, logging full trajectories and outcome signals.
2. **Diagnose and propose.** The meta-agent inspects trajectories plus the current harness code and edits it. It gets four inputs: annotated trajectories, the harness implementation, a **search memory** of past proposals and their effects (so it stops rediscovering the same dead-end fix), and design-space documentation of available components. Proposals pass a cheap sanity check before evaluation, with a repair loop.
3. **Validate and keep.** If the edited harness improves on the batch, it gets a full validation run and enters the pool.

The loop is cheap by design: $20 of optimizer budget per task-model pair, three runs each, best validation harness kept — $1,260 total for the whole study. That number matters, because the entire economics of the paper rests on it: **a one-time offline cost, amortized across every production run.** The authors calculate the $20 is recovered after 13 runs on average.

## What the numbers say

The headline table compares three SLMs — `gemma-4-26b-a4b` (MoE, 4B active), `qwen3-coder-30b-a3b` (MoE, 3B active), `ministral3-8b` (dense) — with a generic harness and an optimized harness, against the frontier baseline, across seven tasks: attendance-auditing, budget-approval, stock-alert, anomaly-detection, playwright-testing, website-management, and code-refactoring (grounded in TheAgentCompany, LOCA-Bench, WebGenBench, WebArena, and RefactorBench, 20/20/60 train/validation/test splits):

| Configuration | Avg. accuracy | Avg. cost/instance | Avg. latency |
|---|---|---|---|
| `gemini-3.1-pro-preview` (generic harness) | 89.7 | $1.735 | 181s |
| `gemma-4-26b-a4b` generic | 31.4 | $0.043 | 328s |
| `gemma-4-26b-a4b` **optimized** | **80.2** | $0.071 | **135s** |
| `qwen3-coder-30b-a3b` generic | 26.9 | $0.085 | 107s |
| `qwen3-coder-30b-a3b` **optimized** | **74.8** | $0.064 | **76s** |
| `ministral3-8b` generic | 9.5 | $0.110 | 194s |
| `ministral3-8b` **optimized** | **25.0** | $0.099 | 172s |

**16 of 21 task-SLM pairs improved significantly with an optimized harness; seven pairs closed the SLM-LLM gap entirely.** The best result is `gemma-4-26b-a4b`: 89.7% of LLM performance at 4% of the cost, with 25% lower latency. The per-task numbers are wilder than the averages — the same SLM that scored 5.0% on anomaly-detection and 9.9% on playwright-testing with a generic harness hits **99.4% and 98.1%** optimized; `qwen3-coder` goes from 1.1% to 93.9% on anomaly-detection.

Two subtleties worth noting before celebrating. First, the cost win is mostly *not* lower per-token prices — the adapted harness is sometimes slightly more expensive per instance (added context, more deterministic scaffolding: $0.043 → $0.071 for gemma). The win is that the agent *finishes the task*, in fewer, shorter loops: latency collapses from 328s to 135s, and failures (which in production mean human triage or re-runs) mostly disappear. Second, the remainders are instructive: website-management stays hard even optimized (45.6% gemma, 24.4% qwen), and `ministral3-8b` — the weakest model — stays at 0.0% on stock-alert and anomaly-detection no matter what the harness does. **Harnesses can absorb task difficulty, but they cannot manufacture capability that isn't there.**

## When adaptation works

Two research questions pin down the boundary conditions, and both confirm the "shared difficulty" thesis:

**Repetitive tasks adapt well (RQ2).** The authors measure task diversity as the average pairwise Levenshtein distance between tool-call sequences across instances, and find a strong negative correlation with optimized performance (Spearman ρ = −0.96). A controlled experiment that varies the number of workflow templates in a task confirms causality: going from 3 templates to 20 drops optimized performance from 89.1% to 68.0%. A budget-approval agent where every instance follows the same procedure is a gift to harness adaptation; a code-refactoring task where every repository and query differs resists it. This is the quantitative version of this blog's [harnesses need tasks that fight back](https://blog.hackspree.com/#harnesses-need-tasks-that-fight-back) — low-entropy tasks are exactly the ones where a harness can win the fight.

**Capable SLMs benefit more (RQ3).** Using Artificial Analysis benchmark scores as a capability proxy, stronger models both perform better and improve *more* from adaptation (+48.8% vs +15.5% for the weakest). The harness offloads the *repetitive* parts of the task; the model must still follow the scaffolded instructions, use the tools correctly, and handle whatever cannot be delegated — and that residual is what separates SLMs.

## What the adaptations actually are

Analyzing the optimized harnesses through the framework (RQ4): the dominant addressed failure modes are **instruction-following (81%) and knowledge (81%)**, followed by tool-use (62%) and long-context (33%). The dominant strategies are **adding contexts (86%), creating tools (43%), and managing tools (29%)**. The best anomaly-detection harness, for example, combines all three: a custom `query_mock_bigquery` tool that sidesteps the default tool's long-context behavior, a filter from 40+ MCP tools down to seven, and the environment's table-naming convention externalized into the system prompt. The budget-approval harness hands the SLM a fixed plan, a filtered tool set, and a hook that refuses to let the agent send the same message twice.

Two findings here deserve emphasis. First, **harness adaptations do not transfer across models.** `ministral3-8b` struggles with file-editing tools; `qwen3-coder` occasionally emits raw XML where the protocol demands JSON tool calls. Each model's failure patterns demand tailored fixes — the playwright-testing harness for gemma adds contexts and pytest-enforcement hooks, while the one for ministral creates new tools to work around its editing failures. There is no universal SLM harness; optimization must be re-run per model (and, the paper's stance implies, per model *version*). Second, a notable negative result: **no optimized harness successfully used sub-agents.** The authors attribute this to current SLMs' limited ability to track and coordinate sub-agent work — which is a bracing counterpoint to the industry reflex that answers every SLM weakness with a multi-agent topology.

## Lessons for building harness optimizers

The meta-agent loop is itself a system, and the paper's development log doubles as practical guidance:

- **Diagnosis quality is the bottleneck.** Spending budget on evaluations or iterations only pays off if the meta-agent proposes good candidates. The authors found that passing *raw JSON* trajectories to the meta-agent beat post-processed markdown summaries — less human-friendly, but richer evidence for the frontier model to reason over.
- **The meta-agent's intelligence is worth its price.** A cheaper meta-agent (`gemini-3.1-flash`) afforded more iterations but produced worse final harnesses; lower-quality diagnoses more than offset the extra exploration. Interestingly, hand-crafting heuristics (showing successful frontier trajectories, manually mapping failure modes to fixes) did *not* consistently help — a strong frontier meta-agent, given faithful evidence and an editable harness, infers targeted repairs on its own.
- **Explore diverse regions, not one long trajectory.** Without search memory, the meta-agent rediscovers the same fixes and burns budget. And several independent shorter searches beat one long search: multiple runs cover more of the design space, raising the chance at least one discovers something strong.

## What this means

The paper's reframing — **the deployable unit is the model+harness pair, not the model** — is the takeaway that matters most. The industry's "model swap" mental model says you replace one model with another and re-validate. This work says the harness is a derived artifact, entangled with the model it scaffolds: the same task, same SLM, generic vs. adapted harness, spans 31% → 80% average accuracy. Swap in a new SLM and you must re-derive the harness; the optimization is part of deployment, not a one-time event. This also reframes *which* SLM to buy: the paper's heuristic — small MoE models like `gemma-4-26b-a4b`, chosen by benchmark score as a capability proxy — is simple and, on this evidence, effective.

Seen through the lens of [every token has a price tag](https://blog.hackspree.com/#every-token-has-a-price-tag) and this blog's running argument that harnesses are where agent economics are decided, the paper is a distillation story: a frontier meta-agent spends $20 once, compresses the shared structure of a task into a harness, and every subsequent run of the cheap model executes that structure nearly free. The expensive intelligence is amortized away; the cheap model becomes the run-time executor of distilled task knowledge. That is the same economics as [agents are distillation at scale](https://blog.hackspree.com/#agents-are-distillation-at-scale) — the harness is the distilled artifact and the frontier model is the teacher.

The honest limits are stated by the authors themselves: seven tasks with clean, verifiable success metrics; one optimizer implementation; a closed-source meta-agent that may drift as APIs evolve; list-price cost accounting. And the paper's own suggestion for the diverse tasks it fails — **mixtures of harnesses**, collections of specialized harnesses paired with a router, plus online adaptation as task distributions drift — is an open research direction, not a delivered result.

For practitioners the decision rule is clean. If you deploy agents on routine, repetitive business workflows and the bill for frontier tokens is the constraint, the paper's evidence says: take the best cheap MoE SLM you can find, spend a few dozen dollars on an automated harness search, and expect to recover most of the frontier model's accuracy at ~5% of the cost — provided you are willing to treat "re-optimize the harness" as a normal part of every model upgrade. The cheap agent is not a model you buy; it is a system you build. The harness is the product, and now it can be searched for automatically.

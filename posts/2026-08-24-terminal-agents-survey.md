---
title: "Terminal Agents: The Terminal Is the Substrate"
date: 2026-08-24
slug: terminal-agents-survey
summary: "A close reading of the Terminal Agents survey (arXiv:2608.20485): the terminal as execution substrate, the seven-dimension competence profile, and why benchmark scores hide process quality."
tags: terminal-agents, surveys, harness, evaluation, agents, swe-bench, cli
---

[Terminal Agents: A Survey of AI Agents in Command-Line Environments](https://arxiv.org/abs/2608.20485) (Yi Bin, Xiaoyang Yuan, Haoxi Zeng, Wencheng Ye, et al., arXiv:2608.20485, 52 pages) is the first survey that treats terminal-mediated execution as an object of study in its own right, rather than scattering it across software-engineering, tool-use, and computer-use literature. It is worth a close read because it formalizes a claim this blog has been circling for months: **the terminal is not an interface the agent uses, it is the substrate the agent lives in** (see [Agentic-First CLI Design](https://blog.hackspree.com/#agentic-first-cli-design) and the canonical [Harness Engineering: Best Practices for Reliable Agent Systems](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)).

The companion repo is [awesome-terminal-agents](https://github.com/EnigmaYYYY/awesome-terminal-agents).

## The scope move: substrate, not surface

The survey's organizing lens is deliberately narrow: a *terminal agent* is a system whose **dominant progress-bearing action–observation loop is mediated by terminal command execution, textual feedback, and stateful environment interaction**. Three workload-level boundary tests operationalize "dominant":

1. **Primary execution substrate** — command execution is the workload's main means of progress.
2. **Iterative command feedback** — outputs, errors, logs, diffs, return codes, or state changes materially shape later actions.
3. **Terminal dependence** — removing terminal access would change the workload's core behavior.

The consequences are sharp. SWE-agent is in scope; Agentless is not (static patch pipelines lack iterative execution); OSWorld is not (visual feedback is the progress-bearing substrate); a CLI-packaged assistant that forwards requests to a non-terminal workflow is not. The tests apply per-workload, not per-product — a platform can be in scope for repository repair and out of scope for its browser workloads.

This matters for harness engineers because **most "terminal agents" in industry are hybrid systems, and the boundary tests tell you which parts of them deserve terminal-specific evaluation**.

## Seven dimensions of terminal competence

The survey's analytical backbone is a seven-dimensional competence profile. These are not model capabilities; they are *system-level responsibilities* distributed across the model, interface, harness, runtime, and environment:

1. **Command and action formulation** — translating goals into executable commands, edits, and build/test/run actions.
2. **Feedback and artifact interpretation** — extracting evidence from stdout, stderr, exit codes, logs, diffs, stack traces.
3. **Runtime and environment management** — preparing, configuring, and repairing dependencies, services, containers, remote machines.
4. **State, task, and context tracking** — maintaining environment state and interaction history across extended sessions.
5. **Progress verification** — designing checks of intermediate validity and completion conditions.
6. **Recovery and adaptation** — diagnosing failures, replanning, retrying from grounded evidence.
7. **Governance and side-effect control** — permissions, sandboxes, approvals, resource limits, destructive-action prevention.

Two features of this list are worth internalizing. First, verification (5) is separated from recovery (6), and both from governance (7) — conflating them is why so many harnesses discover that "it passed the tests" and "it did something terrible" are simultaneous truths (see [Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents) and [Always-On Agents](https://blog.hackspree.com/#always-on-agents), which frames governing and recovering as first-class state concerns). Second, the dimensions are trace-observable in principle: planning appears as executable action sequences, memory is tested against persistent state, adaptation is grounded in external feedback. That traceability is the hook for everything that follows.

## Architecture: the harness is a performance-shaping component

The survey traces four shifts in design emphasis: tool-augmented prompting (ReAct, Toolformer, CodeAct) → structured executable actions (SWE-agent's ACI primitives, OpenHands) → terminal-mediated agency as a first-class target (Terminal-Bench, CLI-Gym, Endless Terminals) → **runtime- and harness-centered design** (Meta-Harness, AutoHarness, Agentic Harness Engineering), where context compaction, approval rules, observation shaping, and observability-driven harness evolution are treated as first-class variables.

Responsibilities are organized into four layers: interface and observation, runtime and workspace, control/verification/recovery/governance, and harness/context. The recurring tensions:

- **Expressiveness vs. recoverability** — raw command access is expressive but noisy and hard to roll back; ACI mediation is reliable but narrows the task surface.
- **Generality vs. task discipline** — platform runtimes span heterogeneous work but give less structured feedback.
- **Automation vs. inspectability** — permission gates and approval checkpoints aid auditability but add latency; governance is the least systematically addressed dimension.

And the one that should shape every comparison you publish: **attribution difficulty**. Controlled skill injection can add substantial token overhead without improving pass rate, and Agentless shows static pipelines rivaling interactive agents on some repository-repair tasks. Gains under an optimized harness may come from context management, observation shaping, retry policy, or injected procedural knowledge — not from a better model. **The complete model–harness–runtime configuration is the unit of comparison, not the model.**

## Acquisition: trajectories are the unit, failure is the signal

The relevant learning unit is a *stateful trajectory* — actions, observations, state changes, verification, recovery — not a prompt–response pair. The survey maps the acquisition levers:

- **SFT on successful traces** teaches common commands and workflows, but offers almost no supervision for diagnosis, rollback, or recovery after wrong assumptions.
- **RL with executable rewards** (Endless Terminals, ECHO, SWE-Master, SWE-Gym, Tmax) aligns behavior with completion, but sparse rewards can reinforce brittle or unsafe behavior.
- **Process-aware / verifier-guided optimization** supervises intermediate decisions via judgments, rankings, or hindsight validation (AgentHER is the key example — terminal failures usually appear early through stderr, failed tests, or inconsistent state).
- **Failure-conditioned training** (TRACE, AgentHER, AgentForesight) preserves diagnosis and recovery, but suffers noisy labels and repair-loop overfit.
- **Synthetic generation** broadens coverage of rare commands but risks teaching synthetic regularities; **runtime memory** (Memento, Context-Folding, TACO) sustains long horizons but risks persisting incorrect assumptions.

The uncomfortable finding: **recoverable failures remain underrepresented**. Failed installs, version conflicts, rollback decisions, and dead-end repairs are exactly the interactions needed to learn recovery — and successful-trace filtering removes them. If your training pipeline only keeps traces that worked, you are training an agent that cannot diagnose.

## Evaluation: outcomes are the floor, not the ceiling

The evaluation section is the most useful part of the paper for working engineers. Benchmark families expose *different things*:

| Emphasis | Representative examples | Blind spot |
|---|---|---|
| Repository repair | SWE-bench, SWE-PolyBench | Conflates repair with terminal competence |
| CLI/terminal-centered | Terminal-Bench, TerminalWorld, LongCLI-Bench | Mixes terminal-native and repo-mediated tasks |
| Setup | SetupBench | Isolated from end-to-end workflows |
| Process | OctoBench, ProcBench, AppWorld | No standardized process scoring |
| Long horizon | SWE-Bench Pro, LoCoEval, LifelongAgentBench | Costly, hard to reproduce |
| Safety/governance | BashArena, ClawSafety, AgentHazard | Task success may conceal harmful actions |
| Production | ProdCodeBench | Limited public access |

The evidence hierarchy goes beyond binary correctness: outcome evidence (task completion), process evidence (how execution happened), environment evidence (runtime validity), trace evidence (inspection/replay), and governance evidence (permissions, containment, side effects). Three protocol facts deserve wide circulation:

- **SWE-EVO-style long-horizon decay**: 65–73% on SWE-Bench Verified drops to 21–25% on multi-file evolution in its reported setting. Short tasks hide failure.
- **Reward hacking is measurable**: an audit reports 16% of tasks across five terminal-agent benchmarks are reward-hackable.
- **Harness effects dominate**: WildClawBench reports an 18-point harness-conditioned gap; SWE-rebench finds evidence consistent with contamination inflation on static tasks.

**Final success alone does not reveal whether an agent preserved state, recovered from failure, verified completion, or respected execution constraints.**

## The survey runs its own diagnostics

Rather than stopping at synthesis, the survey contributes bounded fixed-condition experiments — and these are the most concrete part of the paper. On a fixed configuration (mini-SWE-agent + DeepSeek-V4-Flash), it computes seven trace-derived process indicators (P1–P7) across four benchmark families:

| Benchmark | Tasks | Outcome | Final-verify (P5) | Gov. trigger (P7) |
|---|---|---|---|---|
| Terminal-Bench 2.1 | 241 | 52.6% | 29.5% | 0.5% |
| SetupBench | 93 | 59.1% | 36.6% | 1.4% |
| LongCLI-Bench | 21 | 23.8% | 23.8% | 4.5% |
| BashArena | 640 | 41.8% | 66.2% | 3.0% |

Key results: rule-matched invocation failures are rare (P1 ≈ 0.0%), local feedback use is high (76–82%), but **final-window verification spans 23.8%–36.6% on three benchmarks versus 66.2% on BashArena** — the same agent shows different process behavior because the benchmarks foreground different demands. The matched system comparison is even more pointed: SWE-agent wins every block, but the gap between systems ranges from at most 7.00 points on SWE-bench Lite to 21.25 points on Claw-SWE-Bench Lite, and the *ordering* of OpenHands and mini-SWE-agent **reverses across benchmarks**. Meanwhile the Flash-vs-Pro model variant never moves outcomes by more than 2.50 points. **Benchmark choice changes what you can see, and system comparison without benchmark context is meaningless.**

The trace cases ground the aggregates — the `magsac-install` case shows a dependency loop with a 38.5% environment-command non-zero-exit rate that never closes because the agent ends before an evaluator-relevant import check. Local feedback without verification and recovery leaves the loop unclosed.

```go
type TraceIndicators struct {
	InvocationFailures  float64 // P1: rule-matched command failure rate
	FeedbackUse         float64 // P2: helpful-use among feedback episodes
	EnvironmentExits    float64 // P3: non-zero exit on env-management commands
	StateErrors         float64 // P4: consequential state errors
	FinalWindowVerify   float64 // P5: verification in final window
	Recovery            float64 // P6: task-relevant recovery success
	GovernanceTriggers  float64 // P7: irreversible/overprivileged/secret-handling actions
}
```

This is the closest thing the survey offers to a portable process-evaluation schema — and it is exactly the shape of instrumentation this blog argues harnesses should ship with (see [Harness Engineering: Best Practices for Reliable Agent Systems](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents)).

## What to take away

The survey's own conclusion is the one to steal for your evaluation practice: **report system and runtime conditions explicitly, pair task outcomes with process evidence, and publish replayable traces.** The four research priorities it derives — cross-domain competence, fresh and replayable process-level evaluation, runtime governability, and controlled model–harness attribution — map one-to-one onto the gaps that keep showing up in real agent deployments.

For harness engineers the operative reading is: treat the terminal as substrate, instrument the seven dimensions, keep failure traces, and never publish a model comparison without pinning the harness, the runtime, and the benchmark. This is the survey version of "tasks that fight back" — except now the tasks are graded on process, not just outcomes, and the benchmark is treated as a measurement instrument that must resist gaming, the argument made in [Empirical Game Theory for Agents](https://blog.hackspree.com/#empirical-games-and-algorithmic-game-theory-for-agents). If you only track resolution rates, you are measuring the flattering version of your system. (Relevant adjacent reading: [DeepSeek Harness Notes](https://blog.hackspree.com/#deepseek-harness) and the [Harness Patterns for Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) index.)

## References

- Yi Bin, Xiaoyang Yuan, Haoxi Zeng, Wencheng Ye, Wenqi Shao, Chen Qian, Wei Ye, Yujuan Ding, Zheng Wang, Pengpeng Zeng, Jingkuan Song, Heng Tao Shen. [Terminal Agents: A Survey of AI Agents in Command-Line Environments](https://arxiv.org/abs/2608.20485). arXiv:2608.20485, submitted 20 Aug 2026. [Companion repo](https://github.com/EnigmaYYYY/awesome-terminal-agents).
- Key benchmarks discussed: [Terminal-Bench](https://arxiv.org/abs/2601.11868), [SWE-bench](https://arxiv.org/abs/2310.06770), [SetupBench](https://arxiv.org/abs/2507.09063), [BashArena](https://arxiv.org/abs/2512.15688), [LongCLI-Bench](https://arxiv.org/abs/2602.14337), [Claw-SWE-Bench](https://arxiv.org/abs/2606.12344), [ClawSafety](https://arxiv.org/abs/2604.01438).

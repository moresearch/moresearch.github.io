---
title: Agents Are Too Stochastic for Intuition
date: 2026-07-21
slug: data-driven-design-swe-agents
summary: You can't think your way to a better agent prompt. The only reliable design method for stochastic systems is measurement, experimentation, and data — the same method that turned databases from black boxes into predictable infrastructure.
tags: data-driven-design, agents, evaluation, harness, software-engineering, metrics
---

Databases used to be black boxes. You wrote a query, waited, and hoped the optimizer did something reasonable. When it didn't, you tried a different index and hoped again. The expertise was real but the method was folk — passed from senior to junior as rules of thumb: "always index foreign keys," "avoid `SELECT *`," "`EXPLAIN` will tell you what's wrong."

That era ended when database teams started treating query performance as a measurement problem. Instrument the optimizer. Collect query plans. Compare actual vs. estimated row counts. Build dashboards. The output of all that instrumentation is a design loop: you don't guess why a query is slow. You look at the data. The data tells you.

We are at the beginning of the same transition for software engineering agents. Right now, most agent design is folk expertise: "be specific in your prompts," "few-shot examples help," "chain-of-thought improves reasoning." These are all true, sometimes, for some models, on some tasks. The problem is that nobody knows which ones.

> The defining quality of an LLM is that it is stochastic. Given the same input, you get different outputs. Given slightly different inputs, you get very different outputs. The only way to reason about a stochastic system at scale is with data.

## Why intuition fails

Your intuition about what makes a good agent prompt is shaped by a few dozen interactions. You try a prompt, it works, you generalize. You try another, it fails, you adjust. This is the scientific method at N=1 — useful for generating hypotheses, useless for drawing conclusions.

The problem compounds because agents are not just stochastic — they are *stateful, tool-using, multi-turn stochastic systems*. An agent that picks the wrong tool on step 2 fails on step 7 in a way that looks unrelated to the original mistake. Intuition attributes the failure to the step-7 behavior. The data would tell you to fix step 2.

> Debugging agents by reading transcripts is like debugging a database by reading query text. You might spot the obvious problems. You will miss everything that matters at scale.

## What data-driven design looks like for agents

The method is straightforward, even if the infrastructure isn't:

**One: Instrument every decision.** The agent chose a tool. Was it the right tool? The agent produced a diff. Did it compile? Did the tests pass? Did it introduce a vulnerability? The agent looped. How many iterations? When did it converge — or fail to? Every one of these is a data point. Collect them.

**Two: Define success numerically, not narratively.** "The agent handled the task well" is a narrative. "The agent solved 73% of tasks in the benchmark, with a median of 4 tool calls per task and a 6% hallucinated-tool-call rate" is data. Narratives are for demos. Data is for design.

**Three: Compare at scale.** You changed the prompt. Did the solve rate go up? By how much? Across how many tasks? With what variance? A single run of a stochastic system is a data point. A thousand runs is a distribution. Design decisions should be made on distributions, not data points.

**Four: Close the loop.** The best agent systems — like the best database administrators — don't make one-off optimizations. They build pipelines. Deploy. Observe. Compare. Improve. Deploy again. The loop is the design process. Speed of iteration matters more than brilliance of a single change.

## The harness is your measurement instrument

A theme that runs through this blog is that the harness is the most under-invested part of an agent system. Data-driven design makes it concrete why: the harness is your measurement instrument. If your harness can't produce reliable, comparable metrics at scale, you cannot do data-driven design. You are back to intuition and anecdotes.

> A harness that can only tell you "pass" or "fail" is a thermometer with one marking. You need to know what happened, where it went wrong, and whether this run was better than the last run in a way that generalizes.

The harness engineering practices that matter for data-driven design:

- **Deterministic replay** — if you can't reproduce a run, you can't measure a change.
- **Layered observability** — prompt selection, tool choice, state transitions, final outcome. Each layer produces its own metrics.
- **Statistical rigor** — a 2% improvement on 10 tasks is noise. A 2% improvement on 1,000 tasks might be real. The harness needs to support sample sizes that make signals distinguishable from noise.

## Why this matters more for agents than for any previous software

Traditional software is deterministic. You change the code, you run the tests, you observe a binary outcome. The design loop is: implement, test, fix, repeat. Data helps, but the feedback signal is strong enough that intuition often suffices.

Agents are fundamentally different. The same agent, on the same task, with the same prompt, will produce different results on different runs. The space of possible behaviors is too large to explore by reasoning alone. The only way to understand an agent is to observe it across many runs and let the patterns emerge from the data.

This is uncomfortable. It means admitting that you cannot fully understand the system you are building by reading the code and the prompts. It means ceding authority to measurement. It means designing experiments instead of designing solutions — because the solution emerges from the experimental data, not from your head.

> The transition from folk expertise to data-driven design is the transition from craftsmanship to engineering. Both produce good outcomes. Only one scales.

## The cost of not doing it

The alternative to data-driven design is design by anecdote. Someone runs the agent on a task they care about. It works. They ship. Someone else runs it on a different task. It fails. Nobody knows why. The team debates prompts instead of looking at data. Decisions are made by the most senior person in the room, not by the person with the best evidence.

This is how most teams work today. It is also how most teams will fail at building agents, because the complexity of stochastic, tool-using, multi-turn systems exceeds what any individual's intuition can model.

> Data-driven design is not about having data. It's about making decisions as if the data matters more than your opinion.

Databases went through this transition. So did compilers. So did networking. Every infrastructure layer that we now treat as predictable, measurable, and engineer-able went through a phase where the experts relied on intuition and the results were inconsistent. Agents are in that phase now. The teams that instrument, measure, and close the loop will build infrastructure. The teams that rely on folk wisdom and anecdote will build demos that don't survive contact with real workloads.

---

**References:**

- Selinger, P. G. et al. (1979). ["Access Path Selection in a Relational Database Management System."](https://doi.org/10.1145/582095.582099) *ACM SIGMOD.* — The paper that made query optimization a measurement problem: cost-based plan selection using catalog statistics instead of heuristics. The origin of `EXPLAIN`.
- Kohavi, R., Tang, D. & Xu, Y. (2020). *Trustworthy Online Controlled Experiments: A Practical Guide to A/B Testing.* — The modern methodology for data-driven design decisions at scale, from the team that built Microsoft and Google's experimentation platforms.
- Sigelman, B. H. et al. (2010). ["Dapper, a Large-Scale Distributed Systems Tracing Infrastructure."](https://research.google/pubs/dapper-a-large-scale-distributed-systems-tracing-infrastructure/) *Google Technical Report.* — The paper that established distributed tracing as a measurement primitive, now industry standard (OpenTelemetry).
- Kleppmann, M. (2017). *Designing Data-Intensive Applications.* — Chapters on observability, metrics, and the shift from intuition to measurement in distributed systems design.
- Related: [Harness Engineering: Best Practices for Reliable Agent Systems](https://blog.hackspree.com/#harness-engineering-best-practices-for-ai-agents) — This blog's framework for building evaluation harnesses that produce the reliable, comparable metrics data-driven design requires.
- Related: [Correctness First: What OpenBSD Teaches Agent Builders](https://blog.hackspree.com/#correctness-first) — The argument that correctness can't be eyeballed — it must be verified systematically, which requires measurement.
- Related: [In the Land of AI Agents, the Verifiers Are King](https://blog.hackspree.com/#verifiers-are-king-sonar-acdc) — Sonar's AC/DC framework as a productization of the same principle: verification infrastructure is the measurement layer for agent quality.

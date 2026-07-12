---
draft: true
title: Harness Engineering (Martin Fowler)
date: 2026-05-17
slug: harness-engineering-fowler
summary: "Reflection and practical takeaways from Martin Fowler's 'Harness Engineering' — designing feedforward/feedback controls and harness templates to govern coding agents."
tags: [harness, synthesis, martin-fowler]
---

This post has been consolidated into [Harness Engineering: Best Practices for Reliable Agent Systems](#harness-engineering-best-practices-for-ai-agents). See that canonical post for the full, merged guidance.
Martin Fowler's essay "Harness Engineering" reframes the problem of trustworthy coding agents as a control-engineering task: the harness is the set of guides and sensors that steer, test, and correct agent behavior so humans can rely on the outputs.

This post summarizes the article and extracts practical recommendations teams can start applying today.

Key idea: harness = guides + sensors

- Guides (feedforward controls) shape agent behavior before it acts: clear instructions, bootstrapping docs, templates, and computational tooling that bias the agent toward desirable outputs.
- Sensors (feedback controls) observe results and enable self-correction: linters, tests, static analysis, and higher-level inferential judges (AI reviewers) that can produce signals optimized for the agent.

Computational vs inferential

- Computational controls are deterministic and fast (linters, unit tests, static analysis). They are cheap enough to run on every change and give reliable results.
- Inferential controls are semantic, non-deterministic, and more expensive (LLM-based code review, AI judges). They add richness and can catch problems that structural checks miss, but they should be used judiciously.

Steering loop and timing

Fowler emphasizes the human role: the team must "steer" the harness by iterating on guides and sensors whenever recurring failures are observed. A practical cadence is:

- Shift cheap, computational checks as far left as possible (pre-commit, IDE/LSP, early CI).
- Run heavier inferential sensors in integration or post-merge pipelines.
- Use continuous drift detection (coverage, dependency analyses, runtime SLO sampling) outside the commit lifecycle to surface gradual regressions.

Three regulation categories

Fowler proposes thinking of harness objectives in three orthogonal dimensions:

- Maintainability harness — tools that regulate code quality and internal consistency (duplicate code, complexity, coverage).
- Architecture fitness harness — fitness functions and structural checks that enforce architectural constraints and observable patterns.
- Behaviour harness — the hardest: verifying that the system functionally does what users need; often needs a mix of tests, approved fixtures, and human checks.

Harnessability and templates

Not every codebase is equally easy to harness. Greenfield projects can bake harnessability in (types, clear module boundaries); legacy codebases may need incremental investments. To scale effort, organizations can capture common topologies as harness templates — bundles of guides and sensors tailored to a topology — while being mindful of version drift and maintenance costs.

Role of humans

Humans bring tacit knowledge, social accountability, and judgement that agents lack. Harnesses are partly an attempt to externalize that experience; they should therefore aim to reduce repetitive oversight while leaving nuanced judgement where people add the most value.

Practical checklist

- Start with computational feedforward controls: LSPs, linters, and targeted pre-commit hooks.
- Add fast feedback: unit tests, structural tests, policy linters that run in CI.
- Design inferential sensors only where they measurably reduce supervision cost; run them at lower frequency or post-merge.
- Prioritize harness cases that map to production failures — convert incidents into permanent harness cases.
- Instrument for replay and for agent-friendly error messages so sensors can be used as correction signals.
- Monitor continuous drift and add sensors that surface slowly accumulating problems.

Conclusion

Fowler's framework makes the abstract idea of "a harness" actionable: treat it as a system of guides and sensors that you iteratively improve, and prefer cheap deterministic controls left in the pipeline while layering inferential checks where they add value. Harness engineering is not a one-off project; it's an ongoing practice that makes coding agents a practical part of the engineering workflow.

References

- Martin Fowler, "Harness Engineering" — https://martinfowler.com/articles/harness-engineering.html


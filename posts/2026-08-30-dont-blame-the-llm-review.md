---
title: "Don't Blame the LLM: A Review of the First Controlled Longitudinal Study of Harness Evolution"
date: 2026-08-30
slug: dont-blame-the-llm-review
summary: "Review of Ben Sghaier, Li, Adams & Hassan, 'Don't Blame the Large Language Model: How Agent Harness Evolution Shapes Coding Agent Quality' (arXiv:2607.03691) — the first controlled longitudinal study that fixes the model and varies only the harness: 35 sequential Qwen Code CLI releases on SWE-bench Verified. Verdict: resolve rates are flat while token consumption nearly doubles, every regression passed CI, and the culprit is the absence of agentic QA. The scaffold moves the quality; the model was held still."
tags: [review, harness-engineering, agent-harness, longitudinal-study, swe-bench, software-evolution, quality-regression, token-economics, empirical-software-engineering, qwen-code]
---

*[Don't Blame the Large Language Model: How Agent Harness Evolution Shapes Coding Agent Quality](https://arxiv.org/abs/2607.03691)* — Oussama Ben Sghaier, Hao Li, Bram Adams, Ahmed E. Hassan (Queen's University), arXiv:2607.03691v2 (cs.SE), July 2026, submitted to ACM TOSEM.

## The complaint this paper formalizes

Every coding-agent user knows the ritual. An update arrives — often silently, because most agent CLIs are configured to auto-update or heavily prompt you to upgrade. The agent starts feeling worse: dumber in conversation, hungrier for tokens, stuck in loops it used to escape. You open the issue tracker and write exactly what everyone else wrote: "the model got worse." The GitHub threads the paper cites are a genre: [Cursor is getting worse and worse](https://forum.cursor.com/t/cursor-is-getting-worse-and-worse/66070), [Claude Code token consumption spike](https://github.com/anthropics/claude-code/issues/16856), [Codex harness change causing quality cliff](https://github.com/openai/codex/issues/8272). In rare cases, a user recommends pinning an old harness version. Everyone blames the model.

This paper is the controlled version of the counter-claim in its title: *don't blame the LLM.* Where every prior study fixed the harness and varied the model, Ben Sghaier et al. do the mirror — **fix the model, vary only the harness** — across 35 sequential releases of the Qwen Code CLI, and measure what moves. The answer: the harness moves quality, and the model was held still. It is the longitudinal companion to the [Inside the Scaffold taxonomy](https://blog.hackspree.com/#inside-the-scaffold-review) reviewed here last week: that paper mapped the scaffold's design space at a point in time; this one watches a single scaffold evolve and measures what the evolution does. And it is the exact empirical form of this blog's standing claim that the [model-swap failure is a harness artifact](https://blog.hackspree.com/#better-harnesses-smaller-models) — except here the harness is the moving part.

## The experiment

The design is clean and expensive. The authors self-host Qwen3-Next-80B-A3B-Instruct (an 80B MoE, 3B active) on a dedicated server via vLLM, exposing an OpenAI-compatible endpoint, so the LLM is bit-for-bit identical across every release — no silent model updates, no API rate-limit noise, no prompt-caching drift. Against that fixed model, they run 35 sequential releases of the Qwen Code CLI (v0.0.10 through v0.10.3 — chosen because Qwen Code natively supports local endpoints and is *designed for* the Qwen model family, making the pairing favorable rather than adversarial), each evaluated on 50 tasks stratified by difficulty from SWE-bench Verified, twice per release. That is 3,500 inference runs plus full patch evaluation in Docker. They measure effectiveness (resolve rate) and efficiency (token consumption, tool calls), and they verify their ground first: the two runs agree on 87.7% of task outcomes, with no statistical difference between run distributions — the fluctuations they report are harness effects, not LLM stochasticity. Most agent evaluations run once; this one earns its confidence interval.

## RQ0 — Hyper-churn: the scaffold is the fastest-moving software you deploy

Before measuring quality, the paper measures velocity, across five harnesses (Gemini CLI, Codex, OpenCode, OpenHands CLI, Qwen Code) against two mature baselines (VS Code, GitHub CLI) over the same calendar window. The result is a new term: **hyper-churn.**

- **Releases:** OpenCode ships 18.0 releases per week (median 0.12 days between releases, peaking at 136 in a single month), Codex 12.4, Gemini CLI 10.3, Qwen Code 10.0 — 13–28× the cadence of VS Code (0.8/week) and GitHub CLI (0.6/week). OpenCode averaged 39.1 patch releases per minor version.
- **Commits and review:** 13–34 merged commits per day, with median PR review times under four hours, and only 7–14 commits per release (VS Code: 421.4). Changes move from merged to shipped in under a day.
- **Backlogs:** Gemini CLI accumulated 9,951 issues in 224 days; OpenCode closes only 54% and its backlog approaches 4,000.

The practical consequence is one this blog has spent a lot of time on: the harness is not infrastructure you install once. It is a fast-moving software system that **reaches your machine on its own release schedule**, and auto-update means you ride every release whether or not you were asked. The question the rest of the paper answers: does that ride improve the agent?

## RQ1 — The flat quality line

The headline result deserves the headline treatment:

> **Across 35 releases, resolve rate shows no statistically significant trend (Spearman ρ = 0.208, p = 0.231), fluctuating around a mean of 30.5% — while token consumption rises significantly (ρ = 0.743, p < 0.0001), from ~391K tokens/task in the first nine releases to ~668K in the latest: a 70% increase with no corresponding gain in resolution.**

The quality peaked early — 39% at v0.0.14 — and never recovered. Later releases are not better; some of the earliest are the best. And the resource inflation is not an artifact of harder task sampling: the authors normalize each task against its own per-task baseline across all 35 versions, and the trend survives (latest releases +19.6% above baseline, early ones −24.2% below). Tool calls climb the same way, from 6.9 to 14.3 per task.

The mechanism of the inflation is worth stating precisely, because it is architectural, not incidental. The initial prompt payload (system prompt + tool schemas) grew ~8% across releases — with the *task description held constant*, so the growth is pure harness bloat. Meanwhile newer releases require 18% more LLM turns, and since the harness prepends the full conversation history at every API call, a bigger base prompt is re-incurred on every turn. Token consumption and turns correlate at ρ = 0.941: the cost compounds.

Two case studies make it concrete. **v0.1.4 → v0.1.5**: token consumption surges 52% (216.6K → 329.9K) while resolve rate stays exactly 26.0%. The inflation traces to two PRs — a search-tools rewrite (+834/−776) and a tool-output-formatting restructure (+795/−607). **v0.2.3 → v0.3.0**: tokens more than double (+131%, 259K → 599.6K) via a +14,550-line streaming-layer PR and +4,740 lines of i18n infrastructure injected into the prompt path. Both releases passed every unit and integration test in CI. No automated check flagged a 52% or a 131% cost increase, because no check measured cost.

There is also a grim efficiency fact buried in the results: **unresolved tasks consume 2.7× the tokens and 1.8× the tool calls of resolved ones** (697.7K vs 258.7K tokens; 12.95 vs 7.2 calls). Correct fixes are recognized and executed quickly; failing tasks trap the agent in fruitless edit-test-read loops. This is the [token snowball](https://blog.hackspree.com/#every-token-has-a-price-tag) made measurable: the agent's own failure loops are where the budget goes.

## RQ2 — Which release patterns explain the shifts

The paper then asks what kind of releases produce the good and bad versions, across 22 release-level factors (churn, commit composition, contributor activity, issue health), using Spearman correlations plus Good/Neutral/Bad tier comparisons with Cliff's delta, all Benjamini-Hochberg corrected. The findings are a development-policy primer for harness maintainers:

- **Feature-heavy releases are the only thing that improves resolve — at a cost.** Feature churn correlates with higher resolve (ρ = 0.438) but also with higher tokens (ρ = 0.401, d = −0.573) and tool calls (d = −0.900). The single largest improvement in the study, v0.5.0 → v0.5.1 (27% → 34%), came from a *five-line change* to sampling-parameter defaults plus tool-schema compliance — a reminder that features ≠ quality.
- **Fix-heavy releases cost without gaining.** More bug-fix churn predicts higher token consumption with no resolve benefit; the largest delta effect in the study (d = −0.833). Edge-case handling accumulates prompt burden.
- **Refactoring is negatively associated with resolve** (ρ = −0.372) and buys no efficiency.
- **Consolidation helps.** Releases that ship larger, coherent PRs cost less (mean PR size negatively correlates with tokens, ρ = −0.441, and tool calls, ρ = −0.484, in the delta analysis) — many small PRs add configuration and edge-case overhead; bigger PRs are self-contained. And **deleting code pays**: more deletions predicts lower token consumption (ρ = −0.351); the v0.4.1 → v0.5.0 cleanup that removed 817 lines of unused harness code cut tokens 12.9%.

## RQ3 — Where the regressions live

The final analysis maps every commit to a ten-component reference architecture (derived, in the tradition of Hassan and Holt, from all five harnesses: UI, Orchestrator, LLM Provider, Tool System, Context Management, Persistence, Security, Extensibility, Config, Communication Backbone) and computes partial Spearman correlations controlling for total release churn, so component effects aren't confounded by release size. The result is a risk map:

- **High-risk zones: the LLM Provider and Context Management.** These are the layers that govern what reaches the model, and they carry the most regression risk. The v0.4.1 → v0.5.0 transition is the cautionary tale: a +346/−8-line change to the OpenAI request/response converter, invisible to tool-level integration tests, dropped resolve from 39.4% to 32.5%. Context Management expansion correlates with *worse* token efficiency (ρ = −0.346) — more context logic means more information presented per turn, without better reasoning.
- **Safe zones: Extensibility and Security fixes.** Hook and plugin work correlates with better token efficiency; targeted security fixes correlate with efficiency gains on both tokens (ρ = 0.346) and tool calls (ρ = 0.341) without regressions — security fixes tend to tighten code paths rather than widen them.
- **Persistence rework pays.** Stabilized session/state management correlates with efficiency gains, plausibly because the LLM stops re-processing redundant history across turns.

## The diagnosis: no agentic QA

The unifying explanation is the paper's most actionable contribution. Qwen Code maintains **300+ unit, integration, and end-to-end tests — and none of them evaluate the agent**. No resolve-rate check, no token budget, no tool-call ceiling. Every quality regression the paper documents — the 52% token spike, the 131% spike, the converter-caused resolve drop — passed all existing automated checks, because those checks verify that the software *functions*, not that the agent *performs*. The authors call the missing practice **non-functional agentic regression testing**: automated evaluation of effectiveness and efficiency metrics across versions, run as part of CI. The reason it doesn't exist yet is honest and structural: running representative agent evaluations at a 10-releases-per-week cadence is expensive. But the cost of not running them is what the paper measured — quality that drifts in production while the dashboard stays green.

This is the [empirical-SE discipline](https://blog.hackspree.com/#empirical-se-what-studies-say) applied to the one place this blog keeps insisting it must be applied: the harness. If the deployable unit is the [model+harness pair](https://blog.hackspree.com/#better-harnesses-smaller-models), then harness changes are quality changes, and they must be gated on harness-level evaluation — not on whether `npm test` passes. The paper's prescription is the engineering form of this blog's [harness-patterns series](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems) thesis: the harness is the product, so the harness needs its own QA loop, with its own budgets.

## What's missing

The paper is unusually disciplined about its own limits, and the review should name them too.

**One harness, one model, one benchmark.** The longitudinal study is a single system — Qwen Code CLI — paired with a single, *favorable* model (Qwen's own, on Qwen's own harness). This is the right design for control and the wrong one for breadth: the authors acknowledge that a stronger model might be more robust to prompt bloat and a weaker one more sensitive, and that the quality findings may not transfer to Codex, OpenHands, or SWE-agent. RQ0 shows the *conditions* (hyper-churn) are industry-wide, but the *effects* were measured on one system.

**Statistical power is modest.** 35 releases (34 transitions for the delta analysis) with BH correction is a small sample; the paper reports effect sizes precisely because significance is hard to come by at this n. The flat-quality finding is robust; the RQ2/RQ3 correlations are suggestive and internally consistent rather than definitive.

**Quality means SWE-bench.** Resolve rate is a binary pass/fail with no partial credit, task difficulty is proxied by human fix time, and the whole quality definition is benchmark-bound. Real-world interactive quality — whether the agent makes a developer faster or more pleasant, whether UX regressions matter — is out of scope. And the 600-second per-task timeout, a pragmatic cost bound, interacts with the token-burn finding: runaway loops are capped, so the *measured* inflation is if anything conservative.

**The replication package is pending.** The paper cites its package as "to be made publicly available upon acceptance" — the PR-level traces and per-release logs are promised, not yet public. Given how much of the paper's credibility rests on checkable evidence, that is a gap to close.

## Bottom line

This paper converts the industry's most common blame story into a controlled result: hold the model still, and the harness alone moves resolve rates, doubles token consumption, and ships regressions that pass every existing test. The finding is not "don't update your harness" — it is that **the harness is quality-critical software whose quality is currently invisible**, and the fix is a QA loop that measures agent-level outcomes across releases, with budgets on tokens and tool calls, before auto-update pushes the next release onto everyone's machine.

Paired with the [Inside the Scaffold taxonomy](https://blog.hackspree.com/#inside-the-scaffold-review), the two papers form a complete picture: the first is the map of what harnesses are made of; this one is the clock showing that the map's features change faster than anyone is measuring. The scaffold moves the quality. The model was held still. The dashboard was green the whole time.

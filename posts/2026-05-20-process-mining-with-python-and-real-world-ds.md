---
title: Process Mining with Python and Solving Real‑World Data Science Tasks
date: 2026-05-20
summary: Practical notes combining process mining techniques in Python with pragmatic data‑science workflows; inspired by two Medium posts.
tags: [data-science, process-mining, python, tutorial]
---

TL;DR

Process mining turns event data into process models and performance insights; Python (pandas + PM4Py) makes it accessible. Pair process‑mining features (throughput, wait times, activity counts) with standard data‑science pipelines (EDA, feature engineering, modeling) to solve real‑world problems like delay prediction and bottleneck analysis. This post synthesizes practical steps and code pointers inspired by two Medium articles: "Process Mining with Python" and "Solving a real‑world data science task with Python." Links in References.

Introduction

Two approachable Medium posts highlight hands‑on ways to extract insights from logs and run pragmatic data‑science projects from end to end. This post synthesizes their practical guidance into a compact recipe: how to extract event logs, discover process models, compute process features, and use them in predictive workflows.

1) From raw events to an event log

Key columns: case id (process instance), activity name, timestamp. Start by loading data with pandas, parsing timestamps, and normalizing column names for PM4Py interoperability.

Example:

```python
import pandas as pd
from pm4py.objects.conversion.log import factory as log_converter

df = pd.read_csv('events.csv', parse_dates=['timestamp'])
# rename columns for PM4Py
df = df.rename(columns={'case_id':'case:concept:name', 'activity':'concept:name', 'timestamp':'time:timestamp'})
log = log_converter.apply(df)
```

2) Discovering process models and visualizing

Use discovery algorithms (e.g., Inductive Miner, Heuristics Miner) to build models. PM4Py supports several miners and visualization backends.

```python
from pm4py.algo.discovery.inductive import factory as inductive_miner
from pm4py.visualization.petrinet import factory as pn_vis

net, im, fm = inductive_miner.apply(log)
gviz = pn_vis.apply(net, im, fm)
pn_vis.view(gviz)
```

3) Feature engineering for ML

Process mining yields rich features per case: total throughput time, activity counts, time between specific activities, resource load, and frequency of rare paths. These make strong predictors when combined with static attributes from the business data.

Practical features:
- case_duration = max(timestamp) - min(timestamp)
- activity_counts: how many times each activity appears per case
- waiting_times: mean/median time between consecutive activities
- path_signature: compressed representation of the activity sequence

4) A pragmatic modeling loop

Apply typical data‑science steps: split by case, build features, train/test, and validate with time‑aware splitting to avoid leakage. For production, monitor model drift and re-run process feature extraction as logs evolve.

5) Putting process mining inside a real project

The Medium examples emphasize real‑world concerns: messy timestamps, missing case identifiers, and schema drift. Good practices:
- validate and canonicalize timestamps early
- infer case IDs when absent (grouping heuristics)
- keep a reproducible ETL script for event extraction

6) When to use process features vs raw sequence models

Simple tabular models with hand‑crafted process features are often more interpretable and cheaper to maintain than sequence models. Use sequence models (RNNs/transformers over activities) when history encoding clearly improves predictive performance and the team can maintain the complexity.

Checklist to get started

- [ ] Identify the event sources and the case id column
- [ ] Export a sample CSV with: case_id, activity, timestamp, and any static attributes
- [ ] Run PM4Py discovery on the sample; inspect model and logs for obvious issues
- [ ] Create per‑case features and run exploratory modeling (time‑aware CV)
- [ ] Add monitoring: data schema checks and drift detection

Skill curation and SkillOS: making pipelines live

Google's SkillOS thread (explained in AVB's Paper Breakdown) describes a two-part architecture: a frozen executor that solves tasks by loading reusable "skills" from a SkillRepo, and a trainable Curator that observes executor trajectories and issues structured edits to the SkillRepo (insert/update/delete). The Curator is trained with a group-based curriculum and a composite reward that measures downstream task success, function-call validity, information compression, and content quality.

For process‑mining pipelines the Curator can distill robust ETL and feature‑engineering recipes into SKILL.md files (frontmatter + concise description used for BM25 retrieval, step‑by‑step workflow, worked example, and "when not to use"). Example skills: `extract_event_log`, `feature_engineer_case_features`, `build_delay_model`.

Benefits: repeated runs produce distilled, versioned recipes that accelerate reproducible pipelines and improve executor reliability while keeping instructions modular and auditable.

Operational notes: require human review before promoting automated updates; avoid embedding dataset‑specific constants; maintain test tasks to evaluate curator changes.

References

- Hussam Alhumsi: "Process Mining with Python" — https://hussamalhumsi-21111.medium.com/process-mining-with-python-6ca1d733b3e6
- İ. Cem Özçelik: "Solving a Real‑World Data Science Task with Python" — https://medium.com/@i.cemozcelik/solving-a-real-world-data-science-tasks-with-python-c43aa7d654d1
- AVB (SkillOS thread / breakdown): https://x.com/neural_avb/status/2053873358853591435?s=20

Notes: This post paraphrases and synthesizes practical advice from the referenced Medium posts and general process‑mining best practices. For full, article‑level detail, consult the original posts.


Engineering is the application of knowledge to solve problems within constraints. The constraint here is the central fact. Understanding the constraint is understanding the problem. The solution follows from the constraint. This is the engineering method: name the constraint, design within it, verify the design works.

---
title: Veracity
date: 2026-07-12
slug: v-veracity
summary: "Veracity is the trustworthiness of data. A pipeline that runs successfully can produce garbage. The tests can pass and the data can still be wrong. Veracity is the gap between what the data says and what is true. Closing that gap is the hardest problem in data engineering."
tags: data-engineering, veracity, data-quality, trust, testing
series: the-vs-of-data
part: 4
---

> Veracity is not about whether the data is correct. It is about whether you believe it is correct. Trust is the product. Trust is earned through testing, monitoring, and the accumulated evidence of not being wrong when it mattered.

Veracity is the quality and trustworthiness of data. It is the hardest V because it is not purely technical. A pipeline can run successfully and produce garbage. The tests can pass and the data can still be wrong — the tests test what you thought to test, and you didn't think of everything. The consumers of data — analysts, executives, ML models — trust the numbers until they don't. The moment they stop trusting, every number produced by the data platform becomes suspect. Restoring trust is harder than destroying it. Most organizations never fully restore it.

A weather forecasting system illustrates the challenge. Data arrives from ground stations (accurate but sparse), satellites (global coverage but lower resolution), weather balloons (vertical profiles, twice daily), and citizen reports (abundant but unreliable — people report hail when they hear acorns on the roof). The system must fuse data from sources of varying veracity, weighting each by its historical accuracy. The fusion is the easy part. The hard part is maintaining the weights as source quality changes — a ground station's sensor drifts, a satellite's calibration degrades, a citizen reporting network grows. The weights must adapt. The adaptation must be automated. The automation must be verified against ground truth. The ground truth is expensive to obtain. The expense is the cost of veracity.

The data quality testing pyramid — column-level tests (nulls, ranges, allowed values), table-level tests (uniqueness, referential integrity), cross-table tests (reconciliation across sources), business logic tests (known reference values) — is the engineering response to veracity. Tests catch errors before users do. Users catching errors is the worst outcome because it destroys trust. Each error that reaches a user reduces the user's confidence in the platform. The reduction is cumulative. After enough errors, the user stops using the platform. The platform has failed. The failure is not that the data was wrong. The failure is that the wrongness was invisible until it reached the user.

Data lineage — the ability to trace any number in any dashboard back to its source, through every transformation — is the mechanism for debugging veracity failures. When the CFO asks why revenue is down, the data engineer must be able to trace the revenue number backward: this table → that pipeline → this source → that extraction. The trace is the answer. The inability to trace is the failure. Lineage is the data engineer's call stack. Without it, debugging is archaeology. With it, debugging is engineering.

*See: Tom Redman, "Data Driven: Profiting from Your Most Important Business Asset" (Harvard Business Review Press, 2008) — the foundational text on data quality as a management discipline. Barr Moses et al., "Data Quality Management at Scale" (Monte Carlo, 2022) — modern operational practices for data observability. Wenfei Fan and Floris Geerts, "Foundations of Data Quality Management" (Morgan & Claypool, 2012) — a formal treatment of data consistency, currency, and accuracy.*

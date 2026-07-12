---
title: Data quality and the problem of truth
date: 2026-07-12
slug: data-engineering-quality
summary: "Data quality is the hardest problem in data engineering because it is not purely technical. Bad data looks like good data. Errors propagate silently. The people who know the data is wrong are not the people who can fix it. The discipline of data quality is the discipline of making wrongness visible before anyone makes a decision based on it."
tags: data-engineering, data-quality, governance, testing, lineage
series: data-engineering
part: 4
---

Data quality is the problem of ensuring that data is correct, complete, consistent, and timely. It is the hardest problem in data engineering because it is not purely technical. A pipeline can run successfully and produce garbage. The tests can pass and the data can still be wrong — the tests test what you thought to test, and you didn't think of everything. The consumers of the data — the analysts, the dashboard viewers, the ML models — trust the data until they don't. The moment they stop trusting it, every number produced by the data platform becomes suspect. Restoring trust is harder than maintaining it. Most organizations never fully restore it.

## The dimensions of quality

Data quality has multiple dimensions. Each is necessary. None is sufficient.

**Accuracy.** The data reflects reality. The sales amount in the warehouse matches the sales amount in the source system. The match should be exact. Approximate is not good enough — a 1% error on revenue, compounded across months, produces financial statements that don't reconcile. Accuracy is verified by reconciliation: comparing warehouse aggregates to source system aggregates. The reconciliation should be automated. It rarely is.

**Completeness.** All the expected data is present. No missing rows, no missing columns, no missing values where values should exist. Completeness is verified by volume checks: does today's row count fall within the historical range? Completeness failures are the most common data quality issue and the most embarrassing — "the dashboard shows zero revenue because the pipeline didn't run."

**Consistency.** The same data has the same meaning across tables. Customer ID 123 in the orders table refers to the same customer as Customer ID 123 in the customers table. Consistency is verified by referential integrity checks: do all foreign keys resolve to existing primary keys? Consistency failures produce silent errors — queries that return wrong results without complaining.

**Timeliness.** The data is available when it is needed. A dashboard that shows yesterday's data at 9 AM is useful. A dashboard that shows yesterday's data at 4 PM is not — the decisions were already made. Timeliness is verified by freshness checks: did the pipeline run within its SLA window? Timeliness failures are operations failures dressed as data failures.

**Uniqueness.** No duplicate records. The same event should not appear twice in the same table. Uniqueness is verified by primary key checks. Duplication failures are insidious — they inflate counts, double revenue, and are invisible to checks that only verify that values are within expected ranges.

## Why data rots

Data rots because the world changes and the data model doesn't. A source system is upgraded. A column is renamed. A new product category is added. A business rule changes — "we now count subscriptions as revenue when billed, not when collected." The pipeline wasn't updated. The data in the warehouse no longer matches the data in the source or the expectations of the business. The rot is gradual. The rot is invisible until someone compares two numbers that should match and finds they don't.

Data rots because humans make errors and the errors accumulate. An analyst writes a SQL query with a join condition that accidentally drops rows. The query becomes a view. The view becomes a dashboard. Six months later, someone notices the numbers are low. The error was in the original query. The query was never tested. The test would have caught the error. The test didn't exist.

Data rots because documentation drifts from reality. The data dictionary says the `status` column contains 'active', 'inactive', and 'pending'. The application added 'suspended' six months ago. The documentation wasn't updated. The analyst's query filters to `status IN ('active', 'inactive', 'pending')`. Suspended customers are invisible. The invisibility is a data quality failure caused by a documentation failure.

## The testing pyramid for data

Software engineering has the testing pyramid: unit tests at the bottom, integration tests in the middle, end-to-end tests at the top. Data engineering needs its own pyramid.

**Column-level tests.** Does this column contain nulls? Are values within expected ranges? Are all values from the expected set? Column-level tests catch the most common failures: missing data, out-of-range data, unexpected values. They are cheap to write and fast to run. Every column that matters should have at least one test.

**Table-level tests.** Is the primary key unique? Are there more rows than the minimum expected? Fewer rows than the maximum expected? Do foreign keys resolve? Table-level tests catch structural failures: duplicates, missing data, referential integrity violations.

**Cross-table tests.** Does the revenue in the orders table match the revenue in the payments table? Does the customer count in the warehouse match the customer count in the source? Cross-table tests catch reconciliation failures. They are the most valuable tests and the least common because they require understanding the relationships between systems.

**Business logic tests.** Does the total revenue for March 2026 match the known value from the audited financial statements? Business logic tests catch semantic errors — the data is structurally correct but wrong. They require known reference values. The reference values must come from outside the data platform — from the accounting system, from the source application, from manual audit. The independence of the reference value is what makes the test valid.

## Governance

Data governance is the set of policies and processes that ensure data quality at scale. Governance answers: who owns this data? Who can access it? What does it mean? How long is it retained? What are the quality standards?

Governance fails when it is imposed by a central team without the authority to enforce it. The central team writes policies. The domain teams ignore them. The policies are documents. The documents are unread. Governance succeeds when it is embedded in the platform. Access controls are enforced by the warehouse, not by policy. Data classification is enforced by automated scanning, not by manual tagging. Quality standards are enforced by automated testing, not by review checklists. The platform is the enforcement mechanism. The mechanism is the governance.

The modern approach to governance is the data catalog. A catalog — Alation, Atlan, DataHub, Amundsen — indexes all data assets across the organization. It shows lineage — where data came from and where it goes. It shows ownership — who is responsible for this dataset. It shows quality metrics — when was this dataset last validated, what tests does it pass. It shows usage — who queries this dataset, how often. The catalog makes data discoverable. Discoverability is the first step toward quality. You cannot improve what you cannot find.

## The principle

Data quality is not a project. It is a practice. There is no point at which you are "done" with data quality. The sources change. The business changes. The consumers' expectations rise. The practice must be continuous: test, monitor, fix, repeat. The practice must be owned by the people closest to the data. The platform must make the practice easy. The alternative is a data platform that produces numbers nobody trusts. A platform that produces untrusted numbers is a platform that has failed. The failure is not technical. It is organizational. The fix is not a tool. It is a commitment.

---

**References:**
- Barr Moses, "The Data Quality Hierarchy of Needs," Monte Carlo Blog, 2021.
- dbt Labs, "Testing," dbt Documentation.
- DataHub, "Data Catalog," DataHub Documentation.
- Related posts: [No solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs), [Scarcity in Practice](https://blog.hackspree.com/#scarcity-in-practice)

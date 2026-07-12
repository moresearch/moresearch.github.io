---
title: Pipelines, ETL, and the art of moving data
date: 2026-07-12
slug: data-engineering-pipelines
summary: "A data pipeline is a program that moves data from A to B, transforming it along the way. It sounds simple. It is not. Pipelines fail silently, corrupt data, run late, and cost more to maintain than to build. The discipline of pipeline engineering is the discipline of making data movement boring."
tags: data-engineering, etl, elt, pipelines, orchestration
series: data-engineering
part: 3
---

A data pipeline is a program that moves data from somewhere to somewhere else, transforming it along the way. It is the simplest concept in data engineering. It is also where most of the pain lives. Pipelines break. Pipelines run late. Pipelines produce wrong data and nobody notices for three weeks. Pipelines cost ten times more to maintain than to build. The discipline of pipeline engineering is the discipline of making data movement boring. Boring is the highest compliment a pipeline can receive.

## ETL vs. ELT

The traditional pattern is ETL: Extract, Transform, Load. Extract data from source systems. Transform it — clean, enrich, aggregate — using an external processing engine. Load the transformed data into the warehouse. The transformation happens outside the warehouse, in a pipeline tool or custom code.

The modern pattern is ELT: Extract, Load, Transform. Extract data from sources. Load it into the warehouse raw, in its original format. Transform it inside the warehouse using SQL. The transformation happens after loading, using the warehouse's compute.

The shift from ETL to ELT was driven by the cloud data warehouse. Snowflake and BigQuery separate storage from compute. You can run transformations on massive datasets without provisioning infrastructure. The warehouse scales compute elastically. The transformation is just SQL. The SQL is version-controlled, tested, and documented by tools like dbt. The analyst who writes the query can also write the transformation. The engineer who built the pipeline doesn't need to understand the business logic.

ELT has a deeper advantage: the raw data is preserved. If the transformation logic is wrong, you can fix it and re-transform from the raw data. In ETL, if the transformation is wrong and the raw data was discarded, you must re-extract from the source — if the source still has the data. The raw layer is insurance. The insurance costs storage. Storage is cheap. The insurance is worth it.

## The properties of a good pipeline

**Idempotent.** Running the pipeline twice produces the same result as running it once. If a pipeline fails halfway through and you restart it, you should not get duplicate data. Idempotency is achieved by deduplication — checking whether a record already exists before inserting it — or by overwriting the output partition entirely. Overwriting is simpler. Deduplication is harder but necessary when you can't afford to recompute the entire output.

**Retryable.** If the pipeline fails, it can be restarted without manual intervention. The retry should be automatic, with exponential backoff, up to a maximum number of attempts. The failure should be logged, with enough context to debug it. The alert should fire only after all retries are exhausted. Alerting on the first failure generates noise. Alerting after retry exhaustion generates signal.

**Observable.** You can answer: did the pipeline run? Did it succeed? How many rows did it process? How long did it take? Were there any anomalies — zero rows when there should be thousands, ten times the usual row count, nulls in columns that should never be null? Observability requires metrics, logs, and alerts. The metrics must be queryable historically — "is this week's row count unusual compared to the last twelve weeks?" The alert threshold must be tuned to avoid false positives while catching genuine anomalies.

**Testable.** You can verify that the pipeline produces correct output. Schema tests: does the output table have the expected columns with the expected types? Data tests: are primary keys unique? Are foreign keys present in the referenced table? Are values within expected ranges? Business logic tests: does the revenue column sum to the expected value given known inputs? Tests catch errors before users do. Users catching errors is the worst outcome.

**Lineage-tracked.** You can trace any row in any dashboard back to its source. The trace requires metadata: which pipeline produced this table, from which source tables, using which transformation logic, running at which time. Lineage is essential for debugging — "this number looks wrong, where did it come from?" — and for impact analysis — "if we change this source schema, what downstream tables are affected?" Lineage is the data engineer's call stack.

## The patterns

**Full refresh.** Drop the output table. Rebuild it from scratch. Simple. Correct. Expensive for large tables. The full refresh is the default for small datasets and the fallback when incremental logic fails.

**Incremental.** Process only the data that changed since the last run. Requires a reliable way to identify changed records — a timestamp column, a change log, a CDC feed. More complex than full refresh. More efficient. The incremental pattern is necessary for large tables but introduces the risk of drift: over time, incremental updates accumulate errors that a full refresh would eliminate. Periodic full refreshes — weekly, monthly — reset the drift.

**CDC (Change Data Capture).** Read the database's transaction log directly. Capture every insert, update, and delete as it happens. Debezium for PostgreSQL, MySQL, MongoDB. The CDC feed is a stream of events. The pipeline consumes the stream and applies changes to the warehouse. CDC is the gold standard for freshness — the warehouse is seconds behind the source. It is also the most operationally complex. The CDC connector can fail. The transaction log can be purged before the connector reads it. The schema can change and break the connector's parsing. CDC is powerful. CDC is not free.

**Lambda architecture.** Maintain two parallel pipelines: a batch layer that processes all historical data and produces accurate but delayed results, and a speed layer that processes recent data in real time with approximate results. The serving layer merges both. The lambda architecture was popular in the Hadoop era. It has been largely replaced by stream processing systems (Kafka Streams, Flink) that can handle both real-time and historical processing in a single framework.

**Kappa architecture.** Process everything as a stream. Historical data is replayed from the stream's retention log. New data arrives in real time. The same code handles both. Simpler than lambda. Requires a stream platform with long-term retention — Kafka with tiered storage, or a cloud-native equivalent. Kappa is the modern default for organizations that have invested in streaming infrastructure.

## The economics

Pipelines cost more to maintain than to build. The initial build is a week. The maintenance is years. Every source system change — a renamed column, a new data type, a deprecated API — requires a pipeline update. Every business logic change requires a transformation update. Every scale increase — more data, more users, more dashboards — requires performance optimization. The maintenance cost is proportional to the number of pipelines and the rate of change of their dependencies.

The economics favor fewer pipelines, simpler pipelines, and pipelines owned by the people who understand the data. The centralized data team that builds pipelines for every department becomes a bottleneck. The data mesh model — each domain owns its pipelines — distributes the maintenance cost to the teams that benefit from the data. The distribution is the economics of the data mesh. The same economics that favor microservices over monoliths — independent deployability, domain ownership, reduced coordination cost — favor the data mesh over the centralized warehouse. The principles are identical. The domain is different.

---

**References:**
- Maxime Beauchemin, "Functional Data Engineering — a modern paradigm for batch data processing," 2018.
- Jay Kreps, "Questioning the Lambda Architecture," 2014.
- dbt Labs, "What is dbt?" dbt Documentation.
- Related posts: [The Unix philosophy](https://blog.hackspree.com/#unix-philosophy), [No solutions, only trade-offs](https://blog.hackspree.com/#no-solutions-only-tradeoffs)

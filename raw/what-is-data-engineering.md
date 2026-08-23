---
title: What data engineering actually is
date: 2026-07-12
slug: what-is-data-engineering
summary: "Data engineering is the discipline of building systems that collect, store, transform, and serve data. It is not glamorous. It is not AI. It is the plumbing that makes AI possible. Without it, models have nothing to train on and dashboards show yesterday's numbers."
tags: data-engineering, etl, pipelines, infrastructure
series: data-engineering
part: 1
---

Data engineering is the discipline of building systems that collect, store, transform, and serve data. It sits between the systems that produce data — applications, sensors, user interactions — and the systems that consume data — dashboards, machine learning models, analysts. It is the plumbing. Plumbing is unglamorous. Plumbing is essential. Without plumbing, the house is uninhabitable.

The field emerged from a specific historical sequence. In the 1980s and 1990s, organizations built data warehouses — centralized repositories that aggregated data from operational systems for reporting and analysis. Ralph Kimball and Bill Inmon developed competing methodologies. Kimball advocated dimensional modeling — star schemas with fact tables and dimension tables, optimized for query performance. Inmon advocated the Corporate Information Factory — a normalized enterprise data warehouse feeding departmental data marts. The debate was religious. Both approaches worked. Both assumed that data was structured, that schemas were stable, and that the warehouse team could enforce standards.

The 2000s broke these assumptions. The volume of data exploded. The variety of data exploded — logs, JSON, sensor readings, social media feeds, clickstreams. The velocity of data increased — real-time streams replaced nightly batch loads. The old warehouse architectures couldn't keep up. Hadoop emerged. MapReduce provided a programming model for distributed data processing. HDFS provided a distributed filesystem. The ecosystem was complex, Java-heavy, and operated by a priesthood of engineers who understood the arcana of YARN configuration and NameNode failover. It worked. It was unpleasant.

The 2010s simplified the stack. Apache Spark replaced MapReduce with an in-memory processing engine that was faster and easier to program. Cloud data warehouses — Snowflake, BigQuery, Redshift — made the warehouse model viable again at cloud scale. The ELT pattern (Extract, Load, Transform) replaced ETL (Extract, Transform, Load): load raw data into the warehouse first, transform it later using the warehouse's own compute. The shift moved transformation logic from external pipelines into SQL, where analysts could contribute.

The 2020s are the era of the modern data stack. Fivetran and Airbyte handle extraction. dbt handles transformation — SQL-based, version-controlled, tested. Snowflake, BigQuery, and Databricks handle storage and query. Airflow and Prefect handle orchestration. The tools are better. The principles are the same: get data from where it is to where it needs to be, reliably, at the right time, in the right shape.

## The core problems

Data engineering has five core problems. Every tool, every architecture, every methodology is a response to one or more of them.

**Ingestion.** Getting data into the system. From databases (CDC — change data capture), from APIs (REST, GraphQL), from files (CSV, JSON, Parquet), from streams (Kafka, Kinesis). The data arrives in different formats, at different cadences, with different reliability characteristics. The ingestion layer must handle all of them without losing data, duplicating data, or falling behind.

**Storage.** Keeping data somewhere it can be accessed. The choice of storage format — row-oriented vs. columnar, compressed vs. uncompressed, partitioned vs. monolithic — determines query performance, storage cost, and the ability to evolve schemas over time. The choice of storage engine — data warehouse vs. data lake vs. lakehouse — determines who can query the data and with what tools.

**Transformation.** Turning raw data into useful data. Cleaning — removing duplicates, fixing nulls, standardizing formats. Enriching — joining with reference data, computing derived fields, applying business logic. Aggregating — rolling up to daily, weekly, monthly levels for dashboards. The transformation layer is where most of the engineering effort goes. It is also where most of the bugs are.

**Orchestration.** Making everything run at the right time, in the right order, with the right dependencies. Pipeline A must finish before Pipeline B starts. Pipeline B must not run if Pipeline A produced bad data. The orchestration layer manages schedules, dependencies, retries, alerts, and backfills. It is the conductor. When the conductor fails, the orchestra plays anyway — out of sync, producing cacophony.

**Serving.** Getting data to the people who need it. Dashboards (Looker, Tableau, Metabase). Ad-hoc queries (SQL editors, notebooks). Machine learning feature stores. Reverse ETL — sending transformed data back to operational systems (CRM, email, advertising). The serving layer determines whether the data is actually used. The best pipeline in the world is worthless if nobody looks at its output.

## How it differs from software engineering

Data engineering is software engineering with different constraints. Software engineering optimizes for correctness, latency, and throughput of application logic. Data engineering optimizes for correctness, latency, and throughput of data movement and transformation. The difference is the nature of the bugs.

A software bug produces a wrong output for a specific input. A data engineering bug produces a wrong output for millions of records, discovered three weeks later when the CFO asks why revenue is down. The blast radius is larger. The debugging is harder — you must trace the error backward through multiple transformation steps, each of which may have run days or weeks ago. The fix requires not just correcting the code but reprocessing the affected data, which may take hours or days. The operational complexity of data engineering is higher than application engineering because the state — the data — is larger, more persistent, and harder to repair.

Data engineering also has a different failure mode: silent corruption. A software system crashes visibly — errors, exceptions, downtime. A data pipeline can produce wrong numbers silently, for weeks, before anyone notices. The pipeline didn't fail. It ran successfully. The data is wrong. The wrongness is invisible until someone looks at the numbers and says "that doesn't seem right." The delay between corruption and detection is the most dangerous property of data systems.

## The data engineer's mindset

The data engineer thinks in terms of data flows, not control flows. The question is not "what does this function return?" but "where does this data come from, what happens to it along the way, and who consumes the output?" The data engineer traces lineage forward and backward through the system. Forward: if this source data changes, what downstream tables are affected? Backward: if this dashboard number is wrong, which pipeline produced it, from which source, using which transformation logic? The ability to trace lineage is the data engineer's superpower. The inability to trace lineage is why data projects fail.

The data engineer is paranoid about state. Every pipeline should be idempotent — running it twice produces the same result as running it once. Every pipeline should be retryable — if it fails, it can be restarted without corrupting the output. Every pipeline should be testable — a small sample of input data should produce a predictable output. Every pipeline should be monitored — if it produces zero rows, or ten times the usual number of rows, or rows with nulls where there should be values, someone should be alerted. The paranoia is not anxiety. It is engineering.

---

**References:**
- Ralph Kimball, *The Data Warehouse Toolkit*, Wiley, 1996.
- Bill Inmon, *Building the Data Warehouse*, Wiley, 1992.
- Maxime Beauchemin, "The Rise of the Data Engineer," 2017.
- Related posts: [The Unix philosophy](https://blog.hackspree.com/#unix-philosophy), [Engineering is economics](https://blog.hackspree.com/#engineering-is-economics)

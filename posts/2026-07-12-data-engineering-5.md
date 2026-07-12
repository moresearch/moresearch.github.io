---
title: Storage engines and the physics of query
date: 2026-07-12
slug: data-engineering-storage-engines
summary: "The choice of storage format determines everything above it: query speed, storage cost, schema flexibility, concurrent access. Row-oriented for transactions. Columnar for analytics. Parquet and Iceberg for the lakehouse. DuckDB for the laptop. The physics of storage is the physics of data."
tags: data-engineering, storage, parquet, iceberg, columnar, duckdb
series: data-engineering
part: 5
---

Data is stored as bytes on disk. How those bytes are organized determines how fast they can be read, how much space they occupy, and what kinds of queries are possible. The physics of storage is the foundation of data engineering. Everything above — the pipelines, the transformations, the dashboards — depends on choices made at the storage layer. The choices are irreversible enough that getting them wrong is expensive.

## Row vs. column

Data can be stored row-by-row or column-by-column. In row-oriented storage, all the fields of a single record are stored together. Record 1: (name, age, city, salary). Record 2: (name, age, city, salary). Row storage is optimal for transactional workloads — inserting, updating, deleting individual records. The database reads or writes the entire row at once. PostgreSQL, MySQL, and most OLTP databases use row storage.

In column-oriented storage, all the values of a single column are stored together. Column name: (Alice, Bob, Carol, Dave). Column age: (30, 25, 35, 28). Column storage is optimal for analytical workloads — aggregating, filtering, grouping across many rows. A query that computes the average salary reads only the salary column. It skips name, age, and city. The reduction in I/O is the performance gain. Columnar databases — Snowflake, BigQuery, Redshift, ClickHouse — are designed for analytics.

The columnar advantage comes from three properties. First, columnar compression is more effective than row compression because values within a column are similar — all integers, all strings of similar length, all dates. Run-length encoding, dictionary encoding, and delta encoding compress columns efficiently. Second, columnar storage enables vectorized execution — the query engine operates on batches of values rather than individual rows, using SIMD instructions for parallelism within a single core. Third, columnar storage enables late materialization — the query engine can filter on one column before reading other columns, reducing the total I/O.

The trade-off: columnar storage is bad for point queries. Finding a single row by ID requires reading every column independently and reassembling the row. Row storage does this in a single read. The right storage format depends on the workload. Most data engineering workloads are analytical. Columnar is the default.

## File formats

**CSV.** Comma-separated values. Human-readable. Universally supported. No schema enforcement. No type information. No compression. No nested data. CSV is the lowest common denominator. It is the format you use when you need to exchange data with a system that speaks nothing else. It is not a format for production storage.

**JSON.** Nested, semi-structured. Self-describing — each record contains its field names. Human-readable (for small records). Widely supported. Inefficient — field names are repeated in every record, numbers are stored as text. JSON is the format of APIs and event streams. It is the format you ingest, not the format you query.

**Parquet.** Columnar, compressed, with schema embedded in the file footer. Developed by Twitter and Cloudera in 2013. The dominant format for analytical workloads in the Hadoop and cloud ecosystems. Parquet stores data in row groups — chunks of rows stored column-by-column within the chunk. The row group size balances read parallelism (more groups = more parallelism) with columnar efficiency (larger groups = better compression). Parquet supports predicate pushdown — the query engine reads the file footer to determine which row groups contain data matching the query's filters, and skips the rest. The skipping is the performance.

**Avro.** Row-oriented, with schema stored in the file header. Developed by Doug Cutting for Hadoop. Used primarily for streaming data and as a serialization format for Kafka messages. Avro supports schema evolution — adding, removing, or modifying fields over time while maintaining backward compatibility. The schema evolution is the reason Avro is used for streams. Streams are long-lived. Schemas change. Avro handles the change.

**ORC.** Optimized Row Columnar. Developed by Hortonworks for Hive. Similar to Parquet in design. Better compression. Less ecosystem support. Parquet won the format war. ORC is still used in the Hive ecosystem but is not the default for new projects.

## Table formats

File formats solve the problem of how to store a single file. Table formats solve the problem of how to manage a collection of files as a single table — partitioning, schema evolution, time travel, concurrent writes.

**Apache Iceberg.** Developed at Netflix. The dominant table format in 2024-2026. Iceberg tracks table metadata — the list of files that comprise the table, the schema, the partitions, the statistics — in a manifest. Queries read the manifest to determine which files to scan. Inserts, updates, and deletes produce new files. Old files are garbage-collected. The table is a logical abstraction over a collection of physical files. Iceberg supports hidden partitioning — the partition scheme is stored in metadata, not in the file path. Changing the partition scheme does not require rewriting the data. Schema evolution is additive — adding a column is a metadata change. Time travel — querying the table as of a past timestamp — is a manifest lookup. Iceberg works with multiple query engines — Spark, Trino, Flink, Snowflake, BigQuery, DuckDB. The engine independence is the architectural win.

**Delta Lake.** Developed at Databricks. Similar to Iceberg in concept. Deeper integration with Spark. Uses a transaction log rather than manifests. Supports ACID transactions, schema enforcement, and versioning. The competition between Iceberg and Delta Lake is the format war of the 2020s. Both are winning. The real winner is the user, who gets a standard abstraction over object storage.

**Apache Hudi.** Developed at Uber. Designed for streaming and incremental processing. Supports record-level upserts and deletes. More complex than Iceberg or Delta Lake. Used when the workload requires frequent updates to individual records — a streaming pipeline that must correct errors in previously written data.

## Query engines

**Data warehouses (Snowflake, BigQuery, Redshift).** Fully managed, SQL-based, designed for analytical queries on structured data. Separate storage from compute. Scale elastically. Handle concurrency, security, and administration. The data warehouse is the default for organizations that have data analysts who write SQL and need a managed platform.

**Query engines (Trino, Presto, Starburst).** Federated SQL engines that query data in place — Parquet files in S3, tables in PostgreSQL, streams in Kafka. No data loading. No ETL. The query engine pushes computation to the data. Trino is the engine behind many data lake architectures. It is fast. It is complex to operate at scale.

**Embedded engines (DuckDB).** An in-process analytical database. No server. No configuration. Runs on a laptop. Reads Parquet, CSV, JSON directly. DuckDB is the SQLite of analytics. It is the engine you use when you need to query a Parquet file on your laptop and don't want to set up a warehouse. It is also the engine that is eating the low end of the analytical market — queries that would have required a warehouse five years ago now run in DuckDB on a MacBook.

**Stream processors (Flink, Kafka Streams, RisingWave).** Process data as it arrives, producing results continuously. Designed for real-time dashboards, alerting, and event-driven applications. Maintain state in memory and on disk. Provide exactly-once semantics. The stream processor is the complement to the warehouse — the warehouse handles historical analysis, the stream processor handles real-time decisions.

## The physics

The physics of storage is the physics of data engineering. Data at rest must be organized so that data in motion can be processed efficiently. The organization is the schema. The format is the encoding. The table format is the abstraction. The query engine is the processor. Each layer constrains the layer above. The constraints are not limitations. They are the design. The design determines what is possible. What is possible determines what is built.

---

**References:**
- Apache Iceberg, "Specification," Iceberg Documentation.
- DuckDB, "Why DuckDB," DuckDB Documentation.
- Netflix Tech Blog, "Apache Iceberg: An Architectural Look Under the Covers," 2020.
- Related posts: [The Unix philosophy](https://blog.hackspree.com/#unix-philosophy), [libp2p](https://blog.hackspree.com/#libp2p)

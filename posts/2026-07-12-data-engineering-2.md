---
title: Data modeling is the hard part
date: 2026-07-12
slug: data-engineering-modeling
summary: "Data modeling is the art of deciding how to structure data so it can be queried efficiently and understood by humans. The models change. The principles don't. From Codd's relational model to Kimball's star schemas to the modern wide table, the problem is always the same: how do you represent reality in tables?"
tags: data-engineering, data-modeling, kimball, star-schema, normalization
series: data-engineering
part: 2
---

Data modeling is the hardest part of data engineering because it is the part that requires judgment. You can learn a tool in a week. You can learn a pipeline pattern in a day. Data modeling takes years because the feedback loop is slow — you design a model, people query it for months, and only then do you discover what you got wrong. The wrongness is expensive to fix because downstream dashboards, ML models, and business processes have been built on the original model. Changing the model breaks them. The breakage is the cost of the original design error.

## The relational model

Edgar Codd published "A Relational Model of Data for Large Shared Data Banks" in 1970. The paper introduced the idea that data should be stored in relations — tables — with well-defined operations for querying and manipulating them. The relational model separated the logical structure of data from its physical storage. Before Codd, databases were navigational — you followed pointers from record to record. The query path was baked into the storage structure. Codd's insight was that queries should be declarative: you specify what you want, not how to get it. SQL is the realization of that insight.

The relational model introduced normalization — the process of organizing data to minimize redundancy. First normal form: no repeating groups. Second normal form: no partial dependencies on a composite key. Third normal form: no transitive dependencies. The normal forms are a hierarchy of increasingly strict constraints on table design. A database in third normal form has minimal duplication. Every fact is stored exactly once. Changes to a fact require updating a single row.

Normalization is elegant. It is also slow for analytical queries. A normalized schema requires joins to reconstruct the original business entities. Joins are expensive on large tables. The tension between normalization (write efficiency, data integrity) and denormalization (read efficiency, query simplicity) is the central tension of data modeling.

## Dimensional modeling

Ralph Kimball resolved the tension by designing for queries, not writes. In a dimensional model, data is organized into fact tables and dimension tables. Fact tables contain measurements — sales amounts, page views, sensor readings. Each row is an event. Dimension tables contain descriptions — customer names, product categories, date attributes. Each row describes an entity. The fact table references dimension tables through foreign keys.

The star schema is the simplest dimensional model. A central fact table surrounded by dimension tables, like points of a star. A sales fact table has foreign keys to date, customer, product, and store dimensions. A query joins the fact to any subset of dimensions. The joins are simple — each dimension is one hop from the fact. The simplicity makes the schema understandable to business users who write SQL.

The star schema's power is that it separates the what (measurements) from the who, what, when, and where (dimensions). You can ask any question that starts with "how many X by Y" — how many sales by product by month? How many page views by country by device? The answer is a join between the fact and the relevant dimensions. The model constrains the questions you can ask to the questions the business needs answered. The constraint is the design.

Kimball's methodology includes slowly changing dimensions (SCDs) — how to handle changes to dimension attributes over time. Type 1: overwrite the old value. Type 2: add a new row with the new value, preserving history. Type 3: add a new column for the new value. Each type trades query complexity for historical accuracy. The choice is a business decision, not a technical one. The business must decide whether historical accuracy matters enough to justify the complexity.

## The modern wide table

The modern data stack has shifted toward wide, denormalized tables. The warehouse engines — Snowflake, BigQuery — are fast enough that joins are less expensive than they were. The transformation layer — dbt — makes it easy to build and maintain derived tables. The result is the One Big Table (OBT) pattern: a single table with hundreds of columns, pre-joined, pre-aggregated, ready for the dashboard to query with a simple `SELECT *`.

The wide table is a response to the reality that most business users cannot write joins. They can write `SELECT * FROM orders_wide WHERE date > '2026-01-01'`. The wide table makes the data accessible. The cost is storage (denormalized data is larger), maintenance (the wide table must be rebuilt when source schemas change), and lineage opacity (it's harder to trace where each column came from). The trade-off is economic: the cost of storage and compute is lower than the cost of analyst time spent writing joins incorrectly. The economics favor the wide table. The wide table is the default.

## Data mesh and domain ownership

Zhamak Dehghani's data mesh (2019) challenges the centralized data warehouse model. The argument: data should be owned by the domains that produce it, not by a central data team. Each domain publishes data products — curated datasets with defined schemas, quality guarantees, and SLAs. The central team provides the platform — infrastructure, tooling, governance standards. The domains own the data.

The data mesh is a response to the scaling problems of centralized data teams. As the number of data sources grows, the central team becomes a bottleneck. Every new data source requires the central team to understand the domain, model the data, build the pipeline, and maintain it. The domain expert who understands the data is not the person building the pipeline. The knowledge gap produces errors. The bottleneck produces delays. The data mesh solves both by moving the pipeline ownership to the domain.

The trade-off: domain teams must now hire data engineering skills. The central team must build a platform that makes it easy for domain teams to publish data products. The governance must be federated — standards enforced by the platform, content owned by the domains. The data mesh is an organizational pattern, not a technology. The technology is the same as the centralized model. The organization is different. The difference is the innovation.

## The principle

The specific model — star schema, wide table, data mesh — matters less than the principle: data must be structured so that the people who need it can find it, understand it, and trust it. The structure that achieves this for a five-person startup is different from the structure that achieves it for a five-thousand-person enterprise. The principle is the same. The implementation varies. The variation is the work.

---

**References:**
- Edgar Codd, "A Relational Model of Data for Large Shared Data Banks," *Communications of the ACM*, 1970.
- Ralph Kimball, *The Data Warehouse Toolkit*, Wiley, 1996.
- Zhamak Dehghani, "How to Move Beyond a Monolithic Data Lake to a Distributed Data Mesh," 2019.
- Related posts: [Parnas's Information Hiding](https://blog.hackspree.com/#parnas-information-hiding), [Scarcity and Software Economics](https://blog.hackspree.com/#scarcity-and-software-economics)

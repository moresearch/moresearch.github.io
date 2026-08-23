---
title: On Data Variety: the cost of integration
date: 2026-07-12
slug: v-variety
summary: "Variety is the diversity of data formats, structures, and sources. Every new source requires a pipeline. Every pipeline requires maintenance. The cost of variety is the cost of integration, and integration is the unglamorous work that makes data useful."
tags: data-engineering, variety, integration, schema, etl
series: the-vs-of-data
part: 3
---

> Variety is not about how many formats you have. It is about how many pipelines you must maintain. Each format is a promise. Each pipeline is the cost of keeping that promise. The cost compounds.

Variety is the diversity of data formats, structures, and sources. It is the dimension that makes data integration hard. Data arrives as relational tables, JSON documents, XML messages, CSV files, Parquet files, Avro records, Protobuf messages, images, videos, log files, and free-text documents. Each format has its own schema, its own semantics, its own quirks. Each source has its own update cadence, its own reliability characteristics, its own failure modes. Integrating them requires understanding each one. Understanding each one requires time. Time is the scarcest resource in data engineering.

A hospital illustrates the constraint. Patient data spans electronic health records (structured, HL7 format), lab systems (semi-structured, ASTM), imaging systems (unstructured, DICOM), patient surveys (free text), and wearable devices (JSON streams). A single patient's record crosses five systems, four formats, three data models. Integrating them requires mapping each source to a common model, resolving identifier conflicts (is patient 123 in the EHR the same person as patient 456 in the lab system?), and handling different update cadences. The hospital's data engineering team spends 80% of its time on integration. The 80% is the cost of variety.

Variety is the reason data engineering teams grow faster than the data they manage. Each new data source adds a pipeline. Each pipeline adds a maintenance burden: the source schema changes, the pipeline breaks, the engineer fixes it. The breakage is not a one-time event. It is continuous. Every source system upgrade, every new column, every deprecated field produces a pipeline failure. The failures accumulate. The maintenance burden grows. The team grows to handle the burden. The growth is the cost of variety.

The modern response to variety is the ELT pattern with a schema-on-read approach. Extract the data in its native format. Load it into the data lake or warehouse without transformation. Apply schema at query time. The raw data is preserved. The transformation logic is version-controlled SQL. The approach decouples ingestion (fast, reliable, format-agnostic) from transformation (flexible, query-time, iterative). The decoupling is the architectural insight: you cannot predict how the data will be used, so preserve it in its original form and let consumers apply their own schemas.

*See: Serge Abiteboul et al., "Data on the Web: From Relations to Semistructured Data and XML" (Morgan Kaufmann, 1999) — the classic text on the shift from structured to semi-structured data. Joe Reis and Matt Housley, "Fundamentals of Data Engineering" (O'Reilly, 2022), Chapter 4, on the modern integration landscape. James Serra, "Deciphering Data Architectures" (O'Reilly, 2024) — on choosing between data warehouse, data lake, lakehouse, and data mesh for variety-heavy environments.*


*This post is part of a series on [The Many Vs of Data](https://blog.hackspree.com/#many-vs-of-data), originating from Doug Laney's 2001 Gartner note. Each V names a dimension of why data is hard.*

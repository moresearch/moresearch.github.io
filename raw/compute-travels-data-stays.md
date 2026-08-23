---
title: Compute Travels. Data Stays.
date: 2026-07-20
slug: compute-travels-data-stays
summary: Bacalhau inverts the cloud model: instead of moving petabytes to a central cluster, it sends compute to where data lives. That inversion is the foundation of data sovereignty — and the architecture the agent era needs.
tags: data-sovereignty, bacalhau, distributed-compute, agents, edge, privacy
---

The cloud won. For two decades, the answer to "where should compute happen?" was "ship the data to us." Centralize. Aggregate. Process. The result: petabytes of data flowing into a handful of regions, owned by a handful of companies, governed by laws that don't match the topology.

The cloud won, but the cloud model is breaking. Not technically — the hyperscalers work fine. But legally, economically, and architecturally, the assumption that data should move to compute is no longer true.

> GDPR requires data residency. The EU AI Act layers new constraints. Countries from Brazil to India to Indonesia are writing sovereignty into law. The pipe dream of "store everything in us-east-1" is dead. The cloud didn't account for borders.

This is where Bacalhau enters. Not as a replacement for the cloud. As an inversion of its founding assumption.

## What Bacalhau does

Bacalhau is an open-source distributed compute orchestrator. The idea is simple: instead of moving data across the network to a central compute cluster, you send a small job description — a Docker container, a Wasm binary, a shell script — to wherever the data already lives. The job runs next to the data. Only the results come back.

![Bacalhau logo — the cod fish, Portuguese "bacalhau," a metaphor for preservation without centralization](/images/bacalhau-logo.svg)

> Compute travels. Data stays. That's the whole architecture.

It's released under Apache 2.0, built by Expanso, a company founded in 2023 by David Aronchick — previously co-founder of Kubeflow at Google, head of open-source ML strategy at Microsoft Azure. The project won the Data Breakthrough Award in 2024, raised $7.5M in seed funding led by General Catalyst, and landed a strategic investment from Samsung Next. It's available on the Google Cloud Marketplace. It's a single Go binary — no cluster to bootstrap, no control plane to subscribe to.

The architecture is an orchestrator-compute model. You label nodes by region (`region=eu`, `region=us`). You submit a job with constraints. The orchestrator routes work to nodes near the data. Jobs are parallel by default — split into partitions that run independently, with isolated failure handling. If a partition fails, it retries. If the network between the orchestrator and a compute node drops, the node keeps working.

## The name matters

Bacalhau is Portuguese for dried salted cod. Before refrigeration, cod was preserved by salting and drying — it could travel long distances without spoiling, sustaining entire maritime economies for centuries. The metaphor is deliberate: preserve the data where it originates. Compute can travel. Data doesn't need to.

This matters because data is heavy and data is governed. Moving a petabyte of logs to a central warehouse is expensive. Moving genomic data across national borders is illegal in an increasing number of jurisdictions. Moving medical records, financial transactions, or personally identifiable information into a third-party cloud triggers compliance obligations that most teams underestimate until an auditor shows up.

Bacalhau's answer: don't move it. Send the analysis to the data.

## How it actually works

A Bacalhau deployment runs a single binary in different modes: orchestrator nodes manage job lifecycles, compute nodes execute workloads. Both can be the same machine. At the edge, a compute node might be a Raspberry Pi on a factory floor. In the cloud, it might be a VM in a specific AWS region.

Jobs are submitted declaratively or via CLI. They target data in S3, IPFS, HTTP endpoints, or local storage. The orchestrator schedules work based on node labels — `region=eu`, `gpu=true`, `tier=production`. The workload runs inside a Docker container or a Wasm sandbox. Output lands wherever you configure: local disk, S3 bucket, the next pipeline stage.

A partitioned job splits across N nodes. Each partition gets an index, a count, and its own slice of the data. Failures are per-partition, not per-job. A node in Frankfurt might process EU customer records while a node in São Paulo processes Brazilian records and a node in Mumbai handles Indian data — all from the same job submission, with constraints that ensure compliance.

## The Genomic Data Proof Point

A 2025 academic study tested Bacalhau with IPFS Cluster and AES-256 encryption for decentralized genomic computation. The distributed architecture achieved **100% job completion under network chaos** — nodes disconnecting, links dropping, partitions reforming. The centralized baseline fell apart under the same conditions.

Under ideal network conditions, the distributed setup added about 30% overhead (49 seconds vs. 37 seconds). That 12-second difference is the price of sovereignty. In the centralized case, you also move the data first — a cost the study's baseline conveniently excluded.

> The paper's conclusion: a proven model for privacy-critical decentralized science collaborations, prioritizing data sovereignty and high availability over raw throughput. Twelve seconds to keep genomic data inside hospital walls.

## Why the agent era needs this

Software engineering agents generate code. That's the story everyone tells. The less-told story is where they run.

An agent debugging a production issue needs access to logs. Those logs live in a specific region, governed by specific laws. An agent analyzing customer behavior needs to read data that cannot legally leave the country. An agent optimizing a factory floor needs to process sensor data at millisecond latency — waiting for a round-trip to the cloud is not an option.

The default architecture for agent platforms today is: ship everything to a central LLM provider. Your code, your logs, your database schema, your customer PII — all of it crosses the wire to a model endpoint in a jurisdiction you didn't choose.

Bacalhau suggests a different architecture: ship the agent to the data. Run the model where the data lives. The agent is a job. The data is stationary. The compliance boundary is the node label, not a legal review of every prompt.

## The broader lesson

Data sovereignty sounds like a legal problem. It becomes an architectural problem the moment you try to build a real system. If your architecture requires data to move to a central location before anything useful can happen, you have already lost the sovereignty argument — all that remains is how many exceptions you'll need and how much the fines will cost.

The alternative is compute over data. It's not a new idea. MapReduce did it. Edge computing does it. Bacalhau makes it general — any workload, any data source, any execution engine, one binary, open source.

> The cloud taught us to centralize. The law is teaching us to distribute. Bacalhau is infrastructure for that transition.

It's not that central compute will disappear. It's that centralization stops being the default. When data must stay where it is, the architecture follows. Compute travels. Data stays. That inversion is the foundation of sovereignty — and the infrastructure the next decade of software will be built on.

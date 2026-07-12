---
title: Velocity
date: 2026-07-12
slug: v-velocity
summary: "Velocity is the speed at which data arrives and the speed at which it must be acted upon. Batch is forgiving. Streaming is not. The choice between them is the most consequential architectural decision in data engineering."
tags: data-engineering, velocity, streaming, batch, real-time
series: the-vs-of-data
part: 2
---

> Velocity does not ask "how fast can you process?" It asks "how fast must you decide?" The speed of data is irrelevant if the decision can wait. The speed matters when the decision cannot.

Velocity is the dimension that separates batch from streaming. Batch processing: data arrives in files, accumulates for hours or days, and is processed in a single job that reads all available data and produces output. Streaming processing: data arrives as messages, each message is processed within milliseconds of arrival, and the output updates continuously. The architectures are different. The failure modes are different. The organizational capabilities required are different. Choosing between them is the most consequential architectural decision in data engineering.

A fraud detection system at a payment processor illustrates the velocity constraint. Each transaction generates an event. The event must be evaluated within 100 milliseconds — is this transaction fraudulent? The decision requires consulting historical patterns (has this card been used in this location before?), real-time aggregates (how many transactions has this card made in the past hour?), and ML model inference (what is the fraud score?). The velocity constraint forces the architecture: historical data pre-computed and served from an in-memory cache, real-time aggregates maintained by a stream processor, model served with sub-millisecond latency. A batch system that evaluated fraud once per hour would approve fraudulent transactions for up to 59 minutes before detection. The 59 minutes is the cost of batch.

Streaming is harder than batch. Batch systems are forgiving — if a job fails, you restart it from the last checkpoint. The input data is immutable. The output is overwritten. The failure is a delay. Streaming systems are unforgiving — if a message is missed, the aggregate is wrong. If the stream processor crashes, the in-memory state is lost. If two messages arrive out of order, the windowed computation is incorrect. Streaming requires exactly-once semantics (each message processed exactly once, even across failures), watermark handling (how to handle late-arriving data in windowed computations), and state management (how to recover in-memory state after a crash). The complexity is the price of low latency. The price is worth paying when the decision must be made now.

The modern synthesis is the Lambda architecture and its successor, Kappa. Lambda: maintain two parallel pipelines — a batch layer for accurate but delayed results, a speed layer for approximate but immediate results, a serving layer that merges them. Kappa: process everything as a stream, replaying historical data from the stream's retention log when reprocessing is needed. Kappa won because it simplifies the operational burden — one code path, not two. The simplification is the engineering insight: the stream is the source of truth. The batch is a special case of the stream — a stream with a very long window.

*See: Tyler Akidau et al., "The Dataflow Model" (VLDB, 2015) — the paper that unified batch and streaming under event-time windowing, now the basis for Apache Beam. Jay Kreps, "Questioning the Lambda Architecture" (O'Reilly, 2014) — why maintaining two code paths is a maintenance disaster. Martin Kleppmann, "Designing Data-Intensive Applications" (O'Reilly, 2017), Chapter 11, on stream processing and the Kappa architecture.*

---
title: Go for Disaggregated Serving
date: 2026-02-22
slug: go-is-a-natural-control-plane-for-disaggregated-serving
summary: Disaggregated prefill and decode pipelines need schedulers, backpressure, and observability more than they need another complex runtime.
tags: golang, inference, llm-serving
---

One of the clearest recent serving ideas is in [DistServe: Disaggregating Prefill and Decoding for Goodput-optimized Large Language Model Serving](https://arxiv.org/abs/2401.09670). The paper shows why prefill and decode interfere with each other and why separating them can improve goodput under real latency targets.

That architecture has a very Go-shaped seam in it.

If prefill and decode become distinct pools, somebody has to own:

- request admission,
- routing policy,
- streaming state,
- deadline propagation.

That "somebody" does not need to be the same runtime that executes the kernels. A small Go service is often the right place to implement the policy layer because it stays deployable, observable, and easy to debug when traffic gets weird.

## Separation creates coordination work

Disaggregation is attractive because it stops different phases of inference from fighting each other quite so directly. But separating them does not remove complexity. It relocates complexity into scheduling, buffering, and operational policy.

That is exactly the kind of work that benefits from a clean control plane.

Once there are distinct pools, the system needs a component that can make understandable decisions when load shifts. It has to answer practical questions: which requests get admitted, what deadlines matter most, and how streaming state is preserved without turning every incident into a forensic exercise.

## Why Go fits the seam

A Go service is not interesting here because it is fashionable. It is useful because the control-plane job rewards plain engineering:

- predictable concurrency,
- simple deployment units,
- straightforward observability,
- code paths operators can still follow during an incident.

The serving runtime can stay specialized around execution. The control layer can stay specialized around policy.

That separation of responsibilities feels healthy to me. DistServe highlights why inference phases deserve different treatment. The systems lesson is that once you accept that split, you should also accept a clear policy layer around it. Go is often a very practical place to put that layer.

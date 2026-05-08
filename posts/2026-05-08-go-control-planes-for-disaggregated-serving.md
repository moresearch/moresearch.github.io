---
title: Go is a natural control plane for disaggregated serving
date: 2026-05-08
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

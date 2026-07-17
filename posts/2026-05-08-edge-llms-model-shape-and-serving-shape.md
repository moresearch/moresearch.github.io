---
title: "Edge LLMs: Model Shape and Serving Shape Are One Decision"
date: 2026-01-29
slug: edge-llms-model-shape-and-serving-shape
summary: On-device inference fails in two separate meetings — the one where a cloud model gets picked for compression, and the one where "the serving layer" is treated as a generic runtime beneath it. Two papers, MobileLLM and Fast On-device LLM Inference with NPUs, are usually read as answers to different questions. They are two halves of the same one — the model's shape and the serving stack's shape are co-designed against the same hardware budget, and teams that decide them separately ship the compromises.
tags: edge-llms, mobile, npu, architecture, serving
---

The promise of on-device inference is easy to say and hard to ship. And when it fails to ship, the failure usually traces back to an org-chart artifact: model architecture was decided in one meeting and serving strategy in another, as if the phone will respect the boundary between them. It won't. On a constrained device, the shape of the model and the shape of the serving stack are a single design decision evaluated against a single budget — memory, latency, and thermals.

Two papers make the halves of this argument, and they're worth reading as one.

## Small models are not shrunk models

[MobileLLM: Optimizing Sub-billion Parameter Language Models for On-Device Use Cases](https://arxiv.org/abs/2402.14905) focuses on the architecture details that matter at the small end: depth, width, and parameter allocation.

A lot of edge planning still begins with the assumption that the main job is to squeeze an existing large model into a smaller box — quantize it, prune it, distill it, ship it. That instinct is understandable, and it leads to awkward compromises. A model that works well in the cloud carries structural assumptions — wide layers that amortize beautifully across an A100's memory bandwidth, attention patterns that assume KV-cache is cheap — that stop making sense once memory pressure, latency budgets, and thermals become the real boss. MobileLLM's finding that thin-and-deep beats wide-and-shallow at sub-billion scale is exactly the kind of result you never discover by compressing downward; you only find it by designing for the target from the start.

That is why model shape matters. Depth, width, and parameter allocation are not abstract architecture debates when the target is a phone. They are part of the deployment contract. Sometimes the winning move is not to compress a cloud model at all — it is to start from a shape that was designed for the edge in the first place.

## NPUs change the serving problem, not just the speed

[Fast On-device LLM Inference with NPUs](https://arxiv.org/abs/2407.05858) supplies the other half. When teams first hear "NPU," it is tempting to translate that into "faster inference" and move on. In practice, specialized acceleration changes the system's shape more than the marketing shorthand suggests. Latency can improve, but the path to *predictable* latency depends on the runtime making better decisions — which is why the paper's ideas like prompt chunking and hardware-aware scheduling matter. Edge serving is rarely "run the same pipeline on smaller hardware." It is a different scheduling problem.

An edge stack that ignores device variation ends up either fragile or overly conservative. Different hardware profiles push the software toward different execution strategies, so the product needs explicit handling for:

- uneven latency behavior across the device fleet,
- fallback paths when the preferred accelerator path is unavailable,
- request shaping that matches the device instead of an abstract average.

This is where "works on my test phone" demos break down. The demo path is a single happy execution route. The product path needs to survive a fleet.

## One budget, one decision

Put the two papers side by side and the shared structure is obvious. MobileLLM says: the hardware budget determines the right model shape. The NPU paper says: the hardware substrate determines the right serving shape. Both are the same claim aimed at different layers — **the constraint shapes the solution space**, and the constraint is the device.

This is hardware-software co-design, and it runs in a loop: the hardware determines what software is efficient, and the software requirements determine what hardware gets built. Platforms have always worked this way — iOS constraints shaped mobile app design, cloud instance economics shaped distributed systems design. The NPU is just the newest constraint in the oldest loop.

The practical consequence is that the two decisions cannot be made in sequence, because each is an input to the other:

- A model shape that behaves well under tight resource limits simplifies the serving layer: fewer ugly runtime compromises, less dependence on aggressive fallback behavior, more predictable performance across device classes.
- A serving strategy that understands the accelerator changes which model shapes are viable: a shape that quantizes cleanly onto the NPU's supported ops beats a nominally better shape that keeps falling back to CPU.

A bad fit in either direction makes the other layer compensate. The serving stack papers over an architecture that was never comfortable on the hardware, or the model gets contorted to survive a runtime that was designed for an abstract average device. Either way, the user gets the compromise.

So the test for an edge effort is not "how small is the model?" or "does it use the NPU?" It is: **are model shape and serving shape being evaluated in the same meeting, against the same measured budget?** Measure the budget, design within it, test at the boundary — the same discipline as fitting firmware to ROM, just with a token stream on top.

Model choice and serving choice are still only two of the coordinated decisions a real deployment needs — compression, fallback policy, and field observability round out the list, which is the subject of [On-Device LLMs: Systems Design](#on-device-llms-are-a-systems-design-problem). And when the budget itself moves — as [1-bit models](#one-bit-models-may-rewrite-the-edge-budget) may move it — both shapes get renegotiated together. That is the point: they were never separate.

## References

1. Zechun Liu et al., [MobileLLM: Optimizing Sub-billion Parameter Language Models for On-Device Use Cases](https://arxiv.org/abs/2402.14905) (2024).
2. Daliang Xu et al., [Fast On-device LLM Inference with NPUs](https://arxiv.org/abs/2407.05858) (2024).
3. Jiajun Xu et al., [On-Device Language Models: A Comprehensive Review](https://arxiv.org/abs/2409.00088) (2024).

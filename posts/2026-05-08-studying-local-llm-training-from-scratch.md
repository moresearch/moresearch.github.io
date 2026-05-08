---
title: Studying local LLM training from scratch
date: 2026-05-08
slug: studying-local-llm-training-from-scratch
summary: I spent time studying Angelos Perivolaropoulos's workshop on training a small LLM locally, and the most important lesson was that tokenization and training discipline matter more than people usually admit.
tags: llms, training, reflections
---

I spent time studying [Angelos Perivolaropoulos's workshop on training an LLM from scratch locally](https://www.youtube.com/watch?v=UsB70Tf5zcE), plus the companion [llm-from-scratch repository](https://github.com/angelos-p/llm-from-scratch), and I think it is one of the better introductions to the topic precisely because it refuses to mystify the process.

[![Training an LLM from Scratch, Locally thumbnail](https://img.youtube.com/vi/UsB70Tf5zcE/maxresdefault.jpg)](https://www.youtube.com/watch?v=UsB70Tf5zcE)

The workshop is not really about “making a tiny ChatGPT.” It is about seeing the transformer pipeline stripped down far enough that every moving part becomes legible: tokenizer, embeddings, self-attention, MLP blocks, residual paths, layer norm, training loop, validation, and sampling.

What I liked most is that the workshop keeps the model small enough to run locally while still preserving the real structure of modern GPT-style training. The companion repo uses a family of configs from tiny to medium, with the default workshop setup landing around a 6-layer, 6-head, 384-dimensional model. That scale is small enough to make experimentation local, but large enough that the important engineering questions still show up.

## The tokenizer is not a preprocessing detail

The first thing that stood out to me is how correctly the workshop treats tokenization as a core modeling decision rather than a boring preprocessing step.

For Shakespeare-scale data, the choice of a character-level tokenizer is not a toy simplification. It is the right systems choice.

The repo makes the case very clearly:

- Shakespeare has about 65 unique characters,
- that means only 65² = 4,225 possible character bigrams,
- those transitions are dense enough that a small model can actually learn them,
- and the embedding/output layers stay tiny.

That last point matters more than many people realize. With `vocab_size=65` and `n_embd=384`, the token embedding table is only about 25K parameters. If you swap in GPT-2's 50,257-token BPE vocabulary at the same embedding width, the embedding table alone jumps to roughly 19.3 million parameters. On a workshop-scale model, that is not a small implementation detail. That is the architecture.

The deeper lesson is that tokenizer choice is really about matching representational granularity to data scale. On a tiny corpus, BPE gives you a vocabulary that is too sparse to learn useful transition structure. Character-level modeling makes the sequence longer, but it gives the model a denser statistical world.

That tradeoff clicked for me very hard while studying the workshop: **sequence length, vocabulary size, and learnable statistics are all coupled**.

## The transformer itself is not the mysterious part

The model architecture is intentionally GPT-2-like:

1. token embeddings,
2. position embeddings,
3. repeated transformer blocks,
4. each block containing causal self-attention plus an MLP,
5. residual connections around both sublayers,
6. layer norm for stability,
7. a projection back to vocabulary logits.

There is nothing magical here, and that is exactly why the workshop is useful.

A lot of people still speak about LLMs as if the mystery lives inside some impossibly exotic block. But once you write the forward path down, the core mechanics are straightforward. Attention produces context-aware token representations; the MLP mixes features position-wise; residual paths preserve gradient flow; layer norm keeps activations sane.

What is more interesting is how these pieces constrain each other. If `n_embd=384` and `n_head=6`, each attention head gets 64 dimensions. That is not just a shape check. It defines the capacity per head, the cost of attention, and the granularity of the similarity computation. Small-model design is mostly about these tradeoffs rather than about novelty.

## The training loop is where the real engineering starts

The strongest message in the workshop is that the training loop matters more than architecture tweaks, and I think that is exactly right.

The objective is standard next-token prediction: input `[t0, t1, ..., tn]`, predict `[t1, t2, ..., tn+1]`. But the workshop makes the practical consequences visible:

- batch construction matters,
- train/validation splits matter,
- the learning-rate schedule matters,
- gradient clipping matters,
- sample generation during training matters,
- and watching validation loss is not optional if you care about overfitting.

This is the kind of detail that separates “I ran a notebook” from “I understand what the model is doing.”

One small piece I kept thinking about is how simple the batch builder is. It just samples random starting offsets, slices `block_size` tokens for `x`, and shifts by one token for `y`. That is conceptually simple, but it encodes the whole autoregressive learning problem:

```go
// makeBatch slices paired input and target windows for next-token prediction.
func makeBatch(tokens []int, starts []int, blockSize int) ([][]int, [][]int) {
	x := make([][]int, 0, len(starts))
	y := make([][]int, 0, len(starts))

	for _, start := range starts {
		// The target window is shifted by one token relative to the input.
		input := append([]int(nil), tokens[start:start+blockSize]...)
		target := append([]int(nil), tokens[start+1:start+blockSize+1]...)
		x = append(x, input)
		y = append(y, target)
	}

	return x, y
}
```

That tiny shift is the whole learning signal. The model is never told about syntax, style, or Shakespearean rhythm directly. It gets only the pressure to predict the next token well, repeatedly, at scale.

## Cosine decay is not cosmetic

I also appreciated that the workshop does not hand-wave optimization. The repo uses warmup, cosine decay, AdamW, and gradient clipping. That is already enough to show why training methodology dominates a lot of outcomes people wrongly attribute to “model intelligence.”

Warmup exists because early optimization steps are fragile. Cosine decay exists because the job changes over time: early on you want exploration and rapid movement; later you want refinement. Gradient clipping exists because small instabilities can still wreck a run, especially when you are learning interactively and changing things quickly.

A lot of frontier-model discussion hides these basics behind scale. This workshop does the opposite. It makes the loop visible enough that you can see the shape of the problem.

## Validation loss and sampling are both debugging tools

One subtle but important point in the workshop is that generation is not only a flashy demo. It is a diagnostic.

If validation loss is improving but samples are still collapsing into garbage, you learn something. If the text starts looking structured and then later begins to regurgitate training fragments, you learn something else. The repo even calls out that peak sample quality often arrives before the end of training, which is a clean reminder that “longer training” is not the same as “better model.”

That is the kind of habit I wish more people carried into practical model work: do not evaluate only through a final scalar loss and do not evaluate only through vibes. Use both.

## The workshop also clarifies what transfers to reasoning and multimodality

I found the later discussion especially useful because it shows how these ideas generalize. A transformer expects sequences of vectors. Once that clicks, it becomes much easier to reason about why language, audio, and other modalities can all fit into related architectures.

The point is not that all modalities are the same. The point is that if you can map them into the right embedding space and preserve the relevant sequence structure, the downstream transformer machinery becomes reusable.

That is also why the workshop feels valuable beyond this exact Shakespeare example. It is teaching the shape of the abstraction, not just one toy exercise.

## My main reflection

The biggest thing I took from studying this workshop is that small local training is useful not because it competes with frontier models, but because it teaches where the real leverage lives.

It lives in:

- tokenizer/data fit,
- parameter budgeting,
- optimization discipline,
- loss interpretation,
- and the relationship between training signals and generated behavior.

If you understand those pieces, larger model systems become much less mystical.

That is why I liked this workshop. It keeps the model compact, but it does not fake the important parts. It shows that even a local, laptop-scale transformer is still a serious engineering object. And once you internalize that, a lot of current LLM discourse starts sounding less like magic and more like systems work.

## Sources

- [Angelos Perivolaropoulos, Training an LLM from Scratch, Locally](https://www.youtube.com/watch?v=UsB70Tf5zcE)
- [angelos-p/llm-from-scratch](https://github.com/angelos-p/llm-from-scratch)
- [Attention Is All You Need](https://arxiv.org/abs/1706.03762)

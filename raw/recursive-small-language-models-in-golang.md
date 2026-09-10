---
title: "Recursive Small Language Models in Golang"
date: 2026-09-10
slug: recursive-small-language-models-in-golang
summary: "How small can the policy inside a Recursive Language Model be before recursion stops paying for itself? This post poses that question, states the working hypotheses, and lays out the method: distill RLM behavior from a big teacher into ibm-granite/granite-4.0-350m, serve it, and harness it in Go with rlm-go and Yaegi. No results yet — this is a research plan, with the experiment written so it can fail."
tags: rlm, recursive-language-models, golang, granite, ibm-granite, fine-tuning, lora, sft, synthetic-data, distillation, yaegi, repl, coding-agents, on-device, quantization, experiments, todo
---

A frontier model is not what makes the Recursive Language Model pattern work. Recursion is. The [RLM paper](https://arxiv.org/abs/2512.24601) from MIT CSAIL (Zhang, Kraska, Khattab) stores the long prompt in a programmatic environment, hands the model a REPL, and lets it write code that inspects the context, slices it, and calls *itself* over the slices. The policy that has to execute those moves — search, decompose, dispatch, submit — is a small policy. It is the knowledge that is large, and the whole design exists to keep the knowledge outside the model.

That observation is what makes a **recursive small language model** practical. The alphaXiv team proved the scaffold is trainable by RL fine-tuning a 4B model to native RLM behavior and matching Claude Sonnet 4.6 on the same harness ([Reinforcing Recursive Language Models](https://www.alphaxiv.org/blog/reinforcement-learning-for-rlms)). But their pipeline ran on a single 8×H200 node with GRPO, which is not how most people will get one. The route that fits a laptop, a single consumer GPU, and a weekend is the one the original paper already used: **distill the behavior from a big teacher into a small student.** Teacher generates RLM trajectories, you filter them, you supervised-fine-tune a 350M model on them, you serve it, you wrap it in a Go RLM harness.

This post is an attempt to answer one question with that pipeline. It is written as a research plan, not a report: the method is assembled from documented parts, the hypotheses are stated so they can be wrong, and the experiment at the end is designed to fail loudly. The student is [`ibm-granite/granite-4.0-350m`](https://huggingface.co/ibm-granite/granite-4.0-350m); the harness is [`rlm-go`](https://github.com/XiaoConstantine/rlm-go), whose interpreter is [Yaegi](https://github.com/traefik/yaegi) — a Go interpreter wired to a model that writes Go, which is the closest Go has come to Python-style interactive programming.

## Research question

> **How small can the policy inside a Recursive Language Model be before the recursion stops paying for itself?**

The RLM paper showed that a frontier model can *use* recursion with no training at all. The alphaXiv result showed that a 4B model can be *trained* to. The open question is where the floor sits when the only affordable training signal is supervised distillation — and whether the answer moves when the harness runs in-process Go instead of a Python subprocess.

Three sub-questions define the edges of it:

1. **Is distillation sufficient?** The alphaXiv recipe reaches native RLM behavior with RL — GRPO over a parent/child policy tree on 8×H200 — and its *first* stage is cold-start SFT on filtered teacher traces. Does that supervised stage alone produce a coherent policy, or is the RL stage what makes it hold together?
2. **Where is the size floor?** The published student is 8B and the trained policy is 4B. Does the scaffold survive at 350M, a model that is a competent instruction follower and a mediocre reasoner? If it survives, what breaks first — syntax, strategy, or routing?
3. **Does the harness move the floor?** If the cost of a recursion is dominated by the number of turns, and every hop across a process boundary is taxed, then a harness with in-process function injection should make deeper and wider recursion affordable. Does that shift the feasible model size, or is the floor set by the model and only the *price* set by the plumbing?

## Working hypotheses

Nothing below is a finding. These are the hypotheses the method is designed to test; each is falsifiable, and several predict failure.

1. **The recursion is the capability, not the weights.** You do not need a model that knows the corpus. You need a model that knows how to *navigate* it: search, slice, recurse, submit. That policy is small enough for 350M parameters.
2. **RL is the optimization; distillation is the on-ramp.** The alphaXiv recipe's first stage is a cold-start SFT on teacher traces. If you stop there, you still get a working RLM — you just cap it at the teacher's demonstrated strategy instead of letting it discover its own.
3. **Synthesize the dataset by running the teacher *as an RLM*.** Generate trajectories through the same harness you will deploy, log them, and train on the logs. The student then learns the exact syntax the harness parses, not a paraphrase of it.
4. **Turn-level data, not rollout-level.** The RLM harness rewrites the user turn each iteration rather than accumulating it, so successive turns share no prefix. An N-turn trajectory is N training samples. This is the single most common way to build the dataset wrong.
5. **Cold start is real at this size.** A 350M instruct model will not reliably emit valid REPL code, call `FINAL`, or stop, no matter how good the prompt is. The alphaXiv team measured 0 pass@16 for a 4B model in the scaffold; at 350M, assume the same until proven otherwise.
6. **Keep the supervised set small and brutal.** The alphaXiv run filtered teacher rollouts and SFT'd on a few dozen held-out examples because training on the full set caused entropy collapse. Filter for *executable and correct*, not for volume.
7. **Split the policy: small root, big children.** `rlm.New` takes separate clients for root orchestration and sub-calls. A 350M fine-tune can be the cheap, fast decomposer while a frontier model does the hard extraction — the recursion hides the size mismatch.
8. **A Go harness with a Go interpreter is a Go REPL.** Yaegi evaluates Go against the standard library with shared state; `rlm-go` injects `Query`/`FINAL` into it and lets the model write the programs. That combination is the Python-like interactive loop Go never shipped.
9. **The pipeline is a plan, not a result.** Every stage is documented and standard, but the *composition* is untested. The [method](#method-the-experiment-todo) section below states the hypothesis, the metrics, and the conditions that would falsify it — because a recursive small model is exactly the kind of idea that sounds good until you measure the valid-code rate.

## Why 350M is a serious number

Granite-4.0-350M is not a toy distil of a bigger sibling. It is a 352,379,904-parameter decoder-only dense transformer — 28 attention layers, embedding size 1024, grouped-query attention, SwiGLU MLPs, RMSNorm, tied input/output embeddings, 32K context — released by IBM under Apache 2.0 on October 28, 2025, and explicitly aimed at on-device use and domain fine-tuning without massive compute.

Its published numbers set expectations honestly: MMLU 35.01, IFEval average 55.4, HumanEval 39, MBPP 48, BFCL v3 39.32 for tool calling. It is a competent instruction follower and a mediocre reasoner. That is exactly the profile the RLM pattern wants, because the pattern removes the need to reason over the whole corpus at once.

| What a 350M model is bad at | What the RLM pattern asks of it instead |
|---|---|
| Holding a 200-page document in context | Slicing the document with code and delegating |
| Multi-hop reasoning over facts | Deciding *which* slice to recurse into |
| Deep, long-chain inference | Emitting valid REPL syntax and terminating |
| Knowing your domain | Choosing a query, a keyword, or an index |

The bet is that navigation is a lower-dimensional skill than knowledge. A teacher can demonstrate navigation; the student cannot be made to memorize the corpus. Distillation is therefore a much better fit here than it usually is.

## The pipeline

```text
  frontier teacher ──► synthetic RLM trajectories ──► filter ──► SFT samples
       (big)              (logged JSONL)              (correct,   (turn-level)
                                                        clean)          │
                                                                        ▼
                                              ibm-granite/granite-4.0-350m
                                                   (LoRA / full SFT)
                                                                        │
                                                                        ▼
  Go coding agent  ◄──── rlm-go + Yaegi REPL ◄──── vLLM or Ollama
                    (root=350M, children=bigger)
```

Four stages, each independently useful. You can run stage 4 against the stock Granite model today and see how much stage 2 buys you.

## Stage 1 — Synthesize the dataset with a bigger model

### Pick a task that is checkable

Distillation needs a filter, and filtering needs a signal. Choose a task where a teacher trajectory either ends with a verifiable answer or does not. Repository question-answering is a good one: given a question and a repository dump in the REPL's context, return the file, line range, and symbol that answers it. It is decomposable (one child per file or package), it forces real use of the REPL, and correctness is checkable by string match against a gold answer plus a quick human pass.

It also doubles as the coding-agent eval, which is why this is worth more than an academic exercise.

### Run the teacher as an RLM, through the harness you will deploy

The cheapest correct way to generate trajectories is to let the teacher *be* an RLM using `rlm-go`, so the logs are already in the deployment turn format:

```bash
# Install the harness (the command from the README)
go install github.com/XiaoConstantine/rlm-go/cmd/rlm@latest

export ANTHROPIC_API_KEY=sk-...

# One question, one repo dump, one recorded trajectory.
rlm -model claude-sonnet-4-20250514 \
    -context "$(cat repo.txt)" \
    -query "Which function enforces the max-iteration guard, and where is the loop it bounds?" \
    -max-iterations 12 \
    -json \
    -log-dir ./teacher-logs
```

Run that over a few hundred questions and you have a corpus. The `-log-dir` JSONL record is the important part: it captures session metadata, per-iteration prompts and responses, the executed code, and the sub-call records with token counts. It was built as a debugging and visualization aid — it is compatible with the Python RLM visualizer — but it is also a perfectly good training-set recorder.

### Filter, then cut the trajectory into turn-level samples

Two filters, both mandatory:

- **Executable.** Drop any trajectory where the logged code produced a REPL error or a panic. A student trained on broken Go learns broken Go.
- **Correct.** Keep only trajectories whose final answer matches the gold answer. For a repo task, that means final file and symbol match, not a fuzzy judge, because you generated the labels yourself.

Then reconstruct training samples per turn. This is insight #4 and it is easy to get wrong:

```python
import json, glob

SYSTEM = open("system_prompt.txt").read()  # the exact harness prompt
samples = []

for path in glob.glob("teacher-logs/*.jsonl"):
    turns = [json.loads(l) for l in open(path)]
    if not turns or turns[-1].get("final_correct") is not True:
        continue
    for t in turns:
        if t.get("role") != "assistant" or not t.get("content"):
            continue
        if t.get("repl_error"):
            continue
        # The harness re-appends the user turn each iteration: no shared prefix.
        samples.append({
            "messages": [
                {"role": "system", "content": SYSTEM},
                {"role": "user", "content": t["user_turn"]},   # query + REPL output
                {"role": "assistant", "content": t["content"]},  # a ```go block or FINAL(...)
            ]
        })

with open("granite_rlm_sft.jsonl", "w") as f:
    for s in samples:
        f.write(json.dumps(s) + "\n")

print(len(samples), "turn-level samples")
```

The `user_turn` must be exactly what the harness will send — the re-appended query plus the `REPL output:` block — because the student will only ever see that shape. A sample record looks like this:

```json
{"messages": [
  {"role": "system", "content": "You are a coding RLM. Inspect `context` with Go code. ..."},
  {"role": "user", "content": "Which function enforces the max-iteration guard? ...\n\nREPL output:\n(pending)"},
  {"role": "assistant", "content": "```go\nhits := []string{}\nfor i, line := range strings.Split(context, \"\\n\") {\n    if strings.Contains(line, \"maxIterations\") {\n        hits = append(hits, fmt.Sprintf(\"%d: %s\", i+1, line))\n    }\n}\nfmt.Println(strings.Join(hits, \"\\n\"))\n```"}
]}
```

**How much data.** Start with a few hundred trajectories and a few thousand turn-level samples. The alphaXiv group deliberately SFT'd on a *few dozen* held-out examples for cold start and warned that supervising on the full training set caused entropy collapse and instability. Volume is not the goal; executable, correct, in-format trajectories are. Held-out slices should stay human-audited so you can tell when the student is confident and wrong.

## Stage 2 — Fine-tune granite-4.0-350m

A 350M model in bf16 is roughly 0.7 GB of weights. That changes the economics completely: LoRA fits on a small consumer GPU, full fine-tuning fits on a single 24 GB card, and neither needs a training cluster. Do not train it on CPU unless you enjoy waiting.

The standard path is Hugging Face `transformers` + `peft` + `trl`:

```python
from datasets import load_dataset
from peft import LoraConfig
from trl import SFTConfig, SFTTrainer

ds = load_dataset("json", data_files="granite_rlm_sft.jsonl", split="train")

peft_config = LoraConfig(
    r=32,
    lora_alpha=64,
    lora_dropout=0.05,
    target_modules="all-linear",   # adapter on every linear projection, incl. GQA and MLP
    task_type="CAUSAL_LM",
)

args = SFTConfig(
    output_dir="granite-rlm-350m",
    num_train_epochs=2,
    per_device_train_batch_size=8,
    gradient_accumulation_steps=4,
    learning_rate=1e-4,            # LoRA; use ~1e-5 for full fine-tuning
    bf16=True,
    max_length=8192,               # RLM turns are long; 4K truncates the useful part
    packing=False,                 # keep each turn a separate example
    logging_steps=10,
    save_strategy="epoch",
    # Mask the loss to assistant turns so the student never learns to
    # generate the user turn or the REPL output. Requires a chat template
    # with generation markers; otherwise mask manually before training.
    assistant_only_loss=True,
)

trainer = SFTTrainer(
    model="ibm-granite/granite-4.0-350m",
    args=args,
    train_dataset=ds,
    peft_config=peft_config,
)
trainer.train()
trainer.save_model("granite-rlm-350m-lora")
```

Three things to get right beyond the hyperparameters.

**Use the model's own chat template.** Granite 4.0 speaks `<|start_of_role|>assistant<|end_of_role|>...<|end_of_text|>`, and the tokenizer knows it. Let `tokenizer.apply_chat_template` build the strings rather than hand-rolling the special tokens, or the fine-tune will be trained on a format the serving stack does not reproduce. `SFTTrainer` does this for you when your dataset is in `messages` form — which is why stage 1 wrote it that way.

**Mask the loss to assistant tokens.** The harness re-sends the system prompt and the user turn; the student should learn only the Go block. Training on the user turn teaches it to hallucinate REPL output, which is a spectacular failure mode to debug. If your chat template lacks generation markers, apply the mask by hand rather than skipping it.

**Merge before serving.** LoRA adapters have to be folded into the base weights for most serving paths:

```python
from peft import PeftModel
from transformers import AutoModelForCausalLM

base = AutoModelForCausalLM.from_pretrained("ibm-granite/granite-4.0-350m")
model = PeftModel.from_pretrained(base, "granite-rlm-350m-lora").merge_and_unload()
model.save_pretrained("granite-rlm-350m-merged")
```

On the role of this stage: it is *imitation*. It teaches the harness contract — valid Go blocks, the exposed functions, `FINAL`, terminating — and a task strategy the teacher demonstrated. It will not discover a better strategy, and it will not exceed the teacher. That is what the alphaXiv team's RL stage is for, and it is a genuinely different project: SkyRL, a Python REPL environment, GRPO over root rollouts with advantages inherited by children. Distillation gets you a deployable recursive model; RL is how you make it inventive.

## Stage 3 — Serve it

For development and eval, an OpenAI-compatible server is the shortest path:

```bash
pip install vllm

vllm serve ./granite-rlm-350m-merged \
  --served-model-name granite-rlm \
  --max-model-len 32768 \
  --port 8000
```

For the edge case that motivated the model in the first place, convert to GGUF and run it under Ollama. Granite 4.0 is already in the Ollama library (`granite4:350m`), so you have a stock baseline to diff against, and a custom build is a `Modelfile` away:

```bash
# after converting the merged model to GGUF with llama.cpp
ollama create granite-rlm -f Modelfile
ollama run granite-rlm "Which function enforces the max-iteration guard?"
```

I would test both against the same held-out questions as the stock `granite4:350m`. The delta between them is the only honest measure of whether the synthetic dataset worked.

## Stage 4 — Harness it: rlm-go, Yaegi, and a Go coding agent

Now the deployment harness. `rlm-go` reflects its REPL functions directly into a [Yaegi](https://github.com/traefik/yaegi) interpreter in the same process — no sockets, no IPC, no subprocess — and the project claims roughly 100× less latency per sub-LLM call than the socket-based Python design. In a recursive agent, whose cost is a product over turns, per-hop overhead compounds where model cost does not.

The API takes **two clients**, and that is the architectural sweet spot for a small model:

```go
// root client   → the fine-tuned 350M (cheap, fast, decomposes)
// repl client   → a frontier model (does the hard child reasoning)
r := rlm.New(graniteClient, frontierClient,
    rlm.WithMaxIterations(12),
    rlm.WithMaxRecursionDepth(2),
    rlm.WithSandbox(),
)
```

The 350M student runs the loop — inspecting context, splitting it, deciding where to recurse, calling `FINAL`. The expensive model only ever sees the small, already-selected slices that the student hands it. That is the mixed policy the recursion makes possible, and it is a better use of a distilled 350M than hoping it can do extraction well.

You need a client for your local endpoint. `rlm-go` ships Anthropic, Gemini, and OpenAI providers, but the OpenAI client's base URL is set internally, so for a local server you implement the two interfaces directly — they are small:

```go
type GraniteClient struct {
    baseURL string // e.g. "http://localhost:8000"
    model   string // "granite-rlm"
    http    *http.Client
}

// Root orchestration.
func (c *GraniteClient) Complete(ctx context.Context, msgs []core.Message) (core.LLMResponse, error) {
    out, pt, ct, err := c.chat(ctx, msgs)
    return core.LLMResponse{Content: out, PromptTokens: pt, CompletionTokens: ct}, err
}

// Sub-LLM calls made from inside the REPL.
func (c *GraniteClient) Query(ctx context.Context, prompt string) (core.QueryResponse, error) {
    out, pt, ct, err := c.chat(ctx, []core.Message{{Role: "user", Content: prompt}})
    return core.QueryResponse{Response: out, PromptTokens: pt, CompletionTokens: ct}, err
}

// Concurrent child rollouts — one goroutine per prompt.
func (c *GraniteClient) QueryBatched(ctx context.Context, prompts []string) ([]core.QueryResponse, error) {
    out := make([]core.QueryResponse, len(prompts))
    g, ctx := errgroup.WithContext(ctx)
    for i, p := range prompts {
        i, p := i, p
        g.Go(func() error {
            r, err := c.Query(ctx, p)
            out[i] = r
            return err
        })
    }
    return out, g.Wait()
}
```

The `chat` helper is a thin `POST /v1/chat/completions`. Register your repo tools next, using the harness's injection point:

```go
r := rlm.New(graniteClient, frontierClient,
    rlm.WithSandbox(),
    rlm.WithREPLSetup(func(env *repl.REPL) error {
        return env.InjectSymbols(map[string]reflect.Value{
            "ReadFile": reflect.ValueOf(readFileTool),
            "GitGrep":  reflect.ValueOf(gitGrepTool),
            "RunTests": reflect.ValueOf(runTestsTool),
        })
    }),
)

result, err := r.Complete(ctx, repoDump, "Where is the max-iteration guard enforced?")
```

`InjectSymbols` merges your Go functions into the interpreter's namespace, dot-imported so the model calls `RunTests("./pkg/rlm")` as an unqualified function inside generated Go. Collisions with builtins are rejected rather than silently shadowed, which is the right default when the tool surface is the safety surface.

And inside the REPL, the model's turns look like this — code that is exactly what stage 1 taught it to emit:

```go
// Locate candidate call sites, then fan out one child RLM per file.
hits := []string{}
for i, line := range strings.Split(context, "\n") {
    if strings.Contains(line, "WithMaxIterations") || strings.Contains(line, "maxIterations") {
        hits = append(hits, fmt.Sprintf("%d: %s", i+1, strings.TrimSpace(line)))
    }
}
fmt.Println(strings.Join(hits, "\n"))

answers := QueryBatched(hits)   // small slices → frontier children
fmt.Println(strings.Join(answers, "\n---\n"))
FINAL(hits[0])
```

`QueryBatched` is the prefix-trap escape from the paper applied to a repository: instead of letting the model anchor on the first file it reads, it dispatches several children at once. On a single-process Go harness, those are goroutines, not spawned processes.

## The Yaegi part: Go finally has a REPL, and RLM is what it's for

Go has never had a good interactive story. It is a compiled language; `go run` starts a process, and there is no built-in way to evaluate an expression against live state the way Python, Ruby, and Lisp do. Yaegi exists precisely to fill that gap: `New()` / `Eval()` / `Use()`, an embedded interpreter with "complete support of the Go specification," plus a command-line REPL of its own.

On its own, though, a Go interpreter is a solution looking for a user. An interactive shell only matters if there is something worth interacting with, and most Go work is better served by compiling. What `rlm-go` does is supply the missing half: it loads Yaegi with the standard library, injects `Query`, `QueryBatched`, `FINAL`, and your tools into the interpreter's scope, and then puts a model behind the keyboard. The result is a stateful session that can read the filesystem, call functions, observe real output, and decide what to do next — which is precisely the loop a Python user takes for granted and a Go user has never had in the same way.

That is the claim, stated carefully: **the RLM harness over Yaegi is one of the closest things Go has to Python-like interactivity.** Not because Yaegi is a better shell than the alternatives, but because it turns "evaluate Go against shared state with real side effects" into a primitive that an autonomous agent can drive. The state is the `context` variable. The function calls are your tools. The recursion is `Query`. The completion signal is `FINAL`. And it is all in one process, in your binary, with no service to run alongside it.

Go's concurrency model is the other half of why this lands. The recursion tree *is* a fan-out; `QueryBatched` is an `errgroup`; per-child timeouts are `context.WithTimeout`. The alphaXiv team's training run hit 512 concurrent rollouts and found race conditions around child REPL timeouts — a Go problem wearing ML clothes. In Go, those are solved problems with standard tooling.

## Method: the experiment (TODO)

**Status: not run.** Everything above is assembled from documented parts — the paper, the RL blog, the Granite model card, and the `rlm-go` API — but the composition is a hypothesis, not a report. This is the experiment that would answer the research question, written so that it can fail.

**Hypothesis.** A 350M Granite RLM, SFT'd on a few hundred filtered teacher trajectories and used as the *root* decomposer with a frontier model as the *child*, recovers the bulk of a frontier-only RLM's accuracy on repository QA at a fraction of the token cost and wall-clock time.

**Setup.** One harness (`rlm-go` + Yaegi), one held-out and human-audited question set, one repository corpus. Four arms:

| Arm | Root | Children | What it tests |
|---|---|---|---|
| **A** | stock `granite4:350m` | frozen 350M | Can a 350M model do RLM behavior at all without training? (prediction: valid-code rate collapses — insight #5) |
| **B** | fine-tuned 350M | fine-tuned 350M | Pure small recursion: how far does distillation alone get? |
| **C** | fine-tuned 350M | frontier | **The bet:** cheap decomposition plus expensive extraction |
| **D** | frontier | frontier | Ceiling: the published RLM behavior |

**Metrics per arm:** task accuracy (final file and symbol match), total tokens across root and children, wall-clock time, turns to `FINAL`, **valid-Go-block rate**, **REPL error rate**, and **cost per query**. The last three are the ones that matter — a recursive small model fails in ways accuracy alone hides, and arm A is the control that shows the difference training made.

**Falsification.** Kill the hypothesis if arm C does not reach roughly 90% of arm D's accuracy at 30% or less of D's token cost, or if arm B's valid-Go-block rate plateaus below ~80% after fine-tuning. Either outcome is still a post; that is the point of writing the criteria down first.

**The ablations I actually care about**, each isolating one claim from the RL blog:

- **Turn-level vs rollout-level samples.** Train two students, one on per-turn samples and one on whole rollouts. Prediction: the rollout-level student learns to hallucinate REPL output, because it is being taught to predict turns it will never be asked to produce.
- **Data volume sweep.** A few dozen → a few thousand trajectories, with the held-out eval run at each point. This traces the entropy-collapse curve the alphaXiv team hit; I want to see where it bends at 350M rather than at 4B.
- **Loss masking on/off.** Does masking the user turn matter more at 350M than at 4B? The guess is that at this size it is the difference between a usable model and a broken one.
- **Prompt length, 1500 → 200 tokens.** The direct test of whether strategy moved into the weights. The alphaXiv ablation converged slightly *below* the long-prompt run with the short one; at 350M the gap should be wider, and that gap is the honest price of skipping the RL stage.
- **LoRA vs full fine-tune.** 350M is small enough that full fine-tuning is affordable, so if LoRA leaves measurable quality on the table here, that is worth knowing before anyone scales the recipe up.

**Why this is a good weekend project.** Every stage is independently runnable and independently useful: stage 1 produces a log corpus whether or not you train anything; stage 3 leaves you with a servable model; stage 4 is a working agent you can point at the stock Granite model today. If the fine-tune flops, you still have a Go RLM harness and a synthetic trajectory dataset.

**What would make me abandon it.** If arm B's valid-Go-block rate plateaus well below arm D even after filtering and the data sweep, then 350M is simply too small for this scaffold on this task, and the correct response is a 1B or 4B student — not more cleverness in the harness.

## What to expect, honestly

> **Note.** We start small and try bigger models after. The whole point of the LM in this pattern is recursive reasoning across iterations — a 350M model will produce incoherent Go code inside the Yaegi REPL and will not complete a real analysis. The 350M run is the floor test, not the destination: the point of doing it first is to find exactly where the recursion breaks before paying for a bigger student. Treat any 350M result as a lower bound on the harness, not an upper bound on the idea.

**The 350M student learns the scaffold, not the knowledge.** Expect it to emit valid Go, call the right functions, slice sensibly, and terminate. Expect its extraction and reasoning to be weak — MMLU 35 and HumanEval 39 are the honest priors. Split the policy as shown and let the frontier model handle the child work.

**SFT caps at the teacher.** Imitation cannot exceed the demonstrations, and it inherits their biases, including any leakage in the synthetic set. Keep a human-audited held-out set, and evaluate with a real metric, not vibes.

**This is not the RL result.** The alphaXiv team's number — a 4B model matching Sonnet 4.6 on the same harness — came from GRPO with inherited advantages across a parent/child policy tree, not from SFT. Distillation is the cold start. If you want strategy discovery, you need the RL stack (they published the SkyRL environment), and you would be rebuilding the harness environment in Python.

**Yaegi is an interpreter, not the Go toolchain.** It has real limitations: no cgo, no `//go:embed` directives, no Go modules inside the interpreter, and slower execution than compiled code. It will run the model's code; it will not build your repository. `go build`, `go test`, and `git` have to be wired in as injected tools.

**Sandbox model-generated code.** In-process Yaegi is fast and unisolated; the interpreter does not export `unsafe` or `syscall` by default, but your injected tools still run with your process's privileges. `rlm.WithSandbox()` auto-detects Podman or Docker, disables the network by default, and pays 50–200 ms per execution for a real boundary. For a coding agent that runs model-authored Go, the sandbox should be the default rather than the escape hatch.

**Tokens are the meter.** `rlm-go` aggregates token usage across the root and every sub-call, so you can price a recursion tree instead of estimating it. A recursive small model is cheap per call and expensive per query; measure both.

## What an answer would change

The question is worth the weekend because both answers are useful, and they point in different directions.

**If 350M works,** recursive small models become a default local pattern rather than a demo: an on-device agent that decomposes, slices, and routes, with a frontier model reserved for the handful of child calls that actually need it. The interesting moat then moves to the teacher and the harness — the data you can synthesize and the environment you can check — which is the same conclusion the RLM-as-mechanism argument reached from the training side.

**If 350M fails,** the result is still a boundary: it tells you the floor is set by the model and not the plumbing, and that "small" in recursive small language models means 1B–4B, not 350M. That would also separate two claims this blog has been making together — that harnesses are reward infrastructure (capability) and that harness latency is on the critical path (cost). A failure here would show that harness engineering buys you cost and latency, not capability, which is a sharper and more honest version of the thesis.

**Either way, the missing measurement is the same one.** Nobody has published an SFT-only, 350M, Go-harnessed RLM with a cost/quality frontier against a frontier baseline. That frontier — accuracy plotted against tokens and wall-clock, across the four arms — is the actual contribution, and it does not depend on the hypothesis surviving.

## What this question changed in my view

I had filed the RLM training story under "needs a lab" — 8×H200, GRPO, a policy tree, advantage inheritance. Reading the alphaXiv blog closely changed that: their own recipe starts with a cold-start SFT on filtered teacher traces, which is distillation, which is a weekend on a consumer GPU. The RL is what makes the model inventive. The distillation is what makes it *work*, and it is available to anyone with an API key and a 350M base model.

The second change is about Go. I came in thinking of the RLM harness as plumbing — an interpreter, some function injection, a loop. The Yaegi angle reframed it: Go's weakness has always been that it has no REPL, no place where you define a function, inspect state, and iterate interactively. `rlm-go` accidentally ships one, because a language model writing Go into an embedded interpreter *is* interactive programming. The model supplies the intent, Yaegi supplies the evaluation, and the context object supplies the state. It is the Python-like loop Go never had, aimed at a user who is not a human.

That is the practical version of the paper's claim. Context is an environment; the model is a program in it; the program can be small; and the environment can be one binary in the language you already deploy. Fine-tune the 350M, and the loop is yours.

## References

- Zhang, Kraska, Khattab. [Recursive Language Models](https://arxiv.org/abs/2512.24601) (MIT CSAIL) — context as an external environment, recursive self-calls, two orders of magnitude beyond the context window, RLM-Qwen3-8B at +28.3% over Qwen3-8B.
- Kim, Ahmad. [Reinforcing Recursive Language Models](https://www.alphaxiv.org/blog/reinforcement-learning-for-rlms) (alphaXiv, May 2026) — RL fine-tuning 4B RLMs, one shared parent/child policy, inherited advantages, stepwise GRPO, rubric judges, cold-start SFT, entropy collapse at scale.
- Granite Team, IBM. [`ibm-granite/granite-4.0-350m`](https://huggingface.co/ibm-granite/granite-4.0-350m) — Apache 2.0, 352M parameters, 32K context, tool calling; base model [`granite-4.0-350m-base`](https://huggingface.co/ibm-granite/granite-4.0-350m-base); [Nano language models repo](https://github.com/ibm-granite/granite-4.0-nano-language-models).
- XiaoConstantine. [rlm-go](https://github.com/XiaoConstantine/rlm-go) — Go RLM harness, direct function injection via Yaegi, `InjectSymbols`, Podman/Docker sandbox, JSONL session logging; install with `go install github.com/XiaoConstantine/rlm-go/cmd/rlm@latest`.
- Traefik. [Yaegi](https://github.com/traefik/yaegi) — embedded Go interpreter and REPL; see its documented limitations (no cgo, no embed directives, no modules in the interpreter).
- Hugging Face. [`transformers`](https://github.com/huggingface/transformers), [`peft`](https://github.com/huggingface/peft), [`trl`](https://github.com/huggingface/trl) — the SFT/LoRA stack used in stage 2; [Unsloth](https://github.com/unslothai/unsloth) is the other common drop-in for low-VRAM fine-tuning.
- vLLM. [OpenAI-compatible server](https://docs.vllm.ai/) — `vllm serve`; Ollama's [`granite4:350m`](https://ollama.com/library/granite4) for the local/edge path.
- NovaSky-AI. [SkyRL](https://github.com/NovaSky-AI/SkyRL) — the RL stack and RLM environment behind the alphaXiv run, if you want to go past distillation.
- Rajani et al. [Scalpel vs. Hammer: GRPO Amplifies Existing Capabilities, SFT Replaces Them](https://arxiv.org/abs/2507.10616) — why a supervised phase sets the ceiling rather than raising it.
- This blog: [RLMs Are the New Reasoning Models](https://blog.hackspree.com/#rlms-are-the-new-reasoning-models), [Harness Patterns for Agentic AI Systems](https://blog.hackspree.com/#harness-patterns-for-agentic-ai-systems), [Go Is Good for Harness Pipelines](https://blog.hackspree.com/#go-is-good-for-harness-pipelines), [Go Can Keep Structured LLM Runtimes Boring](https://blog.hackspree.com/#go-can-keep-structured-llm-runtimes-boring), [Sandboxing AI Agents](https://blog.hackspree.com/#sandboxing-ai-agents), [Coding Agent Harnesses Need Real Repositories](https://blog.hackspree.com/#coding-agent-harnesses-need-real-repositories), [Verification Is the Bottleneck](https://blog.hackspree.com/#fowler-retreat-verification-harness-engineering).

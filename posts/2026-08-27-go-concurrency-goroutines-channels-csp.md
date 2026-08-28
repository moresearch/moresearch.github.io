---
title: "Go Concurrency: Goroutines, Channels, and the CSP Model That Ships"
date: 2026-08-27
slug: go-concurrency
summary: "Three insights: the goroutine makes concurrency cheap enough to be a design default — a language feature, not a library; channels and select make coordination compositional, so most patterns are channel patterns; sync, context, and the race detector are where production concurrency bugs actually live. Plus the six books that teach it properly."
tags: [golang, concurrency, csp, channels, goroutines, sync, context, books, engineering]
---

Concurrency is a language feature in Go, not a library. Most languages bolt threading onto the runtime and let the ecosystem figure out coordination; Go embeds the two ideas — lightweight units of execution and typed message passing — into the syntax itself. The design descends from Hoare's Communicating Sequential Processes (1978) via the language's own canonical advice: *do not communicate by sharing memory; instead, share memory by communicating*. Three insights follow.

## Insight 1 — The goroutine is the unit of concurrency, not the thread

**A goroutine is cheap enough that concurrency can be the default design, not the special case.** A goroutine is a function call with `go` in front of it: it starts on a stack of a few kilobytes that grows and shrinks on demand, and the runtime's M:N scheduler multiplexes millions of them onto a handful of OS threads.

The cost ceiling is what changes the design. In thread-based languages you wrap code in threads sparingly, because threads are expensive to create, expensive to switch, and limited in number. In Go the default assumption is that every independent unit of work gets its own goroutine — the `net/http` server, for instance, handles each connection in one. The smallest complete example is the fan-out:

```go
var wg sync.WaitGroup
for _, url := range urls {
    wg.Add(1)                 // count this unit of work *before* the goroutine starts
    go func(u string) {
        defer wg.Done()       // decrement when this goroutine finishes
        fetch(u)              // each independent unit of work runs in its own goroutine
    }(url)
}
wg.Wait() // blocks until every goroutine has called Done
```

Two rules follow. First, `Add` is called before the goroutine runs, never inside it — the classic `WaitGroup` race is counting work that has already finished. Second, goroutines share the heap, so sharing is only safe when it is coordinated. Passing the loop variable as a parameter (as above) keeps each goroutine's data private; that is the point of the next insight.

## Insight 2 — Channels make coordination compositional; select is the arbiter

**Channels move ownership instead of sharing state, and `select` composes them into anything.** A channel is a typed conduit. An unbuffered channel is a rendezvous — sender and receiver block until both are present — so it is a synchronization primitive as much as a data structure. A buffered channel decouples producers from consumers up to a bounded queue, which is how you add backpressure without locks.

Most Go concurrency patterns are channel patterns: worker pools, pipelines, fan-in/fan-out. The worker pool is the pattern you will reach for first:

```go
// jobs is unbuffered (each worker blocks until work exists);
// results is buffered so consumers can drain without waiting on producers
jobs := make(chan int)
results := make(chan int, 8)

// worker reads jobs until the channel is closed, then exits
worker := func(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs { // range ends when the producer calls close(jobs)
        results <- id * job
    }
}

// start a fixed pool, then feed it and shut it down cleanly
const n = 4
for i := 0; i < n; i++ {
    go worker(i, jobs, results)
}
for j := 0; j < 10; j++ {
    jobs <- j
}
close(jobs) // closing the channel is the shutdown signal
```

The built-in `select` statement is the construct other languages still lack: it waits on many channel operations at once and runs the first one that is ready, which turns timeouts, cancellation, and multiplexing into syntax instead of libraries:

```go
// select runs the first ready case; the timeout case guards a hung channel
select {
case job := <-jobs:
    process(job)
case <-time.After(2 * time.Second):
    log.Println("no job arrived within 2s")
case <-ctx.Done(): // cancellation wins even when jobs are ready
    return ctx.Err()
}
```

The "share memory by communicating" slogan is not a moral stance; it is a practical one. Passing a value over a channel transfers ownership at a point where both sides can observe it, and the race detector — see Insight 3 — agrees: values sent over channels are the easiest shared state to prove correct.

## Insight 3 — When channels don't fit, sync is where production bugs live

**The race detector is non-negotiable, and `context` is how cancellation composes.** Channels cover the composition cases; the `sync` package covers the rest — `Mutex` and `RWMutex` for shared state you cannot avoid, `Once` for lazy initialization, `Pool` for reusable buffers, `WaitGroup` for fan-out.

Two practices separate "mostly working" from production-grade. First, the race detector: `go test -race` belongs in every CI run, because data races are undefined behavior that only manifests under load. Second, `context.Context`: cancellation and deadlines are values that must flow down the goroutine tree, or a single hung request leaks goroutines forever. The `errgroup` package (golang.org/x/sync/errgroup) is the idiomatic join of both — fan-out with error propagation and shared cancellation:

```go
// errgroup joins fan-out with error propagation and group-wide cancellation
g, ctx := errgroup.WithContext(ctx)
for _, task := range tasks {
    // task is per-iteration since Go 1.22; pass it explicitly anyway for clarity
    g.Go(func() error {
        select {
        case <-ctx.Done():
            return ctx.Err() // a sibling failed: stop early
        default:
        }
        return run(task) // the first non-nil error cancels the group
    })
}
if err := g.Wait(); err != nil { // blocks until every goroutine returns
    log.Fatalf("one task failed, the group was cancelled: %v", err)
}
```

The failure modes are well catalogued: goroutines blocked forever on a channel nobody closes, `WaitGroup.Add` called inside the goroutine instead of before it, mutexes copied after first use, nil channels silently disabling a `select` case. One book below devotes a third of its chapters to exactly this catalogue — because these are the bugs that ship.

## Insight 4 — Concurrency is not parallelism; the model is the moat

**Design for composition; the runtime delivers parallelism — and the books teach the model, not just the API.** Rob Pike's 2012 talk drew the line that still matters: concurrency is the *composition* of independently executing computations; parallelism is *simultaneous execution*. Go gives you the composition primitives; the scheduler decides about parallelism. Most production bugs come from forgetting the first half, not the second.

That is why the books matter more here than in most languages: the API surface is small, so the difference between a mediocre and a great Go engineer is the mental model — CSP, ownership, the memory model — not the number of packages memorized. It is also why this blog keeps returning to Go for AI-adjacent infrastructure: Go is the tool that stays boring and readable when the system grows new branches ([why Go still matters in the AI era](https://blog.hackspree.com/#why-go-still-matters-in-the-ai-era), [Go can keep structured LLM runtimes boring](https://blog.hackspree.com/#go-can-keep-structured-llm-runtimes-boring)). Concurrency is where that bet pays off first.

## The books that teach it properly

| Book | What it teaches | Why it matters |
|---|---|---|
| **The Go Programming Language** — Donovan & Kernighan | The language, including Ch. 8 (goroutines & channels) and Ch. 9 (shared variables, the race detector) | The canonical foundation before any concurrency book |
| **Concurrency in Go** — Katherine Cox-Buday | The only book devoted entirely to the topic: CSP mental models, channels and select, sync, worker pools, context | The core recommendation; read it cover to cover |
| **100 Go Mistakes and How to Avoid Them** — Teiva Harsanyi | Part 3 (Concurrency): goroutine leaks, WaitGroup misuse, channel semantics, mutex ownership | The checklist to run before shipping |
| **Go in Action** — Kennedy, Ketelsen, St. Martin | Ch. 7: concurrency patterns — work queues, pipelines, fan-out — in a pragmatic walk-through | The patterns primer when you want working code first |
| **Distributed Services with Go** — Travis Jeffery | Concurrency under real network load: gRPC servers, log replication, Raft | The production view: what goroutines do when a network is involved |
| **The Art of Multiprocessor Programming** — Herlihy, Shavit, et al. | The theory behind the primitives: why mutexes are hard, memory models, lock-free structures | The why-behind-the-tooling; explains why the race detector exists |

## Key insight

**Concurrency is a design default in Go, not an optimization.** Cheap goroutines make composition the natural style; channels make ownership explicit; `sync`, `context`, and the race detector make the residual complexity survivable. The books matter because the model — CSP, composition over simultaneity — is the actual product.

## References

- The Go Programming Language (Donovan & Kernighan) — https://www.gopl.io/
- Concurrency in Go: Tools and Techniques for Developers (Cox-Buday, O'Reilly) — https://www.oreilly.com/library/view/concurrency-in-go/9781491941195/
- 100 Go Mistakes and How to Avoid Them (Harsanyi, Manning) — https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
- Go in Action (Kennedy, Ketelsen, St. Martin, Manning) — https://www.manning.com/books/go-in-action
- Distributed Services with Go (Jeffery, Pragmatic Bookshelf) — https://pragprog.com/titles/tjgo/distributed-services-with-go/
- The Art of Multiprocessor Programming, 2nd ed. (Herlihy, Shavit, Luchangco, Spear) — https://www.elsevier.com/books/the-art-of-multiprocessor-programming/herlihy/978-0-12-397337-5
- Effective Go — https://go.dev/doc/effective_go
- Share Memory By Communicating — https://go.dev/blog/codelab-share
- Data Race Detector — https://go.dev/doc/articles/race_detector
- Concurrency Is Not Parallelism (Rob Pike) — https://go.dev/blog/waza-talk

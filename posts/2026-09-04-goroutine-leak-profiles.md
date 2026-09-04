---
title: "The Garbage Collector Can Prove Your Goroutine Leaked"
date: 2026-09-04
slug: goroutine-leak-profiles
summary: "Go 1.27 ships a goroutine leak profiler that stops guessing: it reuses the garbage collector's reachability machinery to prove, for goroutines blocked on channels and sync primitives, which ones can never be unblocked. What the article on go.dev announces, why the design is a decision procedure instead of another dashboard, and where the proof deliberately ends."
tags: golang, go1.27, concurrency, goroutines, pprof, garbage-collection, runtime, debugging, observability
---

[Goroutine Leak Profiles](https://go.dev/blog/goroutine-leak-profiles) is one of those Go blog posts where the headline feature is almost less interesting than the trick behind it. Go 1.27 introduces a profiler that finds goroutine leaks — and the detection mechanism is not a new monitoring heuristics. It is a proof, computed by the garbage collector, that a given goroutine can never run again.

That inverts how leak hunting has always worked, and it is worth understanding exactly why it works, where it stops working, and what it costs.

## Key insights

Six claims stand alone; the rest of the post is the argument behind them.

1. **A goroutine leak is garbage that is still scheduled.** Whether a blocked goroutine can ever run again is a reachability question: its blocking condition must be reachable from some goroutine that can still act on it. The garbage collector — not a monitoring heuristics — already computes the answer.
2. **The `goroutineleak` profile proves; it does not guess.** A goroutine profile hands you counts and stacks and asks you to infer a leak. The new profile computes an exact answer for the class it covers: leaked, or not leaked. No baseline, no trend analysis, no traffic normalization.
3. **The proof only covers blocking that memory can decide.** Goroutines permanently blocked on channels or on `sync` primitives (`Mutex`, `RWMutex`, `WaitGroup`, `Cond`) are within reach. I/O waits are not — their wake-up condition lives outside the heap — so they are never reported as leaked.
4. **Reachable-but-dead blockers are the built-in blind spot.** If a channel or lock stays reachable through a global variable or a merely runnable goroutine, waiters on it are never flagged, even when nothing will ever touch it again. The runtime cannot tell "reachable" from "will be used".
5. **Leaked goroutines never un-leak, so sample rarely.** A leak, once observed, stays observed for the lifetime of the process. Collecting the profile hourly instead of continuously costs almost nothing and misses almost nothing; alarm on a count above zero, not on a growing trend.
6. **Unit tests and production catch different halves of the problem.** `goleak` and `synctest` reproduce and catch leaks deterministically at test time; the goroutine leak profile catches what tests never exercised, in the running system. The winning setup layers both.

## The leak that the other tools miss

A goroutine leak is a goroutine that is blocked forever: the condition that would unblock it can never be met. A worker parked on an unbuffered channel send whose receiver has already returned, a mutex waiter whose lock holder broke out of its loop without unlocking, a `WaitGroup` waiter whose counter will never reach zero.

The cost is compounding. Each leaked goroutine pins its stack and everything it references, so memory grows without bound while the garbage collector pays to scan it — worse under [`GOMEMLIMIT`](https://go.dev/doc/gc-guide), where the GC starts thrashing against a hard ceiling. In production the symptom shows up as a slow memory creep and then an incident, and the diagnosis is usually painful.

The existing tooling covers two layers but not the third. [`goleak`](https://github.com/uber-go/goleak) instruments unit tests and flags any goroutine still alive when a test ends — excellent at the developer's desk. [`synctest`](https://go.dev/blog/synctest) (Go 1.25) makes concurrent tests deterministic so those leaks actually reproduce. But neither runs in production, and the production tool that *does* exist, the goroutine profile, only gives you counts and stacks. A microservice under load will show thousands of goroutines blocked in a channel receive *by design*; a leak of two goroutines per request fan-out is invisible against that baseline until the memory graph turns into a hockey stick. The classic workflow — diff goroutine dumps over time, read stacks, hypothesize — is human pattern-matching on top of sampling noise.

## Liveness is reachability in disguise

The Go 1.27 feature, announced on the Go blog by Vlad Saioc ([Goroutine Leak Profiles](https://go.dev/blog/goroutine-leak-profiles)), starts from an observation that sounds almost too simple: if a goroutine is blocked on a concurrency primitive that nothing else can reach, it is obviously leaked. Generalize that into an inductive definition of *liveness*:

- a goroutine is **live** if it is not blocked by a concurrency primitive, or
- if at least one concurrency primitive that blocks it is referenced by another live goroutine.

The base case is unblocked goroutines — they can always act. The inductive step assumes a live goroutine will eventually touch the primitives it references, unblocking whoever waits on them. Everything reachable from live goroutines, transitively, is live; everything blocked and unreachable is leaked.

Now notice what that definition is. "Reachable from a set of roots, transitively" is precisely what a garbage collector computes for memory. A goroutine leak, under this definition, is garbage that is still scheduled. The GC just needed its mark roots redefined: instead of rooting all goroutines (so they are never collected), root only the *unblocked* ones, mark what they reference, then iteratively add any blocked goroutine whose blocking primitive got marked, and repeat until the closure is stable. Whatever blocked goroutines remain are leaked. It is deadlock detection by conservation of reachability — the runtime already had the reachability engine; the paper behind this just pointed it at a different question.

That reframing is what makes this a *decision procedure* rather than another dashboard. Previous approaches required a human to look at counts and stacks and decide "these are growing, therefore suspicious." This feature computes, for the class of leaks it targets, an answer: leaked, or not leaked. No baseline, no traffic normalization, no trend analysis.

## The cut is principled, not arbitrary

The profiler only considers goroutines blocked on channels — send, receive, blocking `select` without a default case, even `select {}` and nil channels — and on the `sync` primitives `Mutex`, `RWMutex`, `WaitGroup`, and `Cond` (the [sync package](https://pkg.go.dev/sync)). Everything else — file and network I/O, system calls, hand-rolled spin locks — is never considered leaked.

This is the part of the design I find genuinely elegant, because the restriction is not a conservative v1 choice. It is the exact boundary where the definition stays decidable. For a goroutine blocked on I/O, unblocking depends on the state of the *world* — a socket, a file, another process — none of which lives in the heap. For a goroutine blocked on a channel or a mutex, unblocking depends only on memory: who else holds a reference to this primitive, and can that goroutine act on it. Memory-based liveness analysis is sound exactly when the wake-up condition is memory-based too. The cut follows the physics of the runtime.

And empirically it covers a lot: the article's examples are drawn from real production code — [CockroachDB](https://github.com/cockroachdb/cockroach/pull/584) (a missing unlock before `break`), [etcd](https://github.com/etcd-io/etcd/pull/6857) (a `Status` goroutine losing a race with `Stop`), [Kubernetes](https://github.com/kubernetes/kubernetes/pull/6632) and [Moby](https://github.com/moby/moby/pull/28462) (channels and mutexes deadlocking each other), [Moby](https://github.com/moby/moby/pull/25384) (a `Wait` inside the loop it should follow) — plus the double-send, early-return, and timeout patterns that show up in code review everywhere. The canonical mistake is an unbuffered channel where a sender outlives its receiver:

```go
type result struct {
	res workResult
	err error
}

func processWorkItems(ws []workItem) ([]workResult, error) {
	// Unbuffered: each send blocks until the main goroutine receives.
	ch := make(chan result)
	for _, w := range ws {
		go func() {
			res, err := processWorkItem(w)
			ch <- result{res, err} // blocks forever if we return early below
		}()
	}

	// Stop receiving at the first error...
	var results []workResult
	for range len(ws) {
		r := <-ch
		if r.err != nil {
			return nil, r.err // ...and every worker still sending leaks.
		}
		results = append(results, r.res)
	}
	return results, nil
}
```

Uber apparently found this exact shape in real services. In Go 1.27 you do not need to guess: the profile is exposed as the `goroutineleak` profile type in [runtime/pprof](https://pkg.go.dev/runtime/pprof), and if the [net/http/pprof](https://pkg.go.dev/net/http/pprof) handlers are already registered, the new `/debug/pprof/goroutineleak` endpoint appears automatically — no code change required. It is also part of the [Go 1.27 release notes](https://go.dev/doc/go1.27). `go tool pprof` will then show you precisely the leaked goroutines — the ones the collector proved unreachable. For the full runnable demo program in the article, the report looks like this:

```text
$ curl http://localhost:6060/debug/pprof/goroutineleak > leak.prof
$ go tool pprof leak.prof
Type: goroutineleak
(pprof) list processWorkItems
Total: 116
ROUTINE ======================== main.processWorkItems.func1 in .../main.go
         0        116 (flat, cum)   100% of Total
         .          .     31:           go func() {
         .          .     32:                   res, err := processWorkItem(w)
         .        116     33:                   ch <- result{res, err}
         .          .     34:           }()
```

In that run, 116 leaked goroutines, all parked on the channel send (line 33 of the demo program), all proven unreachable. The fix is the boring one — `ch := make(chan result, len(ws))` — which is the point. The profiler's job is not to be clever about fixes; it is to make the leak *visible* at a scale where eyeballing goroutine dumps no longer works.

## Where the proof ends

Because the mechanism is a proof, it has proof-shaped blind spots, and the article is refreshingly explicit about them.

**Reachability overreach.** If a blocking primitive stays reachable — through a global variable, or through a goroutine that is merely *runnable* — then goroutines blocked on it are never reported as leaked, even when nothing will ever touch that primitive again. A channel parked in a package-level variable with a worker blocked on it is, by the collector's definition, not garbage. This is a false-negative class that no GC-based scheme can close, because "reachable" and "will be used" are different predicates. The mitigation is architectural: stop letting concurrency primitives outlive their lifecycle through globals.

**Non-standard blocking.** Anything not blocked on a first-class channel or `sync` primitive is invisible. Network I/O waits, file reads, and user-space spin locks fall outside the proof. That matches the decidable-class argument above, but it means the feature has a specific shape: it finds *structural* concurrency leaks, not lifetime bugs where a goroutine waits forever on a resource that will never arrive.

**Non-determinism.** The collector reports leaks after they happen; it cannot predict them. Flaky, timing-dependent leaks still need the deterministic hammer — `synctest` — to reproduce in the first place. The article's own recommendation is to layer the approaches: leak profiles in production for detection, `goleak` and `synctest` in tests for diagnosis.

## What proving liveness costs

The detection rides on the GC's marking phase, and the price shows up in exactly one pathological shape, which the article calls the "daisy-chain": a runnable goroutine referencing a primitive that blocks G₁, which references a primitive blocking G₂, and so on. Proving Gᵢ₊₁ live requires first proving Gᵢ live, which serializes the marking rounds — and each round currently rescans all blocked goroutines, for a worst case of O(n²) per cycle where n is the goroutine count. The second cost is optimizable; the first is intrinsic to the analysis.

The practical mitigation is built into the semantics: once a goroutine is observed leaked, it stays leaked for the rest of the process. There is no point sampling this profile continuously. The article suggests periodic collection on the order of hours — the leak is monotonic, so sparse sampling loses almost nothing while keeping the GC cost amortized. That is a nice property and one the profile type itself encourages: `Type: goroutineleak` output that is a near-monotone counter of *proven* leaks, rather than a noisy count of blocked goroutines. In production I would alarm on "goroutineleak count > 0 over any window," which is a far more honest signal than any goroutine-profile growth rule I have ever written.

## From ASPLOS to the standard library

The feature has a research pedigree worth noting: it comes out of *[Dynamic Partial Deadlock Detection and Recovery via Garbage Collection](https://dl.acm.org/doi/pdf/10.1145/3676641.3715990)* (Saioc et al., ASPLOS 2025), a collaboration between Aarhus University, Washington University in St. Louis, and Uber, with Go team guidance from Michael Knyszek, Michael Pratt, and PJ Malloy. Deadlock detection via garbage collection is an old idea in the literature, but the interesting part of this transition is what changed between the paper and the release: the runtime's collector is concurrent and incremental now (the article casually drops that the marking phase is "now with the [Green Tea](https://go.dev/blog/greenteagc) variant"), and the implementation had to stay compatible with that machinery rather than bolt on a separate analysis. Research prototypes of runtime analyses usually die in the gap between "we can detect this offline" and "this must run inside a concurrent, low-pause collector without breaking it." Shipping it as a first-class profile type is a small miracle of persistence.

## What I take from this

Three things.

First, the feature converts a genre of debugging that felt like witchcraft — "is this goroutine dump growing because of traffic or because of a bug?" — into a computed property for a large, real class of failures. That is a bigger deal than the endpoint name suggests. The remaining work for the operator is only the part no runtime can do: keeping primitives' lifetimes aligned with the goroutines that use them.

Second, the boundary of the proof is the most instructive part of the design. Knowing *why* I/O-blocked goroutines cannot be covered — because their wake-up condition lives outside memory — is the difference between understanding the limitation and being surprised by it in an incident.

Third, "the runtime as a verifier" is a pattern with momentum. `synctest` gives tests control over time; this gives production a proof about blocked goroutines; the [race detector](https://go.dev/doc/articles/race_detector) long ago proved a whole bug class at runtime. Each one moves a question that used to require human judgment into machinery that either answers it or states exactly why it cannot.

If you write concurrent Go, read the [full post](https://go.dev/blog/goroutine-leak-profiles) — the [extra examples section](https://go.dev/blog/goroutine-leak-profiles#examples) alone is worth it, and the [Go playground reproductions](https://go.dev/play/p/S4Uw66sMbpj-) make the patterns click fast. Then add `goroutineleak` to your pprof collection rotation. The GC already knew. Now it can tell you.

## References

- Saioc, V. [Goroutine Leak Profiles](https://go.dev/blog/goroutine-leak-profiles) — The Go Blog, 2 September 2026. The source article for this post; the end sections catalogue additional leak patterns ([examples](https://go.dev/blog/goroutine-leak-profiles#examples)) and the full trade-offs ([limitations](https://go.dev/blog/goroutine-leak-profiles#limitations)).
- [Go 1.27 Release Notes](https://go.dev/doc/go1.27) — the "Goroutine leak profile" entry.
- [runtime/pprof](https://pkg.go.dev/runtime/pprof) and [net/http/pprof](https://pkg.go.dev/net/http/pprof) — package docs for the `goroutineleak` profile type and its HTTP endpoint.
- Saioc, V., et al. [Dynamic Partial Deadlock Detection and Recovery via Garbage Collection](https://dl.acm.org/doi/pdf/10.1145/3676641.3715990) — ASPLOS 2025 (ACM), the research the feature ships.
- Go team. [Go GC guide](https://go.dev/doc/gc-guide) — `GOMEMLIMIT` and soft memory limits.
- Go team. [Race detector](https://go.dev/doc/articles/race_detector) article.
- Go Blog. [Testing concurrent code with testing/synctest](https://go.dev/blog/synctest) and [Testing Time (and other asynchronicities)](https://go.dev/blog/testing-time) — the Go 1.25 testing groundwork.
- Knyszek, M., and Clements, A. [Green Tea: A new family of concurrent garbage collectors](https://go.dev/blog/greenteagc) — the concurrent GC, experimental in Go 1.25 and the default collector from Go 1.26 on.
- uber-go. [goleak](https://github.com/uber-go/goleak) — goroutine leak detection in unit tests.
- Real-world leaks used as examples in the source article: [CockroachDB #584](https://github.com/cockroachdb/cockroach/pull/584), [etcd #6857](https://github.com/etcd-io/etcd/pull/6857), [Kubernetes #6632](https://github.com/kubernetes/kubernetes/pull/6632), [Moby #25384](https://github.com/moby/moby/pull/25384), [Moby #28462](https://github.com/moby/moby/pull/28462).
- [Go playground reproduction](https://go.dev/play/p/S4Uw66sMbpj-) of the leak patterns from the article.

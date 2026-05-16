---
title: Codebases for Agentic Engineering
date: 2026-05-08
slug: codebases-in-the-era-of-agentic-software-engineering
summary: Agent-first codebases need explicit contracts, repo-local skills, modular packages, and safe parallel work patterns so multiple coding agents can ship without colliding.
tags: golang, agents, architecture
---

Codebases used to be written primarily for humans. The main readers were the teammates who opened files in an editor, learned the local conventions by trial and error, and built a mental map over weeks or months.

That assumption is breaking down.

In an agentic workflow, the first reader of a codebase is often a coding agent. The second reader may be another agent in a different worktree. The third may be a reviewer agent that only sees a diff. Humans still matter, but the codebase now has to explain itself to fast, literal, parallel workers that do not share much hidden context.

That changes what “well structured” means.

## A modern codebase needs explicit contracts

The strongest recent signal here is the emergence of repository-level instruction files such as [AGENTS.md in OpenAI Codex](https://developers.openai.com/codex/guides/agents-md) and the closely related `CLAUDE.md`, skills, and subagent patterns in [Claude Code best practices](https://code.claude.com/docs/en/best-practices.md), [skills](https://code.claude.com/docs/en/skills.md), and [worktrees](https://code.claude.com/docs/en/worktrees.md).

Those tools all point toward the same lesson: agents do better when the repo contains a compact, explicit contract for:

- what the repo is for,
- which commands are authoritative,
- which paths are safe to change,
- which invariants are non-negotiable,
- where deeper local instructions live.

Humans can absorb ambiguity. Agents mostly amplify it.

If a repo does not declare its rules, every agent run starts by rediscovering them. That wastes context, increases variance, and creates merge friction when multiple agents land changes that were individually reasonable but globally inconsistent.

## AGENTS.md should be a routing layer, not a novel

The worst version of `AGENTS.md` is a giant wall of text. The best version is a routing contract.

At the root, it should state the repo mission, the required workflow, the canonical test/build commands, and the directories where deeper instructions live. Then each major subtree can add a local instruction file with the context only that area needs.

That lets an agent read just enough to work safely instead of loading the entire history of the repository into every run.

For a Go codebase, that usually means:

1. a small root `AGENTS.md`,
2. local contracts for subsystems like `cmd/`, `internal/`, or `pkg/`,
3. repo-local skills for recurring workflows such as adding a handler, expanding a schema, or shipping a release.

The key idea is locality. The closer the instruction is to the code it governs, the easier it is for parallel agents to stay correct.

## Skills are how you turn tribal knowledge into executable guidance

Agent-first repositories should treat skills as first-class assets. A good skill is not motivational prose. It is an operational recipe with:

- purpose,
- inputs,
- exact steps,
- validations,
- failure modes.

That is useful for humans too, but it is especially valuable for agents because it removes guesswork from repeated tasks. Instead of hoping every agent rediscovers the right release flow, migration sequence, or API checklist, the repo can teach that behavior directly.

Skills are the scalable answer to “everyone knows how to do this.” In an agentic repo, that sentence should be treated as a bug report.

## Modularity matters more when multiple agents work in parallel

Parallel agents are most effective when they can work in separate git worktrees with minimal coordination. That only works if the codebase has solid seams.

In Go, the natural seam is the package boundary. If a package exports a small, well-tested interface contract, different agents can work on adjacent layers without constantly reaching across the boundary.

For example, an agent that owns orchestration code should not need to know how persistence is implemented. It should only need a stable interface:

```go
package contracts

import "context"

// SkillsCatalog defines the read-only contract for skill discovery.
type SkillsCatalog interface {
	Load(ctx context.Context, name string) (Skill, error)
	List(ctx context.Context) ([]Skill, error)
}

// WorktreeAllocator isolates parallel work into separate trees.
type WorktreeAllocator interface {
	Reserve(ctx context.Context, branch string) (Worktree, error)
	Release(ctx context.Context, path string) error
}

type Skill struct {
	Name        string
	Description string
}

type Worktree struct {
	Path   string
	Branch string
}
```

This looks simple, but that simplicity is the point. If the contract is small and explicit, one agent can change the allocator implementation while another extends skill discovery without both editing the same files.

## Agent-friendly Go packages should minimize hidden cross-package state

A lot of merge pain in agentic work happens because packages are not really modular. They look modular, but they share config globals, mutate common registries, or depend on side effects that are never written down.

A safer pattern is to make dependencies explicit in constructors:

```go
package planner

import (
	"context"
	"fmt"
)

// TaskStore persists the generated plan steps for later execution.
type TaskStore interface {
	SavePlan(ctx context.Context, id string, steps []string) error
}

type Service struct {
	store TaskStore
}

func New(store TaskStore) *Service {
	return &Service{store: store}
}

func (s *Service) Plan(ctx context.Context, id string, ask string) error {
	// Keep the planning stages explicit so parallel agents share the same flow.
	steps := []string{
		fmt.Sprintf("classify: %s", ask),
		"load local contracts",
		"select skill or subagent",
		"emit scoped plan",
	}
	return s.store.SavePlan(ctx, id, steps)
}
```

This does two things for agentic development:

1. it reduces the number of invisible assumptions,
2. it makes interface contracts testable in isolation.

That means a parallel agent can change a planner, store, or allocator behind the same interface and still merge cleanly.

## Worktrees are safer when interface tests are part of the contract

When agents work in separate git trees, smooth merging depends on more than good intentions. It depends on contract tests.

If package boundaries are meant to stay stable, the repo should enforce them with focused tests:

```go
package contracts_test

import (
	"context"
	"testing"
)

// fakeCatalog is a tiny stand-in that satisfies the contract in tests.
type fakeCatalog struct{}

func (fakeCatalog) Load(context.Context, string) (Skill, error) { return Skill{Name: "go-api"}, nil }
func (fakeCatalog) List(context.Context) ([]Skill, error)       { return []Skill{{Name: "go-api"}}, nil }

func TestCatalogContract(t *testing.T) {
	// Bind the fake to the interface so the consumer-facing seam stays explicit.
	var svc SkillsCatalog = fakeCatalog{}

	skill, err := svc.Load(context.Background(), "go-api")
	if err != nil {
		t.Fatalf("load skill: %v", err)
	}
	if skill.Name == "" {
		t.Fatal("expected skill name")
	}
}
```

The implementation here is tiny, but the idea scales: every important seam should have a small set of tests that define what consumers rely on. Parallel agents can change internals freely when those seams are protected.

## The codebase now needs to optimize for fast onboarding by machines

A human teammate can survive an opaque repo if they are patient and can ask questions. An agent gets a narrower window. It needs a fast path to useful context.

That suggests a different priority order than older codebases often used:

1. explicit contracts before clever abstractions,
2. local instructions before tribal knowledge,
3. interface stability before cross-package reach,
4. reproducible commands before “it works on my laptop.”

This is also why repo-local templates such as [agent-ready-repo](https://github.com/eugenelim/agent-ready-repo) are interesting. They encode the idea that architecture docs, operational skills, and agent-facing contracts belong inside the repo rather than floating around in chat history.

## The merge target is not just correctness, but convergence

The best agent-first codebases do more than help a single agent succeed. They help many agents converge on compatible answers.

That means designing the repo so independent workers can discover the same commands, the same invariants, the same subsystem boundaries, and the same review expectations. When that happens, separate worktrees stop feeling risky. They start feeling like throughput.

The old codebase question was: *can a human figure this out eventually?*

The new question is: *can multiple agents work in parallel, in separate trees, with enough shared contract to merge smoothly later?*

That is a different optimization target. It favors explicitness, modularity, skills, and durable interface contracts. Go is a strong fit for that world because its package boundaries, interfaces, tests, and deployment story make it easier to build systems that are boring in the right places.

And in the era of agentic software engineering, boring seams are a superpower.

## Sources

- [OpenAI Codex guide: AGENTS.md](https://developers.openai.com/codex/guides/agents-md)
- [Claude Code best practices](https://code.claude.com/docs/en/best-practices.md)
- [Claude Code skills](https://code.claude.com/docs/en/skills.md)
- [Claude Code worktrees](https://code.claude.com/docs/en/worktrees.md)
- [agent-ready-repo example](https://github.com/eugenelim/agent-ready-repo)

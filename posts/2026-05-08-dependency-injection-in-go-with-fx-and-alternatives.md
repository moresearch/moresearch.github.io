---
title: Dependency Injection in Go
date: 2026-05-08
slug: dependency-injection-in-go-with-fx-and-alternatives
summary: In Go, dependency injection is usually best when it stays explicit. Uber Fx becomes useful when the application graph and lifecycle are large enough to justify framework help, but it is not the only option.
tags: golang, architecture, dependency-injection
---

Dependency injection in Go is one of those topics where the community is often right for the wrong reason.

People say “just wire things manually,” and a lot of the time that is the correct answer. But the deeper point is not that frameworks are bad. It is that Go already gives you a simple, testable way to express dependencies: constructors, interfaces, and explicit initialization in `main`.

That means a DI framework has to earn its complexity.

When I look at the current Go ecosystem, I think the most useful way to frame the space is:

1. **manual wiring first,**
2. **Dig if you want a runtime container without a full app framework,**
3. **Fx if you want runtime wiring plus lifecycle orchestration,**
4. **Wire if you specifically want compile-time generation, with the caveat that it is now unmaintained.**

That ordering matches how I think about operational risk, not just developer preference.

## What dependency injection should mean in Go

In Go, dependency injection should mostly mean this: constructors receive the collaborators they need, and `main` decides how the graph gets assembled.

That keeps the code honest.

```go
package main

import "log"

type Config struct {
	DSN string
}

type DB struct {
	dsn string
}

// NewDB keeps database setup explicit at the application edge.
func NewDB(cfg Config) *DB {
	return &DB{dsn: cfg.DSN}
}

type UserService struct {
	db *DB
}

// NewUserService injects the database directly through the constructor.
func NewUserService(db *DB) *UserService {
	return &UserService{db: db}
}

func main() {
	cfg := Config{DSN: "postgres://app"}
	db := NewDB(cfg)
	svc := NewUserService(db)

	// Use the fully assembled service graph.
	log.Printf("service ready with %s", svc.db.dsn)
}
```

This style is boring, but it scales farther than people sometimes admit. It is obvious in code review, easy to test, and does not hide object creation behind reflection or generated files.

If the graph is still small enough to fit comfortably in one place, this is usually my favorite option.

## Where Uber Dig fits

[Uber Dig](https://github.com/uber-go/dig) is a runtime DI toolkit, not a full framework. Its own README is pretty clear about the intended scope: it is good for resolving the object graph during process startup and as a building block for a framework like Fx, but not as a user-facing service locator.

That distinction matters.

Dig is useful when you want container-driven wiring without buying into a larger application model. You provide constructors, then invoke a function whose parameters the container fills in.

```go
package main

import (
	"log"

	"go.uber.org/dig"
)

type Config struct {
	DSN string
}

type DB struct {
	dsn string
}

// NewDB builds the shared database dependency.
func NewDB(cfg Config) *DB {
	return &DB{dsn: cfg.DSN}
}

func main() {
	c := dig.New()

	// Register concrete constructors with the container.
	_ = c.Provide(func() Config { return Config{DSN: "postgres://app"} })
	_ = c.Provide(NewDB)

	// Ask Dig to resolve the object graph for this startup function.
	_ = c.Invoke(func(db *DB) {
		log.Printf("connected to %s", db.dsn)
	})
}
```

The upside is less manual wiring in `main`. The downside is that the dependency graph becomes more implicit. You read constructor signatures to understand the graph, but the assembly is no longer plain Go code in one obvious place.

That tradeoff can be fine, but I think it should be deliberate.

## Where Uber Fx becomes compelling

[Uber Fx](https://github.com/uber-go/fx) sits one level higher. It is built on Dig, but it is really an **application framework** for dependency injection plus lifecycle.

That lifecycle part is the reason to care.

Once your application has:

- HTTP servers,
- background workers,
- metrics/reporting,
- shutdown hooks,
- multiple modules owned by different teams,

plain constructor wiring stops being the whole problem. Now you also need deterministic startup ordering, clean shutdown, and a compositional way to express module boundaries.

That is where Fx earns its keep.

```go
package main

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"
)

type ServerParams struct {
	fx.In

	Lifecycle fx.Lifecycle
	Mux       *http.ServeMux
}

// NewMux provides the shared HTTP router for the app.
func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

// RegisterServer binds the HTTP server lifecycle to the Fx app.
func RegisterServer(p ServerParams) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: p.Mux,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(NewMux),
		fx.Invoke(RegisterServer),
	)

	// Run manages startup, signal handling, and shutdown hooks.
	app.Run()
	log.Println("stopped")
}
```

That example captures the real difference. Fx is not just about “inject this dependency for me.” It is about giving the whole application a structured runtime and lifecycle.

If I had a multi-module service with enough startup/shutdown machinery, Fx would be near the top of my list.

## The strongest alternative to Fx is often still manual wiring

This is the part that gets lost in some DI comparisons. The main alternative to Fx is not necessarily another DI library. It is often **explicit module wiring plus a few carefully owned lifecycle abstractions**.

For many Go services, that wins on:

- debuggability,
- ease of onboarding,
- grep-ability,
- and the ability to understand startup by reading one file.

Fx becomes attractive when those benefits are outweighed by graph size and lifecycle complexity. Until then, a framework can be solving a problem you do not actually have yet.

## What about Google Wire?

[Google Wire](https://github.com/google/wire) took a very different approach: compile-time code generation instead of runtime reflection.

That is conceptually appealing in Go because it keeps the final wiring as ordinary generated Go code, with no runtime container and no reflection overhead. It also fits the language's bias toward explicitness better than many DI frameworks do.

```go
//go:build wireinject

package main

import "github.com/google/wire"

type Config struct{}
type DB struct{}
type UserService struct{}

// NewDB constructs the database dependency from configuration.
func NewDB(Config) *DB { return &DB{} }

// NewUserService wires the service against the database dependency.
func NewUserService(*DB) *UserService { return &UserService{} }

// InitializeUserService tells Wire which providers belong in the graph.
func InitializeUserService() *UserService {
	wire.Build(NewDB, NewUserService)
	return nil
}
```

The problem is that Wire is now explicitly marked as **no longer maintained**. Its README says so directly and points users toward forks if they need updates.

That does not make the idea bad. In fact, I still think compile-time wiring is philosophically attractive in Go. But if I were starting something fresh today, I would treat Wire as a useful reference point, not my default foundation.

## My practical ranking

If I were choosing today for a production Go codebase, my rough decision tree would be:

1. **Manual constructors** if the graph is still easy to read in `main`.
2. **Fx** if startup/shutdown lifecycle and module composition are the real pain.
3. **Dig** if I want container-style runtime wiring without the full Fx application model.
4. **Wire** only with eyes open, because the project is no longer maintained.

The important part is not picking the most “advanced” tool. It is matching the tool to the shape of the graph and the operational complexity of the app.

## My reflection

Go dependency injection works best when it preserves the language's bias toward clarity.

That is why I like Fx more than many DI frameworks in other ecosystems: it is not trying to turn Go into annotation soup. It still relies on constructors and explicit parameter types. But I also think Fx is easiest to justify when lifecycle management is the real problem, not just constructor wiring.

If your application is still small, manual wiring is not primitive. It is often the most Go-like answer.

If your application has grown into a startup graph with real operational structure, Fx becomes much more convincing.

## Sources

- [Uber Fx README](https://github.com/uber-go/fx)
- [Uber Fx docs](https://uber-go.github.io/fx/)
- [Uber Dig README](https://github.com/uber-go/dig)
- [Google Wire README](https://github.com/google/wire)


Dependency injection is the engineering response to the problem of coupling. Every component depends on other components. The dependencies must be satisfied for the system to function. Manual wiring creates tight coupling — each component knows exactly which implementation it uses. DI inverts the dependency: the component declares what it needs, the injector provides it. The component doesn't know the implementation. The pattern is information hiding applied to object construction. Parnas argued that modules should hide design decisions from each other. DI hides the decision of which implementation to use. The principle is the same. The granularity is different.

---
title: Temporal vs DBOS for Go: two paths to durable execution
date: 2026-07-10
slug: temporal-vs-dbos-golang
summary: "Temporal gives you a battle-tested orchestration platform with event-sourced replay. DBOS gives you durable workflows in a single Postgres-backed binary. Which one fits your Go stack?"
tags: go, temporal, dbos, durable-execution, workflows, postgres
---

Durable execution — the guarantee that a workflow runs to completion even if the process crashes, the machine dies, or the network partitions — is becoming table stakes for backend systems. Two projects in the Go ecosystem take fundamentally different paths to the same destination: Temporal and DBOS.

Temporal is the incumbent. Separate infrastructure, event-sourced replay, a mature SDK that replaces Go's concurrency primitives with deterministic equivalents. DBOS is the challenger. PostgreSQL-native, checkpoint-based, a thin layer over your existing database that gives you durable workflows, queues, and scheduling without deploying anything new.

Choosing between them is not about which one is "better." It is about which architecture matches your operational constraints.

## Temporal: the platform

Temporal is a distributed orchestration engine. You deploy a Temporal server (or use Temporal Cloud), and your Go workers connect to it. Workflows are deterministic functions that use Temporal's SDK replacements for Go primitives: `workflow.Sleep()` instead of `time.Sleep()`, `workflow.Go()` instead of `go`, `workflow.Channel()` instead of `chan`, `workflow.Selector()` instead of `select`.

The determinism constraint is the defining tradeoff. Because Temporal replays workflow code against event history to rebuild state after a crash, your workflow code must produce the same sequence of decisions given the same history. This means no `time.Now()`, no random numbers, no direct database calls inside a workflow. Non-deterministic work goes into Activities — functions that Temporal calls outside the replay sandbox, with configurable retries and timeouts.

The Go SDK (v1.44.0 as of mid-2026) is mature. Here is a user onboarding workflow — charge a customer, provision services, send a welcome email — with configurable retries and timeouts:

```go
// Workflow: deterministic coordination logic
func OnboardingWorkflow(ctx workflow.Context, input OnboardingInput) error {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: 30 * time.Second,
        RetryPolicy: &temporal.RetryPolicy{
            InitialInterval:    time.Second,
            BackoffCoefficient: 2.0,
            MaximumAttempts:    3,
        },
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    // Step 1: Charge payment
    var chargeID string
    if err := workflow.ExecuteActivity(ctx, ChargeCustomer, input.CustomerID, input.Amount).Get(ctx, &chargeID); err != nil {
        return fmt.Errorf("charge failed: %w", err)
    }

    // Step 2: Provision resources
    var resourceIDs []string
    if err := workflow.ExecuteActivity(ctx, ProvisionResources, input.Plan, input.Region).Get(ctx, &resourceIDs); err != nil {
        // Compensation: refund if provisioning fails
        if refundErr := workflow.ExecuteActivity(ctx, RefundCharge, chargeID).Get(ctx, nil); refundErr != nil {
            return fmt.Errorf("provision failed, refund also failed: %w", refundErr)
        }
        return fmt.Errorf("provision failed, refunded: %w", err)
    }

    // Step 3: Send welcome email (fire-and-forget with short timeout)
    emailCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
        StartToCloseTimeout: 10 * time.Second,
        RetryPolicy:         &temporal.RetryPolicy{MaximumAttempts: 5},
    })
    if err := workflow.ExecuteActivity(emailCtx, SendWelcomeEmail, input.Email, resourceIDs).Get(ctx, nil); err != nil {
        workflow.GetLogger(ctx).Warn("welcome email failed, not fatal", "error", err)
    }

    return nil
}

// Activity: non-deterministic code — safe to call external APIs
func ChargeCustomer(ctx context.Context, customerID string, amount int) (string, error) {
    return billingClient.CreateCharge(ctx, customerID, amount)
}
```

Each `ExecuteActivity` call checkpoints progress in Temporal's event history. If the worker crashes after `ChargeCustomer` succeeds but before `ProvisionResources` starts, the workflow resumes from the charge — it does not double-charge. The compensation logic (refund on provision failure) is explicit in the workflow code, not buried in infrastructure config.

Recent SDK additions include Standalone Activities (durable job processing without a parent workflow), Worker Versioning for safe deploys of workflow code changes, and Nexus for cross-service orchestration. The tooling — `workflowcheck` for static analysis of determinism violations, replay tests for verifying workflow code against historical event histories — helps catch issues before they reach production.

Running Temporal means running a Temporal server. This is a non-trivial operational commitment. The server requires a database (MySQL or PostgreSQL), an Elasticsearch instance for visibility, and a multi-service deployment for production. Temporal Cloud exists, but it is a paid service with its own pricing model. For teams already running significant infrastructure, this is manageable. For small teams or single-service deployments, it is overhead that must be justified.

## DBOS: the embedded alternative

DBOS takes the opposite approach. There is no external server to deploy. You add `dbos-transact-golang` to your Go module, point it at a PostgreSQL database, and you have durable workflows. All state — inputs, outputs, step progress, sleep timers, queue positions, notifications — is checkpointed in Postgres. If your process crashes, on restart all workflows automatically resume from the last completed step.

The programming model is simpler than Temporal's because there is no replay. Workflows are regular Go functions. Non-deterministic operations go into steps via `dbos.RunAsStep()`. Here is the same user onboarding workflow in DBOS:

```go
func OnboardingWorkflow(ctx dbos.DBOSContext, input OnboardingInput) error {
    // Step 1: Charge payment (wrapped — output is checkpointed)
    chargeID, err := dbos.RunAsStep(ctx, func(ctx context.Context) (string, error) {
        return billingClient.CreateCharge(ctx, input.CustomerID, input.Amount)
    }, dbos.WithStepName("charge"), dbos.WithStepMaxRetries(3))
    if err != nil {
        return fmt.Errorf("charge failed: %w", err)
    }

    // Step 2: Provision resources
    resourceIDs, err := dbos.RunAsStep(ctx, func(ctx context.Context) ([]string, error) {
        return provisioningClient.CreateResources(ctx, input.Plan, input.Region)
    }, dbos.WithStepName("provision"), dbos.WithStepMaxRetries(3))
    if err != nil {
        // Compensation: refund
        dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
            return nil, billingClient.Refund(ctx, chargeID)
        }, dbos.WithStepMaxRetries(5))
        return fmt.Errorf("provision failed, refunded: %w", err)
    }

    // Step 3: Send welcome email (fire-and-forget, best-effort)
    dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
        return nil, emailClient.SendWelcome(ctx, input.Email, resourceIDs)
    }, dbos.WithStepName("email"), dbos.WithStepMaxRetries(5))

    return nil
}
```

Each `RunAsStep` checkpoint its return value to Postgres. If the process crashes after charging but before provisioning, the workflow resumes from the next step — the charge is never duplicated. The compensation path is explicit in application code. No separate activity registration, no SDK replacement for `time.Now()`, no sandbox. Just Go functions with step-wrapped I/O and Postgres as the durability layer.

This simplicity extends across the feature set. Durable queues give you concurrency control and rate limiting without a message broker:

```go
queue := dbos.NewWorkflowQueue(ctx, "task_queue",
    dbos.WithWorkerConcurrency(5),
    dbos.WithRateLimiter(&dbos.RateLimiter{
        Limit: 100, Period: 60 * time.Second,
    }))
```

Durable sleep persists wake-up time to Postgres — a workflow can `dbos.Sleep(ctx, 48*time.Hour)` and survive process restarts across those two days. Cron scheduling is a single option on workflow registration. Send/Recv provides durable notifications between workflows with exactly-once semantics. These are not add-ons. They are part of the same library.

The cost is scale. DBOS is bounded by a single Postgres instance (or cluster). Temporal scales horizontally across workers and partitions. For most applications this distinction is theoretical — a well-tuned Postgres instance handles millions of workflow steps — but for the largest deployments, Temporal's distributed architecture is the right call.

## Setup side by side

The difference in operational commitment is visible even at the setup stage.

**Temporal** requires worker registration, activity registration, and a running Temporal server. A minimal worker binary looks like this:

```go
func main() {
    c, _ := temporalclient.NewClient(temporalclient.Options{})
    defer c.Close()

    w := worker.New(c, "onboarding-queue", worker.Options{})
    w.RegisterWorkflow(OnboardingWorkflow)
    w.RegisterActivity(ChargeCustomer)
    w.RegisterActivity(ProvisionResources)
    w.RegisterActivity(RefundCharge)
    w.RegisterActivity(SendWelcomeEmail)

    if err := w.Run(worker.InterruptCh()); err != nil {
        log.Fatal(err)
    }
}
```

Workflows, activities, and the worker are three distinct concerns. You register each activity by name. Temporal's type system enforces the separation — `workflow.Context` for workflows, `context.Context` for activities — which prevents you from accidentally calling non-deterministic code inside a workflow.

**DBOS** collapses this into a single lifecycle block:

```go
func main() {
    ctx, _ := dbos.NewDBOSContext(context.Background(), dbos.Config{
        AppName:     "onboarding",
        DatabaseURL: os.Getenv("DBOS_SYSTEM_DATABASE_URL"),
    })

    dbos.RegisterWorkflow(ctx, OnboardingWorkflow)
    queue := dbos.NewWorkflowQueue(ctx, "onboarding-queue",
        dbos.WithWorkerConcurrency(10))

    if err := dbos.Launch(ctx); err != nil {
        log.Fatal(err)
    }
    defer dbos.Shutdown(ctx, 30*time.Second)
}
```

No separate activity registration. No external server URL. The database connection string is the only infrastructure dependency. The `Launch` call starts the worker loop, and `Shutdown` drains in-flight work gracefully.

The difference is not just lines of code. It is the number of things you need to have running before the code works. Temporal needs a server, a database, and your worker. DBOS needs a database and your binary. For a team evaluating durable execution for the first time, that gap determines whether they ship or shelve.

## The determinism constraint, concretely

The sharpest practical difference is how each system handles non-deterministic Go code inside a workflow. Consider a workflow that needs the current time:

**Temporal** — this fails:

```go
// INSIDE A WORKFLOW — WRONG
now := time.Now() // Non-deterministic! Replay will see a different value.
```

Temporal enforces this at multiple levels. The SDK documentation lists the allowed API surface inside workflows. The `workflowcheck` linter catches violations at build time. The replay test framework re-executes your workflow against recorded event histories and fails if the code path diverges. You must use `workflow.Now()` instead, which returns the timestamp from the event history rather than the system clock:

```go
// CORRECT inside a Temporal workflow
now := workflow.Now(ctx)
```

**DBOS** — this works fine:

```go
// Inside a DBOS workflow — fine, no replay ambiguity
now := time.Now()
```

Because DBOS does not replay workflow code — it checkpoints step *outputs* and resumes from the last checkpoint — `time.Now()` inside a workflow is harmless. The workflow runs once forward. If it crashes after step 3, it resumes by re-executing the workflow function from the top, but `dbos.RunAsStep` returns the cached output for steps 1–3 without re-executing their bodies. The non-deterministic code inside the step closure is never re-run.

The tradeoff is clear. Temporal guarantees that your workflow code produces the same decisions given the same history — a strong correctness property, paid for with a constrained programming model. DBOS guarantees that completed steps are never re-executed — a weaker property, but one that lets you write regular Go. If your workflows are straightforward chains of API calls with compensation on failure, DBOS's model is sufficient and less constraining. If your workflows contain complex branching, signals, or state machines where replay divergence would mean incorrect business outcomes, Temporal's enforcement is worth the ceremony.

## Four stories

Abstract comparisons only go so far. Here are four concrete scenarios — each drawn from real architectural choices — that show how the tradeoffs play out.

### Story 1: The two-person startup

Maya and Carlos are building a billing platform. Their stack is Go, Postgres, and a few Lambda functions. They need order processing to be durable — a customer signs up, a Stripe charge must succeed, a provisioning call to a third-party API must complete, and an invoice must be generated. If any step fails mid-flight, the whole thing must recover without double-charging.

They evaluate Temporal first. The programming model looks great. Then they read the deployment guide: Temporal server, MySQL or PostgreSQL, Elasticsearch for visibility, multi-service setup for production. They have two people and a managed Postgres instance. Running Temporal themselves is a non-starter, and Temporal Cloud adds a line item they cannot justify at this stage.

They try DBOS. `go get github.com/dbos-inc/dbos-transact-golang`, point it at their existing Postgres instance, and they have durable workflows by the end of the day. The order processing saga — three steps with compensation on failure — is 40 lines of Go. They ship within the week.

**What happened:** DBOS won because the operational cost of Temporal exceeded the value of its additional capabilities. For a small team, "just a library" is the right abstraction. The constraint they accepted — bounded by a single Postgres instance — is irrelevant at their scale. If they grow to need Temporal later, the workflow logic ports across because the concepts (steps, compensation, idempotency) are the same.

### Story 2: The enterprise polyglot migration

FinServCo is migrating a monolith to microservices. They have teams writing Go, Java, and Python. Their core business process — account opening — spans 14 steps across 6 services: identity verification (Go), fraud check (Java), credit check (Python), account creation (Go), card issuance (Java), and welcome kit dispatch (Python). The saga must handle partial failures with compensating transactions at each step. It must be visible in a single dashboard. It must survive any individual service going down.

They prototype the saga in DBOS on a single Go service. It works. Then they realize: every non-Go step needs an HTTP wrapper, every service needs to expose a compensation endpoint, and the DBOS workflow becomes a fragile orchestration hub that must coordinate everything over the network. The simplicity they gained by embedding the workflow engine they lose in the integration layer.

They switch to Temporal. Each team writes their step in their own language as a Temporal Activity. The Go service hosts the workflow, but the activities are polyglot — Temporal's SDK handles the cross-service communication. The Temporal UI gives them a single pane of glass across all 14 steps, with retry status, input/output inspection, and stack traces on failure. When the credit check service goes down during a deploy, in-flight account openings pause and resume automatically.

**What happened:** Temporal won because polyglot orchestration and visibility mattered more than operational simplicity. DBOS would have worked technically, but the integration tax of wrapping every non-Go service as an HTTP endpoint would have erased the simplicity advantage. Temporal's architecture — separate server, polyglot workers, unified visibility — matched the organizational structure of the problem.

### Story 3: The team already running Temporal

PlatformCo has run Temporal in production for two years. They have 40 workflow types, a dedicated infrastructure team managing the Temporal cluster, and engineers who know the SDK cold. But they are tired.

Tired of the upgrade cycles. Tired of the Elasticsearch index falling over during traffic spikes. Tired of explaining to new hires why `time.Now()` inside a workflow is a compile error. Tired of the determinism debugging sessions where someone accidentally closed over a map and the replay diverged three weeks later.

They start a new service — a simple notification pipeline (ingest event → enrich with user data → route to channel → record delivery). Four steps, no branching, no signals, no complex state machines. It does not need Temporal's full power. They build it in DBOS. The Temporal cluster stays for the complex 14-step account provisioning sagas. The new service ships faster, debugs easier, and has zero new infrastructure.

**What happened:** They did not migrate off Temporal. They stopped using it for problems that did not need it. The sweet spot for DBOS is the 80% of workflows that are linear chains of API calls with compensation. The sweet spot for Temporal is the 20% that involve branching, signals, child workflows, or cross-service coordination. The platforms coexist.

### Story 4: The 72-hour approval workflow

RegTech builds a compliance review system. A case is submitted. A human must approve or reject within 72 hours. If no response, the case auto-escalates. During those 72 hours, the workflow must survive process restarts, deploys, and database failovers. After approval, the case moves to archival. After rejection, it moves to remediation.

They build it in DBOS first:

```go
func ComplianceReview(ctx dbos.DBOSContext, caseID string) error {
    // Assign reviewer and notify
    reviewer, _ := dbos.RunAsStep(ctx, func(ctx context.Context) (string, error) {
        return assignReviewer(ctx, caseID)
    }, dbos.WithStepName("assign"))
    dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
        return nil, notifyReviewer(ctx, reviewer, caseID)
    }, dbos.WithStepName("notify"))

    // Wait for human approval — 72h timeout, survives everything
    result, err := dbos.Recv[string](ctx, "approval."+caseID, 72*time.Hour)
    if err != nil {
        dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
            return nil, escalate(ctx, caseID)
        })
        return nil
    }

    if result == "approved" {
        dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
            return nil, archive(ctx, caseID)
        })
    } else {
        dbos.RunAsStep(ctx, func(ctx context.Context) (any, error) {
            return nil, remediate(ctx, caseID)
        })
    }
    return nil
}
```

The `dbos.Recv` call pauses the workflow for up to 72 hours. The wake-up time is checkpointed in Postgres. If the server restarts during those three days, the workflow resumes with the timer intact. When the reviewer's HTTP endpoint calls `dbos.Send(ctx, workflowID, "approved", "approval."+caseID)`, the workflow wakes up and continues. No external timer service, no scheduled-job table, no cron polling.

They could build this in Temporal too — with signals and timers. The Temporal version would be equally durable and more scalable. But they would need to run a Temporal cluster for a workflow that fires a few hundred times a day and sleeps for three days each time. The infrastructure-to-business-logic ratio is backward.

**What happened:** DBOS won specifically because of the sleep pattern. `dbos.Sleep` and `dbos.Recv` turn "wait for something, survive anything" into a single function call persisted in Postgres. Temporal does this too, but the operational overhead is harder to amortize when the workflow is mostly waiting.

## Head-to-head

| Property | Temporal | DBOS |
|---|---|---|
| **Architecture** | External server + workers | Embedded library + Postgres |
| **Persistence model** | Event-sourced, replay-based | Checkpoint-based, step-level |
| **Determinism guarantee** | Enforced by SDK + static analysis | Convention, not enforced |
| **Go primitives** | Must use SDK replacements | Regular Go, wrap I/O in steps |
| **Deployment** | Temporal server + DB + ES | Your binary + Postgres |
| **Scaling** | Horizontal (workers + partitions) | Vertical (Postgres) + horizontal workers |
| **Queues** | Task queues built in | `NewWorkflowQueue` |
| **Scheduling** | Via client API / external cron | Built-in cron expressions |
| **Sleep** | `workflow.Sleep()` | `dbos.Sleep()` |
| **Resumability** | Automatic via event replay | Automatic via checkpoint |
| **Maturity** | Battle-tested since 2020 | GA since ~2024 |
| **Go SDK** | `go.temporal.io/sdk` v1.44 | `dbos-transact-golang` v0.18 |
| **License** | MIT | MIT |

## When to choose Temporal

**You already run Temporal.** If your organization has a Temporal deployment, the operational cost is sunk. Adding new workflows in Go is the path of least resistance.

**You need horizontal scale.** Temporal's worker pools and partitioned task queues handle throughput that would overwhelm a single Postgres instance. If you are building the control plane for a high-volume system, Temporal's architecture is designed for that.

**You want enforced determinism.** Temporal's SDK catches non-deterministic patterns at development time. The replay model means your workflow code is tested against real event histories. If correctness under failure is your primary concern, Temporal's constraints are features, not bugs.

**You are multi-language.** Temporal has mature SDKs for Go, Java, TypeScript, Python, and .NET. If your workflows span services in different languages, Temporal's polyglot support matters.

## When to choose DBOS

**You want zero new infrastructure.** If you already run Postgres, you already run DBOS. There is no server to deploy, no new database to provision, no additional service to monitor. The operational simplicity is the primary value proposition.

**Your team is small.** Three developers with a Postgres database should not be running a Temporal cluster. DBOS turns durable execution from an infrastructure problem into a library problem. This is the right abstraction level for most teams.

**You want deadline-driven sleep.** `dbos.Sleep(ctx, 48*time.Hour)` works across process restarts and redeploys. Temporal's `workflow.Sleep()` does the same, but the ergonomics of colocating your sleep logic with your workflow code in a single binary are simpler with DBOS.

**You are building a Go-native service.** DBOS feels like Go. No two-function split between workflow and activity. No SDK replacements for `time.Now()`. Just regular Go functions with step-wrapped I/O. The learning curve is a fraction of Temporal's.

## The deeper split

The choice between Temporal and DBOS reflects a broader tension in backend infrastructure: platform vs library. Temporal is a platform — separate, powerful, operationally significant, and extremely capable. DBOS is a library — embedded, simpler, and limited by the constraints of the database it sits on.

Neither is wrong. The platform approach wins when the problem is large enough to justify the operational cost. The library approach wins when the problem is important but the infrastructure overhead is not. Durable execution is too valuable to skip entirely. The question is whether you want to run a platform for it, or just import a library.

---

**References:**
- [Temporal Go SDK](https://pkg.go.dev/go.temporal.io/sdk) — `go.temporal.io/sdk` v1.44
- [DBOS Transact Golang](https://github.com/dbos-inc/dbos-transact-golang) — `dbos-transact-golang` v0.18
- [DBOS Go Workflow Tutorial](https://docs.dbos.dev/golang/tutorials/workflow-tutorial)
- [Temporal Standalone Activities](https://temporal.io/blog/standalone-activities-durable-job-processing-now-in-public-preview)

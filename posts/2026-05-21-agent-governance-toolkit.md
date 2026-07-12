---
title: Agent Governance Toolkit: deterministic, auditable governance for agent fleets
date: 2026-05-21
slug: agent-governance-toolkit
summary: "A practical look at Microsoft’s Agent Governance Toolkit (AGT): what it solves, core capabilities, and how to get started governing agent actions safely and audibly."
tags: [governance, agents, security, microsoft]
draft: true
---

TL;DR — Microsoft’s Agent Governance Toolkit (AGT) is a production-oriented stack for enforcing deterministic, sub-millisecond policy checks and audit trails on agent actions. It provides a policy engine (YAML / OPA / Cedar), identity & trust primitives, execution sandboxes, SRE primitives, tamper-evident audit logs, and adapters for major agent frameworks. If you’re running agents in anything beyond a toy demo, AGT is one of the most complete open-source starting points for making agent behavior safe, observable, and auditable.

Why this matters
---------------
Agentic systems change the threat model: “please follow the rules” is not reliable. AGT moves enforcement out of prompts and into the application layer — every tool call and resource access is evaluated against policy before execution. That means deterministic allow/deny decisions, consistent enforcement across frameworks, and an auditable trail that teams can use for compliance and post-mortem.

What AGT provides (high level)
-------------------------------
- Policy Engine: deterministic, low-latency allow/deny checks (sub-ms in benchmarks). Supports YAML policy files plus OPA/Rego and Cedar, and defaults to fail-closed.
- Zero-Trust Identity: agent credentials, delegation chains, and behavioral trust scoring so identity and delegation limits are part of enforcement decisions.
- Execution Sandboxing: privilege rings, saga orchestration for multi-step flows, and an agent “kill switch” to stop runaway behavior.
- Agent SRE: SLOs, error budgets, replay debugging, chaos/testing primitives, and OpenTelemetry-native governance events.
- Audit & Compliance: Merkle-chained, tamper-evident audit logs and Decision BOMs; mappings for EU AI Act, SOC 2, HIPAA, GDPR.
- MCP Security Gateway & PromptDefense: tool-poisoning detection, description-drift checks, and multi-vector prompt-injection auditing.
- Framework Integrations: adapters/middleware for LangChain, AutoGen, OpenAI Agents, Semantic Kernel and more — plus cross-language SDKs (Python, TypeScript, .NET, Go, Rust).

Practical example (Python style)
-------------------------------
This is the idiomatic quick usage pattern from the project — an evaluator that denies dangerous tools:

```python
from agent_os.policies import PolicyEvaluator, PolicyDocument, PolicyRule, PolicyCondition, PolicyOperator, PolicyAction, PolicyDefaults

evaluator = PolicyEvaluator(policies=[
  PolicyDocument(
    name="my-policy",
    version="1.0",
    defaults=PolicyDefaults(action=PolicyAction.ALLOW),
    rules=[
      PolicyRule(
        name="block-dangerous-tools",
        condition=PolicyCondition(field="tool_name", operator=PolicyOperator.IN, value=["execute_code","delete_file"]),
        action=PolicyAction.DENY,
        priority=100,
      ),
    ],
  )
])

evaluator.evaluate({"tool_name": "web_search"})   # allowed
evaluator.evaluate({"tool_name": "delete_file"})  # denied
```

Operational checks & CI
-----------------------
AGT ships CLI tools (`agt doctor`, `agt verify`, `agt red-team scan`, `agt lint-policy`) to test policies, run evidence checks, and fail CI on weak evidence. The repo also contains 900+ conformance tests and formal RFC-style specifications for each major component — a good sign for teams that need provable behaviors and repeatable test suites.

When to consider AGT
--------------------
- You operate agent-driven automation that performs non-trivial actions (API calls, code execution, data deletion).
- Compliance, auditability, or SRE-grade observability is required.
- You need consistent enforcement across different agent frameworks or languages.
- You want a spec-driven implementation to anchor testing and long-term maintenance.

Caveats and realities
---------------------
- AGT is ambitious and scoped for production; it’s labeled “Public Preview” and may evolve before GA.
- Integration work is required to adapt your existing agents and tool definitions; expect effort in mapping actions/contexts to policy fields.
- Policy design is its own discipline — start with conservative, fail-closed defaults and iterate with evidence.

Resources & next steps
---------------------
- Repo: https://github.com/microsoft/agent-governance-toolkit  
- Docs / Quickstart: https://microsoft.github.io/agent-governance-toolkit  
- Try: pip install agent-governance-toolkit[full] and run `agt doctor` / `agt verify` on your policy files  
- Read the specs and conformance tests in the repo to understand the behavioral contracts AGT enforces before integrating.

Bottom line
-----------
AGT isn’t a lightweight policy library — it’s an operational stack for teams that need to run agent fleets under strict, auditable controls. For product teams and security engineers, it’s a valuable reference implementation and a practical toolkit to lift agent governance from ad-hoc prompts into reproducible, testable infrastructure.



Engineering is always embedded in a context of constraints — economic, political, organizational. The engineer who ignores the context builds systems that are technically correct and operationally irrelevant. The engineer who understands the context builds systems that survive. The context is part of the specification. The specification is incomplete without it.

# SentinelPay — Agentic Payment Incident Commander

> An agentic payment-operations system that investigates payment failures, determines recovery strategies, and executes only policy-approved actions.

## Overview

Payment failures are rarely just a single error message.

An incident may involve payment state, account state, currency context, invalid parameters, transient conditions, or downstream failures. Operations teams often need to correlate multiple pieces of information before deciding what to do next.

**SentinelPay** is an agentic payment-operations prototype designed to investigate these incidents and determine the safest next action.

The core principle is:

> **Let agents investigate and reason quickly, but make financial actions explicit, bounded, auditable, and verifiable.**

SentinelPay does not give an AI agent unrestricted control over money movement.

Instead, it uses **bounded autonomy**:

* the agent investigates independently,
* decisions are evaluated against explicit policies,
* sensitive actions can require human approval,
* automated actions have verification requirements,
* ambiguous incidents are escalated instead of guessed.

## Architecture

```text
Payment Incident
       |
       v
Incident Detection
       |
       v
Agent Investigation
       |
       +------> Payment Context
       |
       +------> Account Context
       |
       +------> Transaction Context
       |
       v
Root-Cause Analysis
       |
       v
Recovery Decision
       |
       v
Policy & Risk Check
       |
       +------------------+
       |                  |
       v                  v
Safe Action         Human Approval
       |                  |
       +--------+---------+
                |
                v
          Airwallex API
                |
                v
            Verification
                |
          +-----+-----+
          |           |
          v           v
       Resolved    Escalated
```

## Agent Loop

SentinelPay follows an explicit:

**Observe → Reason → Act → Verify**

workflow.

### Observe

The agent gathers relevant payment and account information.

### Reason

The agent correlates the available evidence and determines the most likely failure condition.

### Act

The agent selects a recovery strategy.

The action is not executed directly. It first passes through the policy layer.

### Verify

After an action, SentinelPay checks the resulting state against the expected post-condition.

If the expected state is not reached, the incident is escalated rather than repeatedly retried.

## Safety Model

SentinelPay separates **investigation autonomy** from **action authority**.

| Capability                 | Agent             |
| -------------------------- | ----------------- |
| Gather evidence            | Allowed           |
| Analyze failure            | Allowed           |
| Recommend recovery         | Allowed           |
| Evaluate policy            | Constrained       |
| Execute sensitive action   | Policy controlled |
| Bypass approval            | Not allowed       |
| Ignore failed verification | Not allowed       |
| Continue indefinitely      | Not allowed       |

This prevents probabilistic agent reasoning from becoming unrestricted financial authority.

## Incident State

Incidents use explicit state transitions:

```text
DETECTED
   ↓
INVESTIGATING
   ↓
DECISION
   ↓
┌─────────────────────┐
│                     │
v                     v
EXECUTING       AWAITING_APPROVAL
│                     │
v                     v
VERIFYING        HUMAN_DECISION
│
├──> RESOLVED
│
└──> ESCALATED
```

Explicit state makes the workflow easier to reason about, test, audit, and recover.

## Evaluation

SentinelPay is designed to evaluate **behavioral correctness**, not only API success.

Example scenarios include:

* recoverable transient failures
* non-recoverable failures
* insufficient context
* conflicting transaction state
* repeated failures
* low-confidence decisions
* actions requiring approval
* successful recovery followed by verification
* failed recovery requiring escalation

Each scenario can define an expected outcome.

The evaluator checks whether the agent:

1. identifies the correct incident state,
2. selects an appropriate action,
3. respects authorization boundaries,
4. avoids unsafe repeated actions, and
5. verifies the resulting state.

## Repository Structure

```text
cmd/
  server/              Application entry point

internal/
  incident/            Incident domain and state management
  agent/               Investigation and decision logic
  policy/              Action authorization and risk rules
  airwallex/           Airwallex integration abstraction
  audit/               Audit events and decision history
  evaluation/          Scenario-based behavioral evaluation

pkg/
  api/                 Shared API response types

scenarios/             Deterministic incident scenarios

docs/
  architecture.md      System architecture
  agent-workflow.md    Agent execution workflow

tests/                 Integration and evaluation tests
```

## Current Status

SentinelPay is currently in the **prototype / architecture phase**.

Current work focuses on:

* incident domain modeling
* explicit incident state transitions
* agent investigation interfaces
* bounded action policies
* Airwallex integration abstraction
* auditability
* deterministic scenario evaluation

The implementation will progressively replace mock components with working integrations during the build phase.

## Development

### Requirements

* Go
* Git

### Run

```bash
go run ./cmd/server
```

The development server starts on:

```text
http://localhost:8080
```

Health check:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{
  "status": "ok",
  "service": "sentinelpay"
}
```

## Roadmap

### Phase 1 — Foundation

* [x] Repository structure
* [x] Go backend skeleton
* [x] Incident domain model
* [x] Agent interface
* [x] Policy abstraction
* [x] Airwallex client abstraction
* [x] Audit abstraction
* [x] Scenario definitions

### Phase 2 — Agent Workflow

* [ ] Incident creation API
* [ ] Investigation workflow
* [ ] Root-cause classification
* [ ] Recovery decision engine
* [ ] Explicit state machine
* [ ] Policy evaluation
* [ ] Approval workflow

### Phase 3 — Airwallex Integration

* [ ] Authentication
* [ ] Payment retrieval
* [ ] Account/context retrieval
* [ ] Supported recovery actions
* [ ] API error handling
* [ ] Post-action verification

### Phase 4 — Evaluation

* [ ] Deterministic scenario runner
* [ ] State-transition assertions
* [ ] Policy violation tests
* [ ] Repeated-action protection
* [ ] Verification failure scenarios
* [ ] Evaluation report

### Phase 5 — Demonstration

* [ ] Operations dashboard
* [ ] Incident timeline
* [ ] Agent reasoning/evidence view
* [ ] Approval interface
* [ ] Action audit trail
* [ ] End-to-end demo

## Design Principle

SentinelPay is built around one principle:

> **Agents should be able to reason about financial incidents without being given unrestricted authority over financial actions.**

The system therefore treats **policy, verification, auditability, and escalation as first-class components**, rather than adding them after the agent has already been built.

## Hackathon

SentinelPay is being developed for the **Agentic Banking Hackathon**, hosted by Airwallex.

The prototype is intended to demonstrate how agentic systems can be applied to payment operations while maintaining explicit safety and authorization boundaries.

# Milestone 7 — Release readiness

**Status:** Not started

## Goal (from proposal)

Tests, authorization checks, logging, deployment and operational documentation.

**Completion check:** Core workflows pass tests and the app can be deployed repeatably.

## Tasks (issue-sized)

1. **Harden auth from Milestones 2/5/6** — if earlier milestones used a
   minimal/bootstrap auth mechanism, this is where it gets hardened for real
   use (password hashing strength, session/token expiry, revocation) rather
   than deferred indefinitely.
2. **Authorization test sweep** — explicit tests proving: a vendor cannot
   touch another vendor's menu/orders; a cashier cannot perform
   vendor/manager actions; a guest session cannot act on another guest's
   dining session. The proposal calls these out by name — this milestone is
   where they get collected and verified systematically, not just
   incidentally covered per-feature.
3. **Unit tests for business rules** — vendor order splitting, accepted-total
   computation, commission/ledger math. These are the places correctness
   bugs have the highest cost (wrong money, wrong orders).
4. **Integration tests** — Postgres-backed repository/transaction behavior;
   migration tests on a clean DB and on a DB with prior migrations applied.
5. **End-to-end check** — a customer ordering from multiple vendors and
   receiving independently-tracked deliveries, exercised as one flow.
6. **Logging** — structured logging across handlers (already have zerolog
   wired from Milestone 1 — extend coverage rather than introduce a new
   approach).
7. **Deployment documentation** — how to run this repeatably outside a local
   machine: env vars required, migration step, build step.

## Completion check

Core workflows (checkout → split → accept/reject → pay → fulfill →
commission/ledger) pass tests, and the app can be deployed from a clean
environment using only the documented steps.

# Forkly — Scope & Roadmap

Working breakdown of the project proposal
(`/Users/kyawzinthant/Desktop/Food_Court_Ordering_System_Project_Proposal.pdf`)
into milestones. One file per milestone. Each file lists: goal, tables/endpoints
involved, concrete tasks (issue-sized), completion check, and open questions to
settle before/while building.

Columns/constraints/indexes are finalized per-feature, not here — this is scope,
not schema design.

## Milestones

| # | Milestone | Status |
|---|---|---|
| 1 | [Foundation](01-foundation.md) | Done |
| 2 | [Vendors and menus](02-vendors-and-menus.md) | In progress |
| 3 | [Tables and sessions](03-tables-and-sessions.md) | Not started |
| 4 | [Ordering](04-ordering.md) | Not started |
| 5 | [Counter and fulfillment](05-counter-and-fulfillment.md) | Not started |
| 6 | [Finance](06-finance.md) | Not started |
| 7 | [Release readiness](07-release-readiness.md) | Not started |

## Ground rules (from the proposal, apply across all milestones)

- **Backend-calculated totals, always.** Never trust a client-sent price.
- **Vendor data isolation.** A vendor must never read/write another vendor's
  menu or orders — every vendor-scoped query filters by vendor id derived from
  the authenticated identity, not from a request parameter.
- **Idempotent checkout.** Duplicate submissions must not create duplicate orders.
- **Traceable money.** Every manual payout/ledger entry has an actor + reference.
- **Migrations are additive.** Never rewrite an applied migration; add a new one.
- **Single food court for MVP.** Don't design for multi-food-court until it's asked for.

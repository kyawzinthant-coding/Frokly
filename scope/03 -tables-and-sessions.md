# Milestone 3 — Tables and sessions

**Status:** Not started

## Goal (from proposal)

QR table identification and dining-session lifecycle.

**Completion check:** Guest can access the correct active table session.

## Tables involved

- `food_courts` — likely already seeded (single row) in Milestone 2.
- `tables` — physical tables in the food court, each identifiable via a QR
  code (a stable table id/token encoded in the QR, not the raw DB id).
- `dining_sessions` — one active session per table at a time; tracks when it
  opened, and (later, Milestone 5) when it's closed after settlement.
- `guest_sessions` — ties an anonymous guest's browser (e.g. a session cookie
  or signed token handed back after the QR scan) to a `dining_sessions` row,
  so a guest doesn't need an account.

## Tasks (issue-sized)

1. **Schema migration** — `tables`, `dining_sessions`, `guest_sessions`.
2. **QR → table resolution endpoint** — given the table's QR token, resolve
   the table and either return the existing active dining session or create
   one. Idempotent: scanning twice doesn't open two sessions.
3. **Guest identity handling** — issue whatever token/cookie represents "this
   guest, in this dining session" and require it on subsequent guest-facing
   calls (cart, checkout, order status) added in Milestone 4+.
4. **Manager-facing table/session visibility** (minimal) — list tables and
   whether each currently has an active session. Full manager tooling is
   Milestone 6/7 territory; just enough here to verify the lifecycle works.

## Open questions to settle before/while building

- What identifies a table in the QR code — a random opaque token per table
  (recommended, avoids exposing sequential ids) or the table's primary key
  directly?
- Where does "guest identity" live — signed cookie, bearer token returned to
  the frontend and stored client-side, or both? Whatever's chosen here is
  what Milestone 4's cart/checkout endpoints will require as auth context.
- Does scanning an already-active table's QR from a second device join the
  same session (multiple phones, one table) or start a conflict? Proposal
  implies one active session per table, shared — worth deciding explicitly
  rather than defaulting silently.

## Completion check

Guest can access the correct active table session. A second scan of the same
table's QR resumes the same session rather than creating a duplicate.

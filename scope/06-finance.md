# Milestone 6 — Finance

**Status:** Not started

## Goal (from proposal)

Commission, ledger, withdrawals, manual payouts, manager views.

**Completion check:** Vendor earnings and payouts are traceable.

## Tables involved

- `commission_rules` — configurable percentage commission, likely per vendor
  (or a food-court-wide default with optional per-vendor override).
- `ledger_entries` — append-only financial record: a sale credits the
  vendor's earnings (sale amount minus commission), commission credits the
  platform.
- `withdrawal_requests` — vendor requests a withdrawal of available earnings.
- `payouts` — manager-approved manual payout against a withdrawal request,
  recorded with an actor + reference (per the proposal's explicit risk
  mitigation for payment/payout discrepancies).

## Tasks (issue-sized)

1. **Schema migration** — `commission_rules`, `ledger_entries`,
   `withdrawal_requests`, `payouts`.
2. **Commission configuration (manager role)** — manager sets/updates a
   vendor's commission percentage. First real use of a manager role —
   decide the role model here consistent with however cashier was handled
   in Milestone 5.
3. **Ledger posting on payment** — when a `payments` row is recorded
   (Milestone 5), post the corresponding `ledger_entries` (vendor earning,
   platform commission) in the same transaction as the payment. This is
   where "payment recording" and "finance" connect — the ledger must never
   be computed retroactively from scratch, it's posted at the moment of truth.
4. **Vendor earnings view (vendor role)** — vendor sees its own running
   earnings from the ledger.
5. **Withdrawal request + manual payout (vendor requests, manager approves)**
   — vendor submits a withdrawal request; manager approves/rejects and
   records a manual payout with actor + reference; payout reduces the
   vendor's available (unwithdrawn) balance.

## Open questions to settle before/while building

- Is "available balance" computed live by summing `ledger_entries` minus
  `payouts` each time, or maintained as a running balance column updated
  transactionally? Summing live is simpler and avoids drift bugs — default
  to that unless query volume becomes a real concern.
- Commission: percentage stored as e.g. `1500` (basis points) or a decimal?
  Given the project's existing convention of money as int64 cents, basis
  points (integer) is the consistent choice — avoid floats for commission
  math same as money.

## Completion check

A vendor's earnings, platform commission taken, and any payouts made to that
vendor can all be reconstructed from the ledger with a clear actor + reference
on every payout.

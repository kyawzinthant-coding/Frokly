# Milestone 4 — Ordering

**Status:** Not started

## Goal (from proposal)

Combined cart, checkout, vendor order splitting, status tracking.

**Completion check:** A single checkout produces independent vendor orders correctly.

## Tables involved

- `customer_orders` — one row per checkout (the guest-facing "order").
- `vendor_orders` — one row per vendor represented in that checkout; carries
  its own status (pending/accepted/rejected/...).
- `vendor_order_items` — line items of a `vendor_order`, each pointing at a
  `product` with the price captured **at checkout time** (never re-read live
  from `products` later — prices can change after the fact).
- `order_status_history` — append-only log of status transitions per
  `vendor_order`, for traceability.

## Tasks (issue-sized)

1. **Schema migration** — `customer_orders`, `vendor_orders`,
   `vendor_order_items`, `order_status_history`.
2. **Checkout endpoint** — guest (via the dining-session identity from
   Milestone 3) submits a multi-vendor cart. Backend re-validates every
   price and availability server-side, groups items by vendor, creates one
   `customer_orders` row + one `vendor_orders` row per vendor, in a single
   DB transaction. Idempotent — a retried/duplicate submission must not
   create a second order (idempotency key, or dedupe on a short window).
3. **Vendor order splitting logic** — the core business rule: given a cart
   payload, partition into per-vendor groups and snapshot prices. Pure
   enough logic that it should be unit-tested directly, independent of HTTP.
4. **Status tracking endpoints** — guest-facing: view current order(s) and
   each vendor order's status. Vendor-facing: list own vendor orders. No
   state-transition actions yet (accept/reject/prepare/deliver) — that's
   Milestone 5; this milestone just needs the split + read-back to work.

## Open questions to settle before/while building

- Is "pending" the initial `vendor_orders` status, with accept/reject arriving
  in Milestone 5, or does this milestone need a stub transition just to prove
  the split works end-to-end? Lean toward: this milestone creates orders in
  `pending` and stops there — Milestone 5 owns all transitions out of pending.
- How is checkout idempotency implemented — client-generated idempotency key
  header, or server-side dedupe on (session, cart-hash, short time window)?
  The proposal explicitly calls out "idempotent checkout" as a top risk, so
  this needs a deliberate answer, not an afterthought.

## Completion check

A checkout containing products from 2+ vendors produces exactly one
`customer_orders` row and one `vendor_orders` row per vendor, with correct,
backend-computed totals, and resubmitting the same checkout doesn't duplicate it.

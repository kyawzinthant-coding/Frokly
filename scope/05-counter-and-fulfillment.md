# Milestone 5 — Counter and fulfillment

**Status:** Not started

## Goal (from proposal)

Vendor response deadline, payment recording, preparation and delivery states.

**Completion check:** Accepted totals can be paid and vendor orders tracked through delivery.

## Tables involved

- `vendor_orders` (extends Milestone 4) — gains real status transitions:
  pending → accepted/rejected → (if accepted) preparing → ready → delivered.
- `order_status_history` — now actually populated on every transition.
- `payments` — one row per recorded counter payment against a `customer_orders`
  row's accepted total.
- `payment_events` — audit trail of payment-related actions (e.g. who
  recorded it, when).

## Tasks (issue-sized)

1. **Schema migration** — `payments`, `payment_events`; any columns
   `vendor_orders` needs for the response deadline (e.g. `responded_by`,
   `response_deadline_at`).
2. **10-minute vendor response window** — vendor accepts/rejects within the
   window; a timeout auto-resolves to rejected. This needs a time-based
   mechanism (checked lazily on read, or an actual scheduled sweep) — decide
   which, since "lazily on read" is simpler and avoids adding a
   background-worker dependency this early.
3. **Accepted-total computation** — once all vendor orders on a
   `customer_orders` are resolved (accepted/rejected/timed-out), compute the
   payable total from only the accepted portions.
4. **Payment recording endpoint (cashier role)** — cashier confirms payment
   received for the accepted total; requires cashier-role auth (first real
   use of a role beyond vendor staff — decide now whether this reuses the
   vendor-staff identity table with a role column, or is a separate minimal
   table, consistent with how Milestone 2 structured vendor auth).
5. **Preparation/ready/delivered status endpoints (vendor role)** — vendor
   moves its own accepted `vendor_orders` through preparing → ready →
   delivered. Vendor-scoped the same way menu management is.

## Open questions to settle before/while building

- Auto-reject on timeout: computed on next read (no payment/status check
  proceeds past a stale pending order past its deadline) vs. an actual
  background sweep updating rows proactively. Lazy-on-read is simpler and
  sufficient for MVP — note that reasoning here so it's not re-litigated later.
- Does a cashier see one unified "pay this total" view across vendors for a
  checkout, or per-vendor? Proposal implies one combined payment per
  checkout's accepted total, not per vendor — confirm before building the
  cashier UI.

## Completion check

Once a checkout's vendor responses are all resolved, the correct accepted
total can be recorded as paid by a cashier, and each accepted vendor order
can be independently advanced through preparing/ready/delivered by its vendor.

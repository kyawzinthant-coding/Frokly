# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

- **Vendor staff** (primary for now): people working a single food stall inside a food court. They
  sign in, see incoming orders for their stall only, accept or reject them within a 10-minute
  window, and move accepted orders through preparing → ready → delivered. They work during busy
  service, often with wet or busy hands, glancing at the screen between cooking and serving.
- **Cashier**: collects payment at the counter for accepted orders.
- **Manager**: runs the food court — vendors, tables, commission, payouts.
- **Customers** order from their table by scanning a QR code; they are guests with no account.
  Customer screens are a later surface.

## Product Purpose

A QR-based, multi-vendor food court ordering system. One customer checkout is split into one
order per vendor; each vendor handles only its own orders. Success for the vendor UI: staff can
see what to cook next and act on it in one tap, without training.

## Operating Context

- Staff use phones, tablets and laptops — every screen must work on all three.
- Noisy, busy counter environment; glanced at, not studied.
- Payment is collected at the counter (no payment gateway in MVP).
- Internal tool: staff are signed in with a vendor staff username and password.

## Capabilities and Constraints

- Built today: vendor staff login (HttpOnly cookie session) and a "who am I" profile
  (vendor name, staff username). Logout is not built yet (button exists, no action).
- Planned: vendor orders, menus/products, tables & dining sessions, cashier payment,
  manager finance (commission, ledger, payouts).
- Strict vendor isolation: a vendor never sees another vendor's orders or menu.
- English only for now. Thai/Burmese may come later — don't paint the type system into a corner.
- Stack: Next.js App Router, Tailwind v4, shadcn/ui (base-nova), TanStack Query with
  Orval-generated hooks from the OpenAPI contract.

## Brand Commitments

Name: **Forkly**. No logo, colors or typography committed yet.

## Evidence on Hand

No real vendor data, photos, menus or testimonials exist. Seed data is one vendor
("Test Vendor", staff "vendor1"). Do not fabricate orders, stats or metrics on screens.

## Product Principles

1. **One glance, one tap.** The next action is obvious and reachable without hunting.
2. **Only your stall.** Every vendor screen is scoped to the signed-in vendor, and says so.
3. **Honest screens.** Show real data or a clear empty state — never placeholder numbers.
4. **Calm under rush.** Nothing on screen competes with the orders for attention.

## Accessibility & Inclusion

Large touch targets, high contrast readable at arm's length, visible keyboard focus,
no information carried by color alone (order statuses also need text).

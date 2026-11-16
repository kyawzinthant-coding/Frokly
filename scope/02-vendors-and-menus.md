# Milestone 2 — Vendors and menus

**Status:** In progress

## Goal (from proposal)

Vendor/menu retrieval, vendor authentication, menu management.

**Completion check:** Vendor can manage products; customer can browse available menus.

## Tables involved

- `vendors` — one row per independent food vendor in the food court.
- `menu_categories` — groups products within a vendor's menu (e.g. "Drinks").
- `products` — name, price (int64 cents), category, vendor id, availability flag.
- Identity for vendor staff: a minimal `users` (or `vendor_staff`) table +
  something to tie a logged-in identity to a `vendor_id`. Keep this minimal —
  full `roles`/`user_roles` from the proposal's identity domain can wait until
  a milestone actually needs multiple roles sharing one `users` table (cashier,
  manager). For now a vendor-staff-only auth table is enough.

## Tasks (issue-sized)

1. **Vendor & menu schema migration**
   Goose migration for `vendors`, `menu_categories`, `products`, plus whatever
   minimal table backs vendor-staff login (see above). Money as integer cents.

2. **Vendor authentication**
   Vendor staff login → session/token that carries a vendor id. This is the
   only identity mechanism needed this milestone; don't build a generic
   multi-role auth system yet — that's scope creep ahead of when it's needed.

3. **Vendor menu management (authenticated, vendor-scoped)**
   CRUD on own products/categories, availability toggle. Every query filters
   by the vendor id derived from the authenticated session — never from a
   path/body param the client controls. This is the highest-risk piece per
   the proposal's "vendor data leakage" risk — needs an explicit test proving
   vendor A cannot read/write vendor B's menu.

4. **Public menu browsing (unauthenticated)**
   List vendors; list a vendor's categories/products. Only rows with
   availability = true are exposed to this path.

## Open questions to settle before/while building

- Does "vendor authentication" need password hashing + sessions, or is a
  simpler bootstrap (e.g. a seeded vendor + static token) acceptable for this
  milestone, with real auth hardening deferred to Milestone 7 (release
  readiness)? Lean toward real but minimal (hashed password, server-side
  session or signed token) — retrofitting auth later touches every
  vendor-scoped query again.
- `food_courts` and `tables` are listed under the proposal's "Food court"
  domain alongside `vendors`, but table/session logic is Milestone 3. Decide
  whether `vendors` needs a `food_court_id` now (even hardcoded to one row)
  or can be added in Milestone 3 without breaking this milestone's queries.
  Given "single food court for MVP," simplest is: seed one `food_courts` row
  now, give `vendors.food_court_id` a FK to it, and not build any food-court
  management UI/endpoints until they're actually needed.

## Completion check

- Vendor can log in and manage (create/update/toggle availability) its own products.
- A vendor cannot see or modify another vendor's menu (tested).
- An unauthenticated customer-facing endpoint lists vendors and their available products.

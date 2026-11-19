---
version: 1
slug: "app-vendor"
primary_target: "app/vendor"
related_targets: []
---

# Surface: vendor (login + dashboard)

Mode: Operate. Vendor staff at a food-court stall, phone/tablet/laptop, glanced at mid-rush.
Job now: sign in, confirm which stall you are, sign out (button only, no API yet).
Later on this surface: incoming orders, accept/reject, prep states.
Constraints: shadcn/ui standard controls, no fabricated orders or stats, English only.

## Direction contract

THESIS: Vendor screens are a stack of tiers like a tiffin carrier: one job per tier, opened by a steel band that names it. Refuses the grey admin shell with one blue accent and floating cards.

OWN-WORLD: Jade melamine #2E7D5B owns the top rail and primary actions; ground is a jade-tinted near-white, panels white, neutrals all from one jade-tinted ramp. Steel bands (#DDE4E0 field, ink text, weight-led) head every tier. Rose #E58E9A only as a tiny rim mark. Flat: no gradients, glow, glass, rivets or metal texture. One sans (Figtree), rank by weight not size.

STORY: Staff see their stall name first, know they are signed in, and see where orders will land.

FIRST VIEWPORT: Login: flat food court floor plan (numbered stalls, tables) fills the screen; the form is the one highlighted stall in the middle (3px jade border, 'Your stall' pill), heading 'Sign in' + wordmark, 48px fields, full-width jade button. No role-specific copy: owners and staff both sign in here. Dashboard: jade rail (Forkly, stall name, account menu with username + Sign out); tiers below: "Orders" band with count 0 and an empty state that teaches.

FORM: candidate 5 of 7 (melamine tableware & tiffin carrier), seed:6037ff2cc006798f. Raises: one tonal ramp (exposure record); weight not size (timetable); flat fields (garden map); whole-cell repack (event catalog).

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance

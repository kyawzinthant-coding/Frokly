---
name: Forkly
description: Food court ordering — vendor, cashier and manager screens in the jade melamine world.
colors:
  jade: "#2e7d5b"
  on-jade: "#ffffff"
  ground: "#f4f7f5"
  surface: "#ffffff"
  ink: "#1d2622"
  ink-muted: "#53625b"
  steel-band: "#dde4e0"
  jade-wash: "#e6ede9"
  line: "#d3dcd7"
  input-line: "#b9c6bf"
  rim: "#e58e9a"
  error: "#b3261e"
typography:
  page-title:
    fontFamily: "Figtree, system-ui, sans-serif"
    fontSize: "1.5rem"
    fontWeight: 600
    lineHeight: 1.33
  rail-title:
    fontFamily: "Figtree, system-ui, sans-serif"
    fontSize: "1.125rem"
    fontWeight: 600
  body:
    fontFamily: "Figtree, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
    lineHeight: 1.5
  band-label:
    fontFamily: "Figtree, system-ui, sans-serif"
    fontSize: "0.875rem"
    fontWeight: 600
  wordmark:
    fontFamily: "Figtree, system-ui, sans-serif"
    fontSize: "1.25rem"
    fontWeight: 700
    letterSpacing: "-0.025em"
rounded:
  control: "12px"
  tier: "16px"
  pill: "9999px"
spacing:
  tier-pad: "20px"
  form-pad: "24px"
  page-gutter: "16px"
  page-gutter-sm: "24px"
components:
  rail:
    backgroundColor: "{colors.jade}"
    textColor: "{colors.on-jade}"
    height: "64px"
  tier:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.tier}"
    padding: "{spacing.tier-pad}"
  tier-band:
    backgroundColor: "{colors.steel-band}"
    textColor: "{colors.ink}"
    typography: "{typography.band-label}"
    padding: "10px 20px"
  button-primary:
    backgroundColor: "{colors.jade}"
    textColor: "{colors.on-jade}"
    rounded: "{rounded.control}"
    height: "48px"
  input:
    backgroundColor: "{colors.surface}"
    textColor: "{colors.ink}"
    rounded: "{rounded.control}"
    height: "48px"
---

# Design

## Overview

Forkly's working screens borrow from a food court's jade melamine plates and stacked tiffin
carriers. Every screen is a stack of **tiers**: one job per tier, each opened by a flat steel
band that names it. Jade owns the top rail and primary actions. Everything is flat — the world
lends palette, density and the tier band; layout, navigation and controls stay standard.

## Colors

- **Jade** `#2e7d5b` — top rail, primary buttons, focus ring, icons in empty states. White text
  on jade passes 4.5:1; never put tinted/secondary text on jade, use white and change weight.
- **Ground** `#f4f7f5` page background; **Surface** `#ffffff` tiers, cards, menus.
- **Steel band** `#dde4e0` — tier headings only.
- All neutrals come from one jade-tinted ramp (`jade-wash`, `line`, `input-line`, `ink-muted`).
  No stock greys.
- **Rim** `#e58e9a` — only the wordmark's dot. Never a fill, never text.
- **Error** `#b3261e` — error text and invalid input borders, always with an icon and words.
- Light theme only for MVP (bright food-court lighting).

## Typography

One family, Figtree, at a tight scale. Rank comes from weight and the band, not size jumps:
page title 24/600, rail stall name 18/600, body 16/400, band labels 14/600. Inputs and buttons
use 16px text so phones don't zoom. Thai/Burmese will need a fallback face later.

## Layout

Content sits in a centred column, max 1152px, 16px gutters (24px from `sm`). Tiers stack in one
column on phones and tablets; from `lg` they sit in a 3-column grid (main tier spans 2), aligned
to the top and repacked in whole tiers — never half-cut. Login form is max 384px, centred.

**Login** sits on a flat food court floor plan (`components/vendor/floor-plan.tsx`): numbered
stalls around the edge, tables as outlined circles, all in `jade-wash` / `line`. The form is drawn
as the one highlighted stall: white, 3px jade border, a jade "Your stall" pill on its top edge.
Three drawings (landscape, tablet portrait, phone) are scaled to fit, never cropped.

## Elevation & Depth

Flat. Tiers and the login card declare elevation once with a 1px `line` ring and no shadow.
Only floating menus get a shadow (from the shadcn popover). No gradients, glow or glass.

## Shapes

Tiers and the login card 16px radius; inputs and buttons 12px; avatar initials fully round.

## Components

- **Rail** (`app/vendor/(dashboard)/layout.tsx`): wordmark, divider, stall name, account menu
  on the right. 64px tall.
- **Tier** (`components/vendor/tier.tsx`): steel band heading + white body.
- **Account menu** (`components/vendor/account-menu.tsx`): initial avatar + username trigger,
  44px tall; menu shows "Signed in as …" and **Sign out**.
- **Empty state**: jade icon, bold one-line title, one sentence that teaches what will appear.
- **Primary button / input**: 48px tall, full width in forms.
- Loading: skeletons in place of the content they stand for, never a centred spinner.

## Do's and Don'ts

- Do give every new section its own tier with a named band.
- Do show real data or an empty state that teaches — never placeholder numbers.
- Don't add rivets, metal textures, plate rims or gradients: the world stays flat.
- Don't use rim pink or jade for decoration; colour marks actions and state.
- Don't put an eyebrow label above a heading; the band is the only label.

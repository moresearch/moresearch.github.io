---
title: "Corne 4x6: Miryoku, Colemak-DH, and 24 macro keys"
date: 2026-10-21
slug: corne-4x6-miryoku-colemak-dh
summary: "A 60-key Corne 4x6, Miryoku, Colemak-DH, and the 24 keys left over for custom macros."
tags: keyboards, corne, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a Corne 4x6. I wanted a metal case, could not find one in Brazil, and settled for non-metal until I make my own custom one. The layout matters more than the case.

![Corne 4x6 — 60 keys, split, column-staggered](images/corne-4x6.jpg)

## Layout

**Matrix.** My board has 60 keys. Firmware defines what each key sends, so every key is programmable.

**Colemak-DH.** Moves `D` and `H` from Colemak's centre columns to the bottom row under the index fingers, `G` and `M` to home, and `B` to the top. `Z`, `X`, `C`, and `V` stay put. Use the matrix variant (formerly DHm); no Angle Mod is needed. It is also Miryoku's default alpha layer.

**Miryoku.** A 36-key design: five columns, three rows, and three thumbs per hand. Layers handle navigation, numbers, symbols, functions, media, and mouse; hold a thumb on the opposite hand to activate one. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the resting fingers.

## 60 keys − 36 = 24 macro keys

Miryoku's core is 36 keys: 30 alpha keys (five columns × three rows per hand) plus six thumb keys. My board has 60, so **24 keys are left for macros** — 12 per hand.

| Piece | Keys |
|---|---:|
| Board | 60 |
| Colemak-DH alphas (5×3 × 2) | 30 |
| Miryoku thumbs (3×2) | 6 |
| **Miryoku core** | **36** |
| **Left for macros** | **24** |

This is not the six-key situation of a 42-key Corne, where the only extras are one outer column per hand. On a 60-key board the surplus is four times larger. If you count only the Colemak-DH alpha block, it looks like 30 keys are free, but six of those become the Miryoku thumb/layer keys — so the real number is 24.

They are not all equally comfortable. The inner keys are easier to reach; the outer and top keys are not. So the easy ones get the daily macros, the awkward ones get the rare ones, and some can stay unbound:

- **Daily six:** `git status`, `make test`, comment line, screenshot, Portuguese accents, play/pause.
- **Next tier:** markdown link, email signature, window switch, terminal, app launcher, duplicate line.
- **Remainder:** app-specific macros, experiments, or nothing.

With QMK or ZMK, a key can send a string or a chord and change per layer, so the 24 keys can become a dev layer, a writing layer, or a dedicated macro layer.

**Practice:** Colemak Club, Keybr, and Monkeytype. Start in Colemak-DH from day one — no QWERTY phase — then add Miryoku layers, then macros. I will follow up with speed, accuracy, and which macros stick.

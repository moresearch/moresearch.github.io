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

## 24 spare keys

Miryoku's core uses 36 keys. My board has 60. That leaves **24 keys** for macros — 12 per hand on a symmetric split.

They are not all equally comfortable. The inner keys are easier to reach; the outer ones are not. But every one is programmable, so the easy ones get the frequent macros and the awkward ones get the rare ones.

Ideas: `git status`, `make test`, comment line, screenshot, markdown link, `á ã ç é ê í ó ô ú`, play/pause. With QMK or ZMK, a key can send a string or a chord and change per layer, so the 24 keys can become a dev layer, a writing layer, or a dedicated macro layer.

**Practice:** Colemak Club, Keybr, and Monkeytype.

Plan: QWERTY first, then Colemak-DH drills, then Miryoku layers, then macros. I will follow up with speed, accuracy, and which macros stick.

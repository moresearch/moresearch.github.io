---
title: "Corne 4x6: Miryoku, Colemak-DH, and six macro keys"
date: 2026-10-21
slug: corne-4x6-miryoku-colemak-dh
summary: "A non-metal Corne 4x6, Miryoku, Colemak-DH, six spare macro keys, and the tools I will use to learn the layout."
tags: keyboards, corne, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a Corne 4x6. I wanted a metal case, could not find one in Brazil, and settled for non-metal until I make my own custom one. The layout matters more than the case.

![Corne 4x6 — 42 keys, split, column-staggered](images/corne-4x6.jpg)

## Layout

**Matrix.** 42 keys: three rows of six columns per hand, plus three thumbs. Firmware defines every key, so every key is programmable.

**Colemak-DH.** Moves `D` and `H` from Colemak's centre columns to the bottom row under the index fingers, `G` and `M` to home, and `B` to the top. `Z`, `X`, `C`, and `V` stay put. Use the matrix variant (formerly DHm); no Angle Mod is needed. It is Miryoku's default alpha layer.

**Miryoku.** A 36-key design: five columns, three rows, three thumbs per hand. Layers handle navigation, numbers, symbols, functions, media, and mouse; hold a thumb on the opposite hand to activate one. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the resting fingers.

## Six spare keys

Miryoku uses five columns per hand; the Corne has six. The extra outer column — three keys per side, six total — is awkward to reach, which makes it ideal for infrequent macros.

Ideas: `git status`, `make test`, comment line, screenshot, markdown link, `á ã ç é ê í ó ô ú`, play/pause. QMK or ZMK can send strings or chords and change per layer, so the six can become a dev row, a writing row, or a macro layer.

## Practice

- **Colemak Club** — best for beginners. Unlocks keys level by level from the home row `R S T / N E I`. Select Colemak-DH (or Mod-DH) in the top menu.
- **Keybr** — adaptive pseudo-words targeting weak keys. Settings > Keyboard Layout > Layout: Colemak, then the DH variant.
- **Monkeytype** — speed and endurance. Settings > Layout > Colemak DH / Mod-DH; it shows an on-screen keymap.

Plan: QWERTY first, then Colemak-DH drills, then Miryoku layers, then macros.

I will follow up with speed, accuracy, and which macros stick.

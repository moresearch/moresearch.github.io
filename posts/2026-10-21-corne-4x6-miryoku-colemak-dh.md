---
title: "Corne 4x6: Miryoku, Colemak-DH, and six macro keys"
date: 2026-10-21
slug: corne-4x6-miryoku-colemak-dh
summary: "A non-metal Corne 4x6, Miryoku, Colemak-DH, six spare macro keys, and the tools I will use to learn the layout."
tags: keyboards, corne, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a Corne 4x6. I wanted a metal case, could not find one in Brazil, and settled for a non-metal build until I make my own custom metal case one day. The case can wait; the layout is the interesting part.

![Corne 4x6 — 42 keys, split, column-staggered](images/corne-4x6.jpg)

## Layout

**Matrix.** The Corne is a 42-key split: three rows of six columns per hand, plus three thumbs. Firmware decides what each key sends, so every key is programmable.

**Colemak-DH.** Moves `D` and `H` from Colemak's centre columns to the bottom row under the index fingers, puts `G` and `M` on home and `B` on top. `Z`, `X`, `C`, and `V` stay put. On a matrix board use the matrix variant (formerly DHm); no Angle Mod is needed. It is also Miryoku's default alpha layer.

**Miryoku.** A 36-key design: five columns, three rows, and three thumbs per hand. Layers — Base, Nav, Mouse, Num, Sym, Fun, Media, Button — replace reaching and are held from the opposite thumb. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the resting fingers. It covers a TKL plus media and mouse.

## Six spare keys

Miryoku uses five columns per hand; the Corne has six. The extra outer column — three keys per side, six total — sits outside the core. It is awkward to reach, which makes it perfect for infrequent macros.

Ideas: `git status`, `git diff`, `make test`; duplicate/comment line; screenshot; markdown link; `á ã ç é ê í ó ô ú`; play/pause. Because QMK or ZMK lets a key send a string, chord, or sequence and change per layer, the six can become a dev row, a writing row, or a dedicated macro layer.

## Practice

- **Colemak Club** — best for absolute beginners. Gradually unlocks keys level by level, starting with the home row `R S T / N E I`. Select Colemak-DH (or Mod-DH) in the top menu.
- **Keybr** — adaptive learning. Generates pseudo-words targeting your weakest keys. Settings > Keyboard Layout > Layout: Colemak, then select the DH variant.
- **Monkeytype** — speed and endurance. Settings > Layout > Colemak DH / Mod-DH. It shows an on-screen keymap, so you do not have to look at your hands.

Plan: QWERTY on the Corne for a few days, then Colemak-DH drills, then Miryoku layers one at a time, then macros.

I will follow up with speed, accuracy, and which macros earned a permanent place.

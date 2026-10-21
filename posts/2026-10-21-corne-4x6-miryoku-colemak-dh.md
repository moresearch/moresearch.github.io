---
title: "Corne 4x6: Miryoku, Colemak-DH, and six macro keys"
date: 2026-10-21
slug: corne-4x6-miryoku-colemak-dh
summary: "I bought a non-metal Corne 4x6 after failing to find a metal case in Brazil. The layout is the real project: Miryoku, Colemak-DH, and six spare keys for custom macros."
tags: keyboards, corne, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a Corne 4x6. I wanted a metal case, could not find one in Brazil, and settled for a non-metal build until I make my own custom metal case one day. The case can wait. The layout is the interesting part.

![Corne 4x6 — 42 keys, split, column-staggered](images/corne-4x6.jpg)

## Matrix

The Corne is a 42-key split: three rows of six columns per hand, plus three thumb keys. "4x6" refers to the switch matrix. The physical matrix is separate from the logical layout — firmware decides what each key sends — and that is what makes custom macros possible.

## Colemak-DH

Standard Colemak puts `D` and `H` in the centre columns, forcing lateral index-finger movement. Colemak-DH moves them to the bottom row under the index fingers, moves `G` and `M` to home, and `B` to the top. `Z`, `X`, `C`, and `V` stay put, so common shortcuts still work. On a matrix board, use the matrix variant (formerly DHm); the Angle Mod is unnecessary. Miryoku's default alpha layer is Colemak Mod-DH, so the two fit together.

## Miryoku

Miryoku is a 36-key design: five columns, three rows, and three thumbs per hand. Layers replace reaching — Base, Nav, Mouse, Num, Sym, Fun, Media, Button — and a layer is activated by holding a thumb key on the opposite hand. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the fingers that already rest there. Left thumbs are `Space`, `Tab`, `Esc`; right thumbs are `Backspace`, `Enter`, `Delete`. It covers a full TKL keyboard, plus media and mouse emulation.

Expected benefits: less finger travel, fewer pinky stretches, useful thumbs, straighter wrists, and navigation without leaving the keyboard.

## Six spare keys

Miryoku uses five columns per hand. The Corne has six. The extra outer column on each side — three keys per hand, six total — sits outside the core layout. It is awkward to reach, which makes it ideal for infrequent actions, but every key is a real switch and fully programmable.

Macro ideas:

- **Development:** `git status`, `git diff`, `make test`, `docker compose up -d`, `cargo check`, a debug print.
- **Editing:** duplicate, comment, or select line; move line up/down; go to definition.
- **Desktop:** screenshot, lock screen, terminal, window switcher.
- **Writing:** signature, email address, markdown link template, common phrase.
- **Portuguese:** `á`, `ã`, `ç`, `é`, `ê`, `í`, `ó`, `ô`, `ú` via dead-key sequences or Unicode macros, depending on OS and layout.
- **Media and apps:** play/pause, mute, OBS scene, app launcher.

Because the board runs QMK or ZMK, a key can send a string, a chord, or a sequence, and the same key can do different things per layer. The six keys could become a dev row, a writing row, or a dedicated macro layer.

## Why this setup

- The **matrix** makes every key programmable.
- **Miryoku** needs only 36 keys, leaving room to add without crowding the home position.
- **Colemak-DH** is already Miryoku's default, so the layout and layers agree.
- The **4x6 Corne** has six spare keys; a 3x5 board would have none.

## Learning plan

1. Type on the Corne in QWERTY for a few days to learn the geometry.
2. Switch to Colemak-DH and drill until the home row feels normal.
3. Add Miryoku layers one at a time.
4. Add macros last, once the core layout is automatic.

I will write a follow-up with speed, accuracy, and which macros earned a permanent place.

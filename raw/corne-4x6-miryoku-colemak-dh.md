---
title: "Corne 4x6: Miryoku, Colemak-DH, and six macro keys"
date: 2026-10-21
slug: corne-4x6-miryoku-colemak-dh
summary: "A non-metal Corne 4x6, Miryoku, Colemak-DH, and six spare keys I can wire into custom macros."
tags: keyboards, corne, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a Corne 4x6. I wanted a metal case, could not find one in Brazil, and settled for a non-metal build until I make my own custom metal case one day. The case can wait; the layout is the interesting part.

![Corne 4x6 — 42 keys, split, column-staggered](images/corne-4x6.jpg)

## Matrix

The Corne is a 42-key split: three rows of six columns per hand, plus three thumbs. "4x6" is the switch matrix. The physical matrix is separate from the logical layout — firmware decides what each key sends — which is what makes custom macros possible.

## Colemak-DH

Standard Colemak puts `D` and `H` in the centre columns, forcing lateral index-finger movement. Colemak-DH moves them to the bottom row under the index fingers, puts `G` and `M` on home, and `B` on top. `Z`, `X`, `C`, and `V` stay put, so shortcuts still work. On a matrix board, use the matrix variant (formerly DHm); no Angle Mod is needed. Miryoku's default alpha layer is Colemak Mod-DH, so the two fit together.

## Miryoku

Miryoku is a 36-key design: five columns, three rows, and three thumbs per hand. Layers replace reaching — Base, Nav, Mouse, Num, Sym, Fun, Media, Button — activated by holding a thumb on the opposite hand. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the resting fingers. It covers a full TKL keyboard, plus media and mouse emulation. The payoff is less finger travel, fewer pinky stretches, and navigation without leaving the keyboard.

## Six spare keys

Miryoku uses five columns per hand; the Corne has six. The extra outer column — three keys per side, six total — sits outside the core layout. It is awkward to reach, which makes it perfect for infrequent actions, and every key is fully programmable through QMK or ZMK.

**Macro key ideas:**

- **Development:** `git status`, `git diff`, `make test`, `docker compose up -d`, `cargo check`, a debug print.
- **Editing:** duplicate, comment, or select line; move line; go to definition.
- **Desktop:** screenshot, lock screen, terminal, window switcher.
- **Writing:** signature, email address, markdown link template, common phrase.
- **Portuguese:** `á`, `ã`, `ç`, `é`, `ê`, `í`, `ó`, `ô`, `ú` via dead keys or Unicode macros.
- **Media and apps:** play/pause, mute, OBS scene, app launcher.

A possible loadout for the six keys: screenshot, `git status`, `make test`, markdown link, a `ç`/`ã` macro, and play/pause. Since a key can send a string, a chord, or a sequence — and can change per layer — the six keys could also become a dev row, a writing row, or a dedicated macro layer.

## Learning plan

1. Type in QWERTY on the Corne for a few days.
2. Switch to Colemak-DH and drill until home row feels normal.
3. Add Miryoku layers one at a time.
4. Add macros last, once the core layout is automatic.

I will follow up with speed, accuracy, and which macros earned a permanent place.

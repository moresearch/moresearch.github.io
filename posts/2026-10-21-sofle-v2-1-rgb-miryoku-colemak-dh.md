---
title: "Sofle V2.1 RGB: Miryoku 36 for typing, 24 controls for macros"
date: 2026-10-21
slug: sofle-v2-1-rgb-miryoku-colemak-dh
summary: "A 58-key SZRKBD Sofle V2.1 RGB running Miryoku on 36 pure typing keys, with every extra key and encoder reserved for macros."
tags: keyboards, sofle, miryoku, colemak-dh, qmk, zmk, macros, ergonomics, split-keyboard, brazil
---

I just bought a split keyboard — specifically the **SZRKBD Sofle V2.1 RGB**. It has **58 keys and two rotary encoders**; SZRKBD markets this version as a 60-key board. I wanted a metal case, could not find one in Brazil, and settled for non-metal until I make my own custom one. The layout matters more than the case.

![Sofle V2.1 RGB — 58 keys plus two encoders, split, column-staggered](images/sofle-v2-1-rgb.jpg)

The plan is deliberately split in two:

- **Miryoku 36 is pure typing.** The 36-key Miryoku core stays clean: alphas, thumbs, home-row mods, and the layers Miryoku already defines.
- **Every extra control is a macro.** The 22 keys outside the Miryoku core plus the two encoders — 24 controls in total — are 100% macros, and their macros change with the active layer.

[![34 keys is all you need: an ergonomic keyboard journey](images/yt-unMXQTSQEak.jpg)](https://www.youtube.com/watch?v=unMXQTSQEak)

*Video: [34 keys is all you need: an ergonomic keyboard journey](https://www.youtube.com/watch?v=unMXQTSQEak).*

## Why this wild journey

I am not doing this for novelty. The default keyboard is a set of compromises I never chose: QWERTY, row stagger, pinky modifiers, and a number row that pulls the hands off home. I type for hours every day, so small inefficiencies compound. This board lets me fix the geometry and the layout at once.

The history matters because this is not gadget churn. QWERTY dates to Sholes and Remington in the 1870s and was shaped by mechanical typebars, not hand comfort; the story that it was designed to slow typists down is a myth, but it was never an ergonomic layout. Splits are older than the hobby suggests: Klockenberg argued in 1926 that the two hand clusters should be separated and angled so the wrists stay straight, Kroemer suggested in 1972 that adjustable splits could reduce pain, and commercial splits arrived in the 1990s. The Sofle belongs to that lineage: a 6×4 column-staggered grid plus a five-key thumb cluster per hand, 58 keys, descended from the Lily58, Corne, and Helix family. Colemak (2006) and Colemak-DH (2014) put common letters under stronger fingers; Miryoku (2019) packages Colemak-DH into a 36-key layer system with home-row mods and six thumb keys.

The point of buying a board this size is not to fill every key with letters. Miryoku needs 36 keys. The Sofle gives me 58 plus two encoders. The surplus is the macro surface.

The three changes reinforce each other. Miryoku assumes Colemak-DH. The Sofle's 58 keys leave enough room for Miryoku's 36 plus a clean block of macro keys. Learning them together is harder up front, but it avoids relearning the same habits twice. The keymap is firmware, not hardware, so the investment moves to my next board — including the custom metal case I still want to build.

## Layout

**Matrix.** The SZRKBD Sofle V2.1 RGB has 58 keys and two encoders, marketed as 60. Each hand gets six columns, four rows, and five thumb keys. Firmware defines what each key sends, so every key and encoder is programmable.

**Colemak-DH.** Moves `D` and `H` from Colemak's centre columns to the bottom row under the index fingers, `G` and `M` to home, and `B` to the top. `Z`, `X`, `C`, and `V` stay put. Use the matrix variant (formerly DHm); no Angle Mod is needed. It is also Miryoku's default alpha layer.

**Miryoku.** A 36-key design: five columns, three rows, and three thumbs per hand. Layers handle navigation, numbers, symbols, functions, media, and mouse; hold a thumb on the opposite hand to activate one. Home-row mods put `Ctrl`, `Alt`, `Shift`, and `GUI` under the resting fingers.

[![Sofle V2 Build Guide (2025): Complete Soldering Walkthrough](images/yt-_3pe236--hA.jpg)](https://www.youtube.com/watch?v=_3pe236--hA)

*Build reference: [Sofle V2 Build Guide (2025): Complete Soldering Walkthrough](https://www.youtube.com/watch?v=_3pe236--hA).*

## 58 keys − 36 = 22 macro keys, plus two encoders

Miryoku's core is 36 keys: 30 alpha keys (five columns × three rows per hand) plus six thumb keys. The Sofle has 58 keys, so **22 keys are left over**, plus the two encoders. That is **24 macro controls** — 11 keys and one encoder per hand.

| Piece | Controls |
|---|---:|
| Board (58 keys + 2 encoders) | 60 |
| Colemak-DH alphas (5×3 × 2) | 30 |
| Miryoku thumbs (3×2) | 6 |
| **Miryoku core** | **36** |
| **Left for macros** | **24 (22 keys + 2 encoders)** |

Per hand, the 11 extra keys are easy to locate: the unused sixth column (four keys), the fourth-row key above each of the five alpha columns (five keys), and two of the five thumb keys. None of them is part of Miryoku.

They are not all equally comfortable. The inner keys are easier to reach; the outer and top keys are not. So the easy ones get the daily macros, the awkward ones get the rare ones, and some can stay unbound:

- **Daily six:** `git status`, `make test`, comment line, screenshot, Portuguese accents, play/pause.
- **Next tier:** markdown link, email signature, window switch, terminal, app launcher, duplicate line.
- **Remainder:** app-specific macros, experiments, or nothing.

The separation is the point. **Layer switching stays on the Miryoku core**, where Miryoku already puts it. The 22 extra keys never switch layers and never type letters. They only fire macros. That keeps the typing layer pure and means an accidental macro key cannot drop me into the wrong layer.

Because every extra key is macro-only, the same key can carry a different macro on each active Miryoku layer. Miryoku has eight main layers — Base, Nav, Mouse, Button, Media, Num, Sym, Fun — so the spare keys represent **24 × 8 = 192 possible macro assignments**. Only 24 are reachable at once, but the same physical key can run a dev macro on Base, a window command on Nav, and a transport control on Media. The two encoders follow the same rule: volume or scrolling on Base, zoom or tab switching on another layer.

## What I expect

**First month:** slower typing, wrong layers, accidental home-row mods, and a lot of macro tuning. **By month three:** Colemak-DH comfortable, the main Miryoku layers automatic, the extra keys mapping to muscle memory, and a handful of macros in daily use. **Long term:** less finger travel and pinky reaching, more comfortable long sessions, faster navigation, and a keyboard shaped around my workflow instead of the other way around.

I do not expect magic. Speed may dip before it recovers, and comfort matters more than WPM. If it fails, I can flash QWERTY back. But the point is to stop treating it as the default.

**Practice:** [Colemak Club](https://gnusenpai.net/colemakclub/), [Keybr](https://www.keybr.com/), and [Monkeytype](https://monkeytype.com/). Start in Colemak-DH from day one — no QWERTY phase — then add the Miryoku 36, then map the extras as macros. I will follow up with speed, accuracy, and which macros stick.

---
title: Launching the Hackspree blog
date: 2026-05-08
slug: launching-the-hackspree-blog
summary: A fresh single-page blog, generated from Markdown and published straight to GitHub Pages.
tags: launch, blog, build
---

This repo now ships the blog as a single generated page instead of a demo layout.

The writing workflow is intentionally small:

1. Drop a new Markdown file into `posts/`.
2. Run `make`.
3. Commit the regenerated `index.html`.

That keeps publishing dead simple while still making the source of truth pleasant to edit.

> Hack first, polish second, keep shipping either way.

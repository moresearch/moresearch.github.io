---
title: Why the blog stays single-page
date: 2026-05-07
slug: why-the-blog-stays-single-page
summary: Keeping every post on one page makes the build trivial, the hosting static, and the writing flow fast.
tags: architecture, static-site
---

The goal here is speed, not ceremony.

By keeping the blog as one generated `index.html`, the site stays easy to host on GitHub Pages and easy to inspect in a diff. There is no runtime, no JavaScript hydration step, and no template sprawl to maintain.

Markdown still gives enough structure for notes, lists, code snippets, and short essays:

```sh
make
git add posts index.html
git commit -m "feat(blog): publish new post"
```

For now, that is enough surface area to keep writing instead of maintaining tooling.

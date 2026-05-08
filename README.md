# blog.hackspree.com

Single-page blog source for `blog.hackspree.com`.

## Write a post

1. Add a new Markdown file under `posts/` with front matter:

   ```md
   ---
   title: Your post title
   date: 2026-05-08
   slug: optional-slug
   summary: One-sentence summary
   tags: engineering, build, notes
   ---
   ```

2. Run `make`.
3. Commit the updated Markdown and generated `index.html`.

## Build

`make` uses the Go toolchain to regenerate the published `index.html`.

The site is intentionally JavaScript-free and published directly by GitHub Pages from the repository root.

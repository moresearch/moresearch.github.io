# Post Specs

## User-facing behavior

- Each post source Markdown file in `posts/` becomes a rendered blog section inside the single-page site.
- Posts display a title, publication date, optional summary, and the rendered Markdown body.
- Posts appear newest first.

## Implementation rules

- Post source files must use the `.md` extension.
- `posts/specs.md` is reserved for contributor guidance and is never published as a post.
- Each post must start with front matter delimited by `---` lines.
- Required front matter fields are `title` and `date`.
- Optional front matter fields are `slug`, `summary`, and `tags`.
- `date` must use ISO-8601 calendar format (`YYYY-MM-DD`).
- The source `date` remains ISO-8601 in front matter even though the published site renders it as a larger Unix timestamp with a UTC hover tooltip.
- Front matter dates control the published chronology of the archive, so authors should set them deliberately.
- `slug`, when provided, must be URL-safe lowercase text using letters, numbers, and hyphens.
- If `slug` is omitted, the build derives it from the title.
- Markdown content begins after the closing front matter delimiter.
- Posts may cite external sources with standard Markdown links in the body, including direct links to arXiv paper pages.
- Fenced code blocks with language info strings are allowed and should be used when posts include code examples.
- Highlighted code blocks inherit the shared site-wide syntax theme automatically; authors should only supply the correct language info string.
- Markdown links inside post bodies are rendered to open in a separate browser tab/window by default.
- Rendered post bodies use the site's configured monospace body font; post authors should not rely on custom font styling in Markdown.
- Rendered post-body links follow the site theme and should not rely on underline styling for visibility.
- Each post slug also drives the in-page navigation state, so slugs must stay stable and unique.
- Post length may vary, including medium-length essays, as long as the Markdown stays readable in the single-page archive.

## Constraints and invariants

- Two posts must never resolve to the same slug.
- Draft handling is out of scope until explicitly added; every valid post source Markdown file in `posts/` is published.
- Post bodies may include standard Markdown links, lists, headings, blockquotes, and fenced code blocks.

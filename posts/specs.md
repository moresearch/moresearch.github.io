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
- `slug`, when provided, must be URL-safe lowercase text using letters, numbers, and hyphens.
- If `slug` is omitted, the build derives it from the title.
- Markdown content begins after the closing front matter delimiter.

## Constraints and invariants

- Two posts must never resolve to the same slug.
- Draft handling is out of scope until explicitly added; every valid post source Markdown file in `posts/` is published.
- Post bodies may include standard Markdown links, lists, headings, blockquotes, and fenced code blocks.

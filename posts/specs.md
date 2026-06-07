# Post Specs

## User-facing behavior

- Each post source Markdown file in `posts/` becomes a rendered blog section inside the single-page site.
- Posts display a title, publication date, optional summary, and the rendered Markdown body.
- Posts appear newest first.

## Implementation rules

- Post source files must use the `.md` extension.
- `posts/specs.md` is reserved for contributor guidance and is never published as a post.
- Each post must start with front matter delimited by `---` lines.
- Required front matter fields are `title` and `date`. Keys are case-sensitive and must be lowercase.
- Optional front matter fields are `slug`, `summary`, and `tags`.
- `date` must use ISO-8601 calendar format (`YYYY-MM-DD`).
- The source `date` remains ISO-8601 in front matter even though the published site renders it as a larger Unix timestamp with a UTC hover tooltip.
- Front matter dates control the published chronology of the archive, so authors should set them deliberately.
- `slug`, when provided, must be URL-safe lowercase text using letters, numbers, and hyphens.
- If `slug` is omitted, the build derives it from the title.
- Markdown content begins after the closing front matter delimiter.
- Posts may cite external sources with standard Markdown links in the body, including direct links to arXiv paper pages.
- Posts may also cite workshops, talks, videos, and companion repositories when a reflection post is grounded in their technical content.
- Posts that cite YouTube videos should include a linked thumbnail image near the relevant discussion rather than a text-only reference. Use the YouTube thumbnail URL format `https://img.youtube.com/vi/<VIDEO_ID>/hqdefault.jpg` and wrap the image in a Markdown link to the video. When referencing a specific scene, include a timestamp (MM:SS) or range (MM:SS–MM:SS) and link to the video with that timestamp.
- Posts may embed remote publisher-hosted book-cover images when discussing influential books tied directly to the post topic.
- Posts may compare concrete Go libraries or frameworks when the comparison is grounded in current official docs and includes code examples.
- Posts may connect personal skill development to games or other systems when the essay stays specific about the engineering habits or mental models that transferred.
- Fenced code blocks with language info strings are allowed and should be used when posts include code examples.
- Highlighted code blocks inherit the shared site-wide syntax theme automatically; authors should only supply the correct language info string.
- When a post includes code, the shared snippet styling should render it in a monochrome `vim-256noir`-style palette and Sudo font without post-level overrides.
- Go code examples should include brief comments so the example remains self-explanatory inside the single-page archive.
- Code-block scrollbars should only become visibly styled when the reader hovers the block.
- Markdown links inside post bodies are rendered to open in a separate browser tab/window by default.
- Rendered post bodies use the site's configured monospace body font; post authors should not rely on custom font styling in Markdown.
- Rendered post-body links follow the site theme and should not rely on underline styling for visibility.
- Each post slug also drives the in-page navigation state, so slugs must stay stable and unique.
- Post length may vary, including medium-length essays, as long as the Markdown stays readable in the single-page archive.
- Reflection posts should explain what specifically was learned, what tradeoffs mattered, and why those details changed the author's view of the topic.

## Constraints and invariants

- Two posts must never resolve to the same slug.
- Draft handling: Posts may include optional front matter `draft: true` to exclude them from the published archive. When consolidating posts, create a canonical merged post (with a stable slug) and mark original files with `draft: true` and a short note linking to the canonical post. The generator will skip files marked `draft: true`. 
- Post bodies may include standard Markdown links, lists, headings, blockquotes, and fenced code blocks.

## Change log:
- Added post 2026-05-17-task-harness-engineering.md: "Task Harness Engineering" (source: YouTube C_GG5g38vLU).
- Added post 2026-05-17-harness-engineering-fowler.md: "Harness Engineering (Martin Fowler)" (source: https://martinfowler.com/articles/harness-engineering.html).
- Added post 2026-06-06-ai-sovereignty-or-ai-colony.md: "AI sovereignty or AI colony: why domestic capability matters" (source: https://thenextweb.com/news/japan-risks-becoming-an-ai-colony-its-digital-minister-warns).

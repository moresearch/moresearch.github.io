# Post Specs

## User-facing behavior

- Each post source Markdown file in `posts/` becomes a rendered blog section inside the single-page site.
- Posts display a title, publication date, optional summary, and the rendered Markdown body.
- Posts appear newest first.

## Implementation rules

- Post source files must use the `.md` extension.
- Only files inside `posts/` are published as posts; this spec lives at the repo root (`SPECS.md`) so it can never be mistaken for a draft.
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
- Updated post 2026-06-06-ai-sovereignty-or-ai-colony.md to include the requested article photo: https://media.thenextweb.com/2026/06/Hisashi-Matsumoto.avif.
- Added post 2026-09-23-llm-at-runtime-is-an-antipattern.md: "LLM at Runtime Is an Antipattern" — an LLM in the event loop converts a fixed cost into a variable cost and gives up replayability, auditability, versioning, and determinism; the fix is a hard wall between compile time (LLM as the semantic layer of AutoML, proposing candidate features) and runtime (frozen, hash-bound artifacts plus a small traditional model). Discusses TypeSafe's System One / Jev launch post (https://typesafe.ai/blog/introducing-system-one-models-and-jev) in detail, including three linked figures from it (the published decision DAG, the workflow-eval Pareto plot, and the wrong-tool-call / type-error evaluation, hosted at framerusercontent.com), plus Laya (https://huggingface.co/convaiinnovations/laya) as the open-weight counterpart. References include Brooks (*No Silver Bullet*, https://www.cs.unc.edu/techreports/86-020.pdf), Boehm & Basili (*Software Defect Reduction Top 10 List*, https://web.archive.org/web/2020/https://www.cs.umd.edu/~basili/publications/journals/J81.pdf), CAAFE (https://arxiv.org/abs/2305.03403), Shwartz-Ziv & Armon (https://arxiv.org/abs/2106.03253), Grinsztajn et al. (https://arxiv.org/abs/2207.08815), Sculley et al. hidden technical debt (https://papers.nips.cc/paper/5656-hidden-technical-debt-in-machine-learning-systems.pdf), and internal post cross-links.
- Updated post 2026-09-23-llm-at-runtime-is-an-antipattern.md: tightened for coherence and focus, from 18 sections to 11 (merged the pattern section with the marginal-cost argument, the two decision-model sections into "The Waypoint", the pipeline and hash sections into "The Wall", and "The Point" with "A Call to the Community"), dropped the standalone "Broader Trend" list into a closing paragraph, and cut ~1,500 words of main body plus a handful of orphaned references; the three TypeSafe figures, the Doom/Jevons economics, the Laya ceilings, and the reference apparatus are unchanged.

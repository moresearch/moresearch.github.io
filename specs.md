# Specs (talks: minimal generated index + slideshow)

This documents the new talks workflow and validation.

Behavioral rules:
- /talks/ reuses the homepage shell (fonts, header/logo, spacing) but renders a generated list of talk cards.
- The blog homepage includes a global, icon-only "Talks" link in the top-right header which points to /talks/ (aria-label and title must be set).
- The blog homepage also shows an icon-only "Talks" link immediately above the post list in the content column for quick discovery; use the existing .blog-nav-icon style and include aria-label and title for accessibility.
- Each talk card shows only: date/timestamp, title, first-page preview image ({{id}}-page-1.webp), and icon-only actions: open PDF, download PDF, fullscreen presentation.
- No descriptive text under each talk card.
- Fullscreen opens a presentation viewer that uses generated per-page images ({{id}}-page-N.webp) and supports keyboard navigation (Space/Right, Left, Escape).

Build rules:
- Makefile target `previews` must generate for every talks/{{id}}.pdf:
  - preview images: talks/{{id}}-page-1.webp, talks/{{id}}-page-2.webp, ... (one per PDF page) — these should be sized ~1200–1600px wide for fast page load and encoded with WebP quality 75–85
  - fullscreen images: talks/{{id}}-page-1-full.webp, talks/{{id}}-page-2-full.webp, ... — high-resolution images sized ~2560–3200px (or generated at ~300 DPI) and encoded as WebP quality 90–95 for sharp fullscreen rendering
  - talks/{{id}}.meta.json with fields: id, title, date, page_count
- The site generator or Makefile should ensure these assets are present in the repo for GitHub Pages.

Validation:
- /talks/ shows only the minimal cards (no descriptive paragraphs)
- Clicking fullscreen opens the viewer, Space/ArrowRight advances slide, ArrowLeft goes back, Escape exits
- Open/download icons work
- No horizontal overflow on desktop/mobile

Change log:
- Blog posts: Each post in posts/ must start with YAML front matter (---) with required fields `title` and `date` (YYYY-MM-DD). The generator fails on missing or mis-cased keys; keep keys lowercase. See posts/specs.md for details.

- Added talk 003 (Harness Engineering): generated PDF and preview assets must be present at talks/003.pdf and talks/003-page-1.webp and referenced from /talks/.
- Added post 2026-05-17-task-harness-engineering.md: "Task Harness Engineering" describing Task vs Eval vs Agent harnesses (source: YouTube C_GG5g38vLU).
- Updated talk 003 (Harness Engineering): appended full content of harness-related blog posts into slides and added a References slide linking to the posts (harness-engineering-best-practices-for-ai-agents, coding-agent-harnesses-need-real-repositories, harnesses-need-real-browsers-not-polite-demos, agent-harnesses-need-tasks-that-fight-back, good-harnesses-watch-the-whole-operating-system, go-is-good-for-harness-pipelines, task-harness-engineering).
- Added post 2026-05-17-harness-engineering-fowler.md: "Harness Engineering (Martin Fowler)" summarizing and reflecting on Martin Fowler's article (https://martinfowler.com/articles/harness-engineering.html).
- Added post 2026-05-20-self-improving-agent-workflows.md: "Self‑Improving Agent Workflows (SEP + OpenDev + architecture study)" — mixes SEP's skill-loop approach with OpenDev's terminal harness and an empirical design-space survey. References: https://sep.com/blog/the-workflow-that-teaches-itself-a-self-improving-agent-workflow/, https://arxiv.org/html/2603.05344v1, https://arxiv.org/html/2604.18071v1
- Added post 2026-05-20-process-mining-with-skillos.md: "Process mining + SkillOS" — synthesizes process‑mining pipelines with automatic skill curation (SkillOS thread). References: https://hussamalhumsi-21111.medium.com/process-mining-with-python-6ca1d733b3e6, https://medium.com/@i.cemozcelik/solving-a-real-world-data-science-tasks-with-python-c43aa7d654d1, https://x.com/neural_avb/status/2053873358853591435?s=20

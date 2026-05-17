# Specs (talks: minimal generated index + slideshow)

This documents the new talks workflow and validation.

Behavioral rules:
- /talks/ reuses the homepage shell (fonts, header/logo, spacing) but renders a generated list of talk cards.
- The blog homepage includes a global, icon-only "Talks" link in the top-right header which points to /talks/ (aria-label and title must be set).
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
- Added talk 003 (Harness Engineering): generated PDF and preview assets must be present at talks/003.pdf and talks/003-page-1.webp and referenced from /talks/.
- Added post 2026-05-17-task-harness-engineering.md: "Task Harness Engineering" describing Task vs Eval vs Agent harnesses (source: YouTube C_GG5g38vLU).
- Updated talk 003 (Harness Engineering): appended full content of harness-related blog posts into slides and added a References slide linking to the posts (harness-engineering-best-practices-for-ai-agents, coding-agent-harnesses-need-real-repositories, harnesses-need-real-browsers-not-polite-demos, agent-harnesses-need-tasks-that-fight-back, good-harnesses-watch-the-whole-operating-system, go-is-good-for-harness-pipelines, task-harness-engineering).

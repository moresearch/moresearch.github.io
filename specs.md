# Specs (talks: minimal generated index + slideshow)

This documents the new talks workflow and validation.

Behavioral rules:
- /talks/ reuses the homepage shell (fonts, header/logo, spacing) but renders a generated list of talk cards.
- Each talk card shows only: date/timestamp, title, first-page preview image ({{id}}-page-1.webp), and icon-only actions: open PDF, download PDF, fullscreen presentation.
- No descriptive text under each talk card.
- Fullscreen opens a presentation viewer that uses generated per-page images ({{id}}-page-N.webp) and supports keyboard navigation (Space/Right, Left, Escape).

Build rules:
- Makefile target `previews` must generate for every talks/{{id}}.pdf:
  - talks/{{id}}-page-1.webp, talks/{{id}}-page-2.webp, ... (one per PDF page)
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

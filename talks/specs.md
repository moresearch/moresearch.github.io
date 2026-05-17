# Talks specs (generated)

- /talks/ is a generated index of talks produced by the Makefile from talks/*.pdf and metadata extraction.
- The blog homepage must include a global, icon-only "Talks" link in the top-right header which points to /talks/ (aria-label and title must be set). The link must use the Font Awesome Regular rectangle-list icon.
- For each talk PDF the following must be generated and tracked:
  - {{id}}.pdf
  - {{id}}-page-1.webp, {{id}}-page-2.webp, ...
  - {{id}}.meta.json with keys: id (string), title (string), date (ISO8601 string), page_count (integer)
- The index page lists each talk as a card with: time, title (h2), first-page preview image, and three icon-only actions (open, download, fullscreen).
- Fullscreen viewer is an overlay that consumes keyboard events while active, uses the low-resolution {{id}}-page-1.webp for the in-page preview, and uses high-resolution fullscreen images {{id}}-page-<n>-full.webp for fullscreen presentation (recommended: 2560–3200px width or ~300 DPI, WebP quality 90–95).
- Font Awesome used only for icons; do not change typography or the homepage logo implementation.

Validation checklist:
- talks/*.meta.json exists and contains correct page_count
- talks/{{id}}-page-1.webp exists and is referenced by index.html
- Fullscreen navigation works across common browsers

Change log:
- Added talk 003 (Harness Engineering): metadata and page images generated during build.

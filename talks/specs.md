# Talks specs (generated)

- /talks/ is a generated index of talks produced by the Makefile from talks/*.pdf and metadata extraction.
- For each talk PDF the following must be generated and tracked:
  - {{id}}.pdf
  - {{id}}-page-1.webp, {{id}}-page-2.webp, ...
  - {{id}}.meta.json with keys: id (string), title (string), date (ISO8601 string), page_count (integer)
- The index page lists each talk as a card with: time, title (h2), first-page preview image, and three icon-only actions (open, download, fullscreen).
- Fullscreen viewer is an overlay that consumes keyboard events while active and uses the -page-N.webp images for navigation.
- Font Awesome used only for icons; do not change typography or the homepage logo implementation.

Validation checklist:
- talks/*.meta.json exists and contains correct page_count
- talks/{{id}}-page-1.webp exists and is referenced by index.html
- Fullscreen navigation works across common browsers

Change log:
- Added talk 003 (Harness Engineering): metadata and page images generated during build.

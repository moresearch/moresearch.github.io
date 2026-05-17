# Talks template (PDF-first)

This canonical spec describes the PDF-first talks viewer used to generate /talks/NNN/ pages.

Design
- Two-column layout on desktop: left sidebar 300px with branding/meta/actions; right column is the PDF viewer and visually dominant.
- Mobile stacks: branding and actions first, then PDF viewer full width.

Implementation
- Template file: talks/template.html. Generated pages are made by replacing {{ID}} with the zero-padded talk id and copying the PDF to /talks/NNN/NNN.pdf.
- Sidebar must include a real <img> for the logo. Use https://hackspree.com/logo.png if not copying the asset locally.
- Right column must use `min-width: 0` so the iframe cannot cause horizontal overflow.
- Viewer sizing: `.talk-viewer { height: calc(100dvh - 96px); }` and the iframe fills the container.

Validation
- Logo image request returns 200
- PDF request returns 200
- Viewer width equals right-column width; no horizontal scrolling
- Mobile stacks correctly

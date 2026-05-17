# Talks (single-preview)

This page documents the current /talks/ behavior: the page reuses the blog homepage shell (fonts, logo, layout) and displays a single presentation preview card (Talk 001).

Requirements:
- Reuse exact homepage logo markup and visual treatment.
- Show only one presentation card with first-page preview image (./001-page-1.webp).
- Icon-only controls: Open PDF (opens ./001.pdf), Download PDF (downloads ./001.pdf), Fullscreen preview (requestFullscreen on #talk-preview-frame).
- Do not include a Back button; logo links to https://hackspree.com/.

Validation:
- /talks/ renders visually similar to / with identical logo appearance.
- Preview image and PDF return 200.
- Fullscreen hides controls and expands the preview.
- Mobile stacks cleanly and no horizontal overflow.

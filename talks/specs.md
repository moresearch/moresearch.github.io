# Talks template (first-page preview)

This spec describes a simplified talks page showing a first-page preview image instead of an embedded PDF viewer.

Behavior
- Two-column layout: left column (220px) shows only the centered Hackspree logo; right column shows a first-page preview image and icon-only controls (Back/Open/Download).
- No iframe is used for preview; the full PDF remains available via the open/download actions.

Implementation
- Template: talks/template.html. Replace {{ID}} with zero-padded talk id when generating pages.
- Build: generate a first-page PNG named `{{ID}}-page-1.png` into each /talks/NNN/ directory using pdftoppm (or ImageMagick convert as fallback).
- Font Awesome may be loaded from CDN for icons only. Do not use Font Awesome as a page font.
- Accessibility: all icon-only controls must include aria-label and title attributes.

Validation
- Logo, preview image, and PDF must return 200 in the network tab.
- No horizontal overflow; mobile stacks cleanly.

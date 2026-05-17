# Talks template (first-page preview)

This spec describes a simplified talks page showing a first-page preview image instead of an embedded PDF viewer.

Behavior
- Two-column layout: left column (220px) shows only the centered Hackspree logo (using the homepage logo markup/CSS); right column shows a first-page preview image and icon-only controls (Open/Fullscreen/Download).
- No iframe is used for preview; the full PDF remains available via the open/download actions.

Implementation
- Template: talks/template.html. Replace {{ID}} with zero-padded talk id when generating pages. For the consolidated talks index (single talk), the generated page is /talks/index.html referencing ./001.pdf and ./001-page-1.webp.
- Build: generate a first-page WebP (preferred) or PNG named `{{ID}}-page-1.webp` or `{{ID}}-page-1.png` into each /talks/NNN/ directory using pdftoppm + cwebp or ImageMagick convert as a fallback. Ensure the generated asset path resolves relative to /talks/.
- Fullscreen: the Fullscreen control must call requestFullscreen() on the `.talk-preview` container (fallback to `.preview-link`) and the stylesheet must hide `.talk-actions` when the container is in fullscreen using :fullscreen and vendor-prefixed pseudo-classes. Double-clicking the preview image should also enter fullscreen.
- Font Awesome may be loaded from CDN or self-hosted for icons only. Do not use Font Awesome as a page font.
- Accessibility: all icon-only controls must include aria-label and title attributes.

Validation
- Logo, preview image, and PDF must return 200 in the network tab.
- Entering fullscreen must hide action controls (verify .talk-actions display becomes none when .talk-preview is fullscreen).
- No horizontal overflow; mobile stacks cleanly.

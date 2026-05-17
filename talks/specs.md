# Talks template (embedded PDF viewer)

This spec describes a simplified talks page showing an embedded PDF viewer for multiple presentations rather than first-page preview images.

Behavior
- Two-column layout: left column (220px) shows only the centered Hackspree logo (using the homepage logo markup/CSS); right column shows an embedded PDF viewer (iframe) and icon-only controls (Open/Fullscreen/Download) plus lightweight talk-selection controls (e.g., 001, 002).
- No preview images are used. The viewer loads PDFs in-place from `/talks/NNN.pdf`.

Implementation
- Template: talks/index.html is the canonical viewer page for the consolidated talks host. The page must reference `./001.pdf` (default) and `./002.pdf` and provide controls to switch the iframe src between them.
- Fullscreen: the Fullscreen control must call `requestFullscreen()` on the `.talk-preview` container and the stylesheet must hide `.talk-actions` when that container is in fullscreen using `:fullscreen` and vendor-prefixed pseudo-classes.
- Controls: talk-selection buttons update the iframe `src`, update the Open and Download button `href` values, and reflect selection with `aria-pressed`.
- Font Awesome may be loaded from CDN or self-hosted for icons only. Do not use Font Awesome as a page font.
- Accessibility: all icon-only controls must include `aria-label` and `title` attributes.

Validation
- `/talks/001.pdf` and `/talks/002.pdf` must return 200 in network tab.
- The iframe should display the selected PDF; entering fullscreen must hide `.talk-actions` and expand the viewer.
- No horizontal overflow; layout stacks cleanly on narrow viewports.

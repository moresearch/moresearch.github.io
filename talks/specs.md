# Talks template specs

This file documents the canonical talks viewer template and constraints for files under /talks/.

User-facing behavior
- Two-column layout on desktop: left sidebar (360px) with branding and actions, right content column minmax(0,1fr) with the PDF viewer.
- Mobile stacks into a single column: logo on top, actions, then viewer.
- Logo must be a real <img> and must match the homepage visual treatment. Prefer using `https://hackspree.com/logo.png` unless the asset is copied into the blog build output.

Implementation rules
- Template file: talks/template.html (canonical). Generated pages are created by replacing `{{ID}}` with zero-padded talk ID and copying the corresponding PDF to /talks/NNN/NNN.pdf.
- CSS layout: use CSS Grid for two columns. Left column fixed at 360px, right column `minmax(0,1fr)` and *must* include `min-width:0` to avoid overflow.
- Logo: use an explicit <img> element with `alt="Hackspree"`, link to `https://hackspree.com/`, and copied homepage `.logo` CSS rules for parity. Do not use background-image for the logo.
- PDF viewer: `.pdf-container { width:100%; max-width:100%; aspect-ratio:210/297; max-height: calc(100dvh - 160px); overflow:hidden; }` and the iframe/object must fill the container width:100% height:100%.
- Avoid `overflow:hidden` on ancestor containers that could clip the rotated logo. Prefer `overflow:visible` on `.wrap` or top-level containers.

Accessibility and UX
- Action controls: Back (link to /talks/), Download (direct link to PDF), Fullscreen (requestFullscreen on the viewer container). Icons-only allowed with sr-only labels for screen readers.
- Responsive rules: at max-width:900px the layout becomes single column and the logo stops being rotated and uses `max-width:252px; max-height:206px`.

Validation checklist
- Logo image request returns 200 (network tab)
- PDF file request returns 200
- Viewer computed width equals the right column width (no horizontal overflow)
- Mobile stacks correctly and fullscreen still works



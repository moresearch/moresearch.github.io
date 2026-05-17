# Talks template (masthead)

This spec defines the compact masthead talks viewer.

Design
- Compact top masthead (logo + talk label + actions), PDF viewer below as the dominant element.
- Dark, minimal aesthetic matching hackspree.com.

Implementation
- Template: talks/template.html. Replace {{ID}} with zero-padded talk id when generating pages.
- Header contains a real <img> logo (prefer https://hackspree.com/logo.png), small talk label, title, and actions (Back / Download / Fullscreen).
- No large sidebar; viewer width is controlled by container and uses `min-width:0` to prevent overflow.
- Viewer: `.talk-viewer { height: calc(100dvh - var(--header-h) - 48px); }` and iframe fills it.

Validation
- Logo loads (200), PDF loads (200), header compact, no horizontal overflow, mobile stacks correctly.

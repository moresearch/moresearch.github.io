# Talks index (presentations listing)

This spec describes the talks index page which lists available presentations in a posts-style layout. Each entry links to the presentation PDF and provides short metadata (date, title, optional summary).

Behavior
- Two-column layout: left column (220px) shows the centered Hackspree logo (homepage logo markup/CSS); right column contains a posts-style list of presentations ordered newest-first.
- Each presentation is represented as an article element with a `post-meta` time, an `h2` title anchor linking to the presentation PDF (`./NNN.pdf`), and a short summary or description.
- Icon-only actions (Open, Fullscreen, Download) are available in the page header or per-presentation entry as needed.

Implementation
- Template: talks/index.html is the canonical index page for presentations and must render a `.post-list` containing `<article class="post">` entries similar to the root blog index.html.
- Build: the Makefile must produce `public/slides/NNN.pdf` and a copy or link `talks/NNN.pdf` so the index can link directly to `./NNN.pdf`. When generating a new talk, ensure `talks/NNN.pdf` is added to the repository so GitHub Pages can serve it.
- Accessibility: presentation entries must include accessible anchors and `time` elements with `datetime` and `data-utc` attributes.

Validation
- `/talks/001.pdf` must return 200 in the network tab.
- The talks index renders a single preview card for talk 001 with an inline PDF preview and action icons (Open, Fullscreen, Download). The post-list should include at least the article entry linking to `./001.pdf` with valid metadata.
- The page reuses the blog homepage fonts and logo slot, includes Font Awesome for icons, and has a small fullscreen JS helper to present the preview.
- Mobile stacks cleanly and no horizontal overflow occurs.

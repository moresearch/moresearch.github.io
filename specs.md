# Specs

## User-facing behavior
- Every blog post must focus on a single, well-explained idea and be well-referenced. Posts covering multiple ideas must be split or refactored to ensure clarity and value. This applies to all published posts.
- When posts reference media (articles, YouTube videos, movie scenes), include explicit links, timestamps, and a short explanation of how the referenced scene supports the post's argument. For YouTube videos, include the thumbnail image linked to the video and an optional timestamped note (MM:SS).

- All post titles must follow the template: "{context}: {core idea}". The context sets the domain, technology, or scenario; the core idea states the main insight or contribution. This applies to all published posts, and titles must be updated if requirements change.
- All post summaries must be approximately 3x longer than before, written in an abstract style similar to academic papers but with a lighter, more casual tone. Summaries must clearly state the problem, approach, and key takeaways, providing enough context for a reader to understand the post's value without reading the full article. This applies to all published posts, and summaries must be updated if requirements change.

- Presentations can be authored as LaTeX (.tex) files using a standard beamer template. The Makefile compiles these to PDFs, which are placed in a public directory as zero-padded files (e.g., 001.pdf, 002.pdf).
- The site serves these PDFs at /slides/NNN (e.g., /slides/001), displaying them inline in the browser. These links are not listed in navigation or post listings.
- Only PDFs with matching NNN are served; others return 404. The LaTeX template is standardized for all presentations.

### Blog Post Titles (shortened and improved)
- Platform Engineering: Scale vs Speed
- Go for Mobile LLM Control Planes
- Why Go Still Matters in AI
- Edge LLMs: Model Shape Matters
- Harnesses Need Real Browsers
- Studying Local LLM Training
- Go for Harness Pipelines
- Empirical Game Theory for Agents
- SWE-Agent Economics: My Focus
- Dependency Injection in Go
- Factorio & SC2: Systems Thinking
- Mechanism Design for Agentic Markets
- On-Device LLMs: Systems Design
- Codebases for Agentic Engineering
- Go for Disaggregated Serving
- Agentic Era: An Economic System
- Coding Harnesses Need Real Repos
- Harnesses Need Tasks That Fight Back
- Blockchains for Agentic Software
- Harness Engineering Best Practices
- Edge NPUs Change Serving Shape
- Harnesses Should Observe the OS
- Go for Structured LLM Runtimes
- 1-Bit Models: Edge Budgets
- Startup Lessons: The Two-Task Rule (summary: Startups face a barrage of advice, but few lessons are as universally relevant—and as frequently ignored—as the two-task rule. This post unpacks why the discipline of focusing on only two critical tasks at any given time is the single most important operating principle for early-stage companies. Drawing on Y Combinator’s "The Hardest Lessons for Startups to Learn," we explore the psychological traps that lead founders to overcommit, the operational chaos that results, and the compounding benefits of ruthless prioritization. Readers will gain a practical framework for applying the two-task rule, understand its impact on execution and morale, and learn how to resist the temptation to chase every opportunity at once.)

- Platform Engineering: Scale vs Speed (summary: Nuanced, actionable exploration of balancing scale and speed in platform engineering, with frameworks and real-world examples. Based on https://www.youtube.com/watch?v=iP-qzK4mQuI and https://www.youtube.com/watch?v=5Ai8UGx7QvQ)

- The GitHub Pages site is reachable at `https://moresearch.github.io/`.
- The GitHub Pages site is also published at `https://blog.hackspree.com/`.
- The published site is a single-page blog at the site root, served from `index.html`.
- The site must include a favicon using the provided favicon assets in the `favicon/` directory. The generated index.html must include explicit <link rel="icon" href="/favicon/favicon-32x32.png" sizes="32x32">, <link rel="icon" href="/favicon/favicon-16x16.png" sizes="16x16">, <link rel="apple-touch-icon" href="/favicon/apple-touch-icon.png">, <link rel="manifest" href="/favicon/site.webmanifest"> and a fallback <link rel="shortcut icon" href="/favicon/favicon.ico">. These tags must reference the files under `/favicon/` so browsers and platforms find the correct assets. Additionally, place a `favicon.ico` at the repository root (`/favicon.ico`) as an extra fallback because some clients and crawlers check the site root for the icon. All major favicon sizes and the manifest should be present.
- The page presents all published posts on one scrollable page with stable direct anchors per post.
- The visual language follows `hackspree.com`: dark background, high-contrast text, Orbitron typography, uppercase labels, and restrained muted secondary text.
- The page should feel like `hackspree.com` with a three-column layout: logo rail, body column, and posts navigation rail.
- On larger screens, the left logo rail and the right posts navigation rail both stay fixed in the vertical center while the body column scrolls.
- The visible typography must be smaller than the current Hackspree landing page while remaining comfortably readable, so more blog content fits on screen.
- The rotated Hackspree logo, its effective left-page margin, and the overall desktop/mobile proportions must stay visually close enough to `hackspree.com` that moving between the landing page and blog feels like a seamless transition.
- The primary visible heading inside the right-side navigation rail is only `Engineering blog`, rendered in a smaller treatment with no subtitle.
- The right-side navigation list must not show a `Posts` label, must use smaller link text with clear separators between entries, link to the in-page post anchors, and keep its vertical scrollbar visually hidden until the user hovers the navigation rail.
- As the user scrolls through posts, the matching navigation entry should become visibly active and the navigation rail should automatically keep that active entry near its vertical center.
- Post titles must be substantially larger than the body copy so each article headline stands out clearly in the dense layout.
- Standard post body copy should use a straighter, less rounded professional monospace face than Space Mono and render in muted dark gray rather than pure white, while important emphasis and headings keep bright contrast against the black background.
- The post typography should stay compact overall, but the body copy should be slightly larger, lighter, and more generously spaced for quick readability.
- Post summaries should use a more polished, readable presentation than the metadata labels, with calmer spacing and contrast.
- Post summaries must use a distinct font from the main body and headings. Use 'Sudo' (a tech-inspired monospace font) for summaries to visually separate them, while keeping the Hackspree style. Fallback to 'IBM Plex Mono', monospace if Sudo is unavailable.
- Published post metadata should display the post date as a larger Unix timestamp, and hovering it should reveal the normal UTC date in a custom tooltip that appears to the right of the timestamp.
- The archive chronology is intentionally driven by post dates, and the published mix may include medium-length engineering essays rather than only short notes.
- Posts must be listed in descending date order (most recent first) on the blog page.
- The published archive may include short research-note posts that cite external papers and link to their source pages directly from the post body.
- The published archive may also include first-person technical reflection posts on workshops, talks, or videos, as long as they stay grounded in concrete implementation details rather than generic impressions.
- The published archive may also include first-person essays about how games, habits, or other non-code systems shaped engineering skill, as long as they connect back to concrete technical practice.
- Posts may include fenced code examples with syntax highlighting, and the code presentation should use a 246_noir-inspired dark palette that fits the Hackspree aesthetic.
- Posts may include remote publisher-hosted book-cover images when they materially support the essay content.
- Links rendered inside post bodies must open in a separate browser tab/window by default with safe external-link rel attributes.
- Links rendered inside post bodies should stay minimal and readable without visible underlines.
- New posts become visible on the page after adding a Markdown source file under `posts/` and rebuilding the site with `make`.

## Implementation rules

- The root Makefile must build all .tex files in a presentations/ directory into PDFs using a standard beamer template, outputting them as zero-padded files (e.g., 001.pdf, 002.pdf) into a public/slides/ directory.
- The root Makefile must include a 'talks' target that depends on 'slides' and generates per-talk viewer pages. The target should create /talks/NNN/, copy public/slides/NNN.pdf to /talks/NNN/NNN.pdf (or ensure the iframe path resolves correctly), and write /talks/NNN/index.html by replacing a {{ID}} placeholder in talks/template.html with the zero-padded talk ID (NNN). Generated talk pages and PDFs must be committed to the repository so GitHub Pages can serve them.
- Talks template requirements:
  - Use a two-column layout: left logo rail (360px) and right content column (minmax(0, 1fr) containing the PDF viewer and controls).
  - Left rail shows the Hackspree logo using an absolute URL (https://hackspree.com/logo.png) inside a dark background box to ensure visibility.
  - Action icons (Back to Talks, Download, Fullscreen) must be in the right column, centered horizontally, anchored near the bottom of the content column.
  - The PDF viewer must be constrained to show a single page at a time and must occupy the full width of the right column. Use `.pdf-container { width:100%; max-width:100%; aspect-ratio:210/297; max-height: calc(100dvh - 160px); overflow:hidden; }` and ensure the right column uses `min-width:0` so the viewer never forces horizontal overflow.
  - The template should avoid page scrollbars by sizing the viewer to the viewport height and using natural document flow. Avoid `overflow:hidden` on parent containers that may clip the rotated logo. Use `.wrap { min-height: 100dvh; overflow: visible; }`.
  - The template must reuse the homepage header/logo treatment so the logo is visible and links to https://hackspree.com/.
  - Keep controls icon-only and preserve PDF iframe fragment `#page=1&view=FitH&toolbar=0` to attempt single-page rendering.

- The site generator or static file config must serve /talks/NNN/NNN.pdf and /talks/NNN/index.html from the repository root, so that /talks/NNN/NNN.pdf and /talks/NNN/ work on GitHub Pages.
- /talks/NNN/ serves an HTML page with an embedded PDF viewer for NNN.pdf, including a fullscreen button for presentations, so /talks/NNN/ always works as a user-friendly link.
- Every /talks/NNN/ page must include a small, minimally spaced "Talks" link (not "Back to Talks") at the top of the main content, styled with #ff9800 and the correct font, to return to /talks/. The text above and below the PDF viewer must be visually minimal and small (smaller font, less spacing).
- The PDF viewer on /talks/NNN/ must be visually dominant within the second (right) column. The viewer container must present a single PDF page (page 1) sized to the page aspect ratio (210/297) and scaled to fit. Implement this with `.pdf-container` using `aspect-ratio: 210/297` and `max-height: 36vh`. The embedded viewer (iframe/object) must fill the container (`width:100%; height:100%`) and have scrolling disabled or overflow hidden so only one page is visible at a time. The viewer must never overflow or stretch beyond its column on any screen size. All non-PDF content (navigation, fallback, etc.) must use a smaller font and minimal spacing.
- The fullscreen button must be a small symbol (⛶) only, with no text label, and visually subtle (smaller, less prominent than before). The fullscreen button must only fullscreen the PDF viewer container, not the entire page. All action buttons (download, back, fullscreen) must be icons only, with no visible text labels such as 'Download PDF' or 'Talks'.
- The main talk links in /talks/index.html must use a different font from the rest of the page, chosen from the set of site fonts (Orbitron, IBM Plex Mono, Sudo Var), to visually distinguish them. Font size and style must match hackspree.com: compact, uppercase, and visually consistent. List items and headings must use smaller, denser font sizes than before, matching the main blog and landing page.
- /talks/ is an index page listing all available talks with links to each viewer. The Hackspree logo must be displayed only in the left rail, in the same position and style as the main blog. No other logo should appear anywhere on the page.
- Extensionless URLs like /talks/NNN/ are supported via an index.html viewer in each /talks/NNN/ directory.
- Each /talks/NNN/ viewer page must show the action icons (download, back, fullscreen) in the left logo rail, arranged horizontally in their own row beneath the logo (not overlapping). The left rail must reserve fixed space for the rotated logo (use `.logo-box`) so icons are placed on a separate row. Icons should be small (≈36×36px), centered beneath the logo with at least 32px vertical spacing, and must be icons-only (no visible text). Provide inline SVG fallbacks; do not display text labels.
- Use talks/template.html as the canonical template for all /talks/NNN/ viewer pages. The template implements a two-column layout (left logo rail, right content column), places the action icons (back, fullscreen, download) directly beneath the PDF viewer in the content column, and constrains the embedded PDF to show a single page (aspect-ratio: 210/297, iframe filling the container, scrollbar and toolbar disabled). For each talk, create /talks/NNN/index.html as a copy of talks/template.html with the placeholder ID replaced by NNN and the PDF file set to NNN.pdf. The layout, logo rail, and visual identity must otherwise match blog.hackspree.com and hackspree.com precisely.
- All links in /talks/index.html and /talks/NNN/index.html must use #fff as the primary color and the same font as the main blog.
- The logo in /talks/NNN/ must be visually identical in position, size, and style to hackspree.com, including rotation, margin, and centering. Concretely, copy the `.logo` CSS rules from the site root `index.html` (e.g., `max-width: 156vh; max-height: 117%; transform: rotate(-90deg) scale(1.15);`) so the rendered size matches the canonical site. The HTML and CSS for the logo rail must match `index.html` exactly.
- No dev comments or stray CSS should appear in the output HTML of /talks/NNN/.
- The color theme for all talks pages must be white/grey on black, matching hackspree.com, with no orange or other accent colors except where specified for navigation or icons.
- The workflow for adding a new presentation: add a .tex file to presentations/, run `make slides && make talks`, verify the generated `/talks/NNN/index.html` and `/talks/NNN/NNN.pdf`, then commit the generated files and directories to the repository so GitHub Pages serves them.
- The LaTeX template for presentations must be standardized and documented in the repo.
- /slides/NNN links must not appear in navigation or post listings.

- The repository root must contain a `CNAME` file with the custom domain hostname.
- The `CNAME` file must contain exactly `blog.hackspree.com`.
- The repository root must contain the generated `index.html` that GitHub Pages serves directly.
- The root `Makefile` must build the single-page blog from Markdown files in `posts/`.
- The local build pipeline must use the Go toolchain only.
- The build flow must be deterministic: posts are rendered in descending date order and repeated builds without source changes must not change the output.
- The generated page must keep the site styling self-contained inside `index.html`, except for remote font or icon stylesheets required to match `hackspree.com`.
- The generated page may load remote syntax-highlighting assets when needed for post code examples.
- The syntax-highlighting palette should be defined once in reusable shared tokens so every current and future post inherits the same code styling.
- Code snippets must stay visually close to `vim-256noir`: black background with a monochrome black/gray/white token palette, including grayscale numerals instead of accent colors.
- Code snippets must use Sudo as the dedicated code font, rendered at a sane slightly-larger pixel size so glyphs stay clear without feeling cramped.
- Code-snippet horizontal scrollbars must stay visually hidden until the reader hovers the snippet.
- Go code snippets should include short comments that clarify the example logic instead of presenting uncommented bare code.
- Posts that cite a YouTube video should include that video's thumbnail image in the post body, linked to the video.
- The generated page may use a small inline enhancement script for active-post navigation behavior, but the page must remain usable when JavaScript is unavailable.
- Navigation-centering behavior must not interfere with the main page scroll direction or pull the reader upward while they are scrolling down the post column.
- The build must fail when required post metadata is missing or when two posts resolve to the same slug.
- The repository must include `.nojekyll` so GitHub Pages serves the generated output without Jekyll processing.
- The repository root must contain the Hackspree logo asset used by the fixed side rail.
- The generated page must not render a helper intro card or an on-page index section above the posts.

## Constraints and invariants

- The `CNAME` file must stay at the published site root.
- The custom domain in `CNAME` must match the DNS configuration managed in the Cloudflare repo.
- GitHub Pages does not run the repository `Makefile`; generated artifacts required for publishing must therefore be committed to the repository.
- Every published post must have a stable in-page anchor derived from its slug.
- The blog page must remain functional with JavaScript disabled.
- The desktop fixed rails must gracefully collapse into a stacked top section on narrower screens.
89. - All talk links in /talks/index.html must reliably navigate to their respective /talks/NNN/ pages using absolute paths (e.g., /talks/001/). No link should ever redirect to the homepage.

### Talks: PDF-first redesign
- Updated talks template to a PDF-first two-column layout (300px sidebar + dominant viewer).
- Sidebar contains logo (<img>), talk metadata, and controls (Talks / Download / Fullscreen).
- See talks/specs.md for canonical implementation details and validation checklist.

### Talks: compact masthead layout
- Talks pages use a compact top masthead (logo + metadata + actions) with the PDF viewer as the page's primary content. See talks/specs.md for details.

### Talks: strict two-column layout
- Talks pages now use a strict two-column layout with a left identity rail and a dominant PDF viewer. See talks/specs.md for details and validation checklist.

### Talks: simplified first-page preview
- Talks pages now show a centered Hackspree logo in a left rail and a first-page preview image on the right. See talks/specs.md for details.
- Fullscreen behavior: the Fullscreen control should requestFullscreen() on the talk preview container ('.talk-preview'), hide action controls using :fullscreen CSS, and support double-click-to-present on the preview image.

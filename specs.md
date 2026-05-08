# Specs

## User-facing behavior

- The GitHub Pages site is reachable at `https://moresearch.github.io/`.
- The GitHub Pages site is also published at `https://blog.hackspree.com/`.
- The published site is a single-page blog at the site root, served from `index.html`.
- The page presents all published posts on one scrollable page with stable direct anchors per post.
- The visual language follows `hackspree.com`: dark background, high-contrast text, Orbitron typography, uppercase labels, and restrained muted secondary text.
- The page should feel like `hackspree.com` with a three-column layout: logo rail, body column, and posts navigation rail.
- On larger screens, the left logo rail and the right posts navigation rail both stay fixed in the vertical center while the body column scrolls.
- The visible typography must be smaller than the current Hackspree landing page while remaining comfortably readable, so more blog content fits on screen.
- The rotated Hackspree logo must match the sizing treatment used by `hackspree.com`.
- The primary visible `Engineering Blog` title and subtitle live inside the right-side navigation rail in a smaller treatment than before.
- The `Posts` navigation list must appear in its own right-side rail, use smaller link text with clear separators between entries, link to the in-page post anchors, and scroll vertically without introducing a horizontal scrollbar.
- The post typography should be compressed aggressively relative to the current version to maximize visible content while staying legible.
- New posts become visible on the page after adding a Markdown source file under `posts/` and rebuilding the site with `make`.

## Implementation rules

- The repository root must contain a `CNAME` file with the custom domain hostname.
- The `CNAME` file must contain exactly `blog.hackspree.com`.
- The repository root must contain the generated `index.html` that GitHub Pages serves directly.
- The root `Makefile` must build the single-page blog from Markdown files in `posts/`.
- The local build pipeline must use the Go toolchain only.
- The build flow must be deterministic: posts are rendered in descending date order and repeated builds without source changes must not change the output.
- The generated page must keep the site styling self-contained inside `index.html`, except for remote font or icon stylesheets required to match `hackspree.com`.
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

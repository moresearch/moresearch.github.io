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
- The rotated Hackspree logo and overall desktop/mobile proportions must stay visually close enough to `hackspree.com` that moving between the landing page and blog feels like a seamless transition.
- The primary visible heading inside the right-side navigation rail is only `Engineering blog`, rendered in a smaller treatment with no subtitle.
- The right-side navigation list must not show a `Posts` label, must use smaller link text with clear separators between entries, link to the in-page post anchors, and keep its vertical scrollbar visually hidden until the user hovers the navigation rail.
- As the user scrolls through posts, the matching navigation entry should become visibly active and the navigation rail should automatically keep that active entry near its vertical center.
- Post titles must be substantially larger than the body copy so each article headline stands out clearly in the dense layout.
- Standard post body copy should use a straighter, less rounded professional monospace face than Space Mono and render in muted dark gray rather than pure white, while important emphasis and headings keep bright contrast against the black background.
- The post typography should stay compact overall, but the body copy should remain slightly enlarged for readability.
- Post summaries should use a more polished, readable presentation than the metadata labels, with calmer spacing and contrast.
- The published archive may include short research-note posts that cite external papers and link to their source pages directly from the post body.
- Links rendered inside post bodies must open in a separate browser tab/window by default with safe external-link rel attributes.
- New posts become visible on the page after adding a Markdown source file under `posts/` and rebuilding the site with `make`.

## Implementation rules

- The repository root must contain a `CNAME` file with the custom domain hostname.
- The `CNAME` file must contain exactly `blog.hackspree.com`.
- The repository root must contain the generated `index.html` that GitHub Pages serves directly.
- The root `Makefile` must build the single-page blog from Markdown files in `posts/`.
- The local build pipeline must use the Go toolchain only.
- The build flow must be deterministic: posts are rendered in descending date order and repeated builds without source changes must not change the output.
- The generated page must keep the site styling self-contained inside `index.html`, except for remote font or icon stylesheets required to match `hackspree.com`.
- The generated page may use a small inline enhancement script for active-post navigation behavior, but the page must remain usable when JavaScript is unavailable.
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

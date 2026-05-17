# Specs (talks single-preview)

/ talks / now renders a single presentation preview reusing the homepage shell. This is an intentional design choice: the talks route is a presentation page rather than an index.

- The site must keep the homepage fonts and layout (Orbitron, Sudo Var, IBM Plex Mono) for brand parity.
- Talk assets must be present at /talks/NNN.pdf and /talks/NNN-page-1.webp when referenced.
- The build pipeline (Makefile) should generate and place the preview image under /talks/.
- The page must not include a left logo rail that differs from the homepage; instead it should reuse the homepage logo slot exactly.

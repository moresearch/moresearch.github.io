# Presentations

- Place LaTeX .tex files for presentations in this directory.
- Use the standard beamer template (see template.tex) for all presentations.
- Each .tex file will be compiled to a PDF with a zero-padded filename (e.g., 001.pdf, 002.pdf) and output to public/slides/.
- To add a new presentation:
  1. Copy template.tex to NNN.tex (e.g., 003.tex).
  2. Edit your content.
  3. Run `make slides` from the repo root.
  4. Commit the resulting PDF in public/slides/.

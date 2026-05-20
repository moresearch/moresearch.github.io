
#LATEX = pdflatex
LATEX = xelatex
D2    = d2
INKSCAPE ?= inkscape
PY    ?= python3

# -----------------------------
# Build units (src/*.tex)
# -----------------------------
DOCS  = se_guidelines se_guidelines_pt
DECKS = presentation

TEX_UNITS = $(DOCS) $(DECKS)

# PDFs (in ./pdfs)
PDFS = $(addprefix pdfs/,$(addsuffix .pdf,$(TEX_UNITS)))

# D2 sources and generated figures
D2SRC = $(wildcard src/d2/*.d2)
D2SVG = $(patsubst src/d2/%.d2,src/figs/%.svg,$(D2SRC))
D2PDF = $(patsubst src/d2/%.d2,src/figs/%.pdf,$(D2SRC))

# -----------------------------
# Assets
# -----------------------------
LOGO_LOCAL = src/figs/mda.jpg

.PHONY: all clean d2figs assets docs decks new newdeck figclean

# Keep intermediate SVGs (don’t delete on failure)
.PRECIOUS: src/figs/%.svg

# Default: build assets, D2 figs, and all PDFs
all: assets d2figs $(PDFS)

docs: assets d2figs $(addprefix pdfs/,$(addsuffix .pdf,$(DOCS)))
decks: assets d2figs $(addprefix pdfs/,$(addsuffix .pdf,$(DECKS)))

assets:
	@if [ ! -f "$(LOGO_LOCAL)" ]; then \
		echo "Error: missing logo at $(LOGO_LOCAL)"; \
		exit 1; \
	fi

# Build all D2 figures (PDFs)
d2figs: $(D2PDF)

# -----------------------------
# D2 -> SVG
# -----------------------------
src/figs/%.svg: src/d2/%.d2
	@mkdir -p src/figs
	$(D2) $< $@

# -----------------------------
# SVG -> PDF (robust)
#
# Fixes:
# - D2 SVG sometimes contains garbage after </svg> (same line or extra bytes)
# - Clean SVG by truncating at the LAST </svg> occurrence
# - Export using Inkscape with:
#   - --export-area-drawing (tight crop)
#   - --export-text-to-path (fonts won’t break in PDF/LaTeX)
# -----------------------------
src/figs/%.pdf: src/figs/%.svg
	@set -e; \
	tmp="$@.tmp.svg"; \
	$(PY) -c 'import sys,pathlib; inp=pathlib.Path(sys.argv[1]); out=pathlib.Path(sys.argv[2]); data=inp.read_bytes(); end=b"</svg>"; i=data.rfind(end); \
	(i!=-1) or sys.exit("ERROR: </svg> not found in %s" % inp); out.write_bytes(data[:i+len(end)]+b"\n")' \
	"$<" "$$tmp"; \
	$(INKSCAPE) "$$tmp" --export-area-drawing --export-text-to-path --export-type=pdf --export-filename="$@"; \
	rm -f "$$tmp"

# -----------------------------
# LaTeX
# -----------------------------
pdfs/%.pdf: src/%.tex assets d2figs
	@mkdir -p pdfs build
	cd src && $(LATEX) -interaction=nonstopmode -output-directory ../build $(<F)
	cd src && $(LATEX) -interaction=nonstopmode -output-directory ../build $(<F)
	cp build/$(@F) pdfs/

# Convenience: `make foo.pdf` builds pdfs/foo.pdf and symlinks foo.pdf -> pdfs/foo.pdf
%.pdf: pdfs/%.pdf
	@ln -sf $< $@

# -----------------------------
# Helpers
# -----------------------------
figclean:
	rm -f src/figs/*.svg src/figs/*.pdf

new:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make new NAME=your_doc_name"; \
		exit 1; \
	fi
	@if [ -e "src/$(NAME).tex" ]; then \
		echo "Error: src/$(NAME).tex already exists"; \
		exit 1; \
	fi
	cp src/main.tex src/$(NAME).tex
	@echo "Created src/$(NAME).tex from src/main.tex"

newdeck:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make newdeck NAME=your_deck_name"; \
		exit 1; \
	fi
	@if [ -e "src/$(NAME).tex" ]; then \
		echo "Error: src/$(NAME).tex already exists"; \
		exit 1; \
	fi
	cp src/presentation.tex src/$(NAME).tex
	@echo "Created src/$(NAME).tex from src/presentation.tex"

clean:
	rm -rf build
	rm -f src/figs/*.pdf
	@if [ -d pdfs ]; then find pdfs -type f ! -name '*.pdf' -delete; fi

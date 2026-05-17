GO := go
GOCACHE := $(CURDIR)/.build-cache
POSTS := $(sort $(filter-out posts/specs.md,$(wildcard posts/*.md)))

.PHONY: all build clean rebuild

all: build

build: index.html

slides: 
	mkdir -p public/slides
	for f in presentations/*.tex; do \
	  n=$$(basename $$f .tex); \
	  pdffile=public/slides/$$(printf "%03d.pdf" $$n); \
	  pdflatex -output-directory=public/slides -jobname=$$(printf "%03d" $$n) $$f; \
	done

index.html: build_blog.go go.mod go.sum logo.png $(POSTS)
	GOCACHE=$(GOCACHE) $(GO) run . --input-dir posts --output index.html

clean:
	rm -rf $(GOCACHE)

rebuild: clean build

# generate talks pages from template
talks: slides
	@echo "Generating talks pages from template"
	@mkdir -p talks
	@for f in public/slides/*.pdf; do \
		id=$$(basename $$f .pdf); \
		dir=talks/$$id; \
		mkdir -p $$dir; \
		sed 's/{{ID}}/'"$$id"'/g' talks/template.html > $$dir/index.html; \
		cp -f $$f $$dir/$$id.pdf; \
	done

# generate first-page previews for talks/*.pdf (webp)
previews:
	@echo "Generating talk preview images"
	@for f in talks/*.pdf; do \
		p=$$(basename $$f .pdf); \
		if command -v pdftoppm >/dev/null 2>&1; then \
			pdftoppm -f 1 -l 1 -singlefile -png "$$f" "talks/$$p-page-1"; \
			if command -v cwebp >/dev/null 2>&1; then \
				cwebp -q 80 "talks/$$p-page-1.png" -o "talks/$$p-page-1.webp" >/dev/null 2>&1 || true; \
				rm -f "talks/$$p-page-1.png"; \
			else \
				if command -v convert >/dev/null 2>&1; then \
					convert "talks/$$p-page-1.png" "talks/$$p-page-1.webp" >/dev/null 2>&1 || true; \
					rm -f "talks/$$p-page-1.png"; \
				fi; \
			fi; \
		else \
			echo "pdftoppm not found; skipping $$f"; \
		fi; \
	done

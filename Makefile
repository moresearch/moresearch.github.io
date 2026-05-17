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

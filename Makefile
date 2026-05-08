GO := go
GOCACHE := $(CURDIR)/.build-cache
POSTS := $(sort $(filter-out posts/specs.md,$(wildcard posts/*.md)))

.PHONY: all build clean rebuild

all: build

build: index.html

index.html: build_blog.go go.mod go.sum logo.png $(POSTS)
	GOCACHE=$(GOCACHE) $(GO) run . --input-dir posts --output index.html

clean:
	rm -rf $(GOCACHE)

rebuild: clean build

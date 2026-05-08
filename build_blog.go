package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

const (
	siteTitle       = "Engineering Blog"
	siteURL         = "https://blog.hackspree.com/"
	siteTagline     = "Software engineering notes, experiments, and build logs."
	siteDescription = "Single-page writing log built from Markdown sources."
)

var (
	slugPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	markdowner  = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)
	pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <meta name="color-scheme" content="dark">
  <meta name="description" content="{{.MetaDescription}}">
  <link rel="canonical" href="{{.SiteURL}}">
  <title>{{.PageTitle}}</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Orbitron:wght@500;700&display=swap" rel="stylesheet">
  <style>
    html,
    body {
      margin: 0 !important;
      padding: 0 !important;
      width: 100% !important;
      min-width: 100% !important;
      min-height: 100% !important;
      background: #000 !important;
      color: #fff !important;
    }

    * {
      box-sizing: border-box;
    }

    body {
      position: relative;
      overflow-x: hidden;
      font-family: "Orbitron", Arial, sans-serif;
    }

    a {
      color: inherit;
      text-decoration: none;
    }

    img {
      display: block;
      max-width: 100%;
      height: auto;
    }

    code,
    pre {
      font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
    }

    .bg {
      position: fixed;
      inset: 0;
      background: #000;
      z-index: 0;
    }

    .wrap {
      position: relative;
      z-index: 1;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 32px;
    }

    .layout {
      width: min(1600px, 100%);
      display: grid;
      grid-template-columns: minmax(220px, 260px) minmax(0, 640px) minmax(200px, 240px);
      gap: 40px;
      align-items: start;
      justify-content: center;
    }

    .logo-side {
      display: flex;
      align-items: center;
      justify-content: center;
      overflow: visible;
      padding: 16px 8px 16px 16px;
      box-sizing: border-box;
      position: sticky;
      top: 50vh;
      transform: translateY(-50%);
      align-self: start;
    }

    .logo-rail {
      width: 100%;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
    }

    .logo-link {
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .logo {
      display: block;
      width: auto;
      height: auto;
      max-width: 156vh;
      max-height: 117%;
      transform: rotate(-90deg) scale(1.3);
      transform-origin: center center;
    }

    .content-side {
      display: flex;
      align-items: flex-start;
      justify-content: flex-start;
      padding: 48px 0;
      box-sizing: border-box;
      min-width: 0;
    }

    .content {
      width: 100%;
      margin: 0;
    }

    .nav-side {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 16px 16px 16px 8px;
      box-sizing: border-box;
      position: sticky;
      top: 50vh;
      transform: translateY(-50%);
      align-self: start;
    }

    .post-meta,
    .post-summary,
    .role,
    .tag {
      margin: 0;
      color: #9ca3af;
      font-size: 0.64rem;
      letter-spacing: 0.18em;
      text-transform: uppercase;
    }

    h1,
    h2,
    h3,
    h4 {
      margin: 0;
      font-weight: 600;
      line-height: 1.2;
      text-transform: uppercase;
      letter-spacing: 0.08em;
      color: #fff;
    }

    h1 {
      font-size: clamp(1.35rem, 2.2vw, 1.95rem);
    }

    h2 {
      font-size: clamp(0.92rem, 1.45vw, 1.08rem);
    }

    h3 {
      font-size: 0.92rem;
    }

    .role {
      margin: 10px 0 0;
      max-width: 36ch;
      font-size: 0.68rem;
      letter-spacing: 0.2em;
      line-height: 1.8;
    }

    .posts-nav {
      width: 100%;
      display: grid;
      gap: 10px;
      justify-items: start;
      text-align: left;
    }

    .nav-rail {
      width: min(220px, 100%);
      max-height: calc(100vh - 96px);
      overflow-y: auto;
      padding-right: 8px;
      display: grid;
      gap: 18px;
      scrollbar-width: thin;
      scrollbar-color: #9ca3af transparent;
    }

    .nav-head {
      display: grid;
      gap: 8px;
    }

    .posts-nav-title,
    .posts-nav-empty,
    .posts-nav a {
      color: #9ca3af;
      font-size: 0.64rem;
      letter-spacing: 0.18em;
      text-transform: uppercase;
      line-height: 1.7;
    }

    .posts-nav-title {
      font-size: 0.66rem;
      letter-spacing: 0.2em;
    }

    .posts-nav a:hover {
      color: #fff;
    }

    .posts-nav-links {
      width: 100%;
      display: grid;
    }

    .posts-nav-links > * + * {
      margin-top: 10px;
      padding-top: 10px;
      border-top: 1px solid rgba(255, 255, 255, 0.12);
    }

    .post-list {
      margin-top: 0;
    }

    .post + .post {
      margin-top: 34px;
      padding-top: 34px;
      border-top: 1px solid rgba(255, 255, 255, 0.12);
    }

    .post-header > * + * {
      margin-top: 12px;
    }

    .post-summary {
      max-width: 46ch;
      font-size: 0.62rem;
      line-height: 1.75;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      gap: 12px;
    }

    .tag {
      font-size: 0.66rem;
    }

    .post-body {
      margin-top: 22px;
      color: #fff;
      font-size: 0.88rem;
      letter-spacing: 0.02em;
      line-height: 1.8;
    }

    .post-body > :first-child {
      margin-top: 0;
    }

    .post-body > :last-child {
      margin-bottom: 0;
    }

    .post-body p,
    .post-body ul,
    .post-body ol,
    .post-body blockquote,
    .post-body pre,
    .post-body h2,
    .post-body h3 {
      margin: 0 0 20px;
    }

    .post-body h2,
    .post-body h3 {
      font-size: 0.92rem;
    }

    .post-body ul,
    .post-body ol {
      padding-left: 24px;
    }

    .post-body li + li {
      margin-top: 8px;
    }

    .post-body blockquote {
      max-width: 34ch;
      padding-left: 18px;
      border-left: 2px solid rgba(255, 255, 255, 0.14);
    }

    .post-body code {
      padding: 2px 6px;
      border-radius: 8px;
      background: rgba(255, 255, 255, 0.08);
      font-size: 0.92em;
    }

    .post-body pre {
      overflow-x: auto;
      padding: 18px;
      border-radius: 18px;
      background: rgba(255, 255, 255, 0.05);
      border: 1px solid rgba(255, 255, 255, 0.12);
    }

    .post-body pre code {
      padding: 0;
      background: transparent;
    }

    .empty-state {
      margin-top: 42px;
      max-width: 100%;
      color: #9ca3af;
      font-size: 0.74rem;
      line-height: 1.9;
      letter-spacing: 0.18em;
      text-transform: uppercase;
    }

    @media (max-width: 900px) {
      .wrap {
        padding: 24px;
      }

      .layout {
        grid-template-columns: 1fr;
        width: 100%;
        gap: 24px;
      }

      .logo-side,
      .content-side,
      .nav-side {
        width: 100%;
        min-width: 0;
        justify-content: center;
        position: static;
        transform: none;
      }

      .logo-side {
        min-height: 28vh;
        padding: 24px 16px 8px;
      }

      .logo {
        transform: none;
        max-width: 286px;
        max-height: 234px;
      }

      .content-side {
        padding: 24px 28px 32px;
      }

      .content,
      .role {
        max-width: 100%;
        margin-left: auto;
        margin-right: auto;
      }

      .nav-side {
        padding: 0 28px 32px;
      }

      .nav-rail {
        max-height: none;
        padding-right: 0;
      }

      .posts-nav {
        justify-items: center;
        text-align: center;
      }

      .post-body {
        text-align: left;
      }
    }
  </style>
</head>
<body>
  <div class="bg"></div>
  <main class="wrap">
    <div class="layout">
      <aside class="logo-side">
        <div class="logo-rail">
          <a class="logo-link" href="https://hackspree.com/" aria-label="Hackspree home">
            <img src="logo.png" alt="Hackspree logo" class="logo">
          </a>
        </div>
      </aside>
      <section class="content-side">
        <div class="content">
      {{- if .HasPosts}}
          <div class="post-list">
      {{- range .Posts}}
      <article class="post" id="{{.Slug}}">
        <header class="post-header">
          <p class="post-meta">{{.DateDisplay}}</p>
          <h2><a href="#{{.Slug}}">{{.Title}}</a></h2>
          {{- if .Summary}}<p class="post-summary">{{.Summary}}</p>{{end}}
          {{- if .Tags}}
          <div class="tags">
            {{- range .Tags}}<span class="tag">{{.}}</span>{{end}}
          </div>
          {{- end}}
        </header>
        <div class="post-body">{{.BodyHTML}}</div>
      </article>
      {{- end}}
          </div>
      {{- else}}
          <p class="empty-state">No posts yet.</p>
      {{- end}}
        </div>
      </section>
      <aside class="nav-side">
        <div class="nav-rail">
          <div class="nav-head">
            <h1>{{.SiteTitle}}</h1>
            <p class="role">{{.SiteTagline}}</p>
          </div>
          <nav class="posts-nav" aria-label="Posts navigation">
            <p class="posts-nav-title">Posts</p>
            <div class="posts-nav-links">
              {{- if .HasPosts}}
              {{- range .Posts}}
              <a href="#{{.Slug}}">{{.Title}}</a>
              {{- end}}
              {{- else}}
              <span class="posts-nav-empty">No posts yet</span>
              {{- end}}
            </div>
          </nav>
        </div>
      </aside>
    </div>
  </main>
</body>
</html>
`))
)

type post struct {
	Title       string
	Date        time.Time
	DateDisplay string
	Slug        string
	Summary     string
	Tags        []string
	BodyHTML    template.HTML
}

type pageData struct {
	MetaDescription string
	SiteURL         string
	PageTitle       string
	SiteTitle       string
	SiteTagline     string
	SiteDescription string
	LatestDate      string
	Posts           []post
	HasPosts        bool
}

func main() {
	inputDir := flag.String("input-dir", "", "directory containing markdown posts")
	outputPath := flag.String("output", "", "path to generated html output")
	flag.Parse()

	if *inputDir == "" || *outputPath == "" {
		exitf("both --input-dir and --output are required")
	}

	posts, err := loadPosts(*inputDir)
	if err != nil {
		exitf("%v", err)
	}

	if err := writePage(*outputPath, posts); err != nil {
		exitf("%v", err)
	}
}

func loadPosts(inputDir string) ([]post, error) {
	pattern := filepath.Join(inputDir, "*.md")
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("glob posts: %w", err)
	}

	sort.Strings(paths)

	posts := make([]post, 0, len(paths))
	seen := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		if filepath.Base(path) == "specs.md" {
			continue
		}

		post, err := loadPost(path)
		if err != nil {
			return nil, err
		}

		if _, exists := seen[post.Slug]; exists {
			return nil, fmt.Errorf("duplicate slug %q", post.Slug)
		}
		seen[post.Slug] = struct{}{}
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		if !posts[i].Date.Equal(posts[j].Date) {
			return posts[i].Date.After(posts[j].Date)
		}
		return posts[i].Slug < posts[j].Slug
	})

	return posts, nil
}

func loadPost(path string) (post, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return post{}, fmt.Errorf("read %s: %w", path, err)
	}

	meta, body, err := parseFrontMatter(path, string(content))
	if err != nil {
		return post{}, err
	}

	title := strings.TrimSpace(meta["title"])
	dateText := strings.TrimSpace(meta["date"])
	if title == "" || dateText == "" {
		return post{}, fmt.Errorf("%s: required front matter fields are 'title' and 'date'", path)
	}

	date, err := time.Parse("2006-01-02", dateText)
	if err != nil {
		return post{}, fmt.Errorf("%s: invalid ISO date %q", path, dateText)
	}

	slug := strings.TrimSpace(meta["slug"])
	if slug == "" {
		slug, err = slugify(title)
		if err != nil {
			return post{}, fmt.Errorf("%s: %w", path, err)
		}
	}
	if !slugPattern.MatchString(slug) {
		return post{}, fmt.Errorf("%s: slug %q must use lowercase letters, numbers, and hyphens", path, slug)
	}

	renderedBody, err := renderMarkdown(body)
	if err != nil {
		return post{}, fmt.Errorf("%s: render markdown: %w", path, err)
	}

	return post{
		Title:       title,
		Date:        date,
		DateDisplay: strings.ToUpper(date.Format("Jan 02, 2006")),
		Slug:        slug,
		Summary:     strings.TrimSpace(meta["summary"]),
		Tags:        parseTags(meta["tags"]),
		BodyHTML:    renderedBody,
	}, nil
}

func parseFrontMatter(path, raw string) (map[string]string, string, error) {
	lines := strings.Split(strings.ReplaceAll(raw, "\r\n", "\n"), "\n")
	if len(lines) == 0 || strings.TrimSpace(lines[0]) != "---" {
		return nil, "", fmt.Errorf("%s: missing opening front matter delimiter", path)
	}

	closing := -1
	for index := 1; index < len(lines); index++ {
		if strings.TrimSpace(lines[index]) == "---" {
			closing = index
			break
		}
	}
	if closing == -1 {
		return nil, "", fmt.Errorf("%s: missing closing front matter delimiter", path)
	}

	meta := make(map[string]string, closing-1)
	for _, line := range lines[1:closing] {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		key, value, found := strings.Cut(trimmed, ":")
		if !found {
			return nil, "", fmt.Errorf("%s: invalid front matter line %q", path, line)
		}
		meta[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}

	body := strings.TrimSpace(strings.Join(lines[closing+1:], "\n"))
	return meta, body, nil
}

func renderMarkdown(body string) (template.HTML, error) {
	var buffer bytes.Buffer
	if err := markdowner.Convert([]byte(body), &buffer); err != nil {
		return "", err
	}
	return template.HTML(buffer.String()), nil
}

func parseTags(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}

	if strings.HasPrefix(raw, "[") && strings.HasSuffix(raw, "]") {
		raw = strings.TrimSpace(raw[1 : len(raw)-1])
	}

	parts := strings.Split(raw, ",")
	tags := make([]string, 0, len(parts))
	for _, part := range parts {
		tag := strings.TrimSpace(part)
		if tag != "" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func slugify(value string) (string, error) {
	lower := strings.ToLower(value)
	var builder strings.Builder
	lastHyphen := false
	for _, r := range lower {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			builder.WriteRune(r)
			lastHyphen = false
		case !lastHyphen:
			builder.WriteByte('-')
			lastHyphen = true
		}
	}

	slug := strings.Trim(builder.String(), "-")
	if slug == "" {
		return "", fmt.Errorf("slug cannot be empty")
	}
	return slug, nil
}

func writePage(outputPath string, posts []post) error {
	latestDate := "TBD"
	metaDescription := siteDescription
	if len(posts) > 0 {
		latestDate = posts[0].DateDisplay
		if strings.TrimSpace(posts[0].Summary) != "" {
			metaDescription = posts[0].Summary
		}
	}

	data := pageData{
		MetaDescription: metaDescription,
		SiteURL:         siteURL,
		PageTitle:       siteTitle,
		SiteTitle:       siteTitle,
		SiteTagline:     siteTagline,
		SiteDescription: siteDescription,
		LatestDate:      latestDate,
		Posts:           posts,
		HasPosts:        len(posts) > 0,
	}

	var buffer bytes.Buffer
	if err := pageTemplate.Execute(&buffer, data); err != nil {
		return fmt.Errorf("render page: %w", err)
	}

	buffer.WriteByte('\n')
	if err := os.WriteFile(outputPath, buffer.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", outputPath, err)
	}

	return nil
}

func exitf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

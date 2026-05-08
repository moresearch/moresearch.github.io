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
    :root {
      --bg: #000;
      --panel: rgba(255, 255, 255, 0.03);
      --line: rgba(255, 255, 255, 0.12);
      --text: #fff;
      --muted: #9ca3af;
      --muted-strong: #d1d5db;
      --accent: #f8fafc;
      --code-bg: rgba(255, 255, 255, 0.06);
      --shadow: 0 24px 80px rgba(0, 0, 0, 0.35);
    }

    * {
      box-sizing: border-box;
    }

    html,
    body {
      margin: 0;
      min-height: 100%;
      background: var(--bg);
      color: var(--text);
    }

    body {
      font-family: "Orbitron", Arial, sans-serif;
      line-height: 1.7;
    }

    a {
      color: inherit;
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
      background:
        radial-gradient(circle at top left, rgba(156, 163, 175, 0.12), transparent 30%),
        radial-gradient(circle at bottom right, rgba(255, 255, 255, 0.08), transparent 25%),
        #000;
      z-index: 0;
    }

    .shell {
      position: relative;
      z-index: 1;
      width: min(1400px, 100%);
      margin: 0 auto;
      padding: clamp(20px, 3vw, 32px);
      display: grid;
      grid-template-columns: minmax(190px, 260px) minmax(0, 760px);
      gap: clamp(28px, 5vw, 72px);
      justify-content: center;
    }

    .sidebar {
      position: sticky;
      top: 0;
      align-self: start;
      height: 100vh;
      padding: clamp(18px, 3vw, 28px) 0;
      display: grid;
      grid-template-rows: minmax(220px, 1fr) auto;
      gap: 18px;
    }

    .logo-link {
      display: flex;
      align-items: center;
      justify-content: center;
      overflow: visible;
      min-height: 220px;
      text-decoration: none;
    }

    .logo {
      width: min(420px, 34vw);
      max-width: none;
      transform: rotate(-90deg) scale(0.92);
      transform-origin: center center;
      filter: drop-shadow(0 18px 40px rgba(0, 0, 0, 0.35));
    }

    .sidebar-copy {
      display: grid;
      gap: 14px;
      align-self: end;
    }

    .eyebrow,
    .post-meta,
    .toc-title,
    .stat-label,
    .tag {
      margin: 0;
      color: var(--muted);
      font-size: 0.66rem;
      letter-spacing: 0.18em;
      text-transform: uppercase;
    }

    h1,
    h2,
    h3,
    h4 {
      margin: 0;
      line-height: 1.2;
      text-transform: uppercase;
      letter-spacing: 0.08em;
    }

    h1 {
      font-size: clamp(1.02rem, 1.7vw, 1.32rem);
    }

    h2 {
      font-size: clamp(1.2rem, 2.3vw, 1.65rem);
    }

    h3 {
      font-size: clamp(0.98rem, 1.9vw, 1.28rem);
    }

    .tagline,
    .summary,
    .intro p,
    .post-body,
    .post-summary {
      color: var(--muted-strong);
      font-size: 0.85rem;
      letter-spacing: 0.04em;
    }

    .tagline {
      margin: 0;
      text-transform: uppercase;
    }

    .summary {
      margin: 0;
      max-width: 34ch;
    }

    .stats {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 12px;
      margin: 10px 0 6px;
    }

    .stats strong {
      display: block;
      margin-top: 6px;
      font-size: 0.86rem;
      letter-spacing: 0.08em;
      text-transform: uppercase;
    }

    .toc {
      display: grid;
      gap: 10px;
      padding-top: 18px;
      border-top: 1px solid var(--line);
      max-height: 28vh;
      overflow: auto;
    }

    .toc a {
      color: var(--muted-strong);
      text-decoration: none;
      letter-spacing: 0.08em;
      text-transform: uppercase;
      font-size: 0.74rem;
    }

    .toc a:hover,
    .post-card h3 a:hover {
      color: var(--accent);
    }

    .content {
      max-width: 760px;
      padding: clamp(18px, 3vw, 28px) 0 72px;
    }

    .intro {
      margin-bottom: 24px;
      padding: clamp(20px, 2.8vw, 28px);
      border: 1px solid var(--line);
      border-radius: 24px;
      background: var(--panel);
      box-shadow: var(--shadow);
    }

    .intro p {
      margin: 12px 0 0;
      max-width: 56ch;
    }

    .post-card {
      margin-top: 24px;
      padding: clamp(20px, 2.8vw, 28px);
      border: 1px solid var(--line);
      border-radius: 24px;
      background: var(--panel);
      box-shadow: var(--shadow);
    }

    .post-header > * + * {
      margin-top: 14px;
    }

    .post-card h3 a {
      text-decoration: none;
    }

    .post-summary {
      margin: 0;
      max-width: 58ch;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      gap: 10px;
      margin-top: 18px;
    }

    .tag {
      padding: 4px 10px;
      border: 1px solid var(--line);
      border-radius: 999px;
    }

    .post-body {
      margin-top: 28px;
      line-height: 1.82;
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
    .post-body pre {
      margin: 0 0 18px;
    }

    .post-body ul,
    .post-body ol {
      padding-left: 24px;
    }

    .post-body li + li {
      margin-top: 8px;
    }

    .post-body blockquote {
      padding-left: 18px;
      border-left: 2px solid var(--line);
      color: var(--text);
    }

    .post-body code {
      padding: 2px 6px;
      border-radius: 8px;
      background: var(--code-bg);
      font-size: 0.92em;
    }

    .post-body pre {
      overflow-x: auto;
      padding: 18px;
      border-radius: 16px;
      background: var(--code-bg);
      border: 1px solid var(--line);
    }

    .post-body pre code {
      padding: 0;
      background: transparent;
    }

    .empty-state {
      text-align: center;
    }

    @media (max-width: 1180px) {
      .shell {
        grid-template-columns: minmax(160px, 220px) minmax(0, 1fr);
        gap: clamp(22px, 4vw, 48px);
      }

      .logo {
        width: min(340px, 30vw);
      }
    }

    @media (max-width: 960px) {
      .shell {
        grid-template-columns: 1fr;
        gap: 18px;
        padding: 20px;
      }

      .sidebar {
        position: static;
        height: auto;
        min-height: auto;
        padding: 8px 0 0;
        grid-template-rows: none;
        justify-items: center;
      }

      .logo-link {
        min-height: 0;
      }

      .logo {
        width: min(270px, 74vw);
        transform: none;
      }

      .sidebar-copy {
        width: min(640px, 100%);
        justify-items: center;
        text-align: center;
      }

      .stats {
        width: min(320px, 100%);
      }

      .toc {
        width: 100%;
        max-height: none;
      }

      .summary,
      .intro p,
      .post-summary {
        max-width: none;
      }
    }

    @media (max-width: 640px) {
      .shell {
        padding: 16px;
      }

      .intro,
      .post-card {
        padding: 18px;
      }
    }
  </style>
</head>
<body>
  <div class="bg"></div>
  <div class="shell">
    <aside class="sidebar">
      <a class="logo-link" href="https://hackspree.com/" aria-label="Hackspree home">
        <img src="logo.png" alt="Hackspree logo" class="logo">
      </a>
      <div class="sidebar-copy">
        <p class="eyebrow">single page / markdown-built</p>
        <h1>{{.SiteTitle}}</h1>
        <p class="tagline">{{.SiteTagline}}</p>
        <p class="summary">{{.SiteDescription}}</p>
        <div class="stats">
          <div>
            <span class="stat-label">posts</span>
            <strong>{{len .Posts}}</strong>
          </div>
          <div>
            <span class="stat-label">latest</span>
            <strong>{{.LatestDate}}</strong>
          </div>
        </div>
        <nav class="toc" aria-label="Post index">
          <p class="toc-title">on this page</p>
          {{- if .HasPosts}}
          {{- range .Posts}}
          <a href="#{{.Slug}}">{{.Title}}</a>
          {{- end}}
          {{- else}}
          <a href="#posts">posts coming soon</a>
          {{- end}}
        </nav>
      </div>
    </aside>
    <main class="content">
      <section class="intro">
        <p class="eyebrow">scroll the posts / keep the mark in view</p>
        <h2>All posts live here.</h2>
        <p>
          Add a new file to <code>posts/</code>, run <code>make</code>, and commit the regenerated
          <code>index.html</code>.
        </p>
      </section>
      {{- if .HasPosts}}
      {{- range .Posts}}
      <article class="post-card" id="{{.Slug}}">
        <header class="post-header">
          <p class="post-meta">{{.DateDisplay}}</p>
          <h3><a href="#{{.Slug}}">{{.Title}}</a></h3>
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
      {{- else}}
      <section class="post-card empty-state">
        <header class="post-header">
          <p class="post-meta">no posts yet</p>
          <h3>Add your first markdown file in <code>posts/</code>.</h3>
        </header>
      </section>
      {{- end}}
    </main>
  </div>
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

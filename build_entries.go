package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var postPageTemplate = template.Must(template.New("post").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <meta name="color-scheme" content="dark">
  <meta name="description" content="{{.Summary}}">
  <meta name="robots" content="index, follow">
  <link rel="canonical" href="{{.CanonicalURL}}">
  <title>{{.Title}} — hackspree</title>
  <meta property="og:title" content="{{.Title}}">
  <meta property="og:description" content="{{.Summary}}">
  <meta property="og:url" content="{{.CanonicalURL}}">
  <meta property="og:type" content="article">
  <meta property="og:image" content="https://blog.hackspree.com/logo.png">
  <meta property="og:image:width" content="512">
  <meta property="og:image:height" content="512">
  <meta name="twitter:card" content="summary">
  <meta name="twitter:title" content="{{.Title}}">
  <meta name="twitter:description" content="{{.Summary}}">
  <meta name="twitter:image" content="https://blog.hackspree.com/logo.png">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400&family=Orbitron:wght@500;700&display=swap" rel="stylesheet">
  <link href="https://fonts.cdnfonts.com/css/sudo-var" rel="stylesheet">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/highlight.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/languages/go.min.js"></script>
  <style>
    :root{--syntax-bg:#000;--syntax-normal:#bcbcbc;--syntax-keyword:#eee;--syntax-constant:#d0d0d0;--syntax-string:#8a8a8a;--syntax-comment:#585858;--syntax-number:#d0d0d0;--syntax-error-bg:#870000}
    html,body{margin:0!important;padding:0!important;width:100%!important;min-width:100%!important;min-height:100%!important;background:#000!important;color:#fff!important}
    *{box-sizing:border-box}body{position:relative;overflow-x:hidden;font-family:"Orbitron",Arial,sans-serif}
    a{color:inherit;text-decoration:none}img{display:block;max-width:100%;height:auto}
    code,pre{font-family:"Sudo Var",monospace;font-variant-ligatures:none}
    .bg{position:fixed;inset:0;background:#000;z-index:0}
    .wrap{position:relative;z-index:1;min-height:100vh;display:flex;align-items:center;justify-content:center;padding:32px;box-sizing:border-box}
    .content{width:100%;max-width:640px;margin:0 auto;padding:48px 16px}
    .back-link{display:inline-block;margin-bottom:32px;color:#9ca3af;font-size:0.52rem;letter-spacing:0.18em;text-transform:uppercase;transition:color 0.2s}
    .back-link:hover{color:#fff}
    .post-meta,.tag{margin:0;color:#9ca3af;letter-spacing:0.18em;text-transform:uppercase}
    .post-meta{font-size:0.56rem;line-height:1.4;margin-bottom:9px}
    h1,h2,h3,h4{margin:0;font-weight:600;line-height:1.2;text-transform:uppercase;letter-spacing:0.08em;color:#fff}
    h1{font-size:clamp(1.18rem,1.7vw,1.4rem);margin-bottom:12px}
    h2{font-size:clamp(0.84rem,1.1vw,0.98rem)}h3{font-size:0.72rem}
    .post-summary{max-width:54ch;padding-left:12px;border-left:1px solid rgba(255,255,255,0.14);color:#c0c8d2;font-size:0.62rem;line-height:1.78;letter-spacing:0.04em;font-family:"IBM Plex Mono","SFMono-Regular",Consolas,"Liberation Mono",Menlo,monospace;margin-bottom:28px}
    .tags{display:flex;flex-wrap:wrap;gap:12px;margin-bottom:28px}.tag{font-size:0.4rem}
    .post-body{margin-top:18px;color:#9099a5;font-size:0.74rem;letter-spacing:0.02em;line-height:1.76;font-family:"IBM Plex Mono","SFMono-Regular",Consolas,"Liberation Mono",Menlo,monospace;font-weight:400}
    .post-body>:first-child{margin-top:0}.post-body>:last-child{margin-bottom:0}
    .post-body p,.post-body ul,.post-body ol,.post-body blockquote,.post-body pre,.post-body h2,.post-body h3{margin:0 0 20px}
    .post-body h2,.post-body h3{font-size:0.62rem;color:#fff;font-family:"Orbitron",Arial,sans-serif}
    .post-body a,.post-body strong,.post-body code{color:#fff}.post-body a{text-decoration:none}
    .post-body a:hover{color:#d1d5db}
    .post-body img{width:min(240px,100%);margin:28px 0 10px;border:1px solid rgba(255,255,255,0.08);border-radius:12px;background:#060606;box-shadow:0 20px 48px rgba(0,0,0,0.36)}
    .post-body ul,.post-body ol{padding-left:24px}.post-body li+li{margin-top:8px}
    .post-body blockquote{max-width:34ch;padding-left:18px;border-left:2px solid rgba(255,255,255,0.42);color:#fff;font-weight:600}
    .post-body code{padding:2px 6px;border-radius:0;background:#000;color:var(--syntax-normal);font-family:"Sudo Var",monospace;font-size:14px;line-height:20px}
    .post-body pre{overflow-x:auto;padding:18px;border-radius:0;background:var(--syntax-bg);border:1px solid #121212;color:var(--syntax-normal);font-family:"Sudo Var",monospace;font-size:14px;line-height:20px;scrollbar-width:none;scrollbar-color:#585858 #000}
    .post-body pre code{padding:0;background:transparent;font-family:inherit;font-size:inherit;line-height:inherit}
    .post-body pre code.hljs{display:block;overflow:visible;color:var(--syntax-normal);background:transparent;font-family:inherit;font-size:inherit;line-height:inherit}
    .post-body pre:hover{scrollbar-width:thin}
    .post-body pre::-webkit-scrollbar{width:0;height:0}.post-body pre:hover::-webkit-scrollbar{width:8px;height:8px}
    .post-body pre::-webkit-scrollbar-track{background:#000}.post-body pre::-webkit-scrollbar-thumb{background:#585858}
    .post-body pre:hover::-webkit-scrollbar-thumb{background:#8a8a8a}
    .post-body .hljs-comment,.post-body .hljs-quote,.post-body .hljs-meta{color:var(--syntax-comment)}
    .post-body .hljs-keyword,.post-body .hljs-built_in,.post-body .hljs-type,.post-body .hljs-title.function_,.post-body .hljs-title.class_,.post-body .hljs-function .hljs-title,.post-body .hljs-title,.post-body .hljs-operator,.post-body .hljs-selector-tag,.post-body .hljs-section,.post-body .hljs-link,.post-body .hljs-tag{color:var(--syntax-keyword)}
    .post-body .hljs-literal,.post-body .hljs-variable,.post-body .hljs-property,.post-body .hljs-params,.post-body .hljs-attr,.post-body .hljs-attribute,.post-body .hljs-punctuation{color:var(--syntax-constant)}
    .post-body .hljs-string,.post-body .hljs-symbol,.post-body .hljs-bullet,.post-body .hljs-template-tag,.post-body .hljs-template-variable,.post-body .hljs-addition,.post-body .hljs-subst{color:var(--syntax-string)}
    .post-body .hljs-number,.post-body .hljs-regexp,.post-body .hljs-selector-class,.post-body .hljs-selector-id,.post-body .hljs-char.escape_{color:var(--syntax-number)}
    .post-body .hljs-emphasis,.post-body .hljs-strong{color:var(--syntax-keyword)}
  </style>
</head>
<body>
  <div class="bg"></div>
  <main class="wrap">
    <div class="content">
      <a class="back-link" href="https://blog.hackspree.com/">← Back to all posts</a>
      <article>
        <header>
          <p class="post-meta"><time datetime="{{.DateISO}}">{{.DateISO}}</time></p>
          <h1>{{.Title}}</h1>
          {{- if .Summary}}<p class="post-summary">{{.Summary}}</p>{{end}}
          {{- if .Tags}}
          <div class="tags">{{range .Tags}}<span class="tag">{{.}}</span>{{end}}</div>
          {{end}}
        </header>
        <div class="post-body">{{.BodyHTML}}</div>
      </article>
    </div>
  </main>
  <script>if(window.hljs){window.hljs.highlightAll()}</script>
</body>
</html>
`))

func writePostPages(posts []post) error {
	for _, p := range posts {
		dir := filepath.Join("entries", p.Slug)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create dir %s: %w", dir, err)
		}
		var buf bytes.Buffer
		if err := postPageTemplate.Execute(&buf, p); err != nil {
			return fmt.Errorf("render post %s: %w", p.Slug, err)
		}
		buf.WriteByte('\n')
		outPath := filepath.Join(dir, "index.html")
		if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}
	return nil
}

func writeSitemap(posts []post) error {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")
	b.WriteString(`  <url>` + "\n")
	b.WriteString(`    <loc>https://blog.hackspree.com/</loc>` + "\n")
	b.WriteString(`    <lastmod>` + time.Now().Format("2006-01-02") + `</lastmod>` + "\n")
	b.WriteString(`    <changefreq>daily</changefreq>` + "\n")
	b.WriteString(`    <priority>1.0</priority>` + "\n")
	b.WriteString(`  </url>` + "\n")
	for _, p := range posts {
		b.WriteString(`  <url>` + "\n")
		b.WriteString(`    <loc>` + p.CanonicalURL + `</loc>` + "\n")
		b.WriteString(`    <lastmod>` + p.DateISO + `</lastmod>` + "\n")
		b.WriteString(`    <changefreq>weekly</changefreq>` + "\n")
		b.WriteString(`    <priority>0.8</priority>` + "\n")
		b.WriteString(`  </url>` + "\n")
	}
	b.WriteString(`</urlset>` + "\n")
	return os.WriteFile("sitemap.xml", []byte(b.String()), 0o644)
}

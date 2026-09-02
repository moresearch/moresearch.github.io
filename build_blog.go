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
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

const (
	siteTitle       = "Engineering blog"
	siteURL         = "https://blog.hackspree.com/"
	siteDescription = "Single-page writing log built from Markdown sources."
	pageTitle       = "Engineering blog — hackspree"
)

var (
	slugPattern  = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	thumbPattern = regexp.MustCompile(`<img src="([^"]*)" alt="thumb:([^"]*)"(?:/>|>)`)
	markdowner   = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
			parser.WithASTTransformers(util.Prioritized(unpublishedLinkTransformer{}, 100)),
		),
	)
	pageTemplate = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <meta name="color-scheme" content="dark">
  <meta name="theme-color" content="#02391f">
  <meta name="description" content="{{.MetaDescription}}">
  <meta name="robots" content="index, follow">
  <link rel="canonical" href="{{.SiteURL}}">
  <title>{{.PageTitle}}</title>
  <meta property="og:title" content="{{.PageTitle}}">
  <meta property="og:description" content="{{.MetaDescription}}">
  <meta property="og:url" content="{{.SiteURL}}">
  <meta property="og:type" content="website">
  <meta property="og:image" content="{{.SiteURL}}logoicon2.png">
  <meta property="og:image:width" content="257">
  <meta property="og:image:height" content="705">
  <meta name="twitter:card" content="summary">
  <meta name="twitter:title" content="{{.PageTitle}}">
  <meta name="twitter:description" content="{{.MetaDescription}}">
  <meta name="twitter:image" content="{{.SiteURL}}logoicon2.png">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono:wght@400&family=Orbitron:wght@500;700&display=swap" rel="stylesheet">
  <link href="https://fonts.cdnfonts.com/css/sudo-var" rel="stylesheet">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/highlight.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/languages/go.min.js"></script>
  <style>
    :root {
      --syntax-bg: #000000;
      --syntax-normal: #bcbcbc;
      --syntax-keyword: #eeeeee;
      --syntax-constant: #d0d0d0;
      --syntax-string: #8a8a8a;
      --syntax-comment: #585858;
      --syntax-number: #d0d0d0;
      --syntax-error-bg: #870000;
    }

    html,
    body {
      margin: 0 !important;
      padding: 0 !important;
      width: 100% !important;
      min-width: 100% !important;
      min-height: 100% !important;
      background: #02391f !important;
      color: #eafff3 !important;
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
      font-family: "Sudo Var", monospace;
      font-variant-ligatures: none;
    }

    .bg {
      position: fixed;
      inset: 0;
      background: #02391f;
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
      box-sizing: border-box;
    }

    .layout {
      width: min(1400px, 100%);
      display: grid;
      grid-template-columns: minmax(220px, 20%) minmax(0, 1fr) minmax(200px, 228px);
      gap: 36px;
      align-items: start;
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
      max-width: min(240px, 100%);
      max-height: min(58vh, 640px);
    }

    .content-side {
      display: flex;
      align-items: flex-start;
      justify-content: flex-start;
      padding: 48px 48px 48px 8px;
      box-sizing: border-box;
      min-width: 0;
    }

    .content {
      width: 100%;
      max-width: 640px;
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
    .tag {
      margin: 0;
      color: #79c2a0;
      letter-spacing: 0.18em;
      text-transform: uppercase;
    }

    .post-meta {
      font-size: 0.56rem;
      line-height: 1.4;
    }

    .raw-link {
      display: inline-flex;
      align-items: center;
      margin-left: 0.55em;
      color: #4f9e77;
      vertical-align: middle;
      text-decoration: none;
      opacity: 0.9;
      transition: color 0.15s ease, opacity 0.15s ease;
    }

    .raw-link:hover,
    .raw-link:focus-visible {
      color: #00f090;
      opacity: 1;
    }

    .post-meta time {
      position: relative;
      display: inline-flex;
      align-items: center;
      cursor: default;
      outline: none;
    }

    .post-meta time::before,
    .post-meta time::after {
      opacity: 0;
      pointer-events: none;
      transition: opacity 0.18s ease, transform 0.18s ease;
    }

    .post-meta time::before {
      content: "";
      position: absolute;
      left: calc(100% + 8px);
      top: 50%;
      width: 8px;
      height: 8px;
      border-top: 1px solid rgba(0, 224, 128, 0.30);
      border-left: 1px solid rgba(0, 224, 128, 0.30);
      background: rgba(2, 26, 14, 0.96);
      transform: translateY(-50%) rotate(-45deg);
      z-index: 2;
    }

    .post-meta time::after {
      content: attr(data-utc);
      position: absolute;
      left: calc(100% + 12px);
      top: 50%;
      padding: 6px 10px;
      border-radius: 10px;
      border: 1px solid rgba(0, 224, 128, 0.30);
      background: rgba(2, 26, 14, 0.96);
      color: #d9ffe9;
      font-size: 0.44rem;
      letter-spacing: 0.08em;
      line-height: 1.4;
      white-space: nowrap;
      text-transform: uppercase;
      transform: translateY(-50%) translateX(4px);
      z-index: 3;
    }

    .post-meta time:hover::before,
    .post-meta time:focus-visible::before {
      opacity: 1;
    }

    .post-meta time:hover::after,
    .post-meta time:focus-visible::after {
      opacity: 1;
      transform: translateY(-50%) translateX(0);
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
      color: #f3fff8;
    }

    h1 {
      font-size: clamp(0.84rem, 1.1vw, 0.98rem);
    }

    h2 {
      font-size: clamp(1.18rem, 1.7vw, 1.4rem);
    }

    h3 {
      font-size: 0.72rem;
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
      overflow-x: hidden;
      padding-right: 8px;
      display: grid;
      gap: 18px;
      scroll-behavior: smooth;
      scrollbar-width: thin;
      scrollbar-color: transparent transparent;
    }

    .nav-rail:hover {
      scrollbar-color: #79c2a0 transparent;
    }

    .nav-rail::-webkit-scrollbar {
      width: 6px;
    }

    .nav-rail::-webkit-scrollbar-track {
      background: transparent;
    }

    .nav-rail::-webkit-scrollbar-thumb {
      border-radius: 999px;
      background: transparent;
    }

    .nav-rail:hover::-webkit-scrollbar-thumb {
      background: #79c2a0;
    }

    .nav-head {
      display: grid;
      gap: 8px;
    }

    .posts-nav-empty,
    .posts-nav a {
      color: #79c2a0;
      display: block;
      width: 100%;
      padding-left: 10px;
      border-left: 2px solid transparent;
      font-size: 0.44rem;
      letter-spacing: 0.18em;
      text-transform: uppercase;
      line-height: 1.5;
      overflow-wrap: anywhere;
      word-break: break-word;
      transition: color 0.2s ease, border-color 0.2s ease;
    }

    .posts-nav a:hover {
      color: #eafff3;
    }

    .posts-nav a.is-active,
    .posts-nav a[aria-current="true"] {
      color: #f3fff8;
      border-left-color: rgba(0, 240, 144, 0.92);
    }

    .posts-nav-links {
      width: 100%;
      display: grid;
    }

    .posts-nav-links > * + * {
      margin-top: 10px;
      padding-top: 10px;
      border-top: 1px solid rgba(0, 224, 128, 0.22);
    }

    .post-list {
      margin-top: 0;
    }

    .post + .post {
      margin-top: 28px;
      padding-top: 28px;
      border-top: 1px solid rgba(0, 224, 128, 0.16);
    }

    .post-header > * + * {
      margin-top: 9px;
    }

    .post-summary {
      max-width: 54ch;
      padding-left: 12px;
      border-left: 1px solid rgba(0, 224, 128, 0.30);
      color: #d6f5e4;
      font-size: 0.62rem;
      line-height: 1.78;
      letter-spacing: 0.04em;
      text-transform: none;
      font-family: "IBM Plex Mono", "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
    }

    .tags {
      display: flex;
      flex-wrap: wrap;
      gap: 12px;
    }

    .tag {
      font-size: 0.4rem;
    }

    .post-body {
      margin-top: 18px;
      color: #8fd9b4;
      font-size: 0.74rem;
      letter-spacing: 0.02em;
      line-height: 1.76;
      font-family: "IBM Plex Mono", "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
      font-weight: 400;
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
    .post-body h3,
    .post-body table {
      margin: 0 0 20px;
    }

    .post-body table {
      width: 100%;
      border-collapse: collapse;
      font-size: 0.62rem;
      line-height: 1.6;
    }

    .post-body table th,
    .post-body table td {
      border: 1px solid rgba(0, 224, 128, 0.18);
      padding: 6px 8px;
      vertical-align: top;
      text-align: left;
    }

    .post-body table th {
      color: #d6f5e4;
      text-transform: uppercase;
      letter-spacing: 0.08em;
      font-weight: 600;
      background: rgba(0, 224, 128, 0.08);
    }

    .post-body table tr:nth-child(even) td {
      background: rgba(0, 224, 128, 0.035);
    }

    .post-body h2,
    .post-body h3 {
      font-size: 0.62rem;
      color: #f3fff8;
      font-family: "Orbitron", Arial, sans-serif;
    }

    .post-body a,
    .post-body strong,
    .post-body code {
      color: #eafff3;
    }

    .post-body a {
      text-decoration: none;
    }

    .post-body a:hover {
      color: #00f090;
    }

    .post-body img {
      width: min(720px, 100%);
      margin: 28px 0 10px;
      border: 1px solid rgba(0, 224, 128, 0.20);
      border-radius: 12px;
      background: #021a0e;
      box-shadow: 0 20px 48px rgba(0, 0, 0, 0.4);
    }

    .post-body figure.fig-thumb {
      display: inline-block;
      width: 250px;
      margin: 0 12px 14px 0;
      vertical-align: top;
      border: 1px solid rgba(0, 224, 128, 0.25);
      border-radius: 8px;
      background: #021a0e;
      overflow: hidden;
    }

    .post-body figure.fig-thumb img {
      width: 100%;
      height: auto;
      margin: 0;
      border: 0;
      border-radius: 0;
      box-shadow: none;
    }

    .post-body figure.fig-thumb figcaption {
      padding: 7px 9px;
      color: #79c2a0;
      font-size: 9px;
      line-height: 1.55;
      letter-spacing: 0.02em;
      font-family: "IBM Plex Mono", "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
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
      border-left: 2px solid rgba(0, 240, 144, 0.55);
      color: #eafff3;
      font-weight: 600;
    }

    .post-body code {
      padding: 2px 6px;
      border-radius: 0;
      background: #000;
      color: var(--syntax-normal);
      font-family: "Sudo Var", monospace;
      font-size: 14px;
      line-height: 20px;
    }

    .post-body pre {
      overflow-x: auto;
      padding: 18px;
      border-radius: 0;
      background: var(--syntax-bg);
      border: 1px solid rgba(0, 224, 128, 0.20);
      color: var(--syntax-normal);
      font-family: "Sudo Var", monospace;
      font-size: 14px;
      line-height: 20px;
      scrollbar-width: none;
      scrollbar-color: #2e8b57 #000;
    }

    .post-body pre code {
      padding: 0;
      background: transparent;
      font-family: inherit;
      font-size: inherit;
      line-height: inherit;
    }

    .post-body pre code.hljs {
      display: block;
      overflow: visible;
      color: var(--syntax-normal);
      background: transparent;
      font-family: inherit;
      font-size: inherit;
      line-height: inherit;
    }

    .post-body pre:hover {
      scrollbar-width: thin;
    }

    .post-body pre::-webkit-scrollbar {
      width: 0;
      height: 0;
    }

    .post-body pre:hover::-webkit-scrollbar {
      width: 8px;
      height: 8px;
    }

    .post-body pre::-webkit-scrollbar-track {
      background: #000;
    }

    .post-body pre::-webkit-scrollbar-thumb {
      background: #2e8b57;
    }

    .post-body pre:hover::-webkit-scrollbar-thumb {
      background: #5fb98b;
    }

    .post-body .hljs-comment,
    .post-body .hljs-quote,
    .post-body .hljs-meta {
      color: var(--syntax-comment);
    }

    .post-body .hljs-keyword,
    .post-body .hljs-built_in,
    .post-body .hljs-type,
    .post-body .hljs-title.function_,
    .post-body .hljs-title.class_,
    .post-body .hljs-function .hljs-title,
    .post-body .hljs-title,
    .post-body .hljs-operator,
    .post-body .hljs-selector-tag,
    .post-body .hljs-section,
    .post-body .hljs-link,
    .post-body .hljs-tag {
      color: var(--syntax-keyword);
    }

    .post-body .hljs-literal,
    .post-body .hljs-variable,
    .post-body .hljs-property,
    .post-body .hljs-params,
    .post-body .hljs-attr,
    .post-body .hljs-attribute,
    .post-body .hljs-punctuation {
      color: var(--syntax-constant);
    }

    .post-body .hljs-string,
    .post-body .hljs-symbol,
    .post-body .hljs-bullet,
    .post-body .hljs-template-tag,
    .post-body .hljs-template-variable,
    .post-body .hljs-addition,
    .post-body .hljs-subst {
      color: var(--syntax-string);
    }

    .post-body .hljs-number,
    .post-body .hljs-regexp,
    .post-body .hljs-selector-class,
    .post-body .hljs-selector-id,
    .post-body .hljs-char.escape_ {
      color: var(--syntax-number);
    }

    .post-body .hljs-emphasis,
    .post-body .hljs-strong {
      color: var(--syntax-keyword);
    }

    .empty-state {
      margin-top: 42px;
      max-width: 100%;
      color: #79c2a0;
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
        max-width: min(150px, 70vw);
        max-height: 30vh;
      }

      .content-side {
        padding: 24px 28px 32px;
      }

      .content {
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
            <img src="/logoicon2.png" alt="Hackspree logo" class="logo">
          </a>
        </div>
      </aside>
      <section class="content-side">
        <div class="content">
      {{- if .HasPosts}}
          <div class="post-list">
      {{- range .Posts}}
      <article class="post" id="{{.Slug}}">
        <link rel="canonical" href="{{$.SiteURL}}entries/{{.Slug}}/">
        <header class="post-header">
          <p class="post-meta"><time datetime="{{.DateISO}}" data-utc="{{.DateHoverUTC}}" aria-label="{{.DateHoverUTC}}" tabindex="0">{{.DateUnix}}</time></p>
          <h2><a href="#{{.Slug}}">{{.Title}}</a><a class="raw-link" href="{{.RawURL}}" title="Raw markdown source" aria-label="Raw markdown source of {{.Title}}" target="_blank" rel="noopener noreferrer"><svg class="raw-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" width="13" height="13" fill="none" stroke="currentColor" stroke-width="1.1" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M9.5 1.5H4.5a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V4.5L9.5 1.5Z"/><path d="M9.5 1.5v3h3"/><path d="M5.5 8h5"/><path d="M5.5 10.5h5"/><path d="M5.5 13h3"/></svg></a></h2>
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
          </div>
          <nav class="posts-nav" aria-label="Posts navigation">
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
  <script>
    if (window.hljs) {
      window.hljs.highlightAll();
    }

    (() => {
      const posts = Array.from(document.querySelectorAll('.post[id]'));
      const navRail = document.querySelector('.nav-rail');
      const links = Array.from(document.querySelectorAll('.posts-nav-links a[href^="#"]'));
      if (!posts.length || !navRail || !links.length) {
        return;
      }

      const linkById = new Map(
        links.map((link) => [decodeURIComponent(link.getAttribute('href').slice(1)), link]),
      );
      const scrollBehavior = window.matchMedia('(prefers-reduced-motion: reduce)').matches
        ? 'auto'
        : 'smooth';
      let activeID = '';

      const setActive = (id) => {
        if (!id || id === activeID) {
          return;
        }

        if (activeID && linkById.has(activeID)) {
          const previous = linkById.get(activeID);
          previous.classList.remove('is-active');
          previous.removeAttribute('aria-current');
        }

        const next = linkById.get(id);
        if (!next) {
          return;
        }

        next.classList.add('is-active');
        next.setAttribute('aria-current', 'true');
        if (navRail.scrollHeight > navRail.clientHeight) {
          const nextTop =
            next.getBoundingClientRect().top - navRail.getBoundingClientRect().top + navRail.scrollTop;
          const targetTop = Math.max(
            0,
            nextTop - navRail.clientHeight / 2 + next.clientHeight / 2,
          );
          navRail.scrollTo({top: targetTop, behavior: scrollBehavior});
        }
        activeID = id;
      };

      const fromHash = () => {
        const hashID = decodeURIComponent(window.location.hash.slice(1));
        if (linkById.has(hashID)) {
          setActive(hashID);
          return true;
        }
        return false;
      };

      if (!('IntersectionObserver' in window)) {
        if (!fromHash()) {
          setActive(posts[0].id);
        }
        return;
      }

      const visible = new Map();
      const observer = new IntersectionObserver(
        (entries) => {
          for (const entry of entries) {
            if (entry.isIntersecting) {
              visible.set(entry.target.id, entry.intersectionRatio);
            } else {
              visible.delete(entry.target.id);
            }
          }

          let bestID = '';
          let bestRatio = -1;
          for (const [id, ratio] of visible.entries()) {
            if (ratio > bestRatio) {
              bestRatio = ratio;
              bestID = id;
            }
          }

          if (bestID) {
            setActive(bestID);
          }
        },
        {
          rootMargin: '-18% 0px -52% 0px',
          threshold: [0, 0.15, 0.3, 0.45, 0.6, 0.75, 1],
        },
      );

      for (const post of posts) {
        observer.observe(post);
      }

      window.addEventListener('hashchange', fromHash);
      if (!fromHash()) {
        setActive(posts[0].id);
      }
    })();
  </script>
</body>
</html>
`))
)

type post struct {
	Title        string
	Date         time.Time
	DateISO      string
	DateUnix     int64
	DateHoverUTC string
	Slug         string
	Summary      string
	Tags         []string
	BodyHTML     template.HTML
	CanonicalURL string
	ModTime      time.Time
	RawURL       string
	SourcePath   string
}

type pageData struct {
	MetaDescription string
	SiteURL         string
	PageTitle       string
	SiteTitle       string
	SiteDescription string
	LatestDate      string
	Posts           []post
	HasPosts        bool
}

func main() {
	inputDir := flag.String("input-dir", "", "directory containing markdown posts")
	outputPath := flag.String("output", "", "path to generated html output")
	todayFlag := flag.String("today", "", "build date YYYY-MM-DD; defaults to now. Posts dated after this are treated as unpublished (scheduled).")
	flag.Parse()

	if *inputDir == "" || *outputPath == "" {
		exitf("both --input-dir and --output are required")
	}

	today, err := buildDate(*todayFlag)
	if err != nil {
		exitf("%v", err)
	}

	known, published, err := scanSlugs(*inputDir, today)
	if err != nil {
		exitf("%v", err)
	}
	knownSlugs, publishedSlugs = known, published

	posts, err := loadPosts(*inputDir, today)
	if err != nil {
		exitf("%v", err)
	}

	if err := writePage(*outputPath, posts); err != nil {
		exitf("%v", err)
	}
	if err := writePostPages(posts); err != nil {
		exitf("%v", err)
	}
	if err := writeRawPosts(posts); err != nil {
		exitf("%v", err)
	}
	if err := cleanupOrphans(posts); err != nil {
		exitf("%v", err)
	}
	if err := writeSitemap(posts); err != nil {
		exitf("%v", err)
	}
}

// buildDate resolves the effective publication date for the build: the
// --today flag when given, otherwise the wall-clock date. Future-dated posts
// are scheduled content and must not appear in the archive until their date.
func buildDate(todayFlag string) (time.Time, error) {
	if todayFlag != "" {
		t, err := time.ParseInLocation("2006-01-02", todayFlag, time.Local)
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid --today %q (want YYYY-MM-DD): %w", todayFlag, err)
		}
		return t, nil
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local), nil
}

// knownSlugs is every post slug in the source tree (published, scheduled, and
// drafts); publishedSlugs is the subset this build actually rendered. The link
// transformer uses both to decide how internal post links render, so they are
// set once in main before any markdown is converted.
var (
	knownSlugs     map[string]bool
	publishedSlugs map[string]bool
)

// unpublishedLinkTransformer rewrites internal links to other posts based on
// what this build actually published:
//   - link to a published post   -> absolute URL to the archive anchor, which
//     works both on the single-page index and on the standalone entry pages
//   - link to an unpublished post (draft or scheduled) -> unlinked text, so the
//     archive never ships a dangling anchor; the link returns automatically
//     when a scheduled post is published or a draft is promoted
//
// Links to headings within the same post and external links are left untouched.
type unpublishedLinkTransformer struct{}

func (unpublishedLinkTransformer) Transform(doc *ast.Document, _ text.Reader, _ parser.Context) {
	// Classify links first, then mutate after the walk: removing a node while
	// ast.Walk is iterating its siblings breaks the chain and silently skips
	// every link that follows it.
	var rewrite, strip []*ast.Link
	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		link, ok := n.(*ast.Link)
		if !ok {
			return ast.WalkContinue, nil
		}
		frag, ok := postAnchorFragment(string(link.Destination))
		if !ok {
			return ast.WalkContinue, nil
		}
		if publishedSlugs[frag] {
			rewrite = append(rewrite, link)
		} else if knownSlugs[frag] {
			strip = append(strip, link)
		}
		return ast.WalkContinue, nil
	})

	for _, link := range rewrite {
		frag, _ := postAnchorFragment(string(link.Destination))
		link.Destination = []byte(siteURL + "#" + frag)
	}

	for _, link := range strip {
		// Draft or scheduled: keep the link text, drop the anchor.
		parent := link.Parent()
		for child := link.FirstChild(); child != nil; {
			next := child.NextSibling()
			parent.InsertBefore(parent, link, child)
			child = next
		}
		parent.RemoveChild(parent, link)
	}
}

// postAnchorFragment returns the post slug when dest is an internal link to a
// post anchor: a bare "#slug" or an absolute "https://blog.hackspree.com/#slug".
func postAnchorFragment(dest string) (string, bool) {
	switch {
	case strings.HasPrefix(dest, siteURL+"#"):
		return strings.TrimPrefix(dest, siteURL+"#"), true
	case strings.HasPrefix(dest, "#") && len(dest) > 1:
		return dest[1:], true
	}
	return "", false
}

// isDraft reports whether front matter marks the post as a draft. Drafts stay
// in the source tree but are never rendered into the archive.
func isDraft(meta map[string]string) bool {
	d := strings.ToLower(strings.TrimSpace(meta["draft"]))
	return d == "true" || d == "yes"
}

// scanSlugs reads front matter for every post source to build the universe of
// post slugs (published, scheduled, and drafts) without rendering anything.
// Invalid or unparsable files are skipped here; loadPosts reports the real
// error when it processes the file.
func scanSlugs(inputDir string, today time.Time) (known, published map[string]bool, err error) {
	known = make(map[string]bool)
	published = make(map[string]bool)

	paths, err := filepath.Glob(filepath.Join(inputDir, "*.md"))
	if err != nil {
		return nil, nil, fmt.Errorf("glob posts: %w", err)
	}

	for _, path := range paths {
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, nil, fmt.Errorf("read %s: %w", path, err)
		}
		meta, _, err := parseFrontMatter(path, string(content))
		if err != nil {
			continue
		}
		if isDraft(meta) {
			continue
		}

		date, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(meta["date"]), time.Local)
		if err != nil {
			continue
		}

		slug := strings.TrimSpace(meta["slug"])
		if slug == "" {
			slug, _ = slugify(strings.TrimSpace(meta["title"]))
		}
		if slug == "" {
			continue
		}
		known[slug] = true
		if !date.After(today) {
			published[slug] = true
		}
	}
	return known, published, nil
}

func loadPosts(inputDir string, today time.Time) ([]post, error) {
	pattern := filepath.Join(inputDir, "*.md")
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("glob posts: %w", err)
	}

	posts := make([]post, 0, len(paths))
	seen := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		// read front matter early to allow draft and schedule skipping
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}
		meta, _, err := parseFrontMatter(path, string(content))
		if err != nil {
			return nil, err
		}
		if isDraft(meta) {
			// skip draft posts
			continue
		}

		// Scheduled posts: front-matter date in the future means the post is
		// not published yet. It stays in the source tree and appears on the
		// date it is dated for; the archive is always strictly date-ordered.
		postDate, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(meta["date"]), time.Local)
		if err != nil {
			return nil, fmt.Errorf("%s: invalid date %q: %w", path, meta["date"], err)
		}
		if postDate.After(today) {
			continue
		}

		post, err := loadPost(path)
		if err != nil {
			return nil, err
		}
		post.SourcePath = path
		post.RawURL = siteURL + "raw/" + post.Slug + ".md"

		if info, statErr := os.Stat(path); statErr == nil {
			post.ModTime = info.ModTime()
		}

		if _, exists := seen[post.Slug]; exists {
			return nil, fmt.Errorf("duplicate slug %q", post.Slug)
		}
		seen[post.Slug] = struct{}{}
		posts = append(posts, post)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		// Primary: front-matter date, newest first.
		if !posts[i].Date.Equal(posts[j].Date) {
			return posts[i].Date.After(posts[j].Date)
		}
		// Secondary: file mtime preserves same-day publication order.
		if !posts[i].ModTime.Equal(posts[j].ModTime) {
			return posts[i].ModTime.After(posts[j].ModTime)
		}
		// Tertiary: deterministic fallback so fresh clones (equal mtimes)
		// rebuild to the same same-day order as the committed archive.
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
		Title:        title,
		Date:         date,
		DateISO:      date.Format("2006-01-02"),
		DateUnix:     date.Unix(),
		DateHoverUTC: date.UTC().Format("Jan 02, 2006 UTC"),
		Slug:         slug,
		Summary:      strings.TrimSpace(meta["summary"]),
		Tags:         parseTags(meta["tags"]),
		BodyHTML:     renderedBody,
		CanonicalURL: siteURL + "entries/" + slug + "/",
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
		meta[strings.TrimSpace(key)] = unquote(strings.TrimSpace(value))
	}

	body := strings.TrimSpace(strings.Join(lines[closing+1:], "\n"))
	return meta, body, nil
}

// unquote strips a matching pair of surrounding double or single quotes,
// so YAML-style quoted values don't render with literal quote marks.
func unquote(value string) string {
	if len(value) >= 2 {
		first, last := value[0], value[len(value)-1]
		if first == last && (first == '"' || first == '\'') {
			return value[1 : len(value)-1]
		}
	}
	return value
}

func renderMarkdown(body string) (template.HTML, error) {
	var buffer bytes.Buffer
	if err := markdowner.Convert([]byte(body), &buffer); err != nil {
		return "", err
	}
	rendered := buffer.String()
	// Images with alt text prefixed "thumb:" become compact thumbnails with
	// captions instead of full-width figures.
	rendered = thumbPattern.ReplaceAllString(
		rendered,
		`<figure class="fig-thumb"><img src="$1" alt="$2" loading="lazy"><figcaption>$2</figcaption></figure>`,
	)
	// Unwrap the paragraph goldmark wraps around each image so thumbnails can
	// flow as an inline-block gallery. Handles both single and consecutive
	// thumb figures in one paragraph.
	rendered = strings.ReplaceAll(rendered, `<p><figure class="fig-thumb">`, `<figure class="fig-thumb">`)
	rendered = strings.ReplaceAll(rendered, `</figure></p>`, `</figure>`)
	// Root-relative image URLs so images resolve on the single-page index and
	// on the standalone entry pages (which live under /entries/<slug>/).
	rendered = strings.ReplaceAll(rendered, `src="images/`, `src="/images/`)
	rendered = strings.ReplaceAll(
		rendered,
		`<a href=`,
		`<a target="_blank" rel="noopener noreferrer" href=`,
	)
	return template.HTML(rendered), nil
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
		latestDate = fmt.Sprintf("%d", posts[0].DateUnix)
		if strings.TrimSpace(posts[0].Summary) != "" {
			metaDescription = posts[0].Summary
		}
	}

	data := pageData{
		MetaDescription: metaDescription,
		SiteURL:         siteURL,
		PageTitle:       pageTitle,
		SiteTitle:       siteTitle,
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

package main

import (
	"bytes"
	"encoding/xml"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

// Post represents a blog post loaded from a markdown file.
type Post struct {
	Title       string
	Slug        string
	Category    string // "school" | "work" | "home"
	Tags        []string
	Date        time.Time
	Description string
	Content     template.HTML
}

// PageData is passed to every template execution.
type PageData struct {
	SiteName    string
	Title       string
	NavPath     string
	Post        *Post
	Posts       []Post
	PageContent template.HTML
	AllTags     []string
	SelectedTag string
}

var mdParser = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		extension.GFM,
		extension.Typographer,
	),
	goldmark.WithRendererOptions(
		goldmarkhtml.WithHardWraps(),
		goldmarkhtml.WithUnsafe(),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

func parseMarkdownFile(filePath string) (string, map[string]interface{}, error) {
	src, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil, err
	}
	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := mdParser.Convert(src, &buf, parser.WithContext(ctx)); err != nil {
		return "", nil, err
	}
	return buf.String(), meta.Get(ctx), nil
}

func loadPost(filePath, category, slug string) (*Post, error) {
	content, metaData, err := parseMarkdownFile(filePath)
	if err != nil {
		return nil, err
	}

	// If the markdown body opens with an <h1> heading, remove it: the post
	// template already renders the title from frontmatter, so keeping the h1
	// would cause the heading to appear twice on the page.
	trimmed := strings.TrimSpace(content)
	if strings.HasPrefix(trimmed, "<h1") {
		if end := strings.Index(trimmed, "</h1>"); end != -1 {
			trimmed = strings.TrimSpace(trimmed[end+len("</h1>"):])
		}
	}

	p := &Post{
		Slug:     slug,
		Category: category,
		Content:  template.HTML(trimmed),
	}

	if v, ok := metaData["title"].(string); ok {
		p.Title = v
	} else {
		p.Title = titleCase(strings.ReplaceAll(slug, "-", " "))
	}

	if v, ok := metaData["description"].(string); ok {
		p.Description = v
	}

	if v, ok := metaData["date"].(string); ok {
		if t, err := time.Parse("2006-01-02", v); err == nil {
			p.Date = t
		}
	}

	if tags, ok := metaData["tags"].([]interface{}); ok {
		for _, t := range tags {
			if s, ok := t.(string); ok {
				p.Tags = append(p.Tags, strings.ToLower(strings.TrimSpace(s)))
			}
		}
	}
	return p, nil
}

func loadPosts(filterTag, filterCategory string) ([]Post, []string, error) {
	var posts []Post
	tagSet := map[string]bool{}

	for _, category := range []string{"school", "work", "home"} {
		dir := filepath.Join("content", "posts", category)
		entries, err := os.ReadDir(dir)
		if os.IsNotExist(err) {
			continue
		}

		if err != nil {
			return nil, nil, err
		}

		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
				continue
			}

			slug := strings.TrimSuffix(e.Name(), ".md")
			p, err := loadPost(filepath.Join(dir, e.Name()), category, slug)
			if err != nil {
				log.Printf("warn: loading %s: %v", e.Name(), err)

				continue
			}

			for _, t := range p.Tags {
				tagSet[t] = true
			}

			if filterCategory != "" && p.Category != filterCategory {
				continue
			}

			if filterTag != "" && !sliceContains(p.Tags, filterTag) {
				continue
			}

			posts = append(posts, *p)
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	tags := make([]string, 0, len(tagSet))
	for t := range tagSet {
		tags = append(tags, t)
	}
	sort.Strings(tags)

	return posts, tags, nil
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}

	return strings.Join(words, " ")
}

func sliceContains(s []string, v string) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}

var tmplFuncs = template.FuncMap{
	"contains":   sliceContains,
	"lower":      strings.ToLower,
	"join":       strings.Join,
	"formatDate": func(t time.Time) string { return t.Format("2006-01-02") },
	"isZero":     func(t time.Time) bool { return t.IsZero() },
}

// renderPage renders base.html + the given page template (full page).
func renderPage(w http.ResponseWriter, page string, data PageData) {
	tmpl, err := template.New("").Funcs(tmplFuncs).ParseFiles(
		"templates/base.html",
		"templates/"+page+".html",
	)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("template execute error: %v", err)
	}
}

// renderPartial renders a named template block from a single file (for HTMX swaps).
func renderPartial(w http.ResponseWriter, page, block string, data PageData) {
	tmpl, err := template.New("").Funcs(tmplFuncs).ParseFiles("templates/" + page + ".html")
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, block, data); err != nil {
		log.Printf("template execute error: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	posts, _, err := loadPosts("", "")
	if err != nil {
		log.Printf("warn: loading posts for index: %v", err)
	}

	const maxRecent = 3
	if len(posts) > maxRecent {
		posts = posts[:maxRecent]
	}

	renderPage(w, "index", PageData{
		SiteName: siteName,
		Title:    siteName,
		NavPath:  "/",
		Posts:    posts,
	})
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	cat := r.URL.Query().Get("category")
	posts, allTags, err := loadPosts(tag, cat)
	if err != nil {
		http.Error(w, "error loading posts", http.StatusInternalServerError)

		return
	}

	data := PageData{
		SiteName:    siteName,
		Title:       "~/posts/",
		NavPath:     "/posts",
		Posts:       posts,
		AllTags:     allTags,
		SelectedTag: tag,
	}

	if r.Header.Get("HX-Request") == "true" {
		renderPartial(w, "posts", "posts-list", data)

		return
	}

	renderPage(w, "posts", data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// path: /posts/{category}/{slug}
	parts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/posts/"), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		http.NotFound(w, r)

		return
	}

	category, slug := parts[0], parts[1]
	fp := filepath.Join("content", "posts", category, slug+".md")
	p, err := loadPost(fp, category, slug)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "error loading post", http.StatusInternalServerError)
		}

		return
	}

	renderPage(w, "post", PageData{
		SiteName: siteName,
		Title:    p.Title,
		NavPath:  "/posts",
		Post:     p,
	})
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	html, _, err := parseMarkdownFile("content/projects.md")
	if err != nil {
		http.Error(w, "error loading projects", http.StatusInternalServerError)

		return
	}

	renderPage(w, "projects", PageData{
		SiteName:    siteName,
		Title:       "~/projects.md",
		NavPath:     "/projects",
		PageContent: template.HTML(html),
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	html, _, err := parseMarkdownFile("content/about.md")
	if err != nil {
		http.Error(w, "error loading about", http.StatusInternalServerError)

		return
	}

	renderPage(w, "about", PageData{
		SiteName:    siteName,
		Title:       "~/about.md",
		NavPath:     "/about",
		PageContent: template.HTML(html),
	})
}

type rssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	GUID        string   `xml:"guid"`
	Categories  []string `xml:"category"`
}

type rssChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
	posts, _, err := loadPosts("", "")
	if err != nil {
		http.Error(w, "error building feed", http.StatusInternalServerError)
		return
	}

	scheme := "https"
	if r.TLS == nil && r.Header.Get("X-Forwarded-Proto") != "https" {
		scheme = "http"
	}

	baseURL := scheme + "://" + r.Host
	items := make([]rssItem, 0, len(posts))
	for _, p := range posts {
		link := baseURL + "/posts/" + p.Category + "/" + p.Slug
		items = append(items, rssItem{
			Title:       p.Title,
			Link:        link,
			Description: p.Description,
			PubDate:     p.Date.UTC().Format(time.RFC1123Z),
			GUID:        link,
			Categories:  p.Tags,
		})
	}

	lastBuild := time.Now().UTC().Format(time.RFC1123Z)
	if len(posts) > 0 && !posts[0].Date.IsZero() {
		lastBuild = posts[0].Date.UTC().Format(time.RFC1123Z)
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         siteName,
			Link:          baseURL,
			Description:   siteName + " - blog posts",
			Language:      "en",
			LastBuildDate: lastBuild,
			Items:         items,
		},
	}
	out, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		http.Error(w, "error encoding feed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.Write([]byte(xml.Header))
	w.Write(out)
}

const siteName = "Bastian Asmussen"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/posts", postsHandler)
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/posts/" {
			http.Redirect(w, r, "/posts", http.StatusMovedPermanently)

			return
		}
		postHandler(w, r)
	})

	mux.HandleFunc("/projects", projectsHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/feed.xml", rssHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	log.Printf("listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

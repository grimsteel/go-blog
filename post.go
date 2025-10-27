package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Id string
	Date string
	Filename string
	Title string
}

type Comment struct {
	Name string
	Content string
}

func (post *Post) HumanDate() (string) {
	parsedDate, err := time.Parse(time.DateOnly, post.Date)
	check(err)

	return parsedDate.Format("Monday, January _2")
}

func getPostList() ([]Post) {
	postListJson, err := os.ReadFile("posts/posts.json")
	check(err)

	// parse JSON
	var posts []Post
	check(json.Unmarshal(postListJson, &posts))

	return posts
}

// used by the template
func (post *Post) Render() (template.HTML) {
	// read file
	postContents, err := os.ReadFile(fmt.Sprintf("posts/%s", post.Filename))
	check(err)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.Footnotes
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(postContents)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return template.HTML(string(markdown.Render(doc, renderer)))
}

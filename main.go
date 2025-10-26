package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// jetbrains
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// jetbrains
func renderTemplate(data any, templateFile string, w *http.ResponseWriter) {
	t, _ := template.ParseFiles(fmt.Sprintf("templates/%s.html", templateFile))
	check(t.Execute(*w, data))
}

type Post struct {
	Id string
	Date string
	Filename string
	Title string
}

func getPostList() ([]Post) {
	postListJson, err := os.ReadFile("posts.json")
	check(err)

	// parse JSON
	var posts []Post
	check(json.Unmarshal(postListJson, &posts))

	return posts
}

func index(w http.ResponseWriter, r *http.Request) {
	posts := getPostList()
	
	renderTemplate(posts, "index", &w)
}

func main() {
	listen_address := os.Getenv("LISTEN_ADDRESS")
	if len(listen_address) == 0 {
		// set a default address
		listen_address = "127.0.0.1:8080"
	}

  mux := http.NewServeMux()

	// serve static files
	staticPath := "/static/"
  mux.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(http.Dir("./static"))))
	
  mux.HandleFunc("/", index)

  server := &http.Server{
    Addr:     listen_address,
    Handler:  mux,
  }

	// start the server
	fmt.Printf("Listening on %s\n", listen_address)
  check(server.ListenAndServe())
}

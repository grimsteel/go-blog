package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	listen_address := os.Getenv("LISTEN_ADDRESS")
	if len(listen_address) == 0 {
		// set a default address
		listen_address = "127.0.0.1:8080"
	}

	// read posts
	posts := getPostList()

	// store comments in-memory
	comments := make(map[string][]Comment)

  mux := http.NewServeMux()

	// serve static files
	staticPath := "/static/"
  mux.Handle(staticPath, http.StripPrefix(staticPath, http.FileServer(http.Dir("./static"))))
	
  mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		renderTemplate(&posts, "index", w)
	})

	// wildcard recently added
	mux.HandleFunc("/posts/{id}", func (w http.ResponseWriter, r *http.Request) {
		postId := r.PathValue("id")
		// initialize to nil
		var post *Post = nil
		for i := range posts {
			if posts[i].Id == postId {
				post = &posts[i]
				break
			}
		}

		// not found
		if post == nil {
			w.WriteHeader(404)
			renderTemplate(nil, "404", w)
		} else {
			renderTemplate(struct {
				Post *Post
				Comments []Comment
			} {
				Post: post,
				Comments: comments[postId],
			}, "post", w)
		}
	})

	mux.HandleFunc("POST /posts/{id}/comment", func (w http.ResponseWriter, r *http.Request) {
		postId := r.PathValue("id")
		name := r.FormValue("name")
		message := r.FormValue("message")

		// add to comments
		comments[postId] = append(comments[postId], Comment {
			name,
			message,
		})

		// redirect to post
		http.Redirect(w, r, fmt.Sprintf("/posts/%s", postId), http.StatusSeeOther)
	})

  server := &http.Server{
    Addr:     listen_address,
    Handler:  mux,
  }

	// start the server
	fmt.Printf("Listening on %s\n", listen_address)
  check(server.ListenAndServe())
}

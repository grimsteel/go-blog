package main

import (
	"time"
	"os"
	"encoding/json"
)

type Post struct {
	Id string
	Date string
	Filename string
	Title string
}

func formatPostDate(post *Post) {
	parsedDate, err := time.Parse(time.DateOnly, post.Date)
	check(err)

	post.Date = parsedDate.Format("Monday, January _2")
}

func getPostList() ([]Post) {
	postListJson, err := os.ReadFile("posts.json")
	check(err)

	// parse JSON
	var posts []Post
	check(json.Unmarshal(postListJson, &posts))

	// need to use i here because range creates a copy
	for i := range posts {
		formatPostDate(&posts[i])
	}

	return posts
}



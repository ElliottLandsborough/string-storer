package main

import (
	"encoding/json"
	"html"
	"time"
)

type Post struct {
	Title     string
	CreatedAt time.Time
}

var posts []Post

func cleanPost(p *Post) *Post {
	p.Title = html.EscapeString(firstN(p.Title, 100))
	return p
}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func createInitialPost() {
	post := Post{ // b == Student{"Bob", 0}
		Title:     "dynamic string",
		CreatedAt: time.Now().UTC(),
	}

	// Don't append, we only need one title at the moment
	//posts = append(posts, post)
	posts = []Post{post}
}

func createPostFromJson(data []byte) Post {
	// Read post
	post := &Post{}

	err := json.Unmarshal(data, post)
	if err != nil {
		return *post
	}
	post.CreatedAt = time.Now().UTC()

	cleanPost(post)

	// Don't append, we only need one title at the moment
	//posts = append(posts, post)
	posts = []Post{*post}

	return *post
}

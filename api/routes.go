package main

import (
	"io"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, posts)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := createPostFromJson(data)

	if len(post.Title) == 0 {
		responseError(w, err.Error(), http.StatusBadRequest)
	}

	responseJSON(w, post)
}

package main

import (
	"testing"
	"time"
)

var app App

func TestCleanPost(t *testing.T) {
	p := &Post{
		Title:     "something<br />",
		CreatedAt: time.Now().UTC(),
	}

	cleanPost(p)

	actual := p.Title
	expected := "something&lt;br /&gt;"

	if actual != expected {
		t.Errorf("expected 'hello world', got '%s'", actual)
	}
}

func TestCreatePostFromJson(t *testing.T) {
	json := []byte("{\"title\":\"a new title\"}")

	actualPost := createPostFromJson(json)

	expectedPost := &Post{
		Title:     "a new title",
		CreatedAt: time.Now().UTC(),
	}

	actual := actualPost.Title
	expected := expectedPost.Title

	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestCreatePostFromBadJson(t *testing.T) {
	json := []byte("{\"title\":\"a new")

	actualPost := createPostFromJson(json)

	expectedPost := &Post{
		Title:     "",
		CreatedAt: time.Now().UTC(),
	}

	actual := actualPost.Title
	expected := expectedPost.Title

	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

func TestCreateHugePostFromJson(t *testing.T) {
	json := []byte("{\"title\":\"----------------------------------------------------------------------------------------------------Z\"}")

	actualPost := createPostFromJson(json)

	expectedPost := &Post{
		Title:     "----------------------------------------------------------------------------------------------------",
		CreatedAt: time.Now().UTC(),
	}

	actual := actualPost.Title
	expected := expectedPost.Title

	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("template/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	// BUG: Handle .favicon

	SyncPosts()

	temp := make([]*Post, 0)
	for _, p := range posts {
		temp = append(temp, p.Post)
	}

	content := struct {
		Config Config
		Posts  []*Post
	}{
		config,
		temp,
	}

	err := templates.ExecuteTemplate(w, "index", content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Posts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fragment := vars["postFragment"]

	post := GetPost(fragment)

	content := struct {
		Config Config
		Post   *Post
	}{
		config,
		post,
	}

	if post == nil {
		fmt.Println("Not Found: ", fragment)
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := templates.ExecuteTemplate(w, "post", content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Tags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]

	SyncPosts()

	tagPosts := GetPosts(tag)

	content := struct {
		Config Config
		Posts  []*Post
	}{
		config,
		tagPosts,
	}

	err := templates.ExecuteTemplate(w, "tag", content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

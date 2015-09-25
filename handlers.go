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

	content := struct {
		Config Config
		Posts  *[]Post
	}{
		config,
		&posts,
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

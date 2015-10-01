package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates *template.Template

func InitTemplates() {
	var folder = "template/"
	if config.Theme != "" {
		var err error
		folder, err = driver.GetTemplateFolder()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	templates = template.Must(template.ParseGlob(folder + "*"))
}

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

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not Found: ", r.URL)
	content := struct {
		Config Config
	}{
		config,
	}
	err := templates.ExecuteTemplate(w, "404", content)
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
		NotFound(w, r)
	} else {
		err := templates.ExecuteTemplate(w, "post", content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Images(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	image := vars["image"]

	if data, err := driver.GetImage(image); err == nil {
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		if _, err = w.Write(data); err != nil {
			fmt.Println("Write image failure")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
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

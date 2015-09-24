package main

import (
	"net/http"	
	"html/template"
	"github.com/gorilla/mux"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// BUG: Handle .favicon			
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")	
	t.Execute(w, posts)
}

func Posts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fragment := vars["postFragment"]
	
	post := GetPost(fragment)
		
	if post == nil {
		fmt.Println("Not Found: ", fragment)
		w.WriteHeader(http.StatusNotFound)
	} else {
		t := template.New("post.html")
		t, _ = t.ParseFiles("post.html")
		t.Execute(w, *post)
	}	
}
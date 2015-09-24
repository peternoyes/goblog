package main

import (
	"net/http"	
	"html/template"
	"github.com/gorilla/mux"
	"fmt"
)

var templates = template.Must(template.ParseGlob("template/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	// BUG: Handle .favicon				
	err := templates.ExecuteTemplate(w, "index", posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Posts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fragment := vars["postFragment"]
	
	post := GetPost(fragment)
		
	if post == nil {
		fmt.Println("Not Found: ", fragment)
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := templates.ExecuteTemplate(w, "post", post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}	
}
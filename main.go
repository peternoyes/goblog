// Copyright (c) 2015 Peter Noyes

// goblog project main.go
package main

// TODO:
// Paging in index page
// Sort posts by date
// Tags
// Database backend ?
// Cache HTML, have abstract caching mechanism, file system default
// API to update single post
// Images
// Style html
// Theme separate from html?
// File system watcher?
// REST call for posting / updating markdown?

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var config Config

func main() {

	posts = LoadPosts()

	data, err := ioutil.ReadFile("posts/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Title: ", config.Title)

	fmt.Println("Main")
	router := NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	log.Fatal(http.ListenAndServe(":3000", router))
}

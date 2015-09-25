// Copyright (c) 2015 Peter Noyes

// goblog project main.go
package main

// TODO: Minimal Supported Set
// Sort posts by date
// Remove Posts/Config from Repo
// Environment Variable to find posts (File System / S3)
// Images
// Footer with copyright notice and link to github

// TODO: Long Term
// Paging in index page
// Database backend
// Render to static content
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

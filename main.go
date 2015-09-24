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
	"net/http"
	"log"
	"fmt"
)

func main() {
	
	posts = LoadPosts()
	
	fmt.Println("Main")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}

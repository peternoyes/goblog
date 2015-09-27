// Copyright (c) 2015 Peter Noyes

// goblog project main.go
package main

// TODO: Minimal Supported Set
// Sort posts by date
// Images
// Footer with copyright notice and link to github
// 404 Page
// Better date format
// .favicon

// TODO: Long Term
// Paging in index page
// Database backend
// Render to static content
// REST call for posting / updating markdown?

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var config Config
var driver Driver

func main() {
	key := os.Getenv("GOBLOG_DRIVER")
	path := os.Getenv("GOBLOG_DATA")
	region := os.Getenv("GOBLOG_REGION")

	if key == "" || path == "" {
		key = "file"
		path = "posts/"
	}

	fmt.Println("Driver: ", key)
	fmt.Println("Path: ", path)
	fmt.Println("Region: ", region)

	switch key {
	case "aws":
		driver = &DriverS3{path, region, nil}
	default:
		driver = &DriverFile{path}
	}

	driver.New()

	data, err := driver.GetConfig()
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

	LoadPosts()

	fmt.Println("Main")
	router := NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	log.Fatal(http.ListenAndServe(":3000", router))
}

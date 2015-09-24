// Copyright (c) 2015 Peter Noyes

// goblog project main.go
package main

// TODO:
// Separate index page from individual post pages
// Paging in index page
// Sort posts by date
// Tags
// Database backend ?
// Cache HTML, have abstract caching mechanism, file system default
// API to update single post
// Images
// Style html
// Theme separate from html?
// Gorilla mux for better formatting of URLs?
// File system watcher?
// REST call for posting / updating markdown?

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"bufio"
	"bytes"
	"os"
	"log"
	"fmt"
	"github.com/russross/blackfriday"
)

type Post struct {
	Title string
	Date string
	Body template.HTML
	File string
}

func handlerequest(w http.ResponseWriter, r *http.Request){
	// BUG: Handle .favicon
	fmt.Println(r.URL)	
	posts := getPosts()
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")	
	t.Execute(w, posts)
}

func getPosts() []Post{
	fmt.Println("getPosts()")
	a := []Post{}
	files, _ := filepath.Glob("Posts/*")
	for _, f := range files {
		file := strings.Replace(f, "Posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		
		fileStream, err := os.Open(f)		
		if err != nil {
			log.Fatal(err)
		}
		
		defer fileStream.Close()
				
		scanner := bufio.NewScanner(fileStream)
		scanner.Scan()
		firstLine := scanner.Text()
		if firstLine != "---" {
			log.Print("Error")
			continue
		}		
		
		inYaml := true
		var title, date string		
		var buffer bytes.Buffer
		
		for scanner.Scan() {
			line := scanner.Text()					
			if inYaml {
				if line == "---" {
					inYaml = false							
				} else {
					// BUG: Colons in data will cause failure
					tokens := strings.Split(line, ":")					
					if len(tokens) != 2 {
						log.Print("Error")
						continue
					}
										
					switch tokens[0] {
						case "title": title = tokens[1]
						case "date": date = tokens[1]
					}
				}	
			} else {
				buffer.WriteString(line)
				buffer.WriteString("\n")
			}
		}
		
		body := buffer.Bytes()

		markedBody := string(blackfriday.MarkdownCommon(body))		
		a = append(a, Post{title, date, template.HTML(markedBody), file})
	}
	return a;
}

func main() {
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":3000", nil)
}

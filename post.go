// Copyright (c) 2015 Peter Noyes

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var posts []Post

type Post struct {
	UrlFragment string
	Title       string
	Date        time.Time
	Summary     string
	Body        template.HTML
}

func LoadPosts() []Post {
	fmt.Println("getPosts()")
	p := []Post{}
	files, _ := filepath.Glob("posts/*")
	for _, f := range files {
		fileStream, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}

		defer fileStream.Close()

		post := loadPost(fileStream)
		p = append(p, post)

	}
	return p
}

func GetPost(fragment string) *Post {
	u, _ := url.Parse(fragment)
	f := u.EscapedPath()

	for _, post := range posts {
		if post.UrlFragment == f {
			return &post
		}
	}
	return nil
}

func loadPost(reader io.Reader) Post {
	fmt.Println("Load Post")
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	firstLine := scanner.Text()
	if firstLine != "---" {
		log.Print("Error")
	}

	inYaml := true
	var title, summary string
	var date time.Time
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
				case "title":
					title = tokens[1]
				case "date":
					date, _ = time.Parse("2006-Jan-02", tokens[1])
				case "summary":
					summary = tokens[1]
				}
			}
		} else {
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
	}

	body := buffer.Bytes()
	u, _ := url.Parse(title)
	fragment := u.EscapedPath()

	markedBody := string(blackfriday.MarkdownCommon(body))
	return Post{fragment, title, date, summary, template.HTML(markedBody)}
}

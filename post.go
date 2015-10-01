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
	"sort"
	"strings"
	"time"
)

var posts Stubs

type Post struct {
	UrlFragment string
	Title       string
	Date        time.Time
	Summary     string
	Tags        []string
	Excerpt     template.HTML
	Body        template.HTML
}

func (p Post) HasTag(tag string) bool {
	for _, t := range p.Tags {
		if t == tag {
			return true
		}
	}

	return false
}

func LoadPosts() {
	var err error
	posts, err = driver.GlobMarkdown()
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(posts))

	for _, p := range posts {
		fileStream, err := driver.Open(p.Path)
		if err != nil {
			log.Fatal(err)
			continue
		}

		defer fileStream.Close()
		p.Post = loadPost(fileStream)
	}
}

func SyncPosts() {
	tempPosts, err := driver.GlobMarkdown()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range tempPosts {
		fragment := GetUrlFragmentFromTitle(p.Title)
		post := GetStub(fragment)
		if post != nil {
			if p.LastModified != post.LastModified {
				fileStream, err := driver.Open(p.Path)
				if err != nil {
					log.Fatal(err)
					continue
				}
				defer fileStream.Close()
				post.Post = loadPost(fileStream)
			}
		} else {
			fileStream, err := driver.Open(p.Path)
			if err != nil {
				log.Fatal(err)
				continue
			}
			defer fileStream.Close()
			p.Post = loadPost(fileStream)
			posts = append(posts, p)
		}
	}

	sort.Sort(sort.Reverse(posts))
}

func GetStub(fragment string) *PostStub {
	u, _ := url.Parse(fragment)
	f := u.EscapedPath()

	for _, post := range posts {
		if post.Post.UrlFragment == f {
			return post
		}
	}
	return nil
}

func GetPost(fragment string) *Post {
	u, _ := url.Parse(fragment)
	f := u.EscapedPath()

	for _, post := range posts {
		if post.Post.UrlFragment == f {
			return post.Post
		}
	}
	return nil
}

func GetPosts(tag string) []*Post {
	p := make([]*Post, 0)

	for _, post := range posts {
		if post.Post.HasTag(tag) {
			p = append(p, post.Post)
		}
	}

	return p
}

func loadPost(reader io.Reader) *Post {
	fmt.Println("Load Post")
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	firstLine := scanner.Text()
	if firstLine != "---" {
		log.Print("Error")
	}

	inYaml := true
	inExcerpt := true
	var title, summary string
	var date time.Time
	var buffer, excerptBuffer bytes.Buffer
	var tags []string

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
				case "tags":
					tags = strings.Split(tokens[1], ",")
					for i, tag := range tags {
						tags[i] = strings.TrimSpace(tag)
					}
				}
			}
		} else {
			if config.ExcerptTag != "" && line == config.ExcerptTag {
				inExcerpt = false
				buffer.Write(excerptBuffer.Bytes())
			} else {
				if inExcerpt {
					excerptBuffer.WriteString(line)
					excerptBuffer.WriteString("\n")
				} else {
					buffer.WriteString(line)
					buffer.WriteString("\n")
				}
			}
		}
	}

	excerpt := []byte(summary) // Default excerpt
	if inExcerpt {
		buffer = excerptBuffer
	} else {
		excerpt = excerptBuffer.Bytes()
	}

	body := buffer.Bytes()
	fragment := GetUrlFragmentFromTitle(title)

	markedExcerpt := string(blackfriday.MarkdownCommon(excerpt))
	markedBody := string(blackfriday.MarkdownCommon(body))
	return &Post{fragment, title, date, summary, tags, template.HTML(markedExcerpt), template.HTML(markedBody)}
}

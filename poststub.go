package main

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Stubs []*PostStub

type PostStub struct {
	Path         string
	Title        string
	Date         time.Time
	LastModified time.Time
	Post         *Post
}

func (s Stubs) Len() int {
	return len(s)
}

func (s Stubs) Less(i, j int) bool {
	return s[i].Date.Before(s[j].Date)
}

func (s Stubs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func GetUrlFragmentFromTitle(title string) string {
	u, _ := url.Parse(title)
	return u.EscapedPath()
}

func GetDateAndTitleFromFile(file string) (title string, date time.Time, err error) {
	title = ""
	date = time.Now()
	err = nil

	tokens := strings.SplitN(file, "-", 7)
	if len(tokens) != 7 {
		err = errors.New("Malformed string")
		return
	}

	var location *time.Location
	var year, monthNum, day, hour, min, sec int
	var month time.Month

	if year, err = strconv.Atoi(tokens[0]); err != nil {
		return
	}
	if monthNum, err = strconv.Atoi(tokens[1]); err != nil {
		return
	}
	month = time.Month(monthNum)
	if day, err = strconv.Atoi(tokens[2]); err != nil {
		return
	}
	if hour, err = strconv.Atoi(tokens[3]); err != nil {
		return
	}
	if min, err = strconv.Atoi(tokens[4]); err != nil {
		return
	}
	if sec, err = strconv.Atoi(tokens[5]); err != nil {
		return
	}
	if location, err = time.LoadLocation("UTC"); err != nil {
		return
	}

	date = time.Date(year, month, day, hour, min, sec, 0, location)
	title = strings.TrimSuffix(tokens[6], ".md")
	err = nil
	return
}

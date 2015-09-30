// Copyright (c) 2015 Peter Noyes
package main

import (
	"crypto/md5"
	"encoding/hex"
)

type Config struct {
	Title        string   `json:"title"`
	Gravatar     string   `json:"gravatar"`
	Description  string   `json:"description"`
	TopLevelTags []string `json:"topLevelTags"`
	Copyright    string   `json:"copyright"`
	ExcerptTag   string   `json:"excerptTag"`
	Links        []Link   `json:"links"`
}

type Link struct {
	Text string `json:"text" `
	Url  string `json:"url"`
}

func (c Config) GetGravatarURL() string {
	hasher := md5.New()
	hasher.Write([]byte(c.Gravatar))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hasher.Sum(nil)) + "?s=256"
}

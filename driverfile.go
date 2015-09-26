// Copyright (c) 2015 Peter Noyes

package main

import (
	"io"
	"io/ioutil"
	"path/filepath"
	"os"
)

type Driver interface {
	New()
	GetConfig() ([]byte, error)
	GlobMarkdown() ([]string, error)
	Open(string) (io.ReadCloser, error)
}

type DriverFile struct {
	Root string
}

func (d *DriverFile) New() {
	// Noething
}

func (d *DriverFile) GetConfig() ([]byte, error) {
	return ioutil.ReadFile(d.Root +  "config.json")
}

func (d *DriverFile) GlobMarkdown() ([]string, error) {
	return filepath.Glob(d.Root + "*.md")
}

func (d *DriverFile) Open(file string) (io.ReadCloser, error) {
	return os.Open(file)
}
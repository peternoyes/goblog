// Copyright (c) 2015 Peter Noyes

package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Driver interface {
	New()
	GetConfig() ([]byte, error)
	GlobMarkdown() (Stubs, error)
	Open(string) (io.ReadCloser, error)
	GetImage(image string) ([]byte, error)
	GetTemplateFolder() (string, error)
}

type DriverFile struct {
	Root string
}

func (d *DriverFile) New() {
	// Noething
}

func (d *DriverFile) GetConfig() ([]byte, error) {
	return ioutil.ReadFile(d.Root + "config.json")
}

func (d *DriverFile) GetImage(image string) ([]byte, error) {
	return ioutil.ReadFile(d.Root + image)
}

func (d *DriverFile) GetTemplateFolder() (string, error) {
	return config.Theme, nil
}

func (d *DriverFile) GlobMarkdown() (Stubs, error) {
	files, err := filepath.Glob(d.Root + "*.md")
	if err != nil {
		return nil, err
	}

	ret := make(Stubs, 0)

	for _, file := range files {
		key := filepath.Base(file)

		stat, err := os.Stat(file)
		if err != nil {
			return nil, err
		}
		mod := stat.ModTime()

		var title string
		var date time.Time
		title, date, err = GetDateAndTitleFromFile(key)
		if err != nil {
			return nil, err
		}

		ret = append(ret, &PostStub{file, title, date, mod, nil})
	}

	return ret, err
}

func (d *DriverFile) Open(file string) (io.ReadCloser, error) {
	return os.Open(file)
}

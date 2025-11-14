package main

import (
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	return os.WriteFile("data/"+p.Title+".txt", p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"

	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	body, error := os.ReadFile(title + ".txt")

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	page1 := &Page{Title: "Test", Body: []byte("hello!")}
	page1.save()
	page2, error := loadPage("Test")
	if error != nil {
		fmt.Println(string(page2.Body))
	} else {
		fmt.Printf("Encountered an error!!! %v", error.Error())
	}
}

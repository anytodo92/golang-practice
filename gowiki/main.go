package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

var templates = template.Must(template.ParseFiles("view.html", "form.html"))

func loadPage(title string) (*Page, error) {
	body, error := os.ReadFile(title + ".txt")

	if error != nil {
		return nil, error
	}

	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/view/"):]
	page, error := loadPage(title)

	if error != nil {
		// page = &Page{Title: "Test"}
		http.Redirect(w, req, fmt.Sprintf("/edit/%s", title), http.StatusFound)
		return
	}

	renderTemplate("view", w, page)
}

func editHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	page, error := loadPage(title)

	if error != nil {
		// page = &Page{Title: title}
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate("form", w, page)
}

func renderTemplate(fileName string, w http.ResponseWriter, page *Page) {
	// t, error := template.ParseFiles(fmt.Sprintf("%s.html", fileName))

	// if error != nil {
	// 	http.Error(w, error.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// error = t.Execute(w, page)

	error := templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", fileName), page)

	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, req *http.Request) {
	var title string = req.URL.Path[len("/save/"):]
	var body string = req.FormValue("body")
	var page *Page = &Page{Title: title, Body: []byte(body)}

	error := page.save()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, fmt.Sprintf("/view/%v", title), http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

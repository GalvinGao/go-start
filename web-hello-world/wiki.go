package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
)

const LISTEN_ON string = ":8085"
const FILESYSTEM_PREFIX string = "wiki_page_a67bb492_"

type Page struct {
	Title	string
	Body	[]byte
}

func (p *Page) save() error {
	filename := FILESYSTEM_PREFIX + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(PageTitle string) (*Page, error) {
	filename := FILESYSTEM_PREFIX + PageTitle + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err	
	}
	return &Page{Title: PageTitle, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles("template_" + tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.Split(r.URL.Path, "/view/")[1]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.Split(r.URL.Path, "/edit/")[1]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.Split(r.URL.Path, "/save/")[1]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main() {
	fmt.Println("Starting HTTP server at " + string(LISTEN_ON))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	
	log.Fatal(http.ListenAndServe(LISTEN_ON, nil))
}





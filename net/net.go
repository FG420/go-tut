package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	tmpl "example.com/template"

	gowiki "example.com/wiki"
)

var port = ":8080"
var validatePath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func handler(w http.ResponseWriter, r *http.Request) {

	docs, err := filepath.Glob("./*.txt")
	if err != nil {
		fmt.Println("Yikes!")
	}

	for _, file := range docs {

		title := file[:len(file)-len(filepath.Ext(file))]
		p, err := gowiki.LoadPage(title)
		if err != nil {
			fmt.Println("Yikes! x 2")
		}
		tmpl.RenderTemplate(w, "docs", p)
	}
	tmpl.RenderMain(w, "main", "")

}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validatePath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := gowiki.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	tmpl.RenderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := gowiki.LoadPage(title)
	if err != nil {
		p = &gowiki.Page{Title: title}
	}

	tmpl.RenderTemplate(w, "edit", p)
}

func formHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl.RenderMain(w, "create", "")
}

func newDocHandler(w http.ResponseWriter, r *http.Request) {
	docTitle := r.FormValue("title")
	docBody := r.FormValue("body")
	p := &gowiki.Page{Title: docTitle, Body: []byte(docBody)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	p := &gowiki.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {

	log.Println("Server Starting")
	if port != "" {
		log.Println(`Port used: http://localhost` + port)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/create/", formHandler)
	http.HandleFunc("/new/", newDocHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(port, nil))
}

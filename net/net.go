package main

import (
	"log"
	"net/http"
	"regexp"

	tmpl "example.com/template"

	gowiki "example.com/wiki"
)

// var tmpls = template.Must(template.ParseFiles("edit.html", "view.html"))
var validatePath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func renderTemplate(w http.ResponseWriter, tmpl string, p *gowiki.Page) {
// 	err := tmpls.ExecuteTemplate(w, tmpl+".html", p)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there!")

	viewHandler(w, r, "newPage")
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
	// log.Println("Port used:       localhost:8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

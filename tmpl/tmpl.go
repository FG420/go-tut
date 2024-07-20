package tmpl

import (
	"html/template"
	"net/http"

	gowiki "example.com/wiki"
)

var tmpls = template.Must(template.ParseFiles("../tmpl/pages/edit.html",
	"../tmpl/pages/view.html",
	"../tmpl/pages/main.html",
	"../tmpl/pages/create.html",
	"../tmpl/pages/docs.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *gowiki.Page) {
	err := tmpls.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RenderMain(w http.ResponseWriter, tmpl string, data any) {
	err := tmpls.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

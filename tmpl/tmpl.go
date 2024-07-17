package tmpl

import (
	"html/template"
	"net/http"

	gowiki "example.com/wiki"
)

var tmpls = template.Must(template.ParseFiles("../tmpl/edit.html", "../tmpl/view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *gowiki.Page) {
	err := tmpls.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

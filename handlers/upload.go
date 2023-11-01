package handlers

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/upload.html"))

func UploadForm(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

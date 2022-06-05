package controllers

import (
	"net/http"
	"text/template"
)


func AdminAuth(w http.ResponseWriter, r *http.Request) {

	files := []string{"resources/views/admin/auth/login.html"}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, nil)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

package providers

import (
	"net/http"
	"text/template"
)

var (
	layoutFileName string
)

func RenderView(name string, data interface{}, w http.ResponseWriter) {
	
	layoutFileName = "templates/layouts/app.html"
	render(name, data, w)
}

func RenderAdminView(name string, data interface{}, w http.ResponseWriter) {
	layoutFileName = "templates/layouts/admin.html"
	render(name, data, w)
}

func render (name string, data interface{}, w http.ResponseWriter) {
	files := []string{ name, layoutFileName }
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}

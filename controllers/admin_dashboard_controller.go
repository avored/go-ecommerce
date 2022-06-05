package controllers

import (
	"html/template"
	"net/http"

	"github.com/avored/go-ecommerce/providers"
)

//Create a struct that holds information to be displayed in our HTML file
type DashboardViewModel struct {
	providers.BaseFlashes
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// session, _ := providers.GetSessionStore(r)

	dashboardViewModel := DashboardViewModel{}

	files := []string{
		"templates/admin/dashboard.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, dashboardViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

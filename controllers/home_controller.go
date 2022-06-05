package controllers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/avored/go-ecommerce/providers"
)

//Create a struct that holds information to be displayed in our HTML file
type HomeViewModel struct {
	providers.BaseFlashes
	Name   string
	Time   string
	IsAuth func() bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeViewModel := HomeViewModel{Name: "Anonymous", Time: time.Now().Format(time.Stamp)}

	files := []string{
		"templates/home.html",
		"templates/layouts/app.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	if name := r.FormValue("name"); name != "" {
		homeViewModel.Name = name
	}

	homeViewModel.IsAuth = func() bool {
		return true
	}
	err := templates.Execute(w, homeViewModel)

	if err != nil {
		// log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

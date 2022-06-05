package controllers

import (
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

	if name := r.FormValue("name"); name != "" {
		homeViewModel.Name = name
	}

	homeViewModel.IsAuth = func() bool {
		return true
	}
	providers.RenderView("templates/home.html", homeViewModel, w)

}

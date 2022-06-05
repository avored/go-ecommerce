package controllers

import (
	"html/template"
	"net/http"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"github.com/gorilla/mux"
)

//Create a struct that holds information to be displayed in our HTML file
type ProductViewModel struct {
	providers.BaseFlashes
	Product *ent.Product
}

func ProductGetHandler(w http.ResponseWriter, r *http.Request) {

	productModel, _ := services.NewProductOps(r.Context()).ProductGetBySlug(mux.Vars(r)["slug"])

	productIndexViewModel := ProductViewModel{
		Product: productModel,
	}

	productIndexViewModel.Messages = providers.Get(w, r, "messages")
	productIndexViewModel.Errors = providers.Get(w, r, "errors")

	files := []string{
		"templates/catalog/product_view.html",
		"templates/layouts/app.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, productIndexViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

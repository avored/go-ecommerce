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
type CategoryViewModel struct {
	providers.BaseFlashes
	Category         *ent.Category
	CategoryProducts []*ent.Product
}

func CategoryGetHandler(w http.ResponseWriter, r *http.Request) {

	category, _ := services.NewCategoryOps(r.Context()).CategoryGetBySlug(mux.Vars(r)["slug"])
	categoryIndexViewModel := CategoryViewModel{
		Category: category,
	}

	categoryIndexViewModel.CategoryProducts, _ = category.QueryProducts().All(r.Context())

	categoryIndexViewModel.Messages = providers.Get(w, r, "messages")
	categoryIndexViewModel.Errors = providers.Get(w, r, "errors")

	files := []string{
		"templates/catalog/category_view.html",
		"templates/layouts/app.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, categoryIndexViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

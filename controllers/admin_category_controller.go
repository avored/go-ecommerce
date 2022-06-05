package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

//Create a struct that holds information to be displayed in our HTML file
type CategoryIndexViewModel struct {
	providers.BaseFlashes
	Categories []*ent.Category
}

func AdminCategoryGetIndexHandler(w http.ResponseWriter, r *http.Request) {
	pageString := r.FormValue("page")
	pageNumber, _ := strconv.Atoi(pageString)
	itemCount := 10
	offSet := pageNumber * itemCount

	categories, _ := services.NewCategoryOps(r.Context()).CategoryPaginate(itemCount, offSet)

	fmt.Printf("categories %+v\n", categories)

	categoryIndexViewModel := CategoryIndexViewModel{
		Categories: categories,
	}

	categoryIndexViewModel.Messages = providers.Get(w, r, "messages")
	categoryIndexViewModel.Errors = providers.Get(w, r, "errors")

	files := []string{
		"templates/admin/catalog/category/index.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, categoryIndexViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

type CategoryGetCreateViewModel struct {
	providers.BaseFlashes
}

func AdminCategoryGetCreateHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/admin/catalog/category/create.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	category := CategoryGetCreateViewModel{}
	category.Messages = providers.Get(w, r, "messages")
	err := templates.Execute(w, category)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func AdminCategoryPostCreateHandler(w http.ResponseWriter, r *http.Request) {

	var category ent.Category
	category.Name = r.FormValue("name")
	category.Slug = r.FormValue("slug")
	category.MetaTitle = r.FormValue("meta_title")
	category.MetaDescription = r.FormValue("meta_description")

	categoryModel, err := services.NewCategoryOps(r.Context()).CategoryCreate(category)
	if err != nil {
		http.Error(w, "Error while creating an category", http.StatusInternalServerError)
		return
	}

	providers.Set(w, r, "messages", "Category "+categoryModel.Name+"created successfully!")

	NewUrl := "/admin/catalog/category"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)

}

type CategoryGetEditViewModel struct {
	providers.BaseFlashes
	Category *ent.Category
}

func AdminCategoryGetEditHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryId, _ := strconv.Atoi(params["id"])

	categoryModel, err := services.NewCategoryOps(r.Context()).CategoryGetByID(categoryId)
	if err != nil {
		http.Error(w, "Error while fetching a category", http.StatusInternalServerError)
		return
	}

	categoryViewModel := CategoryGetEditViewModel{
		Category: categoryModel,
	}

	files := []string{
		"templates/admin/catalog/category/edit.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	templateError := templates.Execute(w, categoryViewModel)

	if templateError != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func AdminCategoryPostUpdateHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	categoryId, _ := strconv.Atoi(params["id"])

	var category ent.Category

	category.Name = r.FormValue("name")
	category.Slug = r.FormValue("slug")
	category.MetaTitle = r.FormValue("meta_title")
	category.MetaDescription = r.FormValue("meta_description")
	category.ID = categoryId

	categoryUpdated, updateCategoryError := services.NewCategoryOps(r.Context()).CategoryUpdate(category)
	if updateCategoryError != nil {
		http.Error(w, "Error while editing a category", http.StatusInternalServerError)
		return
	}

	providers.Set(w, r, "messages", "Category "+categoryUpdated.Name+" updated successfully!")

	NewUrl := "/admin/catalog/category"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}
func AdminCategoryPostDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	deletedID, err := services.NewCategoryOps(r.Context()).CategoryDelete(id)
	if err != nil {
		http.Error(w, "Error while deleting a category", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Category Deleted %+v\n", deletedID)

	providers.Set(w, r, "errors", "Category destroyed successfully!")

	NewUrl := "/admin/catalog/category"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}

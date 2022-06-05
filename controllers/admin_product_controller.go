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
type ProductIndexViewModel struct {
	providers.BaseFlashes
	Products []*ent.Product
}

func AdminProductGetIndexHandler(w http.ResponseWriter, r *http.Request) {
	pageString := r.FormValue("page")
	pageNumber, _ := strconv.Atoi(pageString)
	itemCount := 10
	offSet := pageNumber * itemCount

	products, _ := services.NewProductOps(r.Context()).ProductPaginate(itemCount, offSet)

	fmt.Printf("products %+v\n", products)

	productIndexViewModel := ProductIndexViewModel{
		Products: products,
	}

	productIndexViewModel.Messages = providers.Get(w, r, "messages")
	productIndexViewModel.Errors = providers.Get(w, r, "errors")

	files := []string{
		"templates/admin/catalog/product/index.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, productIndexViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

type CategoryOptions struct {
	Label      string
	Value      int
	IsSelected bool
}
type ProductGetCreateViewModel struct {
	providers.BaseFlashes
	CategoryOptions []CategoryOptions
}

func AdminProductGetCreateHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/admin/catalog/product/create.html",
		"templates/layouts/admin.html",
	}
	var product = ProductGetCreateViewModel{}

	var categoryOptions CategoryOptions
	allCategories, _ := services.NewCategoryOps(r.Context()).AdminCategoriesGetAll()

	for _, categoryOptionModel := range allCategories {
		categoryOptions.Label = categoryOptionModel.Name
		categoryOptions.Value = categoryOptionModel.ID
		categoryOptions.IsSelected = false

		product.CategoryOptions = append(product.CategoryOptions, categoryOptions)
	}

	templates := template.Must(template.ParseFiles(files...))
	product.Messages = providers.Get(w, r, "messages")
	err := templates.Execute(w, product)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func AdminProductPostCreateHandler(w http.ResponseWriter, r *http.Request) {

	var product ent.Product
	product.Name = r.FormValue("name")
	product.Slug = r.FormValue("slug")
	product.MetaTitle = r.FormValue("meta_title")
	product.MetaDescription = r.FormValue("meta_description")
	var categoryIds []int

	for _, categoryId := range r.Form["categories[]"] {
		intCategoryId, _ := strconv.Atoi(categoryId)
		categoryIds = append(categoryIds, intCategoryId)
	}

	productModel, err := services.NewProductOps(r.Context()).ProductCreate(product, categoryIds)
	if err != nil {
		http.Error(w, "Error while creating an product", http.StatusInternalServerError)
		return
	}

	providers.Set(w, r, "messages", "Product "+productModel.Name+"created successfully!")

	NewUrl := "/admin/catalog/product"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)

}

type ProductGetEditViewModel struct {
	providers.BaseFlashes
	Product         *ent.Product
	CategoryOptions []CategoryOptions
}

func AdminProductGetEditHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId, _ := strconv.Atoi(params["id"])

	productModel, err := services.NewProductOps(r.Context()).ProductGetByID(productId)
	if err != nil {
		http.Error(w, "Error while fetching a product", http.StatusInternalServerError)
		return
	}

	productVEditiewModel := ProductGetEditViewModel{
		Product: productModel,
	}

	var categoryOptions CategoryOptions
	allCategories, _ := services.NewCategoryOps(r.Context()).AdminCategoriesGetAll()

	for _, categoryOptionModel := range allCategories {
		categoryOptions.Label = categoryOptionModel.Name
		categoryOptions.Value = categoryOptionModel.ID
		categoryOptions.IsSelected = false

		for _, productCategoryModel := range productModel.QueryCategories().AllX(r.Context()) {
			if productCategoryModel.ID == categoryOptionModel.ID {
				categoryOptions.IsSelected = true
			}
		}

		productVEditiewModel.CategoryOptions = append(productVEditiewModel.CategoryOptions, categoryOptions)
	}

	files := []string{
		"templates/admin/catalog/product/edit.html",
		"templates/layouts/admin.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	templateError := templates.Execute(w, productVEditiewModel)

	if templateError != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func AdminProductPostUpdateHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	productId, _ := strconv.Atoi(params["id"])

	var product ent.Product

	product.Name = r.FormValue("name")
	product.Slug = r.FormValue("slug")
	product.MetaTitle = r.FormValue("meta_title")
	product.MetaDescription = r.FormValue("meta_description")
	product.ID = productId

	var categoryIds []int

	for _, categoryId := range r.Form["categories[]"] {
		intCategoryId, _ := strconv.Atoi(categoryId)
		categoryIds = append(categoryIds, intCategoryId)
	}

	productUpdated, updateProductError := services.NewProductOps(r.Context()).ProductUpdate(product, categoryIds)
	if updateProductError != nil {
		http.Error(w, "Error while editing a product", http.StatusInternalServerError)
		return
	}

	providers.Set(w, r, "messages", "Product "+productUpdated.Name+" updated successfully!")

	NewUrl := "/admin/catalog/product"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}
func AdminProductPostDeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	deletedID, err := services.NewProductOps(r.Context()).ProductDelete(id)
	if err != nil {
		http.Error(w, "Error while deleting a product", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Product Deleted %+v\n", deletedID)

	providers.Set(w, r, "errors", "Product destroyed successfully!")

	NewUrl := "/admin/catalog/product"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}

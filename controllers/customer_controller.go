package controllers

import (
	"net/http"
	"text/template"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
)

type CustomerDashboardViewModel struct {
	providers.BaseFlashes
	Customer *ent.Customer
}

func MyAccountGetHandler(w http.ResponseWriter, r *http.Request) {
	var customerDashboardViewModel = CustomerDashboardViewModel{}

	session, _ := providers.GetSessionStore(r)
	customerModel, _ := services.NewCustomerOps(r.Context()).GetCustomerByEmail(session.Values["Customer_Email"].(string))

	customerDashboardViewModel.Customer = customerModel
	files := []string{
		"templates/account/dashboard.html",
		"templates/layouts/app.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, customerDashboardViewModel)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}
func MyAccountEditGetHandler(w http.ResponseWriter, r *http.Request) {
	var customerDashboardViewModel = CustomerDashboardViewModel{}

	session, _ := providers.GetSessionStore(r)
	customerModel, _ := services.NewCustomerOps(r.Context()).GetCustomerByEmail(session.Values["Customer_Email"].(string))

	customerDashboardViewModel.Customer = customerModel
	files := []string{
		"templates/account/edit.html",
		"templates/layouts/app.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, customerDashboardViewModel)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}
}
func MyAccountUpdatePostHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := providers.GetSessionStore(r)
	customerModel, _ := services.NewCustomerOps(r.Context()).GetCustomerByEmail(session.Values["Customer_Email"].(string))

	customer := ent.Customer{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     customerModel.Email,
		ID:        customerModel.ID,
	}

	updatedCustomer, _ := services.NewCustomerOps(r.Context()).CustomerUpdate(customer)

	providers.Set(w, r, "messages", "Customer "+updatedCustomer.FirstName+" "+updatedCustomer.LastName+"created successfully!")

	NewUrl := "/my-account"
	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}

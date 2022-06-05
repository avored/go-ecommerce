package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"golang.org/x/crypto/bcrypt"
)

type RegisterViewModel struct {
	providers.BaseFlashes
}

func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	var registerViewModel = RegisterViewModel{}

	files := []string{
		"templates/auth/register.html",
		"templates/layouts/app.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, registerViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	newCustomer := ent.Customer{}

	newCustomer.FirstName = r.FormValue("first_name")
	newCustomer.LastName = r.FormValue("last_name")
	newCustomer.Email = r.FormValue("email")
	newCustomer.Password = r.FormValue("password")

	customerPassword, err := HashPasswordForCustomer(newCustomer.Password)

	if err != nil {
		log.Fatalf("Error while hashing password: %s", err)
	}
	newCustomer.Password = string(customerPassword)

	customer, _ := services.NewCustomerOps(r.Context()).CustomerCreate(newCustomer)

	NewUrl := "/"

	if customer != nil {
		session.Values["Customer_Authenticated"] = true
		session.Values["Customer_FirstName"] = customer.FirstName
		session.Values["Customer_LastName"] = customer.LastName
		session.Values["Customer_Email"] = customer.Email
		session.Save(r, w)
	}

	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}

func HashPasswordForCustomer(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

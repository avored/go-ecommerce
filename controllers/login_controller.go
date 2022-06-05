package controllers

import (
	"html/template"
	"net/http"

	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"golang.org/x/crypto/bcrypt"
)

type LoginViewModel struct {
	providers.BaseFlashes
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	var loginViewModel = LoginViewModel{}

	files := []string{
		"templates/auth/login.html",
		"templates/layouts/app.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, loginViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	customerModel, _ := services.NewCustomerOps(r.Context()).GetCustomerByEmail(r.FormValue("email"))

	match := CheckCustomerPasswordHash(r.FormValue("password"), customerModel.Password)

	NewUrl := "/"

	// Set customer as Authenticated
	if match {
		session.Values["Customer_Authenticated"] = true
		session.Values["Customer_FirstName"] = customerModel.FirstName
		session.Values["Customer_LastName"] = customerModel.LastName
		session.Values["Customer_Email"] = customerModel.Email
		session.Save(r, w)
		NewUrl = "/"
	}

	http.Redirect(w, r, NewUrl, http.StatusSeeOther)
}
func LogoutPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)

	delete(session.Values, "Customer_Authenticated")
	delete(session.Values, "Customer_FirstName")
	delete(session.Values, "Customer_LastName")
	delete(session.Values, "Customer_Email")

	session.Save(r, w)

	newUrl := "/"

	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

func CheckCustomerPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

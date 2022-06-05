package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/address"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"github.com/gorilla/sessions"
)

type CheckoutViewModel struct {
	providers.BaseFlashes
	Session                *sessions.Session
	CartModel              *ent.Cart
	CartItems              []*ent.CartItem
	Customer               *ent.Customer
	DefaultShippingAddress *ent.Address
	DefaultBillingAddress  *ent.Address
}

func CheckoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	var checkoutViewModel = CheckoutViewModel{}

	if session.Values["Customer_Email"] != nil {
		customerModel, _ := services.NewCustomerOps(r.Context()).GetCustomerByEmail(session.Values["Customer_Email"].(string))
		checkoutViewModel.Customer = customerModel

		defaultShippingAddress, _ := customerModel.QueryAddresses().Where(address.IsDefaultEQ(true)).Where(address.TypeEQ("SHIPPING")).Only(r.Context())
		defaultBillingAddress, _ := customerModel.QueryAddresses().Where(address.IsDefaultEQ(true)).Where(address.TypeEQ("Billing")).Only(r.Context())

		checkoutViewModel.DefaultShippingAddress = defaultShippingAddress
		checkoutViewModel.DefaultBillingAddress = defaultBillingAddress
	}

	checkoutViewModel.Session = session
	cartModel, serviceError := services.NewCartOps(r.Context()).CartGetBySessionId(session.Values["ID"].(string))

	if ent.IsNotFound(serviceError) {
		fmt.Println("********   No Items in Cart ******** ")
	}

	checkoutViewModel.CartModel = cartModel

	cartItems, _ := cartModel.QueryCartItems().All(r.Context())
	checkoutViewModel.CartItems = cartItems

	files := []string{
		"templates/checkout/checkout_view.html",
		"templates/layouts/app.html",
	}

	templates := template.Must(template.ParseFiles(files...))

	err := templates.Execute(w, checkoutViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}

}

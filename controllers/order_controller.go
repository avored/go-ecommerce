package controllers

import (
	"fmt"
	"html/template"
	"strconv"

	"net/http"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type OrderViewModel struct {
	providers.BaseFlashes
	Session   *sessions.Session
	CartModel *ent.Cart
	CartItems []*ent.CartItem
	Customer  *ent.Customer
}

func PlaceOrderPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	var customer *ent.Customer
	var newOrder ent.Order

	customerId, _ := strconv.Atoi(r.FormValue("customer[customer_id]"))

	if customerId == 0 {
		var newcustomer ent.Customer

		fmt.Printf("\n********  Created Customer %v ******** ", newcustomer)

	} else {
		dbCustomerModel, customerError := services.NewCustomerOps(r.Context()).CustomerGetByID(customerId)
		if ent.IsNotFound(customerError) {
			http.Error(w, "Customer Not Found", 500)
		}
		customer = dbCustomerModel
	}

	newOrder.CustomerID = customer.ID

	shippingAddressId := r.FormValue("shipping[address_id]")

	if shippingAddressId == "" {
		var shippingAddress ent.Address
		shippingAddress.Type = "SHIPPING"
		shippingAddress.IsDefault = true
		shippingAddress.FirstName = r.FormValue("shipping[first_name]")
		shippingAddress.LastName = r.FormValue("shipping[last_name]")
		shippingAddress.Address1 = r.FormValue("shipping[address1]")
		shippingAddress.Address2 = r.FormValue("shipping[address2]")
		shippingAddress.Suburb = r.FormValue("shipping[suburb]")
		shippingAddress.CompanyName = r.FormValue("shipping[company_name]")
		shippingAddress.City = r.FormValue("shipping[city]")
		shippingAddress.State = r.FormValue("shipping[state]")
		shippingAddress.Postcode = r.FormValue("shipping[postcode]")
		shippingAddress.Phone = r.FormValue("shipping[phone]")
		shippingAddress.CountryID, _ = strconv.Atoi(r.FormValue("shipping[country_id]"))
		shippingAddress.CustomerID = customer.ID

		newShippingAddress, _ := services.NewAddressOps(r.Context()).AddressCreate(shippingAddress)

		fmt.Printf("\n********  New Address Model %v ******** ", newShippingAddress)

		newOrder.ShippingAddressID = newShippingAddress.ID
	} else {
		newOrder.ShippingAddressID, _ = strconv.Atoi(shippingAddressId)
	}

	//@todo add the checkbox condition as well
	billingAddressId := r.FormValue("billing[address_id]")

	if billingAddressId == "" {
		var billingAddress ent.Address
		billingAddress.Type = "BILLING"
		billingAddress.IsDefault = true
		billingAddress.FirstName = r.FormValue("billing[first_name]")
		billingAddress.LastName = r.FormValue("billing[last_name]")
		billingAddress.Address1 = r.FormValue("billing[address1]")
		billingAddress.Address2 = r.FormValue("billing[address2]")
		billingAddress.Suburb = r.FormValue("billing[suburb]")
		billingAddress.CompanyName = r.FormValue("billing[company_name]")
		billingAddress.City = r.FormValue("billing[city]")
		billingAddress.State = r.FormValue("billing[state]")
		billingAddress.Postcode = r.FormValue("billing[postcode]")
		billingAddress.Phone = r.FormValue("billing[phone]")
		billingAddress.CountryID, _ = strconv.Atoi(r.FormValue("billing[country_id]"))
		billingAddress.CustomerID = customer.ID

		newBillingAddress, _ := services.NewAddressOps(r.Context()).AddressCreate(billingAddress)
		newOrder.BillingAddressID = newBillingAddress.ID
	} else {
		newOrder.BillingAddressID, _ = strconv.Atoi(billingAddressId)
	}

	newOrder.PaymentMethod = "CashOnDelivery"
	newOrder.ShippingMethod = "Pickup"

	newOrderModel, _ := services.NewOrderOps(r.Context()).OrderCreate(newOrder)

	cartModel, _ := services.NewCartOps(r.Context()).CartGetBySessionId(session.Values["ID"].(string))

	cartItems, _ := cartModel.QueryCartItems().All(r.Context())

	for _, cartItem := range cartItems {
		var orderProduct ent.OrderProduct
		orderProduct.OrderID = newOrderModel.ID
		orderProduct.ProductID = cartItem.ProductID
		orderProduct.Amount = cartItem.ItemPrice
		orderProduct.Qty = cartItem.Qty
		services.NewOrderProductOps(r.Context()).OrderProductCreate(orderProduct)
	}

	for _, deleteCartItem := range cartItems {
		services.NewCartItemOps(r.Context()).CartItemDelete(deleteCartItem.ID)
	}

	newUrl := "/order/success/" + strconv.Itoa(newOrderModel.ID)

	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

type SuccessOrderViewModel struct {
	providers.BaseFlashes
	Order *ent.Order
}

func SuccessOrderGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId, _ := strconv.Atoi(params["id"])

	orderModel, _ := services.NewOrderOps(r.Context()).OrderGetByID(orderId)

	files := []string{
		"templates/order/success.html",
		"templates/layouts/app.html",
	}

	successOrderViewModel := SuccessOrderViewModel{
		Order: orderModel,
	}
	templates := template.Must(template.ParseFiles(files...))

	templateError := templates.Execute(w, successOrderViewModel)

	if templateError != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}
}

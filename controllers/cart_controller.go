package controllers

import (
	// "fmt"
	"fmt"
	"html/template"
	"net/http"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type EmptyViewModel struct {
	Product   *ent.Product
	Session   *sessions.Session
	CartItem  *ent.CartItem
	CartModel *ent.Cart
}

type CartViewModel struct {
	providers.BaseFlashes
	Session   *sessions.Session
	CartModel *ent.Cart
	CartItems []*ent.CartItem
}

func CartGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	var cartViewModel = CartViewModel{}

	cartViewModel.Session = session
	cartModel, serviceError := services.NewCartOps(r.Context()).CartGetBySessionId(session.Values["ID"].(string))

	if ent.IsNotFound(serviceError) {
		fmt.Println("********   No Items in Cart ******** ")
	}

	fmt.Printf("********  Cart %v ******** ", cartModel)
	cartViewModel.CartModel = cartModel

	cartItems, _ := cartModel.QueryCartItems().All(r.Context())
	cartViewModel.CartItems = cartItems

	files := []string{
		"templates/cart/cart_view.html",
		"templates/layouts/app.html",
	}

	funcMap := template.FuncMap{
		// The name "mul" is what the function will be called in the template text.
		"mul": func(x, y float32) float32 {
			return x * y
		},
	}

	// templates := template.Funcs(funcMap).Must(template.ParseFiles(files...))
	templates := template.Must(template.ParseFiles(files...)).Funcs(funcMap)

	err := templates.Execute(w, cartViewModel)

	if err != nil {
		http.Error(w, "Internal Template Parsing Error", 500)
	}

}
func AddToCartPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := providers.GetSessionStore(r)
	productModel, _ := services.NewProductOps(r.Context()).ProductGetBySlug(mux.Vars(r)["slug"])
	cartModel, _ := services.NewCartOps(r.Context()).CartGetBySessionId(session.Values["ID"].(string))

	if cartModel == nil {
		var cart = ent.Cart{}
		cart.SessionID = session.Values["ID"].(string)
		cart.FullAmount = 56.30
		newCartModel, _ := services.NewCartOps(r.Context()).CartCreate(cart)

		cartModel = newCartModel
	}

	var cartItem = ent.CartItem{}

	cartItemModel, cartItemError := services.NewCartItemOps(r.Context()).GetCartItemByCartAndProductModel(cartModel, productModel)

	if ent.IsNotFound(cartItemError) {
		cartItem.Name = productModel.Name
		cartItem.Qty = 1
		cartItem.ItemPrice = 56.40
		newCartItem, _ := services.NewCartItemOps(r.Context()).CartItemCreate(cartItem, cartModel, productModel)

		cartItemModel = newCartItem
	} else {
		cartItem.Name = cartItemModel.Name
		cartItem.Qty = cartItemModel.Qty + 1
		cartItem.ItemPrice = 56.40
		cartItem.ID = cartItemModel.ID
		cartItemModelUpdated, _ := services.NewCartItemOps(r.Context()).CartItemUpdate(cartItem, cartModel, productModel)
		cartItemModel = cartItemModelUpdated
	}

	newUrl := "/"
	http.Redirect(w, r, newUrl, http.StatusSeeOther)

}

package routers

import (
	"net/http"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/avored/go-ecommerce/providers"
	"github.com/gorilla/mux"
)

func registerAppRouter(r *mux.Router) {

	r.StrictSlash(true)

	// List All Front end routes
	r.HandleFunc("/", providers.MultipleMiddleware(controllers.HomeHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/category/{slug}", providers.MultipleMiddleware(controllers.CategoryGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/product/{slug}", providers.MultipleMiddleware(controllers.ProductGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/add-to-cart/{slug}", providers.MultipleMiddleware(controllers.AddToCartPostHandler, providers.SessionMiddleware)).Methods(http.MethodPost)

	r.HandleFunc("/cart", providers.MultipleMiddleware(controllers.CartGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)

	r.HandleFunc("/checkout", providers.MultipleMiddleware(controllers.CheckoutGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/order/success/{id}", providers.MultipleMiddleware(controllers.SuccessOrderGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)

	r.HandleFunc("/placeorder", providers.MultipleMiddleware(controllers.PlaceOrderPostHandler, providers.SessionMiddleware)).Methods(http.MethodPost)

	r.HandleFunc("/register", providers.MultipleMiddleware(controllers.RegisterGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/register", providers.MultipleMiddleware(controllers.RegisterPostHandler, providers.SessionMiddleware)).Methods(http.MethodPost)

	r.HandleFunc("/login", providers.MultipleMiddleware(controllers.LoginGetHandler, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/login", providers.MultipleMiddleware(controllers.LoginPostHandler, providers.SessionMiddleware)).Methods(http.MethodPost)

	r.HandleFunc("/logout", providers.MultipleMiddleware(controllers.LogoutPostHandler, providers.SessionMiddleware)).Methods(http.MethodPost)

	r.HandleFunc("/my-account", providers.MultipleMiddleware(controllers.MyAccountGetHandler, providers.SessionMiddleware, providers.CustomerAuthMiddleware)).Methods(http.MethodGet)

	r.HandleFunc("/my-account/edit", providers.MultipleMiddleware(controllers.MyAccountEditGetHandler, providers.SessionMiddleware, providers.CustomerAuthMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/my-account/edit", providers.MultipleMiddleware(controllers.MyAccountUpdatePostHandler, providers.SessionMiddleware, providers.CustomerAuthMiddleware)).Methods(http.MethodPost)

}

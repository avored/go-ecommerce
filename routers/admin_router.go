package routers

import (
	"net/http"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/avored/go-ecommerce/providers"
	"github.com/gorilla/mux"
)

func registerAdminUserRouter(r *mux.Router) {

	r.StrictSlash(true)

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("", providers.MultipleMiddleware(controllers.AdminDashboardHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/login", controllers.AdminLoginGetHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc("/login", controllers.AdminLoginPostHandler).Methods(http.MethodPost)

	// List All Route
	r.HandleFunc("/admin/catalog/category", providers.MultipleMiddleware(controllers.AdminCategoryGetIndexHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	// Create Route
	r.HandleFunc("/admin/catalog/category/create", providers.MultipleMiddleware(controllers.AdminCategoryGetCreateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/admin/catalog/category", providers.MultipleMiddleware(controllers.AdminCategoryPostCreateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)
	// Edit Route
	r.HandleFunc("/admin/catalog/category/{id}", providers.MultipleMiddleware(controllers.AdminCategoryGetEditHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/admin/catalog/category/{id}", providers.MultipleMiddleware(controllers.AdminCategoryPostUpdateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)
	// Delete Route
	r.HandleFunc("/admin/catalog/category-delete/{id}", providers.MultipleMiddleware(controllers.AdminCategoryPostDeleteHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)

	// List All Route
	r.HandleFunc("/admin/catalog/product", providers.MultipleMiddleware(controllers.AdminProductGetIndexHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	// Create Route
	r.HandleFunc("/admin/catalog/product/create", providers.MultipleMiddleware(controllers.AdminProductGetCreateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/admin/catalog/product", providers.MultipleMiddleware(controllers.AdminProductPostCreateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)
	// Edit Route
	r.HandleFunc("/admin/catalog/product/{id}", providers.MultipleMiddleware(controllers.AdminProductGetEditHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodGet)
	r.HandleFunc("/admin/catalog/product/{id}", providers.MultipleMiddleware(controllers.AdminProductPostUpdateHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)
	// Delete Route
	r.HandleFunc("/admin/catalog/product-delete/{id}", providers.MultipleMiddleware(controllers.AdminProductPostDeleteHandler, providers.AuthMiddleware, providers.SessionMiddleware)).Methods(http.MethodPost)

	userRouter := r.PathPrefix("/admin-user").Subrouter()
	userRouter.HandleFunc("/profile", controllers.HandleAdminUserProfile).Methods(http.MethodGet)

	// userRouter.HandleFunc("/{id}", controller.UserGetByIDController).Methods(http.MethodGet)
	// userRouter.HandleFunc("/", controller.UserCreateController).Methods(http.MethodPost)
	// userRouter.HandleFunc("/{id}", controller.UserUpdateController).Methods(http.MethodPut)
	// userRouter.HandleFunc("/{id}", controller.UserDeleteController).Methods(http.MethodDelete)
}

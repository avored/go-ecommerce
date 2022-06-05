package providers

import (
	"net/http"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/gorilla/mux"
)
var (
	router *mux.Router
)

// register all available route
func RegisterRouter() *mux.Router {
	router := mux.NewRouter()

	// registerAdminUserRouter(r)
	// registerAppRouter(r)

	router.StrictSlash(true)
	router.HandleFunc("/", controllers.HelloServer).Methods(http.MethodGet)


	
	router.HandleFunc("/admin/login", controllers.AdminAuth).Methods(http.MethodGet)	

	return router
}

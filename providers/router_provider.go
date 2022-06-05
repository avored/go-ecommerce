package providers

import (
	"net/http"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/gorilla/mux"
)
var (
	router *mux.Router
)


func RegisterRouter() *mux.Router {
	router := mux.NewRouter()

	router.StrictSlash(true)
	registerAppRouter(router)
	registerAdminRouter(router)

	return router
}


/**
 * -----------------                               -----------------
 * -----------------                               -----------------
 * -----------------      Register app router    -----------------
 * -----------------                               -----------------
 * -----------------                               -----------------
 */
func registerAppRouter(router *mux.Router) {
	router.HandleFunc("/", controllers.HelloServer).Methods(http.MethodGet)
}




/**
 * -----------------                               -----------------
 * -----------------                               -----------------
 * -----------------     Register ADMIN  router    -----------------
 * -----------------                               -----------------
 * -----------------                               -----------------
 */
func registerAdminRouter(router *mux.Router) {
		router.HandleFunc("/admin/login", controllers.AdminAuth).Methods(http.MethodGet)
}

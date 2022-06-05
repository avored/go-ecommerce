package routers

import (

	"github.com/gorilla/mux"
)

// register all available route
func RegisterRouter(r *mux.Router) {
	registerAdminUserRouter(r)
	registerAppRouter(r)
}

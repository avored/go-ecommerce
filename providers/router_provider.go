package providers

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("/", HelloServer).Methods(http.MethodGet)	

	return router
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World URL: /%s", r.URL.Path[1:])
}

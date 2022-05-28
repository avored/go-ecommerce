package controllers

import (
	"fmt"
	"net/http"
)


func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World URL: /%s", r.URL.Path[1:])
}

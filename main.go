package main

import (
	"log"
	"net/http"
	"time"

	"github.com/avored/go-ecommerce/app/providers"
)

func main() {

    router := providers.RegisterRouter()
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

    port := "8080"
    srv := &http.Server {
		Handler:      router,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}

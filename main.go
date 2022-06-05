package main

import (
	"log"
	"net/http"
	"time"

	"github.com/avored/go-ecommerce/providers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	setupDatabase()

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


func setupDatabase(){
	//initiate Ent Client
	client, err := providers.SetupDatabaseClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}
}

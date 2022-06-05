package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avored/go-ecommerce/cmd"
	"github.com/avored/go-ecommerce/providers"
	"github.com/avored/go-ecommerce/routers"
	"github.com/gorilla/mux"
)

const defaultPort = "8000"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cmd.Execute()

	if err := providers.LoadTranslations(); err != nil {
		log.Fatal(err)
	}

	//initiate Ent Client
	client, err := providers.SetupDatabaseClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client intance to be accessed anywhere through the accessor which is a function
	//named GetClient
	providers.SetDatabaseClient(client)
	providers.SetSession()

	//initiate router and register all the route
	r := mux.NewRouter()
	routers.RegisterRouter(r)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}

package providers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/avored/go-ecommerce/ent"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client
)

func GetDatabaseClient() *ent.Client {
	return client
}

func SetDatabaseClient(newClient *ent.Client) {
	client = newClient
}

func SetupDatabaseClient() (*ent.Client, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

	fmt.Printf("dsn: %s\n", dsn)
	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}

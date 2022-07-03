package main

import (
	// "net/http"
	"log"
	"os"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/avored/go-ecommerce/providers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	providers.SetupServiceProvider()	
	client, err := providers.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()
	server := gin.New()
	server.SetTrustedProxies(nil)

	providers.SetClient(client)

	// loginController := providers.LoginController
	server.POST("/admin/login", controllers.AdminLoginHandler)
	// server.POST("/login", func(ctx *gin.Context) {
	// 	token := providers.LoginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}

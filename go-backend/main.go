package main

import (
	// "net/http"
	"log"
	"os"

	"github.com/avored/go-ecommerce/controllers"
	"github.com/avored/go-ecommerce/middlewares"
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
	cors := middlewares.CORSMiddleware()

	// config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	// config.AllowOrigins = []string{"http://localhost:3000"}

	server.Use(cors)
	server.SetTrustedProxies(nil)

	providers.SetClient(client)

	// loginController := providers.LoginController
	server.POST("/admin/login", controllers.AdminLoginHandler)

	//############### ADMIN USER ROUTES ###############
	server.GET("/admin/user/:id", controllers.GetAdminUserDetails)
	server.POST("/admin/user", controllers.CreateAdminUser)
	server.PUT("/admin/user/:id", controllers.UpdateAdminUser)
	server.DELETE("/admin/user/:id", controllers.DeleteAdminUser)
	
	//############### CATEGORY ROUTES ###############
	server.GET("/admin/category/:id", controllers.GetCategoryDetails)
	server.POST("/admin/category", controllers.CreateCategory)
	server.PUT("/admin/category/:id", controllers.UpdateCategory)
	server.DELETE("/admin/category/:id", controllers.DeleteCategory)



	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}

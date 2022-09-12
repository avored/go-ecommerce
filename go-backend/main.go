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
	oAuth2Middleware := middlewares.GetOauth2Server(server)
	// config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	// config.AllowOrigins = []string{"http://localhost:3000"}

	server.Use(cors)
	
	server.SetTrustedProxies(nil)

	providers.SetClient(client)

	// loginController := providers.LoginController
	server.POST("/admin/login", controllers.AdminLoginHandler)

	adminRoutes := server.Group("/admin")
	{
		adminRoutes.Use(oAuth2Middleware)
		//############### ADMIN USER ROUTES ###############
		adminRoutes.GET("/user/:id", controllers.GetAdminUserDetails)
		adminRoutes.POST("/user", controllers.CreateAdminUser)
		adminRoutes.PUT("/user/:id", controllers.UpdateAdminUser)
		adminRoutes.DELETE("/user/:id", controllers.DeleteAdminUser)
		
		adminRoutes.GET("/category/:id", controllers.GetCategoryDetails)
		adminRoutes.POST("/category", controllers.CreateCategory)
		adminRoutes.PUT("/category/:id", controllers.UpdateCategory)
		adminRoutes.DELETE("/category/:id", controllers.DeleteCategory)
		
		adminRoutes.GET("/role/:id", controllers.GetRoleDetails)
		adminRoutes.POST("/role", controllers.CreateRole)
		adminRoutes.PUT("/role/:id", controllers.UpdateRole)
		adminRoutes.DELETE("/role/:id", controllers.DeleteRole)
		
		adminRoutes.GET("/permission/:id", controllers.GetPermissionDetails)
		adminRoutes.POST("/permission", controllers.CreatePermission)
		adminRoutes.PUT("/permission/:id", controllers.UpdatePermission)
		adminRoutes.DELETE("/permission/:id", controllers.DeletePermission)
	}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}

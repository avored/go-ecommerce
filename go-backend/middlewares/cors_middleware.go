package middlewares

import (
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func CORSMiddleware() gin.HandlerFunc {
    config := cors.DefaultConfig()

    corsBackendUrl := os.Getenv("CORS_BACKEND_BASE_URL")

	config.AllowMethods = []string{"POST", "GET", "PUT", "DELETE"}
	config.AllowOrigins = []string{corsBackendUrl}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Authorization", "Content-Length", "Accept-Encoding"}
	config.AllowCredentials = true
    return cors.New(config)
}

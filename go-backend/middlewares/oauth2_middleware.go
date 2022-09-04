package middlewares

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)


func GetOauth2Server(g *gin.Engine) gin.HandlerFunc {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())


	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	auth := g.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	tokenVerification := ginserver.HandleTokenVerify()
	return tokenVerification
}

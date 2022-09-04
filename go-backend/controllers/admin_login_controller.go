package controllers

import (
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	// "github.com/go-oauth2/gin-server"
	// "github.com/go-oauth2/oauth2/v4"
	// "github.com/go-oauth2/oauth2/v4/manage"
	// "github.com/go-oauth2/oauth2/v4/models"
	// "github.com/go-oauth2/oauth2/v4/server"
	// "github.com/go-oauth2/oauth2/v4/store"
	"golang.org/x/crypto/bcrypt"
)

type LoginJSON struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func AdminLoginHandler(ctx *gin.Context) {
	request := LoginJSON{}
	ctx.Bind(&request)

	adminUserModel := services.FetchAdminUserByEmail(ctx, request.Email)

	match := CheckPasswordHash(request.Password, adminUserModel.Password)
	var token string = "INVALID_TOKEN"
	if match {
		// manager := manage.NewDefaultManager()
		// manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

		// manager.MustTokenStorage(store.NewMemoryTokenStore())

		// clientStore := store.NewClientStore()
		// manager.MapClientStorage(clientStore)

		// srv := server.NewDefaultServer(manager)
		// srv.SetAllowGetAccessRequest(true)
		// srv.SetClientInfoHandler(server.ClientFormHandler)
		// manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

		token = "my_demo_valid_token_add_middleware_jwt_token_via_gin_framework"
	}

	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
			"token":      token,
		},
	)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

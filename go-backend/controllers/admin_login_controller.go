package controllers

import (
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginJSON struct {
    Email     string `form:"email" binding:"required"`
    Password string `form:"password" binding:"required"`
}


func AdminLoginHandler(ctx *gin.Context) {
	request := LoginJSON{}
	ctx.Bind(&request)
	
	adminUserModel := services.FetchAdminUserByEmail(ctx, request.Email)

	match := CheckPasswordHash(request.Password, adminUserModel.Password)
	var token string = "INVALID_TOKEN"
	if match {
		token = "my_demo_valid_token_add_middleware_jwt_token_via_gin_framework"
	}

	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
			"token":      token,
		},
	)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

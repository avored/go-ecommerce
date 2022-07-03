package controllers

import (

	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

func AdminLoginHandler(ctx *gin.Context) {
	email := "admin@admin.com"
	adminUserModel := services.FetchAdminUserByEmail(ctx, email)
	
	// token := "sjsjhekhjskjhsnjkdhjiskjhsjklepoisoiuapol7"
	// AdminUSerModel := "tests"
	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
		},
	)
}

package controllers

import (
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type AdminUserIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

func GetAdminUserDetails(ctx *gin.Context) {
	var adminUserRequest AdminUserIDRequest
	if err := ctx.ShouldBindUri(&adminUserRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	adminUserModel := services.FetchAdminUserByID(ctx, adminUserRequest.ID)

	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
		},
	)
}

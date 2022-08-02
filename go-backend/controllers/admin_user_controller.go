package controllers

import (
	"github.com/avored/go-ecommerce/ent"
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

type AdminUserUpdateRequest struct {
    FirstName string `form:"first_name" binding:"required"`
    LastName string `form:"last_name" binding:"required"`
}

func UpdateAdminUserDetails(ctx *gin.Context) {
	var adminUserRequest AdminUserIDRequest
	if err := ctx.ShouldBindUri(&adminUserRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	request := AdminUserUpdateRequest{}
	ctx.Bind(&request)


	var adminUser ent.AdminUser
	adminUser.ID = adminUserRequest.ID
	adminUser.FirstName = request.FirstName
	adminUser.LastName = request.LastName

	adminUserModel := services.UpdateAdminUserByID(ctx, adminUserRequest.ID, adminUser)

	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
		},
	)
}

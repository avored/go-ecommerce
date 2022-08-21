package controllers

import (
	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type PermissionIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

func GetPermissionDetails(ctx *gin.Context) {
	var getPermissionRequest PermissionIDRequest
	if err := ctx.ShouldBindUri(&getPermissionRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	adminUserModel := services.FetchPermissionByID(ctx, getPermissionRequest.ID)

	ctx.JSON(200,
		gin.H{
			"permission": adminUserModel,
		},
	)
}

type CreatePermissionRequest struct {
	Name       		string `form:"name" binding:"required"`
	Identifier      string `form:"identifier" binding:"required"`
	Description		string `form:"description"`
}

func CreatePermission(ctx *gin.Context) {
	var createPermissionRequest CreatePermissionRequest
	ctx.Bind(&createPermissionRequest)

	var entCreatePermission ent.Permission
	entCreatePermission.Name = createPermissionRequest.Name
	entCreatePermission.Identifier = createPermissionRequest.Identifier
	entCreatePermission.Description = createPermissionRequest.Description

	adminPermissionModel := services.CreatePermission(ctx, entCreatePermission)

	ctx.JSON(200,
		gin.H{
			"permission": adminPermissionModel,
		},
	)
}

type PermissionUpdateRequest struct {
	Name string `form:"name" binding:"required"`
	Identifier  string `form:"identifier" binding:"required"`
	Description  string `form:"description"`
}

func UpdatePermission(ctx *gin.Context) {
	var updatePermissionRequest PermissionIDRequest
	if err := ctx.ShouldBindUri(&updatePermissionRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	request := PermissionUpdateRequest{}
	ctx.Bind(&request)

	var permission ent.Permission
	permission.ID = updatePermissionRequest.ID
	permission.Name = request.Name
	permission.Identifier = request.Identifier
	permission.Description = request.Description

	permissionModel := services.UpdatePermissionByID(ctx, updatePermissionRequest.ID, permission)

	ctx.JSON(200,
		gin.H{
			"permission": permissionModel,
		},
	)
}



type DeletePermissionRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeletePermission(ctx *gin.Context) {
	var deletePermissionRequest DeletePermissionRequest
	if err := ctx.ShouldBindUri(&deletePermissionRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	deletePermissionId := services.DeletePermissionByID(ctx, deletePermissionRequest.ID)

	ctx.JSON(200,
		gin.H{
			"delete_permission_id": deletePermissionId,
		},
	)
}

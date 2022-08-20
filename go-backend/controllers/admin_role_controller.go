package controllers

import (
	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type RoleIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

func GetRoleDetails(ctx *gin.Context) {
	var getRoleRequest RoleIDRequest
	if err := ctx.ShouldBindUri(&getRoleRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	adminUserModel := services.FetchRoleByID(ctx, getRoleRequest.ID)

	ctx.JSON(200,
		gin.H{
			"role": adminUserModel,
		},
	)
}

type CreateRoleRequest struct {
	Name       		string `form:"name" binding:"required"`
	Identifier      string `form:"identifier" binding:"required"`
	Description		string `form:"description"`
}

func CreateRole(ctx *gin.Context) {
	var createRoleRequest CreateRoleRequest
	ctx.Bind(&createRoleRequest)

	var entCreateRole ent.Role
	entCreateRole.Name = createRoleRequest.Name
	entCreateRole.Identifier = createRoleRequest.Identifier
	entCreateRole.Description = createRoleRequest.Description

	adminRoleModel := services.CreateRole(ctx, entCreateRole)

	ctx.JSON(200,
		gin.H{
			"role": adminRoleModel,
		},
	)
}

type RoleUpdateRequest struct {
	Name string `form:"name" binding:"required"`
	Identifier  string `form:"identifier" binding:"required"`
	Description  string `form:"description"`
}

func UpdateRole(ctx *gin.Context) {
	var updateRoleRequest RoleIDRequest
	if err := ctx.ShouldBindUri(&updateRoleRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	request := RoleUpdateRequest{}
	ctx.Bind(&request)

	var role ent.Role
	role.ID = updateRoleRequest.ID
	role.Name = request.Name
	role.Identifier = request.Identifier
	role.Description = request.Description

	roleModel := services.UpdateRoleByID(ctx, updateRoleRequest.ID, role)

	ctx.JSON(200,
		gin.H{
			"role": roleModel,
		},
	)
}



type DeleteRoleRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeleteRole(ctx *gin.Context) {
	var deleteRoleRequest DeleteRoleRequest
	if err := ctx.ShouldBindUri(&deleteRoleRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	deleteRoleId := services.DeleteRoleByID(ctx, deleteRoleRequest.ID)

	ctx.JSON(200,
		gin.H{
			"delete_role_id": deleteRoleId,
		},
	)
}

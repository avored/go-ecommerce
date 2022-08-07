package controllers

import (
	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

type CreateAdminUserRequest struct {
	FirstName       string `form:"first_name" binding:"required"`
	LastName        string `form:"last_name" binding:"required"`
	Email           string `form:"email" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required"`
}

func CreateAdminUser(ctx *gin.Context) {
	var createAdminUserRequest CreateAdminUserRequest
	ctx.Bind(&createAdminUserRequest)

	/*	Hash password	*/
	password, _ := HashPassword(createAdminUserRequest.Password)

	var entCreateAdminUser ent.AdminUser
	entCreateAdminUser.FirstName = createAdminUserRequest.FirstName
	entCreateAdminUser.LastName = createAdminUserRequest.LastName
	entCreateAdminUser.Email = createAdminUserRequest.Email
	entCreateAdminUser.Password = password

	adminUserModel := services.CreateAdminUser(ctx, entCreateAdminUser)

	ctx.JSON(200,
		gin.H{
			"admin_user": adminUserModel,
		},
	)
}

type AdminUserUpdateRequest struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
}

func UpdateAdminUser(ctx *gin.Context) {
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



type DeleteAdminUserRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeleteAdminUser(ctx *gin.Context) {
	var deleteAdminUserRequest DeleteAdminUserRequest
	if err := ctx.ShouldBindUri(&deleteAdminUserRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	deleteAdminUserId := services.DeleteAdminUserByID(ctx, deleteAdminUserRequest.ID)

	ctx.JSON(200,
		gin.H{
			"delete_admin_user_id": deleteAdminUserId,
		},
	)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
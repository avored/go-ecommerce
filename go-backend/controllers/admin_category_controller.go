package controllers

import (
	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/services"
	"github.com/gin-gonic/gin"
)

type CategoryIDRequest struct {
	ID int `uri:"id" binding:"required"`
}

func GetCategoryDetails(ctx *gin.Context) {
	var getCategoryRequest CategoryIDRequest
	if err := ctx.ShouldBindUri(&getCategoryRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	adminUserModel := services.FetchCategoryByID(ctx, getCategoryRequest.ID)

	ctx.JSON(200,
		gin.H{
			"category": adminUserModel,
		},
	)
}

type CreateCategoryRequest struct {
	Name       		string `form:"name" binding:"required"`
	Identifier      string `form:"identifier" binding:"required"`
	MetaTitle       string `form:"meta_title"`
	MetaDescription	string `form:"meta_description"`
}

func CreateCategory(ctx *gin.Context) {
	var createCategoryRequest CreateCategoryRequest
	ctx.Bind(&createCategoryRequest)

	var entCreateCategory ent.Category
	entCreateCategory.Name = createCategoryRequest.Name
	entCreateCategory.Identifier = createCategoryRequest.Identifier
	entCreateCategory.MetaTitle = createCategoryRequest.MetaTitle
	entCreateCategory.MetaDescription = createCategoryRequest.MetaDescription

	adminCategoryModel := services.CreateCategory(ctx, entCreateCategory)

	ctx.JSON(200,
		gin.H{
			"category": adminCategoryModel,
		},
	)
}

type CategoryUpdateRequest struct {
	Name string `form:"name" binding:"required"`
	Identifier  string `form:"identifier" binding:"required"`
	MetaTitle  string `form:"meta_title"`
	MetaDescription  string `form:"meta_description"`
}

func UpdateCategory(ctx *gin.Context) {
	var updateCategoryRequest CategoryIDRequest
	if err := ctx.ShouldBindUri(&updateCategoryRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	request := CategoryUpdateRequest{}
	ctx.Bind(&request)

	var category ent.Category
	category.ID = updateCategoryRequest.ID
	category.Name = request.Name
	category.Identifier = request.Identifier
	category.MetaTitle = request.MetaTitle
	category.MetaDescription = request.MetaDescription

	categoryModel := services.UpdateCategoryByID(ctx, updateCategoryRequest.ID, category)

	ctx.JSON(200,
		gin.H{
			"category": categoryModel,
		},
	)
}



type DeleteCategoryRequest struct {
	ID int `uri:"id" binding:"required"`
}

func DeleteCategory(ctx *gin.Context) {
	var deleteCategoryRequest DeleteCategoryRequest
	if err := ctx.ShouldBindUri(&deleteCategoryRequest); err != nil {
		ctx.JSON(400, gin.H{"msg": err})
		return
	}

	deleteCategoryId := services.DeleteCategoryByID(ctx, deleteCategoryRequest.ID)

	ctx.JSON(200,
		gin.H{
			"delete_category_id": deleteCategoryId,
		},
	)
}

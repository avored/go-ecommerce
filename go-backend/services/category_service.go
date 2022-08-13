package services

import (
	"context"
	"log"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/repository"
)


// func FetchAdminUserByEmail (ctx context.Context, email string) *ent.AdminUser {
// 	adminUserModel, err := repository.NewAdminUserOps(ctx).AdminUserGetByEmail(email)
// 	if err != nil {
// 		log.Printf("err : %s", err)
// 		return nil
// 	}

// 	return adminUserModel
// }


func CreateCategory (ctx context.Context, createCategory ent.Category) *ent.Category {
	createCategoryModel, err := repository.NewCategoryOps(ctx).CreateCategory(createCategory)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return createCategoryModel
}

func FetchCategoryByID (ctx context.Context, id int) *ent.Category {
	fetchCategoryModel, err := repository.NewCategoryOps(ctx).CategoryGetByID(id)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return fetchCategoryModel
}
func UpdateCategoryByID (ctx context.Context, id int, category ent.Category) *ent.Category {
	updatedCategoryModel, err := repository.NewCategoryOps(ctx).UpdateCategoryById(id, category)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return updatedCategoryModel
}


func DeleteCategoryByID (ctx context.Context, id int) int {
	deleteCategoryId, err := repository.NewCategoryOps(ctx).CategoryDelete(id)
	if err != nil {
		log.Printf("err : %s", err)
		return 0
	}

	return deleteCategoryId
}

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


func CreatePermission (ctx context.Context, createPermission ent.Permission) *ent.Permission {
	createPermissionModel, err := repository.NewPermissionOps(ctx).CreatePermission(createPermission)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return createPermissionModel
}

func FetchPermissionByID (ctx context.Context, id int) *ent.Permission {
	fetchPermissionModel, err := repository.NewPermissionOps(ctx).PermissionGetByID(id)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return fetchPermissionModel
}
func UpdatePermissionByID (ctx context.Context, id int, permission ent.Permission) *ent.Permission {
	updatedPermissionModel, err := repository.NewPermissionOps(ctx).UpdatePermissionById(id, permission)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return updatedPermissionModel
}


func DeletePermissionByID (ctx context.Context, id int) int {
	deletePermissionId, err := repository.NewPermissionOps(ctx).PermissionDelete(id)
	if err != nil {
		log.Printf("err : %s", err)
		return 0
	}

	return deletePermissionId
}

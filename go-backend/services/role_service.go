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


func CreateRole (ctx context.Context, createRole ent.Role) *ent.Role {
	createRoleModel, err := repository.NewRoleOps(ctx).CreateRole(createRole)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return createRoleModel
}

func FetchRoleByID (ctx context.Context, id int) *ent.Role {
	fetchRoleModel, err := repository.NewRoleOps(ctx).RoleGetByID(id)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

return fetchRoleModel
}
func UpdateRoleByID (ctx context.Context, id int, role ent.Role) *ent.Role {
	updatedRoleModel, err := repository.NewRoleOps(ctx).UpdateRoleById(id, role)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return updatedRoleModel
}


func DeleteRoleByID (ctx context.Context, id int) int {
	deleteRoleId, err := repository.NewRoleOps(ctx).RoleDelete(id)
	if err != nil {
		log.Printf("err : %s", err)
		return 0
	}

	return deleteRoleId
}

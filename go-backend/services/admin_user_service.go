package services

import (
	"context"
	"log"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/repository"
)


func FetchAdminUserByEmail (ctx context.Context, email string) *ent.AdminUser {
	adminUserModel, err := repository.NewAdminUserOps(ctx).AdminUserGetByEmail(email)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return adminUserModel
}


func CreateAdminUser (ctx context.Context, createAdminUser ent.AdminUser) *ent.AdminUser {
	adminUserModel, err := repository.NewAdminUserOps(ctx).CreateAdminUser(createAdminUser)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return adminUserModel
}

func FetchAdminUserByID (ctx context.Context, id int) *ent.AdminUser {
	adminUserModel, err := repository.NewAdminUserOps(ctx).AdminUserGetByID(id)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return adminUserModel
}
func UpdateAdminUserByID (ctx context.Context, id int, adminUser ent.AdminUser) *ent.AdminUser {
	adminUserModel, err := repository.NewAdminUserOps(ctx).UpdateAdminUserById(id, adminUser)
	if err != nil {
		log.Printf("err : %s", err)
		return nil
	}

	return adminUserModel
}


func DeleteAdminUserByID (ctx context.Context, id int) int {
	deleteAdminUserId, err := repository.NewAdminUserOps(ctx).AdminUserDelete(id)
	if err != nil {
		log.Printf("err : %s", err)
		return 0
	}

	return deleteAdminUserId
}

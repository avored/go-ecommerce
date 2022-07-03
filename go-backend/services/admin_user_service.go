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

package services

import (
	"context"
	"time"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/adminuser"
	"github.com/avored/go-ecommerce/providers"
)

type AdminUserOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewAdminUserOps(ctx context.Context) *AdminUserOps {
	return &AdminUserOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}

func (r *AdminUserOps) AdminAdminUsersGetAll() ([]*ent.AdminUser, error) {

	adminUsers, err := r.client.AdminUser.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return adminUsers, nil
}

func (r *AdminUserOps) AdminUserGetByID(id int) (*ent.AdminUser, error) {

	user, err := r.client.AdminUser.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AdminUserOps) AdminUserGetByEmail(email string) (*ent.AdminUser, error) {

	user, err := r.client.AdminUser.Query().Where(adminuser.EmailEQ(email)).Only(r.ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}

	return user, nil
}

func (r *AdminUserOps) AdminUserCreate(newAdminUser ent.AdminUser) (*ent.AdminUser, error) {

	newCreatedAdminUser, err := r.client.AdminUser.Create().
		SetEmail(newAdminUser.Email).
		SetFirstName(newAdminUser.FirstName).
		SetLastName(newAdminUser.LastName).
		SetPassword(newAdminUser.Password).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedAdminUser, nil
}

func (r *AdminUserOps) AdminUserUpdate(user ent.AdminUser) (*ent.AdminUser, error) {

	updatedAdminUser, err := r.client.AdminUser.UpdateOneID(user.ID).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetUpdatedAt(time.Now()).
		SetLastName(user.LastName).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedAdminUser, nil
}

func (r *AdminUserOps) AdminUserDelete(id int) (int, error) {

	err := r.client.AdminUser.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

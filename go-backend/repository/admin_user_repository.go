package repository

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/adminuser"
	"github.com/avored/go-ecommerce/providers"
)

type AdminUserOps struct {
	ctx context.Context
	client *ent.Client
}

func NewAdminUserOps(ctx context.Context)  *AdminUserOps {
	return &AdminUserOps{
		ctx: ctx,
		client: providers.GetClient(),
	}
}

func (r *AdminUserOps) AdminUserGetByID(id int) (*ent.AdminUser, error) {

	adminUser, err := r.client.AdminUser.Query().
		Where(adminuser.IDEQ(id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return adminUser, nil
}
func (r *AdminUserOps) AdminUserGetByEmail(email string) (*ent.AdminUser, error) {

	adminUser, err := r.client.AdminUser.Query().
		Where(adminuser.EmailEQ(email)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return adminUser, nil
}

func (r *AdminUserOps) CreateAdminUser(createAdminUser ent.AdminUser) (*ent.AdminUser, error) {

	createdAdminUserModel, err := r.client.AdminUser.
						Create().
						SetFirstName(createAdminUser.FirstName).
						SetLastName(createAdminUser.LastName).
						SetEmail(createAdminUser.Email).
						SetPassword(createAdminUser.Password).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return createdAdminUserModel, nil
}
func (r *AdminUserOps) UpdateAdminUserById(id int, updateAdminUser ent.AdminUser) (*ent.AdminUser, error) {

	adminUserModel, err := r.client.AdminUser.
						UpdateOneID(id).
						SetFirstName(updateAdminUser.FirstName).
						SetLastName(updateAdminUser.LastName).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return adminUserModel, nil
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

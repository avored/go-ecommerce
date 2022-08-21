package repository

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/permission"
	"github.com/avored/go-ecommerce/providers"
)

type PermissionOps struct {
	ctx context.Context
	client *ent.Client
}

func NewPermissionOps(ctx context.Context)  *PermissionOps {
	return &PermissionOps{
		ctx: ctx,
		client: providers.GetClient(),
	}
}

func (r *PermissionOps) PermissionGetByID(id int) (*ent.Permission, error) {

	permissionModel, err := r.client.Permission.Query().
		Where(permission.IDEQ(id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return permissionModel, nil
}
// func (r *PermissionOps) PermissionGetByEmail(email string) (*ent.Permission, error) {

// 	adminUser, err := r.client.Permission.Query().
// 		Where(adminuser.EmailEQ(email)).
// 		Only(r.ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return adminUser, nil
// }

func (r *PermissionOps) CreatePermission(createPermission ent.Permission) (*ent.Permission, error) {

	createdPermissionModel, err := r.client.Permission.
						Create().
						SetName(createPermission.Name).
						SetIdentifier(createPermission.Identifier).
						SetDescription(createPermission.Description).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return createdPermissionModel, nil
}
func (r *PermissionOps) UpdatePermissionById(id int, updatePermission ent.Permission) (*ent.Permission, error) {

	updatedPermissionModel, err := r.client.Permission.
						UpdateOneID(id).
						SetName(updatePermission.Name).
						SetIdentifier(updatePermission.Identifier).
						SetDescription(updatePermission.Description).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return updatedPermissionModel, nil
}


func (r *PermissionOps) PermissionDelete(id int) (int, error) {

	err := r.client.Permission.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

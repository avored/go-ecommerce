package repository

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/role"
	"github.com/avored/go-ecommerce/providers"
)

type RoleOps struct {
	ctx context.Context
	client *ent.Client
}

func NewRoleOps(ctx context.Context)  *RoleOps {
	return &RoleOps{
		ctx: ctx,
		client: providers.GetClient(),
	}
}

func (r *RoleOps) RoleGetByID(id int) (*ent.Role, error) {

	roleModel, err := r.client.Role.Query().
		Where(role.IDEQ(id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return roleModel, nil
}
// func (r *RoleOps) RoleGetByEmail(email string) (*ent.Role, error) {

// 	adminUser, err := r.client.Role.Query().
// 		Where(adminuser.EmailEQ(email)).
// 		Only(r.ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return adminUser, nil
// }

func (r *RoleOps) CreateRole(createRole ent.Role) (*ent.Role, error) {

	createdRoleModel, err := r.client.Role.
						Create().
						SetName(createRole.Name).
						SetIdentifier(createRole.Identifier).
						SetDescription(createRole.Description).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return createdRoleModel, nil
}
func (r *RoleOps) UpdateRoleById(id int, updateRole ent.Role) (*ent.Role, error) {

	updatedRoleModel, err := r.client.Role.
						UpdateOneID(id).
						SetName(updateRole.Name).
						SetIdentifier(updateRole.Identifier).
						SetDescription(updateRole.Description).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return updatedRoleModel, nil
}


func (r *RoleOps) RoleDelete(id int) (int, error) {

	err := r.client.Role.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

package repository

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/category"
	"github.com/avored/go-ecommerce/providers"
)

type CategoryOps struct {
	ctx context.Context
	client *ent.Client
}

func NewCategoryOps(ctx context.Context)  *CategoryOps {
	return &CategoryOps{
		ctx: ctx,
		client: providers.GetClient(),
	}
}

func (r *CategoryOps) CategoryGetByID(id int) (*ent.Category, error) {

	categoryModel, err := r.client.Category.Query().
		Where(category.IDEQ(id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return categoryModel, nil
}
// func (r *CategoryOps) CategoryGetByEmail(email string) (*ent.Category, error) {

// 	adminUser, err := r.client.Category.Query().
// 		Where(adminuser.EmailEQ(email)).
// 		Only(r.ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return adminUser, nil
// }

func (r *CategoryOps) CreateCategory(createCategory ent.Category) (*ent.Category, error) {

	createdCategoryModel, err := r.client.Category.
						Create().
						SetName(createCategory.Name).
						SetIdentifier(createCategory.Identifier).
						SetMetaTitle(createCategory.MetaTitle).
						SetMetaDescription(createCategory.MetaDescription).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return createdCategoryModel, nil
}
func (r *CategoryOps) UpdateCategoryById(id int, updateCategory ent.Category) (*ent.Category, error) {

	updatedCategoryModel, err := r.client.Category.
						UpdateOneID(id).
						SetName(updateCategory.Name).
						SetIdentifier(updateCategory.Identifier).
						SetMetaTitle(updateCategory.MetaTitle).
						SetMetaDescription(updateCategory.MetaDescription).
						Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return updatedCategoryModel, nil
}


func (r *CategoryOps) CategoryDelete(id int) (int, error) {

	err := r.client.Category.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

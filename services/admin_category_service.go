package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/category"
	"github.com/avored/go-ecommerce/providers"
)

type CategoryOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCategoryOps(ctx context.Context) *CategoryOps {
	return &CategoryOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *CategoryOps) CategoryPaginate(limit int, offset int) ([]*ent.Category, error) {

	categories, err := r.client.Category.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryOps) AdminCategoriesGetAll() ([]*ent.Category, error) {

	categories, err := r.client.Category.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryOps) CategoryGetByID(id int) (*ent.Category, error) {

	categoryModel, err := r.client.Category.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return categoryModel, nil
}

func (r *CategoryOps) CategoryGetBySlug(slug string) (*ent.Category, error) {

	categoryModel, err := r.client.Category.Query().WithProducts().Where(category.SlugEQ(slug)).Only(r.ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}

	return categoryModel, nil
}

func (r *CategoryOps) CategoryCreate(newCategory ent.Category) (*ent.Category, error) {

	newCreatedCategory, err := r.client.Category.Create().
		SetName(newCategory.Name).
		SetSlug(newCategory.Slug).
		SetMetaTitle(newCategory.MetaTitle).
		SetMetaDescription(newCategory.MetaDescription).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedCategory, nil
}

func (r *CategoryOps) CategoryUpdate(category ent.Category) (*ent.Category, error) {

	updatedCategory, err := r.client.Category.UpdateOneID(category.ID).
		SetName(category.Name).
		SetSlug(category.Slug).
		SetMetaTitle(category.MetaTitle).
		SetMetaDescription(category.MetaDescription).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
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

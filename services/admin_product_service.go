package services

import (
	"context"
	"fmt"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/product"
	"github.com/avored/go-ecommerce/providers"
)

type ProductOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewProductOps(ctx context.Context) *ProductOps {
	return &ProductOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *ProductOps) ProductPaginate(limit int, offset int) ([]*ent.Product, error) {

	products, err := r.client.Product.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductOps) AdminProductsGetAll() ([]*ent.Product, error) {

	products, err := r.client.Product.Query().WithCategories().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductOps) ProductGetByID(id int) (*ent.Product, error) {

	productModel, err := r.client.Product.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return productModel, nil
}

func (r *ProductOps) ProductGetBySlug(slug string) (*ent.Product, error) {

	productModel, err := r.client.Product.Query().Where(product.SlugEQ(slug)).Only(r.ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}

	return productModel, nil
}

func (r *ProductOps) ProductCreate(newProduct ent.Product, categoryIds []int) (*ent.Product, error) {

	newCreatedProduct, err := r.client.Product.Create().
		SetName(newProduct.Name).
		SetSlug(newProduct.Slug).
		SetMetaTitle(newProduct.MetaTitle).
		SetMetaDescription(newProduct.MetaDescription).
		AddCategoryIDs(categoryIds...).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedProduct, nil
}

func (r *ProductOps) ProductUpdate(product ent.Product, categoryIds []int) (*ent.Product, error) {

	clearProductCategoryModel, err := r.client.Product.UpdateOneID(product.ID).
		ClearCategories().
		Save(r.ctx)

	fmt.Printf("clearProductCategoryModel: %+v\n", clearProductCategoryModel)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := r.client.Product.UpdateOneID(product.ID).
		SetName(product.Name).
		SetSlug(product.Slug).
		SetMetaTitle(product.MetaTitle).
		SetMetaDescription(product.MetaDescription).
		AddCategoryIDs(categoryIds...).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (r *ProductOps) ProductDelete(id int) (int, error) {

	err := r.client.Product.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

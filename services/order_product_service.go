package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
)

type OrderProductOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewOrderProductOps(ctx context.Context) *OrderProductOps {
	return &OrderProductOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *OrderProductOps) OrderProductPaginate(limit int, offset int) ([]*ent.OrderProduct, error) {

	orderProduct, err := r.client.OrderProduct.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return orderProduct, nil
}

func (r *OrderProductOps) OrderProductProductsGetAll() ([]*ent.OrderProduct, error) {

	orderProductProducts, err := r.client.OrderProduct.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return orderProductProducts, nil
}

func (r *OrderProductOps) OrderProductGetByID(id int) (*ent.OrderProduct, error) {

	orderProductModel, err := r.client.OrderProduct.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return orderProductModel, nil
}

func (r *OrderProductOps) OrderProductCreate(newOrderProduct ent.OrderProduct) (*ent.OrderProduct, error) {

	newCreatedOrderProduct, err := r.client.OrderProduct.Create().
		SetOrderID(newOrderProduct.OrderID).
		SetProductID(newOrderProduct.ProductID).
		SetAmount(newOrderProduct.Amount).
		SetQty(newOrderProduct.Qty).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedOrderProduct, nil
}

func (r *OrderProductOps) OrderProductUpdate(orderProduct ent.OrderProduct) (*ent.OrderProduct, error) {

	updatedOrderProduct, err := r.client.OrderProduct.UpdateOneID(orderProduct.ID).
		SetOrderID(orderProduct.OrderID).
		SetProductID(orderProduct.ProductID).
		SetAmount(orderProduct.Amount).
		SetQty(orderProduct.Qty).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedOrderProduct, nil
}

func (r *OrderProductOps) OrderProductDelete(id int) (int, error) {

	err := r.client.OrderProduct.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

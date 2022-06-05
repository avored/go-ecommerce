package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
)

type OrderOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewOrderOps(ctx context.Context) *OrderOps {
	return &OrderOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *OrderOps) OrderPaginate(limit int, offset int) ([]*ent.Order, error) {

	order, err := r.client.Order.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderOps) OrdersGetAll() ([]*ent.Order, error) {

	orders, err := r.client.Order.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderOps) OrderGetByID(id int) (*ent.Order, error) {

	orderModel, err := r.client.Order.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return orderModel, nil
}

func (r *OrderOps) OrderCreate(newOrder ent.Order) (*ent.Order, error) {

	newCreatedOrder, err := r.client.Order.Create().
		SetPaymentMethod(newOrder.PaymentMethod).
		SetShippingMethod(newOrder.ShippingMethod).
		SetShippingAddressID(newOrder.ShippingAddressID).
		SetBillingAddressID(newOrder.BillingAddressID).
		SetCustomerID(newOrder.CustomerID).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedOrder, nil
}

func (r *OrderOps) OrderUpdate(order ent.Order) (*ent.Order, error) {

	updatedOrder, err := r.client.Order.UpdateOneID(order.ID).
		SetPaymentMethod(order.PaymentMethod).
		SetShippingMethod(order.ShippingMethod).
		SetShippingAddressID(order.ShippingAddressID).
		SetBillingAddressID(order.BillingAddressID).
		SetCustomerID(order.CustomerID).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedOrder, nil
}

func (r *OrderOps) OrderDelete(id int) (int, error) {

	err := r.client.Order.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

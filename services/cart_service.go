package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/cart"

	"github.com/avored/go-ecommerce/providers"
)

type CartOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCartOps(ctx context.Context) *CartOps {
	return &CartOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *CartOps) CartPaginate(limit int, offset int) ([]*ent.Cart, error) {

	cart, err := r.client.Cart.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartOps) AdminCartsGetAll() ([]*ent.Cart, error) {

	carts, err := r.client.Cart.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *CartOps) CartGetByID(id int) (*ent.Cart, error) {

	cartModel, err := r.client.Cart.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return cartModel, nil
}
func (r *CartOps) CartGetBySessionId(sessionId string) (*ent.Cart, error) {

	cartModel, err := r.client.Cart.Query().Where(cart.SessionIDEQ(sessionId)).Only(r.ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}

	return cartModel, nil
}

func (r *CartOps) CartCreate(newCart ent.Cart) (*ent.Cart, error) {

	newCreatedCart, err := r.client.Cart.Create().
		SetSessionID(newCart.SessionID).
		SetFullAmount(newCart.FullAmount).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedCart, nil
}

func (r *CartOps) CartDelete(id int) (int, error) {

	err := r.client.Cart.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

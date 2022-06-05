package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/cartitem"
	"github.com/avored/go-ecommerce/providers"
)

type CartItemOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCartItemOps(ctx context.Context) *CartItemOps {
	return &CartItemOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *CartItemOps) CartItemPaginate(limit int, offset int) ([]*ent.CartItem, error) {

	cartItem, err := r.client.CartItem.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (r *CartItemOps) AdminCartItemsGetAll() ([]*ent.CartItem, error) {

	cartItems, err := r.client.CartItem.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *CartItemOps) CartItemGetByID(id int) (*ent.CartItem, error) {

	cartItemModel, err := r.client.CartItem.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return cartItemModel, nil
}

func (r *CartItemOps) GetCartItemByCartAndProductModel(cartModel *ent.Cart, productModel *ent.Product) (*ent.CartItem, error) {

	cartItemModel, err := r.client.CartItem.Query().
		Where(cartitem.ProductIDEQ(productModel.ID)).
		Where(cartitem.CartIDEQ(cartModel.ID)).
		// WithCart(func(q *ent.CartQuery) {
		// 	q.Where(cart.IDEQ(cartModel.ID))
		// }).
		// WithProduct(func(pq *ent.ProductQuery) {
		// 	pq.Where(product.IDEQ(productModel.ID))
		// }).
		Only(r.ctx)

	if ent.IsNotFound(err) {
		return nil, err
	}

	return cartItemModel, nil
}

func (r *CartItemOps) CartItemCreate(newCartItem ent.CartItem, newCart *ent.Cart, product *ent.Product) (*ent.CartItem, error) {

	newCreatedCartItem, err := r.client.CartItem.Create().
		SetName(newCartItem.Name).
		SetQty(newCartItem.Qty).
		SetItemPrice(newCartItem.ItemPrice).
		SetCart(newCart).
		SetProduct(product).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedCartItem, nil
}

func (r *CartItemOps) CartItemUpdate(cartItem ent.CartItem, newCart *ent.Cart, product *ent.Product) (*ent.CartItem, error) {

	updatedCartItem, err := r.client.CartItem.UpdateOneID(cartItem.ID).
		SetName(cartItem.Name).
		SetQty(cartItem.Qty).
		SetItemPrice(cartItem.ItemPrice).
		SetCart(newCart).
		SetProduct(product).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}
	return updatedCartItem, nil
}

func (r *CartItemOps) CartItemDelete(id int) (int, error) {

	err := r.client.CartItem.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

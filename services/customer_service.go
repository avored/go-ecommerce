package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/ent/customer"
	"github.com/avored/go-ecommerce/providers"
)

type CustomerOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewCustomerOps(ctx context.Context) *CustomerOps {
	return &CustomerOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *CustomerOps) CustomerPaginate(limit int, offset int) ([]*ent.Customer, error) {

	customer, err := r.client.Customer.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerOps) CustomersGetAll() ([]*ent.Customer, error) {

	customers, err := r.client.Customer.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerOps) CustomerGetByID(id int) (*ent.Customer, error) {

	customerModel, err := r.client.Customer.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return customerModel, nil
}
func (r *CustomerOps) GetCustomerByEmail(email string) (*ent.Customer, error) {

	customerModel, err := r.client.Customer.Query().Where(customer.EmailEQ(email)).Only(r.ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}

	return customerModel, nil
}

func (r *CustomerOps) CustomerCreate(newCustomer ent.Customer) (*ent.Customer, error) {

	newCreatedCustomer, err := r.client.Customer.Create().
		SetFirstName(newCustomer.FirstName).
		SetLastName(newCustomer.LastName).
		SetEmail(newCustomer.Email).
		SetPassword(newCustomer.Password).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedCustomer, nil
}

func (r *CustomerOps) CustomerUpdate(customer ent.Customer) (*ent.Customer, error) {

	updatedCustomer, err := r.client.Customer.UpdateOneID(customer.ID).
		SetFirstName(customer.FirstName).
		SetLastName(customer.LastName).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (r *CustomerOps) CustomerDelete(id int) (int, error) {

	err := r.client.Customer.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

package services

import (
	"context"

	"github.com/avored/go-ecommerce/ent"
	"github.com/avored/go-ecommerce/providers"
)

type AddressOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewAddressOps(ctx context.Context) *AddressOps {
	return &AddressOps{
		ctx:    ctx,
		client: providers.GetDatabaseClient(),
	}
}
func (r *AddressOps) AddressPaginate(limit int, offset int) ([]*ent.Address, error) {

	address, err := r.client.Address.Query().Limit(limit).Offset(offset).All(r.ctx)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AddressOps) AddressesGetAll() ([]*ent.Address, error) {

	addresses, err := r.client.Address.Query().All(r.ctx)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *AddressOps) AddressGetByID(id int) (*ent.Address, error) {

	addressModel, err := r.client.Address.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return addressModel, nil
}

func (r *AddressOps) AddressCreate(newAddress ent.Address) (*ent.Address, error) {

	newCreatedAddress, err := r.client.Address.Create().
		SetType(newAddress.Type).
		SetIsDefault(newAddress.IsDefault).
		SetFirstName(newAddress.FirstName).
		SetLastName(newAddress.LastName).
		SetAddress1(newAddress.Address1).
		SetAddress2(newAddress.Address2).
		SetCompanyName(newAddress.CompanyName).
		SetCity(newAddress.City).
		SetSuburb(newAddress.Suburb).
		SetState(newAddress.State).
		SetPostcode(newAddress.Postcode).
		SetCountryID(newAddress.CountryID).
		SetPhone(newAddress.Phone).
		SetCustomerID(newAddress.CustomerID).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedAddress, nil
}

func (r *AddressOps) AddressUpdate(address ent.Address) (*ent.Address, error) {

	updatedAddress, err := r.client.Address.UpdateOneID(address.ID).
		SetType(address.Type).
		SetIsDefault(address.IsDefault).
		SetAddress1(address.Address1).
		SetAddress2(address.Address2).
		SetCompanyName(address.CompanyName).
		SetSuburb(address.Suburb).
		SetCity(address.City).
		SetState(address.State).
		SetPostcode(address.Postcode).
		SetCountryID(address.CountryID).
		SetCustomer(address.Edges.Customer).
		SetPhone(address.Phone).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedAddress, nil
}

func (r *AddressOps) AddressDelete(id int) (int, error) {

	err := r.client.Address.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

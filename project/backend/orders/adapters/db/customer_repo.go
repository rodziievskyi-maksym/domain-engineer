package db

import (
	"context"
	"eats/backend/common/shared"
	"eats/backend/orders/adapters/db/dbmodels"

	"github.com/jackc/pgx/v5/pgxpool"

	"eats/backend/common"
	"eats/backend/orders/api/http"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	if db == nil {
		panic("db connection pool cannot be nil")
	}

	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) RegisterCustomer(ctx context.Context, customerUUID common.UUID, customer http.RegisterCustomer) error {
	address, err := openapiAddressToSharedAddress(customer.Address)
	if err != nil {
		return err
	}

	args := dbmodels.InsertCustomerParams{
		CustomerUuid: customerUUID,
		Name:         customer.Name,
		Email:        string(customer.Email),
		Address:      address,
		PhoneNumber:  customer.PhoneNumber,
	}

	queries := dbmodels.New(r.db)

	if err := queries.InsertCustomer(ctx, args); err != nil {
		return err
	}

	return nil
}

func openapiAddressToSharedAddress(addr http.Address) (shared.Address, error) {
	return shared.NewAddress(
		addr.Line1,
		addr.Line2,
		addr.PostalCode,
		addr.City,
		addr.CountryCode,
	)
}

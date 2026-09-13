package db

import (
	"context"
	"eats/backend/orders/adapters/db/dbmodels"
	"eats/backend/orders/app"

	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *CustomerRepository) RegisterCustomer(ctx context.Context, customer app.Customer) error {
	args := dbmodels.InsertCustomerParams{
		CustomerUuid: customer.CustomerUUID,
		Name:         customer.Name,
		Email:        customer.Email,
		Address:      customer.Address,
		PhoneNumber:  customer.PhoneNumber,
	}

	queries := dbmodels.New(r.db)

	if err := queries.InsertCustomer(ctx, args); err != nil {
		return err
	}

	return nil
}

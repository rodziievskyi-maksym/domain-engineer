package http

import (
	"context"

	"eats/backend/common"
	"eats/backend/common/shared"
	"eats/backend/orders/app"
)

type Handler struct {
	app *app.Service
}

func NewHandler(app *app.Service) Handler {
	if app == nil {
		panic("order app cannot be nil")
	}

	return Handler{
		app: app,
	}
}

func Register(ctx context.Context, e common.EchoRouter, handler Handler) error {
	RegisterHandlers(e, NewStrictHandler(handler, nil))
	return nil
}

func (h Handler) RegisterCustomer(ctx context.Context, request RegisterCustomerRequestObject) (RegisterCustomerResponseObject, error) {
	customerUUID := app.CustomerUUID{common.NewUUIDv7()}

	commonAddress, err := openapiAddressToSharedAddress(request.Body.Address)
	if err != nil {
		return nil, err
	}

	customer := app.Customer{
		CustomerUUID: customerUUID,
		Name:         request.Body.Name,
		Email:        string(request.Body.Email),
		Address:      commonAddress,
		PhoneNumber:  request.Body.PhoneNumber,
	}

	if err := h.app.RegisterCustomer(context.Background(), customer); err != nil {
		return nil, err
	}

	return RegisterCustomer201JSONResponse{
		CustomerUuid: customerUUID,
	}, nil
}

func openapiAddressToSharedAddress(address Address) (shared.Address, error) {
	commonAddress, err := shared.NewAddress(
		address.Line1,
		address.Line2,
		address.PostalCode,
		address.City,
		address.CountryCode,
	)
	if err != nil {
		return shared.Address{}, err
	}

	return commonAddress, nil
}

package shared

import (
	"eats/backend/common"
	"fmt"
)

type Currency struct {
	common.Enum[CurrencyType]
}

func (c Currency) Code() string {
	return c.String()
}

type CurrencyType string

func (c CurrencyType) Values() []string {
	return []string{"USD", "EUR", "GBP", "JPY", "PLN"}
}

func MustNewCurrency(value string) Currency {
	c := Currency{}
	if err := c.UnmarshalText([]byte(value)); err != nil {
		panic(fmt.Errorf("error unmarshaling currency code value: %s", value))
	}

	return c
}

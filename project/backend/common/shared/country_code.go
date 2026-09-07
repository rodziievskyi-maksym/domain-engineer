package shared

import (
	"fmt"

	"eats/backend/common"
)

type CountryCode struct {
	common.Enum[CountryCodeType]
}

func (c CountryCode) Code() string {
	return c.String()
}

type CountryCodeType string

func (c CountryCodeType) Values() []string {
	return []string{
		"US",
		"DE",
		"GB",
		"JP",
		"PL",
	}
}

func MustNewCountryCode(value string) CountryCode {
	c := CountryCode{}
	err := c.UnmarshalText([]byte(value))
	if err != nil {
		panic(fmt.Errorf("error unmarshalling country code value: %s", value))
	}

	return c
}

type OrderStatus struct {
	common.Enum[OrderStatusType]
}

func (o OrderStatus) Status() string {
	return o.String()
}

type OrderStatusType string

func (o OrderStatusType) Values() []string {
	return []string{
		"CREATED",
		"CONFIRMED",
		"SHIPPED",
		"DELIVERED",
		"CANCELLED",
	}
}

func MustNewOrderStatus(value string) OrderStatus {
	o := OrderStatus{}
	err := o.UnmarshalText([]byte(value))
	if err != nil {
		panic(fmt.Errorf("error unmarshalling order status: %s", value))
	}

	return o
}

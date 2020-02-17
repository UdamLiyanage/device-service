package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (device Device) Validate() error {
	return validation.ValidateStruct(&device,
		validation.Field(&device.Serial, validation.Required, is.Alphanumeric),
		validation.Field(&device.Name, validation.Required, validation.Length(5, 25)),
	)
}

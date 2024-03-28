package validator_test

import (
	"marketplace/model"
	"marketplace/pkg/price"
	"marketplace/pkg/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateRegisterInput(t *testing.T) {
	tests := []struct {
		name          string
		registerInput model.RegisterInput
		expectedErr   error
	}{
		{
			"valid input",
			model.RegisterInput{
				Username: "username",
				Password: "qwerty123",
			},
			nil,
		},
		{
			"empty username",
			model.RegisterInput{
				Username: "",
				Password: "qwerty123",
			},
			validator.ErrInvalidUsername,
		},
		{
			"username with spaces",
			model.RegisterInput{
				Username: "one two",
				Password: "qwerty123",
			},
			validator.ErrInvalidUsername,
		},
		{
			"small password",
			model.RegisterInput{
				Username: "username",
				Password: "12345",
			},
			validator.ErrInvalidPassword,
		},
		{
			"small password",
			model.RegisterInput{
				Username: "username",
				Password: "12345",
			},
			validator.ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := validator.ValidateRegisterInput(tt.registerInput)

			assert.ErrorIs(t, actualErr, tt.expectedErr, "errors must match")
		})
	}
}

func TestValidateAdAPI(t *testing.T) {
	tests := []struct {
		name        string
		ad          model.AdAPI
		expectedErr error
	}{
		{
			"valid input",
			model.AdAPI{
				Name:        "name",
				Description: "description",
				Price:       "0.00",
			},
			nil,
		},
		{
			"empty name",
			model.AdAPI{
				Name:        "",
				Description: "description",
				Price:       "0.00",
			},
			validator.ErrInvalidAdName,
		},
		{
			"long name",
			model.AdAPI{
				Name:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Description: "description",
				Price:       "0.00",
			},
			validator.ErrInvalidAdName,
		},
		{
			"long description",
			model.AdAPI{
				Name:        "name",
				Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				Price:       "0.00",
			},
			validator.ErrInvalidAdDescription,
		},
		{
			"invalid price 1",
			model.AdAPI{
				Name:        "name",
				Description: "description",
				Price:       "price",
			},
			price.ErrInvalidPriceFormat,
		},
		{
			"invalid price 2",
			model.AdAPI{
				Name:        "name",
				Description: "description",
				Price:       "",
			},
			price.ErrInvalidPriceFormat,
		},
		{
			"invalid price 3",
			model.AdAPI{
				Name:        "name",
				Description: "description",
				Price:       "0.0,",
			},
			price.ErrInvalidPriceFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualErr := validator.ValidateAdAPI(tt.ad)

			assert.ErrorIs(t, actualErr, tt.expectedErr, "errors must match")
		})
	}
}

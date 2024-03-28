package price_test

import (
	"marketplace/pkg/price"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriceFromCents(t *testing.T) {
	tests := []struct {
		name        string
		cents       int64
		expected    string
		expectedErr error
	}{
		{"zero cents", 0, "0.00", nil},
		{"one cent", 1, "0.01", nil},
		{"two digits cent", 54, "0.54", nil},
		{"three digits cent", 324, "3.24", nil},
		{"four digits cent", 4324, "43.24", nil},
		{"five digits cent", 54324, "543.24", nil},
		{"with zero cent", 54300, "543.00", nil},
		{"with zero cent", 54305, "543.05", nil},
		{"with zero cent", 54350, "543.50", nil},
		{"invalid cents", -45, "", price.ErrInvalidCents},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := price.FromCents(tt.cents)

			assert.Equal(t, actual, tt.expected, "should be equal")
			assert.ErrorIs(t, actualErr, tt.expectedErr, "errors must match")
		})
	}
}

func TestPriceToCents(t *testing.T) {
	tests := []struct {
		name        string
		price       string
		expected    int64
		expectedErr error
	}{
		{"zero cents", "0.00", 0, nil},
		{"one cent", "0.01", 1, nil},
		{"two digits cent", "0.54", 54, nil},
		{"three digits cent", "3.24", 324, nil},
		{"four digits cent", "43.24", 4324, nil},
		{"five digits cent", "543.24", 54324, nil},
		{"with zero cent", "543.00", 54300, nil},
		{"with zero cent", "543.05", 54305, nil},
		{"with zero cent", "543.50", 54350, nil},
		{"head zeros", "00543.50", 54350, nil},
		{"empty price", "", 0, price.ErrInvalidPriceFormat},
		{"only cents", ".05", 0, price.ErrInvalidPriceFormat},
		{"without period", "546", 0, price.ErrInvalidPriceFormat},
		{"one symbol cents", "5.5", 0, price.ErrInvalidPriceFormat},
		{"multiple periods 1", "5.50.", 0, price.ErrInvalidPriceFormat},
		{"multiple periods 2", "5..50", 0, price.ErrInvalidPriceFormat},
		{"multiple periods 3", "5.50.15", 0, price.ErrInvalidPriceFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualErr := price.ToCents(tt.price)

			assert.Equal(t, actual, tt.expected, "should be equal")
			assert.ErrorIs(t, actualErr, tt.expectedErr, "errors must match")
		})
	}
}

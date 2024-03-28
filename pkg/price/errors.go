package price

import "errors"

var (
	ErrInvalidCents       = errors.New("invalid cents")
	ErrInvalidPriceFormat = errors.New("invalid price format")
)

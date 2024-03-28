package price

import (
	"fmt"
	"strconv"
	"strings"
)

func FromCents(cents int64) (string, error) {
	if cents < 0 {
		return "", ErrInvalidCents
	}
	var integer int64 = cents / 100
	var frac int64 = cents % 100
	return fmt.Sprintf("%d.%02d", integer, frac), nil
}

func ToCents(price string) (int64, error) {
	priceParts := strings.Split(price, ".")
	if len(priceParts) != 2 || len(priceParts[1]) != 2 || len(priceParts[0]) < 1 {
		return 0, ErrInvalidPriceFormat
	}

	integer, err := strconv.ParseInt(priceParts[0], 10, 64)
	if err != nil {
		return 0, ErrInvalidPriceFormat
	}

	fraq, err := strconv.ParseInt(priceParts[1], 10, 64)
	if err != nil {
		return 0, ErrInvalidPriceFormat
	}

	return integer*100 + fraq, nil
}

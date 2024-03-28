package validator

import "marketplace/pkg/price"

const (
	USERNAME_MAX_LEN = 150
	PASSOWRD_MIN_LEN = 6

	AD_NAME_MAX_LEN        = 200
	AD_DESCRIPTION_MAX_LEN = 1000
)

func ValidateUsername(username string) error {
	if len(username) == 0 || len(username) > USERNAME_MAX_LEN {
		return ErrInvalidUsername
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < PASSOWRD_MIN_LEN {
		return ErrInvalidPassword
	}
	return nil
}

func ValidateAdName(name string) error {
	if len(name) > AD_NAME_MAX_LEN {
		return ErrInvalidAdName
	}
	return nil
}

func ValidateAdDescription(description string) error {
	if len(description) > AD_DESCRIPTION_MAX_LEN {
		return ErrInvalidAdDescription
	}
	return nil
}

func ValidateAdPrice(priceStr string) error {
	_, err := price.ToCents(priceStr)
	if err != nil {
		return err
	}
	return nil
}

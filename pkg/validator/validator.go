package validator

const (
	USERNAME_MAX_LEN = 150
	PASSOWRD_MIN_LEN = 6
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

package validator

import "marketplace/model"

func ValidateRegisterInput(registerInput model.RegisterInput) (err error) {
	if err = ValidateUsername(registerInput.Username); err != nil {
		return
	}
	if err = ValidatePassword(registerInput.Password); err != nil {
		return
	}
	return
}

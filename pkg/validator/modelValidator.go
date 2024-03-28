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

func ValidateAdAPI(ad model.AdAPI) (err error) {
	if err = ValidateAdName(ad.Name); err != nil {
		return
	}
	if err = ValidateAdDescription(ad.Description); err != nil {
		return
	}
	if err = ValidateAdPrice(ad.Price); err != nil {
		return
	}
	return
}

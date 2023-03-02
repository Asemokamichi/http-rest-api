package model

import validation "github.com/go-ozzo/ozzo-validation"

func requiredIf(conde bool) validation.RuleFunc {
	return func(value interface{}) error {
		if conde {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

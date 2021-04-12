package validator

import (
	"fmt"
)

type StringValidator struct {
	rules     []StringRule
	required  requiredRule
	fieldName string
}

func NewStringValidator(name string, required bool, r ...StringRule) StringValidator {
	return StringValidator{
		fieldName: name,
		rules:     r,
		required:  StrRequired(required),
	}
}

func (sv StringValidator) Validate(s string) error {
	if err := sv.required.Validate(s); err != nil {
		return fmt.Errorf("%s is invalid: %w", sv.fieldName, err)
	}
	if isEmpty(s) {
		return nil
	}
	for _, r := range sv.rules {
		err := r.Validate(s)
		if err != nil {
			return fmt.Errorf("%s is invalid: %w", sv.fieldName, err)
		}
	}
	return nil
}

func isEmpty(s string) bool {
	return len(s) == 0
}

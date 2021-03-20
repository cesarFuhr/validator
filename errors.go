package validator

import "errors"

var (
	ErrTooShort        = errors.New("should not be shorter than")
	ErrTooLong         = errors.New("should not be longer than")
	ErrRequired        = errors.New("is required")
	ErrRegexNotMatched = errors.New("should satisfy the regex")
	ErrInvalidDate     = errors.New("should be a parseable date string")
	ErrInvalidUUID     = errors.New("should be a parseable uuid string")
)

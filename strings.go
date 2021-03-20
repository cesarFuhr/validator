package validator

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type StringRule interface {
	Validate(string) error
}

func StrLength(min, max int) lengthRule {
	return lengthRule{min, max}
}

type lengthRule struct {
	min int
	max int
}

func (r lengthRule) Validate(s string) error {
	if len(s) < r.min {
		return fmt.Errorf("%w %d", ErrTooShort, r.min)
	}
	if len(s) > r.max {
		return fmt.Errorf("%w %d", ErrTooLong, r.max)
	}

	return nil
}

func StrRequired(required bool) requiredRule {
	return requiredRule{required}
}

type requiredRule struct {
	required bool
}

func (r requiredRule) Validate(s string) error {
	if r.required && len(s) == 0 {
		return ErrRequired
	}

	return nil
}

func StrRegexp(rg *regexp.Regexp) regexRule {
	return regexRule{
		rg: rg,
	}
}

type regexRule struct {
	rg *regexp.Regexp
}

func (r regexRule) Validate(s string) error {
	if !r.rg.MatchString(s) {
		return fmt.Errorf("%w %v", ErrRegexNotMatched, r.rg)
	}

	return nil
}

func StrDate(tf string) dateRule {
	return dateRule{tf}
}

type dateRule struct {
	tf string
}

func (r dateRule) Validate(s string) error {
	if _, err := time.Parse(r.tf, s); err != nil {
		return fmt.Errorf("%w, %s", ErrInvalidDate, err.Error())
	}

	return nil
}

func StrUUID() uuidRule {
	return uuidRule{}
}

type uuidRule struct {
}

func (r uuidRule) Validate(s string) error {
	if _, err := uuid.Parse(s); err != nil {
		return fmt.Errorf("%w, %s", ErrInvalidUUID, err.Error())
	}

	return nil
}

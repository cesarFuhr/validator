package validator

import (
	"errors"
	"regexp"
	"testing"
	"time"
)

func TestLengthRule(t *testing.T) {
	type in struct {
		s        string
		min, max int
	}
	cases := []struct {
		d     string
		input in
		want  error
	}{
		{d: "returns an error if the string is shorter than min", input: in{"12", 3, 5}, want: ErrTooShort},
		{d: "returns an error if the string is longer than max", input: in{"121212", 3, 5}, want: ErrTooLong},
		{d: "returns nil if the string is inside the minimum length", input: in{"123", 3, 5}, want: nil},
		{d: "returns nil if the string is inside the maximum length", input: in{"12345", 3, 5}, want: nil},
	}
	for _, c := range cases {
		t.Run(c.d, func(t *testing.T) {
			rule := lengthRule{c.input.min, c.input.max}
			got := rule.Validate(c.input.s)

			if !errors.Is(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func BenchmarkLengthSuccess(b *testing.B) {
	rule := lengthRule{1, 2}
	s := "1"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}
func BenchmarkLengthError(b *testing.B) {
	rule := lengthRule{1, 2}
	s := "1123"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func TestRequiredRule(t *testing.T) {
	type in struct {
		s string
		r bool
	}
	cases := []struct {
		d     string
		input in
		want  error
	}{
		{d: "returns an error if the string is empty and required", input: in{"", true}, want: ErrRequired},
		{d: "returns nil if the string is empty but not required", input: in{"", false}, want: nil},
		{d: "returns nil if the string is not empty and required", input: in{"123", true}, want: nil},
		{d: "returns nil if the string is not empty and not required", input: in{"123", false}, want: nil},
	}

	for _, c := range cases {
		t.Run(c.d, func(t *testing.T) {
			rule := requiredRule{c.input.r}
			got := rule.Validate(c.input.s)

			if !errors.Is(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func BenchmarkRequiredSuccess(b *testing.B) {
	rule := requiredRule{}
	s := "1"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func BenchmarkRequiredError(b *testing.B) {
	rule := requiredRule{}
	s := ""
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func TestRegexRule(t *testing.T) {
	type in struct {
		s  string
		rg *regexp.Regexp
	}
	rgxTest, _ := regexp.Compile("123")
	cases := []struct {
		d     string
		input in
		want  error
	}{
		{d: "returns an error if the regexp did not match", input: in{"321", rgxTest}, want: ErrRegexNotMatched},
		{d: "returns nil if the regexp have matched", input: in{"123", rgxTest}, want: nil},
	}

	for _, c := range cases {
		t.Run(c.d, func(t *testing.T) {
			rule := regexRule{c.input.rg}
			got := rule.Validate(c.input.s)

			if !errors.Is(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func BenchmarkRegexpSuccess(b *testing.B) {
	rgx, _ := regexp.Compile("string")
	rule := regexRule{rgx}
	s := "string"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func BenchmarkRegexpError(b *testing.B) {
	rgx, _ := regexp.Compile("string")
	rule := regexRule{rgx}
	s := "123456"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func TestDateRule(t *testing.T) {
	type in struct {
		s  string
		tf string
	}
	cases := []struct {
		d     string
		input in
		want  error
	}{
		{d: "returns nil if the date is parseable", input: in{"2021-11-20T18:01:24.100+00:00", time.RFC3339}, want: nil},
		{d: "returns error if the date is not parseable", input: in{"123", time.RFC3339}, want: ErrInvalidDate},
	}

	for _, c := range cases {
		t.Run(c.d, func(t *testing.T) {
			rule := dateRule{c.input.tf}
			got := rule.Validate(c.input.s)

			if !errors.Is(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func BenchmarkDateRuleSuccess(b *testing.B) {
	rule := dateRule{time.RFC3339}
	s := "2021-11-20T18:01:24.100+00:00"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func BenchmarkDateRuleError(b *testing.B) {
	rule := dateRule{time.RFC3339}
	s := "123456"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func TestUUIDRule(t *testing.T) {
	cases := []struct {
		d     string
		input string
		want  error
	}{
		{d: "returns nil if the uuid is parseable", input: "4b9e7348-bdda-4584-88c1-a1e9ac4c6595", want: nil},
		{d: "returns error if the uuid is not parseable", input: "123", want: ErrInvalidUUID},
	}

	for _, c := range cases {
		t.Run(c.d, func(t *testing.T) {
			rule := uuidRule{}
			got := rule.Validate(c.input)

			if !errors.Is(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}

func BenchmarkUUIDRuleSuccess(b *testing.B) {
	rule := uuidRule{}
	s := "4b9e7348-bdda-4584-88c1-a1e9ac4c6595"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

func BenchmarkUUIDRuleError(b *testing.B) {
	rule := uuidRule{}
	s := "123456"
	for i := 0; i < b.N; i++ {
		rule.Validate(s)
	}
}

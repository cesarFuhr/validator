package validator

import (
	"errors"
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	type input struct {
		s  string
		sv StringValidator
	}
	testCases := []struct {
		desc string
		in   input
		want error
	}{
		{
			desc: "success: validate required and single length",
			in:   input{"12", NewStringValidator("fieldName", true, StrLength(1, 2))},
			want: nil,
		},
		{
			desc: "too long: validate required and single length",
			in:   input{"123", NewStringValidator("fieldName", true, StrLength(1, 2))},
			want: ErrTooLong,
		},
		{
			desc: "too short: validate required and single length",
			in:   input{"1", NewStringValidator("fieldName", true, StrLength(2, 2))},
			want: ErrTooShort,
		},
		{
			desc: "required: validate required",
			in:   input{"", NewStringValidator("fieldName", true, StrLength(2, 2))},
			want: ErrRequired,
		},
		{
			desc: "success: not required",
			in:   input{"", NewStringValidator("fieldName", false)},
			want: nil,
		},
		{
			desc: "success: not required with other rules",
			in:   input{"", NewStringValidator("fieldName", false, StrLength(1, 2))},
			want: nil,
		},
		{
			desc: "error: validate date string",
			in:   input{"2021-11-20T18:01", NewStringValidator("fieldName", false, StrDate(time.RFC3339))},
			want: ErrInvalidDate,
		},
		{
			desc: "Success: validate date string",
			in:   input{"2021-11-20T18:01:24.100+00:00", NewStringValidator("fieldName", false, StrDate(time.RFC3339))},
			want: nil,
		},
		{
			desc: "Success: validate uuid string and length",
			in:   input{"4b9e7348-bdda-4584-88c1-a1e9ac4c6595", NewStringValidator("fieldName", true, StrUUID(), StrLength(1, 36))},
			want: nil,
		},
		{
			desc: "Error: validate uuid string",
			in:   input{"12314", NewStringValidator("fieldName", true, StrUUID())},
			want: ErrInvalidUUID,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.in.sv.Validate(tC.in.s)

			if !errors.Is(got, tC.want) {
				t.Errorf("got %v, want %v", got, tC.want)
			}
		})
	}
}

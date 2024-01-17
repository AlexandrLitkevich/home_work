package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version        string `validate:"len:5"`
		RegisterNumber string `validate:"regexp:\\d+"`
		Location       string `validate:"in:msk,smr,spb"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
	InvalidRuleValidate struct {
		Name string `validate:"len:one"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		in          interface{}
		expectedErr error
	}{
		{
			name: "struct App",
			in: App{
				Version:        "ver12",
				RegisterNumber: "221100",
				Location:       "smr",
			},
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			_ = tt
			err := Validate(tt.in)
			require.NoError(t, err)
		})
	}
}

func TestValidateError(t *testing.T) {
	tests := []struct {
		name        string
		in          interface{}
		expectedErr error
	}{
		{
			name: "struct App InvalidLen",
			in: App{
				Version:        "ver",
				RegisterNumber: "221100",
				Location:       "smr",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   InvalidLength,
				},
			},
		},
		{
			name: "struct App InvalidLen|InvalidRegexp|InvalidValue",
			in: App{
				Version:        "ver",
				RegisterNumber: "notNumber",
				Location:       "smr1",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   InvalidLength,
				},
				ValidationError{
					Field: "RegisterNumber",
					Err:   InvalidRegexp,
				},
				ValidationError{
					Field: "Location",
					Err:   InvalidValue,
				},
			},
		},
		{
			name: "Invalid rule validate",
			in: InvalidRuleValidate{
				Name: "petr",
			},
			expectedErr: InvalidRule,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			_ = tt
			err := Validate(tt.in)
			require.Error(t, err)
			t.Log("this err", err)

			require.Equal(t, tt.expectedErr, err)

		})
	}
}

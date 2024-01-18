package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
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

	Num struct {
		Age  int `validate:"min:10|max:30"`
		Code int `validate:"in:200,404,500"`
		One  int `validate:"min:10"`
		Two  int `validate:"max:30"`
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

	SliceStruct struct {
		Name []string `validate:"len:6"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in interface{}
	}{
		{
			in: App{
				Version:        "ver12",
				RegisterNumber: "221100",
				Location:       "smr",
			},
		},
		{
			in: Num{
				Age:  22,
				Code: 200,
				One:  15,
				Two:  20,
			},
		},
		{
			in: Response{
				Code: 500,
			},
		},
		{
			in: SliceStruct{
				Name: []string{"123456", "123456"},
			},
		},
		{
			in: User{
				ID:     "caf6e4cb-5a38-47f8-92d9-7296ce85ae78",
				Name:   "Alex",
				Age:    22,
				Email:  "thismail@mail.com",
				Role:   "admin",
				Phones: []string{"12345678999", "12345678999", "12345678999", "12345678999"},
			},
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
					Err:   ErrInvalidLength,
				},
			},
		},
		{
			name: "struct App InvalidLen|ErrInvalidRegexp|ErrInvalidValue",
			in: App{
				Version:        "ver",
				RegisterNumber: "notNumber",
				Location:       "smr1",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   ErrInvalidLength,
				},
				ValidationError{
					Field: "RegisterNumber",
					Err:   ErrInvalidRegexp,
				},
				ValidationError{
					Field: "Location",
					Err:   ErrInvalidValue,
				},
			},
		},
		{
			name: "Invalid rule validate",
			in: InvalidRuleValidate{
				Name: "petr",
			},
			expectedErr: ErrInvalidRule,
		},
		{
			name: "struct Age InvalidMax|ErrInvalidValue|InvalidMin",
			in: Num{
				Age:  222,
				Code: 100,
				One:  2,
				Two:  20,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   ErrInvalidMaxValue,
				},
				ValidationError{
					Field: "Code",
					Err:   ErrInvalidValue,
				},
				ValidationError{
					Field: "One",
					Err:   ErrInvalidMinValue,
				},
			},
		},
		{
			name: "User struct ",
			in: User{
				ID:     "caf6e4cb-5a38-47f8",
				Name:   "Alex",
				Age:    17,
				Email:  "thisma",
				Role:   "employee",
				Phones: []string{"12", "12345678999", "12345678999", "12345678999"},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrInvalidLength,
				},
				ValidationError{
					Field: "Age",
					Err:   ErrInvalidMinValue,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrInvalidRegexp,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrInvalidValue,
				},
				ValidationError{
					Field: "Phones index:0",
					Err:   ErrInvalidLength,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			_ = tt
			err := Validate(tt.in)
			require.Error(t, err)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

package hw09structvalidator

import (
	"testing"
)

func TestValidateStringError(t *testing.T) {
	type args struct {
		field          FieldData
		validateErrors *ValidationErrors
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "errLen",
			args: args{
				field: FieldData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = ValidateString(tt.args.field, tt.args.validateErrors)
			t.Log(tt.args.validateErrors)
			//require.True(t, len(tt.args.validateErrors) > 0)
		})
	}
}

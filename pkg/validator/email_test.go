package validator

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsEmailValid_Validate(t *testing.T) {
	tests := []struct {
		email         string
		expectedError error
	}{
		// invalid
		{
			email:         "abc",
			expectedError: ErrInvalidEmail,
		},
		// valid
		{
			email:         "me@kaanksc.com",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ValidateEmail_%s", tt.email), func(t *testing.T) {
			vld := IsEmailValid{Email: tt.email}

			err := vld.Validate()

			if !errors.Is(err, tt.expectedError) {
				t.Error(err)
			}
		})
	}
}
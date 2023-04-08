package validator

import "fmt"

var _ Interface = (*IsEmailValid)(nil)

type IsEmailValid struct {
	Email string
}

func (i *IsEmailValid) Validate() error {
	if !emailRegexp.MatchString(i.Email) {
		return fmt.Errorf("is email valid : %w", ErrInvalidEmail)
	}

	return nil
}

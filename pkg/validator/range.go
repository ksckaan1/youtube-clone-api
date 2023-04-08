package validator

import "fmt"

var _ Interface = (*IsInRange)(nil)

type IsInRange struct {
	Text string
	Min  int
	Max  int
}

func (i *IsInRange) Validate() error {
	if !(len(i.Text) >= i.Min && len(i.Text) <= i.Max) {
		return fmt.Errorf("is in range : %w", ErrInvalidRange)
	}

	return nil
}

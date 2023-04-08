package validator

import (
	_ "regexp"
)

type Interface interface {
	Validate() error
}

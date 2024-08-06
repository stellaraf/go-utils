package errs

import "errors"

// IsOneOf determines if the first error matches the chain of any following errors.
func IsOneOf(err error, oneOf ...error) bool {
	for _, e := range oneOf {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

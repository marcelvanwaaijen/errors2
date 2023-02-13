package errors2

import (
	"errors"
)

// OnErr calls function fn when err is of (wrapped) type T and is not nil. If err is nil or of another type, fn will not be called.
func OnErr[T error](err error, fn func(err T)) {
	if retval := ErrAs[T](err); retval != nil {
		fn(*retval)
	}
}

// ErrAs returns the (wrapped) error as type T, or nil if param 'err' is of another type
func ErrAs[T error](err error) *T {
	switch retval := err.(type) {
	case T:
		return &retval
	default:
		werr := errors.Unwrap(err)
		if werr != nil {
			return ErrAs[T](werr)
		}
		return nil
	}
}

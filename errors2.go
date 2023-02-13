package errors2

import (
	"errors"
	"log"
)

// OnErr calls function fn when err is of (wrapped) type T and is not nil. If err is nil or of another type, fn will not be called.
func OnErr[T error](err error, fn func(err T)) {
	if retval := ErrAs[T](err); retval != nil {
		fn(*retval)
	}
}

// Log return a function that calls log.Fatalf using the signature that can be used in OnErr using the generic error interface as the type.
func Log(v ...any) func(err error) {
	return func(err error) {
		log.Print(v...)
	}
}

// Fatal return a function that calls log.Fatalf using the signature that can be used in OnErr using the generic error interface as the type.
func Fatal(v ...any) func(err error) {
	return func(err error) {
		log.Fatal(v...)
	}
}

// Logf return a function that calls log.Fatalf using the signature that can be used in OnErr using the generic error interface as the type.
func Logf(format string, v ...any) func(err error) {
	return func(err error) {
		log.Printf(format, v...)
	}
}

// Fatalf return a function that calls log.Fatalf using the signature that can be used in OnErr using the generic error interface as the type.
func Fatalf(format string, v ...any) func(err error) {
	return func(err error) {
		log.Fatalf(format, v...)
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

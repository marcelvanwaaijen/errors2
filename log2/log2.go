package log2

import "log"

// Log return a function that calls log.Fatalf using the signature that can be used in log2.OnErr using the generic error interface as the type.
func Log(v ...any) func(err error) {
	return func(err error) {
		log.Print(v...)
	}
}

// Fatal return a function that calls log.Fatalf using the signature that can be used in log2.OnErr using the generic error interface as the type.
func Fatal(v ...any) func(err error) {
	return func(err error) {
		log.Fatal(v...)
	}
}

// Logf return a function that calls log.Fatalf using the signature that can be used in log2.OnErr using the generic error interface as the type.
func Logf(format string, v ...any) func(err error) {
	return func(err error) {
		log.Printf(format, v...)
	}
}

// Fatalf return a function that calls log.Fatalf using the signature that can be used in log2.OnErr using the generic error interface as the type.
func Fatalf(format string, v ...any) func(err error) {
	return func(err error) {
		log.Fatalf(format, v...)
	}
}

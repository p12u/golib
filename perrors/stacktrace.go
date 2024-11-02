package perrors

import (
	"errors"
	"fmt"
	"runtime"
)

func withErrLocation(message string, location string) func(err error) error {
	return func(err error) error {
		if err == nil {
			return nil
		}

		return &withLocation{
			message:  message,
			err:      err,
			location: location,
		}
	}
}

// wraps an error with location
// this is needed so we can accurately report
// the locations where an error was created because
// we are wrapping fault, which does not allow configuring
// the trace attached to it.
type withLocation struct {
	message  string
	err      error
	location string
}

func (e *withLocation) Error() string {
	return "<ploc>"
}

func (e *withLocation) Unwrap() error {
	return e.err
}

// prints stacktrace based on the location
// if attached
func Stacktrace(err error) []string {
	result := []string{}
	if err == nil {
		return result
	}

	for err != nil {
		if locationErr, ok := err.(*withLocation); ok {
			result = append(
				result,
				fmt.Sprintf("%s - %s", locationErr.message, locationErr.location),
			)
		}
		err = errors.Unwrap(err)
	}

	return result
}

func getLocation(skipFrame int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skipFrame, pc)
	cf := runtime.CallersFrames(pc)
	f, _ := cf.Next()

	return fmt.Sprintf("%s:%d", f.File, f.Line)
}

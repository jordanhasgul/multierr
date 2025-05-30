// Package multierr is a Go package for aggregating multiple errors into
// a single error.
package multierr

import (
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"

	"github.com/jordanhasgul/multierr/prefix"
)

// Error is an error type that is used to aggregate multiple errors into
// a single error.
type Error struct {
	errs []error
}

// New returns a new Error that contains errs. Any nil errors contained
// within errs are removed.
func New(errs ...error) *Error {
	errs = removeNilErrors(errs)
	return &Error{errs: errs}
}

// Append returns an Error by appending errs onto err.
//
// If err is not an Error then it will be turned into one. Any nil errors
// contained within errs are removed.
func Append(err error, errs ...error) *Error {
	errs = removeNilErrors(errs)
	switch err := err.(type) {
	case *Error:
		if err == nil {
			return New(errs...)
		}

		err.errs = append(err.errs, errs...)
		return err
	default:
		if err == nil {
			return New(errs...)
		}

		errs = append(errs, nil)
		copy(errs[1:], errs)
		errs[0] = err
		return New(errs...)
	}
}

func removeNilErrors(errs []error) []error {
	isErrNil := func(err error) bool {
		if err == nil {
			return true
		}

		var (
			value = reflect.ValueOf(err)
			kind  = value.Kind()

			nillable = kind == reflect.Ptr || kind == reflect.UnsafePointer ||
				kind == reflect.Func || kind == reflect.Map || kind == reflect.Slice ||
				kind == reflect.Chan || kind == reflect.Interface
		)
		return nillable && value.IsNil()
	}
	return slices.DeleteFunc(errs, isErrNil)
}

// Error returns the string representation of an Error.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if len(e.errs) == 0 {
		return ""
	}

	var sb strings.Builder
	fprintError(&sb, &sb, e)
	return sb.String()
}

func fprintError(currWriter, prevWriter io.Writer, e *Error) {
	_, _ = fmt.Fprintf(prevWriter, "%d error(s) occurred:\n", len(e.errs))

	for i, err := range e.errs {
		var (
			pipe = "├── "
			sep  = "│   "
		)
		if i == len(e.errs)-1 {
			pipe = "└── "
			sep = "    "
		}

		switch err := err.(type) {
		case *Error:
			var (
				prevWriter = prefix.New(currWriter, pipe)
				currWriter = prefix.New(currWriter, sep)
			)
			fprintError(currWriter, prevWriter, err)
		default:
			_, _ = fmt.Fprintf(currWriter, "%s%s\n", pipe, err)
		}
	}
}

// Unwrap returns the list of errors that this Error wraps.
func (e *Error) Unwrap() []error {
	if e == nil {
		return nil
	}

	return e.errs
}

// Len returns the number of errors that this Error wraps.
func (e *Error) Len() int {
	if e == nil {
		return 0
	}

	return len(e.errs)
}

package multierr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_New(t *testing.T) {
	t.Run("zero-errors", func(t *testing.T) {
		e := New()
		require.Len(t, e.errs, 0)
	})

	t.Run("one-non-nil-error", func(t *testing.T) {
		e := New(errors.New("1"))
		require.Len(t, e.errs, 1)
	})

	t.Run("one-nil-error", func(t *testing.T) {
		e := New(nil)
		require.Len(t, e.errs, 0)
	})

	t.Run("some-non-nil-errors", func(t *testing.T) {
		e := New(
			errors.New("1"),
			errors.New("2"),
			errors.New("3"),
		)
		require.Len(t, e.errs, 3)
	})

	t.Run("some-nil-errors", func(t *testing.T) {
		e := New(
			errors.New("1"),
			errors.New("2"),
			nil,
		)
		require.Len(t, e.errs, 2)
	})
}

func TestError_Append(t *testing.T) {
	t.Run("non-nil-error", func(t *testing.T) {
		t.Run("zero-errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = Append(err)
			)
			require.Len(t, e.errs, 1)
		})

		t.Run("one-non-nil-error", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = Append(err, errors.New("2"))
			)
			require.Len(t, e.errs, 2)
		})

		t.Run("one-nil-error", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = Append(err, nil)
			)
			require.Len(t, e.errs, 1)
		})

		t.Run("some-non-nil-errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = Append(
					err,
					errors.New("2"),
					errors.New("3"),
					errors.New("4"),
				)
			)
			require.Len(t, e.errs, 4)
		})

		t.Run("some-nil-errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = Append(
					err,
					errors.New("2"),
					errors.New("3"),
					nil,
				)
			)
			require.Len(t, e.errs, 3)
		})
	})

	t.Run("nil-error", func(t *testing.T) {
		t.Run("zero-errors", func(t *testing.T) {
			var (
				err error
				e   = Append(err)
			)
			require.Len(t, e.errs, 0)
		})

		t.Run("one-non-nil-error", func(t *testing.T) {
			var (
				err error
				e   = Append(err, errors.New("1"))
			)
			require.Len(t, e.errs, 1)
		})

		t.Run("one-nil-error", func(t *testing.T) {
			var (
				err error
				e   = Append(err, nil)
			)
			require.Len(t, e.errs, 0)
		})

		t.Run("some-non-nil-errors", func(t *testing.T) {
			var (
				err error
				e   = Append(
					err,
					errors.New("1"),
					errors.New("2"),
					errors.New("3"),
				)
			)
			require.Len(t, e.errs, 3)
		})

		t.Run("some-nil-errors", func(t *testing.T) {
			var (
				err error
				e   = Append(
					err,
					errors.New("1"),
					errors.New("2"),
					nil,
				)
			)
			require.Len(t, e.errs, 2)
		})
	})

	t.Run("non-nil-multierror", func(t *testing.T) {
		t.Run("zero-errors", func(t *testing.T) {
			e := New(errors.New("1"))
			e = Append(e)
			require.Len(t, e.errs, 1)
		})

		t.Run("one-non-nil-error", func(t *testing.T) {
			e := New(errors.New("1"))
			e = Append(e, errors.New("2"))
			require.Len(t, e.errs, 2)
		})

		t.Run("one-nil-error", func(t *testing.T) {
			e := New(errors.New("1"))
			e = Append(e, nil)
			require.Len(t, e.errs, 1)
		})

		t.Run("some-non-nil-errors", func(t *testing.T) {
			e := New(errors.New("1"))
			e = Append(
				e,
				errors.New("2"),
				errors.New("3"),
				errors.New("4"),
			)
			require.Len(t, e.errs, 4)
		})

		t.Run("some-nil-errors", func(t *testing.T) {
			e := New(errors.New("1"))
			e = Append(
				e,
				errors.New("2"),
				errors.New("3"),
				nil,
			)
			require.Len(t, e.errs, 3)
		})
	})

	t.Run("nil-multierror", func(t *testing.T) {
		t.Run("zero-errors", func(t *testing.T) {
			var e *Error
			e = Append(e)
			require.Len(t, e.errs, 0)
		})

		t.Run("one-non-nil-error", func(t *testing.T) {
			var e *Error
			e = Append(e, errors.New("1"))
			require.Len(t, e.errs, 1)
		})

		t.Run("one-nil-error", func(t *testing.T) {
			var e *Error
			e = Append(e, nil)
			require.Len(t, e.errs, 0)
		})

		t.Run("some-non-nil-errors", func(t *testing.T) {
			var e *Error
			e = Append(
				e,
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			)
			require.Len(t, e.errs, 3)
		})

		t.Run("some-nil-errors", func(t *testing.T) {
			var e *Error
			e = Append(
				e,
				errors.New("1"),
				errors.New("2"),
				nil,
			)
			require.Len(t, e.errs, 2)
		})
	})
}

func TestError_Error(t *testing.T) {
	t.Run("zero-errors", func(t *testing.T) {
		var (
			e      = New()
			errStr = ""
		)

		require.Equal(t, errStr, e.Error())
	})

	t.Run("one-error", func(t *testing.T) {
		var (
			e      = New(errors.New("1"))
			errStr = "multierr: 1 error(s) occurred:\n└── 1\n"
		)
		require.Equal(t, errStr, e.Error())
	})

	t.Run("some-errors", func(t *testing.T) {
		var (
			e = New(
				errors.New("1"),
				New(
					errors.New("2"),
					errors.New("3"),
				),
				New(errors.New("4")),
			)
			errStr = "" +
				"multierr: 3 error(s) occurred:\n" +
				"├── 1\n" +
				"├── multierr: 2 error(s) occurred:\n" +
				"│   ├── 2\n" +
				"│   └── 3\n" +
				"└── multierr: 1 error(s) occurred:\n" +
				"    └── 4\n"
		)
		require.Equal(t, errStr, e.Error())
	})
}

func TestError_Unwrap(t *testing.T) {
	t.Run("empty-multierr", func(t *testing.T) {
		var (
			e    = New()
			errs = e.Unwrap()
		)
		require.Empty(t, errs)
	})

	t.Run("non-empty-multierr", func(t *testing.T) {
		var (
			e = New(
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			)
			errs = e.Unwrap()
		)
		require.Len(t, errs, 3)
	})
}

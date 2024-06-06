package multierr_test

import (
	"errors"
	"testing"

	"github.com/jordanhasgul/multierr"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("zero errors", func(t *testing.T) {
		e := multierr.New()
		require.Equal(t, e.Len(), 0)
	})

	t.Run("one non nil error", func(t *testing.T) {
		e := multierr.New(errors.New("1"))
		require.Equal(t, e.Len(), 1)
	})

	t.Run("one nil error", func(t *testing.T) {
		e := multierr.New(nil)
		require.Equal(t, e.Len(), 0)
	})

	t.Run("some non nil errors", func(t *testing.T) {
		e := multierr.New(
			errors.New("1"),
			errors.New("2"),
			errors.New("3"),
		)
		require.Equal(t, e.Len(), 3)
	})

	t.Run("some nil errors", func(t *testing.T) {
		e := multierr.New(
			errors.New("1"),
			errors.New("2"),
			nil,
		)
		require.Equal(t, e.Len(), 2)
	})
}

func TestAppend(t *testing.T) {
	t.Run("non nil error", func(t *testing.T) {
		t.Run("zero errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(err)
			)
			require.Equal(t, e.Len(), 1)
		})

		t.Run("one non nil error", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(err, errors.New("2"))
			)
			require.Equal(t, e.Len(), 2)
		})

		t.Run("one nil error", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(err, nil)
			)
			require.Equal(t, e.Len(), 1)
		})

		t.Run("some non nil errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(
					err,
					errors.New("2"),
					errors.New("3"),
					errors.New("4"),
				)
			)
			require.Equal(t, e.Len(), 4)
		})

		t.Run("some nil errors", func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(
					err,
					errors.New("2"),
					errors.New("3"),
					nil,
				)
			)
			require.Equal(t, e.Len(), 3)
		})
	})

	t.Run("nil error", func(t *testing.T) {
		t.Run("zero errors", func(t *testing.T) {
			var (
				err error
				e   = multierr.Append(err)
			)
			require.Equal(t, e.Len(), 0)
		})

		t.Run("one non nil error", func(t *testing.T) {
			var (
				err error
				e   = multierr.Append(err, errors.New("1"))
			)
			require.Equal(t, e.Len(), 1)
		})

		t.Run("one nil error", func(t *testing.T) {
			var (
				err error
				e   = multierr.Append(err, nil)
			)
			require.Equal(t, e.Len(), 0)
		})

		t.Run("some non nil errors", func(t *testing.T) {
			var (
				err error
				e   = multierr.Append(
					err,
					errors.New("1"),
					errors.New("2"),
					errors.New("3"),
				)
			)
			require.Equal(t, e.Len(), 3)
		})

		t.Run("some nil errors", func(t *testing.T) {
			var (
				err error
				e   = multierr.Append(
					err,
					errors.New("1"),
					errors.New("2"),
					nil,
				)
			)
			require.Equal(t, e.Len(), 2)
		})
	})

	t.Run("non nil multierror", func(t *testing.T) {
		t.Run("zero errors", func(t *testing.T) {
			e := multierr.New(errors.New("1"))
			e = multierr.Append(e)
			require.Equal(t, e.Len(), 1)
		})

		t.Run("one non nil error", func(t *testing.T) {
			e := multierr.New(errors.New("1"))
			e = multierr.Append(e, errors.New("2"))
			require.Equal(t, e.Len(), 2)
		})

		t.Run("one nil error", func(t *testing.T) {
			e := multierr.New(errors.New("1"))
			e = multierr.Append(e, nil)
			require.Equal(t, e.Len(), 1)
		})

		t.Run("some non nil errors", func(t *testing.T) {
			e := multierr.New(errors.New("1"))
			e = multierr.Append(
				e,
				errors.New("2"),
				errors.New("3"),
				errors.New("4"),
			)
			require.Equal(t, e.Len(), 4)
		})

		t.Run("some nil errors", func(t *testing.T) {
			e := multierr.New(errors.New("1"))
			e = multierr.Append(
				e,
				errors.New("2"),
				errors.New("3"),
				nil,
			)
			require.Equal(t, e.Len(), 3)
		})
	})

	t.Run("nil multierror", func(t *testing.T) {
		t.Run("zero errors", func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(e)
			require.Equal(t, e.Len(), 0)
		})

		t.Run("one non nil error", func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(e, errors.New("1"))
			require.Equal(t, e.Len(), 1)
		})

		t.Run("one nil error", func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(e, nil)
			require.Equal(t, e.Len(), 0)
		})

		t.Run("some non nil errors", func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(
				e,
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			)
			require.Equal(t, e.Len(), 3)
		})

		t.Run("some nil errors", func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(
				e,
				errors.New("1"),
				errors.New("2"),
				nil,
			)
			require.Equal(t, e.Len(), 2)
		})
	})
}

func TestError_Error(t *testing.T) {
	t.Run("zero errors", func(t *testing.T) {
		var (
			e      = multierr.New()
			errStr = ""
		)

		require.Equal(t, errStr, e.Error())
	})

	t.Run("one error", func(t *testing.T) {
		var (
			e      = multierr.New(errors.New("1"))
			errStr = "1 error(s) occurred:\n└── 1\n"
		)
		require.Equal(t, errStr, e.Error())
	})

	t.Run("some errors", func(t *testing.T) {
		var (
			e = multierr.New(
				errors.New("1"),
				multierr.New(
					errors.New("2"),
					errors.New("3"),
				),
				multierr.New(errors.New("4")),
			)
			errStr = "" +
				"3 error(s) occurred:\n" +
				"├── 1\n" +
				"├── 2 error(s) occurred:\n" +
				"│   ├── 2\n" +
				"│   └── 3\n" +
				"└── 1 error(s) occurred:\n" +
				"    └── 4\n"
		)
		require.Equal(t, errStr, e.Error())
	})
}

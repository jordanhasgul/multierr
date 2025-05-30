package multierr_test

import (
	"errors"
	"testing"

	"github.com/jordanhasgul/multierr"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output int
	}{
		{
			name: "with zero errors",

			input:  []error{},
			output: 0,
		},
		{
			name: "with one nil error",

			input:  []error{nil},
			output: 0,
		},
		{
			name: "with one non-nil error",

			input: []error{
				errors.New("1"),
			},
			output: 1,
		},
		{
			name: "with some nil errors",

			input: []error{
				errors.New("1"),
				errors.New("2"),
				nil,
			},
			output: 2,
		},
		{
			name: "with some non-nil errors",

			input: []error{
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			},
			output: 3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := multierr.New(testCase.input...)
			require.Equal(t, testCase.output, e.Len())
		})
	}
}

func TestAppend_AppendToNilError(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output int
	}{
		{
			name: "append zero errors",

			input:  []error{},
			output: 0,
		},
		{
			name: "append one nil error",

			input:  []error{nil},
			output: 0,
		},
		{
			name: "append one non-nil error",

			input: []error{
				errors.New("1"),
			},
			output: 1,
		},
		{
			name: "append some nil errors",

			input: []error{
				errors.New("1"),
				errors.New("2"),
				nil,
			},
			output: 2,
		},
		{
			name: "append some non-nil errors",
			input: []error{
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			},
			output: 3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := multierr.Append(nil, testCase.input...)
			require.Equal(t, testCase.output, e.Len())
		})
	}
}

func TestAppend_AppendToNonNilError(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output int
	}{
		{
			name: "append zero errors",

			input:  []error{},
			output: 1,
		},
		{
			name: "append one nil error",

			input:  []error{nil},
			output: 1,
		},
		{
			name: "append one non-nil error",

			input: []error{
				errors.New("2"),
			},
			output: 2,
		},
		{
			name: "append some nil errors",

			input: []error{
				errors.New("2"),
				errors.New("3"),
				nil,
			},
			output: 3,
		},
		{
			name: "append some non-nil errors",
			input: []error{
				errors.New("2"),
				errors.New("3"),
				errors.New("4"),
			},
			output: 4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				err = errors.New("1")
				e   = multierr.Append(err, testCase.input...)
			)
			require.Equal(t, testCase.output, e.Len())
		})
	}
}

func TestAppend_AppendToNilMultierr(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output int
	}{
		{
			name: "append zero errors",

			input:  []error{},
			output: 0,
		},
		{
			name: "append one nil error",

			input:  []error{nil},
			output: 0,
		},
		{
			name: "append one non-nil error",

			input: []error{
				errors.New("1"),
			},
			output: 1,
		},
		{
			name: "append some nil errors",

			input: []error{
				errors.New("1"),
				errors.New("2"),
				nil,
			},
			output: 2,
		},
		{
			name: "append some non-nil errors",
			input: []error{
				errors.New("1"),
				errors.New("2"),
				errors.New("3"),
			},
			output: 3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var e *multierr.Error
			e = multierr.Append(e, testCase.input...)
			require.Equal(t, testCase.output, e.Len())
		})
	}
}

func TestAppend_AppendToNonNilMultierr(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output int
	}{
		{
			name: "append zero errors",

			input:  []error{},
			output: 1,
		},
		{
			name: "append one nil error",

			input:  []error{nil},
			output: 1,
		},
		{
			name: "append one non-nil error",

			input: []error{
				errors.New("2"),
			},
			output: 2,
		},
		{
			name: "append some nil errors",

			input: []error{
				errors.New("2"),
				errors.New("3"),
				nil,
			},
			output: 3,
		},
		{
			name: "append some non-nil errors",
			input: []error{
				errors.New("2"),
				errors.New("3"),
				errors.New("4"),
			},
			output: 4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := multierr.New(
				errors.New("1"),
			)
			e = multierr.Append(e, testCase.input...)
			require.Equal(t, testCase.output, e.Len())
		})
	}
}

func TestError_ErrorOnNilMultierr(t *testing.T) {
	var e multierr.Error
	require.Empty(t, e.Error())
}

func TestError_ErrorOnNonNilMultierr(t *testing.T) {
	testCases := []struct {
		name string

		input  []error
		output string
	}{
		{
			name: "append zero errors",

			input:  []error{},
			output: "",
		},
		{
			name: "append one error",

			input: []error{
				errors.New("1"),
			},
			output: "" +
				"1 error(s) occurred:\n" +
				"└── 1\n",
		},
		{
			name: "append some errors",

			input: []error{
				errors.New("1"),
				multierr.New(
					errors.New("2"),
					errors.New("3"),
				),
				multierr.New(errors.New("4")),
			},
			output: "" +
				"3 error(s) occurred:\n" +
				"├── 1\n" +
				"├── 2 error(s) occurred:\n" +
				"│   ├── 2\n" +
				"│   └── 3\n" +
				"└── 1 error(s) occurred:\n" +
				"    └── 4\n",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := multierr.New(testCase.input...)
			require.Equal(t, testCase.output, e.Error())
		})
	}
}

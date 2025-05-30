package prefix_test

import (
	"fmt"
	"github.com/jordanhasgul/multierr/prefix"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	testCases := []struct {
		name string

		input  []string
		output string
	}{
		{
			name: "write nothing",

			input:  []string{},
			output: "",
		},
		{
			name: "write something",

			input: []string{"item 1", "item 2", "item 3"},
			output: "" +
				"* item 1\n" +
				"* item 2\n" +
				"* item 3\n",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				sb strings.Builder
				w  = prefix.New(&sb, "* ")
			)
			for _, input := range testCase.input {
				input = fmt.Sprintf("%s\n", input)
				_, _ = w.Write([]byte(input))
			}
			require.Equal(t, testCase.output, sb.String())
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	b.ReportAllocs()

	items := make([]string, 0, 100)
	for idx := range len(items) {
		item := fmt.Sprintf("item %d\n", idx)
		items = append(items, item)
	}

	var (
		sb strings.Builder
		w  = prefix.New(&sb, "* ")
	)
	for n := 0; n < b.N; n++ {
		for _, item := range items {
			_, _ = w.Write([]byte(item))
		}
	}
}

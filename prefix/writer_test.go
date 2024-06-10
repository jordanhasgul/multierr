package prefix_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"unsafe"

	"github.com/jordanhasgul/multierr/prefix"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	t.Run("write nothing", func(t *testing.T) {
		var (
			want = ""

			sb strings.Builder
			w  = prefix.New(&sb, "* ")
		)
		_, _ = w.Write([]byte(""))
		got := sb.String()

		require.Equal(t, want, got)
	})

	t.Run("write something", func(t *testing.T) {
		var (
			want = "" +
				"* item 1\n" +
				"* item 2\n" +
				"* item 3"

			sb strings.Builder
			w  = prefix.New(&sb, "* ")
		)
		_, _ = w.Write([]byte("item 1\n"))
		_, _ = w.Write([]byte("item 2\n"))
		_, _ = w.Write([]byte("item 3"))
		got := sb.String()

		require.Equal(t, want, got)
	})
}

func BenchmarkWrite(b *testing.B) {
	b.ReportAllocs()

	file, _ := os.CreateTemp("", "write_benchmark")
	defer os.Remove(file.Name())

	items := make([]string, 0, 100)
	for idx := range len(items) {
		item := fmt.Sprintf("item %d\n", idx)
		items = append(items, item)
	}

	writer := prefix.New(file, "* ")
	for n := 0; n < b.N; n++ {
		for _, item := range items {
			itemBytes := unsafe.Slice(
				unsafe.StringData(item),
				len(item),
			)
			_, _ = writer.Write(itemBytes)
		}
	}
}

package prefix

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteSomething(t *testing.T) {
	t.Run("write-nothing", func(t *testing.T) {
		var (
			want = ""

			sb strings.Builder
			w  = New(&sb, "* ")
		)
		_, _ = w.Write([]byte(""))
		got := sb.String()

		require.Equal(t, want, got)
	})

	t.Run("write-something", func(t *testing.T) {
		var (
			want = "" +
				"* item 1\n" +
				"* item 2\n" +
				"* item 3"

			sb strings.Builder
			w  = New(&sb, "* ")
		)
		_, _ = w.Write([]byte("item 1\n"))
		_, _ = w.Write([]byte("item 2\n"))
		_, _ = w.Write([]byte("item 3"))
		got := sb.String()

		require.Equal(t, want, got)
	})
}

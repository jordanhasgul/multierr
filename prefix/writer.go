package prefix

import (
	"bytes"
	"io"
)

var _ io.Writer = (*Writer)(nil)

type Writer struct {
	delegate io.Writer

	buf bytes.Buffer

	prefix      string
	seenNewline bool
}

func New(w io.Writer, prefix string) *Writer {
	return &Writer{
		delegate: w,

		prefix:      prefix,
		seenNewline: true,
	}
}

func (w *Writer) Write(bs []byte) (int, error) {
	for _, b := range bs {
		if w.seenNewline {
			w.buf.WriteString(w.prefix)
			w.seenNewline = false
		}

		w.buf.WriteByte(b)
		if b == '\n' {
			w.seenNewline = true
		}
	}

	_, err := w.buf.WriteTo(w.delegate)
	return len(bs), err
}

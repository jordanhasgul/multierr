package prefix

import (
	"bytes"
	"io"
	"unsafe"
)

var _ io.Writer = (*Writer)(nil)

type Writer struct {
	writer io.Writer

	buffer bytes.Buffer

	prefix       []byte
	insertPrefix bool
}

func New(w io.Writer, prefix string) *Writer {
	return &Writer{
		writer: w,

		prefix: unsafe.Slice(
			unsafe.StringData(prefix),
			len(prefix),
		),
		insertPrefix: true,
	}
}

func (w *Writer) Write(bs []byte) (int, error) {
	for _, b := range bs {
		if w.insertPrefix {
			_, _ = w.buffer.Write(w.prefix)
			w.insertPrefix = false
		}

		_ = w.buffer.WriteByte(b)
		w.insertPrefix = b == '\n'
	}

	_, err := w.buffer.WriteTo(w.writer)
	return len(bs), err
}

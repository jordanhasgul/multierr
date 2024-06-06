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

type Option func(*Writer)

func New(w io.Writer, opts ...Option) *Writer {
	writer := &Writer{
		writer: w,
	}
	for _, opt := range opts {
		opt(writer)
	}

	return writer
}

func WithPrefix(prefix string) Option {
	return func(writer *Writer) {
		writer.prefix = unsafe.Slice(
			unsafe.StringData(prefix),
			len(prefix),
		)
		writer.insertPrefix = true
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

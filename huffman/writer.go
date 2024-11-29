package huffman

import (
	"bufio"
	"io"
)

type writer struct {
	w *bufio.Writer
	b byte
	c byte
}

func newWriter(w io.Writer) *writer {
	return &writer{
		b: 7,
		c: 0,
		w: bufio.NewWriter(w),
	}
}

func (w *writer) writeByte(b byte) error {
	return w.w.WriteByte(b)
}

func (w *writer) writeBit(b byte) error {
	if b == 1 {
		w.c |= (1 << w.b)
	}
	if w.b == 0 {
		w.writeByte(w.c)
		w.b = 7
		w.c = 0
	} else {
		w.b--
	}
	return nil // REVIEW: Should we return something here?
}

func (w *writer) writeString(s string) error {
	_, err := w.w.WriteString(s)
	return err
}

func (w *writer) flush() error {
	return w.w.Flush()
}

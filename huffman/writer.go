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

func (w *writer) write(p []byte) (int, error) {
	return w.w.Write(p)
}

func (w *writer) writeByte(b byte) error {
	return w.w.WriteByte(b)
}

func (w *writer) writeBit(bit bool) error {
	if bit {
		w.c |= (1 << w.b)
	}
	if w.b == 0 {
		if err := w.writeByte(w.c); err != nil {
			return err
		}
		w.b = 7
		w.c = 0
	} else {
		w.b--
	}
	return nil
}

func (w *writer) writeString(s string) (int, error) {
	return w.w.WriteString(s)
}

func (w *writer) flush() error {
	return w.w.Flush()
}

package huffman

import (
	"bufio"
	"io"
)

type writer struct {
	w      *bufio.Writer
	bitIdx byte
	ch     byte
}

func newWriter(w io.Writer) *writer {
	return &writer{
		bitIdx: 7,
		ch:     0,
		w:      bufio.NewWriter(w),
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
		w.ch |= (1 << w.bitIdx)
	}
	if w.bitIdx == 0 {
		if err := w.writeByte(w.ch); err != nil {
			return err
		}
		w.bitIdx = 7
		w.ch = 0
	} else {
		w.bitIdx--
	}
	return nil
}

func (w *writer) writeString(s string) (int, error) {
	return w.w.WriteString(s)
}

func (w *writer) flush() error {
	return w.w.Flush()
}

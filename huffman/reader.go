package huffman

import (
	"bufio"
	"io"
)

type reader struct {
	r      *bufio.Reader
	bitIdx byte
	ch     byte
}

func newReader(r io.Reader) *reader {

	return &reader{
		r:      bufio.NewReader(r),
		bitIdx: 7,
		ch:     0,
	}
}

func (r *reader) read(p []byte) (int, error) {
	return r.r.Read(p)
}

func (r *reader) peekByte() (byte, error) {
	data, err := r.r.Peek(1)
	if err != nil {
		return 0, err
	}
	return data[0], nil
}

func (r *reader) nextByte() (byte, error) {
	return r.r.ReadByte()
}

func (r *reader) nextBit() (byte, error) {
	ch, err := r.peekByte()
	if err != nil {
		return 0, err
	}
	bit := ch & (1 << r.bitIdx)
	if r.bitIdx == 0 {
		r.bitIdx = 7
		r.nextByte()
	} else {
		r.bitIdx--
	}
	if bit != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

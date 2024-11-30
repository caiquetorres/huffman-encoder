package huffman

import (
	"bufio"
	"io"
)

type reader struct {
	r *bufio.Reader
	b byte
	c byte
}

func newReader(r io.Reader) *reader {
	re := bufio.NewReader(r)
	return &reader{
		r: re,
		b: 7,
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
	c, err := r.peekByte()
	if err != nil {
		return 0, err
	}
	bit := c & (1 << r.b)
	if r.b == 0 {
		r.b = 7
		r.nextByte()
	} else {
		r.b--
	}
	if bit != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

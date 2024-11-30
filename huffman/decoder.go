package huffman

import (
	"fmt"
	"io"
	"unicode"
)

type Decoder struct {
	r io.ReadSeeker
	w io.Writer
}

func NewDecoder(r io.ReadSeeker, w io.Writer) *Decoder {
	return &Decoder{r: r, w: w}
}

func (e *Decoder) Decode() error {
	r := newReader(e.r)
	w := newWriter(e.w)
	fileName := make([]byte, 4)
	r.read(fileName)
	if string(fileName) != "HUFF" {
		return fmt.Errorf("file is not huff")
	}
	r.nextByte() // ';'
	countMap := map[byte]uint{}
	byteCount := uint(0)
	for {
		ch, err := r.nextByte()
		if err != nil {
			break
		}
		_, err = r.nextByte() // ':'
		if err != nil {
			break
		}
		freq, err := decodeNumber(r)
		if err != nil {
			break
		}
		byteCount += freq
		countMap[ch] = freq
		ch, err = r.peekByte() // ','
		if err != nil || ch == ';' {
			break
		}
		_, err = r.nextByte()
		if err != nil {
			break
		}
	}
	_, err := r.nextByte() // ';'
	if err != nil {
		return err
	}
	tree := newTree(countMap)
	for range byteCount {
		ch, err := decodeHuff(r, tree.root)
		if err != nil {
			return err
		}
		err = w.writeByte(ch)
		if err != nil {
			return err
		}
	}
	return w.flush()
}

func decodeNumber(r *reader) (uint, error) {
	ans := uint(0)
	for {
		b, err := r.peekByte()
		if err != nil {
			break
		}
		if !unicode.IsNumber(rune(b)) {
			break
		}
		ans += uint(b - '0')
		ans *= 10
		_, err = r.nextByte()
		if err != nil {
			break
		}
	}
	ans /= 10
	return ans, nil
}

func decodeHuff(r *reader, node *node[byte]) (byte, error) {
	if node.l == nil && node.r == nil {
		return node.val, nil
	}
	bit, err := r.nextBit()
	if err != nil {
		return 0, err
	}
	if bit == 0 {
		return decodeHuff(r, node.l)
	} else {
		return decodeHuff(r, node.r)
	}
}

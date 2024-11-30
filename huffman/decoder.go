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
	re := newReader(e.r)
	wr := newWriter(e.w)
	fileName := make([]byte, 4)
	re.read(fileName)
	if string(fileName) != "HUFF" {
		return fmt.Errorf("file is not huff")
	}
	re.nextByte() // ;
	occ := map[byte]uint{}
	byteCount := uint(0)
	for {
		b, err := re.nextByte()
		if err != nil {
			break
		}
		re.nextByte() // :
		freq, err := decodeNumber(re)
		byteCount += freq
		occ[b] = freq
		b, err = re.peekByte() // ,
		if err != nil || b == ';' {
			break
		}
		re.nextByte()
	}
	re.nextByte() // ;
	tree := newTree(occ)
	for range byteCount {
		b, err := decodeHuff(re, tree.r)
		if err != nil {
			return err
		}
		wr.writeByte(b)
	}
	wr.flush()
	return nil
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
		r.nextByte()
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

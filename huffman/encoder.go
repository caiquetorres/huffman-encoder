package huffman

import (
	"io"
	"strconv"
)

type Encoder struct {
	r io.ReadSeeker
	w io.Writer
}

func NewEncoder(r io.ReadSeeker, w io.Writer) *Encoder {
	return &Encoder{r, w}
}

func (e *Encoder) Encode() error {
	w := newWriter(e.w)
	r := newReader(e.r)
	countMap := map[byte]uint{}
	for {
		ch, err := r.nextByte()
		if err != nil {
			break
		}
		if _, ok := countMap[ch]; !ok {
			countMap[ch] = 1
		} else {
			countMap[ch]++
		}
	}
	tree := newTree(countMap)
	_, err := w.writeString("HUFF;")
	if err != nil {
		return err
	}
	hasPrev := false
	for i := range 256 {
		ch := byte(i)
		freq := countMap[ch]
		if freq == 0 {
			continue
		}
		if hasPrev {
			w.writeByte(',')
		}
		hasPrev = true
		_, err := w.write([]byte{ch, ';'})
		if err != nil {
			return err
		}
		_, err = w.writeString(strconv.FormatUint(uint64(freq), 10))
		if err != nil {
			return err
		}
	}
	err = w.writeByte(';')
	if err != nil {
		return err
	}
	_, err = e.r.Seek(int64(0), io.SeekStart)
	if err != nil {
		return err
	}
	bitCount := 0
	for {
		b, err := r.nextByte()
		if err != nil {
			break
		}
		p := tree.path(b)
		for _, d := range p {
			bitCount++
			err = w.writeBit(d)
			if err != nil {
				return err
			}
		}
	}
	remainCount := 8 - bitCount%8
	for range remainCount {
		err = w.writeBit(false)
		if err != nil {
			return err
		}
	}
	return w.flush()
}

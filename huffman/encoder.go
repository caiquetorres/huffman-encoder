package huffman

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Encoder struct {
	r io.ReadSeeker
	w io.Writer
}

func NewEncoder(r io.ReadSeeker, w io.Writer) *Encoder {
	return &Encoder{r: r, w: w}
}

func (e *Encoder) Encode() error {
	wr := newWriter(e.w)
	re := bufio.NewReader(e.r)
	occ := map[byte]uint{}
	for {
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		if _, ok := occ[b]; !ok {
			occ[b] = 1
		} else {
			occ[b]++
		}
	}
	tree := newTree(occ)
	wr.writeString("HUFF;")
	hasPrev := false
	for i := range 256 {
		freq := occ[byte(i)]
		if freq == 0 {
			continue
		}
		if hasPrev {
			wr.writeByte(',')
		}
		hasPrev = true
		wr.writeString(fmt.Sprintf("%s:%s", string(byte(i)), strconv.FormatUint(uint64(freq), 10)))
	}
	wr.writeByte(';')
	e.r.Seek(int64(0), io.SeekStart)
	bitCount := 0
	for {
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		p := tree.Path(b)
		for _, d := range p {
			bitCount++
			wr.writeBit(d - '0')
		}
	}
	remainCount := 8 - bitCount%8
	for range remainCount {
		wr.writeBit(0)
	}
	wr.flush()
	return nil
}

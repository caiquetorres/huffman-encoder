package encoder

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/caiquetorres/compression-tool/huffman"
)

func Encode(r io.ReadSeeker, w io.Writer) error {
	writer := bufio.NewWriter(w)
	reader := bufio.NewReader(r)
	occ := map[byte]uint{}
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if _, ok := occ[b]; !ok {
			occ[b] = 1
		} else {
			occ[b]++
		}
	}
	tree := huffman.NewTree(occ)
	writer.WriteString("HUFF;")
	hasPrev := false
	for i := range 256 {
		freq := occ[byte(i)]
		if freq == 0 {
			continue
		}
		if hasPrev {
			writer.WriteByte(',')
		}
		hasPrev = true
		writer.WriteString(fmt.Sprintf("%s:%s", string(byte(i)), strconv.FormatUint(uint64(freq), 10)))
	}
	writer.WriteByte(';')
	r.Seek(int64(0), io.SeekStart)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		p := tree.Path(b)
		writer.WriteString(string(p)) // we're going to write bits instead
	}
	writer.Flush()
	return nil
}

package huffman

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteBit(t *testing.T) {
	assert := assert.New(t)

	var buffer bytes.Buffer
	writer := newWriter(&buffer)

	bits := []bool{
		true, false, true, false, false, true, false, true,
		true, true, false, false, true, true, false, false,
	}

	for _, bit := range bits {
		err := writer.writeBit(bit)
		assert.NoError(err, "writeBit should not return an error")
	}

	err := writer.flush()
	assert.NoError(err, "flush should not return an error")

	expectedBytes := []byte{0xA5, 0xCC}
	actualBytes := buffer.Bytes()

	assert.Equal(expectedBytes, actualBytes, "The written bytes should match the expected byte sequence")
}

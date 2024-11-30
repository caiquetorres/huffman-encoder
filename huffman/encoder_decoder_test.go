package huffman

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	assert := assert.New(t)

	originalData := []byte("hello huffman encoding and decoding test")

	var encodedBuffer bytes.Buffer
	var decodedBuffer bytes.Buffer

	encoder := NewEncoder(bytes.NewReader(originalData), &encodedBuffer)
	err := encoder.Encode()
	assert.NoError(err, "Encoding should not fail")

	decoder := NewDecoder(bytes.NewReader(encodedBuffer.Bytes()), &decodedBuffer)
	err = decoder.Decode()
	assert.NoError(err, "Decoding should not fail")

	decodedData := decodedBuffer.Bytes()
	assert.Equal(originalData, decodedData, "Decoded data should match the original data")
}

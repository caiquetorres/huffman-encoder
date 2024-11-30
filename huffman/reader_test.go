package huffman

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeekByte(t *testing.T) {
	assert := assert.New(t)

	data := []byte{0xA5, 0xC3} // Example binary data
	reader := newReader(bytes.NewReader(data))

	peekedByte, err := reader.peekByte()
	assert.NoError(err, "peekByte should not return an error")
	assert.Equal(byte(0xA5), peekedByte, "peekByte should return the first byte without consuming it")

	peekedByteAgain, err := reader.peekByte()
	assert.NoError(err, "peekByte should not return an error on repeated calls")
	assert.Equal(byte(0xA5), peekedByteAgain, "peekByte should consistently return the same first byte")
}

func TestNextBit(t *testing.T) {
	assert := assert.New(t)

	data := []byte{0xA5} // Binary: 10100101
	reader := newReader(bytes.NewReader(data))

	expectedBits := []byte{1, 0, 1, 0, 0, 1, 0, 1}
	for i, expectedBit := range expectedBits {
		bit, err := reader.nextBit()
		assert.NoError(err, "nextBit should not return an error")
		assert.Equal(expectedBit, bit, "Bit %d should match expected value", i)
	}

	_, err := reader.nextBit()
	assert.Error(err, "nextBit should return an error when reading beyond available data")
}

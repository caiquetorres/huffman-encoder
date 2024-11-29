package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapPush(t *testing.T) {
	h := newHeap[string]()
	h.push(&node[string]{val: "a", freq: 5})
	h.push(&node[string]{val: "b", freq: 3})
	h.push(&node[string]{val: "c", freq: 8})
	assert.Equal(t, uint(3), h.s, "Expected heap size to be 3")
	assert.Equal(t, uint(3), h.peek().freq, "Expected the root to have frequency 3")
}

func TestHeapPop(t *testing.T) {
	h := newHeap[string]()
	h.push(&node[string]{val: "a", freq: 5})
	h.push(&node[string]{val: "b", freq: 3})
	h.push(&node[string]{val: "c", freq: 8})
	popped := h.pop()
	assert.Equal(t, uint(3), popped.freq, "Expected popped node to have frequency 3")
	assert.Equal(t, uint(2), h.s, "Expected heap size to be 2 after pop")
	assert.Equal(t, uint(5), h.peek().freq, "Expected the root to have frequency 5 after pop")
}

func TestHeapPeek(t *testing.T) {
	h := newHeap[string]()
	h.push(&node[string]{val: "a", freq: 10})
	h.push(&node[string]{val: "b", freq: 1})
	peeked := h.peek()
	assert.Equal(t, uint(1), peeked.freq, "Expected top frequency to be 1")
	assert.Equal(t, uint(2), h.s, "Expected heap size to be 2")
}

func TestHeapEmpty(t *testing.T) {
	h := newHeap[string]()
	assert.Nil(t, h.peek(), "Expected peek to return nil for empty heap")
	assert.Nil(t, h.pop(), "Expected pop to return nil for empty heap")
}

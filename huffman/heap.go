package huffman

type node[T comparable] struct {
	val  T
	freq uint
	l    *node[T]
	r    *node[T]
}

type heap[T comparable] struct {
	size uint
	buf  [256]*node[T]
}

func newHeap[T comparable]() *heap[T] {
	return &heap[T]{}
}

func (h *heap[T]) push(node *node[T]) {
	h.size++
	i := h.size - 1
	h.buf[i] = node
	for i > 0 && h.buf[h.parent(i)].freq > h.buf[i].freq {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *heap[T]) pop() *node[T] {
	if h.size == 0 {
		return nil
	}
	root := h.buf[0]
	h.buf[0] = h.buf[h.size-1]
	h.size--
	h.heapify(0)
	return root
}

func (h *heap[T]) peek() *node[T] {
	if h.size == 0 {
		return nil
	}
	return h.buf[0]
}

func (h *heap[T]) parent(i uint) uint {
	return (i - 1) / 2
}

func (h *heap[T]) left(i uint) uint {
	return 2*i + 1
}

func (h *heap[T]) right(i uint) uint {
	return 2*i + 2
}

func (h *heap[T]) swap(a uint, b uint) {
	temp := h.buf[a]
	h.buf[a] = h.buf[b]
	h.buf[b] = temp
}

func (h *heap[T]) heapify(i uint) {
	l := h.left(i)
	r := h.right(i)
	smallest := i
	if l < h.size && h.buf[l].freq < h.buf[i].freq {
		smallest = l
	}
	if r < h.size && h.buf[r].freq < h.buf[smallest].freq {
		smallest = r
	}
	if smallest != i {
		h.swap(i, smallest)
		h.heapify(smallest)
	}
}

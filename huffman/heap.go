package huffman

type node[T comparable] struct {
	val  T
	freq uint
	l    *node[T]
	r    *node[T]
}

type heap[T comparable] struct {
	s   uint
	buf [256]*node[T]
}

func newHeap[T comparable]() *heap[T] {
	return &heap[T]{}
}

func (h *heap[T]) push(node *node[T]) {
	h.s++
	i := h.s - 1
	h.buf[i] = node
	for i > 0 && h.buf[h.parent(i)].freq > h.buf[i].freq {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *heap[T]) pop() *node[T] {
	if h.s == 0 {
		return nil
	}
	root := h.buf[0]
	h.buf[0] = h.buf[h.s-1]
	h.s--
	h.heapify(0)
	return root
}

func (h *heap[T]) peek() *node[T] {
	if h.s == 0 {
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
	s := i
	if l < h.s && h.buf[l].freq < h.buf[i].freq {
		s = l
	}
	if r < h.s && h.buf[r].freq < h.buf[s].freq {
		s = r
	}
	if s != i {
		h.swap(i, s)
		h.heapify(s)
	}
}

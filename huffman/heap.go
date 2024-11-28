package huffman

type heap struct {
	s   uint
	buf [256]*node
}

func newHeap() *heap {
	return &heap{}
}

func (h *heap) push(node *node) {
	h.s++
	i := h.s - 1
	h.buf[i] = node
	for i > 0 && h.buf[h.parent(i)].freq > h.buf[i].freq {
		h.swap(i, h.parent(i))
		i = h.parent(i)
	}
}

func (h *heap) pop() *node {
	if h.s == 0 {
		return nil
	}
	root := h.buf[0]
	h.buf[0] = h.buf[h.s-1]
	h.s--
	h.heapify(0)
	return root
}

func (h *heap) peek() *node {
	if h.s == 0 {
		return nil
	}
	return h.buf[0]
}

func (h *heap) parent(i uint) uint {
	return (i - 1) / 2
}

func (h *heap) left(i uint) uint {
	return 2*i + 1
}

func (h *heap) right(i uint) uint {
	return 2*i + 2
}

func (h *heap) swap(a uint, b uint) {
	temp := h.buf[a]
	h.buf[a] = h.buf[b]
	h.buf[b] = temp
}

func (h *heap) heapify(i uint) {
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

package huffman

type path = []byte

type Tree struct {
	r *node
	m map[byte]path
}

func NewTree(o map[byte]uint) *Tree {
	h := newHeap()
	for i := range 256 {
		b := byte(i)
		if o[b] != 0 {
			n := &node{val: b, freq: o[b], l: nil, r: nil}
			h.push(n)
		}
	}
	for h.s > 1 {
		l := h.pop()
		r := h.pop()
		freq := l.freq + r.freq
		n := &node{val: ' ', freq: freq, l: l, r: r}
		h.push(n)
	}
	m := map[byte]path{}
	r := h.peek()
	fill(m, path{}, r)
	return &Tree{r: r, m: m}
}

func (t *Tree) Path(b byte) path {
	return t.m[b]
}

func fill(m map[byte]path, p path, n *node) {
	if n == nil {
		return
	}
	if n.l == nil && n.r == nil {
		dst := make(path, len(p))
		for i, v := range p {
			dst[i] = v
		}
		m[n.val] = dst
	}
	p = append(p, '0')
	fill(m, p, n.l)
	p = p[:len(p)-1]
	p = append(p, '1')
	fill(m, p, n.r)
	p = p[:len(p)-1]
}

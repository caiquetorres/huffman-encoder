package huffman

type path = []byte

type tree struct {
	r *node[byte]
	m map[byte]path
}

func newTree(o map[byte]uint) *tree {
	h := newHeap[byte]()
	for i := range 256 {
		b := byte(i)
		if o[b] != 0 {
			n := &node[byte]{val: b, freq: o[b], l: nil, r: nil}
			h.push(n)
		}
	}
	for h.s > 1 {
		l := h.pop()
		r := h.pop()
		freq := l.freq + r.freq
		n := &node[byte]{val: ' ', freq: freq, l: l, r: r}
		h.push(n)
	}
	m := map[byte]path{}
	r := h.peek()
	fill(m, path{}, r)
	return &tree{r: r, m: m}
}

func (t *tree) Path(b byte) path {
	return t.m[b]
}

func fill(m map[byte]path, p path, n *node[byte]) {
	if n == nil {
		return
	}
	if n.l == nil && n.r == nil {
		m[n.val] = make(path, len(p))
		for i, v := range p {
			m[n.val][i] = v
		}
	}
	p = append(p, '0')
	fill(m, p, n.l)
	p = p[:len(p)-1]
	p = append(p, '1')
	fill(m, p, n.r)
	p = p[:len(p)-1]
}

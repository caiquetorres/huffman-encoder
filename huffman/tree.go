package huffman

type path = []bool

type tree struct {
	root    *node[byte]
	codeMap map[byte]path
}

func newTree(o map[byte]uint) *tree {
	h := newHeap[byte]()
	for i := range 256 {
		ch := byte(i)
		if o[ch] != 0 {
			n := &node[byte]{val: ch, freq: o[ch], l: nil, r: nil}
			h.push(n)
		}
	}
	for h.size > 1 {
		l := h.pop()
		r := h.pop()
		freq := l.freq + r.freq
		n := &node[byte]{val: ' ', freq: freq, l: l, r: r}
		h.push(n)
	}
	codeMap := map[byte]path{}
	r := h.peek()
	fill(codeMap, path{}, r)
	return &tree{r, codeMap}
}

func (t *tree) path(b byte) path {
	return t.codeMap[b]
}

func fill(codeMap map[byte]path, p path, n *node[byte]) {
	if n == nil {
		return
	}
	if n.l == nil && n.r == nil {
		codeMap[n.val] = make(path, len(p))
		for i, v := range p {
			codeMap[n.val][i] = v
		}
	}
	p = append(p, false)
	fill(codeMap, p, n.l)
	p = p[:len(p)-1]
	p = append(p, true)
	fill(codeMap, p, n.r)
	p = p[:len(p)-1]
}

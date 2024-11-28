package huffman

type node struct {
	val  byte
	freq uint
	l    *node
	r    *node
}

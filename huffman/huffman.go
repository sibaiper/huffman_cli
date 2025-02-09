package huffman

// Huffman Node structure
type HuffmanNode struct {
	Char  rune
	Freq  int
	Left  *HuffmanNode
	Right *HuffmanNode
}

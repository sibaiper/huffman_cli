package huffman

import (
	"fmt"
	"io"
)

// build Huffman Tree
func BuildHuffmanTree(text string) *HuffmanNode {
	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	pq := PriorityQueue{}
	for char, freq := range freqMap {
		pq.Push(&HuffmanNode{Char: char, Freq: freq})
	}

	for len(pq) > 1 {
		left := pq.Pop()
		right := pq.Pop()
		pq.Push(&HuffmanNode{Freq: left.Freq + right.Freq, Left: left, Right: right})
	}
	return pq.Pop()
}

// serialize Huffman Tree
func SerializeTree(node *HuffmanNode) string {
	if node == nil {
		return ""
	}
	if node.Char != 0 {
		return fmt.Sprintf("1%c", node.Char)
	}
	return "0" + SerializeTree(node.Left) + SerializeTree(node.Right)
}

func DeserializeTree(reader io.Reader) *HuffmanNode {
	var buf [1]byte

	// read first byte (it determines if it's a leaf or internal node)
	if _, err := reader.Read(buf[:]); err != nil {
		return nil
	}
	if buf[0] == '1' { // Leaf node
		// read character and create leaf node
		if _, err := reader.Read(buf[:]); err != nil {
			return nil
		}
		return &HuffmanNode{Char: rune(buf[0])}
	} else if buf[0] == '0' { // internal node
		// recursively deserialize left and right children
		left := DeserializeTree(reader)
		right := DeserializeTree(reader)
		return &HuffmanNode{Left: left, Right: right}
	}

	// if neither '0' nor '1', return nil (unexpected case)
	return nil
}

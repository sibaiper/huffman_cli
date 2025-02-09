package huffman

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// generate Huffman codes
func GenerateHuffmanCodes(node *HuffmanNode, prefix string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.Char != 0 {
		codes[node.Char] = prefix
	}
	GenerateHuffmanCodes(node.Left, prefix+"0", codes)
	GenerateHuffmanCodes(node.Right, prefix+"1", codes)
}

// pack bit string into bytes
func PackBits(bitString string) []byte {
	var buf bytes.Buffer
	bitLen := len(bitString)
	for i := 0; i < bitLen; i += 8 {
		var byteVal byte
		for j := 0; j < 8 && i+j < bitLen; j++ {
			if bitString[i+j] == '1' {
				byteVal |= 1 << (7 - j)
			}
		}
		buf.WriteByte(byteVal)
	}
	return buf.Bytes()
}

/*
.huff File Format:

[4 bytes]  Magic Number (e.g., "HUFF")
[4 bytes]  Tree Length (int32)
[N bytes]  Serialized Huffman Tree
[4 bytes]  Original Text Length (int32)
[4 bytes]  Bit Length (int32)
[M bytes]  Huffman-encoded binary data
*/

// save compressed file
func SaveCompressedFile(filename string, treeData string, compressedData []byte, originalSize int, bitLength int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// write Magic Number
	file.Write([]byte("HUFF"))

	// write Metadata
	binary.Write(file, binary.LittleEndian, int32(len(treeData))) // Tree Length
	file.WriteString(treeData)                                    // Serialized Tree
	binary.Write(file, binary.LittleEndian, int32(originalSize))  // Original file size
	binary.Write(file, binary.LittleEndian, int32(bitLength))     // Bit length

	// write Compressed Data
	file.Write(compressedData)

	return nil
}

func HuffmanEncoding(text string, inputfile string, outputfile string) error {
	// build Huffman Tree
	root := BuildHuffmanTree(text)

	// generate Huffman Codes
	codes := make(map[rune]string)
	GenerateHuffmanCodes(root, "", codes)

	// encode the Text
	var encodedText string
	for _, char := range text {
		encodedText += codes[char]
	}

	// serialize the Tree
	treeData := SerializeTree(root)

	// pack the Bits
	compressedData := PackBits(encodedText)

	// save the Compressed File
	err := SaveCompressedFile(outputfile, treeData, compressedData, len(text), len(encodedText))
	if err != nil {
		return err
	}

	// print statements to indicate compression success
	fmt.Printf("File '%s' compressed successfully and saved as '%s'.\n", inputfile, outputfile)
	fmt.Printf("Original size: %d bytes\n", len(text))
	fmt.Printf("Compressed size: %d bytes\n", len(compressedData)+4+4+len(treeData)+4+4) // calc and print the full compressed file size
	return nil
}

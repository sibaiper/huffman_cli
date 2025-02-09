package huffman

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// unpack bytes into a bitstring
func UnpackBits(data []byte, bitLength int) string {
	var bitString string
	for i, b := range data {
		for j := 7; j >= 0; j-- {
			if i*8+(7-j) >= bitLength {
				break
			}
			if (b>>j)&1 == 1 {
				bitString += "1"
			} else {
				bitString += "0"
			}
		}
	}
	return bitString
}

// decode compressed data using Huffman tree
func DecodeHuffmanData(bitString string, root *HuffmanNode, originalSize int) []byte {
	var decodedData []byte
	node := root
	count := 0

	for _, bit := range bitString {
		if bit == '0' {
			node = node.Left
		} else {
			node = node.Right
		}

		// if we reach a leaf node
		if node.Left == nil && node.Right == nil {
			decodedData = append(decodedData, byte(node.Char))
			node = root
			count++
			if count >= originalSize {
				break
			}
		}
	}
	return decodedData
}

func LoadCompressedFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// read Magic Number
	magic := make([]byte, 4)
	if _, err := file.Read(magic); err != nil || string(magic) != "HUFF" {
		return nil, fmt.Errorf("invalid file format or magic number mismatch")
	}

	// read Tree Length
	var treeLength int32
	if err := binary.Read(file, binary.LittleEndian, &treeLength); err != nil {
		return nil, fmt.Errorf("failed to read tree length: %w", err)
	}

	// read Serialized Huffman Tree
	treeData := make([]byte, treeLength)
	if _, err := io.ReadFull(file, treeData); err != nil {
		return nil, fmt.Errorf("failed to read serialized tree: %w", err)
	}

	// rebuild the Huffman tree
	root := DeserializeTree(bytes.NewReader(treeData))
	if root == nil {
		return nil, fmt.Errorf("failed to deserialize Huffman tree")
	}

	// read Original File Size
	var originalSize int32
	if err := binary.Read(file, binary.LittleEndian, &originalSize); err != nil {
		return nil, fmt.Errorf("failed to read original size: %w", err)
	}

	// read Bit Length
	var bitLength int32
	if err := binary.Read(file, binary.LittleEndian, &bitLength); err != nil {
		return nil, fmt.Errorf("failed to read bit length: %w", err)
	}

	// read Compressed Data
	compressedData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read compressed data: %w", err)
	}

	// convert compressed data into bit string
	bitString := UnpackBits(compressedData, int(bitLength))

	// cecode bit string using Huffman tree
	decodedData := DecodeHuffmanData(bitString, root, int(originalSize))

	return decodedData, nil
}

// HuffmanDecode - reads a compressed file and saves the decoded text
func HuffmanDecode(inputFile, outputFile string) error {
	decodedData, err := LoadCompressedFile(inputFile)
	if err != nil {
		return err
	}

	// write to output file
	err = os.WriteFile(outputFile, decodedData, 0644)
	if err != nil {
		return err
	}

	//fmt.Println("Decoded file saved as:", outputFile)
	fmt.Printf("File '%s' de-compressed successfully to '%s'.\n", inputFile, outputFile)
	return nil
}

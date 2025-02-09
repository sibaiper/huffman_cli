/*
MIT License
Copyright © 2025 Sibai
*/
package cmd

import (
	"fmt"
	"huffman-cli/huffman"

	"os"

	"github.com/spf13/cobra"
)

var inputFile, outputFile string

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encodes a text file using Huffman encoding",
	Long:  `encodes a text file using the Huffman algorithm. Do not use this command to encode binary files or files with images or anything other than text.`,
	Run: func(cmd *cobra.Command, args []string) {
		// validate input file
		data, err := os.ReadFile(inputFile)
		if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		// perform Huffman Encoding
		err = huffman.HuffmanEncoding(string(data), inputFile, outputFile)
		if err != nil {
			fmt.Println("Error during encoding:", err)
		}

	},
}

/*
Encode a text file
./huffman-cli encode -i input.txt -o encoded.bin
*/

func init() {
	rootCmd.AddCommand(encodeCmd)
	// default cobra generated examples
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	encodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file (required)")
	encodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (required)")
	encodeCmd.MarkFlagRequired("input")
	encodeCmd.MarkFlagRequired("output")
}

/*
MIT License
Copyright © 2025 Sibai
*/
package cmd

import (
	"fmt"
	"huffman-cli/huffman"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a Huffman encoded file",
	Long:  `Decodes a file that was encoded using the Huffman algorithm.`,
	Run: func(cmd *cobra.Command, args []string) {
		// validate input file
		if inputFile == "" || outputFile == "" {
			fmt.Println("Both input and output files are required")
			return
		}

		// perform Huffman decoding
		err := huffman.HuffmanDecode(inputFile, outputFile)
		if err != nil {
			fmt.Println("Error during decoding:", err)
			return
		}

	},
}

/*
Decode an encoded file
./huffman-cli decode -i encoded.bin -o decoded.txt
*/

func init() {
	rootCmd.AddCommand(decodeCmd)

	// default cobra generated examples
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	decodeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file (required)")
	decodeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (required)")
	decodeCmd.MarkFlagRequired("input")
	decodeCmd.MarkFlagRequired("output")
}

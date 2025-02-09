# HUFFMAN ENCODING/DECODING CLI

***this is merely a hobby project and should not be used in any professional environment. This tool only decodes files encoded using itself. It will probably produce nonsense if used elsewhere***   

A simple CLI tool that encodes and decodes (text) files using the huffman compressing algorithm.


## BUILDING
1. clone this repository
2. in the root directory run `go build -o huffman-cli`
3. wait
4. move to the next step (USAGE)


## USAGE

```bash
./huffman-cli <action> -i <input-file-name> -o <output-file-name>
```
example usage:  
```bash
./huffman-cli encode -i example.txt -o encoded.bin
```

```bash
./huffman-cli decode -i encoded.bin -o decoded.txt
```

---
From few observations, the compression works best on larger files. On smaller files, it produces bigger files than the original since it is storing meta-data in the encoded file (for decoding).

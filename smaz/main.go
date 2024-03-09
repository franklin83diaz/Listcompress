package main

import (
	"fmt"

	"github.com/kjk/smaz"
)

func main() {
	s := "se acabo, es super bueno y muy divertido, pero no es para mi"
	compressed := smaz.Encode(nil, []byte(s))
	decompressed, err := smaz.Decode(nil, compressed)
	if err != nil {
		fmt.Printf("decompressed: %s\n", string(decompressed))
	}
	//length of compressed data
	fmt.Printf("compressed: %d\n", len(compressed))
	//length of original data
	fmt.Printf("original: %d\n", len(s))

	rate := float64(len(compressed)) / float64(len(s))
	fmt.Printf("compression rate: %.2f\n", rate)

	fmt.Printf("compressed: %s\n", string(decompressed))
}

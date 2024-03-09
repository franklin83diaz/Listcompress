package main

import (
	"fmt"

	"github.com/klauspost/compress/zstd"
)

func compressBytes(src []byte) ([]byte, error) {
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}
	return encoder.EncodeAll(src, nil), nil
}

func decompressBytes(compressed []byte) ([]byte, error) {
	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}
	defer decoder.Close()
	return decoder.DecodeAll(compressed, nil)
}

func main() {
	originalBytes := []byte("Estos son algunos bytes que queremos comprimir")
	fmt.Println("Original bytes:", originalBytes)

	compressedBytes, err := compressBytes(originalBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Compressed bytes:", compressedBytes)

	decompressedBytes, err := decompressBytes(compressedBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decompressed bytes:", decompressedBytes)
}

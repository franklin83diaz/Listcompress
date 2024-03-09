package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

// compressString utiliza gzip para comprimir un string y devuelve un slice de bytes.
func compressString(str string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(str))
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// decompressString toma un slice de bytes comprimido con gzip y devuelve el string original.
func decompressString(compressed []byte) (string, error) {
	buf := bytes.NewBuffer(compressed)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer gz.Close()

	str, err := ioutil.ReadAll(gz)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func main() {
	originalString := "El texto que quieres comprimir va aqui, asi es como funciona el algoritmo de compresion."
	fmt.Println("Original string:", originalString)
	//length of original string
	fmt.Println("Length of original string:", len(originalString))

	compressedString, err := compressString(originalString)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Compressed string:", compressedString)
	//length of compressed string
	fmt.Println("Length of compressed string:", len(compressedString))

	decompressedString, err := decompressString(compressedString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decompressed string:", decompressedString)
}

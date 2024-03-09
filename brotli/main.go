package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/andybalholm/brotli"
)

func main() {
	// Datos de ejemplo para comprimir
	inputData := []byte("dfs689fi43tdfs689fi4689fi43tdf689fi4s689fi43t")

	// Comprimir los datos
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.DefaultCompression)
	_, err := w.Write(inputData)
	if err != nil {
		panic(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}
	compressedData := b.Bytes()

	// Descomprimir los datos
	r := brotli.NewReader(bytes.NewReader(compressedData))
	decompressedData, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Datos descomprimidos:", string(decompressedData))

	// Imprimir la longitud de los datos comprimidos y descomprimidos
	fmt.Printf("Datos comprimidos: %d\n", len(compressedData))
	fmt.Printf("Datos originales: %d\n", len(inputData))
	rate := float64(len(compressedData)) / float64(len(inputData))
	fmt.Printf("Tasa de compresión: %.2f\n", rate)

}

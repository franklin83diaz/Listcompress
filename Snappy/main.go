package main

import (
	"fmt"
	"log"

	"github.com/golang/snappy"
)

func main() {
	// Define la cadena de texto original
	originalText := "Este es un texto de ejemplo que será comprimido y luego descomprimido usando Snappy."

	// Comprime la cadena de texto
	compressed := snappy.Encode(nil, []byte(originalText))

	fmt.Println("Texto comprimido:", compressed)

	// Descomprime la cadena de texto
	decompressed, err := snappy.Decode(nil, compressed)
	if err != nil {
		log.Fatalf("Error al descomprimir: %v", err)
	}

	// Convierte el resultado descomprimido a una cadena de texto
	decompressedText := string(decompressed)

	fmt.Println("Texto descomprimido:", decompressedText)

	// Imprime la longitud de los datos comprimidos y descomprimidos
	fmt.Printf("Texto comprimido: %d\n", len(compressed))
	fmt.Printf("Texto original: %d\n", len(originalText))
	rate := float64(len(compressed)) / float64(len(originalText))
	fmt.Printf("Tasa de compresión: %.2f\n", rate)
}

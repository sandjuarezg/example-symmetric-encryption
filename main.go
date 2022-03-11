package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func main() {
	text := []byte("This message is secret >u<")
	key := []byte("123 123 123 123 123 123 123 123!") // 32 bytes

	// create new cipher block
	// 16 bytes - AES-128
	// 24 bytes - AES-192
	// 32 bytes - AES-256 - the safer
	ciph, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// slice with block size
	aux := make([]byte, ciph.BlockSize())

	// encrypt
	ciph.Encrypt(aux, text)
	fmt.Println("encrypt:", string(aux), len(aux))

	// decrypt
	ciph.Decrypt(text, aux)
	fmt.Println("decrypt:", string(text), len(text))
}

package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func main() {
	key := []byte("123456789abcdefghijklmno")
	text := []byte("This message is secret wiu top secret")

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encrypt := make([]byte, block.BlockSize())
	block.Encrypt(encrypt, text)
	fmt.Println(string(encrypt))

	decrypt := make([]byte, 1024)
	block.Decrypt(decrypt, encrypt)
	fmt.Println(string(decrypt))
}

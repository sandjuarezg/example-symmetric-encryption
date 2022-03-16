package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
)

func main() {
	key := []byte("123456789abcdefghijklmno")          // 16, 24, 32
	iv := []byte("1234567890000000")                   // 16
	text := []byte("This message is secret wiiiiiiiu") // multiple of block size (16, 32, 48...)
	encrypt := make([]byte, len(text))
	decrypt := make([]byte, len(text))

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// encrypt
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypt, text)
	fmt.Println(string(encrypt))

	// decrypt
	mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypt, encrypt)
	fmt.Println(string(decrypt))
}

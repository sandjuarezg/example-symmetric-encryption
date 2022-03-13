package main

import (
	"crypto/aes"
	"crypto/cipher"
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

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, aesgcm.NonceSize())

	// encrypt
	encrypt := aesgcm.Seal(nil, nonce, text, nil)
	fmt.Println(string(encrypt))

	//decrypt
	decrypt, err := aesgcm.Open(nil, nonce, encrypt, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decrypt))
}

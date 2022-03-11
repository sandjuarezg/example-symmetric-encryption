package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func main() {
	text := []byte("This message is secret wiuuuuuuuuuuui top secret")
	key := []byte("123456789abcdefghijklmnopqrstuvw")

	// encrypt
	// create a block cipher
	cpbl, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// cipher package - implements standard block cipher
	//
	gcm, err := cipher.NewGCM(cpbl)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Fatal(err)
	}

	encrypt := gcm.Seal(nonce, nonce, text, nil)

	fmt.Println(string(encrypt))

	// decrypt
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err = cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encrypt) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := encrypt[:nonceSize], encrypt[nonceSize:]
	decrypt, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(decrypt))
}

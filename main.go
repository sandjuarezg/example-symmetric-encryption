package main

import (
	"crypto/rc4"
	"fmt"
	"log"
)

func main() {
	key := []byte("1234")
	text := []byte("This message is secret wiu")
	encrypt := make([]byte, len(text))
	decrypt := make([]byte, len(text))

	// encrypt
	rc4C, err := rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	rc4C.XORKeyStream(encrypt, text)
	fmt.Println(string(encrypt))

	// decrypt
	rc4C, err = rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	rc4C.XORKeyStream(decrypt, encrypt)
	fmt.Println(string(decrypt))
}

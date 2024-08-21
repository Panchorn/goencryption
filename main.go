package main

import (
	"encryptiondynamicnonce/encryption"
	"fmt"
)

func main() {
	key := []byte("thisisakey123456") // เปลี่ยนเป็นกุญแจของคุณ
	data := []byte("my name is panchorn.")

	// เข้ารหัส
	ciphertext, err := encryption.Encrypt(data, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// ถอดรหัส
	plaintext, err := encryption.Decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Plaintext: %s\n", plaintext)
}

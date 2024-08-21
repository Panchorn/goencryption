package encryption

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		t.Fatalf("Failed to generate random key: %v", err)
	}
	testcases := []struct {
		name      string
		plaintext string
	}{
		{"Short message", "Hello, world!"},
		{"Long message", "This is a much longer test message."},
		{"Empty message", ""},
		{"Name", "Mr.Somchai"},
	}

	for _, tc := range testcases {
		t.Run(tc.plaintext, func(test *testing.T) {
			ciphertext, err := Encrypt([]byte(tc.plaintext), key)

			// Encode the ciphertext to base64 for storage or transmission
			ciphertextEncoded := base64.StdEncoding.EncodeToString(ciphertext)
			fmt.Printf("%v\n", ciphertext)
			fmt.Printf("%v\n", ciphertextEncoded)

			// Decode the ciphertext from base64 before decryption
			ciphertextDecoded, err := base64.StdEncoding.DecodeString(ciphertextEncoded)
			if err != nil {
				t.Fatalf("Decoding failed: %v", err)
			}

			plaintextDecrypted, err := Decrypt(ciphertextDecoded, key)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			if string(plaintextDecrypted) != tc.plaintext {
				t.Errorf("Decrypted plaintext doesn't match original: expected %q, got %q", tc.plaintext, string(plaintextDecrypted))
			}
		})
	}
}

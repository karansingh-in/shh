package main

import (
	bytes "bytes"
	"fmt"
	testing "testing"
)

func TestDeriveKey(t *testing.T) {
	password := []byte("my password")
	salt1 := []byte("my salt")
	salt2 := []byte("different salt")

	key1 := DeriveKey(password, salt1)
	key2 := DeriveKey(password, salt1) // same password, same salt
	key3 := DeriveKey(password, salt2) // same password, different salt

	if len(key1) != 32 {
		t.Errorf("expected key length 32, got %d", len(key1))
	}
	if !bytes.Equal(key1, key2) {
		t.Error("same password+salt should produce identical keys")
	}
	if bytes.Equal(key1, key3) {
		t.Error("different salts should produce different keys")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	password := []byte("my password")
	salt1 := []byte("my salt")
	plaintext1 := "this is the plaintext"

	key1 := DeriveKey(password, salt1)

	ciphertext, nonce, err := Encrypt(key1, []byte(plaintext1))
	if err != nil {
		t.Errorf("encryption failed: %v", err)
	}
	fmt.Printf("%s", ciphertext)
	fmt.Printf("%s", nonce)

	plaintext2, err := Decrypt(key1, nonce, ciphertext)
	if err != nil {
		t.Errorf("decryption failed: %v", err)
	}
	if !bytes.Equal([]byte(plaintext1), plaintext2) {
		t.Errorf("files do not match after encryption decryption")
	}
}

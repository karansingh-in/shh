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

func TestDecryptWrongKey(t *testing.T) {
	password := []byte("my password")
	wrongPassword := []byte("not my password")
	salt := []byte("my salt")
	plaintext := "this is the plaintext"

	key := DeriveKey(password, salt)
	wrongKey := DeriveKey(wrongPassword, salt)

	ciphertext, nonce, err := Encrypt(key, []byte(plaintext))
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	_, err = Decrypt(wrongKey, nonce, ciphertext)
	if err == nil {
		t.Error("expected decryption with wrong key to fail, but it succeeded")
	}
}

func TestSaveLoadVaultRoundTrip(t *testing.T) {
	path := "test_vault.enc"

	password := []byte("my real password")

	// first run: file doesn't exist yet
	v, err := LoadVault(password, path)
	if err != nil {
		t.Fatalf("file not found: %v", err)
	}

	entry := Entry{Name: "test entry", Body: "some journal content"}
	if err := v.Add(entry); err != nil {
		t.Fatalf("add failed: %v", err)
	}

	if err := SaveVault(v, password, path); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	// load it back with the correct password
	loaded, err := LoadVault(password, path)
	if err != nil {
		t.Fatalf("load with correct password failed: %v", err)
	}

	got, err := loaded.Get("test entry")
	if err != nil {
		t.Fatalf("entry missing after reload: %v", err)
	}
	if got.Body != entry.Body {
		t.Errorf("body mismatch: expected %q, got %q", entry.Body, got.Body)
	}

	// checking access with a WRONG password
	wrongPassword := []byte("not the real password")
	_, err = LoadVault(wrongPassword, path)
	if err == nil {
		t.Error("expected error loading with wrong password, got nil")
	}
}

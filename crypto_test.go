package main

import (
	bytes "bytes"
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

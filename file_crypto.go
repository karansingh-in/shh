package main

import (
	"fmt"
	"os"
	"strings"
)

func cmdEncryptFile(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: shh encrypt <path>")
	}
	filePath := args[0]

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	salt, err := GenerateSalt()
	if err != nil {
		return err
	}

	key := DeriveKey(password, salt)

	ciphertext, nonce, err := Encrypt(key, fileBytes)
	if err != nil {
		return err
	}

	var out []byte
	out = append(out, salt...)
	out = append(out, nonce...)
	out = append(out, ciphertext...)

	outPath := filePath + ".shh"
	if err := os.WriteFile(outPath, out, 0600); err != nil {
		return err
	}

	fmt.Println("encrypted:", outPath)
	return nil
}

func cmdDecryptFile(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: shh decrypt <path>")
	}
	filePath := args[0]

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if len(data) < saltSize+nonceSize {
		return fmt.Errorf("file is too small to be a valid .shh encrypted file")
	}
	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+nonceSize]
	ciphertext := data[saltSize+nonceSize:]

	key := DeriveKey(password, salt)

	plaintext, err := Decrypt(key, nonce, ciphertext)
	if err != nil {
		return err
	}

	outPath := strings.TrimSuffix(filePath, ".shh")
	if outPath == filePath {
		outPath = filePath + ".decrypted"
	}

	if err := os.WriteFile(outPath, plaintext, 0600); err != nil {
		return err
	}

	fmt.Println("decrypted:", outPath)
	return nil
}

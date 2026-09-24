package main

import (
	"encoding/json"
	"os"
)

func SaveVault(v *Vault, path string) error {
	data, err := json.Marshal(v.secrets)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func LoadVault(path string) (*Vault, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return NewVault(), nil
		}
		return nil, err
	}
	secrets := make(map[string]Secret)
	err = json.Unmarshal(data, &secrets)

	return &Vault{secrets: secrets}, nil
}

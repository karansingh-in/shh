package main

import (
	"encoding/json"
	"os"
)

func SaveVault(v *Vault, path string) error {
	// converting go's structure to json file
	data, err := json.Marshal(v.entries)
	if err != nil {
		return err
	}
	os.WriteFile(path, data, 0600)
	return nil
}

func LoadVault(path string) (*Vault, error) {
	// reading the file
	data, err := os.ReadFile(path)
	if err != nil {
		// if the error says that file doesn't exist, then create a new vault
		if os.IsNotExist(err) {
			return NewVault(), nil
		}
		// else show the error message
		return &Vault{}, err
	}
	loaded_entries := make(map[string]Entry)
	err = json.Unmarshal(data, loaded_entries)

	return &Vault{entries: loaded_entries}, nil
}

package main

import (
	"fmt"
)

// declaring the variables inside a data structure named Secret
type Secret struct {
	Name     string
	Username string
	Password string
	Notes    string
}

// create a vault hashmap
type Vault struct {
	secrets map[string]Secret
}

// NewVault returns a pointer to vault, initializing Secret structure, this is the constructor
func NewVault() *Vault {
	return &Vault{
		secrets: make(map[string]Secret),
	}
}

// taking input of the data type Secret, returning errors, the function belongs to a vault pointer
func (v *Vault) Add(s Secret) error {
	// checking before adding a new username if it already exists
	if _, exists := v.secrets[s.Name]; exists {
		return fmt.Errorf("%s has already been used", s.Name)
	}
	v.secrets[s.Name] = s
	return nil // no errors
}

func (v *Vault) Get(name string) (Secret, error) {
	s, exists := v.secrets[name]
	if !exists {
		return Secret{}, fmt.Errorf("name %s not found", name)
	}
	return s, nil // returning the found name
}

func (v *Vault) Update(name string, s Secret) error {
	// check if name exists
	_, exists := v.secrets[name]
	if !exists {
		return fmt.Errorf("name %s does not exist", name)
	}
	v.secrets[name] = s
	return nil
}

func (v *Vault) Delete(name string) error {
	// check if it exists
	_, exists := v.secrets[name]
	if !exists {
		return fmt.Errorf("name %s not found", name)
	}
	delete(v.secrets, name)
	return nil
}

func (v *Vault) List() []Secret {
	result := make([]Secret, 0, len(v.secrets))
	for _, s := range v.secrets {
		result = append(result, s)
	}
	return result
}

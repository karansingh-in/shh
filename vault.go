package main

import (
	"fmt"
	"time"
)

// creating a data structure named Entry
type Entry struct {
	Name    string
	Body    string
	Created time.Time
}

// hashmap for entries
type Vault struct {
	entries map[string]Entry
}

// constructor for the data structure
func NewVault() *Vault {
	return &Vault{
		make(map[string]Entry),
	}
}

func (v *Vault) Add(e Entry) error {
	// check if it already exists
	_, exists := v.entries[e.Name]
	if exists {
		return fmt.Errorf("a file with this filename already exists")
	}
	v.entries[e.Name] = e
	return nil
}

func (v *Vault) Delete(e Entry) error {
	// check if the file exists before deleting it
	_, exists := v.entries[e.Name]
	if !exists {
		return fmt.Errorf("the file does not exist")
	}
	delete(v.entries, e.Name)
	return nil
}

func (v *Vault) Update(name string, e Entry) error {
	// check if the file exists
	_, exists := v.entries[name]
	if !exists {
		return fmt.Errorf("the file does not exist, create the file first")
	}
	v.entries[name] = e
	return nil
}

func (v *Vault) Get(name string) (Entry, error) {
	e, exists := v.entries[name]
	if !exists {
		return Entry{}, fmt.Errorf("the file does not exist")
	}
	return e, nil
}

func (v *Vault) List() []Entry {
	result := make([]Entry, 0, len(v.entries))
	for _, e := range v.entries {
		result = append(result, e)
	}
	return result
}

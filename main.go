package main

import (
	"fmt"
)

func main() {
	// --- build a fresh vault and add a couple entries ---
	v := NewVault()

	v.Add(Secret{Name: "github", Username: "spartan", Password: "hunter2", Notes: "work account"})
	v.Add(Secret{Name: "gmail", Username: "spartan@gmail.com", Password: "changeme", Notes: ""})

	// --- save it to disk ---
	if err := SaveVault(v, "test.json"); err != nil {
		fmt.Println("save failed:", err)
		return
	}
	fmt.Println("saved ok")

	// --- load it back into a SEPARATE variable, proving it round-tripped ---
	loaded, err := LoadVault("test.json")
	if err != nil {
		fmt.Println("load failed:", err)
		return
	}

	fmt.Println("loaded secrets:")
	for _, s := range loaded.List() {
		fmt.Printf("  %+v\n", s)
	}

	// --- sanity check: does a specific entry match what we saved? ---
	gh, err := loaded.Get("github")
	if err != nil {
		fmt.Println("github entry missing after reload:", err)
		return
	}
	fmt.Println("github password after reload:", gh.Password)
}

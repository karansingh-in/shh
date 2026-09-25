package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/term"
)

func VaultPath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".shh")
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "vault")

	return path, nil
}

func promptPassword(prompt string) ([]byte, error) {

	// to avoid the password echo in the terminal
	fmt.Print(prompt)
	input := os.Stdin.Fd()
	password, err := term.ReadPassword(int(input))
	if err != nil {
		return nil, err
	}
	fmt.Println()
	return password, nil
}

func promptLine(prompt string) (string, error) {

	// for all other echo
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func cmdAdd(args []string) error {
	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}
	path, err := VaultPath()
	if err != nil {
		return err
	}
	v, err := LoadVault(password, path)
	if err != nil {
		return err
	}
	name, err := promptLine("Name: ")
	if err != nil {
		return err
	}

	body, err := promptLine("Body: ")
	if err != nil {
		return err
	}
	entry := Entry{
		Name:    name,
		Body:    body,
		Created: time.Now(),
	}
	if err := v.Add(entry); err != nil {
		return err
	}
	if err := SaveVault(v, password, path); err != nil {
		return err
	}
	return nil
}
func cmdGet(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: shh get <name>")
	}

	name := args[0]

	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	path, err := VaultPath()
	if err != nil {
		return err
	}

	v, err := LoadVault(password, path)
	if err != nil {
		return err
	}

	entry, err := v.Get(name)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", entry.Name)
	fmt.Printf("Body: %s\n", entry.Body)
	fmt.Printf("Created: %s\n", entry.Created)

	return nil
}

func cmdList(args []string) error {
	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	path, err := VaultPath()
	if err != nil {
		return err
	}

	v, err := LoadVault(password, path)
	if err != nil {
		return err
	}

	entries := v.List()

	for _, entry := range entries {
		fmt.Println(entry.Name)
	}

	return nil
}

func cmdDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: shh delete <name>")
	}

	name := args[0]

	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	path, err := VaultPath()
	if err != nil {
		return err
	}

	v, err := LoadVault(password, path)
	if err != nil {
		return err
	}

	entry, err := v.Get(name)
	if err != nil {
		return err
	}

	if err := v.Delete(entry); err != nil {
		return err
	}

	if err := SaveVault(v, password, path); err != nil {
		return err
	}

	return nil
}

func cmdUpdate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: shh update <name>")
	}

	name := args[0]

	password, err := promptPassword("Master password: ")
	if err != nil {
		return err
	}

	path, err := VaultPath()
	if err != nil {
		return err
	}

	v, err := LoadVault(password, path)
	if err != nil {
		return err
	}

	entry, err := v.Get(name)
	if err != nil {
		return err
	}

	body, err := promptLine("New body: ")
	if err != nil {
		return err
	}

	entry.Body = body

	if err := v.Update(name, entry); err != nil {
		return err
	}

	if err := SaveVault(v, password, path); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: shh <add|get|list|delete|update> [args]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	var err error
	switch command {
	case "add":
		err = cmdAdd(args)
	case "get":
		err = cmdGet(args)
	case "list":
		err = cmdList(args)
	case "delete":
		err = cmdDelete(args)
	case "update":
		err = cmdUpdate(args)
	default:
		fmt.Println("unknown command:", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// defining components
type model struct {
	input     textinput.Model
	submitted string
	errMsg    string
	unlocked  bool
}

// initializing components
func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Master password"
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '#'
	ti.Focus()

	return model{input: ti}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			password := m.input.Value()
			m.input.SetValue("")
			path, _ := VaultPath()
			return m, tryUnlock(password, path)
		}
	case unlockResultMsg:
		if msg.err != nil {
			m.errMsg = msg.err.Error()
		} else {
			m.unlocked = true
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.unlocked {
		return "unlocked\n"
	}
	s := "Master password:\n\n" + m.input.View() + "\n\n(esc to quit)\n"
	if m.errMsg != "" {
		s += "\nerror: " + m.errMsg + "\n"
	}
	return s
}

type unlockResultMsg struct {
	vault *Vault
	key   []byte
	err   error
}

func tryUnlock(password string, path string) tea.Cmd {
	return func() tea.Msg {

		// using captured password to open vault
		pwBytes := []byte(password)
		v, err := LoadVault(pwBytes, path)
		if err != nil {
			return unlockResultMsg{err: err}
		}
		// reusing the saved salt
		key := DeriveKey(pwBytes, v.salt)
		return unlockResultMsg{vault: v, key: key}
	}
}

func tuiTEST() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
	}
}

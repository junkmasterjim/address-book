package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type addContactForm struct {
	fields  []string
	answers map[string]string
	done    bool
	menu    *menu
}

func (c addContactForm) Init() tea.Cmd {
	return nil
}

func (c addContactForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "esc":
			return c.menu, cmd
		}
	}
	return c, cmd
}

func (c addContactForm) View() string {
	s := strings.Join(c.fields, ", ")
	return s
}

func InitAddContactForm(m *menu) addContactForm {
	f := []string{
		"First name: ",
		"Last name: ",
		"Phone number: ",
		"Email: ",
	}
	return addContactForm{
		fields: f,
		menu:   m,
	}
}

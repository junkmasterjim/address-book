package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type editContactForm struct {
	fields  []string
	answers map[string]string
	done    bool
	table   *tableModel
}

func (c editContactForm) Init() tea.Cmd {
	return nil
}

func (c editContactForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "q", "esc":
			return c.table, cmd
		}
	}
	return c, cmd
}

func (c editContactForm) View() string {
	s := strings.Join(c.fields, ", ")
	return s
}

func InitEditContactForm(m *menu) editContactForm {
	f := []string{
		"First name: ",
		"Last name: ",
		"Phone number: ",
		"Email: ",
	}
	return editContactForm{
		fields: f,
		table:  &m.table,
	}
}

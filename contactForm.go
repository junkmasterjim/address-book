package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type contactForm struct {
	fields  []string
	answers map[string]string
	done    bool
	table   *tableModel
}

func (c contactForm) Init() tea.Cmd {
	return nil
}

func (c contactForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (c contactForm) View() string {
	s := strings.Join(c.fields, ", ")
	return s
}

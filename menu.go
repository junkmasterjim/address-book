package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type menu struct {
	choices   []string
	cursor    int
	contacts  []Contact
	selection string
	table     tableModel
}

func (m menu) Init() tea.Cmd {
	return func() tea.Msg {
		m.contacts = parseFile(DB_PATH)
		return "Success"
	}
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.contacts = parseFile(DB_PATH)
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.selection = msg.String()

		switch m.selection {
		case tea.KeyEnter.String():
			s := string(m.choices[m.cursor][0])
			switch s {
			case "q":
				return m, tea.Quit
			case "s":
				m.table = tableModel{
					table: buildTable(m.contacts),
					mode:  m.selection,
					menu:  &m,
				}
				return m.table, cmd
			}

		case tea.KeyDown.String(), "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case tea.KeyUp.String(), "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.choices) - 1
			}

		case tea.KeyCtrlC.String(), "q":
			return m, tea.Quit

		case "s":
			m.table = tableModel{
				table: buildTable(m.contacts),
				mode:  m.selection,
				menu:  &m,
			}
			return m.table, cmd

		case "a":
		}
	}

	return m, cmd
}

func (m menu) View() string {
	s := "menu\n\n"
	for i, c := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = CURSOR_CHAR
		}
		s += fmt.Sprintf("%s %s\n", cursor, c)
	}

	return s
}

func InitMainMenu() *menu {
	choices := []string{
		"s - show contacts",
		"a - add a contact",
		"q - quit",
	}
	mainMenu := menu{
		choices: choices,
	}
	return &mainMenu
}

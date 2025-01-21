package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableModel struct {
	table table.Model
	mode  string
	menu  *menu
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m.menu, cmd
		}

		switch m.mode {
		case "d":
		case "e":
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	switch m.mode {
	case "d":
		return "Select a contact to delete\n\n" +
			m.table.View() +
			"\n\n q - return to main menu" +
			"\n enter - select"

	case "e":
		return "Select a contact to edit\n\n" +
			m.table.View() +
			"\n\n q - return to main menu" +
			"\n enter - select"
	}

	return "Browsing contacts\n\n" +
		m.table.View() +
		"\n\n q - return to main menu"
}

func buildTable(contacts []Contact) table.Model {
	var fw, lw, pw, ew int
	var rows []table.Row

	for _, c := range contacts {
		if len(c.FirstName) > fw {
			fw = len(c.FirstName)
		}
		if len(c.LastName) > lw {
			lw = len(c.LastName)
		}
		if len(c.PhoneNumber) > pw {
			pw = len(c.PhoneNumber)
		}
		if len(c.Email) > ew {
			ew = len(c.Email)
		}

		rows = append(rows, table.Row{c.ID, c.FirstName, c.LastName, c.PhoneNumber, c.Email})
	}

	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "First", Width: fw},
		{Title: "Last", Width: lw},
		{Title: "Phone", Width: pw},
		{Title: "Email", Width: ew},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(contacts)+1),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

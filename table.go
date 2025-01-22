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
	form  editContactForm
}

func (t tableModel) Init() tea.Cmd {
	return nil
}

func (t tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		switch key {
		case tea.KeyEnter.String(), "q", "esc":
			return t.menu, cmd

		case "e":
			t.form = InitEditContactForm(t.menu)
			return t.form, cmd

		case "d":

		}
	}

	t.table, cmd = t.table.Update(msg)
	return t, cmd
}

func (t tableModel) View() string {
	return "Browsing contacts\n\n" +
		t.table.View() +
		"\n\nenter, q - return to main menu" +
		"\ne - edit contact" +
		"\nd - delete contact"
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

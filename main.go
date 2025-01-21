package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	DB_PATH     = "db.csv"
	CURSOR_CHAR = "=>"
)

type Contact struct {
	ID          string
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
}

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
			case "s", "e", "d":
				m.table = tableModel{buildTable(m.contacts), s, &m}
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

		case "s", "e", "d":
			m.table = tableModel{buildTable(m.contacts), m.selection, &m}
			return m.table, cmd
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

// Parse file at i, create file if it doesnt exist. Returns []Contact
func parseFile(path string) []Contact {
	file, err := os.Open(path)
	if err != nil {
		f, e := os.Create(path)
		if e != nil {
			log.Fatal(e)
		}
		// create file if not found
		f.Write([]byte("ID,First Name,Last Name,Phone Number,Email\n"))
	}
	s := bufio.NewScanner(file)
	lineNum := 0
	contacts := []Contact{}
	for s.Scan() {
		if lineNum != 0 {
			contact := strings.Split(s.Text(), ",")
			id := contact[0]
			fn := contact[1]
			ln := contact[2]
			ph := contact[3]
			em := contact[4]
			contacts = append(contacts, Contact{id, fn, ln, ph, em})
		}
		lineNum++
	}
	file.Close()
	return contacts
}

func initMainMenu() *menu {
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

func main() {
	menu := initMainMenu()

	if _, err := tea.NewProgram(menu, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("good bye")
}

/*
	needs work below
	needs work below
	needs work below
	needs work below
*/

// Appends string `toAppend` to file at string `path`.
func appendToFile(path string, toAppend string) {
	f, _ := os.Open(path)
	s := bufio.NewScanner(f)
	content := []string{}
	for s.Scan() {
		content = append(content, s.Text())
	}
	content = append(content, toAppend)

	fc := strings.Join(content, "\n")
	os.WriteFile(path, []byte(fc), 0o666)
	f.Close()
}

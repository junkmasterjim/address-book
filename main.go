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

func main() {
	menu := InitMainMenu()

	if _, err := tea.NewProgram(menu, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("good bye")
}

/* HELPER FUNCTIONS */

// Parse file at `path`, create file if it doesnt exist.
// Returns a slice of contacts.
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

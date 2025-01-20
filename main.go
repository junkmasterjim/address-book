package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const DB_PATH = "db.csv"

// Represents a contact in the address book.
type Contact struct {
  FirstName   string
  LastName    string
  PhoneNumber string
  Email       string
}

// Stores application state
type Model struct {
  choices   []string
  cursor    int
  contacts  []Contact
  state_map map[string]int
  state	    int	// 0 - main
		// 1 - add a contact
		// 2 - view contacts
		// 3 - delete a contact
		// 4 - edit a contact
}

// Returns an initial model for use in the program.
func initModel() Model {
  return Model{
    choices: []string{
      "a - add a contact", 
      "v - view contacts", 
      "d - delete a contact",
      "e - edit a contact",
      "q - quit",
    },
    cursor: 0,
    contacts: []Contact{},
    state_map: map[string]int{
      "": 0,
      "a": 1,
      "v": 2,
      "d": 3,
      "e": 4,
    },
    state: 0,
  }
}

// Init returns a `Cmd` which can perform some initial I/O. Returning `nil`
// ignores any I/O.
func (m Model) Init() tea.Cmd {
  return func() tea.Msg {
    m.contacts = parseFile(DB_PATH)
    return "Success"
  }
}

// Update is called when "things happen." Returns an updated model based on
// tea.Msg. Can also return a tea.Cmd to make more things happen.
//
// Typically, we check the type of message with a switch(msg), but type
// assertion works as well.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  m.contacts = parseFile(DB_PATH)
  var cmd tea.Cmd = nil

  switch msg := msg.(type){
  case tea.KeyMsg : // Key press
    switch msg.String() { // Which key was pressed?
    case "ctrl-c", "q": // Quit the application
      switch m.state {
      case 0: // Main menu
	return m, tea.Quit

      default: // Sub-menu. return to main menu
	m.state = 0
	return m, nil
      }

    case "up", "k": // Move the cursor up
      if m.cursor > 0 {
	m.cursor--
      } else {
	switch m.state {
	case 0:
	  m.cursor = len(m.choices) - 1
	case 2, 3, 4:
	  m.cursor = len(m.contacts) - 1
	}
      }

    case "down", "j": // Move the cursor down
      switch m.state {
      case 0: 
	if m.cursor < len(m.choices) - 1 {
	  m.cursor++
	} else {
	  m.cursor = 0
	}

      case 2, 3, 4:
	if m.cursor < len(m.contacts) - 1 {
	  m.cursor++
	} else {
	  m.cursor = 0
	}
      }

    case "enter": // Select an option
      switch m.state {
      case 0:
	opt := strings.Split(m.choices[m.cursor], "")[0]
	switch opt {
	case "a", "v", "d", "e":
	  m.state = m.state_map[opt]
	case "q":
	  cmd = tea.Quit
	}

      case 2, 3, 4:

      }

    m.cursor = 0
    }
  }
  return m, cmd
}

// Renders our UI to the terminal. Returns a string representing the entirity of
// the UI.
func (m Model) View() string {
  var s string

  switch m.state {
  case 0: // Print the main menu
    // Initial string header
    s = "--- Main Menu ---\n\n"
    
    // Iterate over our choices
    for i, choice := range m.choices {
  
      // Is the cursor pointing at this choice?
      cursor := " " // no cursor
      if m.cursor == i {
        cursor = ">" // cursor!
      }
  
      // Render the row
      s += fmt.Sprintf("%s %s\n", cursor, choice)
    }
  
    
  case 2, 3, 4: // Select a contact
    switch m.state {
    case 2:
      s = "--- All Contacts ---\n\n"
    case 3:
      s = "--- Delete a Contact ---\n\n"
    case 4:
      s = "--- Edit a Contact ---\n\n"
    }

    for i, contact := range m.contacts {
      cursor := " "
      if m.cursor == i {
	cursor = ">"
      }

      s += fmt.Sprintf("%s %s, %s\n", cursor, contact.LastName, contact.FirstName)
    }
  }

  if m.state != 0 {
    s += "\npress q to return to menu"
  }

  return s
}


func main(){
  p := tea.NewProgram(initModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("There was an error: %v", err)
    os.Exit(1)
  }
}

// --- Helper Functions ------------

// Parse file at i, create file if it doesnt exist. Returns []Contact
func parseFile(i string) []Contact {
  file, err := os.Open(i)
  if err != nil{
    f, e := os.Create(i)
    if e != nil {
    	log.Fatal(e)
    }
    // create file if not found
    f.Write([]byte("First Name,Last Name,Phone Number,Email\n"))
  }
  s := bufio.NewScanner(file)
  lineNum := 0
  contacts := []Contact{}
  for s.Scan(){
   if lineNum != 0 {
      contact := strings.Split(s.Text(), ",")
      fn := contact[0]
      ln := contact[1]
      ph := contact[2]
      em := contact[3]
      contacts = append(contacts, Contact{fn, ln, ph, em})
    } 
    lineNum++
  }
  file.Close()
  return contacts
}

// Appends string `toAppend` to file at string `path`.
func appendToFile(path string, toAppend string){
  f, _ := os.Open(path)
  s := bufio.NewScanner(f)
  content := []string{}
  for s.Scan(){
    content = append(content, s.Text())
  }
  content = append(content, toAppend)
  
  fc := strings.Join(content, "\n")
  os.WriteFile(path, []byte(fc),  0666)
  f.Close()
}

// returns a new contact based on user input
// i can clean this up using a different logic flow
func newContact() *Contact{
  r:= bufio.NewReader(os.Stdin)
 
  fmt.Print("First name: ")
  first, _, _:= r.ReadLine()
  fn := string(first)

  fmt.Print("Last name: ")
  last, _, _ := r.ReadLine()
  ln := string(last)

  fmt.Print("Phone number: ")
  phone, _, _ := r.ReadLine()
  ph := string(phone)

  fmt.Print("Email: ")
  email, _, _ := r.ReadLine()
  em := string(email)
  
  return &Contact{fn, ln, ph, em}
}



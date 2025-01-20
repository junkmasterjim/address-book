package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const DB_PATH = "db.csv"
const CURSOR_CHAR = "=>"

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
  input	    textinput.Model
  smap	    map[string]int
  state	    int	// 0 - main
		// 1 - add a contact
		// 2 - view contacts
		// 3 - delete a contact
		// 4 - edit a contact
}

// Returns an initial model for use in the program.
func initModel() *Model {
  return &Model{
    choices: []string{
      "a - add a contact", 
      "s - show contacts", 
      "d - delete a contact",
      "e - edit a contact",
      "q - quit",
    },
    cursor: 0,
    contacts: []Contact{},
    smap: map[string]int{
      "": 0,
      "a": 1,
      "s": 2,
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
    switch m.state { // Check menu state
    case 0: // Main menu
      return m.handleMainMenu(msg)

    case 1:
      return m.handleAddContact(msg)
      // return m.handleSubMenu(msg)

    case 2, 3, 4:
      return m.handleSubMenu(msg)
    }
  }
  return m, cmd
}

// Renders our UI to the terminal. Returns a string representing the entirity of
// the UI.
func (m Model) View() string {
  var s string

  switch m.state {
  case m.smap[""]: // Print the main menu
    // Initial string header
    s = "--- Main Menu ---\n\n"
    // Iterate over our choices
    for i, choice := range m.choices {
      // Is the cursor pointing at this choice?
      cursor := " " // no cursor
      if m.cursor == i {
        cursor = CURSOR_CHAR // cursor!
      }
      // Render the row
      s += fmt.Sprintf("%s %s\n", cursor, choice)
    }
  
    
  default: // anything but the main menu state
    switch m.state {
    case m.smap["a"]:
      return fmt.Sprintf("Add a new contact\n%s\n%s", m.input.View(), "[esc to cancel]")
      return "'add contact' not ready yet\npress q to return to menu\n\n"
    case m.smap["s"]:
      s = "--- All Contacts ---\n\n"
    case m.smap["d"]:
      return "'delete contact 'not ready yet\npress q to return to menu\n\n"
      s = "--- Delete a Contact ---\n\n"
    case m.smap["e"]:
      return "'edit contact' not ready yet\npress q to return to menu\n\n"
      s = "--- Edit a Contact ---\n\n"
    }


    for i, contact := range m.contacts {
      cursor := " "
      if m.cursor == i {
	cursor = CURSOR_CHAR
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
  p := tea.NewProgram(initModel(), tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("There was an error: %v", err)
    os.Exit(1)
  }
}

// --- Model Helpers ---------------
func (m *Model) reset() {
  m.state = 0
  m.cursor = 0
}

func (m Model) handleAddContact(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd = nil
  key := msg.String()

  m.input = textinput.New()
  m.input.Placeholder = "First Name"
  m.input.Focus()
  m.input.CharLimit = 15
  m.input.Width = 20
  
  switch key {
  case "esc":
    m.reset()
  }
  
  m.input, cmd = m.input.Update(msg)
  return m, cmd
}

func (m Model) handleMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
  key := msg.String()
  var cmd tea.Cmd = nil

  switch key {
  case "ctrl-c", "q":
    return m, tea.Quit

  case "a":
    m.state = m.smap[key]
  case "s":
    m.state = m.smap[key]
  case "d":
    m.state = m.smap[key]
  case "e":
    m.state = m.smap[key]
  
  case "j", "down":
    if m.cursor < len(m.choices) - 1 {
      m.cursor++
    } else {
      m.cursor = 0
    }
    
  case "k", "up":
    if m.cursor > 0 {
      m.cursor--
    } else{
      m.cursor = len(m.choices) - 1
    }
  

  case "enter":
    opt := strings.Split(m.choices[m.cursor], "")[0]
    switch opt {
      case "a", "s", "d", "e":
	m.state = m.smap[opt]
	m.cursor = 0
      case "q":
	return m, tea.Quit
    }
  }

  return m, cmd
}

func (m Model) handleSubMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd){
  key := msg.String()
  switch key{
  case "ctrl-c":
    return m, tea.Quit

  case "q":
    m.state = 0
    m.cursor = 0
    return m, nil
  
  case "j", "down":
    if m.cursor < len(m.contacts) - 1 {
      m.cursor++
    } else {
      m.cursor = 0
    }
  
  case "k", "up":
    if m.cursor > 0 {
      m.cursor--
    } else {
      m.cursor = len(m.contacts) - 1
    }

  case "enter":
    switch m.state{
    case m.smap["d"]: // delete contact
      // remove contact from file and the model will update on parse
    case m.smap["e"]: // edit contact
      // edit contact in file and the model will update on parse
    }
  }

  return m, nil
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


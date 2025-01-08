package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"golang.org/x/term"
)

type Contact struct {
  FirstName   string
  LastName    string
  PhoneNumber string
  Email       string
}

const DB_PATH = "db.csv"
const MENU_STRING = "-------Main Menu-------\r\n" +
		    "(a)dd a contact\r\n" +
		    "(s)how all contacts\r\n" +
		    "(d)elete a contact\r\n" +
		    "(e)dit a contact\r\n" +
		    "(q)uit\r\n" +
		    "------------------------\r"

func main(){
  contacts := parseFile(DB_PATH)
  for {
    switch getCharInput(MENU_STRING){
    case "q":
      fmt.Print("Good bye\r\n\n")
      os.Exit(0)
    case "a":
      contacts = append(contacts, newContact())
      c := contacts[len(contacts)-1]
      contactString := fmt.Sprintf("%v,%v,%v,%v\n", c.FirstName, c.LastName, c.PhoneNumber, c.Email)
      appendToFile(DB_PATH, contactString)
    case "s":
      printTable(contacts)
    case "d":
      break
    case "e":
      break
    default:
      fmt.Println("Please select one of the menu options\n\r")
    }
  }
}

// parse DB file and create if it doesnt exist
func parseFile(input string) ([]Contact){
  file, err := os.Open(input)
  if err != nil{
    f, e := os.Create(input)
    if e != nil {
    	log.Fatal(e)  // this is just insane
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

// appends string to file
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

// set terminal to raw mode and reads single keypress input from a user
// prompts the user with the prompt parameter
func getCharInput(prompt string) string{
  oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
  if err != nil {
    log.Fatal(err)
  }
  defer term.Restore(int(os.Stdin.Fd()), oldState)

  fmt.Println(prompt)
  
  r := bufio.NewReader(os.Stdin)
  charPressed, _,err := r.ReadRune()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Println()
  return string(charPressed)
}

// returns a new contact based on user input
// i can clean this up using a different flow logic
func newContact() Contact{
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
  
  return Contact{fn, ln, ph, em}
}

func printTable(c []Contact){
  w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
  fmt.Fprintln(w, "First Name\tLast Name\tPhone Number\tEmail\r")
  for i:= range c {
    fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\r\n", c[i].FirstName, c[i].LastName, c[i].PhoneNumber, c[i].Email)
  }
  fmt.Fprintln(w)
  w.Flush()
}

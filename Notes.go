package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Note is a struct that represents a single note, with a title and body.
type Note struct {
	Title string
	Body  string
}

func main() {
	// Create a map to store the notes
	notes := make(map[string]Note)

	// Read input from the command line
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command (add, view, remove, or quit): ")
		scanner.Scan()
		input := scanner.Text()

		// Check the command and perform the appropriate action
		switch strings.ToLower(input) {
		case "add":
			// Add a new note
			fmt.Print("Enter the title of the note: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter the body of the note: ")
			scanner.Scan()
			body := scanner.Text()
			notes[title] = Note{Title: title, Body: body}
			fmt.Println("Note added.")
		case "view":
			// View all notes or a specific note
			fmt.Print("Enter the title of the note (leave blank to view all notes): ")
			scanner.Scan()
			title := scanner.Text()
			if title == "" {
				// View all notes
				fmt.Println("All notes:")
				for _, note := range notes {
					fmt.Println(note.Title)
				}
			} else {
				// View a specific note
				note, ok := notes[title]
				if !ok {
					fmt.Println("Note not found.")
				} else {
					fmt.Println(note.Title)
					fmt.Println(note.Body)
				}
			}
		case "remove":
			// Remove a note
			fmt.Print("Enter the title of the note: ")
			scanner.Scan()
			title := scanner.Text()
			if _, ok := notes[title]; !ok {
				fmt.Println("Note not found.")
			} else {
				delete(notes, title)
				fmt.Println("Note removed.")
			}
		case "quit":
			// Quit the program
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println

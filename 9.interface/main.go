package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface

type saver interface {
	Save() error
}

// type displayer interface{
// 	Display()
// }

// interface embedding
type outputtable interface {
	saver
	// displayer
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.1)
	printSomething("strom")
	title, content := getNoteData()
	todoText := getUserInput("todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
	outputData(userNote)

}

// this function will accept any type value string int etc
func printSomething(value interface{}) {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("integer: ", typedVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("float64: ", floatVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("integer: ", value)
	// case float64:
	// 	fmt.Println("float: ", value)
	// case string:
	// 	fmt.Println("string: ", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note successful")

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}

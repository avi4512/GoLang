package main

import (
	"booksProject/note"
	"booksProject/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	SaveTOjson() error
}

func main() {

	title, content := getNoteData()
	text := getData("Todo Text:")

	userTodo, err := todo.New(text)

	if err != nil {
		fmt.Print(err, " Something Wrong!!!")
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err, " Something Wrong!!!")
		return
	}

	userNote.Output()
	err = saveData(userNote)

	if err != nil {
		fmt.Println("Error in saveData..")
		fmt.Println(err)
	}

	userTodo.Output()
	err = saveData(userTodo)

	if err != nil {
		fmt.Println("Error in saveData..")
		fmt.Println(err)
	}

}

func saveData(data saver) error {

	err := data.SaveTOjson()

	if err != nil {
		fmt.Println("File Not Saved:")
		return err
	}

	note_path, err := os.Getwd()

	if err != nil {
		fmt.Println("File Not Saved")
		return err
	}

	fmt.Printf("File Saved Successfully..at %s\n", note_path)
	return nil

}

func getNoteData() (string, string) {

	title := getData("Enter the tile:")
	Content := getData("Enter the Content: ")
	return title, Content

}
func getData(text string) string {
	fmt.Print(text)
	read := bufio.NewReader(os.Stdin)
	text, err := read.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	if err != nil {
		fmt.Println(err)
	}
	return text
}

package main

import (
	"booksProject/note"
	"booksProject/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()
	text := getData("Todo Text:")

	userTodo, err := todo.New(text)

	if err != nil {
		fmt.Print(err, " Something Wrong!!!")
		return
	}

	userTodo.Output()
	err = userTodo.SaveTOjson()

	if err != nil {
		fmt.Println("File Not Saved:")
		return
	}

	todo_path, _ := os.Getwd()

	fmt.Printf("File Saved Successfully.. at %s\n", todo_path)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err, " Something Wrong!!!")
		return
	}

	userNote.Output()
	err = userNote.SaveTOjson()

	if err != nil {
		fmt.Println("File Not Saved:")
		return
	}

	note_path, err := os.Getwd()

	if err != nil {
		fmt.Println("File Not Saved")
		return
	}

	fmt.Printf("File Saved Successfully..at %s\n", note_path)

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

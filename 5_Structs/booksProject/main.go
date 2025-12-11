package main

import (
	"booksProject/note"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()
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

	path, err := os.Getwd()

	if err != nil {
		fmt.Println("File Not Saved")
		return
	}

	fmt.Printf("File Saved Successfully..at %s", path)

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

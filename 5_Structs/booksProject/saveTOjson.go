package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on"`
}

func New(title, contenet string) note {

	return note{
		Title:     title,
		Content:   contenet,
		CreatedOn: time.Now(),
	}

}

func (n note) display() {
	fmt.Println("Title is:", n.Title)
	fmt.Println("Content is:", n.Content)
}

func (n note) save() error {

	fileName := "xyz.json"

	json_byte, err := json.Marshal(n)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json_byte, 0644)

	return nil

}

func main() {

	t := getUserData("Enter the Title: ")
	c := getUserData("Enter the Content: ")

	appNote := New(t, c)

	appNote.display()
	err := appNote.save()

	if err != nil {
		fmt.Println("File Not Saved!!!!!!!!")
	}

	path, err := os.Getwd()

	if err != nil {
		fmt.Println("File Not Saved!!!!!!!!")
	}

	fmt.Printf("File Saved Successfully... at %v", path)

}

func getUserData(text string) string {
	fmt.Print(text)

	read := bufio.NewReader(os.Stdin)
	data, err := read.ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
	data = strings.TrimSuffix(data, "\r")

	if err != nil {
		fmt.Println(err)
	}

	return data
}

package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string
	Content  string
	CreateOn time.Time
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("enter all details")
	}
	return Note{
		Title:    title,
		Content:  content,
		CreateOn: time.Now(),
	}, nil

}

func (n Note) SaveTOjson() error {

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	json_file, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json_file, 0644)

}

func (n Note) Output() {
	//fmt.Printf("Title of the Book is: %s \nContent is: %s", n.title, n.content)
	fmt.Println("Tiltel of the Book:", n.Title)
	fmt.Println("Content is:", n.Content)

}

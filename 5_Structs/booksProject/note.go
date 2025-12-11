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
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateOn time.Time `json:"created_on"`
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
	fileName = strings.ToLower(fileName) + ".json"

	json_byte, err := json.Marshal(n)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json_byte, 0644)

	return nil

}

func (n Note) Output() {
	//fmt.Printf("Title of the Book is: %s \nContent is: %s", n.title, n.content)
	fmt.Println("Tiltel of the Book:", n.Title)
	fmt.Println("Content is:", n.Content)

}

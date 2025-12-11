package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("enter all details")
	}
	return Todo{
		Text: content,
	}, nil

}

func (todo Todo) SaveTOjson() error {

	fileName := "todo.json"

	json_byte, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json_byte, 0644)

	return nil

}

func (todo Todo) Output() {

	fmt.Println("Text is:", todo.Text)

}

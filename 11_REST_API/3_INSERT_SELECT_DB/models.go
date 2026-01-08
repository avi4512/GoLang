package models

import (
	"apis/db"
	"fmt"
	"time"
)

type Events struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserId      int       `json:"userId"`
}

var events = []Events{}

func (e Events) Save() error {

	query := `INSERT INTO event(name,description,location,dateTime,user_id)
	VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("invalid query!")
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	id, err := result.LastInsertId()

	e.Id = id
	return err

}

func GetAllEvents() ([]Events, error) {

	query := "SELECT * FROM event"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Events
	for rows.Next() {

		var e Events
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}


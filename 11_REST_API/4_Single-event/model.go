package models

import (
	"fmt"
	"rest/db"
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

	query := `INSERT INTO event (name, description, location, dateTime, user_id)
	VALUES(?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("Insert Query is Wrong!!")
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		fmt.Println(err)
		return err
	}

	id, err := result.LastInsertId()

	e.Id = id
	return err
}

func GetAllEvents() ([]Events, error) {

	query := "SELECT * FROM event"

	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var events []Events
	for rows.Next() {

		var e Events
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil
}

func GetEventFromId(id int64) (Events, error) {

	query := "SELECT * FROM event WHERE id = ?"

	var e Events
	row := db.DB.QueryRow(query, id)

	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

	if err != nil {
		fmt.Println(err)
		return Events{}, nil
	}
	return e, nil

}

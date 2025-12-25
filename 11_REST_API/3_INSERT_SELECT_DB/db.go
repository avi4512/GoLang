package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("dataBase Not Connected!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTable()

}

func CreateTable() {

	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME,
    user_id INTEGER
)
	`

	_, err := DB.Exec(createEventTable)

	if err != nil {
		fmt.Println(err)
		panic("Table is not Created!")
	}
}

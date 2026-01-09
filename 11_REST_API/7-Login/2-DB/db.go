package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("sqlite3", "eventsAPI.db")

	if err != nil {
		fmt.Println("DataBase Not Established!!")
		fmt.Println(err)
		panic("Somthing wrong in BD establishment..!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTable()
}

func createTable() {

	createUserTable := `
						CREATE TABLE IF NOT EXISTS users (
    					id INTEGER PRIMARY KEY AUTOINCREMENT,
						email TEXT UNIQUE,
						password TEXT NOT NULL)

						`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("User Table NOT Created!")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME,
    user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		fmt.Println("Table NOT created!!!")
		fmt.Println(err)
		return
	}

}

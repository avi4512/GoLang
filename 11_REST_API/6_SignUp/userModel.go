package models

import (
	"fmt"
	"rest/db"
	"rest/hashing"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := "INSERT INTO users(email,password) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hasedPassword, err := hashing.HashingPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hasedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.Id = id
	return err
}

func GetAllUsers() ([]User, error) {

	query := "SELECT * FROM users"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {

		var u User

		err := rows.Scan(&u.Id, &u.Email, &u.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	if len(users) == 0 {
		fmt.Println("No users found in DB")
	}
	return users, nil
}

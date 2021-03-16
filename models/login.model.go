package models

import (
	"AICLoginService/db"
	"AICLoginService/helpers"
	"database/sql"
	"fmt"
)

type Users struct {
	Email    int    `json:"email"`
	Password string `json:"password"`
}

func CheckLogin(Email, password string) (bool, error) {

	var obj Users
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE email = ?"

	err := con.QueryRow(sqlStatement, Email).Scan(
		&obj.Email, &obj.Password, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email Not Found")
		return false, err
	}
	if err != nil {
		fmt.Println("Querry Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match")
		return false, err
	}
	return true, nil
}

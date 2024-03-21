package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type infomation struct {
	UserName string `json:"Username"`
	Password string `json:"Password"`
	FullName string `json:"Fullname"`
	Email    string `json:"Email"`
	CreateAt string `json:"CreateAt"`
	UpdateAt string `json:"UpdateAt"`
}

func sign_Up(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	hashedPassword := encode_data(Account.Email, Account.Password, 2)

	db, err := sql.Open("mysql", "root:")
	check_err(err)

	// add account into database
	_, err = db.Exec("INSERT INTO account(email, password) VALUES(?, ?)", Account.Email, hashedPassword)
	check_err(err)
	defer db.Close()

	check := true
	JsonData, err := json.Marshal(check)
	check_err(err)

	return string(JsonData)
}

func sign_In(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	hashPassword := encode_data(Account.Password, Account.Password, 2)

	// check account from database
	check := check_login(Account.Email, hashPassword)
	JsonData, err := json.Marshal(check)
	check_err(err)

	return string(JsonData)
}

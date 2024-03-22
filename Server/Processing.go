package main

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type quantity struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
	RoleID int `json:"roleID"`
}

type infomation struct {
	Name string `json:"name"`
	Password string `json:"Password"`
	FullName string `json:"Fullname"`
	Gender 	 string `json:"Gender"`
	Email    string `json:"Email"`
	CreateAt string `json:"CreateAt"`
	UpdateAt string `json:"UpdateAt"`
}

func create_account(jsonData []byte) string {
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

func get_Account(jsonData []byte) string {
	var object quantity
	err := json.Unmarshal(jsonData, &object)
	check_err(err)

	list := get_account(object.RoleID, object.Start, object.Limit)
	JsonData, err := json.Marshal(list)
	check_err(err)

	return string(JsonData)
}
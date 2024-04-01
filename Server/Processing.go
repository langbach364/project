package main

import (
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type quantity struct {
	Start  int `json:"start"`
	Limit  int `json:"limit"`
	RoleID int `json:"roleID"`
}

type CheckRole struct {
	ID_USER  int `json:"IdUser"`
	ID_CHECK int `json:"IdCheck"`
}

type infomation struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}


func create_Account(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	check := create_account(Account.Email, Account.Password)
	JsonData, err := json.Marshal(check)
	check_err(err)

	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	
	var Infomation infomation
	_, err = db.Exec("INSERT INTO Informations(id, name, password, fullname, gender, email, createAt, updateAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		Infomation.Name, Infomation.Password, Infomation.FullName, Infomation.Gender, Infomation.Email, Infomation.CreateAt, Infomation.UpdateAt)
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

func check_Role(jsonData []byte) string {
	var check CheckRole
	err := json.Unmarshal(jsonData, &check)
	check_err(err)

	result := check_role(check.ID_USER, check.ID_CHECK)
	JsonData, err := json.Marshal(result)
	check_err(err)

	return string(JsonData)
}

package main

import (
	
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func check_err(err error) {
	if err != nil {
		println(err)
		log.Fatal(err)
	}
}

type Account struct {
	Email    string
	Password string
}


type List_Account struct {
	Name     string
	Password string
	FullName string
	Gender   string
	Email    string
	CreateAt time.Time
	UpdateAt time.Time
}



func check_role(id_user int, id_check int) bool {
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(192.168.108.129)/Manager")
	check_err(err)
	defer db.Close()

	var role_user int
	err = db.QueryRow("SELECT RoleID from Role WHERE id = ?", id_user).Scan(&role_user)
	check_err(err)

	var role_check int
	err = db.QueryRow("SELECT RoleID from Role WHERE id = ?", id_check).Scan(&role_check)
	check_err(err)

	return role_user < role_check
}

func create_account(email string, password string) bool {
	var Account Account
	Account.Email = email
	Account.Password = password

	hashedPassword := encode_data(Account.Email, Account.Password, 2)

	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(192.168.108.129)/Manager")
	check_err(err)
	defer db.Close()

	var id *int
	err = db.QueryRow("SELECT MAX(id) FROM Account;").Scan(&id)
	check_err(err)

	newId := 1
	if id != nil {
		newId = *id + 1
	}

	_, err = db.Exec("INSERT INTO Account(id, email, password) VALUES(?, ?, ?)", newId, Account.Email, hashedPassword)
	check_err(err)

	defer db.Close()
	return true
}

func get_account(roleID int, start int, limit int) []List_Account {
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(192.168.108.129)/Manage")
	check_err(err)
	defer db.Close()

	result := make([]List_Account, 0)
	query1 := "SELECT Informations.name, Informations.fullname, Informations.gender, Informations.email, Informations.CreateAt, Informations.UpdateAt"
	query2 := " from Informations join Role ON Informations.id = Role.id WHERE Role.roleID = ? LIMIT ? OFFSET ?"
	rows, err := db.Query(query1+query2, roleID, limit, start-1)
	check_err(err)
	defer rows.Close()

	for rows.Next() {
		var account List_Account
		err := rows.Scan(&account.Name, &account.FullName, &account.Gender, &account.Email, &account.CreateAt, &account.UpdateAt)
		check_err(err)
		result = append(result, account)
	}

	return result
}

func change_infomation(id_user int, info infomation) bool {
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(192.168.108.129)/Manager")
	check_err(err)
	defer db.Close()

	_, err = db.Exec("UPDATE Users SET Name = ?, Password = ?, FullName = ?, Gender = ?, Email = ?, UpdateAt = ? WHERE id = ?",
		info.Name, info.Password, info.FullName, info.Gender, info.Email, info.UpdateAt, id_user)
	check_err(err)
	return true
}

func update_infomation(id_user int, info infomation) bool {
	db, err := sql.Open("mysql", "root")
	check_err(err)
	defer db.Close()
	var check bool = false
	check = change_infomation(id_user, info)
	return check
}

func delete_account(id_user int, id_delete int) bool {
	db, err := sql.Open("mysql", "root:@ztegc4DF9F4E@tcp(192.168.108.129)/Manager")
	check_err(err)
	defer db.Close()

	if check_role(id_user, id_delete) {
		_, err = db.Exec("DELETE from Account WHERE id = ?", id_delete)
		check_err(err)
		return true
	}
	return false
}

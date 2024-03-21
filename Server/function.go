package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func check_err(err error) {
	if err != nil {
		println(err)
		log.Fatal(err)
	}
}

type List_Account struct {
	Name     string
	Password string
	FullName string
	Gender   string
	Email    string
	CreateAt string
	UpdateAt string
}

func check_login(email string, Password string) bool {
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	defer db.Close()

	var storedPassword string
	err = db.QueryRow("SELECT password FROM Account WHERE email = ?", email).Scan(&storedPassword)

	switch {
	case err == sql.ErrNoRows:
		return false

	case err != nil:
		check_err(err)

	default:
		if Password != storedPassword {
			return false
		}
	}
	return true
}

func check_role(id_user int, id_delete int) bool {
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	defer db.Close()

	var role_user int
	err = db.QueryRow("SELECT RoleID from Role WHERE id = ?", id_user).Scan(&role_user)
	check_err(err)

	var role_delete int
	err = db.QueryRow("SELECT RoleID from Role WHERE id = ?", id_delete).Scan(&role_delete)
	check_err(err)

	return role_user < role_delete
}

func change_infomation(id_user int, info infomation) bool {
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
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
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	defer db.Close()

	if check_role(id_user, id_delete) {
		_, err = db.Exec("DELETE from Account WHERE id = ?", id_delete)
		check_err(err)
		return true
	}
	return false
}

func get_account(sum_account int) []List_Account {
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	defer db.Close()

	var result []List_Account

	for i := 0; i < sum_account; i++ {
		err = db.QueryRow("SELECT name, fullname, gender, email, CreateAt, UpdateAt from Informations").Scan(&result[i])
		check_err(err)
	}
	return result
}

func list_account(char string, start int, limit int) [][]List_Account {
	db, err := sql.Open("mysql", "user:@ztegc4DF9F4E@tcp(localhost)/Manager")
	check_err(err)
	defer db.Close()

	var List_role [][]List_Account
	var sum_role int
	err = db.QueryRow("SELECT SUM(roleID) from Role").Scan(&sum_role)
	var sum_account int
	err = db.QueryRow("SELECT SUM(id) from Account").Scan(&sum_account)

	for i := 0 ; i < sum_role ; i++ {
		List_role[i] = get_account(sum_account)
	}
	return List_role
}

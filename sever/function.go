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
	UserName string
	Password string
	FullName string
	Email    string
	CreateAt string
	UpdateAt string
}

func check_login(email string, Password string) bool {
	db, err := sql.Open("mysql", "root:")
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
	db, err := sql.Open("mysql", "root")
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
	db, err := sql.Open("mysql", "root")
	check_err(err)
	defer db.Close()

	_, err = db.Exec("UPDATE Users SET Username = ?, Password = ?, FullName = ?, Email = ?, UpdateAt = ? WHERE id = ?",
		info.UserName, info.Password, info.FullName, info.Email, info.UpdateAt, id_user)
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
	db, err := sql.Open("mysql", "root")
	check_err(err)
	defer db.Close()

	if check_role(id_user, id_delete) {
		_, err = db.Exec("DELETE from Account WHERE id = ?", id_delete)
		check_err(err)
		return true
	}
	return false
}

func list_account(char string, start int, limit int) []List_Account {
	db, err := sql.Open("mysql", "root")
	check_err(err)
	defer db.Close()

	var List []List_Account

	if char != "ALL" {
		rows, err := db.Query("SELECT UserName, FROM Users WHERE LEFT(UserName, 1) LIKE ? LIMIT ? OFFSET ?", char, limit, start)
		check_err(err)
		for rows.Next() {
			var p List_Account
			rows.Scan(&p.UserName)
			List = append(List, p)
		}
		return List

	} else {
		rows, err := db.Query("SELECT UserName, FROM Users LIMIT ? OFFSET ?", char, limit, start)
		check_err(err)
		for rows.Next() {
			var p List_Account
			rows.Scan(&p.UserName)
			List = append(List, p)
		}
		return List
	}

}

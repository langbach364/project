package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type infomation struct {
	Id       int
	Name     string
	Password string
	FullName string
	Gender   string
	Email    string
	CreateAt time.Time
	UpdateAt time.Time
}

type Account struct {
	Email    string
	Password string
}

func check_err(err error) {
	if err != nil {
		println(err)
		log.Fatal(err)
	}
}

func create_account(email string, password string) bool {
	var Account Account
	Account.Email = email
	Account.Password = password

	hashedPassword := encode_data(Account.Email, Account.Password, 2)

	db, err := sql.Open("mysql", "id21200239_bachlang364:@ztegc4DF9F4E@botmiw59tjefgodpkrik-mysql.services.clever-cloud.com/id21200239_intern")
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

	info := infomation{
		Name:     "username",
		FullName: "Full Name",
		Gender:   "Male",
		Email:    "email@example.com",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	info.Id = newId

	_, err = db.Exec("INSERT INTO Informations(id, name, password, fullname, gender, email, createAt, updateAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?)",
		info.Id, info.Name, info.Password, info.FullName, info.Gender, info.Email, info.CreateAt, info.UpdateAt)
	check_err(err)
	defer db.Close()
	return true
}

func main() {
	check := create_account("langbach364", "123456")
	println(check)
}

package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

type table struct {
	TableName string `json:"table"`
}

type query struct {
	Query string `json:"query"`
}

type object struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
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


func get_Table_Columns(tableName string) []string {
	db, err := Connect()
	check_err(err)
	defer db.DB.Close()

	query := fmt.Sprintf("SELECT * FROM %s LIMIT 0", tableName)
	rows, err := db.DB.Query(query)
	check_err(err)
	defer rows.Close()
	columns, err := rows.Columns()
	check_err(err)
	return columns
}

func queryNotAllowed() map[string]bool {
	Attricbute := map[string]bool{
		"password": true,
	}
	return Attricbute
}

func queryAllowed(tableName string) map[string]bool {
	columns := get_Table_Columns(tableName)
	attribute_Allowed := make(map[string]bool)
	attribute_No_Allowed := queryNotAllowed()

	for _, column := range columns {
		switch attribute_No_Allowed[column] {
		case true:
			attribute_Allowed[column] = false
		default:
			attribute_Allowed[column] = true
		}
	}
	return attribute_Allowed
}

func check_allowed(attribute_No_Allowed map[string]bool, column string) bool {
	return !attribute_No_Allowed[column]
}

func select_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var name table
				err = json.Unmarshal(body, &name)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				rows, err := dbInfo.DB.Query(string(query.Query))
				check_err(err)
				defer rows.Close()

				columns, err := rows.Columns()
				check_err(err)

				Attricbute := queryAllowed(name.TableName)
				var check_no_allowed bool
				for _, column := range columns {
					check_no_allowed = check_allowed(Attricbute, column)
				}

				if !check_no_allowed {
					result := make([]map[string]interface{}, 0)
					values := make([]interface{}, len(columns))
					pointers := make([]interface{}, len(columns))

					for i := range values {
						pointers[i] = &values[i]
					}

					for rows.Next() {
						err := rows.Scan(pointers...)
						check_err(err)

						row := make(map[string]interface{})
						for i, colName := range columns {
							value := check_string(values[i])
							row[colName] = value
						}

						result = append(result, row)
					}
					json.NewEncoder(w).Encode(result)
				} else {
					json.NewEncoder(w).Encode("Not Allowed")
				}
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

func check_passowrd(id int, password string) bool {
	db, err := Connect()
	check_err(err)

	var password_object string
	err = db.DB.QueryRow("SELECT password FROM Account WHERE id = ?", id).Scan(&password_object)
	check_err(err)

	db.DB.Close()

	return password_object == password
}

func get_email(id int) string {
	db, err := Connect()
	var email string

	db.DB.QueryRow("SELECT email FROM Account WHERE id = ?", id).Scan(&email)
	check_err(err)
	db.DB.Close()

	return email
}

func delete_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				var Object object
				err = json.Unmarshal(body, &Object)
				check_err(err)

				Object.Password = encode_data(get_email(Object.ID), Object.Password, 2)

				if check_passowrd(Object.ID, Object.Password) {
					dbInfo.DB.Exec("DELETE FROM Account WHERE id = ?", Object.ID)
					json.NewEncoder(w).Encode("True")
				} else {
					json.NewEncoder(w).Encode("Password or Id is not correct")
				}
			}
		default:
			http.Error(w, "Method is not used", http.StatusMethodNotAllowed)
		}
	}
}

GRANT SELECT ON employees TO user1;
REVOKE SELECT ON employees FROM user1;


func update_Handler(dbInfo *DBInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := io.ReadAll(r.Body)
				check_err(err)

				var query query
				err = json.Unmarshal(body, &query)
				check_err(err)

				var Object object
				err = json.Unmarshal(body, &Object)
				check_err(err)

				Object.Password = encode_data(get_email(Object.ID), Object.Password, 2)

				if check_passowrd(Object.ID, Object.Password) {
					dbInfo.DB.Exec(string(query.Query), Object.ID, Object.Password)
					json.NewEncoder(w).Encode("True")
				} else {
					json.NewEncoder(w).Encode("Password or Id is not correct")
				}
			}
		}
	}
}

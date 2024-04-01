package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"crypto/rand"
	"encoding/hex"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

type DBInfo struct {
	DB *sql.DB
}

type account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func enable_middleware_cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors := cors.New(cors.Options{
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
			AllowedMethods:   []string{"POST"},
			AllowedOrigins:   []string{"http://127.0.0.1:5500"},
			AllowCredentials: true,
			Debug:            true,
		})
		Cors.ServeHTTP(w, r, next.ServeHTTP)
	})
}

func Connect() (*DBInfo, error) {
	connStr := "root:@ztegc4DF9F4E@tcp(localhost:3306)/Manager"
	db, err := sql.Open("mysql", connStr)
	check_err(err)
	return &DBInfo{DB: db}, nil
}

func check_string(str interface{}) interface{} {
	switch v := str.(type) {
	case []byte:
		return string(v)
	case string:
		return v
	default:
		return str
	}
}

func randomToken() string {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	check_err(err)
	return hex.EncodeToString(bytes)
}


func create_cookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    randomToken(),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}


func check_login(email string, Password string) bool {
	db, err := Connect()
	check_err(err)
	var storedPassword string
	err = db.DB.QueryRow("SELECT password FROM Account WHERE email = ?", email).Scan(&storedPassword)

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

func Router_login(router *http.ServeMux) {
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			check := login(data)
			create_cookie(w)
			fmt.Fprintln(w, check)
		case "GET":
			fmt.Println("Method is not used")
		}
	})
}

func login(jsonData []byte) string {
	var Account account
	err := json.Unmarshal(jsonData, &Account)
	check_err(err)

	hashPassword := encode_data(Account.Email, Account.Password, 2)

	check := check_login(Account.Email, hashPassword)
	JsonData, err := json.Marshal(check)
	check_err(err)

	return string(JsonData)
}


func muxtiplexer_router(router *http.ServeMux) {
	Router_login(router)
	dbInfo, err := Connect()
	check_err(err)
	router.HandleFunc("/select", select_Handler(dbInfo))
	router.HandleFunc("/delete", delete_Handler(dbInfo))
	router.HandleFunc("/update", update_Handler(dbInfo))
}

func Create_server() {
	router := http.NewServeMux()
	muxtiplexer_router(router)

	server := http.Server{
		Addr:    ":5050",
		Handler: enable_middleware_cors(router),
	}
	server.ListenAndServe()
}

func main() {
	Create_server()
}

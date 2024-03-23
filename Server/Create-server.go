package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/cors"
)

func Enable_cors(handler http.Handler) http.Handler {
	return cors.Default().Handler(handler)
}

func enable_middleware_cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Cors := cors.New(cors.Options{
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Language", "Content-Type"},
			AllowedMethods:   []string{"POST", "GET"},
			AllowedOrigins:   []string{"http://127.0.0.1:5500"},
			AllowCredentials: true,
			Debug:            true,
		})
		Cors.ServeHTTP(w, r, next.ServeHTTP)
	})
}

func Router_login(router *http.ServeMux) {
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			check := login(data)
			fmt.Fprintln(w, check)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_create_account(router *http.ServeMux) {
	router.HandleFunc("/CreateAccount", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			check := create_account(data)
			fmt.Fprintln(w, check)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_get_account(router *http.ServeMux) {
	router.HandleFunc("/GetAccount", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			list := get_Account(data)
			fmt.Fprintln(w, list)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func Router_check_role(router *http.ServeMux) {
	router.HandleFunc("/CheckRole", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "POST":
			data, err := io.ReadAll(r.Body)
			check_err(err)

			check_role := check_Role(data)
			fmt.Fprintln(w, check_role)
		case "GET":
			fmt.Println("Get method is not used")
		}
	})
}

func muxtiplexer_router(router *http.ServeMux) {
	Router_create_account(router)
	Router_get_account(router)
	Router_check_role(router)
}

func Create_server() {
	router := http.NewServeMux()
	muxtiplexer_router(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "welcome to server my server")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: enable_middleware_cors(router),
	}
	server.ListenAndServe()
}

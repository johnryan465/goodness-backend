package main

import (
	"fmt"
	"goodness/app"
	"goodness/controllers"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/scores", controllers.GetScore).Methods("POST")
	router.HandleFunc("/api/bank/token", controllers.SetToken).Methods("POST")
	router.HandleFunc("/api/bank/transactions", controllers.GetTransactions).Methods("POST")
	router.HandleFunc("/api/data", controllers.GetData).Methods("POST")
	router.HandleFunc("/api/auth", controllers.GetToken).Methods("GET")
	router.HandleFunc("/api/init", controllers.BerkInit).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port,handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		fmt.Print(err)
	}
}

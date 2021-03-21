package main

import (
	"coffeapp/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/coffes", handlers.GetCoffes).Methods("GET")
	r.HandleFunc("/coffe/{id}", handlers.GetCoffe).Methods("GET")
	r.HandleFunc("/coffe", handlers.CreateCoffe).Methods("POST")
	r.HandleFunc("/coffe/{id}", handlers.UpdateCoffe).Methods("PUT")
	r.HandleFunc("/coffe/{id}", handlers.DeleteCoffe).Methods("DELETE")

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handlers.InitialMigration()
	initializeRouter()
}

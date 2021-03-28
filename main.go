package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dionarya23/coffee-app-backend/coffee"
	"github.com/dionarya23/coffee-app-backend/config"
	"github.com/dionarya23/coffee-app-backend/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	db, err := config.ConnectToDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	coffeeRepository := coffee.NewRepository(db)
	coffeeService := coffee.NewService(coffeeRepository)
	coffeeHandler := handler.NewCoffeeHandler(coffeeService)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/coffe", coffeeHandler.CreateCoffee).Methods("POST")
	r.HandleFunc("/api/v1/coffes", coffeeHandler.GetCoffees).Methods("GET")
	r.HandleFunc("/api/v1/coffe/{id}", coffeeHandler.GetCoffeeByID).Methods("GET")
	r.HandleFunc("/api/v1/coffe/{id}", coffeeHandler.UpdateCoffeByID).Methods("PUT")
	r.HandleFunc("/api/v1/coffe/{id}", coffeeHandler.DeleteCoffeByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dionarya23/coffee-app-backend/coffee"
	"github.com/dionarya23/coffee-app-backend/helper"
	"github.com/gorilla/mux"
)

type coffeeHandler struct {
	coffeeService coffee.Service
}

func NewCoffeeHandler(coffeeService coffee.Service) *coffeeHandler {
	return &coffeeHandler{coffeeService}
}

func (h *coffeeHandler) CreateCoffee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input coffee.CreateNewCoffeeInput

	json.NewDecoder(r.Body).Decode(&input)
	newCoffee, err := h.coffeeService.CreateCoffee(input)

	if err != nil {
		responseFail := helper.APIResponse("Error ketika membuat data coffe baru", http.StatusBadRequest, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	formatter := coffee.FormatCoffee(newCoffee)
	response := helper.APIResponse("success", http.StatusCreated, formatter)
	json.NewEncoder(w).Encode(&response)
}

func (h *coffeeHandler) GetCoffees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coffees, err := h.coffeeService.GetCoffees()

	if err != nil {
		responseFail := helper.APIResponse("Error ketika mengambil data coffe", http.StatusBadRequest, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	formatter := coffee.FormatCoffees(coffees)
	response := helper.APIResponse("success", http.StatusCreated, formatter)
	json.NewEncoder(w).Encode(&response)
}

func (h *coffeeHandler) GetCoffeeByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	coffee_, err := h.coffeeService.GetCoffee(ID)

	if err != nil {
		responseFail := helper.APIResponse("Error ketika mengambil data coffe", http.StatusBadRequest, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	if coffee_.ID == 0 {
		responseFail := helper.APIResponse("Error ketika mengambil data coffe", http.StatusNotFound, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	formatter := coffee.FormatCoffee(coffee_)
	response := helper.APIResponse("success", http.StatusCreated, formatter)
	json.NewEncoder(w).Encode(&response)
}

func (h *coffeeHandler) UpdateCoffeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input coffee.CreateNewCoffeeInput

	json.NewDecoder(r.Body).Decode(&input)
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])

	updatedCoffee, err := h.coffeeService.UpdateCoffee(ID, input)

	if err != nil {
		responseFail := helper.APIResponse("Error ketika menghapus data coffe", http.StatusBadRequest, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	formatter := coffee.FormatCoffee(updatedCoffee)
	response := helper.APIResponse("success", http.StatusCreated, formatter)
	json.NewEncoder(w).Encode(&response)
}

func (h *coffeeHandler) DeleteCoffeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])

	_, err := h.coffeeService.DeleteCoffee(ID)
	if err != nil {
		responseFail := helper.APIResponse("Error ketika menghapus data coffe", http.StatusBadRequest, nil)
		json.NewEncoder(w).Encode(&responseFail)
		return
	}

	// formatter := coffee.FormatCoffee(updatedCoffee)
	response := helper.APIResponse("success", http.StatusCreated, nil)
	json.NewEncoder(w).Encode(&response)
}

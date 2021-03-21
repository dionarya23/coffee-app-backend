package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Coffe struct {
	gorm.Model
	Price       float32 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Latitude    float32 `json:"latitude"`
	Longtitude  float32 `json:"longtitude"`
	Address     string  `json:"address"`
	Type        string  `gorm:"type:enum('Espresso', 'Macchiato', 'Piccolo', 'Cappucino', 'Latte', 'Ristretto', 'Americano', 'Cortado', 'Flat White', 'Affogato', 'Moccacino');default:'Cappucino'" json:"type"`
}

type MessageResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Coffe `json:"data"`
}

type MessageResponseSingle struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Coffe  `json:"data"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Coffe{})
}

func GetCoffes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var coffes []Coffe

	DB.Find(&coffes)
	json.NewEncoder(w).Encode(&MessageResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    coffes,
	})
}

func GetCoffe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var coffe Coffe
	DB.First(&coffe, params["id"])
	json.NewEncoder(w).Encode(&MessageResponseSingle{
		Status:  http.StatusOK,
		Message: "success",
		Data:    coffe,
	})
}

func CreateCoffe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var coffe Coffe
	json.NewDecoder(r.Body).Decode(&coffe)
	DB.Create(&coffe)
	json.NewEncoder(w).Encode(&MessageResponseSingle{
		Status:  http.StatusCreated,
		Message: "created",
		Data:    coffe,
	})
}

func UpdateCoffe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var coffe Coffe
	DB.First(&coffe, params["id"])
	json.NewDecoder(r.Body).Decode(&coffe)
	DB.Save(&coffe)
	json.NewEncoder(w).Encode(&MessageResponseSingle{
		Status:  http.StatusOK,
		Message: "success",
		Data:    coffe,
	})
}

func DeleteCoffe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var coffe Coffe
	DB.Delete(&coffe, params["id"])
	json.NewEncoder(w).Encode(&MessageResponseSingle{
		Status:  http.StatusOK,
		Message: "success deleted",
		Data:    coffe,
	})
}

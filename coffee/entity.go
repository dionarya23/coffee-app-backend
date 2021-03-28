package coffee

import "gorm.io/gorm"

type Coffee struct {
	gorm.Model
	Price       float32 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Latitude    float32 `json:"latitude"`
	Longtitude  float32 `json:"longtitude"`
	Address     string  `json:"address"`
	Type        string  `gorm:"type:enum('Espresso', 'Macchiato', 'Piccolo', 'Cappucino', 'Latte', 'Ristretto', 'Americano', 'Cortado', 'Flat White', 'Affogato', 'Moccacino');default:'Cappucino'" json:"type"`
}

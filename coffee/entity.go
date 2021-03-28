package coffee

import "gorm.io/gorm"

type Coffee struct {
	gorm.Model
	Price       float32
	Thumbnail   string
	Description string
	Latitude    float32
	Longtitude  float32
	Address     string
	Type        string
}

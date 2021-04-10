package coffee

type CoffeeFormatter struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Thumbnail   string  `json:"thumbnail"`
	Description string  `json:"description"`
	Latitude    float32 `json:"latitude"`
	Longtitude  float32 `json:"longtitude"`
	Address     string  `json:"address"`
	Type        string  `json:"type"`
}

func FormatCoffee(coffee Coffee) CoffeeFormatter {
	formatter := CoffeeFormatter{
		ID:          coffee.ID,
		Name:        coffee.Name,
		Price:       coffee.Price,
		Thumbnail:   coffee.Thumbnail,
		Description: coffee.Description,
		Latitude:    coffee.Latitude,
		Longtitude:  coffee.Longtitude,
		Address:     coffee.Address,
		Type:        coffee.Type,
	}

	return formatter
}

func FormatCoffees(coffees []Coffee) []CoffeeFormatter {

	coffeesFormatter := []CoffeeFormatter{}

	for _, coffee := range coffees {
		coffeeFormatter := FormatCoffee(coffee)
		coffeesFormatter = append(coffeesFormatter, coffeeFormatter)
	}

	return coffeesFormatter
}

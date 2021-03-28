package coffee

import "errors"

type Service interface {
	CreateCoffee(input CreateNewCoffeeInput) (Coffee, error)
	GetCoffees() ([]Coffee, error)
	GetCoffee(ID int) (Coffee, error)
	UpdateCoffee(ID int, Updatedcoffee CreateNewCoffeeInput) (Coffee, error)
	DeleteCoffee(ID int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCoffee(input CreateNewCoffeeInput) (Coffee, error) {
	coffee_ := Coffee{}

	coffee_.Price = input.Price
	coffee_.Thumbnail = input.Thumbnail
	coffee_.Description = input.Description
	coffee_.Latitude = input.Latitude
	coffee_.Longtitude = input.Longtitude
	coffee_.Address = input.Address
	coffee_.Type = input.Type

	newCoffee, err := s.repository.Save(coffee_)
	if err != nil {
		return newCoffee, err
	}

	return newCoffee, nil
}

func (s *service) GetCoffees() ([]Coffee, error) {
	coffee_, err := s.repository.FindAll()

	if err != nil {
		return coffee_, err
	}

	return coffee_, err
}

func (s *service) GetCoffee(ID int) (Coffee, error) {
	coffee_, err := s.repository.FinByID(ID)

	if err != nil {
		return coffee_, err
	}

	return coffee_, nil
}

func (s *service) UpdateCoffee(ID int, updatedcoffee CreateNewCoffeeInput) (Coffee, error) {
	coffee_, err := s.repository.FinByID(ID)

	if err != nil {
		return coffee_, err
	}

	if coffee_.ID == 0 {
		return coffee_, errors.New("coffee not found")
	}

	coffee_.Price = updatedcoffee.Price
	coffee_.Thumbnail = updatedcoffee.Thumbnail
	coffee_.Description = updatedcoffee.Description
	coffee_.Latitude = updatedcoffee.Latitude
	coffee_.Longtitude = updatedcoffee.Longtitude
	coffee_.Address = updatedcoffee.Address
	coffee_.Type = updatedcoffee.Type

	update, err := s.repository.Update(coffee_)
	if err != nil {
		return coffee_, err
	}

	return update, nil
}

func (s *service) DeleteCoffee(ID int) (bool, error) {
	coffee_, err := s.repository.FinByID(ID)

	if err != nil {
		return false, err
	}

	if coffee_.ID == 0 {
		return false, errors.New("coffee not found")
	}

	s.repository.Delete(ID)

	return true, nil
}

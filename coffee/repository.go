package coffee

import "gorm.io/gorm"

type Repository interface {
	Save(coffee_ Coffee) (Coffee, error)
	FindAll() ([]Coffee, error)
	FinByID(ID int) (Coffee, error)
	Update(coffee_ Coffee) (Coffee, error)
	Delete(ID int)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(coffee_ Coffee) (Coffee, error) {
	err := r.db.Create(&coffee_).Error

	if err != nil {
		return coffee_, err
	}

	return coffee_, nil
}

func (r *repository) FindAll() ([]Coffee, error) {
	var coffee_ []Coffee

	err := r.db.Find(&coffee_).Error
	if err != nil {
		return coffee_, err
	}
	return coffee_, nil
}

func (r *repository) FinByID(ID int) (Coffee, error) {
	var coffee_ Coffee
	err := r.db.Where("id=?", ID).Find(&coffee_).Error
	if err != nil {
		return coffee_, err
	}

	return coffee_, nil
}

func (r *repository) Update(coffee_ Coffee) (Coffee, error) {
	err := r.db.Save(&coffee_).Error

	if err != nil {
		return coffee_, err
	}

	return coffee_, nil
}

func (r *repository) Delete(ID int) {
	var coffee_ Coffee
	r.db.Delete(&coffee_, ID)
}

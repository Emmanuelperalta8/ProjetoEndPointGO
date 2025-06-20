package services

import (
	"coffee-delivery/database"
	"coffee-delivery/models"
	"errors"

	"gorm.io/gorm"
)

func CreateCoffee(cafe models.Coffee) error {
	return database.DB.Create(&cafe).Error
}

func FindAllCoffees() ([]models.Coffee, error) {
	var cafes []models.Coffee
	err := database.DB.Preload("Tags").Preload("Pedidos").Find(&cafes).Error
	return cafes, err
}

func DeleteCoffee(id uint) error {
	var cafe models.Coffee
	if err := database.DB.First(&cafe, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return database.DB.Delete(&cafe).Error
}

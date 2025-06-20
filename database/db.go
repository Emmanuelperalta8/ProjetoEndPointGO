package database

import (
	"coffee-delivery/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("coffees.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
	DB.AutoMigrate(&models.Coffee{}, &models.Tag{}, &models.Pedido{})
}

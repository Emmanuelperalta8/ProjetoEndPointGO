package database

import (
	"log"

	"coffee-delivery/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	// Usando SQLite com suporte a foreign keys ativado (sem CGO)
	DB, err = gorm.Open(sqlite.Open("file:coffees.db?_pragma=foreign_keys=on"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = DB.AutoMigrate(&models.Coffee{}, &models.Tag{}, &models.Pedido{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco:", err)
	}
}

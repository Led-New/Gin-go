package database

import (
	"log"

	"github.com/Led-New/Gin-Go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	stringConnection := "host=localhost user=Darksanpaio password=1234 dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		log.Panic("erro no BAnco")
	}
	DB.AutoMigrate(&models.Aluno{})
}

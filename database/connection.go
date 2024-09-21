package database

import (
	"log"

	"github.com/Msad668/library-management/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global database connection
var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
    config := configs.GetConfig()
    var err error
    DB, err = gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    log.Println("Database connection established")
}
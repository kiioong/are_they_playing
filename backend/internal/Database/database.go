package Database

import (
	"log"
	"os"

	hash "github.com/kiioong/are_they_playing/internal/Hash"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

var DB *gorm.DB

func InitDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})

	password, err := hash.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return
	}

	DB.Create(&User{Username: "admin", Password: password})
}

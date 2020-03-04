package migrate

import (
	"fmt"
	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var db *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	createdAt time.Time
	updatedAt time.Time
}

func InitDB(config DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		config.Host, config.Port, config.User, config.DBName, config.Password))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
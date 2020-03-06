package migrate

import (
	"fmt"
	"time"

	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	ID        uint `gorm:"primary_key" json:"id"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
}

func InitDB(config DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		config.Host, config.Port, config.User, config.DBName, config.Password))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	return db, nil
}

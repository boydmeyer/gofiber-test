package database

import (
	"fmt"
	"strconv"

	"github.com/boydmeyer/fiber/app/config"
	"github.com/boydmeyer/fiber/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	var err error
	p := config.Config("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s  port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Post{}, &model.User{})
	fmt.Println("Database Migrated")
}

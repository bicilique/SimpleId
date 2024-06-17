package config

import (
	"SimpleId/internal/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(cfg *Config) *gorm.DB {
	Db := connectDB(cfg)
	return Db
}

func connectDB(cfg *Config) *gorm.DB {
	username := cfg.DBUser
	password := cfg.DBPassword
	dbHost := "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" //"@tcp(localhost:3306)/"
	dbName := cfg.DBName
	dsn := username + ":" + password + dbHost + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Request{},
		&entity.Shared{},
	)
	return db
}

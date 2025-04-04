package config

import (
	"log"

	"horizen-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB connects to MySQL using environment variables.
func InitDB() *gorm.DB {
	dsn := "root:my_secure_password@tcp(localhost)/horizen_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	DB = db
	return db
}

// MigrateDB performs auto migration if you wish to let GORM handle schema creation.
// Comment this out if you manage schema manually.
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

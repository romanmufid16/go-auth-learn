package config

import (
	"github.com/romanmufid16/go-auth-learn/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func SetupDatabaseConnection() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Menggunakan log.New untuk membuat writer
		logger.Config{
			SlowThreshold: time.Second, // Jika query lambat lebih dari 1 detik
			LogLevel:      logger.Info, // Menampilkan log pada level Info
			Colorful:      true,        // Menampilkan log berwarna
		},
	)
	var err error
	dsn := os.Getenv("DATABASE_URI")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connecting database")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}

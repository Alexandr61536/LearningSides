package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/internal/app"
	"backend/internal/database/migrations"
	"backend/internal/models"
)

func InitDB() {
	log.Println("Starting DB")
	db_name, exists := os.LookupEnv("DB_NAME")
	var db *gorm.DB

	if !exists {
		log.Fatal("No DB name in .env")
		panic("No DB name in .env")
	}

	if _, err := os.Stat(db_name); os.IsNotExist(err) {
		log.Println("DB doesn`t exist, creating...")

		master_login, exists := os.LookupEnv("MASTER_LOGIN")
		if !exists {
			log.Fatal("No master login in .env")
			panic("No master login in .env")
		}
		master_password, exists := os.LookupEnv("MASTER_PASSWORD")
		if !exists {
			log.Fatal("No master password in .env")
			panic("No master password in .env")
		}

		db, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting DB:", err)
		}

		log.Println("Connected successfully")

		if migration_result := migrations.AutoMigrateAll(db); !migration_result.Status {
			log.Fatal("Migrations failed", migration_result.Err)
			panic("Migrations failed")
		} else {
			log.Println("Migrated successfully")
		}

		hash_pass, err := app.HashPassword(master_password)

		if err != nil {
			log.Fatal("Can`t hash master password")
		}

		db.Create(&models.User{Login: master_login, Password: hash_pass, Role: "admin"})
		log.Println("Created master account")

	} else {
		log.Println("DB exists, connecting...")

		db, err = gorm.Open(sqlite.Open(db_name), &gorm.Config{})
		if err != nil {
			log.Fatal("Error connecting DB:", err)
		}

		log.Println("Connected successfully")

		if migration_result := migrations.AutoMigrateAll(db); !migration_result.Status {
			log.Fatal("Migrations failed", migration_result.Err)
			panic("Migrations failed")
		} else {
			log.Println("Migrated successfully")
		}
	}

	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
}

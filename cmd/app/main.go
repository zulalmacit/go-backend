package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-backend/internal/database"
	"go-backend/internal/server"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=backend_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Veritabanı bağlantı hatası:", err)
	}

	// Otomatik migration
	if err := database.Migrate(db); err != nil {
		log.Fatal("❌ Migration hatası:", err)
	}

	// Sunucuyu başlat
	server.StartServer(db)
}

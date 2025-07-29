package main

import (
	"fmt"
	"log"

	"github.com/zulal/go-backend/internal/database"
	"github.com/zulal/go-backend/internal/domain"
	"github.com/zulal/go-backend/internal/repository"
	"github.com/zulal/go-backend/internal/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Veritabanı bağlantı dizesi (DSN)
	dsn := "host=localhost user=postgres password=postgres dbname=backend_db port=5432 sslmode=disable"

	// PostgreSQL'e bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Veritabanına bağlanılamadı:", err)
	}

	// Otomatik migration
	err = database.Migrate(db)
	if err != nil {
		log.Fatal("❌ Migration hatası:", err)
	}
	fmt.Println("✅ Migration başarılı!")

	// Repository ve service başlat
	userRepo := repository.NewUserGormRepository(db)
	userService := service.NewUserService(userRepo)

	// Örnek kullanıcı oluştur
	user := &domain.User{
		Name:         "Zulal",
		Email:        "zulal@example.com",
		PasswordHash: "test1234", // ham şifre, servis içinde hashlenecek
	}

	// Kullanıcı kaydı
	err = userService.Register(user)
	if err != nil {
		log.Println("❌ Kullanıcı oluşturulamadı:", err)
	} else {
		fmt.Println("🎉 Kullanıcı başarıyla oluşturuldu!")
	}
}

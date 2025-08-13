package server

import (
	"fmt"
	"log"
	"net/http"

	"go-backend/internal/router"

	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) {
	r := router.SetupRouter(db)

	fmt.Println("🚀 Sunucu http://localhost:8080 adresinde başlatılıyor...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("❌ Sunucu başlatılamadı: %v", err)
	}
}

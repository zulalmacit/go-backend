package router

import (
	"net/http"

	"github.com/zulal/go-backend/internal/handler"
	"github.com/zulal/go-backend/internal/middleware"
	"github.com/zulal/go-backend/internal/repository"
	"github.com/zulal/go-backend/internal/service"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	// Ortak middleware'leri yükle
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Katmanları başlat
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Route'ları ayarla
	r.Route("/api", func(api chi.Router) {
		api.Post("/register", userHandler.Register)
		api.Post("/login", userHandler.Login)
		// buraya başka endpoint'ler de eklenebilir
	})

	return r
}

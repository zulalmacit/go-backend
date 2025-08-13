package handler

import (
	"encoding/json"
	"net/http"

	"go-backend/internal/domain"
	"go-backend/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Geçersiz istek", http.StatusBadRequest)
		return
	}

	if err := h.userService.Register(&user); err != nil {
		http.Error(w, "Kayıt başarısız: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kayıt başarılı"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Geçersiz istek", http.StatusBadRequest)
		return
	}

	user, token, err := h.userService.Login(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Giriş başarısız: "+err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  user,
		"token": token,
	})
}

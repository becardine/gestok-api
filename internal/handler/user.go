package handler

import (
	"encoding/json"
	"net/http"

	"github.com/becardine/gestock-api/internal/entity"
	"github.com/becardine/gestock-api/internal/service"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	user, err := h.userService.FindUserByEmail(email)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

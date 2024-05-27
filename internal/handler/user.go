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

// @BasePath /api/v1
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body entity.User true "Create User Input"
// @Success 201
// @Failure 400 {object} errors.ErrorHandler
// @Failure 500 {object} errors.ErrorHandler
// @Router /users [post]
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

// @BasePath /api/v1
// @Summary Get user by email
// @Description Get user by email
// @Tags users
// @Accept  json
// @Produce  json
// @Param email path string true "User Email"
// @Success 200 {object} entity.User
// @Failure 404 {object} errors.ErrorHandler
// @Failure 500 {object} errors.ErrorHandler
// @Router /users/{email} [get]
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

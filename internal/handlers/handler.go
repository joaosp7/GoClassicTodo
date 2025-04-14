package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaosp7/GoClassicTodo/internal/dto"
	"github.com/joaosp7/GoClassicTodo/internal/services"
)

type Handler struct {
	userService *services.UserService
}

func NewHandler (userService *services.UserService) *Handler{
	return &Handler{userService: userService}
}

func (h *Handler) CreateUser( w http.ResponseWriter, r *http.Request) {
	var input dto.UserRequestDto	
	err := json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	output := h.userService.CreateAccount(input)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
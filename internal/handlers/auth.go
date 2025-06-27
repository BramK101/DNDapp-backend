package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BramK101/DNDapp-backend/internal/models"
	"github.com/BramK101/DNDapp-backend/internal/utils"
)

func (h *Handlers) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, valid := h.Services.ValidateUser(req.Email, req.Password)
	if !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateJWT(uint(user.ID), user.Email)
	if err != nil {
		http.Error(w, "Token generation failed", http.StatusInternalServerError)
		return
	}

	response := models.LoginResponse{Token: token, User: user}
	json.NewEncoder(w).Encode(response)
}

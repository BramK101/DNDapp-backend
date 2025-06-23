package handlers

import (
	"log"
	"net/http"

	"github.com/BramK101/DNDapp-backend/internal/services"
	"github.com/BramK101/DNDapp-backend/internal/utils"
)

func (h *Handlers) SetupRoutes() {
    http.HandleFunc("/users", h.createUser)
    http.HandleFunc("/users/", h.getUser)
	http.HandleFunc("/login", h.loginHandler)

	http.HandleFunc("/dashboard", utils.AuthMiddleware(h.loginHandler))
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Handlers struct {
    Services *services.Services
}

func NewHandlers(Services *services.Services) *Handlers {
	return &Handlers{Services: Services}
}
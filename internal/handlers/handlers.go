package handlers

import (
	"log"
	"net/http"

	"github.com/BramK101/DNDapp-backend/internal/config"
	"github.com/BramK101/DNDapp-backend/internal/middleware"
	"github.com/BramK101/DNDapp-backend/internal/services"
)

func (h *Handlers) SetupRoutes(cfg *config.Config) {
	http.HandleFunc("/register", middleware.CORS(h.createUser))
	http.HandleFunc("/users/", middleware.CORS(middleware.Auth(h.getUser)))
	http.HandleFunc("/login", middleware.CORS(h.loginHandler))

	log.Fatal(http.ListenAndServe(cfg.UrlPort, nil))
}

type Handlers struct {
	Services *services.Services
}

func NewHandlers(Services *services.Services) *Handlers {
	return &Handlers{Services: Services}
}

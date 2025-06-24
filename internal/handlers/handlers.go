package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BramK101/DNDapp-backend/internal/middleware"
	"github.com/BramK101/DNDapp-backend/internal/services"
)

func (h *Handlers) SetupRoutes() {
    http.HandleFunc("/users", middleware.CORS(h.createUser))
    http.HandleFunc("/users/", middleware.CORS(h.getUser))
	http.HandleFunc("/login", middleware.CORS(h.loginHandler))

	http.HandleFunc("/dashboard", middleware.CORS(middleware.Auth(h.showDashboard)))
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type Handlers struct {
    Services *services.Services
}

func NewHandlers(Services *services.Services) *Handlers {
	return &Handlers{Services: Services}
}

func (h *Handlers) showDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Dashboard")
}
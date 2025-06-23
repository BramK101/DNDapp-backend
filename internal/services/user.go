package services

import (
	"github.com/BramK101/DNDapp-backend/internal/models"
)

func (s *Services) CreateUser(username, email, password string) (*models.User, error) {
    query := `
        INSERT INTO users (username, email, password, created_at) 
        VALUES ($1, $2, $3, NOW()) 
        RETURNING id, username, email, created_at`
    
    user := &models.User{}
    err := s.db.QueryRow(query, username, email, password).Scan(
        &user.ID, &user.Username, &user.Email, &user.CreatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *Services) GetUserByID(id int) (*models.User, error) {
    query := `SELECT id, username, email, created_at FROM users WHERE id = $1`
    
    user := &models.User{}
    err := s.db.QueryRow(query, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.CreatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *Services) ValidateUser(username string, password string) (models.User, bool) {
    query := `SELECT id, username, email, password, created_at FROM users WHERE username = $1`
    
    user := models.User{}
    err := s.db.QueryRow(query, username).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
    )

    if err != nil {
        return user, false
    }

    if user.Password == password {
        return user, true
    }
    
    return user, false
}
package services

import (
	"log"

	"github.com/BramK101/DNDapp-backend/internal/models"
	"github.com/BramK101/DNDapp-backend/internal/utils"
)

func (s *Services) CreateUser(username, email, password string) (*models.User, error) {
    query := `
        INSERT INTO users (username, email, password, created_at) 
        VALUES ($1, $2, $3, NOW()) 
        RETURNING id, username, email, created_at`
    
    user := &models.User{}
    hashedPassword, errHash := utils.HashPassword(password)
    if errHash != nil {
        return nil, errHash
    }
    err := s.db.QueryRow(query, username, email, hashedPassword).Scan(
        &user.ID, &user.Username, &user.Email, &user.CreatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *Services) GetUserByID(id int) (*models.User, error) {
    query := `SELECT id, username, email, password, created_at FROM users WHERE id = $1`
    
    user := &models.User{}
    err := s.db.QueryRow(query, id).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *Services) ValidateUser(email string, password string) (models.User, bool) {
    query := `SELECT id, username, email, password, created_at FROM users WHERE email = $1`
    
    user := models.User{}
    err := s.db.QueryRow(query, email).Scan(
        &user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
    )

    if err != nil {
        log.Print("User not found!", err)
        return user, false
    }

    err1 := utils.VerifyPassword(user.Password, password)
	if err1 != nil {
        log.Print("Verify password hash went wrong: ", err1)
		return user, false
	}

    
    return user, true
}
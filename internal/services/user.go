package services

import (
	"log"

	"github.com/BramK101/DNDapp-backend/internal/models"
	"github.com/BramK101/DNDapp-backend/internal/utils"
)

func (s *Services) CreateUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Services) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	if err := s.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Services) ValidateUser(email string, password string) (models.User, bool) {
	user := models.User{}
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Print("User not found!", err)
		return user, false
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		log.Print("Verify password hash went wrong: ", err)
		return user, false
	}

	return user, true
}

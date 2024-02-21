package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUser(email string) (*models.User, error)
}

type AuthReoisitory struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthReoisitory{db: db}
}

func (r *AuthReoisitory) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AuthReoisitory) FindUser(email string) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, "email =?", email)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("User not found")
		}
	}

	return &user, nil
}

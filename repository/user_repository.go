package repository

import (
	"github.com/romanmufid16/go-auth-learn/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) error
}

// Implementasi Repository
type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// Implementasi method Create
func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Implementasi method GetById
func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Implementasi method FindByEmail
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Implementasi metode Update
func (r *userRepository) Update(user *models.User) (*models.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Implementasi metode Delete
func (r *userRepository) Delete(id uint) error {
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

package service

import (
	"errors"
	"github.com/romanmufid16/go-auth-learn/dto"
	"github.com/romanmufid16/go-auth-learn/models"
	"github.com/romanmufid16/go-auth-learn/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userDTO *dto.RegisterUser) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService untuk membuat instance userService
func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepository,
	}
}

func (s *userService) Register(userDTO *dto.RegisterUser) (*dto.UserResponse, error) {
	existingUser, _ := s.userRepo.FindByEmail(userDTO.Email)
	if existingUser != nil {
		return nil, errors.New("Email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), 10)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	userResponse := dto.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	return &userResponse, nil
}
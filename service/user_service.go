package service

import (
	"errors"
	"github.com/romanmufid16/go-auth-learn/dto"
	"github.com/romanmufid16/go-auth-learn/models"
	"github.com/romanmufid16/go-auth-learn/repository"
	"github.com/romanmufid16/go-auth-learn/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userDTO *dto.RegisterUser) (*dto.UserResponse, error)
	Login(userDTO *dto.LoginUser) (*dto.TokenResponse, error)
	GetUser(id uint) (*dto.UserResponse, error)
	UpdateUser(id uint, userDTO *dto.UpdateUser) (*dto.UserResponse, error)
	DeleteUser(id uint) (bool, error)
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

func (s *userService) Login(userDTO *dto.LoginUser) (*dto.TokenResponse, error) {
	existingUser, _ := s.userRepo.FindByEmail(userDTO.Email)
	if existingUser == nil {
		return nil, errors.New("Invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userDTO.Password))
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	token, err := utils.GenerateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return nil, err
	}

	tokenResponse := dto.TokenResponse{
		Token: token,
	}

	return &tokenResponse, nil
}

func (s *userService) GetUser(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	// Transformasi ke DTO
	userResponse := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userResponse, nil
}

func (s *userService) UpdateUser(id uint, userDTO *dto.UpdateUser) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		return nil, errors.New("User Not found")
	}

	if userDTO.Name != "" {
		user.Name = userDTO.Name
	}

	if userDTO.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), 10)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	updatedUser, err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}

	return userResponse, nil
}

func (s *userService) DeleteUser(id uint) (bool, error) {
	_, err := s.userRepo.GetById(id)
	if err != nil {
		return false, errors.New("User not found")
	}

	err = s.userRepo.Delete(id)
	if err != nil {
		return false, errors.New("Delete failed")
	}

	return true, nil
}

package service

import (
	"backend-boking-ticket/internal/entity"
	"backend-boking-ticket/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req entity.RegisterRequest) (entity.AuthResponse, error)
	Login(req entity.LoginRequest) (entity.AuthResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(req entity.RegisterRequest) (entity.AuthResponse, error) {
	// 1. Cek User Exist
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err == nil && existingUser.ID != "" {
		return entity.AuthResponse{}, errors.New("email already registered")
	}

	// 2. Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.AuthResponse{}, err
	}

	// 3. Create User
	newUser := entity.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err = s.repo.Create(newUser)
	if err != nil {
		return entity.AuthResponse{}, err
	}

	return entity.AuthResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (s *authService) Login(req entity.LoginRequest) (entity.AuthResponse, error) {
	// 1. Cari User
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return entity.AuthResponse{}, errors.New("invalid email or password")
	}

	// 2. Cek Password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return entity.AuthResponse{}, errors.New("invalid email or password")
	}

	// 3. Login sukses
	return entity.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		// Token: "JWT_TOKEN_HERE", // Nanti akan diisi JWT
	}, nil
}

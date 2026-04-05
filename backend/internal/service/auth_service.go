package service

import (
	"context"
	"errors"

	"github.com/muhammetalicaliskan/libranet/internal/config"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
	jwtpkg "github.com/muhammetalicaliskan/libranet/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailAlreadyExists = errors.New("bu e-posta adresi zaten kayıtlı")
	ErrInvalidCredentials = errors.New("e-posta veya şifre hatalı")
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{userRepo: userRepo, cfg: cfg}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	existing, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     model.RoleUser,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Email, string(user.Role), s.cfg.JWT.Secret, s.cfg.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Email, string(user.Role), s.cfg.JWT.Secret, s.cfg.JWT.Expiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

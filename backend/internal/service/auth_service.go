package service

import (
	"context"
	"errors"
	"time"

	"github.com/muhammetalicaliskan/libranet/internal/config"
	"github.com/muhammetalicaliskan/libranet/internal/dto"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/muhammetalicaliskan/libranet/internal/repository"
	jwtpkg "github.com/muhammetalicaliskan/libranet/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaxFailedLoginAttempts = 5
	AccountLockDuration    = 15 * time.Minute
)

var (
	ErrEmailAlreadyExists = errors.New("bu e-posta adresi zaten kayıtlı")
	ErrInvalidCredentials = errors.New("e-posta veya şifre hatalı")
	ErrAccountLocked      = errors.New("hesabınız geçici olarak kilitlendi, lütfen daha sonra tekrar deneyin")
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

	// Check if account is locked
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, ErrAccountLocked
	}

	// If lock expired, reset
	if user.LockedUntil != nil && user.LockedUntil.Before(time.Now()) {
		user.FailedLoginAttempts = 0
		user.LockedUntil = nil
		_ = s.userRepo.Update(ctx, user)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		user.FailedLoginAttempts++
		if user.FailedLoginAttempts >= MaxFailedLoginAttempts {
			lockUntil := time.Now().Add(AccountLockDuration)
			user.LockedUntil = &lockUntil
		}
		_ = s.userRepo.Update(ctx, user)

		if user.FailedLoginAttempts >= MaxFailedLoginAttempts {
			return nil, ErrAccountLocked
		}
		return nil, ErrInvalidCredentials
	}

	// Successful login: reset failed attempts
	if user.FailedLoginAttempts > 0 {
		user.FailedLoginAttempts = 0
		user.LockedUntil = nil
		_ = s.userRepo.Update(ctx, user)
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

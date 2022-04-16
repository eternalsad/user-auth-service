package service

import (
	"github.com/eternalsad/user-auth-service/config"
	"github.com/eternalsad/user-auth-service/internal/models"
	"github.com/eternalsad/user-auth-service/internal/repository"
	"go.uber.org/zap"
)

type AuthService struct {
	repo repository.Repository
	cfg  *config.Configuration
	log  *zap.SugaredLogger
}

func NewAuthService(repo repository.Repository, cfg *config.Configuration, log *zap.SugaredLogger) *AuthService {
	return &AuthService{repo, cfg, log}
}

func (s *AuthService) CreateUser(u *models.User) (int64, error) {
	return s.repo.CreateUser(u)
}

func (s *AuthService) SignUp() {
}

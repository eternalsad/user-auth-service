package service

import (
	"github.com/eternalsad/user-auth-service/config"
	"github.com/eternalsad/user-auth-service/internal/models"
	"github.com/eternalsad/user-auth-service/internal/repository"
	"go.uber.org/zap"
)

type Authorization interface {
	CreateUser(*models.User) (int64, error)
	SignUp()
}

type Service struct {
	Authorization
}

func NewService(repo repository.Repository, cfg *config.Configuration, log *zap.SugaredLogger) *Service {
	return &Service{
		NewAuthService(repo, cfg, log),
	}
}

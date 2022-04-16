package repository

import (
	"github.com/eternalsad/user-auth-service/config"
	"github.com/eternalsad/user-auth-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type AuthRepo struct {
	db  *pgxpool.Pool
	cfg *config.Configuration
	log *zap.SugaredLogger
}

func NewAuthRepo(db *pgxpool.Pool, cfg *config.Configuration, log *zap.SugaredLogger) *AuthRepo {
	return &AuthRepo{db, cfg, log}
}

func (a *AuthRepo) CreateUser(u *models.User) (int64, error) {
	return 0, nil
}

func (a *AuthRepo) SignUp() {
}

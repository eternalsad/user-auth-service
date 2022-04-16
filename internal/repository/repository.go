package repository

import "github.com/eternalsad/user-auth-service/internal/models"

type Repository interface {
	CreateUser(*models.User) (int64, error)
	SignUp()
}

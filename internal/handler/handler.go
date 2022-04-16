package handler

import (
	"github.com/eternalsad/user-auth-service/config"
	"github.com/eternalsad/user-auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service.Authorization
	cfg *config.Configuration
	log *zap.SugaredLogger
}

func NewHandler(service service.Authorization, cfg *config.Configuration, log *zap.SugaredLogger) *Handler {
	return &Handler{service, cfg, log}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
	}
	return router
}

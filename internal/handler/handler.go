package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

// func NewHandler() *Handler {

// }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
	}
	return router
}

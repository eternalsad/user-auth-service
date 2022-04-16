package handler

import (
	"net/http"

	"github.com/eternalsad/user-auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
}

func (h *Handler) signUp(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, CreateResponse(nil, models.ErrIncorrectJSON))
	}
	id, err := h.Authorization.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.Status)
	}
}

func CreateResponse(data interface{}, err error) gin.H {
	if err != nil {
		return gin.H{
			"data": nil,
			"error": gin.H{
				"message": err.Error(),
			},
		}
	}
	return gin.H{
		"data":  data,
		"error": nil,
	}
}

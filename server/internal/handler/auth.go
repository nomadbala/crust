package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nomadbala/crust/server/internal/domain/auth"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var request auth.RegistrationRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := auth.ParseRegistrationRequest(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.services.AuthService.SignUp(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) SignIn(c *gin.Context) {
	var request auth.LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := h.services.AuthService.SignIn(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

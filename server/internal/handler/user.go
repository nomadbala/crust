package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nomadbala/crust/server/internal/domain/user"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.UsersService.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) SignUp(c *gin.Context) {
	var request user.RegistrationRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.services.UsersService.SignUp(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

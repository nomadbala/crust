package handler

import (
	"github.com/gin-gonic/gin"
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

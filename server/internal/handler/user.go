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

//func (h *Handler) SendVerificationEmail(c *gin.Context) {
//	id, err := getUserId(c)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
//		return
//	}
//
//	success, err := h.services.UsersService.SendEmailVerification(id)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
//		return
//	}
//
//	if !success {
//		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "failed to send verification email"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "successfully sent verification email"})
//}

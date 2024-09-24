package handler

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Middleware(c *gin.Context) {
	const authHeaderPrefix = "Bearer "

	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
		return
	}

	if !strings.HasPrefix(header, authHeaderPrefix) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}

	token := strings.TrimPrefix(header, authHeaderPrefix)
	userId, err := h.services.AuthService.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (uuid.UUID, error) {
	id, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No user id found"})
		return uuid.Nil, errors.New("user id not found")
	}

	userId, ok := id.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id format"})
		return uuid.Nil, errors.New("invalid user id format")
	}

	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid UUID"})
		return uuid.Nil, errors.New("invalid UUID format")
	}

	return parsedUUID, nil
}

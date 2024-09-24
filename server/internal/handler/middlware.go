package handler

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderPrefix = "Bearer "
)

var (
	ErrorUserIdNotFound      = errors.New("user id not found")
	ErrorInvalidUserIdFormat = errors.New("invalid user id format")
)

func (h *Handler) Middleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
		return
	}

	if !strings.HasPrefix(header, authorizationHeaderPrefix) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}

	token := strings.TrimPrefix(header, authorizationHeaderPrefix)
	userId, err := h.services.AuthService.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("userId", userId)
}

func GetUserIdFromAccessToken(c *gin.Context) (uuid.UUID, error) {
	id, exists := c.Get("userId")
	if !exists {
		return uuid.Nil, ErrorUserIdNotFound
	}

	userId, ok := id.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrorInvalidUserIdFormat
	}

	return userId, nil
}

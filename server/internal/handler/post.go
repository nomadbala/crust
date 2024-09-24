package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nomadbala/crust/server/db/postgres/sqlc"
	"net/http"
)

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.services.PostsService.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(posts) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPostById(c *gin.Context) {
	param := c.Param("id")

	id, err := uuid.Parse(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post, err := h.services.PostsService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePost(c *gin.Context) {
	var param sqlc.CreatePostParams

	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post, err := h.services.PostsService.Create(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

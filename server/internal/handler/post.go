package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPostById(c *gin.Context) {
	idStr := c.Param("id") // Получаем id из параметров URL

	var id pgtype.UUID
	err := id.Scan(idStr) // Используем метод Scan для преобразования
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	post, err := h.services.PostsService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("qwqwe: %s", post.Content)

	c.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePost(c *gin.Context) {
	var params sqlc.CreatePostParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	post, err := h.services.PostsService.Create(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, post)

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nomadbala/crust/server/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) ConfigureRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/", h.SignUp)
			}

			users := v1.Group("/users")
			{
				users.GET("/", h.GetAllUsers)
			}
		}
	}

	return router
}

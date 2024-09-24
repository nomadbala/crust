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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.Middleware)
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/", h.GetAllUsers)
				//users.POST("/send_verification_email", h.SendVerificationEmail)
			}

			posts := v1.Group("/posts")
			{
				posts.POST("/", h.CreatePost)
				posts.GET("/", h.GetAllPosts)
				posts.GET("/:id", h.GetPostById)
			}
		}
	}

	return router
}

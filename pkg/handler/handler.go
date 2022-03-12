package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/max-sanch/AuthJWT/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-in", h.singIn)
		auth.POST("/refresh", h.refresh)
	}

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
			users.DELETE("/:id", h.deleteUser)
		}
	}

	return router
}

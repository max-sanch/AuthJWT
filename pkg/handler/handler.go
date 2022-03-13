package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/max-sanch/AuthJWT/pkg/service"
	"github.com/spf13/viper"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	if viper.GetString("releaseMode") == "True" {
		gin.SetMode(gin.ReleaseMode)
	}

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
			users.DELETE("/", h.deleteUser)
		}
	}

	return router
}

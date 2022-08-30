package handler

import (
	"github.com/TelitsynNikita/bookkeeper-backend/pkg/service"
	"github.com/gin-gonic/gin"
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		request := api.Group("request")
		{
			request.GET("/", h.GetAllRequests)
			request.GET("/:id", h.GetOneRequest)
			request.POST("/", h.CreateRequest)
			request.PATCH("/", h.UpdateRequest)
			request.DELETE("/:id", h.DeleteRequest)
		}
	}

	return router
}

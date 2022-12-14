package handler

import (
	"secrets_keeper/app/pkg/service"

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
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLFiles(
		"templates/index.html",
		"templates/key.html",
		"templates/message.html",
		"templates/404.html",
		"templates/500.html",
		"templates/400.html",
	)

	router.GET("/", h.GetIndexPage)
	router.GET("/message/:key", h.GetMessage)
	router.POST("/", h.SetMessage)

	return router
}

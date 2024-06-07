package pkg

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *Service
}

func NewHandler(services *Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("../templates/*")

	router.GET("/", h.middleware)
	router.POST("/",h.floyd);
	return router
}

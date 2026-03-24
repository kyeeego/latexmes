package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kyeeego/latexmes/delivery/ws"
	"github.com/kyeeego/latexmes/service"
)

type Handler struct {
	services  *service.Service
	wsHandler *ws.Handler
}

func New(service *service.Service, wsHandler *ws.Handler) *Handler {
	return &Handler{service, wsHandler}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	a := router.Group("/")
	{
		a.POST("/signup", h.HandleRegister)
		a.POST("/login", h.HandleLogin)
		a.GET("/authorize", h.HandleAuthorize)
		a.GET("/id/:id", h.HandleGet)
	}

	chat := router.Group("/ws")
	{
		chat.GET("/", h.wsHandler.Connect)
		chat.GET("/test", ws.Test)
	}
}

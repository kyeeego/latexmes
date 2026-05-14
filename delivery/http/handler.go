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
	}

	api := router.Group("/api")
	{
		api.GET("/connections", h.authorize, h.handleGetConnections)
		api.GET("/messages/:conn", h.authorize, h.handleGetConnectionMessages)
		api.GET("/users", h.authorize, h.handleSearchUsers)
		// api.GET("/friendrequests", h.handleGetFriendRequests)
		api.POST("/request/:recid", h.authorize, h.handleFriendRequest)
		api.POST("/reply", h.authorize, h.handleReplyToRequest)
	}

	chat := router.Group("/ws")
	{
		chat.GET("/", h.authorize, h.wsHandler.Connect)
		chat.GET("/test/:userid", ws.Test)
	}
}

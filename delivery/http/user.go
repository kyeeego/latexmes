package http

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyeeego/latexmes/domain/dto"
)

func (h *Handler) HandleRegister(c *gin.Context) {
	var body dto.LoginDto
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	err := h.services.User.Register(body.Username, body.Password)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(201, map[string]string{"username": body.Username})
}

func (h *Handler) HandleLogin(c *gin.Context) {
	var body dto.LoginDto
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	tokens, err := h.services.Auth.Login(body)
	if err != nil {
		c.AbortWithStatusJSON(401, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, tokens)
}

func (h *Handler) authorize(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(401, map[string]error{"error": errors.New("no auth header")})
		return
	}
	a := strings.Split(authHeader, " ")
	if a[0] != "Bearer" || len(a) < 2 {
		c.AbortWithStatusJSON(401, map[string]error{"error": errors.New("invalid auth header")})
		return
	}

	token := a[1]
	userId, err := h.services.Auth.Authorize(token)
	if err != nil {
		c.AbortWithStatusJSON(403, map[string]error{"error": errors.New("invalid token")})
		return
	}

	c.Set("userid", uint(userId))
}

func (h *Handler) handleSearchUsers(c *gin.Context) {
	q := c.Query("q")

	dtos, err := h.services.User.Search(q)
	if err != nil {
		c.JSON(404, map[string]error{"error": err})
		return
	}

	c.JSON(200, dtos)
}

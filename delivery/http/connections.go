package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kyeeego/latexmes/domain/dto"
)

func (h *Handler) handleGetConnections(c *gin.Context) {
	userid := c.GetUint("userid")

	conn, err := h.services.Connection.GetUserConnections(userid)
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, conn)
}

func (h *Handler) handleGetConnectionMessages(c *gin.Context) {
	userid := c.GetUint("userid")
	connstr := c.Param("conn")
	connid, _ := strconv.ParseUint(connstr, 10, 32)

	conn, err := h.services.Connection.GetById(uint(connid))
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	if conn.U1 != userid && conn.U2 != userid {
		c.AbortWithStatusJSON(403, map[string]string{"error": "not yo chat bihhhhh"})
		return
	}

	msgs, err := h.services.Chat.GetConnectionMessages(uint(connid))
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, msgs)
}

func (h *Handler) handleFriendRequest(c *gin.Context) {
	userid := c.GetUint("userid")
	recidstr := c.Param("recid")
	recid, _ := strconv.ParseUint(recidstr, 10, 32)

	if exists, _ := h.services.Connection.Exists(userid, uint(recid)); exists {
		c.AbortWithStatusJSON(409, map[string]string{"error": "exists"})
		return
	}

	req, err := h.services.MsgReq.Create(userid, uint(recid))
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
		return
	}

	h.wsHandler.SendFriendRequest(req, uint(recid))

	c.Status(201)
}

func (h *Handler) handleReplyToRequest(c *gin.Context) {
	var body dto.MessageRequestResponse
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"error": err.Error()})
		return
	}

	if body.IsAccepted {
		req, err := h.services.MsgReq.GetById(body.ReqID)
		if err != nil {
			c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
			return
		}

		if exists, _ := h.services.Connection.Exists(req.From, req.To); !exists {
			err = h.services.Connection.Create(req.From, req.To)
			if err != nil {
				c.AbortWithStatusJSON(500, map[string]string{"error": err.Error()})
				return
			}
		}
	}

	h.services.MsgReq.Delete(body.ReqID)

	c.Status(201)
}

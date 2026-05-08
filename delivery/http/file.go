package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SaveFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	header := form.File["file"][0]
	file, err := header.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	url, err := h.services.S3.UploadFile(c.GetUint("userid"), header.Filename, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"url": url})
}

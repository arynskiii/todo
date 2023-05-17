package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllItem(c *gin.Context) {

}

func (h *Handler) getItembyId(c *gin.Context) {

}
func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}
